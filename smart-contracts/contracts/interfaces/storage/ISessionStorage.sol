// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface ISessionStorage {
    struct Session {
        address user;
        bytes32 bidId;
        uint256 stake;
        bytes closeoutReceipt;
        // TODO: Use enum?
        uint256 closeoutType;
        // Amount of funds that was already withdrawn by provider (we allow to withdraw for the previous day)
        uint256 providerWithdrawnAmount;
        uint128 openedAt;
        // Expected end time considering the stake provided
        uint128 endsAt;
        uint128 closedAt;
        bool isActive;
    }

    struct OnHold {
        uint256 amount;
        // In epoch seconds. TODO: consider using hours to reduce storage cost
        uint128 releaseAt;
    }

    struct Pool {
        uint256 initialReward;
        uint256 rewardDecrease;
        uint128 payoutStart;
        uint128 decreaseInterval;
    }

    function getSession(bytes32 sessionId_) external view returns (Session memory);

    function getUserSessions(address user, uint256 offset_, uint256 limit_) external view returns (bytes32[] memory);

    function getProviderSessions(
        address provider_,
        uint256 offset_,
        uint256 limit_
    ) external view returns (bytes32[] memory);

    function getModelSessions(
        bytes32 modelId_,
        uint256 offset_,
        uint256 limit_
    ) external view returns (bytes32[] memory);

    function getPools() external view returns (Pool[] memory);

    function getPool(uint256 index_) external view returns (Pool memory);

    function getFundingAccount() external view returns (address);

    function getTotalSessions(address providerAddr_) external view returns (uint256);

    function getProvidersTotalClaimed() external view returns (uint256);

    function getIsProviderApprovalUsed(bytes memory approval_) external view returns (bool);
}
