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

// PricingABI is the input ABI used to generate the binding from.
const PricingABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"alerter\",\"type\":\"address\"}],\"name\":\"removeAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"currentBlockNumber\",\"type\":\"uint256\"},{\"name\":\"buy\",\"type\":\"bool\"},{\"name\":\"qty\",\"type\":\"uint256\"}],\"name\":\"getPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"reserve\",\"type\":\"address\"}],\"name\":\"setReserveAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"disableTokenTrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"enableTokenTrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"buyAmount\",\"type\":\"int256\"},{\"name\":\"priceUpdateBlock\",\"type\":\"uint256\"},{\"name\":\"currentBlock\",\"type\":\"uint256\"}],\"name\":\"recoredImbalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\"},{\"name\":\"baseBuy\",\"type\":\"uint256[]\"},{\"name\":\"baseSell\",\"type\":\"uint256[]\"},{\"name\":\"buy\",\"type\":\"bytes14[]\"},{\"name\":\"sell\",\"type\":\"bytes14[]\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"setBasePrice\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"alertersGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAlerter\",\"type\":\"address\"}],\"name\":\"addAlerter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPriceUpdateBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numTokensInCurrentCompactData\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validPriceDurationInBlocks\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"buy\",\"type\":\"bytes14[]\"},{\"name\":\"sell\",\"type\":\"bytes14[]\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"setCompactData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setValidPriceDurationInBlocks\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"claimAdmin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"xBuy\",\"type\":\"int256[]\"},{\"name\":\"yBuy\",\"type\":\"int256[]\"},{\"name\":\"xSell\",\"type\":\"int256[]\"},{\"name\":\"ySell\",\"type\":\"int256[]\"}],\"name\":\"setQtyStepFunction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"buy\",\"type\":\"bool\"}],\"name\":\"getBasicPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOperator\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"removeOperator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"xBuy\",\"type\":\"int256[]\"},{\"name\":\"yBuy\",\"type\":\"int256[]\"},{\"name\":\"xSell\",\"type\":\"int256[]\"},{\"name\":\"ySell\",\"type\":\"int256[]\"}],\"name\":\"setImbalanceStepFunction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"minimalRecordResolution\",\"type\":\"uint256\"},{\"name\":\"maxPerBlockImbalance\",\"type\":\"uint256\"},{\"name\":\"maxTotalImbalance\",\"type\":\"uint256\"}],\"name\":\"setTokenControlInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getCompactData\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes1\"},{\"name\":\"\",\"type\":\"bytes1\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenControlInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"operatorsGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_admin\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"WithdrawToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sendTo\",\"type\":\"address\"}],\"name\":\"WithdrawEther\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"pendingAdmin\",\"type\":\"address\"}],\"name\":\"TransferAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAdmin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"previousAdmin\",\"type\":\"address\"}],\"name\":\"ClaimAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newAlerter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"AddAlerter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOperator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isAdd\",\"type\":\"bool\"}],\"name\":\"AddOperator\",\"type\":\"event\"}]"

// Pricing is an auto generated Go binding around an Ethereum contract.
type Pricing struct {
	PricingCaller     // Read-only binding to the contract
	PricingTransactor // Write-only binding to the contract
}

