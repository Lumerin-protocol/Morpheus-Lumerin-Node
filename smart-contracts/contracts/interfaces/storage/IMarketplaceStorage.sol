// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface IMarketplaceStorage {
    function getBidFee() external view returns (uint256);

    function getFeeBalance() external view returns (uint256);
}
