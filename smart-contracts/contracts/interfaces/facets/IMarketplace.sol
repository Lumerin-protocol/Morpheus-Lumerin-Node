// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {IMarketplaceStorage} from "../storage/IMarketplaceStorage.sol";

interface IMarketplace is IMarketplaceStorage {
    event MarketplaceBidPosted(address indexed provider, bytes32 indexed modelId, uint256 nonce);
    event MarketplaceBidDeleted(address indexed provider, bytes32 indexed modelId, uint256 nonce);
    event MaretplaceFeeUpdated(uint256 bidFee);

    error MarketplaceProviderNotFound();
    error MarketplaceModelNotFound();
    error MarketplaceActiveBidNotFound();

    function __Marketplace_init(address token_) external;

    function setMarketplaceBidFee(uint256 bidFee_) external;

    function postModelBid(bytes32 modelId_, uint256 pricePerSecond_) external returns (bytes32);

    function deleteModelBid(bytes32 bidId_) external;

    function withdraw(address recipient_, uint256 amount_) external;

    function getBidId(address provider_, bytes32 modelId_, uint256 nonce_) external view returns (bytes32);

    function getProviderModelId(address provider_, bytes32 modelId_) external view returns (bytes32);
}