// PricingCaller is an auto generated read-only Go binding around an Ethereum contract.
type PricingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PricingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PricingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PricingSession struct {
	Contract     *Pricing          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PricingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PricingCallerSession struct {
	Contract *PricingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PricingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PricingTransactorSession struct {
	Contract     *PricingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PricingRaw is an auto generated low-level Go binding around an Ethereum contract.
type PricingRaw struct {
	Contract *Pricing // Generic contract binding to access the raw methods on
}

// PricingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PricingCallerRaw struct {
	Contract *PricingCaller // Generic read-only contract binding to access the raw methods on
}

// PricingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PricingTransactorRaw struct {
	Contract *PricingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPricing creates a new instance of Pricing, bound to a specific deployed contract.
func NewPricing(address common.Address, backend bind.ContractBackend) (*Pricing, error) {
	contract, err := bindPricing(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pricing{PricingCaller: PricingCaller{contract: contract}, PricingTransactor: PricingTransactor{contract: contract}}, nil
}

// NewPricingCaller creates a new read-only instance of Pricing, bound to a specific deployed contract.
func NewPricingCaller(address common.Address, caller bind.ContractCaller) (*PricingCaller, error) {
	contract, err := bindPricing(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &PricingCaller{contract: contract}, nil
}

// NewPricingTransactor creates a new write-only instance of Pricing, bound to a specific deployed contract.
func NewPricingTransactor(address common.Address, transactor bind.ContractTransactor) (*PricingTransactor, error) {
	contract, err := bindPricing(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &PricingTransactor{contract: contract}, nil
}

// bindPricing binds a generic wrapper to an already deployed contract.
func bindPricing(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PricingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pricing *PricingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pricing.Contract.PricingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pricing *PricingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pricing.Contract.PricingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pricing *PricingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pricing.Contract.PricingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pricing *PricingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Pricing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pricing *PricingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pricing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pricing *PricingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pricing.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Pricing *PricingCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pricing.contract.Call(opts, out, "admin")
	return *ret0, err
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Pricing *PricingSession) Admin() (common.Address, error) {
	return _Pricing.Contract.Admin(&_Pricing.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() constant returns(address)
func (_Pricing *PricingCallerSession) Admin() (common.Address, error) {
	return _Pricing.Contract.Admin(&_Pricing.CallOpts)
}

// AlertersGroup is a free data retrieval call binding the contract method 0x3a7a1081.
//
// Solidity: function alertersGroup( uint256) constant returns(address)
func (_Pricing *PricingCaller) AlertersGroup(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pricing.contract.Call(opts, out, "alertersGroup", arg0)
	return *ret0, err
}

// AlertersGroup is a free data retrieval call binding the contract method 0x3a7a1081.
//
// Solidity: function alertersGroup( uint256) constant returns(address)
func (_Pricing *PricingSession) AlertersGroup(arg0 *big.Int) (common.Address, error) {
	return _Pricing.Contract.AlertersGroup(&_Pricing.CallOpts, arg0)
}

// AlertersGroup is a free data retrieval call binding the contract method 0x3a7a1081.
//
// Solidity: function alertersGroup( uint256) constant returns(address)
func (_Pricing *PricingCallerSession) AlertersGroup(arg0 *big.Int) (common.Address, error) {
	return _Pricing.Contract.AlertersGroup(&_Pricing.CallOpts, arg0)
}

// GetBasicPrice is a free data retrieval call binding the contract method 0x9306fce5.
//
// Solidity: function getBasicPrice(token address, buy bool) constant returns(uint256)
func (_Pricing *PricingCaller) GetBasicPrice(opts *bind.CallOpts, token common.Address, buy bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pricing.contract.Call(opts, out, "getBasicPrice", token, buy)
	return *ret0, err
}

// GetBasicPrice is a free data retrieval call binding the contract method 0x9306fce5.
//
// Solidity: function getBasicPrice(token address, buy bool) constant returns(uint256)
func (_Pricing *PricingSession) GetBasicPrice(token common.Address, buy bool) (*big.Int, error) {
	return _Pricing.Contract.GetBasicPrice(&_Pricing.CallOpts, token, buy)
}

// GetBasicPrice is a free data retrieval call binding the contract method 0x9306fce5.
//
// Solidity: function getBasicPrice(token address, buy bool) constant returns(uint256)
func (_Pricing *PricingCallerSession) GetBasicPrice(token common.Address, buy bool) (*big.Int, error) {
	return _Pricing.Contract.GetBasicPrice(&_Pricing.CallOpts, token, buy)
}

// GetCompactData is a free data retrieval call binding the contract method 0xe4a2ac62.
//
// Solidity: function getCompactData(token address) constant returns(uint256, uint256, bytes1, bytes1)
func (_Pricing *PricingCaller) GetCompactData(opts *bind.CallOpts, token common.Address) (*big.Int, *big.Int, [1]byte, [1]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new([1]byte)
		ret3 = new([1]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Pricing.contract.Call(opts, out, "getCompactData", token)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetCompactData is a free data retrieval call binding the contract method 0xe4a2ac62.
//
// Solidity: function getCompactData(token address) constant returns(uint256, uint256, bytes1, bytes1)
func (_Pricing *PricingSession) GetCompactData(token common.Address) (*big.Int, *big.Int, [1]byte, [1]byte, error) {
	return _Pricing.Contract.GetCompactData(&_Pricing.CallOpts, token)
}

// GetCompactData is a free data retrieval call binding the contract method 0xe4a2ac62.
//
// Solidity: function getCompactData(token address) constant returns(uint256, uint256, bytes1, bytes1)
func (_Pricing *PricingCallerSession) GetCompactData(token common.Address) (*big.Int, *big.Int, [1]byte, [1]byte, error) {
	return _Pricing.Contract.GetCompactData(&_Pricing.CallOpts, token)
}

// GetPrice is a free data retrieval call binding the contract method 0x13e2fdd3.
//
// Solidity: function getPrice(token address, currentBlockNumber uint256, buy bool, qty uint256) constant returns(uint256)
func (_Pricing *PricingCaller) GetPrice(opts *bind.CallOpts, token common.Address, currentBlockNumber *big.Int, buy bool, qty *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pricing.contract.Call(opts, out, "getPrice", token, currentBlockNumber, buy, qty)
	return *ret0, err
}

// GetPrice is a free data retrieval call binding the contract method 0x13e2fdd3.
//
// Solidity: function getPrice(token address, currentBlockNumber uint256, buy bool, qty uint256) constant returns(uint256)
func (_Pricing *PricingSession) GetPrice(token common.Address, currentBlockNumber *big.Int, buy bool, qty *big.Int) (*big.Int, error) {
	return _Pricing.Contract.GetPrice(&_Pricing.CallOpts, token, currentBlockNumber, buy, qty)
}

// GetPrice is a free data retrieval call binding the contract method 0x13e2fdd3.
//
// Solidity: function getPrice(token address, currentBlockNumber uint256, buy bool, qty uint256) constant returns(uint256)
func (_Pricing *PricingCallerSession) GetPrice(token common.Address, currentBlockNumber *big.Int, buy bool, qty *big.Int) (*big.Int, error) {
	return _Pricing.Contract.GetPrice(&_Pricing.CallOpts, token, currentBlockNumber, buy, qty)
}

// GetPriceUpdateBlock is a free data retrieval call binding the contract method 0x4132cf6f.
//
// Solidity: function getPriceUpdateBlock(token address) constant returns(uint256)
func (_Pricing *PricingCaller) GetPriceUpdateBlock(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pricing.contract.Call(opts, out, "getPriceUpdateBlock", token)
	return *ret0, err
}

// GetPriceUpdateBlock is a free data retrieval call binding the contract method 0x4132cf6f.
//
// Solidity: function getPriceUpdateBlock(token address) constant returns(uint256)
func (_Pricing *PricingSession) GetPriceUpdateBlock(token common.Address) (*big.Int, error) {
	return _Pricing.Contract.GetPriceUpdateBlock(&_Pricing.CallOpts, token)
}

// GetPriceUpdateBlock is a free data retrieval call binding the contract method 0x4132cf6f.
//
// Solidity: function getPriceUpdateBlock(token address) constant returns(uint256)
func (_Pricing *PricingCallerSession) GetPriceUpdateBlock(token common.Address) (*big.Int, error) {
	return _Pricing.Contract.GetPriceUpdateBlock(&_Pricing.CallOpts, token)
}

// GetTokenControlInfo is a free data retrieval call binding the contract method 0xe7d4fd91.
//
// Solidity: function getTokenControlInfo(token address) constant returns(uint256, uint256, uint256)
func (_Pricing *PricingCaller) GetTokenControlInfo(opts *bind.CallOpts, token common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _Pricing.contract.Call(opts, out, "getTokenControlInfo", token)
	return *ret0, *ret1, *ret2, err
}

// GetTokenControlInfo is a free data retrieval call binding the contract method 0xe7d4fd91.
//
// Solidity: function getTokenControlInfo(token address) constant returns(uint256, uint256, uint256)
func (_Pricing *PricingSession) GetTokenControlInfo(token common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Pricing.Contract.GetTokenControlInfo(&_Pricing.CallOpts, token)
}

// GetTokenControlInfo is a free data retrieval call binding the contract method 0xe7d4fd91.
//
// Solidity: function getTokenControlInfo(token address) constant returns(uint256, uint256, uint256)
func (_Pricing *PricingCallerSession) GetTokenControlInfo(token common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Pricing.Contract.GetTokenControlInfo(&_Pricing.CallOpts, token)
}

// NumTokensInCurrentCompactData is a free data retrieval call binding the contract method 0x5085c9f1.
//
// Solidity: function numTokensInCurrentCompactData() constant returns(uint256)
func (_Pricing *PricingCaller) NumTokensInCurrentCompactData(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pricing.contract.Call(opts, out, "numTokensInCurrentCompactData")
	return *ret0, err
}

// NumTokensInCurrentCompactData is a free data retrieval call binding the contract method 0x5085c9f1.
//
// Solidity: function numTokensInCurrentCompactData() constant returns(uint256)
func (_Pricing *PricingSession) NumTokensInCurrentCompactData() (*big.Int, error) {
	return _Pricing.Contract.NumTokensInCurrentCompactData(&_Pricing.CallOpts)
}

// NumTokensInCurrentCompactData is a free data retrieval call binding the contract method 0x5085c9f1.
//
// Solidity: function numTokensInCurrentCompactData() constant returns(uint256)
func (_Pricing *PricingCallerSession) NumTokensInCurrentCompactData() (*big.Int, error) {
	return _Pricing.Contract.NumTokensInCurrentCompactData(&_Pricing.CallOpts)
}

// OperatorsGroup is a free data retrieval call binding the contract method 0xfc5bf0f2.
//
// Solidity: function operatorsGroup( uint256) constant returns(address)
func (_Pricing *PricingCaller) OperatorsGroup(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pricing.contract.Call(opts, out, "operatorsGroup", arg0)
	return *ret0, err
}

// OperatorsGroup is a free data retrieval call binding the contract method 0xfc5bf0f2.
//
// Solidity: function operatorsGroup( uint256) constant returns(address)
func (_Pricing *PricingSession) OperatorsGroup(arg0 *big.Int) (common.Address, error) {
	return _Pricing.Contract.OperatorsGroup(&_Pricing.CallOpts, arg0)
}

// OperatorsGroup is a free data retrieval call binding the contract method 0xfc5bf0f2.
//
// Solidity: function operatorsGroup( uint256) constant returns(address)
func (_Pricing *PricingCallerSession) OperatorsGroup(arg0 *big.Int) (common.Address, error) {
	return _Pricing.Contract.OperatorsGroup(&_Pricing.CallOpts, arg0)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_Pricing *PricingCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pricing.contract.Call(opts, out, "pendingAdmin")
	return *ret0, err
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_Pricing *PricingSession) PendingAdmin() (common.Address, error) {
	return _Pricing.Contract.PendingAdmin(&_Pricing.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() constant returns(address)
func (_Pricing *PricingCallerSession) PendingAdmin() (common.Address, error) {
	return _Pricing.Contract.PendingAdmin(&_Pricing.CallOpts)
}

// ReserveContract is a free data retrieval call binding the contract method 0xa7f43acd.
//
// Solidity: function reserveContract() constant returns(address)
func (_Pricing *PricingCaller) ReserveContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Pricing.contract.Call(opts, out, "reserveContract")
	return *ret0, err
}

// ReserveContract is a free data retrieval call binding the contract method 0xa7f43acd.
//
// Solidity: function reserveContract() constant returns(address)
func (_Pricing *PricingSession) ReserveContract() (common.Address, error) {
	return _Pricing.Contract.ReserveContract(&_Pricing.CallOpts)
}

// ReserveContract is a free data retrieval call binding the contract method 0xa7f43acd.
//
// Solidity: function reserveContract() constant returns(address)
func (_Pricing *PricingCallerSession) ReserveContract() (common.Address, error) {
	return _Pricing.Contract.ReserveContract(&_Pricing.CallOpts)
}

// ValidPriceDurationInBlocks is a free data retrieval call binding the contract method 0x5b125641.
//
// Solidity: function validPriceDurationInBlocks() constant returns(uint256)
func (_Pricing *PricingCaller) ValidPriceDurationInBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Pricing.contract.Call(opts, out, "validPriceDurationInBlocks")
	return *ret0, err
}

// ValidPriceDurationInBlocks is a free data retrieval call binding the contract method 0x5b125641.
//
// Solidity: function validPriceDurationInBlocks() constant returns(uint256)
func (_Pricing *PricingSession) ValidPriceDurationInBlocks() (*big.Int, error) {
	return _Pricing.Contract.ValidPriceDurationInBlocks(&_Pricing.CallOpts)
}

// ValidPriceDurationInBlocks is a free data retrieval call binding the contract method 0x5b125641.
//
// Solidity: function validPriceDurationInBlocks() constant returns(uint256)
func (_Pricing *PricingCallerSession) ValidPriceDurationInBlocks() (*big.Int, error) {
	return _Pricing.Contract.ValidPriceDurationInBlocks(&_Pricing.CallOpts)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(newAlerter address) returns()
func (_Pricing *PricingTransactor) AddAlerter(opts *bind.TransactOpts, newAlerter common.Address) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "addAlerter", newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(newAlerter address) returns()
func (_Pricing *PricingSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.AddAlerter(&_Pricing.TransactOpts, newAlerter)
}

// AddAlerter is a paid mutator transaction binding the contract method 0x408ee7fe.
//
// Solidity: function addAlerter(newAlerter address) returns()
func (_Pricing *PricingTransactorSession) AddAlerter(newAlerter common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.AddAlerter(&_Pricing.TransactOpts, newAlerter)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(newOperator address) returns()
func (_Pricing *PricingTransactor) AddOperator(opts *bind.TransactOpts, newOperator common.Address) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "addOperator", newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(newOperator address) returns()
func (_Pricing *PricingSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.AddOperator(&_Pricing.TransactOpts, newOperator)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(newOperator address) returns()
func (_Pricing *PricingTransactorSession) AddOperator(newOperator common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.AddOperator(&_Pricing.TransactOpts, newOperator)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(token address) returns()
func (_Pricing *PricingTransactor) AddToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "addToken", token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(token address) returns()
func (_Pricing *PricingSession) AddToken(token common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.AddToken(&_Pricing.TransactOpts, token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(token address) returns()
func (_Pricing *PricingTransactorSession) AddToken(token common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.AddToken(&_Pricing.TransactOpts, token)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_Pricing *PricingTransactor) ClaimAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "claimAdmin")
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_Pricing *PricingSession) ClaimAdmin() (*types.Transaction, error) {
	return _Pricing.Contract.ClaimAdmin(&_Pricing.TransactOpts)
}

// ClaimAdmin is a paid mutator transaction binding the contract method 0x77f50f97.
//
// Solidity: function claimAdmin() returns()
func (_Pricing *PricingTransactorSession) ClaimAdmin() (*types.Transaction, error) {
	return _Pricing.Contract.ClaimAdmin(&_Pricing.TransactOpts)
}

// DisableTokenTrade is a paid mutator transaction binding the contract method 0x158859f7.
//
// Solidity: function disableTokenTrade(token address) returns()
func (_Pricing *PricingTransactor) DisableTokenTrade(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "disableTokenTrade", token)
}

// DisableTokenTrade is a paid mutator transaction binding the contract method 0x158859f7.
//
// Solidity: function disableTokenTrade(token address) returns()
func (_Pricing *PricingSession) DisableTokenTrade(token common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.DisableTokenTrade(&_Pricing.TransactOpts, token)
}

// DisableTokenTrade is a paid mutator transaction binding the contract method 0x158859f7.
//
// Solidity: function disableTokenTrade(token address) returns()
func (_Pricing *PricingTransactorSession) DisableTokenTrade(token common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.DisableTokenTrade(&_Pricing.TransactOpts, token)
}

// EnableTokenTrade is a paid mutator transaction binding the contract method 0x1d6a8bda.
//
// Solidity: function enableTokenTrade(token address) returns()
func (_Pricing *PricingTransactor) EnableTokenTrade(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "enableTokenTrade", token)
}

// EnableTokenTrade is a paid mutator transaction binding the contract method 0x1d6a8bda.
//
// Solidity: function enableTokenTrade(token address) returns()
func (_Pricing *PricingSession) EnableTokenTrade(token common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.EnableTokenTrade(&_Pricing.TransactOpts, token)
}

// EnableTokenTrade is a paid mutator transaction binding the contract method 0x1d6a8bda.
//
// Solidity: function enableTokenTrade(token address) returns()
func (_Pricing *PricingTransactorSession) EnableTokenTrade(token common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.EnableTokenTrade(&_Pricing.TransactOpts, token)
}

// RecoredImbalance is a paid mutator transaction binding the contract method 0x24413bb2.
//
// Solidity: function recoredImbalance(token address, buyAmount int256, priceUpdateBlock uint256, currentBlock uint256) returns()
func (_Pricing *PricingTransactor) RecoredImbalance(opts *bind.TransactOpts, token common.Address, buyAmount *big.Int, priceUpdateBlock *big.Int, currentBlock *big.Int) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "recoredImbalance", token, buyAmount, priceUpdateBlock, currentBlock)
}

// RecoredImbalance is a paid mutator transaction binding the contract method 0x24413bb2.
//
// Solidity: function recoredImbalance(token address, buyAmount int256, priceUpdateBlock uint256, currentBlock uint256) returns()
func (_Pricing *PricingSession) RecoredImbalance(token common.Address, buyAmount *big.Int, priceUpdateBlock *big.Int, currentBlock *big.Int) (*types.Transaction, error) {
	return _Pricing.Contract.RecoredImbalance(&_Pricing.TransactOpts, token, buyAmount, priceUpdateBlock, currentBlock)
}

// RecoredImbalance is a paid mutator transaction binding the contract method 0x24413bb2.
//
// Solidity: function recoredImbalance(token address, buyAmount int256, priceUpdateBlock uint256, currentBlock uint256) returns()
func (_Pricing *PricingTransactorSession) RecoredImbalance(token common.Address, buyAmount *big.Int, priceUpdateBlock *big.Int, currentBlock *big.Int) (*types.Transaction, error) {
	return _Pricing.Contract.RecoredImbalance(&_Pricing.TransactOpts, token, buyAmount, priceUpdateBlock, currentBlock)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(alerter address) returns()
func (_Pricing *PricingTransactor) RemoveAlerter(opts *bind.TransactOpts, alerter common.Address) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "removeAlerter", alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(alerter address) returns()
func (_Pricing *PricingSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.RemoveAlerter(&_Pricing.TransactOpts, alerter)
}

// RemoveAlerter is a paid mutator transaction binding the contract method 0x01a12fd3.
//
// Solidity: function removeAlerter(alerter address) returns()
func (_Pricing *PricingTransactorSession) RemoveAlerter(alerter common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.RemoveAlerter(&_Pricing.TransactOpts, alerter)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(operator address) returns()
func (_Pricing *PricingTransactor) RemoveOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "removeOperator", operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(operator address) returns()
func (_Pricing *PricingSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.RemoveOperator(&_Pricing.TransactOpts, operator)
}

// RemoveOperator is a paid mutator transaction binding the contract method 0xac8a584a.
//
// Solidity: function removeOperator(operator address) returns()
func (_Pricing *PricingTransactorSession) RemoveOperator(operator common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.RemoveOperator(&_Pricing.TransactOpts, operator)
}

// SetBasePrice is a paid mutator transaction binding the contract method 0x247bd540.
//
// Solidity: function setBasePrice(tokens address[], baseBuy uint256[], baseSell uint256[], buy bytes14[], sell bytes14[], blockNumber uint256, indices uint256[]) returns()
func (_Pricing *PricingTransactor) SetBasePrice(opts *bind.TransactOpts, tokens []common.Address, baseBuy []*big.Int, baseSell []*big.Int, buy [][14]byte, sell [][14]byte, blockNumber *big.Int, indices []*big.Int) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "setBasePrice", tokens, baseBuy, baseSell, buy, sell, blockNumber, indices)
}

// SetBasePrice is a paid mutator transaction binding the contract method 0x247bd540.
//
// Solidity: function setBasePrice(tokens address[], baseBuy uint256[], baseSell uint256[], buy bytes14[], sell bytes14[], blockNumber uint256, indices uint256[]) returns()
func (_Pricing *PricingSession) SetBasePrice(tokens []common.Address, baseBuy []*big.Int, baseSell []*big.Int, buy [][14]byte, sell [][14]byte, blockNumber *big.Int, indices []*big.Int) (*types.Transaction, error) {
	return _Pricing.Contract.SetBasePrice(&_Pricing.TransactOpts, tokens, baseBuy, baseSell, buy, sell, blockNumber, indices)
}

// SetBasePrice is a paid mutator transaction binding the contract method 0x247bd540.
//
// Solidity: function setBasePrice(tokens address[], baseBuy uint256[], baseSell uint256[], buy bytes14[], sell bytes14[], blockNumber uint256, indices uint256[]) returns()
func (_Pricing *PricingTransactorSession) SetBasePrice(tokens []common.Address, baseBuy []*big.Int, baseSell []*big.Int, buy [][14]byte, sell [][14]byte, blockNumber *big.Int, indices []*big.Int) (*types.Transaction, error) {
	return _Pricing.Contract.SetBasePrice(&_Pricing.TransactOpts, tokens, baseBuy, baseSell, buy, sell, blockNumber, indices)
}

// SetCompactData is a paid mutator transaction binding the contract method 0x64887334.
//
// Solidity: function setCompactData(buy bytes14[], sell bytes14[], blockNumber uint256, indices uint256[]) returns()
func (_Pricing *PricingTransactor) SetCompactData(opts *bind.TransactOpts, buy [][14]byte, sell [][14]byte, blockNumber *big.Int, indices []*big.Int) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "setCompactData", buy, sell, blockNumber, indices)
}

// SetCompactData is a paid mutator transaction binding the contract method 0x64887334.
//
// Solidity: function setCompactData(buy bytes14[], sell bytes14[], blockNumber uint256, indices uint256[]) returns()
func (_Pricing *PricingSession) SetCompactData(buy [][14]byte, sell [][14]byte, blockNumber *big.Int, indices []*big.Int) (*types.Transaction, error) {
	return _Pricing.Contract.SetCompactData(&_Pricing.TransactOpts, buy, sell, blockNumber, indices)
}

// SetCompactData is a paid mutator transaction binding the contract method 0x64887334.
//
// Solidity: function setCompactData(buy bytes14[], sell bytes14[], blockNumber uint256, indices uint256[]) returns()
func (_Pricing *PricingTransactorSession) SetCompactData(buy [][14]byte, sell [][14]byte, blockNumber *big.Int, indices []*big.Int) (*types.Transaction, error) {
	return _Pricing.Contract.SetCompactData(&_Pricing.TransactOpts, buy, sell, blockNumber, indices)
}

// SetImbalanceStepFunction is a paid mutator transaction binding the contract method 0xbc9cbcc8.
//
// Solidity: function setImbalanceStepFunction(token address, xBuy int256[], yBuy int256[], xSell int256[], ySell int256[]) returns()
func (_Pricing *PricingTransactor) SetImbalanceStepFunction(opts *bind.TransactOpts, token common.Address, xBuy []*big.Int, yBuy []*big.Int, xSell []*big.Int, ySell []*big.Int) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "setImbalanceStepFunction", token, xBuy, yBuy, xSell, ySell)
}

// SetImbalanceStepFunction is a paid mutator transaction binding the contract method 0xbc9cbcc8.
//
// Solidity: function setImbalanceStepFunction(token address, xBuy int256[], yBuy int256[], xSell int256[], ySell int256[]) returns()
func (_Pricing *PricingSession) SetImbalanceStepFunction(token common.Address, xBuy []*big.Int, yBuy []*big.Int, xSell []*big.Int, ySell []*big.Int) (*types.Transaction, error) {
	return _Pricing.Contract.SetImbalanceStepFunction(&_Pricing.TransactOpts, token, xBuy, yBuy, xSell, ySell)
}

// SetImbalanceStepFunction is a paid mutator transaction binding the contract method 0xbc9cbcc8.
//
// Solidity: function setImbalanceStepFunction(token address, xBuy int256[], yBuy int256[], xSell int256[], ySell int256[]) returns()
func (_Pricing *PricingTransactorSession) SetImbalanceStepFunction(token common.Address, xBuy []*big.Int, yBuy []*big.Int, xSell []*big.Int, ySell []*big.Int) (*types.Transaction, error) {
	return _Pricing.Contract.SetImbalanceStepFunction(&_Pricing.TransactOpts, token, xBuy, yBuy, xSell, ySell)
}

// SetQtyStepFunction is a paid mutator transaction binding the contract method 0x80d8b380.
//
// Solidity: function setQtyStepFunction(token address, xBuy int256[], yBuy int256[], xSell int256[], ySell int256[]) returns()
func (_Pricing *PricingTransactor) SetQtyStepFunction(opts *bind.TransactOpts, token common.Address, xBuy []*big.Int, yBuy []*big.Int, xSell []*big.Int, ySell []*big.Int) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "setQtyStepFunction", token, xBuy, yBuy, xSell, ySell)
}

// SetQtyStepFunction is a paid mutator transaction binding the contract method 0x80d8b380.
//
// Solidity: function setQtyStepFunction(token address, xBuy int256[], yBuy int256[], xSell int256[], ySell int256[]) returns()
func (_Pricing *PricingSession) SetQtyStepFunction(token common.Address, xBuy []*big.Int, yBuy []*big.Int, xSell []*big.Int, ySell []*big.Int) (*types.Transaction, error) {
	return _Pricing.Contract.SetQtyStepFunction(&_Pricing.TransactOpts, token, xBuy, yBuy, xSell, ySell)
}

// SetQtyStepFunction is a paid mutator transaction binding the contract method 0x80d8b380.
//
// Solidity: function setQtyStepFunction(token address, xBuy int256[], yBuy int256[], xSell int256[], ySell int256[]) returns()
func (_Pricing *PricingTransactorSession) SetQtyStepFunction(token common.Address, xBuy []*big.Int, yBuy []*big.Int, xSell []*big.Int, ySell []*big.Int) (*types.Transaction, error) {
	return _Pricing.Contract.SetQtyStepFunction(&_Pricing.TransactOpts, token, xBuy, yBuy, xSell, ySell)
}

// SetReserveAddress is a paid mutator transaction binding the contract method 0x14673d31.
//
// Solidity: function setReserveAddress(reserve address) returns()
func (_Pricing *PricingTransactor) SetReserveAddress(opts *bind.TransactOpts, reserve common.Address) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "setReserveAddress", reserve)
}

// SetReserveAddress is a paid mutator transaction binding the contract method 0x14673d31.
//
// Solidity: function setReserveAddress(reserve address) returns()
func (_Pricing *PricingSession) SetReserveAddress(reserve common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.SetReserveAddress(&_Pricing.TransactOpts, reserve)
}

// SetReserveAddress is a paid mutator transaction binding the contract method 0x14673d31.
//
// Solidity: function setReserveAddress(reserve address) returns()
func (_Pricing *PricingTransactorSession) SetReserveAddress(reserve common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.SetReserveAddress(&_Pricing.TransactOpts, reserve)
}

// SetTokenControlInfo is a paid mutator transaction binding the contract method 0xbfee3569.
//
// Solidity: function setTokenControlInfo(token address, minimalRecordResolution uint256, maxPerBlockImbalance uint256, maxTotalImbalance uint256) returns()
func (_Pricing *PricingTransactor) SetTokenControlInfo(opts *bind.TransactOpts, token common.Address, minimalRecordResolution *big.Int, maxPerBlockImbalance *big.Int, maxTotalImbalance *big.Int) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "setTokenControlInfo", token, minimalRecordResolution, maxPerBlockImbalance, maxTotalImbalance)
}

