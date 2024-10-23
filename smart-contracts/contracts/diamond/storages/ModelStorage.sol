// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {Paginator} from "@solarity/solidity-lib/libs/arrays/Paginator.sol";

import {IModelStorage} from "../../interfaces/storage/IModelStorage.sol";

contract ModelStorage is IModelStorage {
    using Paginator for *;
    using EnumerableSet for EnumerableSet.Bytes32Set;

    struct ModelsStorage {
        uint256 modelMinimumStake;
        EnumerableSet.Bytes32Set modelIds;
        mapping(bytes32 modelId => Model) models;
         // TODO: move vars below to the graph in the future
        EnumerableSet.Bytes32Set activeModels;
    }

    bytes32 public constant MODELS_STORAGE_SLOT = keccak256("diamond.standard.models.storage");

    /** PUBLIC, GETTERS */
    function getModel(bytes32 modelId_) external view returns (Model memory) {
        return getModelsStorage().models[modelId_];
    }

    function getModelIds(uint256 offset_, uint256 limit_) external view returns (bytes32[] memory) {
        return getModelsStorage().modelIds.part(offset_, limit_);
    }

    function getModelMinimumStake() external view returns (uint256) {
        return getModelsStorage().modelMinimumStake;
    }

    function getActiveModels(uint256 offset_, uint256 limit_) external view returns (bytes32[] memory) {
         return getModelsStorage().activeModels.part(offset_, limit_);
    }

    function getIsModelActive(bytes32 modelId_) public view returns (bool) {
        return !getModelsStorage().models[modelId_].isDeleted;
    }

    /** INTERNAL */
    function getModelsStorage() internal pure returns (ModelsStorage storage ds) {
        bytes32 slot_ = MODELS_STORAGE_SLOT;

        assembly {
            ds.slot := slot_
        }
    }
}
