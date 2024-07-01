//SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { KeySet, AddressSet, Uint256Set } from "./libraries/KeySet.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";


uint128 constant PROVIDER_REWARD_LIMITER_PERIOD = 365 days; // reward for this period will be limited by the stake

struct Provider {
  string endpoint; // example 'domain.com:1234'
  uint256 stake; // stake amount, which also server as a reward limiter
  uint128 createdAt; // timestamp of the registration
  uint128 limitPeriodEnd; // timestamp of the limiter period end
  uint256 limitPeriodEarned; // total earned during the last limiter period
  bool isDeleted;
}

struct Model {
  bytes32 ipfsCID; // https://docs.ipfs.tech/concepts/content-addressing/#what-is-a-cid
  uint256 fee;
  uint256 stake;
  address owner;
  string name; // limit name length
  string[] tags; // TODO: limit tags amount
  uint128 createdAt;
  bool isDeleted;
}

struct Bid {
  address provider;
  bytes32 modelAgentId;
  uint256 pricePerSecond; // hourly price
  uint256 nonce;
  uint128 createdAt;
  uint128 deletedAt;
}

struct Session {
  bytes32 id;
  address user;
  address provider;
  bytes32 modelAgentId;
  bytes32 bidID;
  uint256 stake;
  uint256 pricePerSecond;
  bytes closeoutReceipt;
  uint256 closeoutType;
  // amount of funds that was already withdrawn by provider (we allow to withdraw for the previous day)
  uint256 providerWithdrawnAmount;
  uint256 openedAt;
  uint256 endsAt; // expected end time considering the stake provided
  uint256 closedAt;
}

struct OnHold {
  uint256 amount;
  uint128 releaseAt; // in epoch seconds TODO: consider using hours to reduce storage cost
}

struct Pool {
  uint256 initialReward;
  uint256 rewardDecrease;
  uint128 payoutStart;
  uint128 decreaseInterval;
}

struct AppStorage {
  //
  // PROVIDER storage
  //
  mapping(address => Provider) providerMap; // provider address => Provider
  address[] providers; // all providers ids
  AddressSet.Set activeProviders; // active providers ids
  //
  // MODEL storage
  //
  mapping(bytes32 => Model) modelMap; // modelId => Model
  // mapping(address => bytes32[]) public modelsByOwner; // owner to modelIds
  bytes32[] models; // all model ids
  KeySet.Set activeModels; // active model ids
  //
  // BID storage
  //
  mapping(bytes32 => Bid) bidMap; // bidId = keccak256(provider, modelAgentId, nonce) => bid
  mapping(bytes32 => uint256) providerModelAgentNonce; // keccak256(provider, modelAgentId) => last nonce
  KeySet.Set activeBids; // all active bidIds
  mapping(address => KeySet.Set) providerActiveBids; // provider => active bidIds
  mapping(bytes32 => KeySet.Set) modelAgentActiveBids; // modelAgentId => active bidIds
  mapping(bytes32 => bytes32[]) modelAgentBids; // keccak256(provider, modelAgentId) => all bidIds
  mapping(address => bytes32[]) providerBids; // provider => all bidIds
  //
  // SESSION storage
  //
  // all sessions
  Session[] sessions;
  mapping(bytes32 => uint256) sessionMap; // sessionId => session index
  mapping(address => uint256[]) userSessions; // user address => all session indexes
  mapping(address => uint256[]) providerSessions; // provider address => all session indexes
  mapping(bytes32 => uint256[]) modelSessions; // modelId => all session indexes
  // active sessions
  mapping(address => Uint256Set.Set) userActiveSessions; // user address => active session indexes
  mapping(address => Uint256Set.Set) providerActiveSessions; // provider address => active session indexes
  mapping(address => OnHold[]) userOnHold; // user address => balance
  mapping(bytes => bool) approvalMap; // provider approval => true if approval was already used
  uint64 activeSessionsCount;
  //
  // OTHER
  //
  IERC20 token; // MOR token
  // Number of seconds to delay the stake return when a user closes out a session using a user signed receipt
  int256 stakeDelay;
  address fundingAccount; // account which stores the MOR tokens with infinite allowance for this contract
  uint256 bidFee;
  uint256 feeBalance; // total fees balance of the contract
  uint256 modelMinStake;
  uint256 providerMinStake;
  Pool[] pools; // distribution pools configuration that mirrors L1 contract
  uint256 totalClaimed; // total amount of MOR claimend by providers
}

library LibAppStorage {
  function appStorage() internal pure returns (AppStorage storage ds) {
    assembly {
      ds.slot := 0
    }
  }
}
