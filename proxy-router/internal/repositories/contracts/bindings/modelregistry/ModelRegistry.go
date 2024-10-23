// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package modelregistry

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

// IModelStorageModel is an auto generated low-level Go binding around an user-defined struct.
type IModelStorageModel struct {
	IpfsCID   [32]byte
	Fee       *big.Int
	Stake     *big.Int
	Owner     common.Address
	Name      string
	Tags      []string
	CreatedAt *big.Int
	IsDeleted bool
}

// ModelRegistryMetaData contains all meta data concerning the ModelRegistry contract.
var ModelRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ModelHasActiveBids\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModelHasAlreadyDeregistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModelNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"}],\"name\":\"ModelStakeTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storageSlot\",\"type\":\"bytes32\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"modelId\",\"type\":\"bytes32\"}],\"name\":\"ModelDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"modelMinimumStake\",\"type\":\"uint256\"}],\"name\":\"ModelMinimumStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"modelId\",\"type\":\"bytes32\"}],\"name\":\"ModelRegisteredUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BIDS_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DIAMOND_OWNABLE_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODELS_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"__ModelRegistry_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getActiveModels\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bidId_\",\"type\":\"bytes32\"}],\"name\":\"getBid\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"modelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"createdAt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deletedAt\",\"type\":\"uint128\"}],\"internalType\":\"structIBidStorage.Bid\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"}],\"name\":\"getIsModelActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"}],\"name\":\"getModel\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ipfsCID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"tags\",\"type\":\"string[]\"},{\"internalType\":\"uint128\",\"name\":\"createdAt\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isDeleted\",\"type\":\"bool\"}],\"internalType\":\"structIModelStorage.Model\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getModelActiveBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getModelBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getModelIds\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModelMinimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getProviderActiveBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getProviderBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bidId_\",\"type\":\"bytes32\"}],\"name\":\"isBidActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"}],\"name\":\"modelDeregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"ipfsCID_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fee_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"tags_\",\"type\":\"string[]\"}],\"name\":\"modelRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"modelMinimumStake_\",\"type\":\"uint256\"}],\"name\":\"modelSetMinStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ModelRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ModelRegistryMetaData.ABI instead.
var ModelRegistryABI = ModelRegistryMetaData.ABI

// ModelRegistry is an auto generated Go binding around an Ethereum contract.
type ModelRegistry struct {
	ModelRegistryCaller     // Read-only binding to the contract
	ModelRegistryTransactor // Write-only binding to the contract
	ModelRegistryFilterer   // Log filterer for contract events
}

// ModelRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModelRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModelRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModelRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModelRegistrySession struct {
	Contract     *ModelRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModelRegistryCallerSession struct {
	Contract *ModelRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ModelRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModelRegistryTransactorSession struct {
	Contract     *ModelRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ModelRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModelRegistryRaw struct {
	Contract *ModelRegistry // Generic contract binding to access the raw methods on
}

// ModelRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModelRegistryCallerRaw struct {
	Contract *ModelRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ModelRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModelRegistryTransactorRaw struct {
	Contract *ModelRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModelRegistry creates a new instance of ModelRegistry, bound to a specific deployed contract.
func NewModelRegistry(address common.Address, backend bind.ContractBackend) (*ModelRegistry, error) {
	contract, err := bindModelRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ModelRegistry{ModelRegistryCaller: ModelRegistryCaller{contract: contract}, ModelRegistryTransactor: ModelRegistryTransactor{contract: contract}, ModelRegistryFilterer: ModelRegistryFilterer{contract: contract}}, nil
}

// NewModelRegistryCaller creates a new read-only instance of ModelRegistry, bound to a specific deployed contract.
func NewModelRegistryCaller(address common.Address, caller bind.ContractCaller) (*ModelRegistryCaller, error) {
	contract, err := bindModelRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModelRegistryCaller{contract: contract}, nil
}

// NewModelRegistryTransactor creates a new write-only instance of ModelRegistry, bound to a specific deployed contract.
func NewModelRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ModelRegistryTransactor, error) {
	contract, err := bindModelRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModelRegistryTransactor{contract: contract}, nil
}

// NewModelRegistryFilterer creates a new log filterer instance of ModelRegistry, bound to a specific deployed contract.
func NewModelRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ModelRegistryFilterer, error) {
	contract, err := bindModelRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModelRegistryFilterer{contract: contract}, nil
}

// bindModelRegistry binds a generic wrapper to an already deployed contract.
func bindModelRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ModelRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModelRegistry *ModelRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModelRegistry.Contract.ModelRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModelRegistry *ModelRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModelRegistry *ModelRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ModelRegistry *ModelRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ModelRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ModelRegistry *ModelRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ModelRegistry *ModelRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ModelRegistry.Contract.contract.Transact(opts, method, params...)
}

