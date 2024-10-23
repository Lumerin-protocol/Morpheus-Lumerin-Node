// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package providerregistry

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

// ProviderRegistryMetaData contains all meta data concerning the ProviderRegistry contract.
var ProviderRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProviderHasActiveBids\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProviderHasAlreadyDeregistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProviderNoStake\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProviderNotDeregistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProviderNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProviderNothingToWithdraw\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"ProviderStakeTooLow\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storageSlot\",\"type\":\"bytes32\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"ProviderDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"providerMinimumStake\",\"type\":\"uint256\"}],\"name\":\"ProviderMinimumStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"ProviderRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ProviderWithdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BIDS_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DIAMOND_OWNABLE_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVIDERS_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__ProviderRegistry_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getActiveProviders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bidId_\",\"type\":\"bytes32\"}],\"name\":\"getBid\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"modelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"createdAt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deletedAt\",\"type\":\"uint128\"}],\"internalType\":\"structIBidStorage.Bid\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"}],\"name\":\"getIsProviderActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getModelActiveBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getModelBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"}],\"name\":\"getProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"createdAt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"limitPeriodEnd\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"limitPeriodEarned\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isDeleted\",\"type\":\"bool\"}],\"internalType\":\"structIProviderStorage.Provider\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getProviderActiveBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getProviderBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProviderMinimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bidId_\",\"type\":\"bytes32\"}],\"name\":\"isBidActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"providerDeregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"endpoint_\",\"type\":\"string\"}],\"name\":\"providerRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"providerMinimumStake_\",\"type\":\"uint256\"}],\"name\":\"providerSetMinStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProviderRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProviderRegistryMetaData.ABI instead.
var ProviderRegistryABI = ProviderRegistryMetaData.ABI

// ProviderRegistry is an auto generated Go binding around an Ethereum contract.
type ProviderRegistry struct {
	ProviderRegistryCaller     // Read-only binding to the contract
	ProviderRegistryTransactor // Write-only binding to the contract
	ProviderRegistryFilterer   // Log filterer for contract events
}

// ProviderRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProviderRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProviderRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProviderRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProviderRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProviderRegistrySession struct {
	Contract     *ProviderRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProviderRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProviderRegistryCallerSession struct {
	Contract *ProviderRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ProviderRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProviderRegistryTransactorSession struct {
	Contract     *ProviderRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ProviderRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProviderRegistryRaw struct {
	Contract *ProviderRegistry // Generic contract binding to access the raw methods on
}

// ProviderRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProviderRegistryCallerRaw struct {
	Contract *ProviderRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ProviderRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProviderRegistryTransactorRaw struct {
	Contract *ProviderRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProviderRegistry creates a new instance of ProviderRegistry, bound to a specific deployed contract.
func NewProviderRegistry(address common.Address, backend bind.ContractBackend) (*ProviderRegistry, error) {
	contract, err := bindProviderRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistry{ProviderRegistryCaller: ProviderRegistryCaller{contract: contract}, ProviderRegistryTransactor: ProviderRegistryTransactor{contract: contract}, ProviderRegistryFilterer: ProviderRegistryFilterer{contract: contract}}, nil
}

// NewProviderRegistryCaller creates a new read-only instance of ProviderRegistry, bound to a specific deployed contract.
func NewProviderRegistryCaller(address common.Address, caller bind.ContractCaller) (*ProviderRegistryCaller, error) {
	contract, err := bindProviderRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryCaller{contract: contract}, nil
}

// NewProviderRegistryTransactor creates a new write-only instance of ProviderRegistry, bound to a specific deployed contract.
func NewProviderRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProviderRegistryTransactor, error) {
	contract, err := bindProviderRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryTransactor{contract: contract}, nil
}

// NewProviderRegistryFilterer creates a new log filterer instance of ProviderRegistry, bound to a specific deployed contract.
func NewProviderRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProviderRegistryFilterer, error) {
	contract, err := bindProviderRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryFilterer{contract: contract}, nil
}

// bindProviderRegistry binds a generic wrapper to an already deployed contract.
func bindProviderRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProviderRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProviderRegistry *ProviderRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProviderRegistry.Contract.ProviderRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProviderRegistry *ProviderRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.ProviderRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProviderRegistry *ProviderRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.ProviderRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProviderRegistry *ProviderRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProviderRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProviderRegistry *ProviderRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProviderRegistry *ProviderRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.contract.Transact(opts, method, params...)
}

