// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface IModelStorage {
    struct Model {
        bytes32 ipfsCID; // https://docs.ipfs.tech/concepts/content-addressing/#what-is-a-cid. Up to the model maintainer to keep up to date
        uint256 fee; // The fee is a royalty placeholder that isn't currently used
        uint256 stake;
        address owner;
        string name; // TODO: Limit name length. Up to the model maintainer to keep up to date
        string[] tags; // TODO: Limit tags amount. Up to the model maintainer to keep up to date
        uint128 createdAt;
        bool isDeleted;
    }

    function getModel(bytes32 modelId_) external view returns (Model memory);

    function getModelIds(uint256 offset_, uint256 limit_) external view returns (bytes32[] memory);

    function getModelMinimumStake() external view returns (uint256);

    function getActiveModels(uint256 offset_, uint256 limit_) external view returns (bytes32[] memory);

    function getIsModelActive(bytes32 modelId_) external view returns (bool);
}
