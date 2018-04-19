// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bbs

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// BbsABI is the input ABI used to generate the binding from.
const BbsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_value\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_value\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// BbsBin is the compiled bytecode used for deploying new contracts.
const BbsBin = `0x6060604052341561000f57600080fd5b6040516103b13803806103b1833981016040528080519091019050600081805161003d929160200190610044565b50506100df565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061008557805160ff19168380011785556100b2565b828001600101855582156100b2579182015b828111156100b2578251825591602001919060010190610097565b506100be9291506100c2565b5090565b6100dc91905b808211156100be57600081556001016100c8565b90565b6102c3806100ee6000396000f30060606040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416634ed3885e81146100505780636d4ce63c146100a3575b600080fd5b341561005b57600080fd5b6100a160046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061012d95505050505050565b005b34156100ae57600080fd5b6100b6610144565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156100f25780820151838201526020016100da565b50505050905090810190601f16801561011f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60008180516101409291602001906101ed565b5050565b61014c61026b565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156101e25780601f106101b7576101008083540402835291602001916101e2565b820191906000526020600020905b8154815290600101906020018083116101c557829003601f168201915b505050505090505b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061022e57805160ff191683800117855561025b565b8280016001018555821561025b579182015b8281111561025b578251825591602001919060010190610240565b5061026792915061027d565b5090565b60206040519081016040526000815290565b6101ea91905b8082111561026757600081556001016102835600a165627a7a723058208279c5c6026f57ccedd52e2941b6c8415e0ec4a15f65a4d86eb32847c51015970029`

// DeployBbs deploys a new Ethereum contract, binding an instance of Bbs to it.
func DeployBbs(auth *bind.TransactOpts, backend bind.ContractBackend, _value string) (common.Address, *types.Transaction, *Bbs, error) {
	parsed, err := abi.JSON(strings.NewReader(BbsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BbsBin), backend, _value)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bbs{BbsCaller: BbsCaller{contract: contract}, BbsTransactor: BbsTransactor{contract: contract}, BbsFilterer: BbsFilterer{contract: contract}}, nil
}

// Bbs is an auto generated Go binding around an Ethereum contract.
type Bbs struct {
	BbsCaller     // Read-only binding to the contract
	BbsTransactor // Write-only binding to the contract
	BbsFilterer   // Log filterer for contract events
}

// BbsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BbsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BbsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BbsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BbsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BbsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BbsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BbsSession struct {
	Contract     *Bbs              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BbsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BbsCallerSession struct {
	Contract *BbsCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BbsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BbsTransactorSession struct {
	Contract     *BbsTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BbsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BbsRaw struct {
	Contract *Bbs // Generic contract binding to access the raw methods on
}

// BbsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BbsCallerRaw struct {
	Contract *BbsCaller // Generic read-only contract binding to access the raw methods on
}

// BbsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BbsTransactorRaw struct {
	Contract *BbsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBbs creates a new instance of Bbs, bound to a specific deployed contract.
func NewBbs(address common.Address, backend bind.ContractBackend) (*Bbs, error) {
	contract, err := bindBbs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bbs{BbsCaller: BbsCaller{contract: contract}, BbsTransactor: BbsTransactor{contract: contract}, BbsFilterer: BbsFilterer{contract: contract}}, nil
}

// NewBbsCaller creates a new read-only instance of Bbs, bound to a specific deployed contract.
func NewBbsCaller(address common.Address, caller bind.ContractCaller) (*BbsCaller, error) {
	contract, err := bindBbs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BbsCaller{contract: contract}, nil
}

// NewBbsTransactor creates a new write-only instance of Bbs, bound to a specific deployed contract.
func NewBbsTransactor(address common.Address, transactor bind.ContractTransactor) (*BbsTransactor, error) {
	contract, err := bindBbs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BbsTransactor{contract: contract}, nil
}

// NewBbsFilterer creates a new log filterer instance of Bbs, bound to a specific deployed contract.
func NewBbsFilterer(address common.Address, filterer bind.ContractFilterer) (*BbsFilterer, error) {
	contract, err := bindBbs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BbsFilterer{contract: contract}, nil
}

// bindBbs binds a generic wrapper to an already deployed contract.
func bindBbs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BbsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bbs *BbsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bbs.Contract.BbsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bbs *BbsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bbs.Contract.BbsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bbs *BbsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bbs.Contract.BbsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bbs *BbsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bbs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bbs *BbsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bbs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bbs *BbsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bbs.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(string)
func (_Bbs *BbsCaller) Get(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Bbs.contract.Call(opts, out, "get")
	return *ret0, err
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(string)
func (_Bbs *BbsSession) Get() (string, error) {
	return _Bbs.Contract.Get(&_Bbs.CallOpts)
}

// Get is a free data retrieval call binding the contract method 0x6d4ce63c.
//
// Solidity: function get() constant returns(string)
func (_Bbs *BbsCallerSession) Get() (string, error) {
	return _Bbs.Contract.Get(&_Bbs.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(_value string) returns()
func (_Bbs *BbsTransactor) Set(opts *bind.TransactOpts, _value string) (*types.Transaction, error) {
	return _Bbs.contract.Transact(opts, "set", _value)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(_value string) returns()
func (_Bbs *BbsSession) Set(_value string) (*types.Transaction, error) {
	return _Bbs.Contract.Set(&_Bbs.TransactOpts, _value)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(_value string) returns()
func (_Bbs *BbsTransactorSession) Set(_value string) (*types.Transaction, error) {
	return _Bbs.Contract.Set(&_Bbs.TransactOpts, _value)
}
