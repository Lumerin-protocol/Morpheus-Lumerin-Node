// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {Paginator} from "@solarity/solidity-lib/libs/arrays/Paginator.sol";

import {IProviderStorage} from "../../interfaces/storage/IProviderStorage.sol";

contract ProviderStorage is IProviderStorage {
    using Paginator for *;
    using EnumerableSet for EnumerableSet.AddressSet;

    struct PovidersStorage {
        uint256 providerMinimumStake;
        mapping(address => Provider) providers;
        // TODO: move vars below to the graph in the future
        EnumerableSet.AddressSet activeProviders;
    }

    // Reward for this period will be limited by the stake
    uint128 constant PROVIDER_REWARD_LIMITER_PERIOD = 365 days;

    bytes32 public constant PROVIDERS_STORAGE_SLOT = keccak256("diamond.standard.providers.storage");

    /** PUBLIC, GETTERS */
    function getProvider(address provider_) external view returns (Provider memory) {
        return getProvidersStorage().providers[provider_];
    }

    function getProviderMinimumStake() external view returns (uint256) {
        return getProvidersStorage().providerMinimumStake;
    }

    function getActiveProviders(uint256 offset_, uint256 limit_) external view returns (address[] memory) {
         return getProvidersStorage().activeProviders.part(offset_, limit_);
    }

    function getIsProviderActive(address provider_) public view returns (bool) {
        return !getProvidersStorage().providers[provider_].isDeleted;
    }

    /** INTERNAL */
    function getProvidersStorage() internal pure returns (PovidersStorage storage ds) {
        bytes32 slot_ = PROVIDERS_STORAGE_SLOT;

        assembly {
            ds.slot := slot_
        }
    }
}