// SetTokenControlInfo is a paid mutator transaction binding the contract method 0xbfee3569.
//
// Solidity: function setTokenControlInfo(token address, minimalRecordResolution uint256, maxPerBlockImbalance uint256, maxTotalImbalance uint256) returns()
func (_Pricing *PricingSession) SetTokenControlInfo(token common.Address, minimalRecordResolution *big.Int, maxPerBlockImbalance *big.Int, maxTotalImbalance *big.Int) (*types.Transaction, error) {
	return _Pricing.Contract.SetTokenControlInfo(&_Pricing.TransactOpts, token, minimalRecordResolution, maxPerBlockImbalance, maxTotalImbalance)
}

// SetTokenControlInfo is a paid mutator transaction binding the contract method 0xbfee3569.
//
// Solidity: function setTokenControlInfo(token address, minimalRecordResolution uint256, maxPerBlockImbalance uint256, maxTotalImbalance uint256) returns()
func (_Pricing *PricingTransactorSession) SetTokenControlInfo(token common.Address, minimalRecordResolution *big.Int, maxPerBlockImbalance *big.Int, maxTotalImbalance *big.Int) (*types.Transaction, error) {
	return _Pricing.Contract.SetTokenControlInfo(&_Pricing.TransactOpts, token, minimalRecordResolution, maxPerBlockImbalance, maxTotalImbalance)
}

