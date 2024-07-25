// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package modelregistry

import (
	"errors"
	"fmt"
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

// LibSDSD is an auto generated low-level Go binding around an user-defined struct.
type LibSDSD struct {
	Mean  int64
	SqSum int64
}

// Model is an auto generated low-level Go binding around an user-defined struct.
type Model struct {
	IpfsCID   [32]byte
	Fee       *big.Int
	Stake     *big.Int
	Owner     common.Address
	Name      string
	Tags      []string
	CreatedAt *big.Int
	IsDeleted bool
}

// ModelStats is an auto generated low-level Go binding around an user-defined struct.
type ModelStats struct {
	TpsScaled1000 LibSDSD
	TtftMs        LibSDSD
	TotalDuration LibSDSD
	Count         uint32
}

// ModelRegistryMetaData contains all meta data concerning the ModelRegistry contract.
var ModelRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"KeyExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeyNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModelHasActiveBids\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ModelNotFound\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_contractOwner\",\"type\":\"address\"}],\"name\":\"NotContractOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSenderOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakeTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroKey\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"modelId\",\"type\":\"bytes32\"}],\"name\":\"ModelDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newStake\",\"type\":\"uint256\"}],\"name\":\"ModelMinStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"modelId\",\"type\":\"bytes32\"}],\"name\":\"ModelRegisteredUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"modelDeregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"modelExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modelGetAll\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ipfsCID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"tags\",\"type\":\"string[]\"},{\"internalType\":\"uint128\",\"name\":\"createdAt\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isDeleted\",\"type\":\"bool\"}],\"internalType\":\"structModel[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"modelGetByIndex\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ipfsCID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"tags\",\"type\":\"string[]\"},{\"internalType\":\"uint128\",\"name\":\"createdAt\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isDeleted\",\"type\":\"bool\"}],\"internalType\":\"structModel\",\"name\":\"model\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modelGetCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modelGetIds\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"modelMap\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"ipfsCID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"tags\",\"type\":\"string[]\"},{\"internalType\":\"uint128\",\"name\":\"createdAt\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isDeleted\",\"type\":\"bool\"}],\"internalType\":\"structModel\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"modelMinStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"modelId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"ipfsCID\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"tags\",\"type\":\"string[]\"}],\"name\":\"modelRegister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"modelResetStats\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minStake\",\"type\":\"uint256\"}],\"name\":\"modelSetMinStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"modelStats\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"int64\",\"name\":\"mean\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"sqSum\",\"type\":\"int64\"}],\"internalType\":\"structLibSD.SD\",\"name\":\"tpsScaled1000\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"mean\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"sqSum\",\"type\":\"int64\"}],\"internalType\":\"structLibSD.SD\",\"name\":\"ttftMs\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int64\",\"name\":\"mean\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"sqSum\",\"type\":\"int64\"}],\"internalType\":\"structLibSD.SD\",\"name\":\"totalDuration\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"count\",\"type\":\"uint32\"}],\"internalType\":\"structModelStats\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"models\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// ModelExists is a free data retrieval call binding the contract method 0x022964d2.
//
// Solidity: function modelExists(bytes32 id) view returns(bool)
func (_ModelRegistry *ModelRegistryCaller) ModelExists(opts *bind.CallOpts, id [32]byte) (bool, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "modelExists", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ModelExists is a free data retrieval call binding the contract method 0x022964d2.
//
// Solidity: function modelExists(bytes32 id) view returns(bool)
func (_ModelRegistry *ModelRegistrySession) ModelExists(id [32]byte) (bool, error) {
	return _ModelRegistry.Contract.ModelExists(&_ModelRegistry.CallOpts, id)
}

// ModelExists is a free data retrieval call binding the contract method 0x022964d2.
//
// Solidity: function modelExists(bytes32 id) view returns(bool)
func (_ModelRegistry *ModelRegistryCallerSession) ModelExists(id [32]byte) (bool, error) {
	return _ModelRegistry.Contract.ModelExists(&_ModelRegistry.CallOpts, id)
}

// ModelGetAll is a free data retrieval call binding the contract method 0xb889c67b.
//
// Solidity: function modelGetAll() view returns(bytes32[], (bytes32,uint256,uint256,address,string,string[],uint128,bool)[])
func (_ModelRegistry *ModelRegistryCaller) ModelGetAll(opts *bind.CallOpts) ([][32]byte, []Model, error) {
	fmt.Println("modelGetAll")
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "modelGetAll")
fmt.Println("modelGetAll error: ", err)
fmt.Println("modelGetAll output: ", out)
	if err != nil {
		return *new([][32]byte), *new([]Model), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)
	out1 := *abi.ConvertType(out[1], new([]Model)).(*[]Model)

	return out0, out1, err

}

// ModelGetAll is a free data retrieval call binding the contract method 0xb889c67b.
//
// Solidity: function modelGetAll() view returns(bytes32[], (bytes32,uint256,uint256,address,string,string[],uint128,bool)[])
func (_ModelRegistry *ModelRegistrySession) ModelGetAll() ([][32]byte, []Model, error) {
	return _ModelRegistry.Contract.ModelGetAll(&_ModelRegistry.CallOpts)
}

// ModelGetAll is a free data retrieval call binding the contract method 0xb889c67b.
//
// Solidity: function modelGetAll() view returns(bytes32[], (bytes32,uint256,uint256,address,string,string[],uint128,bool)[])
func (_ModelRegistry *ModelRegistryCallerSession) ModelGetAll() ([][32]byte, []Model, error) {
	return _ModelRegistry.Contract.ModelGetAll(&_ModelRegistry.CallOpts)
}

// ModelGetByIndex is a free data retrieval call binding the contract method 0x43119765.
//
// Solidity: function modelGetByIndex(uint256 index) view returns(bytes32 modelId, (bytes32,uint256,uint256,address,string,string[],uint128,bool) model)
func (_ModelRegistry *ModelRegistryCaller) ModelGetByIndex(opts *bind.CallOpts, index *big.Int) (struct {
	ModelId [32]byte
	Model   Model
}, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "modelGetByIndex", index)

	outstruct := new(struct {
		ModelId [32]byte
		Model   Model
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ModelId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Model = *abi.ConvertType(out[1], new(Model)).(*Model)

	return *outstruct, err

}

// ModelGetByIndex is a free data retrieval call binding the contract method 0x43119765.
//
// Solidity: function modelGetByIndex(uint256 index) view returns(bytes32 modelId, (bytes32,uint256,uint256,address,string,string[],uint128,bool) model)
func (_ModelRegistry *ModelRegistrySession) ModelGetByIndex(index *big.Int) (struct {
	ModelId [32]byte
	Model   Model
}, error) {
	return _ModelRegistry.Contract.ModelGetByIndex(&_ModelRegistry.CallOpts, index)
}

// ModelGetByIndex is a free data retrieval call binding the contract method 0x43119765.
//
// Solidity: function modelGetByIndex(uint256 index) view returns(bytes32 modelId, (bytes32,uint256,uint256,address,string,string[],uint128,bool) model)
func (_ModelRegistry *ModelRegistryCallerSession) ModelGetByIndex(index *big.Int) (struct {
	ModelId [32]byte
	Model   Model
}, error) {
	return _ModelRegistry.Contract.ModelGetByIndex(&_ModelRegistry.CallOpts, index)
}

// ModelGetCount is a free data retrieval call binding the contract method 0x807b8975.
//
// Solidity: function modelGetCount() view returns(uint256 count)
func (_ModelRegistry *ModelRegistryCaller) ModelGetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "modelGetCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ModelGetCount is a free data retrieval call binding the contract method 0x807b8975.
//
// Solidity: function modelGetCount() view returns(uint256 count)
func (_ModelRegistry *ModelRegistrySession) ModelGetCount() (*big.Int, error) {
	return _ModelRegistry.Contract.ModelGetCount(&_ModelRegistry.CallOpts)
}

// ModelGetCount is a free data retrieval call binding the contract method 0x807b8975.
//
// Solidity: function modelGetCount() view returns(uint256 count)
func (_ModelRegistry *ModelRegistryCallerSession) ModelGetCount() (*big.Int, error) {
	return _ModelRegistry.Contract.ModelGetCount(&_ModelRegistry.CallOpts)
}

// ModelGetIds is a free data retrieval call binding the contract method 0x9f9877fc.
//
// Solidity: function modelGetIds() view returns(bytes32[])
func (_ModelRegistry *ModelRegistryCaller) ModelGetIds(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "modelGetIds")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// ModelGetIds is a free data retrieval call binding the contract method 0x9f9877fc.
//
// Solidity: function modelGetIds() view returns(bytes32[])
func (_ModelRegistry *ModelRegistrySession) ModelGetIds() ([][32]byte, error) {
	return _ModelRegistry.Contract.ModelGetIds(&_ModelRegistry.CallOpts)
}

// ModelGetIds is a free data retrieval call binding the contract method 0x9f9877fc.
//
// Solidity: function modelGetIds() view returns(bytes32[])
func (_ModelRegistry *ModelRegistryCallerSession) ModelGetIds() ([][32]byte, error) {
	return _ModelRegistry.Contract.ModelGetIds(&_ModelRegistry.CallOpts)
}

// ModelMap is a free data retrieval call binding the contract method 0x6e5cbd85.
//
// Solidity: function modelMap(bytes32 id) view returns((bytes32,uint256,uint256,address,string,string[],uint128,bool))
func (_ModelRegistry *ModelRegistryCaller) ModelMap(opts *bind.CallOpts, id [32]byte) (Model, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "modelMap", id)

	if err != nil {
		return *new(Model), err
	}

	out0 := *abi.ConvertType(out[0], new(Model)).(*Model)

	return out0, err

}

// ModelMap is a free data retrieval call binding the contract method 0x6e5cbd85.
//
// Solidity: function modelMap(bytes32 id) view returns((bytes32,uint256,uint256,address,string,string[],uint128,bool))
func (_ModelRegistry *ModelRegistrySession) ModelMap(id [32]byte) (Model, error) {
	return _ModelRegistry.Contract.ModelMap(&_ModelRegistry.CallOpts, id)
}

// ModelMap is a free data retrieval call binding the contract method 0x6e5cbd85.
//
// Solidity: function modelMap(bytes32 id) view returns((bytes32,uint256,uint256,address,string,string[],uint128,bool))
func (_ModelRegistry *ModelRegistryCallerSession) ModelMap(id [32]byte) (Model, error) {
	return _ModelRegistry.Contract.ModelMap(&_ModelRegistry.CallOpts, id)
}

// ModelMinStake is a free data retrieval call binding the contract method 0xe41a6823.
//
// Solidity: function modelMinStake() view returns(uint256)
func (_ModelRegistry *ModelRegistryCaller) ModelMinStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "modelMinStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ModelMinStake is a free data retrieval call binding the contract method 0xe41a6823.
//
// Solidity: function modelMinStake() view returns(uint256)
func (_ModelRegistry *ModelRegistrySession) ModelMinStake() (*big.Int, error) {
	return _ModelRegistry.Contract.ModelMinStake(&_ModelRegistry.CallOpts)
}

// ModelMinStake is a free data retrieval call binding the contract method 0xe41a6823.
//
// Solidity: function modelMinStake() view returns(uint256)
func (_ModelRegistry *ModelRegistryCallerSession) ModelMinStake() (*big.Int, error) {
	return _ModelRegistry.Contract.ModelMinStake(&_ModelRegistry.CallOpts)
}

// ModelStats is a free data retrieval call binding the contract method 0x99644b2c.
//
// Solidity: function modelStats(bytes32 id) view returns(((int64,int64),(int64,int64),(int64,int64),uint32))
func (_ModelRegistry *ModelRegistryCaller) ModelStats(opts *bind.CallOpts, id [32]byte) (ModelStats, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "modelStats", id)

	if err != nil {
		return *new(ModelStats), err
	}

	out0 := *abi.ConvertType(out[0], new(ModelStats)).(*ModelStats)

	return out0, err

}

