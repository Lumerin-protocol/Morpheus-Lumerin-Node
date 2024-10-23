// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract LumerinToken is ERC20 {
    uint256 public constant INITIAL_SUPPLUY = 1_000_000_000 * (10 ** 8);

    constructor(string memory name_, string memory symbol_) ERC20(name_, symbol_) {
        _mint(_msgSender(), INITIAL_SUPPLUY);
    }

    function decimals() public pure override returns (uint8) {
        return 8;
    }
}