// BIDSSTORAGESLOT is a free data retrieval call binding the contract method 0x266ccff0.
//
// Solidity: function BIDS_STORAGE_SLOT() view returns(bytes32)
func (_ProviderRegistry *ProviderRegistryCaller) BIDSSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "BIDS_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BIDSSTORAGESLOT is a free data retrieval call binding the contract method 0x266ccff0.
//
// Solidity: function BIDS_STORAGE_SLOT() view returns(bytes32)
func (_ProviderRegistry *ProviderRegistrySession) BIDSSTORAGESLOT() ([32]byte, error) {
	return _ProviderRegistry.Contract.BIDSSTORAGESLOT(&_ProviderRegistry.CallOpts)
}

// BIDSSTORAGESLOT is a free data retrieval call binding the contract method 0x266ccff0.
//
// Solidity: function BIDS_STORAGE_SLOT() view returns(bytes32)
func (_ProviderRegistry *ProviderRegistryCallerSession) BIDSSTORAGESLOT() ([32]byte, error) {
	return _ProviderRegistry.Contract.BIDSSTORAGESLOT(&_ProviderRegistry.CallOpts)
}

// DIAMONDOWNABLESTORAGESLOT is a free data retrieval call binding the contract method 0x4ac3371e.
//
// Solidity: function DIAMOND_OWNABLE_STORAGE_SLOT() view returns(bytes32)
func (_ProviderRegistry *ProviderRegistryCaller) DIAMONDOWNABLESTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "DIAMOND_OWNABLE_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DIAMONDOWNABLESTORAGESLOT is a free data retrieval call binding the contract method 0x4ac3371e.
//
// Solidity: function DIAMOND_OWNABLE_STORAGE_SLOT() view returns(bytes32)
func (_ProviderRegistry *ProviderRegistrySession) DIAMONDOWNABLESTORAGESLOT() ([32]byte, error) {
	return _ProviderRegistry.Contract.DIAMONDOWNABLESTORAGESLOT(&_ProviderRegistry.CallOpts)
}

// DIAMONDOWNABLESTORAGESLOT is a free data retrieval call binding the contract method 0x4ac3371e.
//
// Solidity: function DIAMOND_OWNABLE_STORAGE_SLOT() view returns(bytes32)
func (_ProviderRegistry *ProviderRegistryCallerSession) DIAMONDOWNABLESTORAGESLOT() ([32]byte, error) {
	return _ProviderRegistry.Contract.DIAMONDOWNABLESTORAGESLOT(&_ProviderRegistry.CallOpts)
}

// PROVIDERSSTORAGESLOT is a free data retrieval call binding the contract method 0xc51830f6.
//
// Solidity: function PROVIDERS_STORAGE_SLOT() view returns(bytes32)
func (_ProviderRegistry *ProviderRegistryCaller) PROVIDERSSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "PROVIDERS_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROVIDERSSTORAGESLOT is a free data retrieval call binding the contract method 0xc51830f6.
//
// Solidity: function PROVIDERS_STORAGE_SLOT() view returns(bytes32)
func (_ProviderRegistry *ProviderRegistrySession) PROVIDERSSTORAGESLOT() ([32]byte, error) {
	return _ProviderRegistry.Contract.PROVIDERSSTORAGESLOT(&_ProviderRegistry.CallOpts)
}

// PROVIDERSSTORAGESLOT is a free data retrieval call binding the contract method 0xc51830f6.
//
// Solidity: function PROVIDERS_STORAGE_SLOT() view returns(bytes32)
func (_ProviderRegistry *ProviderRegistryCallerSession) PROVIDERSSTORAGESLOT() ([32]byte, error) {
	return _ProviderRegistry.Contract.PROVIDERSSTORAGESLOT(&_ProviderRegistry.CallOpts)
}

