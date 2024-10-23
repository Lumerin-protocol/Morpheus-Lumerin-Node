// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface IBidStorage {
    struct Bid {
        address provider;
        bytes32 modelId;
        uint256 pricePerSecond; // Hourly price
        uint256 nonce;
        uint128 createdAt;
        uint128 deletedAt;
    }

    function getBid(bytes32 bidId_) external view returns (Bid memory);

    function getProviderActiveBids(
        address provider_,
        uint256 offset_,
        uint256 limit_
    ) external view returns (bytes32[] memory);

    function getModelActiveBids(
        bytes32 modelId_,
        uint256 offset_,
        uint256 limit_
    ) external view returns (bytes32[] memory);

    function getProviderBids(
        address provider_,
        uint256 offset_,
        uint256 limit_
    ) external view returns (bytes32[] memory);

    function getModelBids(bytes32 modelId_, uint256 offset_, uint256 limit_) external view returns (bytes32[] memory);

    function getToken() external view returns (address);

    function isBidActive(bytes32 bidId_) external view returns (bool);
}
