// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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
)

// LeaderElectionBatchInboxMetaData contains all meta data concerning the LeaderElectionBatchInbox contract.
var LeaderElectionBatchInboxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_leaderId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"isCurrentLeader\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_leaderId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"nextBlocksAsLeader\",\"outputs\":[{\"internalType\":\"enumLeaderElectionBatchInbox.LeaderStatusFlags\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LeaderElectionBatchInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use LeaderElectionBatchInboxMetaData.ABI instead.
var LeaderElectionBatchInboxABI = LeaderElectionBatchInboxMetaData.ABI

// LeaderElectionBatchInbox is an auto generated Go binding around an Ethereum contract.
type LeaderElectionBatchInbox struct {
	LeaderElectionBatchInboxCaller     // Read-only binding to the contract
	LeaderElectionBatchInboxTransactor // Write-only binding to the contract
	LeaderElectionBatchInboxFilterer   // Log filterer for contract events
}

// LeaderElectionBatchInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type LeaderElectionBatchInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeaderElectionBatchInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LeaderElectionBatchInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeaderElectionBatchInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LeaderElectionBatchInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeaderElectionBatchInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LeaderElectionBatchInboxSession struct {
	Contract     *LeaderElectionBatchInbox // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// LeaderElectionBatchInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LeaderElectionBatchInboxCallerSession struct {
	Contract *LeaderElectionBatchInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// LeaderElectionBatchInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LeaderElectionBatchInboxTransactorSession struct {
	Contract     *LeaderElectionBatchInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// LeaderElectionBatchInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type LeaderElectionBatchInboxRaw struct {
	Contract *LeaderElectionBatchInbox // Generic contract binding to access the raw methods on
}

// LeaderElectionBatchInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LeaderElectionBatchInboxCallerRaw struct {
	Contract *LeaderElectionBatchInboxCaller // Generic read-only contract binding to access the raw methods on
}

// LeaderElectionBatchInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LeaderElectionBatchInboxTransactorRaw struct {
	Contract *LeaderElectionBatchInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLeaderElectionBatchInbox creates a new instance of LeaderElectionBatchInbox, bound to a specific deployed contract.
func NewLeaderElectionBatchInbox(address common.Address, backend bind.ContractBackend) (*LeaderElectionBatchInbox, error) {
	contract, err := bindLeaderElectionBatchInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LeaderElectionBatchInbox{LeaderElectionBatchInboxCaller: LeaderElectionBatchInboxCaller{contract: contract}, LeaderElectionBatchInboxTransactor: LeaderElectionBatchInboxTransactor{contract: contract}, LeaderElectionBatchInboxFilterer: LeaderElectionBatchInboxFilterer{contract: contract}}, nil
}

// NewLeaderElectionBatchInboxCaller creates a new read-only instance of LeaderElectionBatchInbox, bound to a specific deployed contract.
func NewLeaderElectionBatchInboxCaller(address common.Address, caller bind.ContractCaller) (*LeaderElectionBatchInboxCaller, error) {
	contract, err := bindLeaderElectionBatchInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LeaderElectionBatchInboxCaller{contract: contract}, nil
}

// NewLeaderElectionBatchInboxTransactor creates a new write-only instance of LeaderElectionBatchInbox, bound to a specific deployed contract.
func NewLeaderElectionBatchInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*LeaderElectionBatchInboxTransactor, error) {
	contract, err := bindLeaderElectionBatchInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LeaderElectionBatchInboxTransactor{contract: contract}, nil
}

// NewLeaderElectionBatchInboxFilterer creates a new log filterer instance of LeaderElectionBatchInbox, bound to a specific deployed contract.
func NewLeaderElectionBatchInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*LeaderElectionBatchInboxFilterer, error) {
	contract, err := bindLeaderElectionBatchInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LeaderElectionBatchInboxFilterer{contract: contract}, nil
}