// GetActiveProviders is a free data retrieval call binding the contract method 0xd5472642.
//
// Solidity: function getActiveProviders(uint256 offset_, uint256 limit_) view returns(address[])
func (_ProviderRegistry *ProviderRegistryCaller) GetActiveProviders(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "getActiveProviders", offset_, limit_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveProviders is a free data retrieval call binding the contract method 0xd5472642.
//
// Solidity: function getActiveProviders(uint256 offset_, uint256 limit_) view returns(address[])
func (_ProviderRegistry *ProviderRegistrySession) GetActiveProviders(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _ProviderRegistry.Contract.GetActiveProviders(&_ProviderRegistry.CallOpts, offset_, limit_)
}

// GetActiveProviders is a free data retrieval call binding the contract method 0xd5472642.
//
// Solidity: function getActiveProviders(uint256 offset_, uint256 limit_) view returns(address[])
func (_ProviderRegistry *ProviderRegistryCallerSession) GetActiveProviders(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _ProviderRegistry.Contract.GetActiveProviders(&_ProviderRegistry.CallOpts, offset_, limit_)
}

// GetBid is a free data retrieval call binding the contract method 0x91704e1e.
//
// Solidity: function getBid(bytes32 bidId_) view returns((address,bytes32,uint256,uint256,uint128,uint128))
func (_ProviderRegistry *ProviderRegistryCaller) GetBid(opts *bind.CallOpts, bidId_ [32]byte) (IBidStorageBid, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "getBid", bidId_)

	if err != nil {
		return *new(IBidStorageBid), err
	}

	out0 := *abi.ConvertType(out[0], new(IBidStorageBid)).(*IBidStorageBid)

	return out0, err

}

// GetBid is a free data retrieval call binding the contract method 0x91704e1e.
//
// Solidity: function getBid(bytes32 bidId_) view returns((address,bytes32,uint256,uint256,uint128,uint128))
func (_ProviderRegistry *ProviderRegistrySession) GetBid(bidId_ [32]byte) (IBidStorageBid, error) {
	return _ProviderRegistry.Contract.GetBid(&_ProviderRegistry.CallOpts, bidId_)
}

// GetBid is a free data retrieval call binding the contract method 0x91704e1e.
//
// Solidity: function getBid(bytes32 bidId_) view returns((address,bytes32,uint256,uint256,uint128,uint128))
func (_ProviderRegistry *ProviderRegistryCallerSession) GetBid(bidId_ [32]byte) (IBidStorageBid, error) {
	return _ProviderRegistry.Contract.GetBid(&_ProviderRegistry.CallOpts, bidId_)
}

// GetIsProviderActive is a free data retrieval call binding the contract method 0x63ef175d.
//
// Solidity: function getIsProviderActive(address provider_) view returns(bool)
func (_ProviderRegistry *ProviderRegistryCaller) GetIsProviderActive(opts *bind.CallOpts, provider_ common.Address) (bool, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "getIsProviderActive", provider_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsProviderActive is a free data retrieval call binding the contract method 0x63ef175d.
//
// Solidity: function getIsProviderActive(address provider_) view returns(bool)
func (_ProviderRegistry *ProviderRegistrySession) GetIsProviderActive(provider_ common.Address) (bool, error) {
	return _ProviderRegistry.Contract.GetIsProviderActive(&_ProviderRegistry.CallOpts, provider_)
}

// GetIsProviderActive is a free data retrieval call binding the contract method 0x63ef175d.
//
// Solidity: function getIsProviderActive(address provider_) view returns(bool)
func (_ProviderRegistry *ProviderRegistryCallerSession) GetIsProviderActive(provider_ common.Address) (bool, error) {
	return _ProviderRegistry.Contract.GetIsProviderActive(&_ProviderRegistry.CallOpts, provider_)
}

// GetModelActiveBids is a free data retrieval call binding the contract method 0x8a683b6e.
//
// Solidity: function getModelActiveBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ProviderRegistry *ProviderRegistryCaller) GetModelActiveBids(opts *bind.CallOpts, modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "getModelActiveBids", modelId_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetModelActiveBids is a free data retrieval call binding the contract method 0x8a683b6e.
//
// Solidity: function getModelActiveBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ProviderRegistry *ProviderRegistrySession) GetModelActiveBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ProviderRegistry.Contract.GetModelActiveBids(&_ProviderRegistry.CallOpts, modelId_, offset_, limit_)
}

// GetModelActiveBids is a free data retrieval call binding the contract method 0x8a683b6e.
//
// Solidity: function getModelActiveBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ProviderRegistry *ProviderRegistryCallerSession) GetModelActiveBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ProviderRegistry.Contract.GetModelActiveBids(&_ProviderRegistry.CallOpts, modelId_, offset_, limit_)
}

// GetModelBids is a free data retrieval call binding the contract method 0xfade17b1.
//
// Solidity: function getModelBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ProviderRegistry *ProviderRegistryCaller) GetModelBids(opts *bind.CallOpts, modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "getModelBids", modelId_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetModelBids is a free data retrieval call binding the contract method 0xfade17b1.
//
// Solidity: function getModelBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ProviderRegistry *ProviderRegistrySession) GetModelBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ProviderRegistry.Contract.GetModelBids(&_ProviderRegistry.CallOpts, modelId_, offset_, limit_)
}

