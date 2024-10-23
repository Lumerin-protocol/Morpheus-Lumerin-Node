// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IProviderStorage} from "../storage/IProviderStorage.sol";

interface IProviderRegistry is IProviderStorage {
    event ProviderRegistered(address indexed provider);
    event ProviderDeregistered(address indexed provider);
    event ProviderMinimumStakeUpdated(uint256 providerMinimumStake);
    event ProviderWithdrawn(address indexed provider, uint256 amount);
    error ProviderStakeTooLow(uint256 amount, uint256 minAmount);
    error ProviderNotDeregistered();
    error ProviderNoStake();
    error ProviderNothingToWithdraw();
    error ProviderHasActiveBids();
    error ProviderNotFound();
    error ProviderHasAlreadyDeregistered();

    function __ProviderRegistry_init() external;

    /**
     * @notice Sets the minimum stake required for a provider
     * @param providerMinimumStake_ The minimal stake
     */
    function providerSetMinStake(uint256 providerMinimumStake_) external;

    /**
     * @notice Register a provider
     * @param amount_ The amount of stake to add
     * @param endpoint_ The provider endpoint (host.com:1234)
     */
    function providerRegister(uint256 amount_, string calldata endpoint_) external;

    /**
     * @notice Deregister a provider
     */
    function providerDeregister() external;
}
