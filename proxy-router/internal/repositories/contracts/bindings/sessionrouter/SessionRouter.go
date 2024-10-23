// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sessionrouter

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IBidStorageBid is an auto generated low-level Go binding around an user-defined struct.
type IBidStorageBid struct {
	Provider       common.Address
	ModelId        [32]byte
	PricePerSecond *big.Int
	Nonce          *big.Int
	CreatedAt      *big.Int
	DeletedAt      *big.Int
}

// IProviderStorageProvider is an auto generated low-level Go binding around an user-defined struct.
type IProviderStorageProvider struct {
	Endpoint          string
	Stake             *big.Int
	CreatedAt         *big.Int
	LimitPeriodEnd    *big.Int
	LimitPeriodEarned *big.Int
	IsDeleted         bool
}

// ISessionStoragePool is an auto generated low-level Go binding around an user-defined struct.
type ISessionStoragePool struct {
	InitialReward    *big.Int
	RewardDecrease   *big.Int
	PayoutStart      *big.Int
	DecreaseInterval *big.Int
}

// ISessionStorageSession is an auto generated low-level Go binding around an user-defined struct.
type ISessionStorageSession struct {
	User                    common.Address
	BidId                   [32]byte
	Stake                   *big.Int
	CloseoutReceipt         []byte
	CloseoutType            *big.Int
	ProviderWithdrawnAmount *big.Int
	OpenedAt                *big.Int
	EndsAt                  *big.Int
	ClosedAt                *big.Int
	IsActive                bool
}

// IStatsStorageModelStats is an auto generated low-level Go binding around an user-defined struct.
type IStatsStorageModelStats struct {
	TpsScaled1000 LibSDSD
	TtftMs        LibSDSD
	TotalDuration LibSDSD
	Count         uint32
}

// IStatsStorageProviderModelStats is an auto generated low-level Go binding around an user-defined struct.
type IStatsStorageProviderModelStats struct {
	TpsScaled1000 LibSDSD
	TtftMs        LibSDSD
	TotalDuration uint32
	SuccessCount  uint32
	TotalCount    uint32
}

// LibSDSD is an auto generated low-level Go binding around an user-defined struct.
type LibSDSD struct {
	Mean  int64
	SqSum int64
}

// SessionRouterMetaData contains all meta data concerning the SessionRouter contract.
var SessionRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SessionAlreadyClosed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SessionApprovedForAnotherUser\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SessionBidNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SessionDuplicateApproval\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SessionNotEndedOrNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SessionPoolIndexOutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SessionProviderNothingToClaimInThisPeriod\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SessionProviderSignatureMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SessionTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SessionUserAmountToWithdrawIsZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SesssionApproveExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SesssionApprovedForAnotherChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SesssionReceiptExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SesssionReceiptForAnotherChainId\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storageSlot\",\"type\":\"bytes32\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sessionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"providerId\",\"type\":\"address\"}],\"name\":\"SessionClosed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sessionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"providerId\",\"type\":\"address\"}],\"name\":\"SessionOpened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"UserWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BIDS_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMPUTE_POOL_INDEX\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DIAMOND_OWNABLE_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SESSION_DURATION\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_SESSION_DURATION\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVIDERS_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SESSIONS_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SIGNATURE_TTL\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STATS_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fundingAccount_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"initialReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDecrease\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"payoutStart\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"decreaseInterval\",\"type\":\"uint128\"}],\"internalType\":\"structISessionStorage.Pool[]\",\"name\":\"pools_\",\"type\":\"tuple[]\"}],\"name\":\"__SessionRouter_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sessionId_\",\"type\":\"bytes32\"}],\"name\":\"claimForProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"receiptEncoded_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"closeSession\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getActiveProviders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bidId_\",\"type\":\"bytes32\"}],\"name\":\"getBid\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"modelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"createdAt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deletedAt\",\"type\":\"uint128\"}],\"internalType\":\"structIBidStorage.Bid\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"timestamp_\",\"type\":\"uint128\"}],\"name\":\"getComputeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFundingAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"}],\"name\":\"getIsProviderActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"approval_\",\"type\":\"bytes\"}],\"name\":\"getIsProviderApprovalUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getModelActiveBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getModelBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getModelSessions\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"}],\"name\":\"getModelStats\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"mean\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"sqSum\",\"type\":\"int64\"}],\"internalType\":\"structLibSD.SD\",\"name\":\"tpsScaled1000\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"mean\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"sqSum\",\"type\":\"int64\"}],\"internalType\":\"structLibSD.SD\",\"name\":\"ttftMs\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"mean\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"sqSum\",\"type\":\"int64\"}],\"internalType\":\"structLibSD.SD\",\"name\":\"totalDuration\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"count\",\"type\":\"uint32\"}],\"internalType\":\"structIStatsStorage.ModelStats\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"}],\"name\":\"getPool\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"initialReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDecrease\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"payoutStart\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"decreaseInterval\",\"type\":\"uint128\"}],\"internalType\":\"structISessionStorage.Pool\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPools\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"initialReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDecrease\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"payoutStart\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"decreaseInterval\",\"type\":\"uint128\"}],\"internalType\":\"structISessionStorage.Pool[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"}],\"name\":\"getProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"createdAt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"limitPeriodEnd\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"limitPeriodEarned\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isDeleted\",\"type\":\"bool\"}],\"internalType\":\"structIProviderStorage.Provider\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getProviderActiveBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getProviderBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProviderMinimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"}],\"name\":\"getProviderModelStats\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"mean\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"sqSum\",\"type\":\"int64\"}],\"internalType\":\"structLibSD.SD\",\"name\":\"tpsScaled1000\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"mean\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"sqSum\",\"type\":\"int64\"}],\"internalType\":\"structLibSD.SD\",\"name\":\"ttftMs\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"totalDuration\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"successCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalCount\",\"type\":\"uint32\"}],\"internalType\":\"structIStatsStorage.ProviderModelStats\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getProviderSessions\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProvidersTotalClaimed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sessionId_\",\"type\":\"bytes32\"}],\"name\":\"getSession\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"bidId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"closeoutReceipt\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"closeoutType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"providerWithdrawnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"openedAt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"endsAt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"closedAt\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structISessionStorage.Session\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond_\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"openedAt_\",\"type\":\"uint128\"}],\"name\":\"getSessionEnd\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"bidId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"sessionNonce_\",\"type\":\"uint256\"}],\"name\":\"getSessionId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"timestamp_\",\"type\":\"uint128\"}],\"name\":\"getTodaysBudget\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"providerAddr_\",\"type\":\"address\"}],\"name\":\"getTotalSessions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getUserSessions\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user_\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"iterations_\",\"type\":\"uint8\"}],\"name\":\"getUserStakesOnHold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"available_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hold_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bidId_\",\"type\":\"bytes32\"}],\"name\":\"isBidActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"approvalEncoded_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"openSession\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index_\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"initialReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardDecrease\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"payoutStart\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"decreaseInterval\",\"type\":\"uint128\"}],\"internalType\":\"structISessionStorage.Pool\",\"name\":\"pool_\",\"type\":\"tuple\"}],\"name\":\"setPoolConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"timestamp_\",\"type\":\"uint128\"}],\"name\":\"stakeToStipend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"timestamp_\",\"type\":\"uint128\"}],\"name\":\"startOfTheDay\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stipend_\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"timestamp_\",\"type\":\"uint128\"}],\"name\":\"stipendToStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"timestamp_\",\"type\":\"uint128\"}],\"name\":\"totalMORSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"iterations_\",\"type\":\"uint8\"}],\"name\":\"withdrawUserStakes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SessionRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use SessionRouterMetaData.ABI instead.
var SessionRouterABI = SessionRouterMetaData.ABI

// SessionRouter is an auto generated Go binding around an Ethereum contract.
type SessionRouter struct {
	SessionRouterCaller     // Read-only binding to the contract
	SessionRouterTransactor // Write-only binding to the contract
	SessionRouterFilterer   // Log filterer for contract events
}

// SessionRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type SessionRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SessionRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SessionRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SessionRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SessionRouterSession struct {
	Contract     *SessionRouter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SessionRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SessionRouterCallerSession struct {
	Contract *SessionRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SessionRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SessionRouterTransactorSession struct {
	Contract     *SessionRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SessionRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type SessionRouterRaw struct {
	Contract *SessionRouter // Generic contract binding to access the raw methods on
}

// SessionRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SessionRouterCallerRaw struct {
	Contract *SessionRouterCaller // Generic read-only contract binding to access the raw methods on
}

// SessionRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SessionRouterTransactorRaw struct {
	Contract *SessionRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSessionRouter creates a new instance of SessionRouter, bound to a specific deployed contract.
func NewSessionRouter(address common.Address, backend bind.ContractBackend) (*SessionRouter, error) {
	contract, err := bindSessionRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SessionRouter{SessionRouterCaller: SessionRouterCaller{contract: contract}, SessionRouterTransactor: SessionRouterTransactor{contract: contract}, SessionRouterFilterer: SessionRouterFilterer{contract: contract}}, nil
}

// NewSessionRouterCaller creates a new read-only instance of SessionRouter, bound to a specific deployed contract.
func NewSessionRouterCaller(address common.Address, caller bind.ContractCaller) (*SessionRouterCaller, error) {
	contract, err := bindSessionRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SessionRouterCaller{contract: contract}, nil
}

// NewSessionRouterTransactor creates a new write-only instance of SessionRouter, bound to a specific deployed contract.
func NewSessionRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*SessionRouterTransactor, error) {
	contract, err := bindSessionRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SessionRouterTransactor{contract: contract}, nil
}