// GetModelBids is a free data retrieval call binding the contract method 0xfade17b1.
//
// Solidity: function getModelBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ProviderRegistry *ProviderRegistryCallerSession) GetModelBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ProviderRegistry.Contract.GetModelBids(&_ProviderRegistry.CallOpts, modelId_, offset_, limit_)
}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address provider_) view returns((string,uint256,uint128,uint128,uint256,bool))
func (_ProviderRegistry *ProviderRegistryCaller) GetProvider(opts *bind.CallOpts, provider_ common.Address) (IProviderStorageProvider, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "getProvider", provider_)

	if err != nil {
		return *new(IProviderStorageProvider), err
	}

	out0 := *abi.ConvertType(out[0], new(IProviderStorageProvider)).(*IProviderStorageProvider)

	return out0, err

}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address provider_) view returns((string,uint256,uint128,uint128,uint256,bool))
func (_ProviderRegistry *ProviderRegistrySession) GetProvider(provider_ common.Address) (IProviderStorageProvider, error) {
	return _ProviderRegistry.Contract.GetProvider(&_ProviderRegistry.CallOpts, provider_)
}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address provider_) view returns((string,uint256,uint128,uint128,uint256,bool))
func (_ProviderRegistry *ProviderRegistryCallerSession) GetProvider(provider_ common.Address) (IProviderStorageProvider, error) {
	return _ProviderRegistry.Contract.GetProvider(&_ProviderRegistry.CallOpts, provider_)
}

// GetProviderActiveBids is a free data retrieval call binding the contract method 0xaf5b77ca.
//
// Solidity: function getProviderActiveBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ProviderRegistry *ProviderRegistryCaller) GetProviderActiveBids(opts *bind.CallOpts, provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "getProviderActiveBids", provider_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProviderActiveBids is a free data retrieval call binding the contract method 0xaf5b77ca.
//
// Solidity: function getProviderActiveBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ProviderRegistry *ProviderRegistrySession) GetProviderActiveBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ProviderRegistry.Contract.GetProviderActiveBids(&_ProviderRegistry.CallOpts, provider_, offset_, limit_)
}

// GetProviderActiveBids is a free data retrieval call binding the contract method 0xaf5b77ca.
//
// Solidity: function getProviderActiveBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ProviderRegistry *ProviderRegistryCallerSession) GetProviderActiveBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ProviderRegistry.Contract.GetProviderActiveBids(&_ProviderRegistry.CallOpts, provider_, offset_, limit_)
}

// GetProviderBids is a free data retrieval call binding the contract method 0x59d435c4.
//
// Solidity: function getProviderBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ProviderRegistry *ProviderRegistryCaller) GetProviderBids(opts *bind.CallOpts, provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "getProviderBids", provider_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProviderBids is a free data retrieval call binding the contract method 0x59d435c4.
//
// Solidity: function getProviderBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ProviderRegistry *ProviderRegistrySession) GetProviderBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ProviderRegistry.Contract.GetProviderBids(&_ProviderRegistry.CallOpts, provider_, offset_, limit_)
}

// GetProviderBids is a free data retrieval call binding the contract method 0x59d435c4.
//
// Solidity: function getProviderBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ProviderRegistry *ProviderRegistryCallerSession) GetProviderBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ProviderRegistry.Contract.GetProviderBids(&_ProviderRegistry.CallOpts, provider_, offset_, limit_)
}

// GetProviderMinimumStake is a free data retrieval call binding the contract method 0x53c029f6.
//
// Solidity: function getProviderMinimumStake() view returns(uint256)
func (_ProviderRegistry *ProviderRegistryCaller) GetProviderMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "getProviderMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProviderMinimumStake is a free data retrieval call binding the contract method 0x53c029f6.
//
// Solidity: function getProviderMinimumStake() view returns(uint256)
func (_ProviderRegistry *ProviderRegistrySession) GetProviderMinimumStake() (*big.Int, error) {
	return _ProviderRegistry.Contract.GetProviderMinimumStake(&_ProviderRegistry.CallOpts)
}

