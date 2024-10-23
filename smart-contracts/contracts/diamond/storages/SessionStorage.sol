// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {Paginator} from "@solarity/solidity-lib/libs/arrays/Paginator.sol";

import {ISessionStorage} from "../../interfaces/storage/ISessionStorage.sol";

contract SessionStorage is ISessionStorage {
    using Paginator for *;
    using EnumerableSet for EnumerableSet.Bytes32Set;

    struct SessionsStorage {
        // Account which stores the MOR tokens with infinite allowance for this contract
        address fundingAccount;
        // Distribution pools configuration that mirrors L1 contract
        Pool[] pools;
        // Total amount of MOR claimed by providers
        uint256 providersTotalClaimed;
        // Used to generate unique session ID
        uint256 sessionNonce;
        mapping(bytes32 sessionId => Session) sessions;
        // Session registry for providers, users and models
        mapping(address user => EnumerableSet.Bytes32Set) userSessions;
        mapping(address provider => EnumerableSet.Bytes32Set) providerSessions;
        mapping(bytes32 modelId => EnumerableSet.Bytes32Set) modelSessions;
        mapping(address user => OnHold[]) userStakesOnHold;
        mapping(bytes providerApproval => bool) isProviderApprovalUsed;
    }

    bytes32 public constant SESSIONS_STORAGE_SLOT = keccak256("diamond.standard.sessions.storage");
    uint32 public constant MIN_SESSION_DURATION = 5 minutes;
    uint32 public constant MAX_SESSION_DURATION = 1 days;
    uint32 public constant SIGNATURE_TTL = 10 minutes;
    uint256 public constant COMPUTE_POOL_INDEX = 3;

    /** PUBLIC, GETTERS */
    function getSession(bytes32 sessionId_) external view returns (Session memory) {
        return getSessionsStorage().sessions[sessionId_];
    }

    function getUserSessions(address user_, uint256 offset_, uint256 limit_) external view returns (bytes32[] memory) {
        return getSessionsStorage().userSessions[user_].part(offset_, limit_);
    }

    function getProviderSessions(
        address provider_,
        uint256 offset_,
        uint256 limit_
    ) external view returns (bytes32[] memory) {
        return getSessionsStorage().providerSessions[provider_].part(offset_, limit_);
    }

    function getModelSessions(
        bytes32 modelId_,
        uint256 offset_,
        uint256 limit_
    ) external view returns (bytes32[] memory) {
        return getSessionsStorage().modelSessions[modelId_].part(offset_, limit_);
    }

    function getPools() external view returns (Pool[] memory) {
        return getSessionsStorage().pools;
    }

    function getPool(uint256 index_) external view returns (Pool memory) {
        return getSessionsStorage().pools[index_];
    }

    function getFundingAccount() external view returns (address) {
        return getSessionsStorage().fundingAccount;
    }

    function getTotalSessions(address providerAddr_) public view returns (uint256) {
        return getSessionsStorage().providerSessions[providerAddr_].length();
    }

    function getProvidersTotalClaimed() external view returns (uint256) {
        return getSessionsStorage().providersTotalClaimed;
    }

    function getIsProviderApprovalUsed(bytes memory approval_) external view returns (bool) {
        return getSessionsStorage().isProviderApprovalUsed[approval_];
    }

    // /** INTERNAL, GETTERS */
    // function pools() internal view returns (Pool[] storage) {
    //     return _getSessionStorage().pools;
    // }

    // function pool(uint256 poolIndex_) internal view returns (Pool storage) {
    //     return _getSessionStorage().pools[poolIndex_];
    // }

    // function userStakesOnHold(address user_) internal view returns (OnHold[] storage) {
    //     return _getSessionStorage().userStakesOnHold[user_];
    // }

    // function sessions(bytes32 sessionId_) internal view returns (Session storage) {
    //     return _getSessionStorage().sessions[sessionId_];
    // }

    /** INTERNAL, SETTERS */
    // function setFundingAccount(address fundingAccount_) internal {
    //     _getSessionStorage().fundingAccount = fundingAccount_;
    // }

    // function setPools(Pool[] calldata pools_) internal {
    //     SNStorage storage s = _getSessionStorage();

    //     for (uint256 i = 0; i < pools_.length; i++) {
    //         s.pools.push(pools_[i]);
    //     }
    // }

    // function setPool(uint256 index_, Pool calldata pool_) internal {
    //     _getSessionStorage().pools[index_] = pool_;
    // }

    // function addUserSessionId(address user_, bytes32 sessionId_) internal {
    //     _getSessionStorage().userSessions[user_].add(sessionId_);
    // }

    // function addProviderSessionId(address provider_, bytes32 sessionId_) internal {
    //     _getSessionStorage().providerSessions[provider_].add(sessionId_);
    // }

    // function addModelSessionId(bytes32 modelId, bytes32 sessionId) internal {
    //     _getSessionStorage().modelSessions[modelId].add(sessionId);
    // }

    // function addUserStakeOnHold(address user, OnHold memory onHold) internal {
    //     _getSessionStorage().userStakesOnHold[user].push(onHold);
    // }

    // function increaseProvidersTotalClaimed(uint256 amount) internal {
    //     _getSessionStorage().providersTotalClaimed += amount;
    // }

    // function incrementSessionNonce() internal returns (uint256) {
    //     return _getSessionStorage().sessionNonce++;
    // }

    // function setIsProviderApprovalUsed(bytes memory approval_, bool isUsed_) internal {
    //     _getSessionStorage().isProviderApprovalUsed[approval_] = isUsed_;
    // }

    /** INTERNAL */
    function getSessionsStorage() internal pure returns (SessionsStorage storage ds) {
        bytes32 slot_ = SESSIONS_STORAGE_SLOT;

        assembly {
            ds.slot := slot_
        }
    }
}