// NewSessionRouterFilterer creates a new log filterer instance of SessionRouter, bound to a specific deployed contract.
func NewSessionRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*SessionRouterFilterer, error) {
	contract, err := bindSessionRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SessionRouterFilterer{contract: contract}, nil
}

// bindSessionRouter binds a generic wrapper to an already deployed contract.
func bindSessionRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SessionRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SessionRouter *SessionRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SessionRouter.Contract.SessionRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SessionRouter *SessionRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SessionRouter.Contract.SessionRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SessionRouter *SessionRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SessionRouter.Contract.SessionRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SessionRouter *SessionRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SessionRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SessionRouter *SessionRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SessionRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SessionRouter *SessionRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SessionRouter.Contract.contract.Transact(opts, method, params...)
}

// BIDSSTORAGESLOT is a free data retrieval call binding the contract method 0x266ccff0.
//
// Solidity: function BIDS_STORAGE_SLOT() view returns(bytes32)
func (_SessionRouter *SessionRouterCaller) BIDSSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "BIDS_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BIDSSTORAGESLOT is a free data retrieval call binding the contract method 0x266ccff0.
//
// Solidity: function BIDS_STORAGE_SLOT() view returns(bytes32)
func (_SessionRouter *SessionRouterSession) BIDSSTORAGESLOT() ([32]byte, error) {
	return _SessionRouter.Contract.BIDSSTORAGESLOT(&_SessionRouter.CallOpts)
}

// BIDSSTORAGESLOT is a free data retrieval call binding the contract method 0x266ccff0.
//
// Solidity: function BIDS_STORAGE_SLOT() view returns(bytes32)
func (_SessionRouter *SessionRouterCallerSession) BIDSSTORAGESLOT() ([32]byte, error) {
	return _SessionRouter.Contract.BIDSSTORAGESLOT(&_SessionRouter.CallOpts)
}

// COMPUTEPOOLINDEX is a free data retrieval call binding the contract method 0xc56d09a0.
//
// Solidity: function COMPUTE_POOL_INDEX() view returns(uint256)
func (_SessionRouter *SessionRouterCaller) COMPUTEPOOLINDEX(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "COMPUTE_POOL_INDEX")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMPUTEPOOLINDEX is a free data retrieval call binding the contract method 0xc56d09a0.
//
// Solidity: function COMPUTE_POOL_INDEX() view returns(uint256)
func (_SessionRouter *SessionRouterSession) COMPUTEPOOLINDEX() (*big.Int, error) {
	return _SessionRouter.Contract.COMPUTEPOOLINDEX(&_SessionRouter.CallOpts)
}

// COMPUTEPOOLINDEX is a free data retrieval call binding the contract method 0xc56d09a0.
//
// Solidity: function COMPUTE_POOL_INDEX() view returns(uint256)
func (_SessionRouter *SessionRouterCallerSession) COMPUTEPOOLINDEX() (*big.Int, error) {
	return _SessionRouter.Contract.COMPUTEPOOLINDEX(&_SessionRouter.CallOpts)
}

// DIAMONDOWNABLESTORAGESLOT is a free data retrieval call binding the contract method 0x4ac3371e.
//
// Solidity: function DIAMOND_OWNABLE_STORAGE_SLOT() view returns(bytes32)
func (_SessionRouter *SessionRouterCaller) DIAMONDOWNABLESTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "DIAMOND_OWNABLE_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DIAMONDOWNABLESTORAGESLOT is a free data retrieval call binding the contract method 0x4ac3371e.
//
// Solidity: function DIAMOND_OWNABLE_STORAGE_SLOT() view returns(bytes32)
func (_SessionRouter *SessionRouterSession) DIAMONDOWNABLESTORAGESLOT() ([32]byte, error) {
	return _SessionRouter.Contract.DIAMONDOWNABLESTORAGESLOT(&_SessionRouter.CallOpts)
}

// DIAMONDOWNABLESTORAGESLOT is a free data retrieval call binding the contract method 0x4ac3371e.
//
// Solidity: function DIAMOND_OWNABLE_STORAGE_SLOT() view returns(bytes32)
func (_SessionRouter *SessionRouterCallerSession) DIAMONDOWNABLESTORAGESLOT() ([32]byte, error) {
	return _SessionRouter.Contract.DIAMONDOWNABLESTORAGESLOT(&_SessionRouter.CallOpts)
}

// MAXSESSIONDURATION is a free data retrieval call binding the contract method 0xcd8cd4ad.
//
// Solidity: function MAX_SESSION_DURATION() view returns(uint32)
func (_SessionRouter *SessionRouterCaller) MAXSESSIONDURATION(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "MAX_SESSION_DURATION")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXSESSIONDURATION is a free data retrieval call binding the contract method 0xcd8cd4ad.
//
// Solidity: function MAX_SESSION_DURATION() view returns(uint32)
func (_SessionRouter *SessionRouterSession) MAXSESSIONDURATION() (uint32, error) {
	return _SessionRouter.Contract.MAXSESSIONDURATION(&_SessionRouter.CallOpts)
}

// MAXSESSIONDURATION is a free data retrieval call binding the contract method 0xcd8cd4ad.
//
// Solidity: function MAX_SESSION_DURATION() view returns(uint32)
func (_SessionRouter *SessionRouterCallerSession) MAXSESSIONDURATION() (uint32, error) {
	return _SessionRouter.Contract.MAXSESSIONDURATION(&_SessionRouter.CallOpts)
}

// MINSESSIONDURATION is a free data retrieval call binding the contract method 0x7d980286.
//
// Solidity: function MIN_SESSION_DURATION() view returns(uint32)
func (_SessionRouter *SessionRouterCaller) MINSESSIONDURATION(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "MIN_SESSION_DURATION")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MINSESSIONDURATION is a free data retrieval call binding the contract method 0x7d980286.
//
// Solidity: function MIN_SESSION_DURATION() view returns(uint32)
func (_SessionRouter *SessionRouterSession) MINSESSIONDURATION() (uint32, error) {
	return _SessionRouter.Contract.MINSESSIONDURATION(&_SessionRouter.CallOpts)
}

// MINSESSIONDURATION is a free data retrieval call binding the contract method 0x7d980286.
//
// Solidity: function MIN_SESSION_DURATION() view returns(uint32)
func (_SessionRouter *SessionRouterCallerSession) MINSESSIONDURATION() (uint32, error) {
	return _SessionRouter.Contract.MINSESSIONDURATION(&_SessionRouter.CallOpts)
}

// PROVIDERSSTORAGESLOT is a free data retrieval call binding the contract method 0xc51830f6.
//
// Solidity: function PROVIDERS_STORAGE_SLOT() view returns(bytes32)
func (_SessionRouter *SessionRouterCaller) PROVIDERSSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "PROVIDERS_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROVIDERSSTORAGESLOT is a free data retrieval call binding the contract method 0xc51830f6.
//
// Solidity: function PROVIDERS_STORAGE_SLOT() view returns(bytes32)
func (_SessionRouter *SessionRouterSession) PROVIDERSSTORAGESLOT() ([32]byte, error) {
	return _SessionRouter.Contract.PROVIDERSSTORAGESLOT(&_SessionRouter.CallOpts)
}

// PROVIDERSSTORAGESLOT is a free data retrieval call binding the contract method 0xc51830f6.
//
// Solidity: function PROVIDERS_STORAGE_SLOT() view returns(bytes32)
func (_SessionRouter *SessionRouterCallerSession) PROVIDERSSTORAGESLOT() ([32]byte, error) {
	return _SessionRouter.Contract.PROVIDERSSTORAGESLOT(&_SessionRouter.CallOpts)
}

// SESSIONSSTORAGESLOT is a free data retrieval call binding the contract method 0xb392636e.
//
// Solidity: function SESSIONS_STORAGE_SLOT() view returns(bytes32)
func (_SessionRouter *SessionRouterCaller) SESSIONSSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "SESSIONS_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SESSIONSSTORAGESLOT is a free data retrieval call binding the contract method 0xb392636e.
//
// Solidity: function SESSIONS_STORAGE_SLOT() view returns(bytes32)
func (_SessionRouter *SessionRouterSession) SESSIONSSTORAGESLOT() ([32]byte, error) {
	return _SessionRouter.Contract.SESSIONSSTORAGESLOT(&_SessionRouter.CallOpts)
}

// SESSIONSSTORAGESLOT is a free data retrieval call binding the contract method 0xb392636e.
//
// Solidity: function SESSIONS_STORAGE_SLOT() view returns(bytes32)
func (_SessionRouter *SessionRouterCallerSession) SESSIONSSTORAGESLOT() ([32]byte, error) {
	return _SessionRouter.Contract.SESSIONSSTORAGESLOT(&_SessionRouter.CallOpts)
}

// SIGNATURETTL is a free data retrieval call binding the contract method 0xe7d791d0.
//
// Solidity: function SIGNATURE_TTL() view returns(uint32)
func (_SessionRouter *SessionRouterCaller) SIGNATURETTL(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "SIGNATURE_TTL")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// SIGNATURETTL is a free data retrieval call binding the contract method 0xe7d791d0.
//
// Solidity: function SIGNATURE_TTL() view returns(uint32)
func (_SessionRouter *SessionRouterSession) SIGNATURETTL() (uint32, error) {
	return _SessionRouter.Contract.SIGNATURETTL(&_SessionRouter.CallOpts)
}

