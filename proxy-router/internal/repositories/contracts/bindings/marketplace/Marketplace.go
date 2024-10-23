// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketplace

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

// IProviderStorageProvider is an auto generated low-level Go binding around an user-defined struct.
type IProviderStorageProvider struct {
	Endpoint          string
	Stake             *big.Int
	CreatedAt         *big.Int
	LimitPeriodEnd    *big.Int
	LimitPeriodEarned *big.Int
	IsDeleted         bool
}

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MarketplaceActiveBidNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketplaceModelNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MarketplaceProviderNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storageSlot\",\"type\":\"bytes32\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidFee\",\"type\":\"uint256\"}],\"name\":\"MaretplaceFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"modelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"MarketplaceBidDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"modelId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"MarketplaceBidPosted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BIDS_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DIAMOND_OWNABLE_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MARKET_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODELS_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROVIDERS_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"__Marketplace_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bidId_\",\"type\":\"bytes32\"}],\"name\":\"deleteModelBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getActiveModels\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getActiveProviders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bidId_\",\"type\":\"bytes32\"}],\"name\":\"getBid\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"modelId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"createdAt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deletedAt\",\"type\":\"uint128\"}],\"internalType\":\"structIBidStorage.Bid\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBidFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nonce_\",\"type\":\"uint256\"}],\"name\":\"getBidId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"}],\"name\":\"getIsModelActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"}],\"name\":\"getIsProviderActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"}],\"name\":\"getModel\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ipfsCID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"tags\",\"type\":\"string[]\"},{\"internalType\":\"uint128\",\"name\":\"createdAt\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isDeleted\",\"type\":\"bool\"}],\"internalType\":\"structIModelStorage.Model\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getModelActiveBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getModelBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getModelIds\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getModelMinimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"}],\"name\":\"getProvider\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"createdAt\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"limitPeriodEnd\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"limitPeriodEarned\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isDeleted\",\"type\":\"bool\"}],\"internalType\":\"structIProviderStorage.Provider\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getProviderActiveBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"offset_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit_\",\"type\":\"uint256\"}],\"name\":\"getProviderBids\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProviderMinimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"provider_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"}],\"name\":\"getProviderModelId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"bidId_\",\"type\":\"bytes32\"}],\"name\":\"isBidActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"pricePerSecond_\",\"type\":\"uint256\"}],\"name\":\"postModelBid\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"bidId\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidFee_\",\"type\":\"uint256\"}],\"name\":\"setMarketplaceBidFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// BIDSSTORAGESLOT is a free data retrieval call binding the contract method 0x266ccff0.
//
// Solidity: function BIDS_STORAGE_SLOT() view returns(bytes32)
func (_Marketplace *MarketplaceCaller) BIDSSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "BIDS_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BIDSSTORAGESLOT is a free data retrieval call binding the contract method 0x266ccff0.
//
// Solidity: function BIDS_STORAGE_SLOT() view returns(bytes32)
func (_Marketplace *MarketplaceSession) BIDSSTORAGESLOT() ([32]byte, error) {
	return _Marketplace.Contract.BIDSSTORAGESLOT(&_Marketplace.CallOpts)
}

// BIDSSTORAGESLOT is a free data retrieval call binding the contract method 0x266ccff0.
//
// Solidity: function BIDS_STORAGE_SLOT() view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) BIDSSTORAGESLOT() ([32]byte, error) {
	return _Marketplace.Contract.BIDSSTORAGESLOT(&_Marketplace.CallOpts)
}

// DIAMONDOWNABLESTORAGESLOT is a free data retrieval call binding the contract method 0x4ac3371e.
//
// Solidity: function DIAMOND_OWNABLE_STORAGE_SLOT() view returns(bytes32)
func (_Marketplace *MarketplaceCaller) DIAMONDOWNABLESTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "DIAMOND_OWNABLE_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DIAMONDOWNABLESTORAGESLOT is a free data retrieval call binding the contract method 0x4ac3371e.
//
// Solidity: function DIAMOND_OWNABLE_STORAGE_SLOT() view returns(bytes32)
func (_Marketplace *MarketplaceSession) DIAMONDOWNABLESTORAGESLOT() ([32]byte, error) {
	return _Marketplace.Contract.DIAMONDOWNABLESTORAGESLOT(&_Marketplace.CallOpts)
}

// DIAMONDOWNABLESTORAGESLOT is a free data retrieval call binding the contract method 0x4ac3371e.
//
// Solidity: function DIAMOND_OWNABLE_STORAGE_SLOT() view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) DIAMONDOWNABLESTORAGESLOT() ([32]byte, error) {
	return _Marketplace.Contract.DIAMONDOWNABLESTORAGESLOT(&_Marketplace.CallOpts)
}

// MARKETSTORAGESLOT is a free data retrieval call binding the contract method 0x2afa2c86.
//
// Solidity: function MARKET_STORAGE_SLOT() view returns(bytes32)
func (_Marketplace *MarketplaceCaller) MARKETSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "MARKET_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MARKETSTORAGESLOT is a free data retrieval call binding the contract method 0x2afa2c86.
//
// Solidity: function MARKET_STORAGE_SLOT() view returns(bytes32)
func (_Marketplace *MarketplaceSession) MARKETSTORAGESLOT() ([32]byte, error) {
	return _Marketplace.Contract.MARKETSTORAGESLOT(&_Marketplace.CallOpts)
}

