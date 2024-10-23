// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {OwnableDiamondStorage} from "../presets/OwnableDiamondStorage.sol";

import {BidStorage} from "../storages/BidStorage.sol";
import {ProviderStorage} from "../storages/ProviderStorage.sol";

import {IProviderRegistry} from "../../interfaces/facets/IProviderRegistry.sol";

contract ProviderRegistry is IProviderRegistry, OwnableDiamondStorage, ProviderStorage, BidStorage {
    using EnumerableSet for EnumerableSet.AddressSet;
    using SafeERC20 for IERC20;

    function __ProviderRegistry_init() external initializer(PROVIDERS_STORAGE_SLOT) {}

    function providerSetMinStake(uint256 providerMinimumStake_) external onlyOwner {
        PovidersStorage storage providersStorage = getProvidersStorage();
        providersStorage.providerMinimumStake = providerMinimumStake_;

        emit ProviderMinimumStakeUpdated(providerMinimumStake_);
    }

    function providerRegister(uint256 amount_, string calldata endpoint_) external {
        BidsStorage storage bidsStorage = getBidsStorage();

        if (amount_ > 0) {
           IERC20(bidsStorage.token).safeTransferFrom(_msgSender(), address(this), amount_);
        }

        PovidersStorage storage providersStorage = getProvidersStorage();
        Provider storage provider = providersStorage.providers[_msgSender()];

        uint256 newStake_ = provider.stake + amount_;
        uint256 minStake_ = providersStorage.providerMinimumStake;
        if (newStake_ < minStake_) {
            revert ProviderStakeTooLow(newStake_, minStake_);
        }

        if (provider.createdAt == 0) {
            provider.endpoint = endpoint_;
            provider.createdAt = uint128(block.timestamp);
            provider.limitPeriodEnd = uint128(block.timestamp) + PROVIDER_REWARD_LIMITER_PERIOD;
        } else if (provider.isDeleted) {
            provider.isDeleted = false;
        }

        provider.endpoint = endpoint_;
        provider.stake = newStake_;

        providersStorage.activeProviders.add(_msgSender());

        emit ProviderRegistered(_msgSender());
    }

    function providerDeregister() external {
        PovidersStorage storage providersStorage = getProvidersStorage();
        Provider storage provider = providersStorage.providers[_msgSender()];

        if (provider.createdAt == 0) {
            revert ProviderNotFound();
        }
        if (!isProviderActiveBidsEmpty(_msgSender())) {
            revert ProviderHasActiveBids();
        }
        if (provider.isDeleted) {
            revert ProviderHasAlreadyDeregistered();
        }

        uint256 withdrawAmount_ = _getWithdrawAmount(provider);

        provider.stake -= withdrawAmount_;
        provider.isDeleted = true;

        providersStorage.activeProviders.remove(_msgSender());

        if (withdrawAmount_ > 0) {
            BidsStorage storage bidsStorage = getBidsStorage();
            IERC20(bidsStorage.token).safeTransfer(_msgSender(), withdrawAmount_);
        }

        emit ProviderDeregistered(_msgSender());
    }

    // /**
    //  *
    //  * @notice Withdraws stake from a provider after it has been deregistered
    //  * Allows to withdraw the stake after provider reward period has ended
    //  */
    // function providerWithdrawStake() external {
    //     Provider storage provider = providers(_msgSender());

    //     if (!provider.isDeleted) {
    //         revert ProviderNotDeregistered();
    //     }
    //     if (provider.stake == 0) {
    //         revert ProviderNoStake();
    //     }

    //     uint256 withdrawAmount_ = _getWithdrawAmount(provider);
    //     if (withdrawAmount_ == 0) {
    //         revert ProviderNothingToWithdraw();
    //     }

    //     provider.stake -= withdrawAmount_;
    //     getToken().safeTransfer(_msgSender(), withdrawAmount_);

    //     emit ProviderWithdrawn(_msgSender(), withdrawAmount_);
    // }

    /**
     * @notice Returns the withdrawable stake for a provider
     * @dev If the provider already earned this period then withdrawable stake
     * is limited by the amount earning that remains in the current period.
     * It is done to prevent the provider from withdrawing and then staking
     * again from a different address, which bypasses the limitation.
     */
    function _getWithdrawAmount(Provider storage provider) private view returns (uint256) {
        if (block.timestamp > provider.limitPeriodEnd) {
            return provider.stake;
        }

        return provider.stake - provider.limitPeriodEarned;
    }
}