// SIGNATURETTL is a free data retrieval call binding the contract method 0xe7d791d0.
//
// Solidity: function SIGNATURE_TTL() view returns(uint32)
func (_SessionRouter *SessionRouterCallerSession) SIGNATURETTL() (uint32, error) {
	return _SessionRouter.Contract.SIGNATURETTL(&_SessionRouter.CallOpts)
}

// STATSSTORAGESLOT is a free data retrieval call binding the contract method 0x87040789.
//
// Solidity: function STATS_STORAGE_SLOT() view returns(bytes32)
func (_SessionRouter *SessionRouterCaller) STATSSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "STATS_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STATSSTORAGESLOT is a free data retrieval call binding the contract method 0x87040789.
//
// Solidity: function STATS_STORAGE_SLOT() view returns(bytes32)
func (_SessionRouter *SessionRouterSession) STATSSTORAGESLOT() ([32]byte, error) {
	return _SessionRouter.Contract.STATSSTORAGESLOT(&_SessionRouter.CallOpts)
}

// STATSSTORAGESLOT is a free data retrieval call binding the contract method 0x87040789.
//
// Solidity: function STATS_STORAGE_SLOT() view returns(bytes32)
func (_SessionRouter *SessionRouterCallerSession) STATSSTORAGESLOT() ([32]byte, error) {
	return _SessionRouter.Contract.STATSSTORAGESLOT(&_SessionRouter.CallOpts)
}

// GetActiveProviders is a free data retrieval call binding the contract method 0xd5472642.
//
// Solidity: function getActiveProviders(uint256 offset_, uint256 limit_) view returns(address[])
func (_SessionRouter *SessionRouterCaller) GetActiveProviders(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getActiveProviders", offset_, limit_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveProviders is a free data retrieval call binding the contract method 0xd5472642.
//
// Solidity: function getActiveProviders(uint256 offset_, uint256 limit_) view returns(address[])
func (_SessionRouter *SessionRouterSession) GetActiveProviders(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _SessionRouter.Contract.GetActiveProviders(&_SessionRouter.CallOpts, offset_, limit_)
}

// GetActiveProviders is a free data retrieval call binding the contract method 0xd5472642.
//
// Solidity: function getActiveProviders(uint256 offset_, uint256 limit_) view returns(address[])
func (_SessionRouter *SessionRouterCallerSession) GetActiveProviders(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _SessionRouter.Contract.GetActiveProviders(&_SessionRouter.CallOpts, offset_, limit_)
}

// GetBid is a free data retrieval call binding the contract method 0x91704e1e.
//
// Solidity: function getBid(bytes32 bidId_) view returns((address,bytes32,uint256,uint256,uint128,uint128))
func (_SessionRouter *SessionRouterCaller) GetBid(opts *bind.CallOpts, bidId_ [32]byte) (IBidStorageBid, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getBid", bidId_)

	if err != nil {
		return *new(IBidStorageBid), err
	}

	out0 := *abi.ConvertType(out[0], new(IBidStorageBid)).(*IBidStorageBid)

	return out0, err

}

// GetBid is a free data retrieval call binding the contract method 0x91704e1e.
//
// Solidity: function getBid(bytes32 bidId_) view returns((address,bytes32,uint256,uint256,uint128,uint128))
func (_SessionRouter *SessionRouterSession) GetBid(bidId_ [32]byte) (IBidStorageBid, error) {
	return _SessionRouter.Contract.GetBid(&_SessionRouter.CallOpts, bidId_)
}

// GetBid is a free data retrieval call binding the contract method 0x91704e1e.
//
// Solidity: function getBid(bytes32 bidId_) view returns((address,bytes32,uint256,uint256,uint128,uint128))
func (_SessionRouter *SessionRouterCallerSession) GetBid(bidId_ [32]byte) (IBidStorageBid, error) {
	return _SessionRouter.Contract.GetBid(&_SessionRouter.CallOpts, bidId_)
}

// GetComputeBalance is a free data retrieval call binding the contract method 0x61ce471a.
//
// Solidity: function getComputeBalance(uint128 timestamp_) view returns(uint256)
func (_SessionRouter *SessionRouterCaller) GetComputeBalance(opts *bind.CallOpts, timestamp_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getComputeBalance", timestamp_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetComputeBalance is a free data retrieval call binding the contract method 0x61ce471a.
//
// Solidity: function getComputeBalance(uint128 timestamp_) view returns(uint256)
func (_SessionRouter *SessionRouterSession) GetComputeBalance(timestamp_ *big.Int) (*big.Int, error) {
	return _SessionRouter.Contract.GetComputeBalance(&_SessionRouter.CallOpts, timestamp_)
}

// GetComputeBalance is a free data retrieval call binding the contract method 0x61ce471a.
//
// Solidity: function getComputeBalance(uint128 timestamp_) view returns(uint256)
func (_SessionRouter *SessionRouterCallerSession) GetComputeBalance(timestamp_ *big.Int) (*big.Int, error) {
	return _SessionRouter.Contract.GetComputeBalance(&_SessionRouter.CallOpts, timestamp_)
}

// GetFundingAccount is a free data retrieval call binding the contract method 0x775c3727.
//
// Solidity: function getFundingAccount() view returns(address)
func (_SessionRouter *SessionRouterCaller) GetFundingAccount(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getFundingAccount")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFundingAccount is a free data retrieval call binding the contract method 0x775c3727.
//
// Solidity: function getFundingAccount() view returns(address)
func (_SessionRouter *SessionRouterSession) GetFundingAccount() (common.Address, error) {
	return _SessionRouter.Contract.GetFundingAccount(&_SessionRouter.CallOpts)
}

// GetFundingAccount is a free data retrieval call binding the contract method 0x775c3727.
//
// Solidity: function getFundingAccount() view returns(address)
func (_SessionRouter *SessionRouterCallerSession) GetFundingAccount() (common.Address, error) {
	return _SessionRouter.Contract.GetFundingAccount(&_SessionRouter.CallOpts)
}

// GetIsProviderActive is a free data retrieval call binding the contract method 0x63ef175d.
//
// Solidity: function getIsProviderActive(address provider_) view returns(bool)
func (_SessionRouter *SessionRouterCaller) GetIsProviderActive(opts *bind.CallOpts, provider_ common.Address) (bool, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getIsProviderActive", provider_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsProviderActive is a free data retrieval call binding the contract method 0x63ef175d.
//
// Solidity: function getIsProviderActive(address provider_) view returns(bool)
func (_SessionRouter *SessionRouterSession) GetIsProviderActive(provider_ common.Address) (bool, error) {
	return _SessionRouter.Contract.GetIsProviderActive(&_SessionRouter.CallOpts, provider_)
}

// GetIsProviderActive is a free data retrieval call binding the contract method 0x63ef175d.
//
// Solidity: function getIsProviderActive(address provider_) view returns(bool)
func (_SessionRouter *SessionRouterCallerSession) GetIsProviderActive(provider_ common.Address) (bool, error) {
	return _SessionRouter.Contract.GetIsProviderActive(&_SessionRouter.CallOpts, provider_)
}

// GetIsProviderApprovalUsed is a free data retrieval call binding the contract method 0xdb1cf1e2.
//
// Solidity: function getIsProviderApprovalUsed(bytes approval_) view returns(bool)
func (_SessionRouter *SessionRouterCaller) GetIsProviderApprovalUsed(opts *bind.CallOpts, approval_ []byte) (bool, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getIsProviderApprovalUsed", approval_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsProviderApprovalUsed is a free data retrieval call binding the contract method 0xdb1cf1e2.
//
// Solidity: function getIsProviderApprovalUsed(bytes approval_) view returns(bool)
func (_SessionRouter *SessionRouterSession) GetIsProviderApprovalUsed(approval_ []byte) (bool, error) {
	return _SessionRouter.Contract.GetIsProviderApprovalUsed(&_SessionRouter.CallOpts, approval_)
}

// GetIsProviderApprovalUsed is a free data retrieval call binding the contract method 0xdb1cf1e2.
//
// Solidity: function getIsProviderApprovalUsed(bytes approval_) view returns(bool)
func (_SessionRouter *SessionRouterCallerSession) GetIsProviderApprovalUsed(approval_ []byte) (bool, error) {
	return _SessionRouter.Contract.GetIsProviderApprovalUsed(&_SessionRouter.CallOpts, approval_)
}

// GetModelActiveBids is a free data retrieval call binding the contract method 0x8a683b6e.
//
// Solidity: function getModelActiveBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterCaller) GetModelActiveBids(opts *bind.CallOpts, modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getModelActiveBids", modelId_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetModelActiveBids is a free data retrieval call binding the contract method 0x8a683b6e.
//
// Solidity: function getModelActiveBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterSession) GetModelActiveBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _SessionRouter.Contract.GetModelActiveBids(&_SessionRouter.CallOpts, modelId_, offset_, limit_)
}

// GetModelActiveBids is a free data retrieval call binding the contract method 0x8a683b6e.
//
// Solidity: function getModelActiveBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterCallerSession) GetModelActiveBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _SessionRouter.Contract.GetModelActiveBids(&_SessionRouter.CallOpts, modelId_, offset_, limit_)
}

// GetModelBids is a free data retrieval call binding the contract method 0xfade17b1.
//
// Solidity: function getModelBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterCaller) GetModelBids(opts *bind.CallOpts, modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getModelBids", modelId_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetModelBids is a free data retrieval call binding the contract method 0xfade17b1.
//
// Solidity: function getModelBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterSession) GetModelBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _SessionRouter.Contract.GetModelBids(&_SessionRouter.CallOpts, modelId_, offset_, limit_)
}

// GetModelBids is a free data retrieval call binding the contract method 0xfade17b1.
//
// Solidity: function getModelBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterCallerSession) GetModelBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _SessionRouter.Contract.GetModelBids(&_SessionRouter.CallOpts, modelId_, offset_, limit_)
}

// GetModelSessions is a free data retrieval call binding the contract method 0x1d78a872.
//
// Solidity: function getModelSessions(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterCaller) GetModelSessions(opts *bind.CallOpts, modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getModelSessions", modelId_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetModelSessions is a free data retrieval call binding the contract method 0x1d78a872.
//
// Solidity: function getModelSessions(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterSession) GetModelSessions(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _SessionRouter.Contract.GetModelSessions(&_SessionRouter.CallOpts, modelId_, offset_, limit_)
}