// ModelStats is a free data retrieval call binding the contract method 0x99644b2c.
//
// Solidity: function modelStats(bytes32 id) view returns(((int64,int64),(int64,int64),(int64,int64),uint32))
func (_ModelRegistry *ModelRegistrySession) ModelStats(id [32]byte) (ModelStats, error) {
	return _ModelRegistry.Contract.ModelStats(&_ModelRegistry.CallOpts, id)
}

// ModelStats is a free data retrieval call binding the contract method 0x99644b2c.
//
// Solidity: function modelStats(bytes32 id) view returns(((int64,int64),(int64,int64),(int64,int64),uint32))
func (_ModelRegistry *ModelRegistryCallerSession) ModelStats(id [32]byte) (ModelStats, error) {
	return _ModelRegistry.Contract.ModelStats(&_ModelRegistry.CallOpts, id)
}

// Models is a free data retrieval call binding the contract method 0x6a030ca9.
//
// Solidity: function models(uint256 index) view returns(bytes32)
func (_ModelRegistry *ModelRegistryCaller) Models(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _ModelRegistry.contract.Call(opts, &out, "models", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Models is a free data retrieval call binding the contract method 0x6a030ca9.
//
// Solidity: function models(uint256 index) view returns(bytes32)
func (_ModelRegistry *ModelRegistrySession) Models(index *big.Int) ([32]byte, error) {
	return _ModelRegistry.Contract.Models(&_ModelRegistry.CallOpts, index)
}

// Models is a free data retrieval call binding the contract method 0x6a030ca9.
//
// Solidity: function models(uint256 index) view returns(bytes32)
func (_ModelRegistry *ModelRegistryCallerSession) Models(index *big.Int) ([32]byte, error) {
	return _ModelRegistry.Contract.Models(&_ModelRegistry.CallOpts, index)
}

// ModelDeregister is a paid mutator transaction binding the contract method 0xd5a245f1.
//
// Solidity: function modelDeregister(bytes32 id) returns()
func (_ModelRegistry *ModelRegistryTransactor) ModelDeregister(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _ModelRegistry.contract.Transact(opts, "modelDeregister", id)
}

// ModelDeregister is a paid mutator transaction binding the contract method 0xd5a245f1.
//
// Solidity: function modelDeregister(bytes32 id) returns()
func (_ModelRegistry *ModelRegistrySession) ModelDeregister(id [32]byte) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelDeregister(&_ModelRegistry.TransactOpts, id)
}

// ModelDeregister is a paid mutator transaction binding the contract method 0xd5a245f1.
//
// Solidity: function modelDeregister(bytes32 id) returns()
func (_ModelRegistry *ModelRegistryTransactorSession) ModelDeregister(id [32]byte) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelDeregister(&_ModelRegistry.TransactOpts, id)
}