// BIDSSTORAGESLOT is a free data retrieval call binding the contract method 0x266ccff0.
//
// Solidity: function BIDS_STORAGE_SLOT() view returns(bytes32)
func (_ModelRegistry *ModelRegistryCaller) BIDSSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "BIDS_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BIDSSTORAGESLOT is a free data retrieval call binding the contract method 0x266ccff0.
//
// Solidity: function BIDS_STORAGE_SLOT() view returns(bytes32)
func (_ModelRegistry *ModelRegistrySession) BIDSSTORAGESLOT() ([32]byte, error) {
	return _ModelRegistry.Contract.BIDSSTORAGESLOT(&_ModelRegistry.CallOpts)
}

// BIDSSTORAGESLOT is a free data retrieval call binding the contract method 0x266ccff0.
//
// Solidity: function BIDS_STORAGE_SLOT() view returns(bytes32)
func (_ModelRegistry *ModelRegistryCallerSession) BIDSSTORAGESLOT() ([32]byte, error) {
	return _ModelRegistry.Contract.BIDSSTORAGESLOT(&_ModelRegistry.CallOpts)
}

// DIAMONDOWNABLESTORAGESLOT is a free data retrieval call binding the contract method 0x4ac3371e.
//
// Solidity: function DIAMOND_OWNABLE_STORAGE_SLOT() view returns(bytes32)
func (_ModelRegistry *ModelRegistryCaller) DIAMONDOWNABLESTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "DIAMOND_OWNABLE_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DIAMONDOWNABLESTORAGESLOT is a free data retrieval call binding the contract method 0x4ac3371e.
//
// Solidity: function DIAMOND_OWNABLE_STORAGE_SLOT() view returns(bytes32)
func (_ModelRegistry *ModelRegistrySession) DIAMONDOWNABLESTORAGESLOT() ([32]byte, error) {
	return _ModelRegistry.Contract.DIAMONDOWNABLESTORAGESLOT(&_ModelRegistry.CallOpts)
}

// DIAMONDOWNABLESTORAGESLOT is a free data retrieval call binding the contract method 0x4ac3371e.
//
// Solidity: function DIAMOND_OWNABLE_STORAGE_SLOT() view returns(bytes32)
func (_ModelRegistry *ModelRegistryCallerSession) DIAMONDOWNABLESTORAGESLOT() ([32]byte, error) {
	return _ModelRegistry.Contract.DIAMONDOWNABLESTORAGESLOT(&_ModelRegistry.CallOpts)
}

// MODELSSTORAGESLOT is a free data retrieval call binding the contract method 0x6f276c1e.
//
// Solidity: function MODELS_STORAGE_SLOT() view returns(bytes32)
func (_ModelRegistry *ModelRegistryCaller) MODELSSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "MODELS_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MODELSSTORAGESLOT is a free data retrieval call binding the contract method 0x6f276c1e.
//
// Solidity: function MODELS_STORAGE_SLOT() view returns(bytes32)
func (_ModelRegistry *ModelRegistrySession) MODELSSTORAGESLOT() ([32]byte, error) {
	return _ModelRegistry.Contract.MODELSSTORAGESLOT(&_ModelRegistry.CallOpts)
}

// MODELSSTORAGESLOT is a free data retrieval call binding the contract method 0x6f276c1e.
//
// Solidity: function MODELS_STORAGE_SLOT() view returns(bytes32)
func (_ModelRegistry *ModelRegistryCallerSession) MODELSSTORAGESLOT() ([32]byte, error) {
	return _ModelRegistry.Contract.MODELSSTORAGESLOT(&_ModelRegistry.CallOpts)
}

// GetActiveModels is a free data retrieval call binding the contract method 0xac59585c.
//
// Solidity: function getActiveModels(uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistryCaller) GetActiveModels(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "getActiveModels", offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetActiveModels is a free data retrieval call binding the contract method 0xac59585c.
//
// Solidity: function getActiveModels(uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistrySession) GetActiveModels(offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ModelRegistry.Contract.GetActiveModels(&_ModelRegistry.CallOpts, offset_, limit_)
}

