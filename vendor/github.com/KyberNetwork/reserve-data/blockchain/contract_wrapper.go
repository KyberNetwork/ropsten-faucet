// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ContractWrapperABI is the input ABI used to generate the binding from.
const ContractWrapperABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"reserve\",\"type\":\"address\"},{\"name\":\"sources\",\"type\":\"address[]\"},{\"name\":\"dests\",\"type\":\"address[]\"}],\"name\":\"getPrices\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"reserve\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getBalances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"type\":\"function\"}]"

// ContractWrapper is an auto generated Go binding around an Ethereum contract.
type ContractWrapper struct {
	ContractWrapperCaller     // Read-only binding to the contract
	ContractWrapperTransactor // Write-only binding to the contract
}

// ContractWrapperCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractWrapperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractWrapperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractWrapperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractWrapperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractWrapperSession struct {
	Contract     *ContractWrapper  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractWrapperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractWrapperCallerSession struct {
	Contract *ContractWrapperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ContractWrapperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractWrapperTransactorSession struct {
	Contract     *ContractWrapperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ContractWrapperRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractWrapperRaw struct {
	Contract *ContractWrapper // Generic contract binding to access the raw methods on
}

// ContractWrapperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractWrapperCallerRaw struct {
	Contract *ContractWrapperCaller // Generic read-only contract binding to access the raw methods on
}

// ContractWrapperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractWrapperTransactorRaw struct {
	Contract *ContractWrapperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractWrapper creates a new instance of ContractWrapper, bound to a specific deployed contract.
func NewContractWrapper(address common.Address, backend bind.ContractBackend) (*ContractWrapper, error) {
	contract, err := bindContractWrapper(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractWrapper{ContractWrapperCaller: ContractWrapperCaller{contract: contract}, ContractWrapperTransactor: ContractWrapperTransactor{contract: contract}}, nil
}

// NewContractWrapperCaller creates a new read-only instance of ContractWrapper, bound to a specific deployed contract.
func NewContractWrapperCaller(address common.Address, caller bind.ContractCaller) (*ContractWrapperCaller, error) {
	contract, err := bindContractWrapper(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ContractWrapperCaller{contract: contract}, nil
}

// NewContractWrapperTransactor creates a new write-only instance of ContractWrapper, bound to a specific deployed contract.
func NewContractWrapperTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractWrapperTransactor, error) {
	contract, err := bindContractWrapper(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ContractWrapperTransactor{contract: contract}, nil
}

// bindContractWrapper binds a generic wrapper to an already deployed contract.
func bindContractWrapper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractWrapperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, nil), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractWrapper *ContractWrapperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractWrapper.Contract.ContractWrapperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractWrapper *ContractWrapperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractWrapper.Contract.ContractWrapperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractWrapper *ContractWrapperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractWrapper.Contract.ContractWrapperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractWrapper *ContractWrapperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractWrapper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractWrapper *ContractWrapperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractWrapper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractWrapper *ContractWrapperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractWrapper.Contract.contract.Transact(opts, method, params...)
}

// GetBalances is a free data retrieval call binding the contract method 0x6a385ae9.
//
// Solidity: function getBalances(reserve address, tokens address[]) constant returns(uint256[])
func (_ContractWrapper *ContractWrapperCaller) GetBalances(opts *bind.CallOpts, reserve common.Address, tokens []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _ContractWrapper.contract.Call(opts, out, "getBalances", reserve, tokens)
	return *ret0, err
}

// GetBalances is a free data retrieval call binding the contract method 0x6a385ae9.
//
// Solidity: function getBalances(reserve address, tokens address[]) constant returns(uint256[])
func (_ContractWrapper *ContractWrapperSession) GetBalances(reserve common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _ContractWrapper.Contract.GetBalances(&_ContractWrapper.CallOpts, reserve, tokens)
}

// GetBalances is a free data retrieval call binding the contract method 0x6a385ae9.
//
// Solidity: function getBalances(reserve address, tokens address[]) constant returns(uint256[])
func (_ContractWrapper *ContractWrapperCallerSession) GetBalances(reserve common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _ContractWrapper.Contract.GetBalances(&_ContractWrapper.CallOpts, reserve, tokens)
}

// GetPrices is a free data retrieval call binding the contract method 0x554d184b.
//
// Solidity: function getPrices(reserve address, sources address[], dests address[]) constant returns(uint256[], uint256[], uint256[])
func (_ContractWrapper *ContractWrapperCaller) GetPrices(opts *bind.CallOpts, reserve common.Address, sources []common.Address, dests []common.Address) ([]*big.Int, []*big.Int, []*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]*big.Int)
		ret2 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ContractWrapper.contract.Call(opts, out, "getPrices", reserve, sources, dests)
	return *ret0, *ret1, *ret2, err
}

// GetPrices is a free data retrieval call binding the contract method 0x554d184b.
//
// Solidity: function getPrices(reserve address, sources address[], dests address[]) constant returns(uint256[], uint256[], uint256[])
func (_ContractWrapper *ContractWrapperSession) GetPrices(reserve common.Address, sources []common.Address, dests []common.Address) ([]*big.Int, []*big.Int, []*big.Int, error) {
	return _ContractWrapper.Contract.GetPrices(&_ContractWrapper.CallOpts, reserve, sources, dests)
}

// GetPrices is a free data retrieval call binding the contract method 0x554d184b.
//
// Solidity: function getPrices(reserve address, sources address[], dests address[]) constant returns(uint256[], uint256[], uint256[])
func (_ContractWrapper *ContractWrapperCallerSession) GetPrices(reserve common.Address, sources []common.Address, dests []common.Address) ([]*big.Int, []*big.Int, []*big.Int, error) {
	return _ContractWrapper.Contract.GetPrices(&_ContractWrapper.CallOpts, reserve, sources, dests)
}