// GetModelSessions is a free data retrieval call binding the contract method 0x1d78a872.
//
// Solidity: function getModelSessions(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterCallerSession) GetModelSessions(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _SessionRouter.Contract.GetModelSessions(&_SessionRouter.CallOpts, modelId_, offset_, limit_)
}

// GetModelStats is a free data retrieval call binding the contract method 0xce535723.
//
// Solidity: function getModelStats(bytes32 modelId_) view returns(((int64,int64),(int64,int64),(int64,int64),uint32))
func (_SessionRouter *SessionRouterCaller) GetModelStats(opts *bind.CallOpts, modelId_ [32]byte) (IStatsStorageModelStats, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getModelStats", modelId_)

	if err != nil {
		return *new(IStatsStorageModelStats), err
	}

	out0 := *abi.ConvertType(out[0], new(IStatsStorageModelStats)).(*IStatsStorageModelStats)

	return out0, err

}

// GetModelStats is a free data retrieval call binding the contract method 0xce535723.
//
// Solidity: function getModelStats(bytes32 modelId_) view returns(((int64,int64),(int64,int64),(int64,int64),uint32))
func (_SessionRouter *SessionRouterSession) GetModelStats(modelId_ [32]byte) (IStatsStorageModelStats, error) {
	return _SessionRouter.Contract.GetModelStats(&_SessionRouter.CallOpts, modelId_)
}

// GetModelStats is a free data retrieval call binding the contract method 0xce535723.
//
// Solidity: function getModelStats(bytes32 modelId_) view returns(((int64,int64),(int64,int64),(int64,int64),uint32))
func (_SessionRouter *SessionRouterCallerSession) GetModelStats(modelId_ [32]byte) (IStatsStorageModelStats, error) {
	return _SessionRouter.Contract.GetModelStats(&_SessionRouter.CallOpts, modelId_)
}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(uint256 index_) view returns((uint256,uint256,uint128,uint128))
func (_SessionRouter *SessionRouterCaller) GetPool(opts *bind.CallOpts, index_ *big.Int) (ISessionStoragePool, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getPool", index_)

	if err != nil {
		return *new(ISessionStoragePool), err
	}

	out0 := *abi.ConvertType(out[0], new(ISessionStoragePool)).(*ISessionStoragePool)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(uint256 index_) view returns((uint256,uint256,uint128,uint128))
func (_SessionRouter *SessionRouterSession) GetPool(index_ *big.Int) (ISessionStoragePool, error) {
	return _SessionRouter.Contract.GetPool(&_SessionRouter.CallOpts, index_)
}

// GetPool is a free data retrieval call binding the contract method 0x068bcd8d.
//
// Solidity: function getPool(uint256 index_) view returns((uint256,uint256,uint128,uint128))
func (_SessionRouter *SessionRouterCallerSession) GetPool(index_ *big.Int) (ISessionStoragePool, error) {
	return _SessionRouter.Contract.GetPool(&_SessionRouter.CallOpts, index_)
}

// GetPools is a free data retrieval call binding the contract method 0x673a2a1f.
//
// Solidity: function getPools() view returns((uint256,uint256,uint128,uint128)[])
func (_SessionRouter *SessionRouterCaller) GetPools(opts *bind.CallOpts) ([]ISessionStoragePool, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getPools")

	if err != nil {
		return *new([]ISessionStoragePool), err
	}

	out0 := *abi.ConvertType(out[0], new([]ISessionStoragePool)).(*[]ISessionStoragePool)

	return out0, err

}

// GetPools is a free data retrieval call binding the contract method 0x673a2a1f.
//
// Solidity: function getPools() view returns((uint256,uint256,uint128,uint128)[])
func (_SessionRouter *SessionRouterSession) GetPools() ([]ISessionStoragePool, error) {
	return _SessionRouter.Contract.GetPools(&_SessionRouter.CallOpts)
}

// GetPools is a free data retrieval call binding the contract method 0x673a2a1f.
//
// Solidity: function getPools() view returns((uint256,uint256,uint128,uint128)[])
func (_SessionRouter *SessionRouterCallerSession) GetPools() ([]ISessionStoragePool, error) {
	return _SessionRouter.Contract.GetPools(&_SessionRouter.CallOpts)
}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address provider_) view returns((string,uint256,uint128,uint128,uint256,bool))
func (_SessionRouter *SessionRouterCaller) GetProvider(opts *bind.CallOpts, provider_ common.Address) (IProviderStorageProvider, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getProvider", provider_)

	if err != nil {
		return *new(IProviderStorageProvider), err
	}

	out0 := *abi.ConvertType(out[0], new(IProviderStorageProvider)).(*IProviderStorageProvider)

	return out0, err

}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address provider_) view returns((string,uint256,uint128,uint128,uint256,bool))
func (_SessionRouter *SessionRouterSession) GetProvider(provider_ common.Address) (IProviderStorageProvider, error) {
	return _SessionRouter.Contract.GetProvider(&_SessionRouter.CallOpts, provider_)
}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address provider_) view returns((string,uint256,uint128,uint128,uint256,bool))
func (_SessionRouter *SessionRouterCallerSession) GetProvider(provider_ common.Address) (IProviderStorageProvider, error) {
	return _SessionRouter.Contract.GetProvider(&_SessionRouter.CallOpts, provider_)
}

// GetProviderActiveBids is a free data retrieval call binding the contract method 0xaf5b77ca.
//
// Solidity: function getProviderActiveBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterCaller) GetProviderActiveBids(opts *bind.CallOpts, provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getProviderActiveBids", provider_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProviderActiveBids is a free data retrieval call binding the contract method 0xaf5b77ca.
//
// Solidity: function getProviderActiveBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterSession) GetProviderActiveBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _SessionRouter.Contract.GetProviderActiveBids(&_SessionRouter.CallOpts, provider_, offset_, limit_)
}

// GetProviderActiveBids is a free data retrieval call binding the contract method 0xaf5b77ca.
//
// Solidity: function getProviderActiveBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterCallerSession) GetProviderActiveBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _SessionRouter.Contract.GetProviderActiveBids(&_SessionRouter.CallOpts, provider_, offset_, limit_)
}

// GetProviderBids is a free data retrieval call binding the contract method 0x59d435c4.
//
// Solidity: function getProviderBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterCaller) GetProviderBids(opts *bind.CallOpts, provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getProviderBids", provider_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProviderBids is a free data retrieval call binding the contract method 0x59d435c4.
//
// Solidity: function getProviderBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterSession) GetProviderBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _SessionRouter.Contract.GetProviderBids(&_SessionRouter.CallOpts, provider_, offset_, limit_)
}

// GetProviderBids is a free data retrieval call binding the contract method 0x59d435c4.
//
// Solidity: function getProviderBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterCallerSession) GetProviderBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _SessionRouter.Contract.GetProviderBids(&_SessionRouter.CallOpts, provider_, offset_, limit_)
}

// GetProviderMinimumStake is a free data retrieval call binding the contract method 0x53c029f6.
//
// Solidity: function getProviderMinimumStake() view returns(uint256)
func (_SessionRouter *SessionRouterCaller) GetProviderMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getProviderMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProviderMinimumStake is a free data retrieval call binding the contract method 0x53c029f6.
//
// Solidity: function getProviderMinimumStake() view returns(uint256)
func (_SessionRouter *SessionRouterSession) GetProviderMinimumStake() (*big.Int, error) {
	return _SessionRouter.Contract.GetProviderMinimumStake(&_SessionRouter.CallOpts)
}

// GetProviderMinimumStake is a free data retrieval call binding the contract method 0x53c029f6.
//
// Solidity: function getProviderMinimumStake() view returns(uint256)
func (_SessionRouter *SessionRouterCallerSession) GetProviderMinimumStake() (*big.Int, error) {
	return _SessionRouter.Contract.GetProviderMinimumStake(&_SessionRouter.CallOpts)
}

// GetProviderModelStats is a free data retrieval call binding the contract method 0x1b26c116.
//
// Solidity: function getProviderModelStats(bytes32 modelId_, address provider_) view returns(((int64,int64),(int64,int64),uint32,uint32,uint32))
func (_SessionRouter *SessionRouterCaller) GetProviderModelStats(opts *bind.CallOpts, modelId_ [32]byte, provider_ common.Address) (IStatsStorageProviderModelStats, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getProviderModelStats", modelId_, provider_)

	if err != nil {
		return *new(IStatsStorageProviderModelStats), err
	}

	out0 := *abi.ConvertType(out[0], new(IStatsStorageProviderModelStats)).(*IStatsStorageProviderModelStats)

	return out0, err

}

