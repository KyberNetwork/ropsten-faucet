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
const ContractWrapperABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ETH_TOKEN_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x\",\"type\":\"bytes14\"},{\"name\":\"byteInd\",\"type\":\"uint256\"}],\"name\":\"getInt8FromByte\",\"outputs\":[{\"name\":\"\",\"type\":\"int8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"reserve\",\"type\":\"address\"},{\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"getBalances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pricingContract\",\"type\":\"address\"},{\"name\":\"tokenList\",\"type\":\"address[]\"}],\"name\":\"getTokenIndicies\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"x\",\"type\":\"bytes14\"},{\"name\":\"byteInd\",\"type\":\"uint256\"}],\"name\":\"getByteFromBytes14\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes1\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pricingContract\",\"type\":\"address\"},{\"name\":\"tokenList\",\"type\":\"address[]\"}],\"name\":\"getTokenRates\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"int8[]\"},{\"name\":\"\",\"type\":\"int8[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

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
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
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

// ETH_TOKEN_ADDRESS is a free data retrieval call binding the contract method 0x1878d1f1.
//
// Solidity: function ETH_TOKEN_ADDRESS() constant returns(address)
func (_ContractWrapper *ContractWrapperCaller) ETH_TOKEN_ADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ContractWrapper.contract.Call(opts, out, "ETH_TOKEN_ADDRESS")
	return *ret0, err
}

// ETH_TOKEN_ADDRESS is a free data retrieval call binding the contract method 0x1878d1f1.
//
// Solidity: function ETH_TOKEN_ADDRESS() constant returns(address)
func (_ContractWrapper *ContractWrapperSession) ETH_TOKEN_ADDRESS() (common.Address, error) {
	return _ContractWrapper.Contract.ETH_TOKEN_ADDRESS(&_ContractWrapper.CallOpts)
}

