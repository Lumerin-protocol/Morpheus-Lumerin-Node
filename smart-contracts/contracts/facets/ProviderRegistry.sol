// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import { AddressSet } from "../libraries/KeySet.sol";
import { AppStorage, Provider } from "../AppStorage.sol";
import { LibOwner } from "../libraries/LibOwner.sol";

contract ProviderRegistry {
  using AddressSet for AddressSet.Set;

  AppStorage internal s;

  event ProviderRegisteredUpdated(address indexed provider);
  event ProviderDeregistered(address indexed provider);
  event ProviderMinStakeUpdated(uint256 newStake);

  error StakeTooLow();

  /**
   * @notice Get provider data by its address
   * @param addr Address of provider
   * @return Provider struct
   */
  function providerMap(address addr) public view returns (Provider memory) {
    return s.providerMap[addr];
  }

  /**
   * @notice Get address of provider by its index in array
   * @param index Index of array
   * @return Address of corresponding provider
   */
  function providers(uint256 index) public view returns (address) {
    return s.providers[index];
  }

  /**
   * @notice Get addresses of all active provider
   * @return Array of addresses
   */
  function providerGetIds() public view returns (address[] memory) {
    return s.activeProviders.keys();
  }

  /**
   * @notice Get count of active providers
   * @return count Active providers count
   */
  function providerGetCount() public view returns (uint count) {
    return s.activeProviders.count();
  }

  /**
   * @notice Get provider's address and struct by its index in array of provider
   * @param index Array index
   * @return addr Provider's address
   * @return provider Provider struct
   */
  function providerGetByIndex(uint index) public view returns (address addr, Provider memory provider) {
    addr = s.activeProviders.keyAtIndex(index);
    return (addr, s.providerMap[addr]);
  }

  /**
   * @notice Get all addresses of providers and structs corresponding to them
   * @return Array of addresses
   * @return Array of provider structs
   */
  function providerGetAll() public view returns (address[] memory, Provider[] memory) {
    uint256 count = s.activeProviders.count();
    address[] memory _addrs = new address[](count);
    Provider[] memory _providers = new Provider[](count);

    for (uint i = 0; i < count; i++) {
      address addr = s.activeProviders.keyAtIndex(i);
      _addrs[i] = addr;
      _providers[i] = s.providerMap[addr];
    }

    return (_addrs, _providers);
  }

  /**
   * @notice Registers a provider. Only callable for sender or contract owner
   * @dev Emits {ProviderRegisteredUpdated}
   * @param addr Provider address
   * @param addStake Amount of tokens to stake
   * @param endpoint Provider's endpoint (host.com:1234)
   */
  function providerRegister(address addr, uint256 addStake, string memory endpoint) public {
    LibOwner._senderOrOwner(addr);
    Provider memory provider = s.providerMap[addr];
    uint256 newStake = provider.stake + addStake;
    if (newStake < s.providerMinStake) {
      revert StakeTooLow();
    }
    if (provider.createdAt == 0) {
      s.activeProviders.insert(addr);
      s.providers.push(addr);
    }

    s.providerMap[addr] = Provider(endpoint, newStake, uint128(block.timestamp), false);

    emit ProviderRegisteredUpdated(addr);

    s.token.transferFrom(msg.sender, address(this), addStake); // reverts with ERC20InsufficientAllowance
  }

  /**
   * @notice Deregister provider and return staked tokens. Only callable for provider or contract owner
   * @dev Emits {ProviderDeregistered}
   * @param addr Address of provider to deregister
   */
  function providerDeregister(address addr) public {
    LibOwner._senderOrOwner(addr);
    s.activeProviders.remove(addr);

    s.providerMap[addr].isDeleted = true;
    uint256 stake = s.providerMap[addr].stake;

    emit ProviderDeregistered(addr);
    s.token.transfer(addr, stake);
  }

  /**
   * @notice Update minimal amount of token to stake for providers. Only callable for contract owner
   * @dev Emits {ProviderMinStakeUpdated}
   * @param _minStake New minimal stake
   */
  function providerSetMinStake(uint256 _minStake) public {
    LibOwner._onlyOwner();
    s.providerMinStake = _minStake;
    emit ProviderMinStakeUpdated(_minStake);
  }

  /**
   * @notice Check if provider exists
   * @param addr Address of provider to check
   * @return True if exists otherwise false
   */
  function providrerExists(address addr) public view returns (bool) {
    return s.activeProviders.exists(addr);
  }

  /**
   * @notice Get current minimal stake required for provider
   * @return Minimal stake
   */
  function providerMinStake() public view returns (uint256) {
    return s.providerMinStake;
  }
}
