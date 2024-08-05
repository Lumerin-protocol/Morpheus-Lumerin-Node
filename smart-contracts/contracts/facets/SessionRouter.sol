// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import { ECDSA } from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import { MessageHashUtils } from "@openzeppelin/contracts/utils/cryptography/MessageHashUtils.sol";
import { KeySet, Uint256Set } from "../libraries/KeySet.sol";
import { AppStorage, Session, Bid, OnHold, Pool, Provider, PROVIDER_REWARD_LIMITER_PERIOD, ProviderModelStats, ModelStats } from "../AppStorage.sol";
import { LibOwner } from "../libraries/LibOwner.sol";
import { LibSD } from "../libraries/LibSD.sol";
import { LinearDistributionIntervalDecrease } from "../libraries/LinearDistributionIntervalDecrease.sol";

contract SessionRouter {
  using KeySet for KeySet.Set;
  using Uint256Set for Uint256Set.Set;
  using LibSD for LibSD.SD;

  AppStorage internal s;

  // constants
  uint32 public constant MIN_SESSION_DURATION = 5 minutes;
  uint32 public constant MAX_SESSION_DURATION = 1 days;
  uint32 public constant SIGNATURE_TTL = 10 minutes;

  // events
  event SessionOpened(address indexed userAddress, bytes32 indexed sessionId, address indexed providerId);
  event SessionClosed(address indexed userAddress, bytes32 indexed sessionId, address indexed providerId);

  // errors
  error NotEnoughWithdrawableBalance(); // means that there is not enough funds at all or some funds are still locked
  error WithdrawableBalanceLimitByStakeReached(); // means that user can't withdraw more funds because of the limit which equals to the stake
  error ProviderSignatureMismatch();
  error SignatureExpired();
  error WrongChaidId();
  error DuplicateApproval();
  error ApprovedForAnotherUser(); // means that approval generated for another user address, protection from front-running

  error SessionTooShort();
  error SessionNotFound();
  error SessionAlreadyClosed();
  error SessionNotClosed();

  error BidNotFound();
  error CannotDecodeAbi();

  //===========================
  //         SESSION
  //===========================

  /// @notice returns session by sessionId
  function getSession(bytes32 sessionId) external view returns (Session memory) {
    return s.sessions[s.sessionMap[sessionId]];
  }

  function getActiveSessionsByUser(address user) external view returns (Session[] memory) {
    Uint256Set.Set storage userSessions = s.userActiveSessions[user];
    uint256 size = userSessions.count();
    Session[] memory sessions = new Session[](size);
    for (uint i = 0; i < size; i++) {
      sessions[i] = s.sessions[userSessions.keyAtIndex(i)];
    }
    return sessions;
  }

  function getActiveSessionsByProvider(address provider) external view returns (Session[] memory) {
    Uint256Set.Set storage providerSessions = s.providerActiveSessions[provider];
    uint256 size = providerSessions.count();
    Session[] memory sessions = new Session[](size);
    for (uint i = 0; i < size; i++) {
      sessions[i] = s.sessions[providerSessions.keyAtIndex(i)];
    }
    return sessions;
  }

  function getSessionsByProvider(
    address provider,
    uint256 offset,
    uint8 limit
  ) external view returns (Session[] memory) {
    return paginate(s.providerSessions[provider], offset, limit);
  }

  function getSessionsByUser(address user, uint256 offset, uint8 limit) external view returns (Session[] memory) {
    return paginate(s.userSessions[user], offset, limit);
  }

  function getSessionsByModel(bytes32 modelId, uint256 offset, uint8 limit) external view returns (Session[] memory) {
    return paginate(s.modelSessions[modelId], offset, limit);
  }

  function paginate(uint256[] memory indexes, uint256 offset, uint8 limit) private view returns (Session[] memory) {
    uint256 length = indexes.length;
    if (length < offset) {
      return (new Session[](0));
    }
    uint8 size = offset + limit > length ? uint8(length - offset) : limit;
    Session[] memory sessions = new Session[](size);
    for (uint i = 0; i < size; i++) {
      uint256 index = length - offset - i - 1;
      sessions[i] = s.sessions[indexes[index]];
    }
    return sessions;
  }

  function activeSessionsCount() external view returns (uint256) {
    return s.activeSessionsCount;
  }

  function sessionsCount() external view returns (uint256) {
    return s.sessions.length;
  }

  function openSession(
    uint256 _stake,
    bytes memory providerApproval,
    bytes memory signature
  ) external returns (bytes32 sessionId) {
    // reverts without specific error if cannot decode abi
    (bytes32 bidId, uint256 chainId, address user, uint128 timestampMs) = abi.decode(
      providerApproval,
      (bytes32, uint256, address, uint128)
    );
    if (user != msg.sender) {
      revert ApprovedForAnotherUser();
    }
    if (chainId != block.chainid) {
      revert WrongChaidId();
    }
    if (timestampMs / 1000 < block.timestamp - SIGNATURE_TTL) {
      revert SignatureExpired();
    }

    Bid memory bid = s.bidMap[bidId];
    if (bid.deletedAt != 0 || bid.createdAt == 0) {
      revert BidNotFound();
    }

    if (!isValidReceipt(bid.provider, providerApproval, signature)) {
      revert ProviderSignatureMismatch();
    }

    if (s.approvalMap[providerApproval]) {
      revert DuplicateApproval();
    }
    s.approvalMap[providerApproval] = true;

    uint256 endsAt = whenSessionEnds(_stake, bid.pricePerSecond, block.timestamp);
    if (endsAt - block.timestamp < MIN_SESSION_DURATION) {
      revert SessionTooShort();
    }

    sessionId = keccak256(abi.encodePacked(msg.sender, bid.provider, _stake, s.sessionNonce++));
    s.sessions.push(
      Session({
        id: sessionId,
        user: msg.sender,
        provider: bid.provider,
        modelAgentId: bid.modelAgentId,
        bidID: bidId,
        stake: _stake,
        pricePerSecond: bid.pricePerSecond,
        closeoutReceipt: "",
        closeoutType: 0,
        providerWithdrawnAmount: 0,
        openedAt: uint128(block.timestamp),
        endsAt: endsAt,
        closedAt: 0
      })
    );

    uint256 sessionIndex = s.sessions.length - 1;
    s.sessionMap[sessionId] = sessionIndex;
    s.userSessions[msg.sender].push(sessionIndex);
    s.providerSessions[bid.provider].push(sessionIndex);
    s.modelSessions[bid.modelAgentId].push(sessionIndex);

    s.userActiveSessions[msg.sender].insert(sessionIndex);
    s.providerActiveSessions[bid.provider].insert(sessionIndex);
    s.activeSessionsCount++;

    emit SessionOpened(msg.sender, sessionId, bid.provider);

    // try to use locked stake first, but limit iterations to 20
    // if user has more than 20 onHold entries, they will have to use withdrawUserStake separately
    uint256 removed = _removeUserStake(_stake, 10);
    _stake -= removed;

    s.token.transferFrom(msg.sender, address(this), _stake); // errors with Insufficient Allowance if not approved

    return sessionId;
  }

  function closeSession(bytes memory receiptEncoded, bytes memory signature) external {
    // reverts without specific error if cannot decode abi
    (bytes32 sessionId, uint256 chainId, uint128 timestampMs, uint32 tpsScaled1000, uint32 ttftMs) = abi.decode(
      receiptEncoded,
      (bytes32, uint256, uint128, uint32, uint32)
    );
    if (chainId != block.chainid) {
      revert WrongChaidId();
    }
    if (timestampMs / 1000 < block.timestamp - SIGNATURE_TTL) {
      revert SignatureExpired();
    }

    uint256 sessionIndex = s.sessionMap[sessionId];
    Session storage session = s.sessions[sessionIndex];
    if (session.openedAt == 0) {
      revert SessionNotFound();
    }
    LibOwner._senderOrOwner(session.user);
    if (session.closedAt != 0) {
      revert SessionAlreadyClosed();
    }

    // update indexes
    s.userActiveSessions[session.user].remove(sessionIndex);
    s.providerActiveSessions[session.provider].remove(sessionIndex);
    s.activeSessionsCount--;

    // update session record
    session.closeoutReceipt = receiptEncoded; //TODO: remove that field in favor of tps and ttftMs
    session.closedAt = uint128(block.timestamp);

    // calculate provider withdraw
    uint256 providerWithdraw;
    uint256 startOfToday = startOfTheDay(block.timestamp);
    bool isClosingLate = startOfToday > startOfTheDay(session.endsAt);
    bool noDispute = isValidReceipt(session.provider, receiptEncoded, signature);

    if (noDispute || isClosingLate) {
      // session was closed without dispute or next day after it expected to end
      uint256 duration = minUint256(block.timestamp, session.endsAt) - session.openedAt;
      uint256 cost = duration * session.pricePerSecond;
      providerWithdraw = cost - session.providerWithdrawnAmount;
    } else {
      // session was closed on the same day or earlier with dispute
      // withdraw all funds except for today's session cost
      uint256 durationTillToday = startOfToday - minUint256(session.openedAt, startOfToday);
      uint256 costTillToday = durationTillToday * session.pricePerSecond;
      providerWithdraw = costTillToday - session.providerWithdrawnAmount;
    }

    // updating provider stats
    ProviderModelStats storage prStats = s.stats[session.modelAgentId][session.provider];
    ModelStats storage modelStats = s.modelStats[session.modelAgentId];

    prStats.totalCount++;

    if (noDispute) {
      if (prStats.successCount > 0) {
        // stats for this provider-model pair already contribute to average model stats
        modelStats.tpsScaled1000.remove(int32(prStats.tpsScaled1000.mean), int32(modelStats.count - 1));
        modelStats.ttftMs.remove(int32(prStats.ttftMs.mean), int32(modelStats.count - 1));
      } else {
        // stats for this provider-model pair do not contribute
        modelStats.count++;
      }

      // update provider-model stats
      prStats.successCount++;
      prStats.totalDuration += uint32(session.closedAt - session.openedAt);
      prStats.tpsScaled1000.add(int32(tpsScaled1000), int32(prStats.successCount));
      prStats.ttftMs.add(int32(ttftMs), int32(prStats.successCount));

      // update model stats
      modelStats.totalDuration.add(int32(prStats.totalDuration), int32(modelStats.count));
      modelStats.tpsScaled1000.add(int32(prStats.tpsScaled1000.mean), int32(modelStats.count));
      modelStats.ttftMs.add(int32(prStats.ttftMs.mean), int32(modelStats.count));
    } else {
      session.closeoutType = 1;
    }

    // we have to lock today's stake so the user won't get the reward twice
    uint256 userStakeToLock = 0;
    if (!isClosingLate) {
      // session was closed on the same day
      // lock today's stake
      uint256 todaysDuration = minUint256(session.endsAt, block.timestamp) - maxUint256(startOfToday, session.openedAt);
      uint256 todaysCost = todaysDuration * session.pricePerSecond;
      userStakeToLock = minUint256(session.stake, stipendToStake(todaysCost, startOfToday));
      s.userOnHold[session.user].push(OnHold({ amount: userStakeToLock, releaseAt: uint128(startOfToday + 1 days) }));
    }
    uint256 userWithdraw = session.stake - userStakeToLock;

    emit SessionClosed(session.user, sessionId, session.provider);

    // withdraw provider
    rewardProvider(session, providerWithdraw, false);

    // withdraw user
    s.token.transfer(session.user, userWithdraw);
  }

  // funds related functions

  /// @notice returns total claimanble balance for the provider for particular session
  function getProviderClaimableBalance(bytes32 sessionId) public view returns (uint256) {
    Session memory session = s.sessions[s.sessionMap[sessionId]];
    if (session.openedAt == 0) {
      revert SessionNotFound();
    }
    return _getProviderClaimableBalance(session);
  }

  /// @notice allows provider to claim their funds
  function claimProviderBalance(bytes32 sessionId, uint256 amountToWithdraw) external {
    Session storage session = s.sessions[s.sessionMap[sessionId]];
    if (session.openedAt == 0) {
      revert SessionNotFound();
    }
    LibOwner._senderOrOwner(session.provider);

    uint256 withdrawableAmount = _getProviderClaimableBalance(session);
    if (amountToWithdraw > withdrawableAmount) {
      revert NotEnoughWithdrawableBalance();
    }

    rewardProvider(session, amountToWithdraw, true);
    return;
  }

  function _getProviderClaimableBalance(Session memory session) private view returns (uint256) {
    // if session was closed with no dispute - provider already got all funds
    //
    // if session was closed with dispute   -
    // if session was ended but not closed  -
    // if session was not ended             - provider can claim all funds except for today's session cost

    uint256 claimIntervalEnd = minUint256(minUint256(startOfTheDay(block.timestamp), session.endsAt), session.closedAt);
    uint256 claimableDuration = maxUint256(claimIntervalEnd, session.openedAt) - session.openedAt;
    uint256 totalCost = claimableDuration * session.pricePerSecond;
    uint256 withdrawableAmount = totalCost - session.providerWithdrawnAmount;

    return withdrawableAmount;
  }

  /// @notice deletes session from the history
  function deleteHistory(bytes32 sessionId) external {
    uint256 sessionIndex = s.sessionMap[sessionId];
    Session storage session = s.sessions[sessionIndex];
    LibOwner._senderOrOwner(session.user);
    if (session.closedAt == 0) {
      revert SessionNotClosed();
    }
    session.user = address(0);
  }

  /// @notice checks if receipt is valid
  function isValidReceipt(address signer, bytes memory receipt, bytes memory signature) private pure returns (bool) {
    if (signature.length == 0) {
      return false;
    }
    bytes32 receiptHash = MessageHashUtils.toEthSignedMessageHash(keccak256(receipt));
    return ECDSA.recover(receiptHash, signature) == signer;
  }

  /// @notice returns amount of withdrawable user stake and one on hold
  function withdrawableUserStake(
    address userAddr,
    uint256 offset,
    uint8 limit
  ) external view returns (uint256 avail, uint256 hold) {
    OnHold[] storage onHold = s.userOnHold[userAddr];

    uint256 length = onHold.length;
    if (length <= offset) {
      return (avail, hold);
    }

    uint8 size = offset + limit > length ? uint8(length - offset) : limit;
    for (uint i = offset; i < offset + size; i++) {
      OnHold storage hh = onHold[i];
      if (block.timestamp < hh.releaseAt) {
        hold += hh.amount;
      } else {
        avail += hh.amount;
      }
    }
    return (avail, hold);
  }

  /// @notice withdraws user stake
  /// @param amountToWithdraw amount of funds to withdraw, maxUint256 means all available
  /// @param iterations number of entries to process
  function withdrawUserStake(uint256 amountToWithdraw, uint8 iterations) external {
    // withdraw all available funds if amountToWithdraw is 0
    if (amountToWithdraw == 0) {
      amountToWithdraw = type(uint256).max;
    }

    uint256 removed = _removeUserStake(amountToWithdraw, iterations);
    if (removed < amountToWithdraw) {
      revert NotEnoughWithdrawableBalance();
    }

    s.token.transfer(msg.sender, amountToWithdraw);
  }

  /// @dev removes user stake amount from onHold entries
  function _removeUserStake(uint256 amountToRemove, uint8 iterations) private returns (uint256) {
    uint256 balance = 0;

    OnHold[] storage onHoldEntries = s.userOnHold[msg.sender];
    iterations = iterations > onHoldEntries.length ? uint8(onHoldEntries.length) : iterations;
    uint i = 0;

    // the only loop that is not avoidable
    while (i < onHoldEntries.length && iterations-- > 0) {
      if (block.timestamp >= onHoldEntries[i].releaseAt) {
        balance += onHoldEntries[i].amount;

        if (balance >= amountToRemove) {
          uint256 delta = balance - amountToRemove;
          onHoldEntries[i].amount = delta;
          return amountToRemove;
        }

        // removes entry from array
        if (onHoldEntries.length > 0) {
          onHoldEntries[i] = onHoldEntries[onHoldEntries.length - 1];
        }
        onHoldEntries.pop();
      } else {
        i++;
      }
    }

    return balance;
  }

  /// @notice returns stipend of user based on their stake
  function stakeToStipend(uint256 sessionStake, uint256 timestamp) public view returns (uint256) {
    // inlined getTodaysBudget call to get a better precision
    return (sessionStake * getComputeBalance(timestamp)) / (totalMORSupply(timestamp) * 100);
  }

  /// @notice returns stake of user based on their stipend
  function stipendToStake(uint256 stipend, uint256 timestamp) public view returns (uint256) {
    // inlined getTodaysBudget call to get a better precision
    // return (stipend * totalMORSupply(timestamp)) / getTodaysBudget(timestamp);
    return (stipend * totalMORSupply(timestamp) * 100) / getComputeBalance(timestamp);
  }

  /// @dev make it pure
  function whenSessionEnds(
    uint256 sessionStake,
    uint256 pricePerSecond,
    uint256 openedAt
  ) public view returns (uint256) {
    // if session stake is more than daily price then session will last for its max duration
    uint256 duration = stakeToStipend(sessionStake, openedAt) / pricePerSecond;
    if (duration >= MAX_SESSION_DURATION) {
      return openedAt + MAX_SESSION_DURATION;
    }

    return openedAt + duration;
  }

  /// @notice returns today's budget in MOR
  function getTodaysBudget(uint256 timestamp) public view returns (uint256) {
    return getComputeBalance(timestamp) / 100; // 1% of Compute Balance
  }

  /// @notice returns today's compute balance in MOR
  function getComputeBalance(uint256 timestamp) public view returns (uint256) {
    Pool memory pool = s.pools[3];
    uint256 periodReward = LinearDistributionIntervalDecrease.getPeriodReward(
      pool.initialReward,
      pool.rewardDecrease,
      pool.payoutStart,
      pool.decreaseInterval,
      pool.payoutStart,
      uint128(startOfTheDay(timestamp))
    );

    return periodReward - s.totalClaimed;
  }

  // returns total amount of MOR tokens that were distributed across all pools
  function totalMORSupply(uint256 timestamp) public view returns (uint256) {
    uint256 totalSupply = 0;
    for (uint i = 0; i < s.pools.length; i++) {
      if (i == 3) continue; // skip compute pool (it's calculated separately)
      Pool memory pool = s.pools[i];
      uint256 sup = LinearDistributionIntervalDecrease.getPeriodReward(
        pool.initialReward,
        pool.rewardDecrease,
        pool.payoutStart,
        pool.decreaseInterval,
        pool.payoutStart,
        uint128(startOfTheDay(timestamp))
      );
      totalSupply += sup;
    }
    return totalSupply + s.totalClaimed;
  }

  /////////////////////////
  //   STATS FUNCTIONS   //
  /////////////////////////

  function totalSessions(address providerAddr) private view returns (uint256) {
    return s.providerSessions[providerAddr].length;
  }

  /// @notice sets distibution pool configuration
  /// @dev parameters should be the same as in Ethereum L1 Distribution contract
  /// @dev at address 0x47176B2Af9885dC6C4575d4eFd63895f7Aaa4790
  /// @dev call 'Distribution.pools(3)' where '3' is a poolId
  function setPoolConfig(uint256 index, Pool calldata pool) public {
    LibOwner._onlyOwner();
    s.pools[index] = pool;
  }

  function maybeResetProviderRewardLimiter(Provider storage provider) private {
    if (block.timestamp > provider.limitPeriodEnd) {
      provider.limitPeriodEnd += PROVIDER_REWARD_LIMITER_PERIOD;
      provider.limitPeriodEarned = 0;
    }
  }

  /// @notice sends provider reward considering stake as the limit for the reward
  /// @param session session storage object
  /// @param reward amount of reward to send
  /// @param revertOnReachingLimit if true function will revert if reward is more than stake, otherwise just limit the reward
  function rewardProvider(Session storage session, uint256 reward, bool revertOnReachingLimit) private {
    Provider storage provider = s.providerMap[session.provider];
    maybeResetProviderRewardLimiter(provider);
    uint256 limit = provider.stake - provider.limitPeriodEarned;

    if (reward > limit) {
      if (revertOnReachingLimit) {
        revert WithdrawableBalanceLimitByStakeReached();
      }
      reward = limit;
    }

    session.providerWithdrawnAmount += reward;
    s.totalClaimed += reward;
    provider.limitPeriodEarned += reward;
    s.token.transferFrom(s.fundingAccount, session.provider, reward);
  }

  function startOfTheDay(uint256 timestamp) private pure returns (uint256) {
    return timestamp - (timestamp % 1 days);
  }

  function minUint256(uint256 a, uint256 b) private pure returns (uint256) {
    return a < b ? a : b;
  }

  function maxUint256(uint256 a, uint256 b) private pure returns (uint256) {
    return a > b ? a : b;
  }
}