// MARKETSTORAGESLOT is a free data retrieval call binding the contract method 0x2afa2c86.
//
// Solidity: function MARKET_STORAGE_SLOT() view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) MARKETSTORAGESLOT() ([32]byte, error) {
	return _Marketplace.Contract.MARKETSTORAGESLOT(&_Marketplace.CallOpts)
}

// MODELSSTORAGESLOT is a free data retrieval call binding the contract method 0x6f276c1e.
//
// Solidity: function MODELS_STORAGE_SLOT() view returns(bytes32)
func (_Marketplace *MarketplaceCaller) MODELSSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "MODELS_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MODELSSTORAGESLOT is a free data retrieval call binding the contract method 0x6f276c1e.
//
// Solidity: function MODELS_STORAGE_SLOT() view returns(bytes32)
func (_Marketplace *MarketplaceSession) MODELSSTORAGESLOT() ([32]byte, error) {
	return _Marketplace.Contract.MODELSSTORAGESLOT(&_Marketplace.CallOpts)
}

// MODELSSTORAGESLOT is a free data retrieval call binding the contract method 0x6f276c1e.
//
// Solidity: function MODELS_STORAGE_SLOT() view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) MODELSSTORAGESLOT() ([32]byte, error) {
	return _Marketplace.Contract.MODELSSTORAGESLOT(&_Marketplace.CallOpts)
}

// PROVIDERSSTORAGESLOT is a free data retrieval call binding the contract method 0xc51830f6.
//
// Solidity: function PROVIDERS_STORAGE_SLOT() view returns(bytes32)
func (_Marketplace *MarketplaceCaller) PROVIDERSSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "PROVIDERS_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROVIDERSSTORAGESLOT is a free data retrieval call binding the contract method 0xc51830f6.
//
// Solidity: function PROVIDERS_STORAGE_SLOT() view returns(bytes32)
func (_Marketplace *MarketplaceSession) PROVIDERSSTORAGESLOT() ([32]byte, error) {
	return _Marketplace.Contract.PROVIDERSSTORAGESLOT(&_Marketplace.CallOpts)
}

// PROVIDERSSTORAGESLOT is a free data retrieval call binding the contract method 0xc51830f6.
//
// Solidity: function PROVIDERS_STORAGE_SLOT() view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) PROVIDERSSTORAGESLOT() ([32]byte, error) {
	return _Marketplace.Contract.PROVIDERSSTORAGESLOT(&_Marketplace.CallOpts)
}

// GetActiveModels is a free data retrieval call binding the contract method 0xac59585c.
//
// Solidity: function getActiveModels(uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceCaller) GetActiveModels(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getActiveModels", offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetActiveModels is a free data retrieval call binding the contract method 0xac59585c.
//
// Solidity: function getActiveModels(uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceSession) GetActiveModels(offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _Marketplace.Contract.GetActiveModels(&_Marketplace.CallOpts, offset_, limit_)
}

// GetActiveModels is a free data retrieval call binding the contract method 0xac59585c.
//
// Solidity: function getActiveModels(uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceCallerSession) GetActiveModels(offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _Marketplace.Contract.GetActiveModels(&_Marketplace.CallOpts, offset_, limit_)
}

// GetActiveProviders is a free data retrieval call binding the contract method 0xd5472642.
//
// Solidity: function getActiveProviders(uint256 offset_, uint256 limit_) view returns(address[])
func (_Marketplace *MarketplaceCaller) GetActiveProviders(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getActiveProviders", offset_, limit_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveProviders is a free data retrieval call binding the contract method 0xd5472642.
//
// Solidity: function getActiveProviders(uint256 offset_, uint256 limit_) view returns(address[])
func (_Marketplace *MarketplaceSession) GetActiveProviders(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _Marketplace.Contract.GetActiveProviders(&_Marketplace.CallOpts, offset_, limit_)
}

// GetActiveProviders is a free data retrieval call binding the contract method 0xd5472642.
//
// Solidity: function getActiveProviders(uint256 offset_, uint256 limit_) view returns(address[])
func (_Marketplace *MarketplaceCallerSession) GetActiveProviders(offset_ *big.Int, limit_ *big.Int) ([]common.Address, error) {
	return _Marketplace.Contract.GetActiveProviders(&_Marketplace.CallOpts, offset_, limit_)
}

// GetBid is a free data retrieval call binding the contract method 0x91704e1e.
//
// Solidity: function getBid(bytes32 bidId_) view returns((address,bytes32,uint256,uint256,uint128,uint128))
func (_Marketplace *MarketplaceCaller) GetBid(opts *bind.CallOpts, bidId_ [32]byte) (IBidStorageBid, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getBid", bidId_)

	if err != nil {
		return *new(IBidStorageBid), err
	}

	out0 := *abi.ConvertType(out[0], new(IBidStorageBid)).(*IBidStorageBid)

	return out0, err

}

// GetBid is a free data retrieval call binding the contract method 0x91704e1e.
//
// Solidity: function getBid(bytes32 bidId_) view returns((address,bytes32,uint256,uint256,uint128,uint128))
func (_Marketplace *MarketplaceSession) GetBid(bidId_ [32]byte) (IBidStorageBid, error) {
	return _Marketplace.Contract.GetBid(&_Marketplace.CallOpts, bidId_)
}

// GetBid is a free data retrieval call binding the contract method 0x91704e1e.
//
// Solidity: function getBid(bytes32 bidId_) view returns((address,bytes32,uint256,uint256,uint128,uint128))
func (_Marketplace *MarketplaceCallerSession) GetBid(bidId_ [32]byte) (IBidStorageBid, error) {
	return _Marketplace.Contract.GetBid(&_Marketplace.CallOpts, bidId_)
}

// GetBidFee is a free data retrieval call binding the contract method 0x8dbb4647.
//
// Solidity: function getBidFee() view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetBidFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getBidFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBidFee is a free data retrieval call binding the contract method 0x8dbb4647.
//
// Solidity: function getBidFee() view returns(uint256)
func (_Marketplace *MarketplaceSession) GetBidFee() (*big.Int, error) {
	return _Marketplace.Contract.GetBidFee(&_Marketplace.CallOpts)
}

// GetBidFee is a free data retrieval call binding the contract method 0x8dbb4647.
//
// Solidity: function getBidFee() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetBidFee() (*big.Int, error) {
	return _Marketplace.Contract.GetBidFee(&_Marketplace.CallOpts)
}

