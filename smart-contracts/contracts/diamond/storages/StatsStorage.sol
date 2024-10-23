// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IStatsStorage} from "../../interfaces/storage/IStatsStorage.sol";

contract StatsStorage is IStatsStorage {
    struct STTSStorage {
        mapping(bytes32 => mapping(address => ProviderModelStats)) providerModelStats; // modelId => provider => stats
        mapping(bytes32 => ModelStats) modelStats;
    }

    bytes32 public constant STATS_STORAGE_SLOT = keccak256("diamond.stats.storage");

    /** PUBLIC, GETTERS */
    function getProviderModelStats(
        bytes32 modelId_,
        address provider_
    ) external view returns (ProviderModelStats memory) {
        return _getStatsStorage().providerModelStats[modelId_][provider_];
    }

    function getModelStats(bytes32 modelId_) external view returns (ModelStats memory) {
        return _getStatsStorage().modelStats[modelId_];
    }

    /** INTERNAL, GETTERS */
    function modelStats(bytes32 modelId) internal view returns (ModelStats storage) {
        return _getStatsStorage().modelStats[modelId];
    }

    function providerModelStats(bytes32 modelId, address provider) internal view returns (ProviderModelStats storage) {
        return _getStatsStorage().providerModelStats[modelId][provider];
    }

    /** PRIVATE */
    function _getStatsStorage() private pure returns (STTSStorage storage ds) {
        bytes32 slot_ = STATS_STORAGE_SLOT;

        assembly {
            ds.slot := slot_
        }
    }
}