// GetProviderModelStats is a free data retrieval call binding the contract method 0x1b26c116.
//
// Solidity: function getProviderModelStats(bytes32 modelId_, address provider_) view returns(((int64,int64),(int64,int64),uint32,uint32,uint32))
func (_SessionRouter *SessionRouterSession) GetProviderModelStats(modelId_ [32]byte, provider_ common.Address) (IStatsStorageProviderModelStats, error) {
	return _SessionRouter.Contract.GetProviderModelStats(&_SessionRouter.CallOpts, modelId_, provider_)
}

// GetProviderModelStats is a free data retrieval call binding the contract method 0x1b26c116.
//
// Solidity: function getProviderModelStats(bytes32 modelId_, address provider_) view returns(((int64,int64),(int64,int64),uint32,uint32,uint32))
func (_SessionRouter *SessionRouterCallerSession) GetProviderModelStats(modelId_ [32]byte, provider_ common.Address) (IStatsStorageProviderModelStats, error) {
	return _SessionRouter.Contract.GetProviderModelStats(&_SessionRouter.CallOpts, modelId_, provider_)
}

// GetProviderSessions is a free data retrieval call binding the contract method 0x87bced7d.
//
// Solidity: function getProviderSessions(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterCaller) GetProviderSessions(opts *bind.CallOpts, provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getProviderSessions", provider_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProviderSessions is a free data retrieval call binding the contract method 0x87bced7d.
//
// Solidity: function getProviderSessions(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterSession) GetProviderSessions(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _SessionRouter.Contract.GetProviderSessions(&_SessionRouter.CallOpts, provider_, offset_, limit_)
}

// GetProviderSessions is a free data retrieval call binding the contract method 0x87bced7d.
//
// Solidity: function getProviderSessions(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterCallerSession) GetProviderSessions(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _SessionRouter.Contract.GetProviderSessions(&_SessionRouter.CallOpts, provider_, offset_, limit_)
}

// GetProvidersTotalClaimed is a free data retrieval call binding the contract method 0xbdfbbada.
//
// Solidity: function getProvidersTotalClaimed() view returns(uint256)
func (_SessionRouter *SessionRouterCaller) GetProvidersTotalClaimed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getProvidersTotalClaimed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProvidersTotalClaimed is a free data retrieval call binding the contract method 0xbdfbbada.
//
// Solidity: function getProvidersTotalClaimed() view returns(uint256)
func (_SessionRouter *SessionRouterSession) GetProvidersTotalClaimed() (*big.Int, error) {
	return _SessionRouter.Contract.GetProvidersTotalClaimed(&_SessionRouter.CallOpts)
}

// GetProvidersTotalClaimed is a free data retrieval call binding the contract method 0xbdfbbada.
//
// Solidity: function getProvidersTotalClaimed() view returns(uint256)
func (_SessionRouter *SessionRouterCallerSession) GetProvidersTotalClaimed() (*big.Int, error) {
	return _SessionRouter.Contract.GetProvidersTotalClaimed(&_SessionRouter.CallOpts)
}

// GetSession is a free data retrieval call binding the contract method 0x39b240bd.
//
// Solidity: function getSession(bytes32 sessionId_) view returns((address,bytes32,uint256,bytes,uint256,uint256,uint128,uint128,uint128,bool))
func (_SessionRouter *SessionRouterCaller) GetSession(opts *bind.CallOpts, sessionId_ [32]byte) (ISessionStorageSession, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getSession", sessionId_)

	if err != nil {
		return *new(ISessionStorageSession), err
	}

	out0 := *abi.ConvertType(out[0], new(ISessionStorageSession)).(*ISessionStorageSession)

	return out0, err

}

// GetSession is a free data retrieval call binding the contract method 0x39b240bd.
//
// Solidity: function getSession(bytes32 sessionId_) view returns((address,bytes32,uint256,bytes,uint256,uint256,uint128,uint128,uint128,bool))
func (_SessionRouter *SessionRouterSession) GetSession(sessionId_ [32]byte) (ISessionStorageSession, error) {
	return _SessionRouter.Contract.GetSession(&_SessionRouter.CallOpts, sessionId_)
}

// GetSession is a free data retrieval call binding the contract method 0x39b240bd.
//
// Solidity: function getSession(bytes32 sessionId_) view returns((address,bytes32,uint256,bytes,uint256,uint256,uint128,uint128,uint128,bool))
func (_SessionRouter *SessionRouterCallerSession) GetSession(sessionId_ [32]byte) (ISessionStorageSession, error) {
	return _SessionRouter.Contract.GetSession(&_SessionRouter.CallOpts, sessionId_)
}