// GetBidId is a free data retrieval call binding the contract method 0x747ddd5b.
//
// Solidity: function getBidId(address provider_, bytes32 modelId_, uint256 nonce_) pure returns(bytes32)
func (_Marketplace *MarketplaceCaller) GetBidId(opts *bind.CallOpts, provider_ common.Address, modelId_ [32]byte, nonce_ *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getBidId", provider_, modelId_, nonce_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBidId is a free data retrieval call binding the contract method 0x747ddd5b.
//
// Solidity: function getBidId(address provider_, bytes32 modelId_, uint256 nonce_) pure returns(bytes32)
func (_Marketplace *MarketplaceSession) GetBidId(provider_ common.Address, modelId_ [32]byte, nonce_ *big.Int) ([32]byte, error) {
	return _Marketplace.Contract.GetBidId(&_Marketplace.CallOpts, provider_, modelId_, nonce_)
}

// GetBidId is a free data retrieval call binding the contract method 0x747ddd5b.
//
// Solidity: function getBidId(address provider_, bytes32 modelId_, uint256 nonce_) pure returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) GetBidId(provider_ common.Address, modelId_ [32]byte, nonce_ *big.Int) ([32]byte, error) {
	return _Marketplace.Contract.GetBidId(&_Marketplace.CallOpts, provider_, modelId_, nonce_)
}

// GetFeeBalance is a free data retrieval call binding the contract method 0xd4c30ceb.
//
// Solidity: function getFeeBalance() view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetFeeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getFeeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeeBalance is a free data retrieval call binding the contract method 0xd4c30ceb.
//
// Solidity: function getFeeBalance() view returns(uint256)
func (_Marketplace *MarketplaceSession) GetFeeBalance() (*big.Int, error) {
	return _Marketplace.Contract.GetFeeBalance(&_Marketplace.CallOpts)
}

// GetFeeBalance is a free data retrieval call binding the contract method 0xd4c30ceb.
//
// Solidity: function getFeeBalance() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetFeeBalance() (*big.Int, error) {
	return _Marketplace.Contract.GetFeeBalance(&_Marketplace.CallOpts)
}

// GetIsModelActive is a free data retrieval call binding the contract method 0xca74b5f3.
//
// Solidity: function getIsModelActive(bytes32 modelId_) view returns(bool)
func (_Marketplace *MarketplaceCaller) GetIsModelActive(opts *bind.CallOpts, modelId_ [32]byte) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getIsModelActive", modelId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsModelActive is a free data retrieval call binding the contract method 0xca74b5f3.
//
// Solidity: function getIsModelActive(bytes32 modelId_) view returns(bool)
func (_Marketplace *MarketplaceSession) GetIsModelActive(modelId_ [32]byte) (bool, error) {
	return _Marketplace.Contract.GetIsModelActive(&_Marketplace.CallOpts, modelId_)
}

// GetIsModelActive is a free data retrieval call binding the contract method 0xca74b5f3.
//
// Solidity: function getIsModelActive(bytes32 modelId_) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) GetIsModelActive(modelId_ [32]byte) (bool, error) {
	return _Marketplace.Contract.GetIsModelActive(&_Marketplace.CallOpts, modelId_)
}

// GetIsProviderActive is a free data retrieval call binding the contract method 0x63ef175d.
//
// Solidity: function getIsProviderActive(address provider_) view returns(bool)
func (_Marketplace *MarketplaceCaller) GetIsProviderActive(opts *bind.CallOpts, provider_ common.Address) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getIsProviderActive", provider_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsProviderActive is a free data retrieval call binding the contract method 0x63ef175d.
//
// Solidity: function getIsProviderActive(address provider_) view returns(bool)
func (_Marketplace *MarketplaceSession) GetIsProviderActive(provider_ common.Address) (bool, error) {
	return _Marketplace.Contract.GetIsProviderActive(&_Marketplace.CallOpts, provider_)
}