// ETH_TOKEN_ADDRESS is a free data retrieval call binding the contract method 0x1878d1f1.
//
// Solidity: function ETH_TOKEN_ADDRESS() constant returns(address)
func (_ContractWrapper *ContractWrapperCallerSession) ETH_TOKEN_ADDRESS() (common.Address, error) {
	return _ContractWrapper.Contract.ETH_TOKEN_ADDRESS(&_ContractWrapper.CallOpts)
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

// GetByteFromBytes14 is a free data retrieval call binding the contract method 0xa609f034.
//
// Solidity: function getByteFromBytes14(x bytes14, byteInd uint256) constant returns(bytes1)
func (_ContractWrapper *ContractWrapperCaller) GetByteFromBytes14(opts *bind.CallOpts, x [14]byte, byteInd *big.Int) ([1]byte, error) {
	var (
		ret0 = new([1]byte)
	)
	out := ret0
	err := _ContractWrapper.contract.Call(opts, out, "getByteFromBytes14", x, byteInd)
	return *ret0, err
}

// GetByteFromBytes14 is a free data retrieval call binding the contract method 0xa609f034.
//
// Solidity: function getByteFromBytes14(x bytes14, byteInd uint256) constant returns(bytes1)
func (_ContractWrapper *ContractWrapperSession) GetByteFromBytes14(x [14]byte, byteInd *big.Int) ([1]byte, error) {
	return _ContractWrapper.Contract.GetByteFromBytes14(&_ContractWrapper.CallOpts, x, byteInd)
}

// GetByteFromBytes14 is a free data retrieval call binding the contract method 0xa609f034.
//
// Solidity: function getByteFromBytes14(x bytes14, byteInd uint256) constant returns(bytes1)
func (_ContractWrapper *ContractWrapperCallerSession) GetByteFromBytes14(x [14]byte, byteInd *big.Int) ([1]byte, error) {
	return _ContractWrapper.Contract.GetByteFromBytes14(&_ContractWrapper.CallOpts, x, byteInd)
}

// GetInt8FromByte is a free data retrieval call binding the contract method 0x67c33c80.
//
// Solidity: function getInt8FromByte(x bytes14, byteInd uint256) constant returns(int8)
func (_ContractWrapper *ContractWrapperCaller) GetInt8FromByte(opts *bind.CallOpts, x [14]byte, byteInd *big.Int) (int8, error) {
	var (
		ret0 = new(int8)
	)
	out := ret0
	err := _ContractWrapper.contract.Call(opts, out, "getInt8FromByte", x, byteInd)
	return *ret0, err
}

// GetInt8FromByte is a free data retrieval call binding the contract method 0x67c33c80.
//
// Solidity: function getInt8FromByte(x bytes14, byteInd uint256) constant returns(int8)
func (_ContractWrapper *ContractWrapperSession) GetInt8FromByte(x [14]byte, byteInd *big.Int) (int8, error) {
	return _ContractWrapper.Contract.GetInt8FromByte(&_ContractWrapper.CallOpts, x, byteInd)
}

// GetInt8FromByte is a free data retrieval call binding the contract method 0x67c33c80.
//
// Solidity: function getInt8FromByte(x bytes14, byteInd uint256) constant returns(int8)
func (_ContractWrapper *ContractWrapperCallerSession) GetInt8FromByte(x [14]byte, byteInd *big.Int) (int8, error) {
	return _ContractWrapper.Contract.GetInt8FromByte(&_ContractWrapper.CallOpts, x, byteInd)
}

// GetTokenIndicies is a free data retrieval call binding the contract method 0x7c80feff.
//
// Solidity: function getTokenIndicies(pricingContract address, tokenList address[]) constant returns(uint256[], uint256[])
func (_ContractWrapper *ContractWrapperCaller) GetTokenIndicies(opts *bind.CallOpts, pricingContract common.Address, tokenList []common.Address) ([]*big.Int, []*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ContractWrapper.contract.Call(opts, out, "getTokenIndicies", pricingContract, tokenList)
	return *ret0, *ret1, err
}

// GetTokenIndicies is a free data retrieval call binding the contract method 0x7c80feff.
//
// Solidity: function getTokenIndicies(pricingContract address, tokenList address[]) constant returns(uint256[], uint256[])
func (_ContractWrapper *ContractWrapperSession) GetTokenIndicies(pricingContract common.Address, tokenList []common.Address) ([]*big.Int, []*big.Int, error) {
	return _ContractWrapper.Contract.GetTokenIndicies(&_ContractWrapper.CallOpts, pricingContract, tokenList)
}

// GetTokenIndicies is a free data retrieval call binding the contract method 0x7c80feff.
//
// Solidity: function getTokenIndicies(pricingContract address, tokenList address[]) constant returns(uint256[], uint256[])
func (_ContractWrapper *ContractWrapperCallerSession) GetTokenIndicies(pricingContract common.Address, tokenList []common.Address) ([]*big.Int, []*big.Int, error) {
	return _ContractWrapper.Contract.GetTokenIndicies(&_ContractWrapper.CallOpts, pricingContract, tokenList)
}

// GetTokenRates is a free data retrieval call binding the contract method 0xf37f8345.
//
// Solidity: function getTokenRates(pricingContract address, tokenList address[]) constant returns(uint256[], uint256[], int8[], int8[], uint256[])
func (_ContractWrapper *ContractWrapperCaller) GetTokenRates(opts *bind.CallOpts, pricingContract common.Address, tokenList []common.Address) ([]*big.Int, []*big.Int, []int8, []int8, []*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]*big.Int)
		ret2 = new([]int8)
		ret3 = new([]int8)
		ret4 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _ContractWrapper.contract.Call(opts, out, "getTokenRates", pricingContract, tokenList)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetTokenRates is a free data retrieval call binding the contract method 0xf37f8345.
//
// Solidity: function getTokenRates(pricingContract address, tokenList address[]) constant returns(uint256[], uint256[], int8[], int8[], uint256[])
func (_ContractWrapper *ContractWrapperSession) GetTokenRates(pricingContract common.Address, tokenList []common.Address) ([]*big.Int, []*big.Int, []int8, []int8, []*big.Int, error) {
	return _ContractWrapper.Contract.GetTokenRates(&_ContractWrapper.CallOpts, pricingContract, tokenList)
}

// GetTokenRates is a free data retrieval call binding the contract method 0xf37f8345.
//
// Solidity: function getTokenRates(pricingContract address, tokenList address[]) constant returns(uint256[], uint256[], int8[], int8[], uint256[])
func (_ContractWrapper *ContractWrapperCallerSession) GetTokenRates(pricingContract common.Address, tokenList []common.Address) ([]*big.Int, []*big.Int, []int8, []int8, []*big.Int, error) {
	return _ContractWrapper.Contract.GetTokenRates(&_ContractWrapper.CallOpts, pricingContract, tokenList)
}
