// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

import {LinearDistributionIntervalDecrease} from "morpheus-smart-contracts/contracts/libs/LinearDistributionIntervalDecrease.sol";

import {OwnableDiamondStorage} from "../presets/OwnableDiamondStorage.sol";

import {BidStorage, EnumerableSet} from "../storages/BidStorage.sol";
import {StatsStorage} from "../storages/StatsStorage.sol";
import {SessionStorage} from "../storages/SessionStorage.sol";
import {ProviderStorage} from "../storages/ProviderStorage.sol";

import {LibSD} from "../../libs/LibSD.sol";

import {ISessionRouter} from "../../interfaces/facets/ISessionRouter.sol";

contract SessionRouter is
    ISessionRouter,
    OwnableDiamondStorage,
    SessionStorage,
    ProviderStorage,
    BidStorage,
    StatsStorage
{
    using Math for *;
    using LibSD for LibSD.SD;
    using SafeERC20 for IERC20;
    using EnumerableSet for EnumerableSet.Bytes32Set;

    function __SessionRouter_init(
        address fundingAccount_,
        Pool[] calldata pools_
    ) external initializer(SESSIONS_STORAGE_SLOT) {
        SessionsStorage storage sessionsStorage = getSessionsStorage();

        sessionsStorage.fundingAccount = fundingAccount_;
        for (uint256 i = 0; i < pools_.length; i++) {
            sessionsStorage.pools.push(pools_[i]);
        }
    }

    ////////////////////////////
    ///   CONTRACT CONFIGS   ///
    ////////////////////////////
    /**
     * @notice Sets distibution pool configuration
     * @dev parameters should be the same as in Ethereum L1 Distribution contract
     * @dev at address 0x47176B2Af9885dC6C4575d4eFd63895f7Aaa4790
     * @dev call 'Distribution.pools(3)' where '3' is a poolId
     */
    function setPoolConfig(uint256 index_, Pool calldata pool_) external onlyOwner {
        if (index_ >= getSessionsStorage().pools.length) {
            revert SessionPoolIndexOutOfBounds();
        }

        getSessionsStorage().pools[index_] = pool_;
    }

    ////////////////////////
    ///   OPEN SESSION   ///
    ////////////////////////
    function openSession(
        uint256 amount_,
        bytes calldata approvalEncoded_,
        bytes calldata signature_
    ) external returns (bytes32) {
        bytes32 bidId_ = _extractProviderApproval(approvalEncoded_);
        if (!isBidActive(bidId_)) {
            revert SessionBidNotFound();
        }

        BidsStorage storage bidsStorage = getBidsStorage();
        SessionsStorage storage sessionsStorage = getSessionsStorage();

        Bid storage bid = bidsStorage.bids[bidId_];
        if (!_isValidProviderReceipt(bid.provider, approvalEncoded_, signature_)) {
            revert SessionProviderSignatureMismatch();
        }
        if (sessionsStorage.isProviderApprovalUsed[approvalEncoded_]) {
            revert SessionDuplicateApproval();
        }

        uint128 endsAt_ = getSessionEnd(amount_, bid.pricePerSecond, uint128(block.timestamp));
        bytes32 sessionId_ = getSessionId(_msgSender(), bid.provider, bidId_, sessionsStorage.sessionNonce++);

        if (endsAt_ - block.timestamp < MIN_SESSION_DURATION) {
            revert SessionTooShort();
        }

        Session storage session = sessionsStorage.sessions[sessionId_];

        session.user = _msgSender();
        session.stake = amount_;
        session.bidId = bidId_;
        session.openedAt = uint128(block.timestamp);
        session.endsAt = endsAt_;
        session.isActive = true;

        sessionsStorage.userSessions[_msgSender()].add(sessionId_);
        sessionsStorage.providerSessions[bid.provider].add(sessionId_);
        sessionsStorage.modelSessions[bid.modelId].add(sessionId_);

        sessionsStorage.isProviderApprovalUsed[approvalEncoded_] = true;

        IERC20(bidsStorage.token).safeTransferFrom(_msgSender(), address(this), amount_);

        emit SessionOpened(_msgSender(), sessionId_, bid.provider);

        return sessionId_;
    }

    function getSessionId(
        address user_,
        address provider_,
        bytes32 bidId_,
        uint256 sessionNonce_
    ) public pure returns (bytes32) {
        return keccak256(abi.encodePacked(user_, provider_, bidId_, sessionNonce_));
    }

    function getSessionEnd(uint256 amount_, uint256 pricePerSecond_, uint128 openedAt_) public view returns (uint128) {
        uint128 duration_ = uint128(stakeToStipend(amount_, openedAt_) / pricePerSecond_);

        if (duration_ > MAX_SESSION_DURATION) {
            duration_ = MAX_SESSION_DURATION;
        }

        return openedAt_ + duration_;
    }

    /**
     * @dev Returns stipend of user based on their stake
     * (User session stake amount / MOR Supply without Compute) * (MOR Compute Supply / 100)
     */
    function stakeToStipend(uint256 amount_, uint128 timestamp_) public view returns (uint256) {
        uint256 totalMorSupply_ = totalMORSupply(timestamp_);
        if (totalMorSupply_ == 0) {
            return 0;
        }

        return (amount_ * getComputeBalance(timestamp_)) / (totalMORSupply(timestamp_) * 100);
    }

    function _extractProviderApproval(bytes calldata providerApproval_) private view returns (bytes32) {
        (bytes32 bidId_, uint256 chainId_, address user_, uint128 timestamp_) = abi.decode(
            providerApproval_,
            (bytes32, uint256, address, uint128)
        );

        if (user_ != _msgSender()) {
            revert SessionApprovedForAnotherUser();
        }
        if (chainId_ != block.chainid) {
            revert SesssionApprovedForAnotherChainId();
        }
        if (block.timestamp > timestamp_ + SIGNATURE_TTL) {
            revert SesssionApproveExpired();
        }

        return bidId_;
    }

    ///////////////////////////////////
    ///   CLOSE SESSION, WITHDRAW   ///
    ///////////////////////////////////
    function closeSession(bytes calldata receiptEncoded_, bytes calldata signature_) external {
        (bytes32 sessionId_, uint32 tpsScaled1000_, uint32 ttftMs_) = _extractProviderReceipt(receiptEncoded_);

        Session storage session = getSessionsStorage().sessions[sessionId_];
        Bid storage bid = getBidsStorage().bids[session.bidId];

        _onlyAccount(session.user);
        if (session.closedAt != 0) {
            revert SessionAlreadyClosed();
        }

        session.isActive = false;
        session.closeoutReceipt = receiptEncoded_; // TODO: Remove that field in favor of tps and ttftMs
        session.closedAt = uint128(block.timestamp);

        //// PROVIDER REWARDS
        uint128 startOfToday_ = startOfTheDay(uint128(block.timestamp));
        // The session should be closed the day after the end of the session to prevent provider rewards locking
        bool isClosingLate_ = startOfToday_ > startOfTheDay(session.endsAt);
        bool noDispute_ = _isValidProviderReceipt(bid.provider, receiptEncoded_, signature_);

        uint128 duration_;
        if (noDispute_ || isClosingLate_) {
            // Session was closed without dispute or next day after it expected to end
            duration_ = uint128(session.endsAt.min(session.closedAt)) - session.openedAt;
        } else {
            // Session was closed on the same day or earlier with dispute
            // withdraw all funds except for today's session cost
            duration_ = startOfToday_ - uint128(session.openedAt.min(uint256(startOfToday_)));
        }
        uint256 providerAmountToWithdraw_ = (duration_ * bid.pricePerSecond) - session.providerWithdrawnAmount;
        _claimForProvider(session, providerAmountToWithdraw_);
        //// END

        //// USER REWARDS
        // We have to lock today's stake so the user won't get the reward twice
        uint256 userStakeToLock_ = 0;
        if (!isClosingLate_) {
            // Session was closed on the same day, lock today's stake
            uint256 userDuration_ = session.endsAt.min(session.closedAt) - session.openedAt.max(startOfToday_);
            uint256 userInitialLock_ = userDuration_ * bid.pricePerSecond;
            userStakeToLock_ = session.stake.min(stipendToStake(userInitialLock_, startOfToday_));

            getSessionsStorage().userStakesOnHold[session.user].push(OnHold(userStakeToLock_, uint128(startOfToday_ + 1 days)));
        }
        uint256 userAmountToWithdraw_ = session.stake - userStakeToLock_;
        IERC20(getBidsStorage().token).safeTransfer(session.user, userAmountToWithdraw_);
        //// END

        //// STATS
        ProviderModelStats storage prStats = providerModelStats(bid.modelId, bid.provider);
        ModelStats storage modelStats = modelStats(bid.modelId);

        prStats.totalCount++;

        if (noDispute_) {
            if (prStats.successCount > 0) {
                // Stats for this provider-model pair already contribute to average model stats
                modelStats.tpsScaled1000.remove(int32(prStats.tpsScaled1000.mean), int32(modelStats.count - 1));
                modelStats.ttftMs.remove(int32(prStats.ttftMs.mean), int32(modelStats.count - 1));
            } else {
                // Stats for this provider-model pair do not contribute
                modelStats.count++;
            }

            // Update provider model stats
            prStats.successCount++;
            prStats.totalDuration += uint32(session.closedAt - session.openedAt);
            prStats.tpsScaled1000.add(int32(tpsScaled1000_), int32(prStats.successCount));
            prStats.ttftMs.add(int32(ttftMs_), int32(prStats.successCount));

            // Update model stats
            modelStats.totalDuration.add(int32(prStats.totalDuration), int32(modelStats.count));
            modelStats.tpsScaled1000.add(int32(prStats.tpsScaled1000.mean), int32(modelStats.count));
            modelStats.ttftMs.add(int32(prStats.ttftMs.mean), int32(modelStats.count));
        } else {
            session.closeoutType = 1;
        }
        //// END

        emit SessionClosed(session.user, sessionId_, bid.provider);
    }

    function _extractProviderReceipt(bytes calldata receiptEncoded_) private view returns (bytes32, uint32, uint32) {
        (bytes32 sessionId_, uint256 chainId_, uint128 timestamp_, uint32 tpsScaled1000_, uint32 ttftMs_) = abi.decode(
            receiptEncoded_,
            (bytes32, uint256, uint128, uint32, uint32)
        );

        if (chainId_ != block.chainid) {
            revert SesssionReceiptForAnotherChainId();
        }
        if (block.timestamp > timestamp_ + SIGNATURE_TTL) {
            revert SesssionReceiptExpired();
        }

        return (sessionId_, tpsScaled1000_, ttftMs_);
    }

    /**
     * @dev Allows providers to receive their funds after the end or closure of the session
     */
    function claimForProvider(bytes32 sessionId_) external {
        Session storage session = getSessionsStorage().sessions[sessionId_];
        Bid storage bid = getBidsStorage().bids[session.bidId];

        _onlyAccount(bid.provider);

        uint256 sessionEnd_ = session.closedAt == 0 ? session.endsAt : session.closedAt;
        if (sessionEnd_ > block.timestamp) {
            revert SessionNotEndedOrNotExist();
        }

        uint256 amount_ = (sessionEnd_ - session.openedAt) * bid.pricePerSecond - session.providerWithdrawnAmount;

        _claimForProvider(session, amount_);
    }

    /**
     * @dev Sends provider reward considering stake as the limit for the reward
     * @param session Storage session object
     * @param amount_ Amount of reward to send
     */
    function _claimForProvider(Session storage session, uint256 amount_) private {
        SessionsStorage storage sessionsStorage = getSessionsStorage();
        BidsStorage storage bidsStorage = getBidsStorage();
        PovidersStorage storage providersStorage = getProvidersStorage();

        Bid storage bid = bidsStorage.bids[session.bidId];
        Provider storage provider = providersStorage.providers[bid.provider];

        if (block.timestamp > provider.limitPeriodEnd) {
            provider.limitPeriodEnd = uint128(block.timestamp) + PROVIDER_REWARD_LIMITER_PERIOD;
            provider.limitPeriodEarned = 0;
        }

        uint256 providerClaimLimit_ = provider.stake - provider.limitPeriodEarned;

        amount_ = amount_.min(providerClaimLimit_);

        if (amount_ == 0) {
            return;
        }

        session.providerWithdrawnAmount += amount_;
        provider.limitPeriodEarned += amount_;
        sessionsStorage.providersTotalClaimed += amount_;

        IERC20(bidsStorage.token).safeTransferFrom(sessionsStorage.fundingAccount, bid.provider, amount_);
    }

    /**
     * @notice Returns stake of user based on their stipend
     */
    function stipendToStake(uint256 stipend_, uint128 timestamp_) public view returns (uint256) {
        return (stipend_ * totalMORSupply(timestamp_) * 100) / getComputeBalance(timestamp_);
    }

    function getUserStakesOnHold(
        address user_,
        uint8 iterations_
    ) external view returns (uint256 available_, uint256 hold_) {
        OnHold[] memory onHold = getSessionsStorage().userStakesOnHold[user_];
        iterations_ = iterations_ > onHold.length ? uint8(onHold.length) : iterations_;

        for (uint256 i = 0; i < onHold.length; i++) {
            uint256 amount = onHold[i].amount;

            if (block.timestamp < onHold[i].releaseAt) {
                hold_ += amount;
            } else {
                available_ += amount;
            }
        }
    }

    function withdrawUserStakes(uint8 iterations_) external {
        uint256 amount_ = 0;

        OnHold[] storage onHoldEntries = getSessionsStorage().userStakesOnHold[_msgSender()];
        uint8 i = iterations_ >= onHoldEntries.length ? uint8(onHoldEntries.length) : iterations_;
        i--;

        while (i >= 0) {
            if (block.timestamp < onHoldEntries[i].releaseAt) {
                if (i == 0) break;
                i--;

                continue;
            }

            amount_ += onHoldEntries[i].amount;
            onHoldEntries.pop();

            if (i == 0) break;
            i--;
        }

        if (amount_ == 0) {
            revert SessionUserAmountToWithdrawIsZero();
        }

        IERC20(getBidsStorage().token).safeTransfer(_msgSender(), amount_);

        emit UserWithdrawn(_msgSender(), amount_);
    }

    ////////////////////////
    ///   GLOBAL PUBLIC  ///
    ////////////////////////

    /**
     * @dev Returns today's budget in MOR. 1%
     */
    function getTodaysBudget(uint128 timestamp_) external view returns (uint256) {
        return getComputeBalance(timestamp_) / 100;
    }

    /**
     * @dev Returns today's compute balance in MOR without claimed amount
     */
    function getComputeBalance(uint128 timestamp_) public view returns (uint256) {
        SessionsStorage storage sessionsStorage = getSessionsStorage();
        Pool storage pool = sessionsStorage.pools[COMPUTE_POOL_INDEX];

        uint256 periodReward = LinearDistributionIntervalDecrease.getPeriodReward(
            pool.initialReward,
            pool.rewardDecrease,
            pool.payoutStart,
            pool.decreaseInterval,
            pool.payoutStart,
            uint128(startOfTheDay(timestamp_))
        );

        return periodReward - sessionsStorage.providersTotalClaimed;
    }

    /**
     * @dev Total amount of MOR tokens that were distributed across all pools
     * without compute pool rewards and with compute claimed rewards
     */
    function totalMORSupply(uint128 timestamp_) public view returns (uint256) {
        uint256 startOfTheDay_ = startOfTheDay(timestamp_);
        uint256 totalSupply_ = 0;

        SessionsStorage storage sessionsStorage = getSessionsStorage();
        Pool[] storage pools = sessionsStorage.pools;

        for (uint256 i = 0; i < pools.length; i++) {
            if (i == COMPUTE_POOL_INDEX) continue;

            totalSupply_ += LinearDistributionIntervalDecrease.getPeriodReward(
                pools[i].initialReward,
                pools[i].rewardDecrease,
                pools[i].payoutStart,
                pools[i].decreaseInterval,
                pools[i].payoutStart,
                uint128(startOfTheDay_)
            );
        }

        return totalSupply_ + sessionsStorage.providersTotalClaimed;
    }

    function startOfTheDay(uint128 timestamp_) public pure returns (uint128) {
        return timestamp_ - (timestamp_ % 1 days);
    }

    //////////////////////////
    ///   GLOBAL PRIVATE   ///
    //////////////////////////

    function _isValidProviderReceipt(
        address provider_,
        bytes calldata receipt_,
        bytes calldata signature_
    ) private pure returns (bool) {
        bytes32 receiptHash_ = ECDSA.toEthSignedMessageHash(keccak256(receipt_));

        return ECDSA.recover(receiptHash_, signature_) == provider_;
    }
}