// ModelRegister is a paid mutator transaction binding the contract method 0x82806227.
//
// Solidity: function modelRegister(bytes32 modelId, bytes32 ipfsCID, uint256 fee, uint256 addStake, address owner, string name, string[] tags) returns()
func (_ModelRegistry *ModelRegistryTransactor) ModelRegister(opts *bind.TransactOpts, modelId [32]byte, ipfsCID [32]byte, fee *big.Int, addStake *big.Int, owner common.Address, name string, tags []string) (*types.Transaction, error) {
	return _ModelRegistry.contract.Transact(opts, "modelRegister", modelId, ipfsCID, fee, addStake, owner, name, tags)
}

// ModelRegister is a paid mutator transaction binding the contract method 0x82806227.
//
// Solidity: function modelRegister(bytes32 modelId, bytes32 ipfsCID, uint256 fee, uint256 addStake, address owner, string name, string[] tags) returns()
func (_ModelRegistry *ModelRegistrySession) ModelRegister(modelId [32]byte, ipfsCID [32]byte, fee *big.Int, addStake *big.Int, owner common.Address, name string, tags []string) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelRegister(&_ModelRegistry.TransactOpts, modelId, ipfsCID, fee, addStake, owner, name, tags)
}

// ModelRegister is a paid mutator transaction binding the contract method 0x82806227.
//
// Solidity: function modelRegister(bytes32 modelId, bytes32 ipfsCID, uint256 fee, uint256 addStake, address owner, string name, string[] tags) returns()
func (_ModelRegistry *ModelRegistryTransactorSession) ModelRegister(modelId [32]byte, ipfsCID [32]byte, fee *big.Int, addStake *big.Int, owner common.Address, name string, tags []string) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelRegister(&_ModelRegistry.TransactOpts, modelId, ipfsCID, fee, addStake, owner, name, tags)
}