// GetIsProviderActive is a free data retrieval call binding the contract method 0x63ef175d.
//
// Solidity: function getIsProviderActive(address provider_) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) GetIsProviderActive(provider_ common.Address) (bool, error) {
	return _Marketplace.Contract.GetIsProviderActive(&_Marketplace.CallOpts, provider_)
}

// GetModel is a free data retrieval call binding the contract method 0x21e7c498.
//
// Solidity: function getModel(bytes32 modelId_) view returns((bytes32,uint256,uint256,address,string,string[],uint128,bool))
func (_Marketplace *MarketplaceCaller) GetModel(opts *bind.CallOpts, modelId_ [32]byte) (IModelStorageModel, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getModel", modelId_)

	if err != nil {
		return *new(IModelStorageModel), err
	}

	out0 := *abi.ConvertType(out[0], new(IModelStorageModel)).(*IModelStorageModel)

	return out0, err

}

// GetModel is a free data retrieval call binding the contract method 0x21e7c498.
//
// Solidity: function getModel(bytes32 modelId_) view returns((bytes32,uint256,uint256,address,string,string[],uint128,bool))
func (_Marketplace *MarketplaceSession) GetModel(modelId_ [32]byte) (IModelStorageModel, error) {
	return _Marketplace.Contract.GetModel(&_Marketplace.CallOpts, modelId_)
}

// GetModel is a free data retrieval call binding the contract method 0x21e7c498.
//
// Solidity: function getModel(bytes32 modelId_) view returns((bytes32,uint256,uint256,address,string,string[],uint128,bool))
func (_Marketplace *MarketplaceCallerSession) GetModel(modelId_ [32]byte) (IModelStorageModel, error) {
	return _Marketplace.Contract.GetModel(&_Marketplace.CallOpts, modelId_)
}

// GetModelActiveBids is a free data retrieval call binding the contract method 0x8a683b6e.
//
// Solidity: function getModelActiveBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceCaller) GetModelActiveBids(opts *bind.CallOpts, modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getModelActiveBids", modelId_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetModelActiveBids is a free data retrieval call binding the contract method 0x8a683b6e.
//
// Solidity: function getModelActiveBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceSession) GetModelActiveBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _Marketplace.Contract.GetModelActiveBids(&_Marketplace.CallOpts, modelId_, offset_, limit_)
}

// GetModelActiveBids is a free data retrieval call binding the contract method 0x8a683b6e.
//
// Solidity: function getModelActiveBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceCallerSession) GetModelActiveBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _Marketplace.Contract.GetModelActiveBids(&_Marketplace.CallOpts, modelId_, offset_, limit_)
}

// GetModelBids is a free data retrieval call binding the contract method 0xfade17b1.
//
// Solidity: function getModelBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceCaller) GetModelBids(opts *bind.CallOpts, modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getModelBids", modelId_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetModelBids is a free data retrieval call binding the contract method 0xfade17b1.
//
// Solidity: function getModelBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceSession) GetModelBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _Marketplace.Contract.GetModelBids(&_Marketplace.CallOpts, modelId_, offset_, limit_)
}

// GetModelBids is a free data retrieval call binding the contract method 0xfade17b1.
//
// Solidity: function getModelBids(bytes32 modelId_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceCallerSession) GetModelBids(modelId_ [32]byte, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _Marketplace.Contract.GetModelBids(&_Marketplace.CallOpts, modelId_, offset_, limit_)
}

// GetModelIds is a free data retrieval call binding the contract method 0x08d0aab4.
//
// Solidity: function getModelIds(uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceCaller) GetModelIds(opts *bind.CallOpts, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getModelIds", offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetModelIds is a free data retrieval call binding the contract method 0x08d0aab4.
//
// Solidity: function getModelIds(uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceSession) GetModelIds(offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _Marketplace.Contract.GetModelIds(&_Marketplace.CallOpts, offset_, limit_)
}

// GetModelIds is a free data retrieval call binding the contract method 0x08d0aab4.
//
// Solidity: function getModelIds(uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceCallerSession) GetModelIds(offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _Marketplace.Contract.GetModelIds(&_Marketplace.CallOpts, offset_, limit_)
}

// GetModelMinimumStake is a free data retrieval call binding the contract method 0xf647ba3d.
//
// Solidity: function getModelMinimumStake() view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetModelMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getModelMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetModelMinimumStake is a free data retrieval call binding the contract method 0xf647ba3d.
//
// Solidity: function getModelMinimumStake() view returns(uint256)
func (_Marketplace *MarketplaceSession) GetModelMinimumStake() (*big.Int, error) {
	return _Marketplace.Contract.GetModelMinimumStake(&_Marketplace.CallOpts)
}

// GetModelMinimumStake is a free data retrieval call binding the contract method 0xf647ba3d.
//
// Solidity: function getModelMinimumStake() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetModelMinimumStake() (*big.Int, error) {
	return _Marketplace.Contract.GetModelMinimumStake(&_Marketplace.CallOpts)
}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address provider_) view returns((string,uint256,uint128,uint128,uint256,bool))
func (_Marketplace *MarketplaceCaller) GetProvider(opts *bind.CallOpts, provider_ common.Address) (IProviderStorageProvider, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getProvider", provider_)

	if err != nil {
		return *new(IProviderStorageProvider), err
	}

	out0 := *abi.ConvertType(out[0], new(IProviderStorageProvider)).(*IProviderStorageProvider)

	return out0, err

}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address provider_) view returns((string,uint256,uint128,uint128,uint256,bool))
func (_Marketplace *MarketplaceSession) GetProvider(provider_ common.Address) (IProviderStorageProvider, error) {
	return _Marketplace.Contract.GetProvider(&_Marketplace.CallOpts, provider_)
}

