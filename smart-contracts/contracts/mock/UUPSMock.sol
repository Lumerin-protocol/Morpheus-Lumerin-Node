// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

contract UUPSMock is UUPSUpgradeable {
    function version() external pure returns (uint256) {
        return 999;
    }

    function _authorizeUpgrade(address) internal view override {}
}
