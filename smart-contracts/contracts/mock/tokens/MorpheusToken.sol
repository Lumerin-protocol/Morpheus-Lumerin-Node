// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract MorpheusToken is ERC20 {
    // set the initial supply to 42 million like in whitepaper
    uint256 public constant INITIAL_SUPPLUY = 42_000_000 ether;

    constructor() ERC20("Morpheus dev", "MOR") {
        _mint(_msgSender(), INITIAL_SUPPLUY);
    }
}
