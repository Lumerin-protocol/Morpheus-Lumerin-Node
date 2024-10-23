// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {Context} from "@openzeppelin/contracts/utils/Context.sol";

import {OwnableDiamond} from "@solarity/solidity-lib/diamond/presets/OwnableDiamond.sol";
import {DiamondOwnableStorage, OwnableDiamondStorage} from "./presets/OwnableDiamondStorage.sol";

contract LumerinDiamond is OwnableDiamond, OwnableDiamondStorage {
    function __LumerinDiamond_init() external initializer(DIAMOND_OWNABLE_STORAGE_SLOT) {
        __DiamondOwnable_init();
    }

    function _onlyOwner() internal view virtual override(DiamondOwnableStorage, OwnableDiamondStorage) {
        OwnableDiamondStorage._onlyOwner();
    }
}
