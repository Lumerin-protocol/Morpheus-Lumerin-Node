// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import { AppStorage, Model } from "../AppStorage.sol";
import { KeySet } from "../libraries/KeySet.sol";
import { LibOwner } from "../libraries/LibOwner.sol";

contract ModelRegistry {
  using KeySet for KeySet.Set;

  AppStorage internal s;

  event ModelRegisteredUpdated(address indexed owner, bytes32 indexed modelId);
  event ModelDeregistered(address indexed owner, bytes32 indexed modelId);
  event ModelMinStakeUpdated(uint256 newStake);

  error ModelNotFound();
  error StakeTooLow();

  /**
   * @notice Get model by its Id
   * @param id Id of model to get
   * @return Struct of model
   */
  function modelMap(bytes32 id) public view returns (Model memory) {
    return s.modelMap[id];
  }

  /**
   * @notice Get model id by key array index
   * @param index Array index
   * @return Corresponding model id
   */
  function models(uint256 index) public view returns (bytes32) {
    return s.models[index];
  }

  /**
   * @notice Get keys of all current active models
   * @return Bytes32 array of ids
   */
  function modelGetIds() public view returns (bytes32[] memory) {
    return s.activeModels.keys();
  }

  /**
   * @notice Get count of active models
   * @return count Active models count
   */
  function modelGetCount() public view returns (uint count) {
    return s.activeModels.count();
  }

  /**
   * @notice Get all ids and models corresponding to them
   * @return Array of model ids
   * @return Array of model structs
   */
  function modelGetAll() public view returns (bytes32[] memory, Model[] memory) {
    uint256 len = s.activeModels.count();
    Model[] memory _models = new Model[](len);
    bytes32[] memory ids = new bytes32[](len);
    for (uint i = 0; i < len; i++) {
      bytes32 id = s.activeModels.keyAtIndex(i);
      ids[i] = id;
      _models[i] = s.modelMap[id];
    }
    return (ids, _models);
  }

  /**
   * @notice Get id and struct of model by key array index
   * @param index Index of key array
   * @return modelId Id of corresponding model
   * @return model Struct of corresponding model
   */
  function modelGetByIndex(uint index) public view returns (bytes32 modelId, Model memory model) {
    modelId = s.activeModels.keyAtIndex(index);
    return (modelId, s.modelMap[modelId]);
  }

  /**
   * @notice Check if model with corresponding id exists
   * @param id Id of model to check
   * @return True if model exists otherwise false
   */
  function modelExists(bytes32 id) public view returns (bool) {
    return s.activeModels.exists(id);
  }

  /**
   * @notice Update model in registry. Only callable for sender or owner
   * @dev Emits {ModelRegisteredUpdated}
   * @param modelId Model's id in bytes32
   * @param ipfsCID IPFS storage CID in bytes32 (https://docs.ipfs.tech/concepts/content-addressing/#what-is-a-cid)
   * @param fee Fee associated with this model
   * @param addStake Amount to stake associated with this agent. Can't be less than `minStake`
   * @param owner Owner of model
   * @param name Model's name
   * @param tags Array of tags
   */
  function modelRegister(
    bytes32 modelId,
    bytes32 ipfsCID,
    uint256 fee,
    uint256 addStake,
    address owner,
    string memory name,
    string[] memory tags
  ) public {
    LibOwner._senderOrOwner(owner);
    Model memory model = s.modelMap[modelId];
    uint256 newStake = model.stake + addStake;
    if (newStake < s.modelMinStake) {
      revert StakeTooLow();
    }
    if (model.stake == 0) {
      s.activeModels.insert(modelId);
      s.models.push(modelId);
    } else {
      LibOwner._senderOrOwner(s.modelMap[modelId].owner);
    }

    s.modelMap[modelId] = Model({
      fee: fee,
      stake: newStake,
      createdAt: uint128(block.timestamp),
      ipfsCID: ipfsCID,
      owner: owner,
      name: name,
      tags: tags,
      isDeleted: false
    });

    emit ModelRegisteredUpdated(owner, modelId);
    s.token.transferFrom(msg.sender, address(this), addStake); // reverts with ERC20InsufficientAllowance()
  }

  /**
   * @notice Deregister model and return staked tokens. Only callable for model or contract owner
   * @dev Emits {ModelDeregistered}
   * @param id Id of model to deregister
   */
  function modelDeregister(bytes32 id) public {
    Model storage model = s.modelMap[id];
    LibOwner._senderOrOwner(model.owner);

    s.activeModels.remove(id); // reverts with KeyNotFound()
    model.isDeleted = true;
    uint256 stake = model.stake;

    emit ModelDeregistered(model.owner, id);
    s.token.transfer(model.owner, stake);
  }

  /**
   * @notice Update minimal stake. Only callable for contract owner
   * @dev Emits {ModelMinStakeUpdated}
   * @param _minStake New minimal stake
   */
  function modelSetMinStake(uint256 _minStake) public {
    LibOwner._onlyOwner();
    s.modelMinStake = _minStake;
    emit ModelMinStakeUpdated(s.modelMinStake);
  }

  /**
   * @notice Get minimal stake for model
   * @return Minimal stake
   */
  function modelMinStake() public view returns (uint256) {
    return s.modelMinStake;
  }

  // TODO: implement these functions
  // function getModelsByOwner(address addr) public view returns (Model[] memory){
  //   Model[] memory _models = new Model[](modelIds.length);
  //   for (uint i = 0; i < modelIds.length; i++) {
  //     if (models[modelIds[i]].owner == addr) {
  //       _models[i] = models[modelIds[i]];
  //     }
  //   }
  //   return _models;
  // }

  // function getModelTypes -- to be implemented when types are defined
  // function getModelsByType
}