// ModelResetStats is a paid mutator transaction binding the contract method 0x9971035a.
//
// Solidity: function modelResetStats(bytes32 id) returns()
func (_ModelRegistry *ModelRegistryTransactor) ModelResetStats(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _ModelRegistry.contract.Transact(opts, "modelResetStats", id)
}

// ModelResetStats is a paid mutator transaction binding the contract method 0x9971035a.
//
// Solidity: function modelResetStats(bytes32 id) returns()
func (_ModelRegistry *ModelRegistrySession) ModelResetStats(id [32]byte) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelResetStats(&_ModelRegistry.TransactOpts, id)
}

// ModelResetStats is a paid mutator transaction binding the contract method 0x9971035a.
//
// Solidity: function modelResetStats(bytes32 id) returns()
func (_ModelRegistry *ModelRegistryTransactorSession) ModelResetStats(id [32]byte) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelResetStats(&_ModelRegistry.TransactOpts, id)
}

// ModelSetMinStake is a paid mutator transaction binding the contract method 0x78998329.
//
// Solidity: function modelSetMinStake(uint256 _minStake) returns()
func (_ModelRegistry *ModelRegistryTransactor) ModelSetMinStake(opts *bind.TransactOpts, _minStake *big.Int) (*types.Transaction, error) {
	return _ModelRegistry.contract.Transact(opts, "modelSetMinStake", _minStake)
}