// GetActiveModels is a free data retrieval call binding the contract method 0xac59585c.
//
// Solidity: function getActiveModels(uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistryCallerSession) GetActiveModels(offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ModelRegistry.Contract.GetActiveModels(&_ModelRegistry.CallOpts, offset_, limit_)
}

// GetBid is a free data retrieval call binding the contract method 0x91704e1e.
//
// Solidity: function getBid(bytes32 bidId_) view returns((address,bytes32,uint256,uint256,uint128,uint128))
func (_ModelRegistry *ModelRegistryCaller) GetBid(opts *bind.CallOpts, bidId_ [32]byte) (IBidStorageBid, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "getBid", bidId_)

	if err != nil {
		return *new(IBidStorageBid), err
	}

	out0 := *abi.ConvertType(out[0], new(IBidStorageBid)).(*IBidStorageBid)

	return out0, err

}

// GetBid is a free data retrieval call binding the contract method 0x91704e1e.
//
// Solidity: function getBid(bytes32 bidId_) view returns((address,bytes32,uint256,uint256,uint128,uint128))
func (_ModelRegistry *ModelRegistrySession) GetBid(bidId_ [32]byte) (IBidStorageBid, error) {
	return _ModelRegistry.Contract.GetBid(&_ModelRegistry.CallOpts, bidId_)
}

// GetBid is a free data retrieval call binding the contract method 0x91704e1e.
//
// Solidity: function getBid(bytes32 bidId_) view returns((address,bytes32,uint256,uint256,uint128,uint128))
func (_ModelRegistry *ModelRegistryCallerSession) GetBid(bidId_ [32]byte) (IBidStorageBid, error) {
	return _ModelRegistry.Contract.GetBid(&_ModelRegistry.CallOpts, bidId_)
}

// GetIsModelActive is a free data retrieval call binding the contract method 0xca74b5f3.
//
// Solidity: function getIsModelActive(bytes32 modelId_) view returns(bool)
func (_ModelRegistry *ModelRegistryCaller) GetIsModelActive(opts *bind.CallOpts, modelId_ [32]byte) (bool, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "getIsModelActive", modelId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsModelActive is a free data retrieval call binding the contract method 0xca74b5f3.
//
// Solidity: function getIsModelActive(bytes32 modelId_) view returns(bool)
func (_ModelRegistry *ModelRegistrySession) GetIsModelActive(modelId_ [32]byte) (bool, error) {
	return _ModelRegistry.Contract.GetIsModelActive(&_ModelRegistry.CallOpts, modelId_)
}

// GetIsModelActive is a free data retrieval call binding the contract method 0xca74b5f3.
//
// Solidity: function getIsModelActive(bytes32 modelId_) view returns(bool)
func (_ModelRegistry *ModelRegistryCallerSession) GetIsModelActive(modelId_ [32]byte) (bool, error) {
	return _ModelRegistry.Contract.GetIsModelActive(&_ModelRegistry.CallOpts, modelId_)
}

// GetModel is a free data retrieval call binding the contract method 0x21e7c498.
//
// Solidity: function getModel(bytes32 modelId_) view returns((bytes32,uint256,uint256,address,string,string[],uint128,bool))
func (_ModelRegistry *ModelRegistryCaller) GetModel(opts *bind.CallOpts, modelId_ [32]byte) (IModelStorageModel, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "getModel", modelId_)

	if err != nil {
		return *new(IModelStorageModel), err
	}

	out0 := *abi.ConvertType(out[0], new(IModelStorageModel)).(*IModelStorageModel)

	return out0, err

}

// GetModel is a free data retrieval call binding the contract method 0x21e7c498.
//
// Solidity: function getModel(bytes32 modelId_) view returns((bytes32,uint256,uint256,address,string,string[],uint128,bool))
func (_ModelRegistry *ModelRegistrySession) GetModel(modelId_ [32]byte) (IModelStorageModel, error) {
	return _ModelRegistry.Contract.GetModel(&_ModelRegistry.CallOpts, modelId_)
}

// GetModel is a free data retrieval call binding the contract method 0x21e7c498.
//
// Solidity: function getModel(bytes32 modelId_) view returns((bytes32,uint256,uint256,address,string,string[],uint128,bool))
func (_ModelRegistry *ModelRegistryCallerSession) GetModel(modelId_ [32]byte) (IModelStorageModel, error) {
	return _ModelRegistry.Contract.GetModel(&_ModelRegistry.CallOpts, modelId_)
}

