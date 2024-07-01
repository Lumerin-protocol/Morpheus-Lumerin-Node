// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.24;

import { AddressSet } from "../libraries/KeySet.sol";
import { AppStorage, Provider, PROVIDER_REWARD_LIMITER_PERIOD } from "../AppStorage.sol";
import { LibOwner } from "../libraries/LibOwner.sol";

contract ProviderRegistry {
  using AddressSet for AddressSet.Set;

  AppStorage internal s;

  event ProviderRegisteredUpdated(address indexed provider);
  event ProviderDeregistered(address indexed provider);
  event ProviderMinStakeUpdated(uint256 newStake);

  error StakeTooLow();
  error ErrProviderNotDeleted();
  error ErrNoStake();
  error ErrNoWithdrawableStake();

  /// @notice Returns provider struct by address
  function providerMap(address addr) external view returns (Provider memory) {
    return s.providerMap[addr];
  }

  /// @notice Returns provider address by index
  function providers(uint256 index) external view returns (address) {
    return s.providers[index];
  }

  /// @notice Returns active (undeleted) provider IDs
  function providerGetIds() external view returns (address[] memory) {
    return s.activeProviders.keys();
  }

  /// @notice Returns count of active providers
  function providerGetCount() external view returns (uint count) {
    return s.activeProviders.count();
  }

  /// @notice Returns provider by index
  function providerGetByIndex(uint index) external view returns (address addr, Provider memory provider) {
    addr = s.activeProviders.keyAtIndex(index);
    return (addr, s.providerMap[addr]);
  }

  /// @notice Returns all providers
  function providerGetAll() external view returns (address[] memory, Provider[] memory) {
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

  /// @notice Registers a provider
  /// @param   addr      provider address
  /// @param   addStake  amount of stake to add
  /// @param   endpoint  provider endpoint (host.com:1234)
  function providerRegister(address addr, uint256 addStake, string memory endpoint) external {
    LibOwner._senderOrOwner(addr);
    Provider memory provider = s.providerMap[addr];
    uint256 newStake = provider.stake + addStake;
    if (newStake < s.providerMinStake) {
      revert StakeTooLow();
    }
    // if we add stake to an existing provider the limiter period is not reset
    uint128 createdAt = provider.createdAt;
    uint128 periodEnd = provider.limitPeriodEnd;
    if (provider.createdAt == 0) {
      s.activeProviders.insert(addr);
      s.providers.push(addr);
      createdAt = uint128(block.timestamp);
      periodEnd = createdAt + PROVIDER_REWARD_LIMITER_PERIOD;
    }

    s.providerMap[addr] = Provider(endpoint, newStake, createdAt, periodEnd, provider.limitPeriodEarned, false);

    emit ProviderRegisteredUpdated(addr);

    s.token.transferFrom(msg.sender, address(this), addStake); // reverts with ERC20InsufficientAllowance
  }

  /// @notice Deregisters a provider
  function providerDeregister(address addr) external {
    LibOwner._senderOrOwner(addr);
    s.activeProviders.remove(addr);

    emit ProviderDeregistered(addr);

    Provider storage p = s.providerMap[addr];
    uint256 withdrawable = getWithdrawableStake(p);
    p.stake -= withdrawable;
    p.isDeleted = true;
    s.token.transfer(addr, withdrawable);
  }

  /// @notice Withdraws stake from a provider after it has been deregistered
  ///         Allows to withdraw the stake after provider reward period has ended
  function providerWithdrawStake(address addr) external {
    Provider storage p = s.providerMap[addr];
    if (!p.isDeleted) {
      revert ErrProviderNotDeleted();
    }

    if (p.stake == 0){
      revert ErrNoStake();
    }

    uint256 withdrawable = getWithdrawableStake(p);
    if (withdrawable == 0) {
      revert ErrNoWithdrawableStake();
    }

    p.stake -= withdrawable;
    s.token.transfer(addr, withdrawable);
  }

  /// @notice Returns the withdrawable stake for a provider
  /// @dev    If the provider already earned this period then withdrawable stake 
  ///         is limited by the amount earning that remains in the current period.
  ///         It is done to prevent the provider from withdrawing and then staking
  ///         again from a different address, which bypasses the limitation. 
  function getWithdrawableStake(Provider memory p) private view returns (uint256) {
    if (uint128(block.timestamp) > p.limitPeriodEnd) {
      return p.stake;
    }
    return p.stake - p.limitPeriodEarned;
  }

  /// @notice Sets the minimum stake required for a provider
  function providerSetMinStake(uint256 _minStake) external {
    LibOwner._onlyOwner();
    s.providerMinStake = _minStake;
    emit ProviderMinStakeUpdated(_minStake);
  }

  /// @notice Checks if a provider exists (is active / not deleted)
  function providerExists(address addr) external view returns (bool) {
    return s.activeProviders.exists(addr);
  }

  /// @notice Returns the minimum stake required for a provider
  function providerMinStake() external view returns (uint256) {
    return s.providerMinStake;
  }
}
