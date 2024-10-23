// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

import {ISessionStorage} from "../storage/ISessionStorage.sol";
import {IBidStorage} from "../storage/IBidStorage.sol";
import {IStatsStorage} from "../storage/IStatsStorage.sol";

interface ISessionRouter is ISessionStorage {
    event SessionOpened(address indexed user, bytes32 indexed sessionId, address indexed providerId);
    event SessionClosed(address indexed user, bytes32 indexed sessionId, address indexed providerId);
    event UserWithdrawn(address indexed user, uint256 amount_);
    error SessionProviderSignatureMismatch();
    error SesssionApproveExpired();
    error SesssionApprovedForAnotherChainId();
    error SessionDuplicateApproval();
    error SessionApprovedForAnotherUser(); // Means that approval generated for another user address, protection from front-running
    error SesssionReceiptForAnotherChainId();
    error SesssionReceiptExpired();
    error SessionTooShort();
    error SessionAlreadyClosed();
    error SessionNotEndedOrNotExist();
    error SessionProviderNothingToClaimInThisPeriod();
    error SessionBidNotFound();
    error SessionPoolIndexOutOfBounds();
    error SessionUserAmountToWithdrawIsZero();

    function __SessionRouter_init(address fundingAccount_, Pool[] calldata pools_) external;

    /**
     * @notice Sets distibution pool configuration
     * @dev parameters should be the same as in Ethereum L1 Distribution contract
     * @dev at address 0x47176B2Af9885dC6C4575d4eFd63895f7Aaa4790
     * @dev call 'Distribution.pools(3)' where '3' is a poolId
     */
    function setPoolConfig(uint256 index_, Pool calldata pool_) external;

    function openSession(
        uint256 amount_,
        bytes calldata approvalEncoded_,
        bytes calldata signature_
    ) external returns (bytes32);

    function getSessionId(
        address user_,
        address provider_,
        bytes32 bidId_,
        uint256 sessionNonce_
    ) external pure returns (bytes32);

    function getSessionEnd(uint256 amount_, uint256 pricePerSecond_, uint128 openedAt_) external view returns (uint128);

    /**
     * @dev Returns stipend of user based on their stake
     * (User session stake amount / MOR Supply without Compute) * (MOR Compute Supply / 100)
     */
    function stakeToStipend(uint256 amount_, uint128 timestamp_) external view returns (uint256);

    function closeSession(bytes calldata receiptEncoded_, bytes calldata signature_) external;

    /**
     * @dev Allows providers to receive their funds after the end or closure of the session
     */
    function claimForProvider(bytes32 sessionId_) external;

    /**
     * @notice Returns stake of user based on their stipend
     */
    function stipendToStake(uint256 stipend_, uint128 timestamp_) external view returns (uint256);

    function getUserStakesOnHold(
        address user_,
        uint8 iterations_
    ) external view returns (uint256 available_, uint256 hold_);

    function withdrawUserStakes(uint8 iterations_) external;

    /**
     * @dev Returns today's budget in MOR. 1%
     */
    function getTodaysBudget(uint128 timestamp_) external view returns (uint256);

    /**
     * @dev Returns today's compute balance in MOR without claimed amount
     */
    function getComputeBalance(uint128 timestamp_) external view returns (uint256);

    /**
     * @dev Total amount of MOR tokens that were distributed across all pools
     * without compute pool rewards and with compute claimed rewards
     */
    function totalMORSupply(uint128 timestamp_) external view returns (uint256);

    function startOfTheDay(uint128 timestamp_) external pure returns (uint128);
}