// GetModelActiveBids is a free data retrieval call binding the contract method 0x8a683b6e.
//
// Solidity: function getModelActiveBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistryCaller) GetModelActiveBids(opts *bind.CallOpts, modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "getModelActiveBids", modelId_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetModelActiveBids is a free data retrieval call binding the contract method 0x8a683b6e.
//
// Solidity: function getModelActiveBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistrySession) GetModelActiveBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ModelRegistry.Contract.GetModelActiveBids(&_ModelRegistry.CallOpts, modelId_, offset_, limit_)
}

// GetModelActiveBids is a free data retrieval call binding the contract method 0x8a683b6e.
//
// Solidity: function getModelActiveBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistryCallerSession) GetModelActiveBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ModelRegistry.Contract.GetModelActiveBids(&_ModelRegistry.CallOpts, modelId_, offset_, limit_)
}

// GetModelBids is a free data retrieval call binding the contract method 0xfade17b1.
//
// Solidity: function getModelBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistryCaller) GetModelBids(opts *bind.CallOpts, modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "getModelBids", modelId_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetModelBids is a free data retrieval call binding the contract method 0xfade17b1.
//
// Solidity: function getModelBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistrySession) GetModelBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ModelRegistry.Contract.GetModelBids(&_ModelRegistry.CallOpts, modelId_, offset_, limit_)
}

// GetModelBids is a free data retrieval call binding the contract method 0xfade17b1.
//
// Solidity: function getModelBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistryCallerSession) GetModelBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ModelRegistry.Contract.GetModelBids(&_ModelRegistry.CallOpts, modelId_, offset_, limit_)
}

// GetModelIds is a free data retrieval call binding the contract method 0x08d0aab4.
//
// Solidity: function getModelIds(uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistryCaller) GetModelIds(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "getModelIds", offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetModelIds is a free data retrieval call binding the contract method 0x08d0aab4.
//
// Solidity: function getModelIds(uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistrySession) GetModelIds(offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ModelRegistry.Contract.GetModelIds(&_ModelRegistry.CallOpts, offset_, limit_)
}

// GetModelIds is a free data retrieval call binding the contract method 0x08d0aab4.
//
// Solidity: function getModelIds(uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistryCallerSession) GetModelIds(offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ModelRegistry.Contract.GetModelIds(&_ModelRegistry.CallOpts, offset_, limit_)
}

// GetModelMinimumStake is a free data retrieval call binding the contract method 0xf647ba3d.
//
// Solidity: function getModelMinimumStake() view returns(uint256)
func (_ModelRegistry *ModelRegistryCaller) GetModelMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "getModelMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModelMinimumStake is a free data retrieval call binding the contract method 0xf647ba3d.
//
// Solidity: function getModelMinimumStake() view returns(uint256)
func (_ModelRegistry *ModelRegistrySession) GetModelMinimumStake() (*big.Int, error) {
	return _ModelRegistry.Contract.GetModelMinimumStake(&_ModelRegistry.CallOpts)
}

// GetModelMinimumStake is a free data retrieval call binding the contract method 0xf647ba3d.
//
// Solidity: function getModelMinimumStake() view returns(uint256)
func (_ModelRegistry *ModelRegistryCallerSession) GetModelMinimumStake() (*big.Int, error) {
	return _ModelRegistry.Contract.GetModelMinimumStake(&_ModelRegistry.CallOpts)
}

// GetProviderActiveBids is a free data retrieval call binding the contract method 0xaf5b77ca.
//
// Solidity: function getProviderActiveBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistryCaller) GetProviderActiveBids(opts *bind.CallOpts, provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "getProviderActiveBids", provider_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProviderActiveBids is a free data retrieval call binding the contract method 0xaf5b77ca.
//
// Solidity: function getProviderActiveBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistrySession) GetProviderActiveBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ModelRegistry.Contract.GetProviderActiveBids(&_ModelRegistry.CallOpts, provider_, offset_, limit_)
}

// GetProviderActiveBids is a free data retrieval call binding the contract method 0xaf5b77ca.
//
// Solidity: function getProviderActiveBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistryCallerSession) GetProviderActiveBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ModelRegistry.Contract.GetProviderActiveBids(&_ModelRegistry.CallOpts, provider_, offset_, limit_)
}