// GetProvider is a free data retrieval call binding the contract method 0x55f21eb7.
//
// Solidity: function getProvider(address provider_) view returns((string,uint256,uint128,uint128,uint256,bool))
func (_Marketplace *MarketplaceCallerSession) GetProvider(provider_ common.Address) (IProviderStorageProvider, error) {
	return _Marketplace.Contract.GetProvider(&_Marketplace.CallOpts, provider_)
}

// GetProviderActiveBids is a free data retrieval call binding the contract method 0xaf5b77ca.
//
// Solidity: function getProviderActiveBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceCaller) GetProviderActiveBids(opts *bind.CallOpts, provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getProviderActiveBids", provider_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProviderActiveBids is a free data retrieval call binding the contract method 0xaf5b77ca.
//
// Solidity: function getProviderActiveBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceSession) GetProviderActiveBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _Marketplace.Contract.GetProviderActiveBids(&_Marketplace.CallOpts, provider_, offset_, limit_)
}

// GetProviderActiveBids is a free data retrieval call binding the contract method 0xaf5b77ca.
//
// Solidity: function getProviderActiveBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceCallerSession) GetProviderActiveBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _Marketplace.Contract.GetProviderActiveBids(&_Marketplace.CallOpts, provider_, offset_, limit_)
}

// GetProviderBids is a free data retrieval call binding the contract method 0x59d435c4.
//
// Solidity: function getProviderBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceCaller) GetProviderBids(opts *bind.CallOpts, provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getProviderBids", provider_, offset_, limit_)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetProviderBids is a free data retrieval call binding the contract method 0x59d435c4.
//
// Solidity: function getProviderBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceSession) GetProviderBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _Marketplace.Contract.GetProviderBids(&_Marketplace.CallOpts, provider_, offset_, limit_)
}

// GetProviderBids is a free data retrieval call binding the contract method 0x59d435c4.
//
// Solidity: function getProviderBids(address provider_, uint256 offset_, uint256 limit_) view returns(bytes32[])
func (_Marketplace *MarketplaceCallerSession) GetProviderBids(provider_ common.Address, offset_ *big.Int, limit_ *big.Int) ([][32]byte, error) {
	return _Marketplace.Contract.GetProviderBids(&_Marketplace.CallOpts, provider_, offset_, limit_)
}

// GetProviderMinimumStake is a free data retrieval call binding the contract method 0x53c029f6.
//
// Solidity: function getProviderMinimumStake() view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetProviderMinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getProviderMinimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProviderMinimumStake is a free data retrieval call binding the contract method 0x53c029f6.
//
// Solidity: function getProviderMinimumStake() view returns(uint256)
func (_Marketplace *MarketplaceSession) GetProviderMinimumStake() (*big.Int, error) {
	return _Marketplace.Contract.GetProviderMinimumStake(&_Marketplace.CallOpts)
}

// GetProviderMinimumStake is a free data retrieval call binding the contract method 0x53c029f6.
//
// Solidity: function getProviderMinimumStake() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetProviderMinimumStake() (*big.Int, error) {
	return _Marketplace.Contract.GetProviderMinimumStake(&_Marketplace.CallOpts)
}