// ModelSetMinStake is a paid mutator transaction binding the contract method 0x78998329.
//
// Solidity: function modelSetMinStake(uint256 _minStake) returns()
func (_ModelRegistry *ModelRegistrySession) ModelSetMinStake(_minStake *big.Int) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelSetMinStake(&_ModelRegistry.TransactOpts, _minStake)
}

// ModelSetMinStake is a paid mutator transaction binding the contract method 0x78998329.
//
// Solidity: function modelSetMinStake(uint256 _minStake) returns()
func (_ModelRegistry *ModelRegistryTransactorSession) ModelSetMinStake(_minStake *big.Int) (*types.Transaction, error) {
	return _ModelRegistry.Contract.ModelSetMinStake(&_ModelRegistry.TransactOpts, _minStake)
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

// ModelRegistryModelMinStakeUpdatedIterator is returned from FilterModelMinStakeUpdated and is used to iterate over the raw logs and unpacked data for ModelMinStakeUpdated events raised by the ModelRegistry contract.
type ModelRegistryModelMinStakeUpdatedIterator struct {
	Event *ModelRegistryModelMinStakeUpdated // Event containing the contract specifics and raw log

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
func (it *ModelRegistryModelMinStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ModelRegistryModelMinStakeUpdated)
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
		it.Event = new(ModelRegistryModelMinStakeUpdated)
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
func (it *ModelRegistryModelMinStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ModelRegistryModelMinStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ModelRegistryModelMinStakeUpdated represents a ModelMinStakeUpdated event raised by the ModelRegistry contract.
type ModelRegistryModelMinStakeUpdated struct {
	NewStake *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterModelMinStakeUpdated is a free log retrieval operation binding the contract event 0xa7facdcb561b9b7c3091d6bea4b06c48f97f719dade07ba7510a687109161f6e.
//
// Solidity: event ModelMinStakeUpdated(uint256 newStake)
func (_ModelRegistry *ModelRegistryFilterer) FilterModelMinStakeUpdated(opts *bind.FilterOpts) (*ModelRegistryModelMinStakeUpdatedIterator, error) {

	logs, sub, err := _ModelRegistry.contract.FilterLogs(opts, "ModelMinStakeUpdated")
	if err != nil {
		return nil, err
	}
	return &ModelRegistryModelMinStakeUpdatedIterator{contract: _ModelRegistry.contract, event: "ModelMinStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchModelMinStakeUpdated is a free log subscription operation binding the contract event 0xa7facdcb561b9b7c3091d6bea4b06c48f97f719dade07ba7510a687109161f6e.
//
// Solidity: event ModelMinStakeUpdated(uint256 newStake)
func (_ModelRegistry *ModelRegistryFilterer) WatchModelMinStakeUpdated(opts *bind.WatchOpts, sink chan<- *ModelRegistryModelMinStakeUpdated) (event.Subscription, error) {

	logs, sub, err := _ModelRegistry.contract.WatchLogs(opts, "ModelMinStakeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ModelRegistryModelMinStakeUpdated)
				if err := _ModelRegistry.contract.UnpackLog(event, "ModelMinStakeUpdated", log); err != nil {
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

// ParseModelMinStakeUpdated is a log parse operation binding the contract event 0xa7facdcb561b9b7c3091d6bea4b06c48f97f719dade07ba7510a687109161f6e.
//
// Solidity: event ModelMinStakeUpdated(uint256 newStake)
func (_ModelRegistry *ModelRegistryFilterer) ParseModelMinStakeUpdated(log types.Log) (*ModelRegistryModelMinStakeUpdated, error) {
	event := new(ModelRegistryModelMinStakeUpdated)
	if err := _ModelRegistry.contract.UnpackLog(event, "ModelMinStakeUpdated", log); err != nil {
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