// bindLeaderElectionBatchInbox binds a generic wrapper to an already deployed contract.
func bindLeaderElectionBatchInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LeaderElectionBatchInboxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeaderElectionBatchInbox *LeaderElectionBatchInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeaderElectionBatchInbox.Contract.LeaderElectionBatchInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeaderElectionBatchInbox *LeaderElectionBatchInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeaderElectionBatchInbox.Contract.LeaderElectionBatchInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeaderElectionBatchInbox *LeaderElectionBatchInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeaderElectionBatchInbox.Contract.LeaderElectionBatchInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LeaderElectionBatchInbox *LeaderElectionBatchInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LeaderElectionBatchInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LeaderElectionBatchInbox *LeaderElectionBatchInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LeaderElectionBatchInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LeaderElectionBatchInbox *LeaderElectionBatchInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LeaderElectionBatchInbox.Contract.contract.Transact(opts, method, params...)
}

// IsCurrentLeader is a free data retrieval call binding the contract method 0xa912546e.
//
// Solidity: function isCurrentLeader(address _leaderId, uint256 _blockNumber) view returns(bool)
func (_LeaderElectionBatchInbox *LeaderElectionBatchInboxCaller) IsCurrentLeader(opts *bind.CallOpts, _leaderId common.Address, _blockNumber *big.Int) (bool, error) {
	var out []interface{}
	err := _LeaderElectionBatchInbox.contract.Call(opts, &out, "isCurrentLeader", _leaderId, _blockNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrentLeader is a free data retrieval call binding the contract method 0xa912546e.
//
// Solidity: function isCurrentLeader(address _leaderId, uint256 _blockNumber) view returns(bool)
func (_LeaderElectionBatchInbox *LeaderElectionBatchInboxSession) IsCurrentLeader(_leaderId common.Address, _blockNumber *big.Int) (bool, error) {
	return _LeaderElectionBatchInbox.Contract.IsCurrentLeader(&_LeaderElectionBatchInbox.CallOpts, _leaderId, _blockNumber)
}

// IsCurrentLeader is a free data retrieval call binding the contract method 0xa912546e.
//
// Solidity: function isCurrentLeader(address _leaderId, uint256 _blockNumber) view returns(bool)
func (_LeaderElectionBatchInbox *LeaderElectionBatchInboxCallerSession) IsCurrentLeader(_leaderId common.Address, _blockNumber *big.Int) (bool, error) {
	return _LeaderElectionBatchInbox.Contract.IsCurrentLeader(&_LeaderElectionBatchInbox.CallOpts, _leaderId, _blockNumber)
}

// NextBlocksAsLeader is a free data retrieval call binding the contract method 0xd52f2d76.
//
// Solidity: function nextBlocksAsLeader(address _leaderId, uint256 _blockNumber) view returns(uint8, bool[])
func (_LeaderElectionBatchInbox *LeaderElectionBatchInboxCaller) NextBlocksAsLeader(opts *bind.CallOpts, _leaderId common.Address, _blockNumber *big.Int) (uint8, []bool, error) {
	var out []interface{}
	err := _LeaderElectionBatchInbox.contract.Call(opts, &out, "nextBlocksAsLeader", _leaderId, _blockNumber)

	if err != nil {
		return *new(uint8), *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new([]bool)).(*[]bool)

	return out0, out1, err

}

// NextBlocksAsLeader is a free data retrieval call binding the contract method 0xd52f2d76.
//
// Solidity: function nextBlocksAsLeader(address _leaderId, uint256 _blockNumber) view returns(uint8, bool[])
func (_LeaderElectionBatchInbox *LeaderElectionBatchInboxSession) NextBlocksAsLeader(_leaderId common.Address, _blockNumber *big.Int) (uint8, []bool, error) {
	return _LeaderElectionBatchInbox.Contract.NextBlocksAsLeader(&_LeaderElectionBatchInbox.CallOpts, _leaderId, _blockNumber)
}

// NextBlocksAsLeader is a free data retrieval call binding the contract method 0xd52f2d76.
//
// Solidity: function nextBlocksAsLeader(address _leaderId, uint256 _blockNumber) view returns(uint8, bool[])
func (_LeaderElectionBatchInbox *LeaderElectionBatchInboxCallerSession) NextBlocksAsLeader(_leaderId common.Address, _blockNumber *big.Int) (uint8, []bool, error) {
	return _LeaderElectionBatchInbox.Contract.NextBlocksAsLeader(&_LeaderElectionBatchInbox.CallOpts, _leaderId, _blockNumber)
}