// GetProviderModelId is a free data retrieval call binding the contract method 0x1cc9de8c.
//
// Solidity: function getProviderModelId(address provider_, bytes32 modelId_) pure returns(bytes32)
func (_Marketplace *MarketplaceCaller) GetProviderModelId(opts *bind.CallOpts, provider_ common.Address, modelId_ [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getProviderModelId", provider_, modelId_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetProviderModelId is a free data retrieval call binding the contract method 0x1cc9de8c.
//
// Solidity: function getProviderModelId(address provider_, bytes32 modelId_) pure returns(bytes32)
func (_Marketplace *MarketplaceSession) GetProviderModelId(provider_ common.Address, modelId_ [32]byte) ([32]byte, error) {
	return _Marketplace.Contract.GetProviderModelId(&_Marketplace.CallOpts, provider_, modelId_)
}

// GetProviderModelId is a free data retrieval call binding the contract method 0x1cc9de8c.
//
// Solidity: function getProviderModelId(address provider_, bytes32 modelId_) pure returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) GetProviderModelId(provider_ common.Address, modelId_ [32]byte) ([32]byte, error) {
	return _Marketplace.Contract.GetProviderModelId(&_Marketplace.CallOpts, provider_, modelId_)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_Marketplace *MarketplaceCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_Marketplace *MarketplaceSession) GetToken() (common.Address, error) {
	return _Marketplace.Contract.GetToken(&_Marketplace.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_Marketplace *MarketplaceCallerSession) GetToken() (common.Address, error) {
	return _Marketplace.Contract.GetToken(&_Marketplace.CallOpts)
}

// IsBidActive is a free data retrieval call binding the contract method 0x1345df58.
//
// Solidity: function isBidActive(bytes32 bidId_) view returns(bool)
func (_Marketplace *MarketplaceCaller) IsBidActive(opts *bind.CallOpts, bidId_ [32]byte) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "isBidActive", bidId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBidActive is a free data retrieval call binding the contract method 0x1345df58.
//
// Solidity: function isBidActive(bytes32 bidId_) view returns(bool)
func (_Marketplace *MarketplaceSession) IsBidActive(bidId_ [32]byte) (bool, error) {
	return _Marketplace.Contract.IsBidActive(&_Marketplace.CallOpts, bidId_)
}

// IsBidActive is a free data retrieval call binding the contract method 0x1345df58.
//
// Solidity: function isBidActive(bytes32 bidId_) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) IsBidActive(bidId_ [32]byte) (bool, error) {
	return _Marketplace.Contract.IsBidActive(&_Marketplace.CallOpts, bidId_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// MarketplaceInit is a paid mutator transaction binding the contract method 0xf6ec33fe.
//
// Solidity: function __Marketplace_init(address token_) returns()
func (_Marketplace *MarketplaceTransactor) MarketplaceInit(opts *bind.TransactOpts, token_ common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "__Marketplace_init", token_)
}

// MarketplaceInit is a paid mutator transaction binding the contract method 0xf6ec33fe.
//
// Solidity: function __Marketplace_init(address token_) returns()
func (_Marketplace *MarketplaceSession) MarketplaceInit(token_ common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceInit(&_Marketplace.TransactOpts, token_)
}

// MarketplaceInit is a paid mutator transaction binding the contract method 0xf6ec33fe.
//
// Solidity: function __Marketplace_init(address token_) returns()
func (_Marketplace *MarketplaceTransactorSession) MarketplaceInit(token_ common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceInit(&_Marketplace.TransactOpts, token_)
}

// DeleteModelBid is a paid mutator transaction binding the contract method 0x8913dcaa.
//
// Solidity: function deleteModelBid(bytes32 bidId_) returns()
func (_Marketplace *MarketplaceTransactor) DeleteModelBid(opts *bind.TransactOpts, bidId_ [32]byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "deleteModelBid", bidId_)
}

// DeleteModelBid is a paid mutator transaction binding the contract method 0x8913dcaa.
//
// Solidity: function deleteModelBid(bytes32 bidId_) returns()
func (_Marketplace *MarketplaceSession) DeleteModelBid(bidId_ [32]byte) (*types.Transaction, error) {
	return _Marketplace.Contract.DeleteModelBid(&_Marketplace.TransactOpts, bidId_)
}

// DeleteModelBid is a paid mutator transaction binding the contract method 0x8913dcaa.
//
// Solidity: function deleteModelBid(bytes32 bidId_) returns()
func (_Marketplace *MarketplaceTransactorSession) DeleteModelBid(bidId_ [32]byte) (*types.Transaction, error) {
	return _Marketplace.Contract.DeleteModelBid(&_Marketplace.TransactOpts, bidId_)
}

// PostModelBid is a paid mutator transaction binding the contract method 0x70ccc166.
//
// Solidity: function postModelBid(bytes32 modelId_, uint256 pricePerSecond_) returns(bytes32 bidId)
func (_Marketplace *MarketplaceTransactor) PostModelBid(opts *bind.TransactOpts, modelId_ [32]byte, pricePerSecond_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "postModelBid", modelId_, pricePerSecond_)
}

// PostModelBid is a paid mutator transaction binding the contract method 0x70ccc166.
//
// Solidity: function postModelBid(bytes32 modelId_, uint256 pricePerSecond_) returns(bytes32 bidId)
func (_Marketplace *MarketplaceSession) PostModelBid(modelId_ [32]byte, pricePerSecond_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.PostModelBid(&_Marketplace.TransactOpts, modelId_, pricePerSecond_)
}

// PostModelBid is a paid mutator transaction binding the contract method 0x70ccc166.
//
// Solidity: function postModelBid(bytes32 modelId_, uint256 pricePerSecond_) returns(bytes32 bidId)
func (_Marketplace *MarketplaceTransactorSession) PostModelBid(modelId_ [32]byte, pricePerSecond_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.PostModelBid(&_Marketplace.TransactOpts, modelId_, pricePerSecond_)
}

// SetMarketplaceBidFee is a paid mutator transaction binding the contract method 0x849c239c.
//
// Solidity: function setMarketplaceBidFee(uint256 bidFee_) returns()
func (_Marketplace *MarketplaceTransactor) SetMarketplaceBidFee(opts *bind.TransactOpts, bidFee_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setMarketplaceBidFee", bidFee_)
}

// SetMarketplaceBidFee is a paid mutator transaction binding the contract method 0x849c239c.
//
// Solidity: function setMarketplaceBidFee(uint256 bidFee_) returns()
func (_Marketplace *MarketplaceSession) SetMarketplaceBidFee(bidFee_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetMarketplaceBidFee(&_Marketplace.TransactOpts, bidFee_)
}