// GetSessionEnd is a free data retrieval call binding the contract method 0xb823dc8f.
//
// Solidity: function getSessionEnd(uint256 amount_, uint256 pricePerSecond_, uint128 openedAt_) view returns(uint128)
func (_SessionRouter *SessionRouterCaller) GetSessionEnd(opts *bind.CallOpts, amount_ *big.Int, pricePerSecond_ *big.Int, openedAt_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getSessionEnd", amount_, pricePerSecond_, openedAt_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSessionEnd is a free data retrieval call binding the contract method 0xb823dc8f.
//
// Solidity: function getSessionEnd(uint256 amount_, uint256 pricePerSecond_, uint128 openedAt_) view returns(uint128)
func (_SessionRouter *SessionRouterSession) GetSessionEnd(amount_ *big.Int, pricePerSecond_ *big.Int, openedAt_ *big.Int) (*big.Int, error) {
	return _SessionRouter.Contract.GetSessionEnd(&_SessionRouter.CallOpts, amount_, pricePerSecond_, openedAt_)
}

// GetSessionEnd is a free data retrieval call binding the contract method 0xb823dc8f.
//
// Solidity: function getSessionEnd(uint256 amount_, uint256 pricePerSecond_, uint128 openedAt_) view returns(uint128)
func (_SessionRouter *SessionRouterCallerSession) GetSessionEnd(amount_ *big.Int, pricePerSecond_ *big.Int, openedAt_ *big.Int) (*big.Int, error) {
	return _SessionRouter.Contract.GetSessionEnd(&_SessionRouter.CallOpts, amount_, pricePerSecond_, openedAt_)
}

// GetSessionId is a free data retrieval call binding the contract method 0x4d689ffd.
//
// Solidity: function getSessionId(address user_, address provider_, bytes32 bidId_, uint256 sessionNonce_) pure returns(bytes32)
func (_SessionRouter *SessionRouterCaller) GetSessionId(opts *bind.CallOpts, user_ common.Address, provider_ common.Address, bidId_ [32]byte, sessionNonce_ *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getSessionId", user_, provider_, bidId_, sessionNonce_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetSessionId is a free data retrieval call binding the contract method 0x4d689ffd.
//
// Solidity: function getSessionId(address user_, address provider_, bytes32 bidId_, uint256 sessionNonce_) pure returns(bytes32)
func (_SessionRouter *SessionRouterSession) GetSessionId(user_ common.Address, provider_ common.Address, bidId_ [32]byte, sessionNonce_ *big.Int) ([32]byte, error) {
	return _SessionRouter.Contract.GetSessionId(&_SessionRouter.CallOpts, user_, provider_, bidId_, sessionNonce_)
}

// GetSessionId is a free data retrieval call binding the contract method 0x4d689ffd.
//
// Solidity: function getSessionId(address user_, address provider_, bytes32 bidId_, uint256 sessionNonce_) pure returns(bytes32)
func (_SessionRouter *SessionRouterCallerSession) GetSessionId(user_ common.Address, provider_ common.Address, bidId_ [32]byte, sessionNonce_ *big.Int) ([32]byte, error) {
	return _SessionRouter.Contract.GetSessionId(&_SessionRouter.CallOpts, user_, provider_, bidId_, sessionNonce_)
}

// GetTodaysBudget is a free data retrieval call binding the contract method 0x40005965.
//
// Solidity: function getTodaysBudget(uint128 timestamp_) view returns(uint256)
func (_SessionRouter *SessionRouterCaller) GetTodaysBudget(opts *bind.CallOpts, timestamp_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getTodaysBudget", timestamp_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTodaysBudget is a free data retrieval call binding the contract method 0x40005965.
//
// Solidity: function getTodaysBudget(uint128 timestamp_) view returns(uint256)
func (_SessionRouter *SessionRouterSession) GetTodaysBudget(timestamp_ *big.Int) (*big.Int, error) {
	return _SessionRouter.Contract.GetTodaysBudget(&_SessionRouter.CallOpts, timestamp_)
}

// GetTodaysBudget is a free data retrieval call binding the contract method 0x40005965.
//
// Solidity: function getTodaysBudget(uint128 timestamp_) view returns(uint256)
func (_SessionRouter *SessionRouterCallerSession) GetTodaysBudget(timestamp_ *big.Int) (*big.Int, error) {
	return _SessionRouter.Contract.GetTodaysBudget(&_SessionRouter.CallOpts, timestamp_)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_SessionRouter *SessionRouterCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_SessionRouter *SessionRouterSession) GetToken() (common.Address, error) {
	return _SessionRouter.Contract.GetToken(&_SessionRouter.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_SessionRouter *SessionRouterCallerSession) GetToken() (common.Address, error) {
	return _SessionRouter.Contract.GetToken(&_SessionRouter.CallOpts)
}

// GetTotalSessions is a free data retrieval call binding the contract method 0x76400935.
//
// Solidity: function getTotalSessions(address providerAddr_) view returns(uint256)
func (_SessionRouter *SessionRouterCaller) GetTotalSessions(opts *bind.CallOpts, providerAddr_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getTotalSessions", providerAddr_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSessions is a free data retrieval call binding the contract method 0x76400935.
//
// Solidity: function getTotalSessions(address providerAddr_) view returns(uint256)
func (_SessionRouter *SessionRouterSession) GetTotalSessions(providerAddr_ common.Address) (*big.Int, error) {
	return _SessionRouter.Contract.GetTotalSessions(&_SessionRouter.CallOpts, providerAddr_)
}

// GetTotalSessions is a free data retrieval call binding the contract method 0x76400935.
//
// Solidity: function getTotalSessions(address providerAddr_) view returns(uint256)
func (_SessionRouter *SessionRouterCallerSession) GetTotalSessions(providerAddr_ common.Address) (*big.Int, error) {
	return _SessionRouter.Contract.GetTotalSessions(&_SessionRouter.CallOpts, providerAddr_)
}

// GetUserSessions is a free data retrieval call binding the contract method 0xeb7764bb.
//
// Solidity: function getUserSessions(address user_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterCaller) GetUserSessions(opts *bind.CallOpts, user_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getUserSessions", user_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetUserSessions is a free data retrieval call binding the contract method 0xeb7764bb.
//
// Solidity: function getUserSessions(address user_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterSession) GetUserSessions(user_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _SessionRouter.Contract.GetUserSessions(&_SessionRouter.CallOpts, user_, offset_, limit_)
}

// GetUserSessions is a free data retrieval call binding the contract method 0xeb7764bb.
//
// Solidity: function getUserSessions(address user_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_SessionRouter *SessionRouterCallerSession) GetUserSessions(user_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _SessionRouter.Contract.GetUserSessions(&_SessionRouter.CallOpts, user_, offset_, limit_)
}

// GetUserStakesOnHold is a free data retrieval call binding the contract method 0x967885df.
//
// Solidity: function getUserStakesOnHold(address user_, uint8 iterations_) view returns(uint256 available_, uint256 hold_)
func (_SessionRouter *SessionRouterCaller) GetUserStakesOnHold(opts *bind.CallOpts, user_ common.Address, iterations_ uint8) (struct {
	Available *big.Int
	Hold      *big.Int
}, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "getUserStakesOnHold", user_, iterations_)

	outstruct := new(struct {
		Available *big.Int
		Hold      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Available = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Hold = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserStakesOnHold is a free data retrieval call binding the contract method 0x967885df.
//
// Solidity: function getUserStakesOnHold(address user_, uint8 iterations_) view returns(uint256 available_, uint256 hold_)
func (_SessionRouter *SessionRouterSession) GetUserStakesOnHold(user_ common.Address, iterations_ uint8) (struct {
	Available *big.Int
	Hold      *big.Int
}, error) {
	return _SessionRouter.Contract.GetUserStakesOnHold(&_SessionRouter.CallOpts, user_, iterations_)
}

// GetUserStakesOnHold is a free data retrieval call binding the contract method 0x967885df.
//
// Solidity: function getUserStakesOnHold(address user_, uint8 iterations_) view returns(uint256 available_, uint256 hold_)
func (_SessionRouter *SessionRouterCallerSession) GetUserStakesOnHold(user_ common.Address, iterations_ uint8) (struct {
	Available *big.Int
	Hold      *big.Int
}, error) {
	return _SessionRouter.Contract.GetUserStakesOnHold(&_SessionRouter.CallOpts, user_, iterations_)
}

// IsBidActive is a free data retrieval call binding the contract method 0x1345df58.
//
// Solidity: function isBidActive(bytes32 bidId_) view returns(bool)
func (_SessionRouter *SessionRouterCaller) IsBidActive(opts *bind.CallOpts, bidId_ [32]byte) (bool, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "isBidActive", bidId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBidActive is a free data retrieval call binding the contract method 0x1345df58.
//
// Solidity: function isBidActive(bytes32 bidId_) view returns(bool)
func (_SessionRouter *SessionRouterSession) IsBidActive(bidId_ [32]byte) (bool, error) {
	return _SessionRouter.Contract.IsBidActive(&_SessionRouter.CallOpts, bidId_)
}

// IsBidActive is a free data retrieval call binding the contract method 0x1345df58.
//
// Solidity: function isBidActive(bytes32 bidId_) view returns(bool)
func (_SessionRouter *SessionRouterCallerSession) IsBidActive(bidId_ [32]byte) (bool, error) {
	return _SessionRouter.Contract.IsBidActive(&_SessionRouter.CallOpts, bidId_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SessionRouter *SessionRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SessionRouter *SessionRouterSession) Owner() (common.Address, error) {
	return _SessionRouter.Contract.Owner(&_SessionRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SessionRouter *SessionRouterCallerSession) Owner() (common.Address, error) {
	return _SessionRouter.Contract.Owner(&_SessionRouter.CallOpts)
}

// StakeToStipend is a free data retrieval call binding the contract method 0xb3cb0d0f.
//
// Solidity: function stakeToStipend(uint256 amount_, uint128 timestamp_) view returns(uint256)
func (_SessionRouter *SessionRouterCaller) StakeToStipend(opts *bind.CallOpts, amount_ *big.Int, timestamp_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "stakeToStipend", amount_, timestamp_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeToStipend is a free data retrieval call binding the contract method 0xb3cb0d0f.
//
// Solidity: function stakeToStipend(uint256 amount_, uint128 timestamp_) view returns(uint256)
func (_SessionRouter *SessionRouterSession) StakeToStipend(amount_ *big.Int, timestamp_ *big.Int) (*big.Int, error) {
	return _SessionRouter.Contract.StakeToStipend(&_SessionRouter.CallOpts, amount_, timestamp_)
}

// StakeToStipend is a free data retrieval call binding the contract method 0xb3cb0d0f.
//
// Solidity: function stakeToStipend(uint256 amount_, uint128 timestamp_) view returns(uint256)
func (_SessionRouter *SessionRouterCallerSession) StakeToStipend(amount_ *big.Int, timestamp_ *big.Int) (*big.Int, error) {
	return _SessionRouter.Contract.StakeToStipend(&_SessionRouter.CallOpts, amount_, timestamp_)
}

// StartOfTheDay is a free data retrieval call binding the contract method 0xba26588c.
//
// Solidity: function startOfTheDay(uint128 timestamp_) pure returns(uint128)
func (_SessionRouter *SessionRouterCaller) StartOfTheDay(opts *bind.CallOpts, timestamp_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "startOfTheDay", timestamp_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartOfTheDay is a free data retrieval call binding the contract method 0xba26588c.
//
// Solidity: function startOfTheDay(uint128 timestamp_) pure returns(uint128)
func (_SessionRouter *SessionRouterSession) StartOfTheDay(timestamp_ *big.Int) (*big.Int, error) {
	return _SessionRouter.Contract.StartOfTheDay(&_SessionRouter.CallOpts, timestamp_)
}

// StartOfTheDay is a free data retrieval call binding the contract method 0xba26588c.
//
// Solidity: function startOfTheDay(uint128 timestamp_) pure returns(uint128)
func (_SessionRouter *SessionRouterCallerSession) StartOfTheDay(timestamp_ *big.Int) (*big.Int, error) {
	return _SessionRouter.Contract.StartOfTheDay(&_SessionRouter.CallOpts, timestamp_)
}

// StipendToStake is a free data retrieval call binding the contract method 0xca40d45f.
//
// Solidity: function stipendToStake(uint256 stipend_, uint128 timestamp_) view returns(uint256)
func (_SessionRouter *SessionRouterCaller) StipendToStake(opts *bind.CallOpts, stipend_ *big.Int, timestamp_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "stipendToStake", stipend_, timestamp_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StipendToStake is a free data retrieval call binding the contract method 0xca40d45f.
//
// Solidity: function stipendToStake(uint256 stipend_, uint128 timestamp_) view returns(uint256)
func (_SessionRouter *SessionRouterSession) StipendToStake(stipend_ *big.Int, timestamp_ *big.Int) (*big.Int, error) {
	return _SessionRouter.Contract.StipendToStake(&_SessionRouter.CallOpts, stipend_, timestamp_)
}

// StipendToStake is a free data retrieval call binding the contract method 0xca40d45f.
//
// Solidity: function stipendToStake(uint256 stipend_, uint128 timestamp_) view returns(uint256)
func (_SessionRouter *SessionRouterCallerSession) StipendToStake(stipend_ *big.Int, timestamp_ *big.Int) (*big.Int, error) {
	return _SessionRouter.Contract.StipendToStake(&_SessionRouter.CallOpts, stipend_, timestamp_)
}

// TotalMORSupply is a free data retrieval call binding the contract method 0x6d0cfe5a.
//
// Solidity: function totalMORSupply(uint128 timestamp_) view returns(uint256)
func (_SessionRouter *SessionRouterCaller) TotalMORSupply(opts *bind.CallOpts, timestamp_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SessionRouter.contract.Call(opts, &out, "totalMORSupply", timestamp_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalMORSupply is a free data retrieval call binding the contract method 0x6d0cfe5a.
//
// Solidity: function totalMORSupply(uint128 timestamp_) view returns(uint256)
func (_SessionRouter *SessionRouterSession) TotalMORSupply(timestamp_ *big.Int) (*big.Int, error) {
	return _SessionRouter.Contract.TotalMORSupply(&_SessionRouter.CallOpts, timestamp_)
}

// TotalMORSupply is a free data retrieval call binding the contract method 0x6d0cfe5a.
//
// Solidity: function totalMORSupply(uint128 timestamp_) view returns(uint256)
func (_SessionRouter *SessionRouterCallerSession) TotalMORSupply(timestamp_ *big.Int) (*big.Int, error) {
	return _SessionRouter.Contract.TotalMORSupply(&_SessionRouter.CallOpts, timestamp_)
}

// SessionRouterInit is a paid mutator transaction binding the contract method 0x44ceb8e0.
//
// Solidity: function __SessionRouter_init(address fundingAccount_, (uint256,uint256,uint128,uint128)[] pools_) returns()
func (_SessionRouter *SessionRouterTransactor) SessionRouterInit(opts *bind.TransactOpts, fundingAccount_ common.Address, pools_ []ISessionStoragePool) (*types.Transaction, error) {
	return _SessionRouter.contract.Transact(opts, "__SessionRouter_init", fundingAccount_, pools_)
}

// SessionRouterInit is a paid mutator transaction binding the contract method 0x44ceb8e0.
//
// Solidity: function __SessionRouter_init(address fundingAccount_, (uint256,uint256,uint128,uint128)[] pools_) returns()
func (_SessionRouter *SessionRouterSession) SessionRouterInit(fundingAccount_ common.Address, pools_ []ISessionStoragePool) (*types.Transaction, error) {
	return _SessionRouter.Contract.SessionRouterInit(&_SessionRouter.TransactOpts, fundingAccount_, pools_)
}

// SessionRouterInit is a paid mutator transaction binding the contract method 0x44ceb8e0.
//
// Solidity: function __SessionRouter_init(address fundingAccount_, (uint256,uint256,uint128,uint128)[] pools_) returns()
func (_SessionRouter *SessionRouterTransactorSession) SessionRouterInit(fundingAccount_ common.Address, pools_ []ISessionStoragePool) (*types.Transaction, error) {
	return _SessionRouter.Contract.SessionRouterInit(&_SessionRouter.TransactOpts, fundingAccount_, pools_)
}

// ClaimForProvider is a paid mutator transaction binding the contract method 0xf2e96e8e.
//
// Solidity: function claimForProvider(bytes32 sessionId_) returns()
func (_SessionRouter *SessionRouterTransactor) ClaimForProvider(opts *bind.TransactOpts, sessionId_ [32]byte) (*types.Transaction, error) {
	return _SessionRouter.contract.Transact(opts, "claimForProvider", sessionId_)
}

// ClaimForProvider is a paid mutator transaction binding the contract method 0xf2e96e8e.
//
// Solidity: function claimForProvider(bytes32 sessionId_) returns()
func (_SessionRouter *SessionRouterSession) ClaimForProvider(sessionId_ [32]byte) (*types.Transaction, error) {
	return _SessionRouter.Contract.ClaimForProvider(&_SessionRouter.TransactOpts, sessionId_)
}

// ClaimForProvider is a paid mutator transaction binding the contract method 0xf2e96e8e.
//
// Solidity: function claimForProvider(bytes32 sessionId_) returns()
func (_SessionRouter *SessionRouterTransactorSession) ClaimForProvider(sessionId_ [32]byte) (*types.Transaction, error) {
	return _SessionRouter.Contract.ClaimForProvider(&_SessionRouter.TransactOpts, sessionId_)
}

// CloseSession is a paid mutator transaction binding the contract method 0x42f77a31.
//
// Solidity: function closeSession(bytes receiptEncoded_, bytes signature_) returns()
func (_SessionRouter *SessionRouterTransactor) CloseSession(opts *bind.TransactOpts, receiptEncoded_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _SessionRouter.contract.Transact(opts, "closeSession", receiptEncoded_, signature_)
}

// CloseSession is a paid mutator transaction binding the contract method 0x42f77a31.
//
// Solidity: function closeSession(bytes receiptEncoded_, bytes signature_) returns()
func (_SessionRouter *SessionRouterSession) CloseSession(receiptEncoded_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _SessionRouter.Contract.CloseSession(&_SessionRouter.TransactOpts, receiptEncoded_, signature_)
}

// CloseSession is a paid mutator transaction binding the contract method 0x42f77a31.
//
// Solidity: function closeSession(bytes receiptEncoded_, bytes signature_) returns()
func (_SessionRouter *SessionRouterTransactorSession) CloseSession(receiptEncoded_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _SessionRouter.Contract.CloseSession(&_SessionRouter.TransactOpts, receiptEncoded_, signature_)
}

// OpenSession is a paid mutator transaction binding the contract method 0x1f71815e.
//
// Solidity: function openSession(uint256 amount_, bytes approvalEncoded_, bytes signature_) returns(bytes32)
func (_SessionRouter *SessionRouterTransactor) OpenSession(opts *bind.TransactOpts, amount_ *big.Int, approvalEncoded_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _SessionRouter.contract.Transact(opts, "openSession", amount_, approvalEncoded_, signature_)
}

// OpenSession is a paid mutator transaction binding the contract method 0x1f71815e.
//
// Solidity: function openSession(uint256 amount_, bytes approvalEncoded_, bytes signature_) returns(bytes32)
func (_SessionRouter *SessionRouterSession) OpenSession(amount_ *big.Int, approvalEncoded_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _SessionRouter.Contract.OpenSession(&_SessionRouter.TransactOpts, amount_, approvalEncoded_, signature_)
}

// OpenSession is a paid mutator transaction binding the contract method 0x1f71815e.
//
// Solidity: function openSession(uint256 amount_, bytes approvalEncoded_, bytes signature_) returns(bytes32)
func (_SessionRouter *SessionRouterTransactorSession) OpenSession(amount_ *big.Int, approvalEncoded_ []byte, signature_ []byte) (*types.Transaction, error) {
	return _SessionRouter.Contract.OpenSession(&_SessionRouter.TransactOpts, amount_, approvalEncoded_, signature_)
}

// SetPoolConfig is a paid mutator transaction binding the contract method 0xd7178753.
//
// Solidity: function setPoolConfig(uint256 index_, (uint256,uint256,uint128,uint128) pool_) returns()
func (_SessionRouter *SessionRouterTransactor) SetPoolConfig(opts *bind.TransactOpts, index_ *big.Int, pool_ ISessionStoragePool) (*types.Transaction, error) {
	return _SessionRouter.contract.Transact(opts, "setPoolConfig", index_, pool_)
}

// SetPoolConfig is a paid mutator transaction binding the contract method 0xd7178753.
//
// Solidity: function setPoolConfig(uint256 index_, (uint256,uint256,uint128,uint128) pool_) returns()
func (_SessionRouter *SessionRouterSession) SetPoolConfig(index_ *big.Int, pool_ ISessionStoragePool) (*types.Transaction, error) {
	return _SessionRouter.Contract.SetPoolConfig(&_SessionRouter.TransactOpts, index_, pool_)
}

// SetPoolConfig is a paid mutator transaction binding the contract method 0xd7178753.
//
// Solidity: function setPoolConfig(uint256 index_, (uint256,uint256,uint128,uint128) pool_) returns()
func (_SessionRouter *SessionRouterTransactorSession) SetPoolConfig(index_ *big.Int, pool_ ISessionStoragePool) (*types.Transaction, error) {
	return _SessionRouter.Contract.SetPoolConfig(&_SessionRouter.TransactOpts, index_, pool_)
}

// WithdrawUserStakes is a paid mutator transaction binding the contract method 0x8293617a.
//
// Solidity: function withdrawUserStakes(uint8 iterations_) returns()
func (_SessionRouter *SessionRouterTransactor) WithdrawUserStakes(opts *bind.TransactOpts, iterations_ uint8) (*types.Transaction, error) {
	return _SessionRouter.contract.Transact(opts, "withdrawUserStakes", iterations_)
}

// WithdrawUserStakes is a paid mutator transaction binding the contract method 0x8293617a.
//
// Solidity: function withdrawUserStakes(uint8 iterations_) returns()
func (_SessionRouter *SessionRouterSession) WithdrawUserStakes(iterations_ uint8) (*types.Transaction, error) {
	return _SessionRouter.Contract.WithdrawUserStakes(&_SessionRouter.TransactOpts, iterations_)
}

// WithdrawUserStakes is a paid mutator transaction binding the contract method 0x8293617a.
//
// Solidity: function withdrawUserStakes(uint8 iterations_) returns()
func (_SessionRouter *SessionRouterTransactorSession) WithdrawUserStakes(iterations_ uint8) (*types.Transaction, error) {
	return _SessionRouter.Contract.WithdrawUserStakes(&_SessionRouter.TransactOpts, iterations_)
}

// SessionRouterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SessionRouter contract.
type SessionRouterInitializedIterator struct {
	Event *SessionRouterInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SessionRouterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionRouterInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SessionRouterInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SessionRouterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionRouterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionRouterInitialized represents a Initialized event raised by the SessionRouter contract.
type SessionRouterInitialized struct {
	StorageSlot [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xdc73717d728bcfa015e8117438a65319aa06e979ca324afa6e1ea645c28ea15d.
//
// Solidity: event Initialized(bytes32 storageSlot)
func (_SessionRouter *SessionRouterFilterer) FilterInitialized(opts *bind.FilterOpts) (*SessionRouterInitializedIterator, error) {

	logs, sub, err := _SessionRouter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SessionRouterInitializedIterator{contract: _SessionRouter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xdc73717d728bcfa015e8117438a65319aa06e979ca324afa6e1ea645c28ea15d.
//
// Solidity: event Initialized(bytes32 storageSlot)
func (_SessionRouter *SessionRouterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SessionRouterInitialized) (event.Subscription, error) {

	logs, sub, err := _SessionRouter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionRouterInitialized)
				if err := _SessionRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xdc73717d728bcfa015e8117438a65319aa06e979ca324afa6e1ea645c28ea15d.
//
// Solidity: event Initialized(bytes32 storageSlot)
func (_SessionRouter *SessionRouterFilterer) ParseInitialized(log types.Log) (*SessionRouterInitialized, error) {
	event := new(SessionRouterInitialized)
	if err := _SessionRouter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionRouterSessionClosedIterator is returned from FilterSessionClosed and is used to iterate over the raw logs and unpacked data for SessionClosed events raised by the SessionRouter contract.
type SessionRouterSessionClosedIterator struct {
	Event *SessionRouterSessionClosed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SessionRouterSessionClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionRouterSessionClosed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SessionRouterSessionClosed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SessionRouterSessionClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionRouterSessionClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionRouterSessionClosed represents a SessionClosed event raised by the SessionRouter contract.
type SessionRouterSessionClosed struct {
	User       common.Address
	SessionId  [32]byte
	ProviderId common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSessionClosed is a free log retrieval operation binding the contract event 0x337fbb0a41a596db800dc836595a57815f967185e3596615c646f2455ac3914a.
//
// Solidity: event SessionClosed(address indexed user, bytes32 indexed sessionId, address indexed providerId)
func (_SessionRouter *SessionRouterFilterer) FilterSessionClosed(opts *bind.FilterOpts, user []common.Address, sessionId [][32]byte, providerId []common.Address) (*SessionRouterSessionClosedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}
	var providerIdRule []interface{}
	for _, providerIdItem := range providerId {
		providerIdRule = append(providerIdRule, providerIdItem)
	}

	logs, sub, err := _SessionRouter.contract.FilterLogs(opts, "SessionClosed", userRule, sessionIdRule, providerIdRule)
	if err != nil {
		return nil, err
	}
	return &SessionRouterSessionClosedIterator{contract: _SessionRouter.contract, event: "SessionClosed", logs: logs, sub: sub}, nil
}

// WatchSessionClosed is a free log subscription operation binding the contract event 0x337fbb0a41a596db800dc836595a57815f967185e3596615c646f2455ac3914a.
//
// Solidity: event SessionClosed(address indexed user, bytes32 indexed sessionId, address indexed providerId)
func (_SessionRouter *SessionRouterFilterer) WatchSessionClosed(opts *bind.WatchOpts, sink chan<- *SessionRouterSessionClosed, user []common.Address, sessionId [][32]byte, providerId []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}
	var providerIdRule []interface{}
	for _, providerIdItem := range providerId {
		providerIdRule = append(providerIdRule, providerIdItem)
	}

	logs, sub, err := _SessionRouter.contract.WatchLogs(opts, "SessionClosed", userRule, sessionIdRule, providerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionRouterSessionClosed)
				if err := _SessionRouter.contract.UnpackLog(event, "SessionClosed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSessionClosed is a log parse operation binding the contract event 0x337fbb0a41a596db800dc836595a57815f967185e3596615c646f2455ac3914a.
//
// Solidity: event SessionClosed(address indexed user, bytes32 indexed sessionId, address indexed providerId)
func (_SessionRouter *SessionRouterFilterer) ParseSessionClosed(log types.Log) (*SessionRouterSessionClosed, error) {
	event := new(SessionRouterSessionClosed)
	if err := _SessionRouter.contract.UnpackLog(event, "SessionClosed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionRouterSessionOpenedIterator is returned from FilterSessionOpened and is used to iterate over the raw logs and unpacked data for SessionOpened events raised by the SessionRouter contract.
type SessionRouterSessionOpenedIterator struct {
	Event *SessionRouterSessionOpened // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SessionRouterSessionOpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionRouterSessionOpened)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SessionRouterSessionOpened)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SessionRouterSessionOpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionRouterSessionOpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionRouterSessionOpened represents a SessionOpened event raised by the SessionRouter contract.
type SessionRouterSessionOpened struct {
	User       common.Address
	SessionId  [32]byte
	ProviderId common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSessionOpened is a free log retrieval operation binding the contract event 0x2bd7c890baf595977d256a6e784512c873ac58ba612b4895dbb7f784bfbf4839.
//
// Solidity: event SessionOpened(address indexed user, bytes32 indexed sessionId, address indexed providerId)
func (_SessionRouter *SessionRouterFilterer) FilterSessionOpened(opts *bind.FilterOpts, user []common.Address, sessionId [][32]byte, providerId []common.Address) (*SessionRouterSessionOpenedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}
	var providerIdRule []interface{}
	for _, providerIdItem := range providerId {
		providerIdRule = append(providerIdRule, providerIdItem)
	}

	logs, sub, err := _SessionRouter.contract.FilterLogs(opts, "SessionOpened", userRule, sessionIdRule, providerIdRule)
	if err != nil {
		return nil, err
	}
	return &SessionRouterSessionOpenedIterator{contract: _SessionRouter.contract, event: "SessionOpened", logs: logs, sub: sub}, nil
}

// WatchSessionOpened is a free log subscription operation binding the contract event 0x2bd7c890baf595977d256a6e784512c873ac58ba612b4895dbb7f784bfbf4839.
//
// Solidity: event SessionOpened(address indexed user, bytes32 indexed sessionId, address indexed providerId)
func (_SessionRouter *SessionRouterFilterer) WatchSessionOpened(opts *bind.WatchOpts, sink chan<- *SessionRouterSessionOpened, user []common.Address, sessionId [][32]byte, providerId []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var sessionIdRule []interface{}
	for _, sessionIdItem := range sessionId {
		sessionIdRule = append(sessionIdRule, sessionIdItem)
	}
	var providerIdRule []interface{}
	for _, providerIdItem := range providerId {
		providerIdRule = append(providerIdRule, providerIdItem)
	}

	logs, sub, err := _SessionRouter.contract.WatchLogs(opts, "SessionOpened", userRule, sessionIdRule, providerIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionRouterSessionOpened)
				if err := _SessionRouter.contract.UnpackLog(event, "SessionOpened", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSessionOpened is a log parse operation binding the contract event 0x2bd7c890baf595977d256a6e784512c873ac58ba612b4895dbb7f784bfbf4839.
//
// Solidity: event SessionOpened(address indexed user, bytes32 indexed sessionId, address indexed providerId)
func (_SessionRouter *SessionRouterFilterer) ParseSessionOpened(log types.Log) (*SessionRouterSessionOpened, error) {
	event := new(SessionRouterSessionOpened)
	if err := _SessionRouter.contract.UnpackLog(event, "SessionOpened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SessionRouterUserWithdrawnIterator is returned from FilterUserWithdrawn and is used to iterate over the raw logs and unpacked data for UserWithdrawn events raised by the SessionRouter contract.
type SessionRouterUserWithdrawnIterator struct {
	Event *SessionRouterUserWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SessionRouterUserWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SessionRouterUserWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SessionRouterUserWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SessionRouterUserWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SessionRouterUserWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SessionRouterUserWithdrawn represents a UserWithdrawn event raised by the SessionRouter contract.
type SessionRouterUserWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUserWithdrawn is a free log retrieval operation binding the contract event 0xe6b386172074b393dc04ed6cb1a352475ffad5dd8cebc76231a3b683141ea6fb.
//
// Solidity: event UserWithdrawn(address indexed user, uint256 amount_)
func (_SessionRouter *SessionRouterFilterer) FilterUserWithdrawn(opts *bind.FilterOpts, user []common.Address) (*SessionRouterUserWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SessionRouter.contract.FilterLogs(opts, "UserWithdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &SessionRouterUserWithdrawnIterator{contract: _SessionRouter.contract, event: "UserWithdrawn", logs: logs, sub: sub}, nil
}

// WatchUserWithdrawn is a free log subscription operation binding the contract event 0xe6b386172074b393dc04ed6cb1a352475ffad5dd8cebc76231a3b683141ea6fb.
//
// Solidity: event UserWithdrawn(address indexed user, uint256 amount_)
func (_SessionRouter *SessionRouterFilterer) WatchUserWithdrawn(opts *bind.WatchOpts, sink chan<- *SessionRouterUserWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _SessionRouter.contract.WatchLogs(opts, "UserWithdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SessionRouterUserWithdrawn)
				if err := _SessionRouter.contract.UnpackLog(event, "UserWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUserWithdrawn is a log parse operation binding the contract event 0xe6b386172074b393dc04ed6cb1a352475ffad5dd8cebc76231a3b683141ea6fb.
//
// Solidity: event UserWithdrawn(address indexed user, uint256 amount_)
func (_SessionRouter *SessionRouterFilterer) ParseUserWithdrawn(log types.Log) (*SessionRouterUserWithdrawn, error) {
	event := new(SessionRouterUserWithdrawn)
	if err := _SessionRouter.contract.UnpackLog(event, "UserWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
