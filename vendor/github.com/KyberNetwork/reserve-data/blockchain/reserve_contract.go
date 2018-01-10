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

// ReserveContractABI is the input ABI used to generate the binding from.
const ReserveContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"sourceToken\",\"type\":\"address\"},{\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"name\":\"destToken\",\"type\":\"address\"},{\"name\":\"destAddress\",\"type\":\"address\"},{\"name\":\"validate\",\"type\":\"bool\"}],\"name\":\"trade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ETH_TOKEN_ADDRESS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"depositToken\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"source\",\"type\":\"address\"},{\"name\":\"dest\",\"type\":\"address\"}],\"name\":\"getPairInfo\",\"outputs\":[{\"name\":\"rate\",\"type\":\"uint256\"},{\"name\":\"expBlock\",\"type\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"depositEther\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kyberNetwork\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"enableTrade\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tradeEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"sources\",\"type\":\"address[]\"},{\"name\":\"dests\",\"type\":\"address[]\"},{\"name\":\"conversionRates\",\"type\":\"uint256[]\"},{\"name\":\"expiryBlocks\",\"type\":\"uint256[]\"},{\"name\":\"validate\",\"type\":\"bool\"}],\"name\":\"setRate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_kyberNetwork\",\"type\":\"address\"},{\"name\":\"_reserveOwner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"origin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"errorInfo\",\"type\":\"uint256\"}],\"name\":\"ErrorReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"origin\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sourceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"destAddress\",\"type\":\"address\"}],\"name\":\"DoTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"source\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dest\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"expiryBlock\",\"type\":\"uint256\"}],\"name\":\"SetRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"EnableTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"}]"

// ReserveContract is an auto generated Go binding around an Ethereum contract.
type ReserveContract struct {
	ReserveContractCaller     // Read-only binding to the contract
	ReserveContractTransactor // Write-only binding to the contract
}

// ReserveContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReserveContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReserveContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReserveContractSession struct {
	Contract     *ReserveContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReserveContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReserveContractCallerSession struct {
	Contract *ReserveContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReserveContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReserveContractTransactorSession struct {
	Contract     *ReserveContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReserveContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReserveContractRaw struct {
	Contract *ReserveContract // Generic contract binding to access the raw methods on
}

// ReserveContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReserveContractCallerRaw struct {
	Contract *ReserveContractCaller // Generic read-only contract binding to access the raw methods on
}

// ReserveContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReserveContractTransactorRaw struct {
	Contract *ReserveContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReserveContract creates a new instance of ReserveContract, bound to a specific deployed contract.
func NewReserveContract(address common.Address, backend bind.ContractBackend) (*ReserveContract, error) {
	contract, err := bindReserveContract(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReserveContract{ReserveContractCaller: ReserveContractCaller{contract: contract}, ReserveContractTransactor: ReserveContractTransactor{contract: contract}}, nil
}

// NewReserveContractCaller creates a new read-only instance of ReserveContract, bound to a specific deployed contract.
func NewReserveContractCaller(address common.Address, caller bind.ContractCaller) (*ReserveContractCaller, error) {
	contract, err := bindReserveContract(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveContractCaller{contract: contract}, nil
}

// NewReserveContractTransactor creates a new write-only instance of ReserveContract, bound to a specific deployed contract.
func NewReserveContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ReserveContractTransactor, error) {
	contract, err := bindReserveContract(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ReserveContractTransactor{contract: contract}, nil
}

// bindReserveContract binds a generic wrapper to an already deployed contract.
func bindReserveContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReserveContract *ReserveContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReserveContract.Contract.ReserveContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReserveContract *ReserveContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveContract.Contract.ReserveContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReserveContract *ReserveContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReserveContract.Contract.ReserveContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReserveContract *ReserveContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReserveContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReserveContract *ReserveContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReserveContract *ReserveContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReserveContract.Contract.contract.Transact(opts, method, params...)
}

// ETH_TOKEN_ADDRESS is a free data retrieval call binding the contract method 0x1878d1f1.
//
// Solidity: function ETH_TOKEN_ADDRESS() constant returns(address)
func (_ReserveContract *ReserveContractCaller) ETH_TOKEN_ADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReserveContract.contract.Call(opts, out, "ETH_TOKEN_ADDRESS")
	return *ret0, err
}

// ETH_TOKEN_ADDRESS is a free data retrieval call binding the contract method 0x1878d1f1.
//
// Solidity: function ETH_TOKEN_ADDRESS() constant returns(address)
func (_ReserveContract *ReserveContractSession) ETH_TOKEN_ADDRESS() (common.Address, error) {
	return _ReserveContract.Contract.ETH_TOKEN_ADDRESS(&_ReserveContract.CallOpts)
}

// ETH_TOKEN_ADDRESS is a free data retrieval call binding the contract method 0x1878d1f1.
//
// Solidity: function ETH_TOKEN_ADDRESS() constant returns(address)
func (_ReserveContract *ReserveContractCallerSession) ETH_TOKEN_ADDRESS() (common.Address, error) {
	return _ReserveContract.Contract.ETH_TOKEN_ADDRESS(&_ReserveContract.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(token address) constant returns(uint256)
func (_ReserveContract *ReserveContractCaller) GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReserveContract.contract.Call(opts, out, "getBalance", token)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(token address) constant returns(uint256)
func (_ReserveContract *ReserveContractSession) GetBalance(token common.Address) (*big.Int, error) {
	return _ReserveContract.Contract.GetBalance(&_ReserveContract.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(token address) constant returns(uint256)
func (_ReserveContract *ReserveContractCallerSession) GetBalance(token common.Address) (*big.Int, error) {
	return _ReserveContract.Contract.GetBalance(&_ReserveContract.CallOpts, token)
}

// GetPairInfo is a free data retrieval call binding the contract method 0x400f7a1e.
//
// Solidity: function getPairInfo(source address, dest address) constant returns(rate uint256, expBlock uint256, balance uint256)
func (_ReserveContract *ReserveContractCaller) GetPairInfo(opts *bind.CallOpts, source common.Address, dest common.Address) (struct {
	Rate     *big.Int
	ExpBlock *big.Int
	Balance  *big.Int
}, error) {
	ret := new(struct {
		Rate     *big.Int
		ExpBlock *big.Int
		Balance  *big.Int
	})
	out := ret
	err := _ReserveContract.contract.Call(opts, out, "getPairInfo", source, dest)
	return *ret, err
}

// GetPairInfo is a free data retrieval call binding the contract method 0x400f7a1e.
//
// Solidity: function getPairInfo(source address, dest address) constant returns(rate uint256, expBlock uint256, balance uint256)
func (_ReserveContract *ReserveContractSession) GetPairInfo(source common.Address, dest common.Address) (struct {
	Rate     *big.Int
	ExpBlock *big.Int
	Balance  *big.Int
}, error) {
	return _ReserveContract.Contract.GetPairInfo(&_ReserveContract.CallOpts, source, dest)
}

// GetPairInfo is a free data retrieval call binding the contract method 0x400f7a1e.
//
// Solidity: function getPairInfo(source address, dest address) constant returns(rate uint256, expBlock uint256, balance uint256)
func (_ReserveContract *ReserveContractCallerSession) GetPairInfo(source common.Address, dest common.Address) (struct {
	Rate     *big.Int
	ExpBlock *big.Int
	Balance  *big.Int
}, error) {
	return _ReserveContract.Contract.GetPairInfo(&_ReserveContract.CallOpts, source, dest)
}

// KyberNetwork is a free data retrieval call binding the contract method 0xb78b842d.
//
// Solidity: function kyberNetwork() constant returns(address)
func (_ReserveContract *ReserveContractCaller) KyberNetwork(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReserveContract.contract.Call(opts, out, "kyberNetwork")
	return *ret0, err
}

// KyberNetwork is a free data retrieval call binding the contract method 0xb78b842d.
//
// Solidity: function kyberNetwork() constant returns(address)
func (_ReserveContract *ReserveContractSession) KyberNetwork() (common.Address, error) {
	return _ReserveContract.Contract.KyberNetwork(&_ReserveContract.CallOpts)
}

// KyberNetwork is a free data retrieval call binding the contract method 0xb78b842d.
//
// Solidity: function kyberNetwork() constant returns(address)
func (_ReserveContract *ReserveContractCallerSession) KyberNetwork() (common.Address, error) {
	return _ReserveContract.Contract.KyberNetwork(&_ReserveContract.CallOpts)
}

// ReserveOwner is a free data retrieval call binding the contract method 0xe61b6557.
//
// Solidity: function reserveOwner() constant returns(address)
func (_ReserveContract *ReserveContractCaller) ReserveOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReserveContract.contract.Call(opts, out, "reserveOwner")
	return *ret0, err
}

// ReserveOwner is a free data retrieval call binding the contract method 0xe61b6557.
//
// Solidity: function reserveOwner() constant returns(address)
func (_ReserveContract *ReserveContractSession) ReserveOwner() (common.Address, error) {
	return _ReserveContract.Contract.ReserveOwner(&_ReserveContract.CallOpts)
}

// ReserveOwner is a free data retrieval call binding the contract method 0xe61b6557.
//
// Solidity: function reserveOwner() constant returns(address)
func (_ReserveContract *ReserveContractCallerSession) ReserveOwner() (common.Address, error) {
	return _ReserveContract.Contract.ReserveOwner(&_ReserveContract.CallOpts)
}

// TradeEnabled is a free data retrieval call binding the contract method 0xd621e813.
//
// Solidity: function tradeEnabled() constant returns(bool)
func (_ReserveContract *ReserveContractCaller) TradeEnabled(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReserveContract.contract.Call(opts, out, "tradeEnabled")
	return *ret0, err
}

// TradeEnabled is a free data retrieval call binding the contract method 0xd621e813.
//
// Solidity: function tradeEnabled() constant returns(bool)
func (_ReserveContract *ReserveContractSession) TradeEnabled() (bool, error) {
	return _ReserveContract.Contract.TradeEnabled(&_ReserveContract.CallOpts)
}

// TradeEnabled is a free data retrieval call binding the contract method 0xd621e813.
//
// Solidity: function tradeEnabled() constant returns(bool)
func (_ReserveContract *ReserveContractCallerSession) TradeEnabled() (bool, error) {
	return _ReserveContract.Contract.TradeEnabled(&_ReserveContract.CallOpts)
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() returns(bool)
func (_ReserveContract *ReserveContractTransactor) DepositEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveContract.contract.Transact(opts, "depositEther")
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() returns(bool)
func (_ReserveContract *ReserveContractSession) DepositEther() (*types.Transaction, error) {
	return _ReserveContract.Contract.DepositEther(&_ReserveContract.TransactOpts)
}

// DepositEther is a paid mutator transaction binding the contract method 0x98ea5fca.
//
// Solidity: function depositEther() returns(bool)
func (_ReserveContract *ReserveContractTransactorSession) DepositEther() (*types.Transaction, error) {
	return _ReserveContract.Contract.DepositEther(&_ReserveContract.TransactOpts)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(token address, amount uint256) returns(bool)
func (_ReserveContract *ReserveContractTransactor) DepositToken(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReserveContract.contract.Transact(opts, "depositToken", token, amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(token address, amount uint256) returns(bool)
func (_ReserveContract *ReserveContractSession) DepositToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReserveContract.Contract.DepositToken(&_ReserveContract.TransactOpts, token, amount)
}

// DepositToken is a paid mutator transaction binding the contract method 0x338b5dea.
//
// Solidity: function depositToken(token address, amount uint256) returns(bool)
func (_ReserveContract *ReserveContractTransactorSession) DepositToken(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ReserveContract.Contract.DepositToken(&_ReserveContract.TransactOpts, token, amount)
}

// EnableTrade is a paid mutator transaction binding the contract method 0xc00f04d1.
//
// Solidity: function enableTrade(enable bool) returns(bool)
func (_ReserveContract *ReserveContractTransactor) EnableTrade(opts *bind.TransactOpts, enable bool) (*types.Transaction, error) {
	return _ReserveContract.contract.Transact(opts, "enableTrade", enable)
}

// EnableTrade is a paid mutator transaction binding the contract method 0xc00f04d1.
//
// Solidity: function enableTrade(enable bool) returns(bool)
func (_ReserveContract *ReserveContractSession) EnableTrade(enable bool) (*types.Transaction, error) {
	return _ReserveContract.Contract.EnableTrade(&_ReserveContract.TransactOpts, enable)
}

// EnableTrade is a paid mutator transaction binding the contract method 0xc00f04d1.
//
// Solidity: function enableTrade(enable bool) returns(bool)
func (_ReserveContract *ReserveContractTransactorSession) EnableTrade(enable bool) (*types.Transaction, error) {
	return _ReserveContract.Contract.EnableTrade(&_ReserveContract.TransactOpts, enable)
}

// SetRate is a paid mutator transaction binding the contract method 0xda6dfea7.
//
// Solidity: function setRate(sources address[], dests address[], conversionRates uint256[], expiryBlocks uint256[], validate bool) returns(bool)
func (_ReserveContract *ReserveContractTransactor) SetRate(opts *bind.TransactOpts, sources []common.Address, dests []common.Address, conversionRates []*big.Int, expiryBlocks []*big.Int, validate bool) (*types.Transaction, error) {
	return _ReserveContract.contract.Transact(opts, "setRate", sources, dests, conversionRates, expiryBlocks, validate)
}

// SetRate is a paid mutator transaction binding the contract method 0xda6dfea7.
//
// Solidity: function setRate(sources address[], dests address[], conversionRates uint256[], expiryBlocks uint256[], validate bool) returns(bool)
func (_ReserveContract *ReserveContractSession) SetRate(sources []common.Address, dests []common.Address, conversionRates []*big.Int, expiryBlocks []*big.Int, validate bool) (*types.Transaction, error) {
	return _ReserveContract.Contract.SetRate(&_ReserveContract.TransactOpts, sources, dests, conversionRates, expiryBlocks, validate)
}

// SetRate is a paid mutator transaction binding the contract method 0xda6dfea7.
//
// Solidity: function setRate(sources address[], dests address[], conversionRates uint256[], expiryBlocks uint256[], validate bool) returns(bool)
func (_ReserveContract *ReserveContractTransactorSession) SetRate(sources []common.Address, dests []common.Address, conversionRates []*big.Int, expiryBlocks []*big.Int, validate bool) (*types.Transaction, error) {
	return _ReserveContract.Contract.SetRate(&_ReserveContract.TransactOpts, sources, dests, conversionRates, expiryBlocks, validate)
}

// Trade is a paid mutator transaction binding the contract method 0x0363b1e1.
//
// Solidity: function trade(sourceToken address, sourceAmount uint256, destToken address, destAddress address, validate bool) returns(bool)
func (_ReserveContract *ReserveContractTransactor) Trade(opts *bind.TransactOpts, sourceToken common.Address, sourceAmount *big.Int, destToken common.Address, destAddress common.Address, validate bool) (*types.Transaction, error) {
	return _ReserveContract.contract.Transact(opts, "trade", sourceToken, sourceAmount, destToken, destAddress, validate)
}

// Trade is a paid mutator transaction binding the contract method 0x0363b1e1.
//
// Solidity: function trade(sourceToken address, sourceAmount uint256, destToken address, destAddress address, validate bool) returns(bool)
func (_ReserveContract *ReserveContractSession) Trade(sourceToken common.Address, sourceAmount *big.Int, destToken common.Address, destAddress common.Address, validate bool) (*types.Transaction, error) {
	return _ReserveContract.Contract.Trade(&_ReserveContract.TransactOpts, sourceToken, sourceAmount, destToken, destAddress, validate)
}

// Trade is a paid mutator transaction binding the contract method 0x0363b1e1.
//
// Solidity: function trade(sourceToken address, sourceAmount uint256, destToken address, destAddress address, validate bool) returns(bool)
func (_ReserveContract *ReserveContractTransactorSession) Trade(sourceToken common.Address, sourceAmount *big.Int, destToken common.Address, destAddress common.Address, validate bool) (*types.Transaction, error) {
	return _ReserveContract.Contract.Trade(&_ReserveContract.TransactOpts, sourceToken, sourceAmount, destToken, destAddress, validate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(token address, amount uint256, destination address) returns(bool)
func (_ReserveContract *ReserveContractTransactor) Withdraw(opts *bind.TransactOpts, token common.Address, amount *big.Int, destination common.Address) (*types.Transaction, error) {
	return _ReserveContract.contract.Transact(opts, "withdraw", token, amount, destination)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(token address, amount uint256, destination address) returns(bool)
func (_ReserveContract *ReserveContractSession) Withdraw(token common.Address, amount *big.Int, destination common.Address) (*types.Transaction, error) {
	return _ReserveContract.Contract.Withdraw(&_ReserveContract.TransactOpts, token, amount, destination)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(token address, amount uint256, destination address) returns(bool)
func (_ReserveContract *ReserveContractTransactorSession) Withdraw(token common.Address, amount *big.Int, destination common.Address) (*types.Transaction, error) {
	return _ReserveContract.Contract.Withdraw(&_ReserveContract.TransactOpts, token, amount, destination)
}
