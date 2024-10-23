// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface IProviderStorage {
    struct Provider {
        string endpoint; // Example 'domain.com:1234'
        uint256 stake; // Stake amount, which also server as a reward limiter
        uint128 createdAt; // Timestamp of the registration
        uint128 limitPeriodEnd; // Timestamp of the limiter period end
        uint256 limitPeriodEarned; // Total earned during the last limiter period
        bool isDeleted;
    }

    function getProvider(address provider_) external view returns (Provider memory);

    function getProviderMinimumStake() external view returns (uint256);

    function getActiveProviders(uint256 offset_, uint256 limit_) external view returns (address[] memory);

    function getIsProviderActive(address provider_) external view returns (bool);
}