// GetProviderMinimumStake is a free data retrieval call binding the contract method 0x53c029f6.
//
// Solidity: function getProviderMinimumStake() view returns(uint256)
func (_ProviderRegistry *ProviderRegistryCallerSession) GetProviderMinimumStake() (*big.Int, error) {
	return _ProviderRegistry.Contract.GetProviderMinimumStake(&_ProviderRegistry.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_ProviderRegistry *ProviderRegistryCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_ProviderRegistry *ProviderRegistrySession) GetToken() (common.Address, error) {
	return _ProviderRegistry.Contract.GetToken(&_ProviderRegistry.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_ProviderRegistry *ProviderRegistryCallerSession) GetToken() (common.Address, error) {
	return _ProviderRegistry.Contract.GetToken(&_ProviderRegistry.CallOpts)
}

// IsBidActive is a free data retrieval call binding the contract method 0x1345df58.
//
// Solidity: function isBidActive(bytes32 bidId_) view returns(bool)
func (_ProviderRegistry *ProviderRegistryCaller) IsBidActive(opts *bind.CallOpts, bidId_ [32]byte) (bool, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "isBidActive", bidId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBidActive is a free data retrieval call binding the contract method 0x1345df58.
//
// Solidity: function isBidActive(bytes32 bidId_) view returns(bool)
func (_ProviderRegistry *ProviderRegistrySession) IsBidActive(bidId_ [32]byte) (bool, error) {
	return _ProviderRegistry.Contract.IsBidActive(&_ProviderRegistry.CallOpts, bidId_)
}

// IsBidActive is a free data retrieval call binding the contract method 0x1345df58.
//
// Solidity: function isBidActive(bytes32 bidId_) view returns(bool)
func (_ProviderRegistry *ProviderRegistryCallerSession) IsBidActive(bidId_ [32]byte) (bool, error) {
	return _ProviderRegistry.Contract.IsBidActive(&_ProviderRegistry.CallOpts, bidId_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProviderRegistry *ProviderRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProviderRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProviderRegistry *ProviderRegistrySession) Owner() (common.Address, error) {
	return _ProviderRegistry.Contract.Owner(&_ProviderRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProviderRegistry *ProviderRegistryCallerSession) Owner() (common.Address, error) {
	return _ProviderRegistry.Contract.Owner(&_ProviderRegistry.CallOpts)
}

// ProviderRegistryInit is a paid mutator transaction binding the contract method 0x5c7ce38b.
//
// Solidity: function __ProviderRegistry_init() returns()
func (_ProviderRegistry *ProviderRegistryTransactor) ProviderRegistryInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "__ProviderRegistry_init")
}

// ProviderRegistryInit is a paid mutator transaction binding the contract method 0x5c7ce38b.
//
// Solidity: function __ProviderRegistry_init() returns()
func (_ProviderRegistry *ProviderRegistrySession) ProviderRegistryInit() (*types.Transaction, error) {
	return _ProviderRegistry.Contract.ProviderRegistryInit(&_ProviderRegistry.TransactOpts)
}

// ProviderRegistryInit is a paid mutator transaction binding the contract method 0x5c7ce38b.
//
// Solidity: function __ProviderRegistry_init() returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) ProviderRegistryInit() (*types.Transaction, error) {
	return _ProviderRegistry.Contract.ProviderRegistryInit(&_ProviderRegistry.TransactOpts)
}

// ProviderDeregister is a paid mutator transaction binding the contract method 0x58e0bd1c.
//
// Solidity: function providerDeregister() returns()
func (_ProviderRegistry *ProviderRegistryTransactor) ProviderDeregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "providerDeregister")
}

// ProviderDeregister is a paid mutator transaction binding the contract method 0x58e0bd1c.
//
// Solidity: function providerDeregister() returns()
func (_ProviderRegistry *ProviderRegistrySession) ProviderDeregister() (*types.Transaction, error) {
	return _ProviderRegistry.Contract.ProviderDeregister(&_ProviderRegistry.TransactOpts)
}

// ProviderDeregister is a paid mutator transaction binding the contract method 0x58e0bd1c.
//
// Solidity: function providerDeregister() returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) ProviderDeregister() (*types.Transaction, error) {
	return _ProviderRegistry.Contract.ProviderDeregister(&_ProviderRegistry.TransactOpts)
}

// ProviderRegister is a paid mutator transaction binding the contract method 0x17028d92.
//
// Solidity: function providerRegister(uint256 amount_, string endpoint_) returns()
func (_ProviderRegistry *ProviderRegistryTransactor) ProviderRegister(opts *bind.TransactOpts, amount_ *big.Int, endpoint_ string) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "providerRegister", amount_, endpoint_)
}

// ProviderRegister is a paid mutator transaction binding the contract method 0x17028d92.
//
// Solidity: function providerRegister(uint256 amount_, string endpoint_) returns()
func (_ProviderRegistry *ProviderRegistrySession) ProviderRegister(amount_ *big.Int, endpoint_ string) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.ProviderRegister(&_ProviderRegistry.TransactOpts, amount_, endpoint_)
}

