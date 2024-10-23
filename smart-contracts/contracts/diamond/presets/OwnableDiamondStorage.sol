// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {DiamondOwnableStorage} from "@solarity/solidity-lib/diamond/access/ownable/DiamondOwnableStorage.sol";

import {Context} from "@openzeppelin/contracts/utils/Context.sol";

abstract contract OwnableDiamondStorage is DiamondOwnableStorage, Context {
    /**
     * @dev The caller account is not authorized to perform an operation as owner.
     */
    error OwnableUnauthorizedAccount(address account_);

    function _onlyOwner() internal view virtual override {
        if (owner() != _msgSender()) {
            revert OwnableUnauthorizedAccount(_msgSender());
        }
    }

    function _onlyAccount(address sender_) internal view {
        if (_msgSender() != sender_) {
            revert OwnableUnauthorizedAccount(_msgSender());
        }
    }
}