// SetMarketplaceBidFee is a paid mutator transaction binding the contract method 0x849c239c.
//
// Solidity: function setMarketplaceBidFee(uint256 bidFee_) returns()
func (_Marketplace *MarketplaceTransactorSession) SetMarketplaceBidFee(bidFee_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetMarketplaceBidFee(&_Marketplace.TransactOpts, bidFee_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient_, uint256 amount_) returns()
func (_Marketplace *MarketplaceTransactor) Withdraw(opts *bind.TransactOpts, recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "withdraw", recipient_, amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient_, uint256 amount_) returns()
func (_Marketplace *MarketplaceSession) Withdraw(recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Withdraw(&_Marketplace.TransactOpts, recipient_, amount_)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient_, uint256 amount_) returns()
func (_Marketplace *MarketplaceTransactorSession) Withdraw(recipient_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.Withdraw(&_Marketplace.TransactOpts, recipient_, amount_)
}

// MarketplaceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Marketplace contract.
type MarketplaceInitializedIterator struct {
	Event *MarketplaceInitialized // Event containing the contract specifics and raw log

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
func (it *MarketplaceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceInitialized)
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
		it.Event = new(MarketplaceInitialized)
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
func (it *MarketplaceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceInitialized represents a Initialized event raised by the Marketplace contract.
type MarketplaceInitialized struct {
	StorageSlot [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xdc73717d728bcfa015e8117438a65319aa06e979ca324afa6e1ea645c28ea15d.
//
// Solidity: event Initialized(bytes32 storageSlot)
func (_Marketplace *MarketplaceFilterer) FilterInitialized(opts *bind.FilterOpts) (*MarketplaceInitializedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MarketplaceInitializedIterator{contract: _Marketplace.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xdc73717d728bcfa015e8117438a65319aa06e979ca324afa6e1ea645c28ea15d.
//
// Solidity: event Initialized(bytes32 storageSlot)
func (_Marketplace *MarketplaceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MarketplaceInitialized) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceInitialized)
				if err := _Marketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseInitialized(log types.Log) (*MarketplaceInitialized, error) {
	event := new(MarketplaceInitialized)
	if err := _Marketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMaretplaceFeeUpdatedIterator is returned from FilterMaretplaceFeeUpdated and is used to iterate over the raw logs and unpacked data for MaretplaceFeeUpdated events raised by the Marketplace contract.
type MarketplaceMaretplaceFeeUpdatedIterator struct {
	Event *MarketplaceMaretplaceFeeUpdated // Event containing the contract specifics and raw log

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
func (it *MarketplaceMaretplaceFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMaretplaceFeeUpdated)
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
		it.Event = new(MarketplaceMaretplaceFeeUpdated)
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
func (it *MarketplaceMaretplaceFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMaretplaceFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMaretplaceFeeUpdated represents a MaretplaceFeeUpdated event raised by the Marketplace contract.
type MarketplaceMaretplaceFeeUpdated struct {
	BidFee *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaretplaceFeeUpdated is a free log retrieval operation binding the contract event 0x9c16b729dd48a231a60f92c4c1a12a5225825dd20699fc221314ea2d73a97cce.
//
// Solidity: event MaretplaceFeeUpdated(uint256 bidFee)
func (_Marketplace *MarketplaceFilterer) FilterMaretplaceFeeUpdated(opts *bind.FilterOpts) (*MarketplaceMaretplaceFeeUpdatedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "MaretplaceFeeUpdated")
	if err != nil {
		return nil, err
	}
	return &MarketplaceMaretplaceFeeUpdatedIterator{contract: _Marketplace.contract, event: "MaretplaceFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchMaretplaceFeeUpdated is a free log subscription operation binding the contract event 0x9c16b729dd48a231a60f92c4c1a12a5225825dd20699fc221314ea2d73a97cce.
//
// Solidity: event MaretplaceFeeUpdated(uint256 bidFee)
func (_Marketplace *MarketplaceFilterer) WatchMaretplaceFeeUpdated(opts *bind.WatchOpts, sink chan<- *MarketplaceMaretplaceFeeUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "MaretplaceFeeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMaretplaceFeeUpdated)
				if err := _Marketplace.contract.UnpackLog(event, "MaretplaceFeeUpdated", log); err != nil {
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

// ParseMaretplaceFeeUpdated is a log parse operation binding the contract event 0x9c16b729dd48a231a60f92c4c1a12a5225825dd20699fc221314ea2d73a97cce.
//
// Solidity: event MaretplaceFeeUpdated(uint256 bidFee)
func (_Marketplace *MarketplaceFilterer) ParseMaretplaceFeeUpdated(log types.Log) (*MarketplaceMaretplaceFeeUpdated, error) {
	event := new(MarketplaceMaretplaceFeeUpdated)
	if err := _Marketplace.contract.UnpackLog(event, "MaretplaceFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMarketplaceBidDeletedIterator is returned from FilterMarketplaceBidDeleted and is used to iterate over the raw logs and unpacked data for MarketplaceBidDeleted events raised by the Marketplace contract.
type MarketplaceMarketplaceBidDeletedIterator struct {
	Event *MarketplaceMarketplaceBidDeleted // Event containing the contract specifics and raw log

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
func (it *MarketplaceMarketplaceBidDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMarketplaceBidDeleted)
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
		it.Event = new(MarketplaceMarketplaceBidDeleted)
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
func (it *MarketplaceMarketplaceBidDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMarketplaceBidDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMarketplaceBidDeleted represents a MarketplaceBidDeleted event raised by the Marketplace contract.
type MarketplaceMarketplaceBidDeleted struct {
	Provider common.Address
	ModelId  [32]byte
	Nonce    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketplaceBidDeleted is a free log retrieval operation binding the contract event 0x409dfc0f98bf6e062e576fbdc63c9f82392d44deff61e3412f8bd256d2814883.
//
// Solidity: event MarketplaceBidDeleted(address indexed provider, bytes32 indexed modelId, uint256 nonce)
func (_Marketplace *MarketplaceFilterer) FilterMarketplaceBidDeleted(opts *bind.FilterOpts, provider []common.Address, modelId [][32]byte) (*MarketplaceMarketplaceBidDeletedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "MarketplaceBidDeleted", providerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMarketplaceBidDeletedIterator{contract: _Marketplace.contract, event: "MarketplaceBidDeleted", logs: logs, sub: sub}, nil
}

// WatchMarketplaceBidDeleted is a free log subscription operation binding the contract event 0x409dfc0f98bf6e062e576fbdc63c9f82392d44deff61e3412f8bd256d2814883.
//
// Solidity: event MarketplaceBidDeleted(address indexed provider, bytes32 indexed modelId, uint256 nonce)
func (_Marketplace *MarketplaceFilterer) WatchMarketplaceBidDeleted(opts *bind.WatchOpts, sink chan<- *MarketplaceMarketplaceBidDeleted, provider []common.Address, modelId [][32]byte) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "MarketplaceBidDeleted", providerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMarketplaceBidDeleted)
				if err := _Marketplace.contract.UnpackLog(event, "MarketplaceBidDeleted", log); err != nil {
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

// ParseMarketplaceBidDeleted is a log parse operation binding the contract event 0x409dfc0f98bf6e062e576fbdc63c9f82392d44deff61e3412f8bd256d2814883.
//
// Solidity: event MarketplaceBidDeleted(address indexed provider, bytes32 indexed modelId, uint256 nonce)
func (_Marketplace *MarketplaceFilterer) ParseMarketplaceBidDeleted(log types.Log) (*MarketplaceMarketplaceBidDeleted, error) {
	event := new(MarketplaceMarketplaceBidDeleted)
	if err := _Marketplace.contract.UnpackLog(event, "MarketplaceBidDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceMarketplaceBidPostedIterator is returned from FilterMarketplaceBidPosted and is used to iterate over the raw logs and unpacked data for MarketplaceBidPosted events raised by the Marketplace contract.
type MarketplaceMarketplaceBidPostedIterator struct {
	Event *MarketplaceMarketplaceBidPosted // Event containing the contract specifics and raw log

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
func (it *MarketplaceMarketplaceBidPostedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceMarketplaceBidPosted)
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
		it.Event = new(MarketplaceMarketplaceBidPosted)
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
func (it *MarketplaceMarketplaceBidPostedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceMarketplaceBidPostedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceMarketplaceBidPosted represents a MarketplaceBidPosted event raised by the Marketplace contract.
type MarketplaceMarketplaceBidPosted struct {
	Provider common.Address
	ModelId  [32]byte
	Nonce    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterMarketplaceBidPosted is a free log retrieval operation binding the contract event 0xfc422bb61cd73bc76f27eecf69cce3db0b05d39769ca3ed0fb314a3d04bff6f6.
//
// Solidity: event MarketplaceBidPosted(address indexed provider, bytes32 indexed modelId, uint256 nonce)
func (_Marketplace *MarketplaceFilterer) FilterMarketplaceBidPosted(opts *bind.FilterOpts, provider []common.Address, modelId [][32]byte) (*MarketplaceMarketplaceBidPostedIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "MarketplaceBidPosted", providerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceMarketplaceBidPostedIterator{contract: _Marketplace.contract, event: "MarketplaceBidPosted", logs: logs, sub: sub}, nil
}

// WatchMarketplaceBidPosted is a free log subscription operation binding the contract event 0xfc422bb61cd73bc76f27eecf69cce3db0b05d39769ca3ed0fb314a3d04bff6f6.
//
// Solidity: event MarketplaceBidPosted(address indexed provider, bytes32 indexed modelId, uint256 nonce)
func (_Marketplace *MarketplaceFilterer) WatchMarketplaceBidPosted(opts *bind.WatchOpts, sink chan<- *MarketplaceMarketplaceBidPosted, provider []common.Address, modelId [][32]byte) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}
	var modelIdRule []interface{}
	for _, modelIdItem := range modelId {
		modelIdRule = append(modelIdRule, modelIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "MarketplaceBidPosted", providerRule, modelIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceMarketplaceBidPosted)
				if err := _Marketplace.contract.UnpackLog(event, "MarketplaceBidPosted", log); err != nil {
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

// ParseMarketplaceBidPosted is a log parse operation binding the contract event 0xfc422bb61cd73bc76f27eecf69cce3db0b05d39769ca3ed0fb314a3d04bff6f6.
//
// Solidity: event MarketplaceBidPosted(address indexed provider, bytes32 indexed modelId, uint256 nonce)
func (_Marketplace *MarketplaceFilterer) ParseMarketplaceBidPosted(log types.Log) (*MarketplaceMarketplaceBidPosted, error) {
	event := new(MarketplaceMarketplaceBidPosted)
	if err := _Marketplace.contract.UnpackLog(event, "MarketplaceBidPosted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