// ProviderRegister is a paid mutator transaction binding the contract method 0x17028d92.
//
// Solidity: function providerRegister(uint256 amount_, string endpoint_) returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) ProviderRegister(amount_ *big.Int, endpoint_ string) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.ProviderRegister(&_ProviderRegistry.TransactOpts, amount_, endpoint_)
}

// ProviderSetMinStake is a paid mutator transaction binding the contract method 0x0b7f94d6.
//
// Solidity: function providerSetMinStake(uint256 providerMinimumStake_) returns()
func (_ProviderRegistry *ProviderRegistryTransactor) ProviderSetMinStake(opts *bind.TransactOpts, providerMinimumStake_ *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.contract.Transact(opts, "providerSetMinStake", providerMinimumStake_)
}

// ProviderSetMinStake is a paid mutator transaction binding the contract method 0x0b7f94d6.
//
// Solidity: function providerSetMinStake(uint256 providerMinimumStake_) returns()
func (_ProviderRegistry *ProviderRegistrySession) ProviderSetMinStake(providerMinimumStake_ *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.ProviderSetMinStake(&_ProviderRegistry.TransactOpts, providerMinimumStake_)
}

// ProviderSetMinStake is a paid mutator transaction binding the contract method 0x0b7f94d6.
//
// Solidity: function providerSetMinStake(uint256 providerMinimumStake_) returns()
func (_ProviderRegistry *ProviderRegistryTransactorSession) ProviderSetMinStake(providerMinimumStake_ *big.Int) (*types.Transaction, error) {
	return _ProviderRegistry.Contract.ProviderSetMinStake(&_ProviderRegistry.TransactOpts, providerMinimumStake_)
}

// ProviderRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ProviderRegistry contract.
type ProviderRegistryInitializedIterator struct {
	Event *ProviderRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ProviderRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRegistryInitialized)
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
		it.Event = new(ProviderRegistryInitialized)
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
func (it *ProviderRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRegistryInitialized represents a Initialized event raised by the ProviderRegistry contract.
type ProviderRegistryInitialized struct {
	StorageSlot [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xdc73717d728bcfa015e8117438a65319aa06e979ca324afa6e1ea645c28ea15d.
//
// Solidity: event Initialized(bytes32 storageSlot)
func (_ProviderRegistry *ProviderRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProviderRegistryInitializedIterator, error) {

	logs, sub, err := _ProviderRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryInitializedIterator{contract: _ProviderRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xdc73717d728bcfa015e8117438a65319aa06e979ca324afa6e1ea645c28ea15d.
//
// Solidity: event Initialized(bytes32 storageSlot)
func (_ProviderRegistry *ProviderRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProviderRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ProviderRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRegistryInitialized)
				if err := _ProviderRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ProviderRegistry *ProviderRegistryFilterer) ParseInitialized(log types.Log) (*ProviderRegistryInitialized, error) {
	event := new(ProviderRegistryInitialized)
	if err := _ProviderRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderRegistryProviderDeregisteredIterator is returned from FilterProviderDeregistered and is used to iterate over the raw logs and unpacked data for ProviderDeregistered events raised by the ProviderRegistry contract.
type ProviderRegistryProviderDeregisteredIterator struct {
	Event *ProviderRegistryProviderDeregistered // Event containing the contract specifics and raw log

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
func (it *ProviderRegistryProviderDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRegistryProviderDeregistered)
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
		it.Event = new(ProviderRegistryProviderDeregistered)
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
func (it *ProviderRegistryProviderDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRegistryProviderDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRegistryProviderDeregistered represents a ProviderDeregistered event raised by the ProviderRegistry contract.
type ProviderRegistryProviderDeregistered struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderDeregistered is a free log retrieval operation binding the contract event 0xf04091b4a187e321a42001e46961e45b6a75b203fc6fb766b7e05505f6080abb.
//
// Solidity: event ProviderDeregistered(address indexed provider)
func (_ProviderRegistry *ProviderRegistryFilterer) FilterProviderDeregistered(opts *bind.FilterOpts, provider []common.Address) (*ProviderRegistryProviderDeregisteredIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.FilterLogs(opts, "ProviderDeregistered", providerRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryProviderDeregisteredIterator{contract: _ProviderRegistry.contract, event: "ProviderDeregistered", logs: logs, sub: sub}, nil
}

// WatchProviderDeregistered is a free log subscription operation binding the contract event 0xf04091b4a187e321a42001e46961e45b6a75b203fc6fb766b7e05505f6080abb.
//
// Solidity: event ProviderDeregistered(address indexed provider)
func (_ProviderRegistry *ProviderRegistryFilterer) WatchProviderDeregistered(opts *bind.WatchOpts, sink chan<- *ProviderRegistryProviderDeregistered, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.WatchLogs(opts, "ProviderDeregistered", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRegistryProviderDeregistered)
				if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderDeregistered", log); err != nil {
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

// ParseProviderDeregistered is a log parse operation binding the contract event 0xf04091b4a187e321a42001e46961e45b6a75b203fc6fb766b7e05505f6080abb.
//
// Solidity: event ProviderDeregistered(address indexed provider)
func (_ProviderRegistry *ProviderRegistryFilterer) ParseProviderDeregistered(log types.Log) (*ProviderRegistryProviderDeregistered, error) {
	event := new(ProviderRegistryProviderDeregistered)
	if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderRegistryProviderMinimumStakeUpdatedIterator is returned from FilterProviderMinimumStakeUpdated and is used to iterate over the raw logs and unpacked data for ProviderMinimumStakeUpdated events raised by the ProviderRegistry contract.
type ProviderRegistryProviderMinimumStakeUpdatedIterator struct {
	Event *ProviderRegistryProviderMinimumStakeUpdated // Event containing the contract specifics and raw log

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
func (it *ProviderRegistryProviderMinimumStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRegistryProviderMinimumStakeUpdated)
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
		it.Event = new(ProviderRegistryProviderMinimumStakeUpdated)
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
func (it *ProviderRegistryProviderMinimumStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRegistryProviderMinimumStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRegistryProviderMinimumStakeUpdated represents a ProviderMinimumStakeUpdated event raised by the ProviderRegistry contract.
type ProviderRegistryProviderMinimumStakeUpdated struct {
	ProviderMinimumStake *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterProviderMinimumStakeUpdated is a free log retrieval operation binding the contract event 0x4d6b08cda70533f8222fea5ffc794d3a1f5dcde2620c6fa74789ef647db19450.
//
// Solidity: event ProviderMinimumStakeUpdated(uint256 providerMinimumStake)
func (_ProviderRegistry *ProviderRegistryFilterer) FilterProviderMinimumStakeUpdated(opts *bind.FilterOpts) (*ProviderRegistryProviderMinimumStakeUpdatedIterator, error) {

	logs, sub, err := _ProviderRegistry.contract.FilterLogs(opts, "ProviderMinimumStakeUpdated")
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryProviderMinimumStakeUpdatedIterator{contract: _ProviderRegistry.contract, event: "ProviderMinimumStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchProviderMinimumStakeUpdated is a free log subscription operation binding the contract event 0x4d6b08cda70533f8222fea5ffc794d3a1f5dcde2620c6fa74789ef647db19450.
//
// Solidity: event ProviderMinimumStakeUpdated(uint256 providerMinimumStake)
func (_ProviderRegistry *ProviderRegistryFilterer) WatchProviderMinimumStakeUpdated(opts *bind.WatchOpts, sink chan<- *ProviderRegistryProviderMinimumStakeUpdated) (event.Subscription, error) {

	logs, sub, err := _ProviderRegistry.contract.WatchLogs(opts, "ProviderMinimumStakeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRegistryProviderMinimumStakeUpdated)
				if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderMinimumStakeUpdated", log); err != nil {
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

// ParseProviderMinimumStakeUpdated is a log parse operation binding the contract event 0x4d6b08cda70533f8222fea5ffc794d3a1f5dcde2620c6fa74789ef647db19450.
//
// Solidity: event ProviderMinimumStakeUpdated(uint256 providerMinimumStake)
func (_ProviderRegistry *ProviderRegistryFilterer) ParseProviderMinimumStakeUpdated(log types.Log) (*ProviderRegistryProviderMinimumStakeUpdated, error) {
	event := new(ProviderRegistryProviderMinimumStakeUpdated)
	if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderMinimumStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderRegistryProviderRegisteredIterator is returned from FilterProviderRegistered and is used to iterate over the raw logs and unpacked data for ProviderRegistered events raised by the ProviderRegistry contract.
type ProviderRegistryProviderRegisteredIterator struct {
	Event *ProviderRegistryProviderRegistered // Event containing the contract specifics and raw log

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
func (it *ProviderRegistryProviderRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRegistryProviderRegistered)
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
		it.Event = new(ProviderRegistryProviderRegistered)
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
func (it *ProviderRegistryProviderRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRegistryProviderRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRegistryProviderRegistered represents a ProviderRegistered event raised by the ProviderRegistry contract.
type ProviderRegistryProviderRegistered struct {
	Provider common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderRegistered is a free log retrieval operation binding the contract event 0x70abce74777b3838ae60a33a6b9a87d9d25532668fe4fea548554c55868579c0.
//
// Solidity: event ProviderRegistered(address indexed provider)
func (_ProviderRegistry *ProviderRegistryFilterer) FilterProviderRegistered(opts *bind.FilterOpts, provider []common.Address) (*ProviderRegistryProviderRegisteredIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.FilterLogs(opts, "ProviderRegistered", providerRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryProviderRegisteredIterator{contract: _ProviderRegistry.contract, event: "ProviderRegistered", logs: logs, sub: sub}, nil
}

// WatchProviderRegistered is a free log subscription operation binding the contract event 0x70abce74777b3838ae60a33a6b9a87d9d25532668fe4fea548554c55868579c0.
//
// Solidity: event ProviderRegistered(address indexed provider)
func (_ProviderRegistry *ProviderRegistryFilterer) WatchProviderRegistered(opts *bind.WatchOpts, sink chan<- *ProviderRegistryProviderRegistered, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.WatchLogs(opts, "ProviderRegistered", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRegistryProviderRegistered)
				if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderRegistered", log); err != nil {
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

// ParseProviderRegistered is a log parse operation binding the contract event 0x70abce74777b3838ae60a33a6b9a87d9d25532668fe4fea548554c55868579c0.
//
// Solidity: event ProviderRegistered(address indexed provider)
func (_ProviderRegistry *ProviderRegistryFilterer) ParseProviderRegistered(log types.Log) (*ProviderRegistryProviderRegistered, error) {
	event := new(ProviderRegistryProviderRegistered)
	if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProviderRegistryProviderWithdrawnIterator is returned from FilterProviderWithdrawn and is used to iterate over the raw logs and unpacked data for ProviderWithdrawn events raised by the ProviderRegistry contract.
type ProviderRegistryProviderWithdrawnIterator struct {
	Event *ProviderRegistryProviderWithdrawn // Event containing the contract specifics and raw log

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
func (it *ProviderRegistryProviderWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProviderRegistryProviderWithdrawn)
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
		it.Event = new(ProviderRegistryProviderWithdrawn)
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
func (it *ProviderRegistryProviderWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProviderRegistryProviderWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProviderRegistryProviderWithdrawn represents a ProviderWithdrawn event raised by the ProviderRegistry contract.
type ProviderRegistryProviderWithdrawn struct {
	Provider common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterProviderWithdrawn is a free log retrieval operation binding the contract event 0x61388fbc2ba003175d3018b8f13b002234cc4e203a332a4a6dadb96bc82c3db2.
//
// Solidity: event ProviderWithdrawn(address indexed provider, uint256 amount)
func (_ProviderRegistry *ProviderRegistryFilterer) FilterProviderWithdrawn(opts *bind.FilterOpts, provider []common.Address) (*ProviderRegistryProviderWithdrawnIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.FilterLogs(opts, "ProviderWithdrawn", providerRule)
	if err != nil {
		return nil, err
	}
	return &ProviderRegistryProviderWithdrawnIterator{contract: _ProviderRegistry.contract, event: "ProviderWithdrawn", logs: logs, sub: sub}, nil
}

// WatchProviderWithdrawn is a free log subscription operation binding the contract event 0x61388fbc2ba003175d3018b8f13b002234cc4e203a332a4a6dadb96bc82c3db2.
//
// Solidity: event ProviderWithdrawn(address indexed provider, uint256 amount)
func (_ProviderRegistry *ProviderRegistryFilterer) WatchProviderWithdrawn(opts *bind.WatchOpts, sink chan<- *ProviderRegistryProviderWithdrawn, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _ProviderRegistry.contract.WatchLogs(opts, "ProviderWithdrawn", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProviderRegistryProviderWithdrawn)
				if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderWithdrawn", log); err != nil {
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

// ParseProviderWithdrawn is a log parse operation binding the contract event 0x61388fbc2ba003175d3018b8f13b002234cc4e203a332a4a6dadb96bc82c3db2.
//
// Solidity: event ProviderWithdrawn(address indexed provider, uint256 amount)
func (_ProviderRegistry *ProviderRegistryFilterer) ParseProviderWithdrawn(log types.Log) (*ProviderRegistryProviderWithdrawn, error) {
	event := new(ProviderRegistryProviderWithdrawn)
	if err := _ProviderRegistry.contract.UnpackLog(event, "ProviderWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
