// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {LibSD} from "../../libs/LibSD.sol";

interface IStatsStorage {
    struct ModelStats {
        LibSD.SD tpsScaled1000;
        LibSD.SD ttftMs;
        LibSD.SD totalDuration;
        uint32 count;
    }

    struct ProviderModelStats {
        LibSD.SD tpsScaled1000; // Tokens per second running average
        LibSD.SD ttftMs; // Time to first token running average in milliseconds
        uint32 totalDuration; // Total duration of sessions
        uint32 successCount; // Number of observations
        uint32 totalCount;
        // TODO: consider adding SD with weldford algorithm
    }

    function getProviderModelStats(
        bytes32 modelId_,
        address provider_
    ) external view returns (ProviderModelStats memory);

    function getModelStats(bytes32 modelId_) external view returns (ModelStats memory);
}