// GetProviderBids is a free data retrieval call binding the contract method 0x59d435c4.
//
// Solidity: function getProviderBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistryCaller) GetProviderBids(opts *bind.CallOpts, provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "getProviderBids", provider_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProviderBids is a free data retrieval call binding the contract method 0x59d435c4.
//
// Solidity: function getProviderBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistrySession) GetProviderBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ModelRegistry.Contract.GetProviderBids(&_ModelRegistry.CallOpts, provider_, offset_, limit_)
}

// GetProviderBids is a free data retrieval call binding the contract method 0x59d435c4.
//
// Solidity: function getProviderBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_ModelRegistry *ModelRegistryCallerSession) GetProviderBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _ModelRegistry.Contract.GetProviderBids(&_ModelRegistry.CallOpts, provider_, offset_, limit_)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_ModelRegistry *ModelRegistryCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_ModelRegistry *ModelRegistrySession) GetToken() (common.Address, error) {
	return _ModelRegistry.Contract.GetToken(&_ModelRegistry.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_ModelRegistry *ModelRegistryCallerSession) GetToken() (common.Address, error) {
	return _ModelRegistry.Contract.GetToken(&_ModelRegistry.CallOpts)
}

// IsBidActive is a free data retrieval call binding the contract method 0x1345df58.
//
// Solidity: function isBidActive(bytes32 bidId_) view returns(bool)
func (_ModelRegistry *ModelRegistryCaller) IsBidActive(opts *bind.CallOpts, bidId_ [32]byte) (bool, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "isBidActive", bidId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBidActive is a free data retrieval call binding the contract method 0x1345df58.
//
// Solidity: function isBidActive(bytes32 bidId_) view returns(bool)
func (_ModelRegistry *ModelRegistrySession) IsBidActive(bidId_ [32]byte) (bool, error) {
	return _ModelRegistry.Contract.IsBidActive(&_ModelRegistry.CallOpts, bidId_)
}

// IsBidActive is a free data retrieval call binding the contract method 0x1345df58.
//
// Solidity: function isBidActive(bytes32 bidId_) view returns(bool)
func (_ModelRegistry *ModelRegistryCallerSession) IsBidActive(bidId_ [32]byte) (bool, error) {
	return _ModelRegistry.Contract.IsBidActive(&_ModelRegistry.CallOpts, bidId_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ModelRegistry *ModelRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ModelRegistry *ModelRegistrySession) Owner() (common.Address, error) {
	return _ModelRegistry.Contract.Owner(&_ModelRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ModelRegistry *ModelRegistryCallerSession) Owner() (common.Address, error) {
	return _ModelRegistry.Contract.Owner(&_ModelRegistry.CallOpts)
}

// ModelRegistryInit is a paid mutator transaction binding the contract method 0xd69bdf30.
//
// Solidity: function __ModelRegistry_init() returns()
func (_ModelRegistry *ModelRegistryTransactor) ModelRegistryInit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ModelRegistry.contract.Transact(opts, "__ModelRegistry_init")
}

// ModelRegistryInit is a paid mutator transaction binding the contract method 0xd69bdf30.
//
// Solidity: function __ModelRegistry_init() returns()
func (_ModelRegistry *ModelRegistrySession) ModelRegistryInit() (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelRegistryInit(&_ModelRegistry.TransactOpts)
}

// ModelRegistryInit is a paid mutator transaction binding the contract method 0xd69bdf30.
//
// Solidity: function __ModelRegistry_init() returns()
func (_ModelRegistry *ModelRegistryTransactorSession) ModelRegistryInit() (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelRegistryInit(&_ModelRegistry.TransactOpts)
}

// ModelDeregister is a paid mutator transaction binding the contract method 0xd5a245f1.
//
// Solidity: function modelDeregister(bytes32 modelId_) returns()
func (_ModelRegistry *ModelRegistryTransactor) ModelDeregister(opts *bind.TransactOpts, modelId_ [32]byte) (*types.Transaction, error) {
	return _ModelRegistry.contract.Transact(opts, "modelDeregister", modelId_)
}

// ModelDeregister is a paid mutator transaction binding the contract method 0xd5a245f1.
//
// Solidity: function modelDeregister(bytes32 modelId_) returns()
func (_ModelRegistry *ModelRegistrySession) ModelDeregister(modelId_ [32]byte) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelDeregister(&_ModelRegistry.TransactOpts, modelId_)
}

// ModelDeregister is a paid mutator transaction binding the contract method 0xd5a245f1.
//
// Solidity: function modelDeregister(bytes32 modelId_) returns()
func (_ModelRegistry *ModelRegistryTransactorSession) ModelDeregister(modelId_ [32]byte) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelDeregister(&_ModelRegistry.TransactOpts, modelId_)
}