// SetValidPriceDurationInBlocks is a paid mutator transaction binding the contract method 0x74daf5cc.
//
// Solidity: function setValidPriceDurationInBlocks(duration uint256) returns()
func (_Pricing *PricingTransactor) SetValidPriceDurationInBlocks(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "setValidPriceDurationInBlocks", duration)
}

// SetValidPriceDurationInBlocks is a paid mutator transaction binding the contract method 0x74daf5cc.
//
// Solidity: function setValidPriceDurationInBlocks(duration uint256) returns()
func (_Pricing *PricingSession) SetValidPriceDurationInBlocks(duration *big.Int) (*types.Transaction, error) {
	return _Pricing.Contract.SetValidPriceDurationInBlocks(&_Pricing.TransactOpts, duration)
}

// SetValidPriceDurationInBlocks is a paid mutator transaction binding the contract method 0x74daf5cc.
//
// Solidity: function setValidPriceDurationInBlocks(duration uint256) returns()
func (_Pricing *PricingTransactorSession) SetValidPriceDurationInBlocks(duration *big.Int) (*types.Transaction, error) {
	return _Pricing.Contract.SetValidPriceDurationInBlocks(&_Pricing.TransactOpts, duration)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(newAdmin address) returns()
func (_Pricing *PricingTransactor) TransferAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "transferAdmin", newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(newAdmin address) returns()
func (_Pricing *PricingSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.TransferAdmin(&_Pricing.TransactOpts, newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0x75829def.
//
// Solidity: function transferAdmin(newAdmin address) returns()
func (_Pricing *PricingTransactorSession) TransferAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.TransferAdmin(&_Pricing.TransactOpts, newAdmin)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(amount uint256, sendTo address) returns()
func (_Pricing *PricingTransactor) WithdrawEther(opts *bind.TransactOpts, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "withdrawEther", amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(amount uint256, sendTo address) returns()
func (_Pricing *PricingSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.WithdrawEther(&_Pricing.TransactOpts, amount, sendTo)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0xce56c454.
//
// Solidity: function withdrawEther(amount uint256, sendTo address) returns()
func (_Pricing *PricingTransactorSession) WithdrawEther(amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.WithdrawEther(&_Pricing.TransactOpts, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(token address, amount uint256, sendTo address) returns()
func (_Pricing *PricingTransactor) WithdrawToken(opts *bind.TransactOpts, token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _Pricing.contract.Transact(opts, "withdrawToken", token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(token address, amount uint256, sendTo address) returns()
func (_Pricing *PricingSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.WithdrawToken(&_Pricing.TransactOpts, token, amount, sendTo)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x3ccdbb28.
//
// Solidity: function withdrawToken(token address, amount uint256, sendTo address) returns()
func (_Pricing *PricingTransactorSession) WithdrawToken(token common.Address, amount *big.Int, sendTo common.Address) (*types.Transaction, error) {
	return _Pricing.Contract.WithdrawToken(&_Pricing.TransactOpts, token, amount, sendTo)
}