// ModelRegister is a paid mutator transaction binding the contract method 0xec2581a1.
//
// Solidity: function modelRegister(bytes32 modelId_, bytes32 ipfsCID_, uint256 fee_, uint256 amount_, string name_, string[] tags_) returns()
func (_ModelRegistry *ModelRegistryTransactor) ModelRegister(opts *bind.TransactOpts, modelId_ [32]byte, ipfsCID_ [32]byte, fee_ *big.Int, amount_ *big.Int, name_ string, tags_ []string) (*types.Transaction, error) {
	return _ModelRegistry.contract.Transact(opts, "modelRegister", modelId_, ipfsCID_, fee_, amount_, name_, tags_)
}

// ModelRegister is a paid mutator transaction binding the contract method 0xec2581a1.
//
// Solidity: function modelRegister(bytes32 modelId_, bytes32 ipfsCID_, uint256 fee_, uint256 amount_, string name_, string[] tags_) returns()
func (_ModelRegistry *ModelRegistrySession) ModelRegister(modelId_ [32]byte, ipfsCID_ [32]byte, fee_ *big.Int, amount_ *big.Int, name_ string, tags_ []string) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelRegister(&_ModelRegistry.TransactOpts, modelId_, ipfsCID_, fee_, amount_, name_, tags_)
}

// ModelRegister is a paid mutator transaction binding the contract method 0xec2581a1.
//
// Solidity: function modelRegister(bytes32 modelId_, bytes32 ipfsCID_, uint256 fee_, uint256 amount_, string name_, string[] tags_) returns()
func (_ModelRegistry *ModelRegistryTransactorSession) ModelRegister(modelId_ [32]byte, ipfsCID_ [32]byte, fee_ *big.Int, amount_ *big.Int, name_ string, tags_ []string) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelRegister(&_ModelRegistry.TransactOpts, modelId_, ipfsCID_, fee_, amount_, name_, tags_)
}

// ModelSetMinStake is a paid mutator transaction binding the contract method 0x78998329.
//
// Solidity: function modelSetMinStake(uint256 modelMinimumStake_) returns()
func (_ModelRegistry *ModelRegistryTransactor) ModelSetMinStake(opts *bind.TransactOpts, modelMinimumStake_ *big.Int) (*types.Transaction, error) {
	return _ModelRegistry.contract.Transact(opts, "modelSetMinStake", modelMinimumStake_)
}

// ModelSetMinStake is a paid mutator transaction binding the contract method 0x78998329.
//
// Solidity: function modelSetMinStake(uint256 modelMinimumStake_) returns()
func (_ModelRegistry *ModelRegistrySession) ModelSetMinStake(modelMinimumStake_ *big.Int) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelSetMinStake(&_ModelRegistry.TransactOpts, modelMinimumStake_)
}

// ModelSetMinStake is a paid mutator transaction binding the contract method 0x78998329.
//
// Solidity: function modelSetMinStake(uint256 modelMinimumStake_) returns()
func (_ModelRegistry *ModelRegistryTransactorSession) ModelSetMinStake(modelMinimumStake_ *big.Int) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelSetMinStake(&_ModelRegistry.TransactOpts, modelMinimumStake_)
}

// ModelRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ModelRegistry contract.
type ModelRegistryInitializedIterator struct {
	Event *ModelRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *ModelRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelRegistryInitialized)
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
		it.Event = new(ModelRegistryInitialized)
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
func (it *ModelRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelRegistryInitialized represents a Initialized event raised by the ModelRegistry contract.
type ModelRegistryInitialized struct {
	StorageSlot [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xdc73717d728bcfa015e8117438a65319aa06e979ca324afa6e1ea645c28ea15d.
//
// Solidity: event Initialized(bytes32 storageSlot)
func (_ModelRegistry *ModelRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ModelRegistryInitializedIterator, error) {

	logs, sub, err := _ModelRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ModelRegistryInitializedIterator{contract: _ModelRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xdc73717d728bcfa015e8117438a65319aa06e979ca324afa6e1ea645c28ea15d.
//
// Solidity: event Initialized(bytes32 storageSlot)
func (_ModelRegistry *ModelRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ModelRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _ModelRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelRegistryInitialized)
				if err := _ModelRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ModelRegistry *ModelRegistryFilterer) ParseInitialized(log types.Log) (*ModelRegistryInitialized, error) {
	event := new(ModelRegistryInitialized)
	if err := _ModelRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelRegistryModelDeregisteredIterator is returned from FilterModelDeregistered and is used to iterate over the raw logs and unpacked data for ModelDeregistered events raised by the ModelRegistry contract.
type ModelRegistryModelDeregisteredIterator struct {
	Event *ModelRegistryModelDeregistered // Event containing the contract specifics and raw log

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
func (it *ModelRegistryModelDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelRegistryModelDeregistered)
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
		it.Event = new(ModelRegistryModelDeregistered)
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
func (it *ModelRegistryModelDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelRegistryModelDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelRegistryModelDeregistered represents a ModelDeregistered event raised by the ModelRegistry contract.
type ModelRegistryModelDeregistered struct {
	Owner   common.Address
	ModelId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelDeregistered is a free log retrieval operation binding the contract event 0x79a9f3017f26694a70f688c1e37f4add042a050660c62fc8351f760b153b888b.
//
// Solidity: event ModelDeregistered(address indexed owner, bytes32 indexed modelId)
func (_ModelRegistry *ModelRegistryFilterer) FilterModelDeregistered(opts *bind.FilterOpts, owner []common.Address, modelId [][32]byte) (*ModelRegistryModelDeregisteredIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _ModelRegistry.contract.FilterLogs(opts, "ModelDeregistered", ownerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &ModelRegistryModelDeregisteredIterator{contract: _ModelRegistry.contract, event: "ModelDeregistered", logs: logs, sub: sub}, nil
}

// WatchModelDeregistered is a free log subscription operation binding the contract event 0x79a9f3017f26694a70f688c1e37f4add042a050660c62fc8351f760b153b888b.
//
// Solidity: event ModelDeregistered(address indexed owner, bytes32 indexed modelId)
func (_ModelRegistry *ModelRegistryFilterer) WatchModelDeregistered(opts *bind.WatchOpts, sink chan<- *ModelRegistryModelDeregistered, owner []common.Address, modelId [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _ModelRegistry.contract.WatchLogs(opts, "ModelDeregistered", ownerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelRegistryModelDeregistered)
				if err := _ModelRegistry.contract.UnpackLog(event, "ModelDeregistered", log); err != nil {
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

// ParseModelDeregistered is a log parse operation binding the contract event 0x79a9f3017f26694a70f688c1e37f4add042a050660c62fc8351f760b153b888b.
//
// Solidity: event ModelDeregistered(address indexed owner, bytes32 indexed modelId)
func (_ModelRegistry *ModelRegistryFilterer) ParseModelDeregistered(log types.Log) (*ModelRegistryModelDeregistered, error) {
	event := new(ModelRegistryModelDeregistered)
	if err := _ModelRegistry.contract.UnpackLog(event, "ModelDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelRegistryModelMinimumStakeUpdatedIterator is returned from FilterModelMinimumStakeUpdated and is used to iterate over the raw logs and unpacked data for ModelMinimumStakeUpdated events raised by the ModelRegistry contract.
type ModelRegistryModelMinimumStakeUpdatedIterator struct {
	Event *ModelRegistryModelMinimumStakeUpdated // Event containing the contract specifics and raw log

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
func (it *ModelRegistryModelMinimumStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelRegistryModelMinimumStakeUpdated)
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
		it.Event = new(ModelRegistryModelMinimumStakeUpdated)
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
func (it *ModelRegistryModelMinimumStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelRegistryModelMinimumStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelRegistryModelMinimumStakeUpdated represents a ModelMinimumStakeUpdated event raised by the ModelRegistry contract.
type ModelRegistryModelMinimumStakeUpdated struct {
	ModelMinimumStake *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterModelMinimumStakeUpdated is a free log retrieval operation binding the contract event 0x136e2b2828baa18257e8eef3c26df62481c7b16303ca0ebe4202865df5e97d7e.
//
// Solidity: event ModelMinimumStakeUpdated(uint256 modelMinimumStake)
func (_ModelRegistry *ModelRegistryFilterer) FilterModelMinimumStakeUpdated(opts *bind.FilterOpts) (*ModelRegistryModelMinimumStakeUpdatedIterator, error) {

	logs, sub, err := _ModelRegistry.contract.FilterLogs(opts, "ModelMinimumStakeUpdated")
	if err != nil {
		return nil, err
	}
	return &ModelRegistryModelMinimumStakeUpdatedIterator{contract: _ModelRegistry.contract, event: "ModelMinimumStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchModelMinimumStakeUpdated is a free log subscription operation binding the contract event 0x136e2b2828baa18257e8eef3c26df62481c7b16303ca0ebe4202865df5e97d7e.
//
// Solidity: event ModelMinimumStakeUpdated(uint256 modelMinimumStake)
func (_ModelRegistry *ModelRegistryFilterer) WatchModelMinimumStakeUpdated(opts *bind.WatchOpts, sink chan<- *ModelRegistryModelMinimumStakeUpdated) (event.Subscription, error) {

	logs, sub, err := _ModelRegistry.contract.WatchLogs(opts, "ModelMinimumStakeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelRegistryModelMinimumStakeUpdated)
				if err := _ModelRegistry.contract.UnpackLog(event, "ModelMinimumStakeUpdated", log); err != nil {
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

// ParseModelMinimumStakeUpdated is a log parse operation binding the contract event 0x136e2b2828baa18257e8eef3c26df62481c7b16303ca0ebe4202865df5e97d7e.
//
// Solidity: event ModelMinimumStakeUpdated(uint256 modelMinimumStake)
func (_ModelRegistry *ModelRegistryFilterer) ParseModelMinimumStakeUpdated(log types.Log) (*ModelRegistryModelMinimumStakeUpdated, error) {
	event := new(ModelRegistryModelMinimumStakeUpdated)
	if err := _ModelRegistry.contract.UnpackLog(event, "ModelMinimumStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ModelRegistryModelRegisteredUpdatedIterator is returned from FilterModelRegisteredUpdated and is used to iterate over the raw logs and unpacked data for ModelRegisteredUpdated events raised by the ModelRegistry contract.
type ModelRegistryModelRegisteredUpdatedIterator struct {
	Event *ModelRegistryModelRegisteredUpdated // Event containing the contract specifics and raw log

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
func (it *ModelRegistryModelRegisteredUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelRegistryModelRegisteredUpdated)
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
		it.Event = new(ModelRegistryModelRegisteredUpdated)
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
func (it *ModelRegistryModelRegisteredUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelRegistryModelRegisteredUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelRegistryModelRegisteredUpdated represents a ModelRegisteredUpdated event raised by the ModelRegistry contract.
type ModelRegistryModelRegisteredUpdated struct {
	Owner   common.Address
	ModelId [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterModelRegisteredUpdated is a free log retrieval operation binding the contract event 0xda9282a68d90c36d31d780a69a65da6c35191a5d8cdd37bd1a1a5aa9fb168e77.
//
// Solidity: event ModelRegisteredUpdated(address indexed owner, bytes32 indexed modelId)
func (_ModelRegistry *ModelRegistryFilterer) FilterModelRegisteredUpdated(opts *bind.FilterOpts, owner []common.Address, modelId [][32]byte) (*ModelRegistryModelRegisteredUpdatedIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _ModelRegistry.contract.FilterLogs(opts, "ModelRegisteredUpdated", ownerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &ModelRegistryModelRegisteredUpdatedIterator{contract: _ModelRegistry.contract, event: "ModelRegisteredUpdated", logs: logs, sub: sub}, nil
}

// WatchModelRegisteredUpdated is a free log subscription operation binding the contract event 0xda9282a68d90c36d31d780a69a65da6c35191a5d8cdd37bd1a1a5aa9fb168e77.
//
// Solidity: event ModelRegisteredUpdated(address indexed owner, bytes32 indexed modelId)
func (_ModelRegistry *ModelRegistryFilterer) WatchModelRegisteredUpdated(opts *bind.WatchOpts, sink chan<- *ModelRegistryModelRegisteredUpdated, owner []common.Address, modelId [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _ModelRegistry.contract.WatchLogs(opts, "ModelRegisteredUpdated", ownerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelRegistryModelRegisteredUpdated)
				if err := _ModelRegistry.contract.UnpackLog(event, "ModelRegisteredUpdated", log); err != nil {
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

// ParseModelRegisteredUpdated is a log parse operation binding the contract event 0xda9282a68d90c36d31d780a69a65da6c35191a5d8cdd37bd1a1a5aa9fb168e77.
//
// Solidity: event ModelRegisteredUpdated(address indexed owner, bytes32 indexed modelId)
func (_ModelRegistry *ModelRegistryFilterer) ParseModelRegisteredUpdated(log types.Log) (*ModelRegistryModelRegisteredUpdated, error) {
	event := new(ModelRegistryModelRegisteredUpdated)
	if err := _ModelRegistry.contract.UnpackLog(event, "ModelRegisteredUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
