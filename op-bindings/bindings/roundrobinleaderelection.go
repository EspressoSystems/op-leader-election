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

// RoundRobinLeaderElectionMetaData contains all meta data concerning the RoundRobinLeaderElection contract.
var RoundRobinLeaderElectionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_n\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"addParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creation_block_number\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_leaderId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"isCurrentLeader\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"is_participant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_leaderId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"nextBlocksAsLeader\",\"outputs\":[{\"internalType\":\"enumILeaderElectionBatchInbox.LeaderStatusFlags\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"participants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b506040516108b63803806108b683398101604081905261002f9161003f565b3360805260005543600455610058565b60006020828403121561005157600080fd5b5051919050565b60805161084361007360003960006103b301526108436000f3fe608060405234801561001057600080fd5b50600436106100725760003560e01c8063a912546e11610050578063a912546e14610121578063d52f2d7614610134578063dfafe10f1461015557600080fd5b80632b9a3fed1461007757806335c1d34914610093578063a7c87067146100ee575b600080fd5b61008060045481565b6040519081526020015b60405180910390f35b6100c96100a13660046105e5565b60026020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161008a565b6101116100fc366004610627565b60036020526000908152604090205460ff1681565b604051901515815260200161008a565b61011161012f366004610649565b61016a565b610147610142366004610649565b6101e0565b60405161008a929190610673565b610168610163366004610627565b61039b565b005b600060045482101561017e575060006101da565b60015460045463ffffffff90911690600090829061019c9086610729565b6101a6919061073c565b60009081526002602052604090205473ffffffffffffffffffffffffffffffffffffffff86811691161492506101da915050565b92915050565b6040517fa7c8706700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152600090606090309063a7c8706790602401602060405180830381865afa158015610250573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102749190610777565b6102845750600290506060610394565b60408051600a8082526101608201909252600091602082016101408036833701905050905060005b600a81101561038c573063a912546e876102c68489610799565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff90921660048301526024820152604401602060405180830381865afa158015610334573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103589190610777565b82828151811061036a5761036a6107ac565b9115156020928302919091019091015280610384816107db565b9150506102ac565b506000925090505b9250929050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161461048b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605360248201527f526f756e64526f62696e4c6561646572456c656374696f6e3a206f6e6c79207460448201527f68652063726561746f72206f66207468697320636f6e74726163742063616e2060648201527f63616c6c20746869732066756e6374696f6e2e00000000000000000000000000608482015260a4015b60405180910390fd5b60005460015463ffffffff1610610524576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f526f756e64526f62696e4c6561646572456c656374696f6e3a206c697374206f60448201527f66207061727469636970616e74732069732066756c6c2e0000000000000000006064820152608401610482565b6001805463ffffffff908116600090815260026020908152604080832080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff881690811790915583526003909152812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001684179055825490911691906105c383610813565b91906101000a81548163ffffffff021916908363ffffffff1602179055505050565b6000602082840312156105f757600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461062257600080fd5b919050565b60006020828403121561063957600080fd5b610642826105fe565b9392505050565b6000806040838503121561065c57600080fd5b610665836105fe565b946020939093013593505050565b600060408201600385106106b0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8483526020604081850152818551808452606086019150828701935060005b818110156106ed5784511515835293830193918301916001016106cf565b5090979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b818103818111156101da576101da6106fa565b600082610772577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500690565b60006020828403121561078957600080fd5b8151801515811461064257600080fd5b808201808211156101da576101da6106fa565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361080c5761080c6106fa565b5060010190565b600063ffffffff80831681810361082c5761082c6106fa565b600101939250505056fea164736f6c6343000813000a",
}

// RoundRobinLeaderElectionABI is the input ABI used to generate the binding from.
// Deprecated: Use RoundRobinLeaderElectionMetaData.ABI instead.
var RoundRobinLeaderElectionABI = RoundRobinLeaderElectionMetaData.ABI

// RoundRobinLeaderElectionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoundRobinLeaderElectionMetaData.Bin instead.
var RoundRobinLeaderElectionBin = RoundRobinLeaderElectionMetaData.Bin

// DeployRoundRobinLeaderElection deploys a new Ethereum contract, binding an instance of RoundRobinLeaderElection to it.
func DeployRoundRobinLeaderElection(auth *bind.TransactOpts, backend bind.ContractBackend, _n *big.Int) (common.Address, *types.Transaction, *RoundRobinLeaderElection, error) {
	parsed, err := RoundRobinLeaderElectionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoundRobinLeaderElectionBin), backend, _n)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RoundRobinLeaderElection{RoundRobinLeaderElectionCaller: RoundRobinLeaderElectionCaller{contract: contract}, RoundRobinLeaderElectionTransactor: RoundRobinLeaderElectionTransactor{contract: contract}, RoundRobinLeaderElectionFilterer: RoundRobinLeaderElectionFilterer{contract: contract}}, nil
}

// RoundRobinLeaderElection is an auto generated Go binding around an Ethereum contract.
type RoundRobinLeaderElection struct {
	RoundRobinLeaderElectionCaller     // Read-only binding to the contract
	RoundRobinLeaderElectionTransactor // Write-only binding to the contract
	RoundRobinLeaderElectionFilterer   // Log filterer for contract events
}

// RoundRobinLeaderElectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type RoundRobinLeaderElectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoundRobinLeaderElectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RoundRobinLeaderElectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoundRobinLeaderElectionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RoundRobinLeaderElectionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RoundRobinLeaderElectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RoundRobinLeaderElectionSession struct {
	Contract     *RoundRobinLeaderElection // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RoundRobinLeaderElectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RoundRobinLeaderElectionCallerSession struct {
	Contract *RoundRobinLeaderElectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// RoundRobinLeaderElectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RoundRobinLeaderElectionTransactorSession struct {
	Contract     *RoundRobinLeaderElectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// RoundRobinLeaderElectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type RoundRobinLeaderElectionRaw struct {
	Contract *RoundRobinLeaderElection // Generic contract binding to access the raw methods on
}

// RoundRobinLeaderElectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RoundRobinLeaderElectionCallerRaw struct {
	Contract *RoundRobinLeaderElectionCaller // Generic read-only contract binding to access the raw methods on
}

// RoundRobinLeaderElectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RoundRobinLeaderElectionTransactorRaw struct {
	Contract *RoundRobinLeaderElectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoundRobinLeaderElection creates a new instance of RoundRobinLeaderElection, bound to a specific deployed contract.
func NewRoundRobinLeaderElection(address common.Address, backend bind.ContractBackend) (*RoundRobinLeaderElection, error) {
	contract, err := bindRoundRobinLeaderElection(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RoundRobinLeaderElection{RoundRobinLeaderElectionCaller: RoundRobinLeaderElectionCaller{contract: contract}, RoundRobinLeaderElectionTransactor: RoundRobinLeaderElectionTransactor{contract: contract}, RoundRobinLeaderElectionFilterer: RoundRobinLeaderElectionFilterer{contract: contract}}, nil
}

// NewRoundRobinLeaderElectionCaller creates a new read-only instance of RoundRobinLeaderElection, bound to a specific deployed contract.
func NewRoundRobinLeaderElectionCaller(address common.Address, caller bind.ContractCaller) (*RoundRobinLeaderElectionCaller, error) {
	contract, err := bindRoundRobinLeaderElection(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RoundRobinLeaderElectionCaller{contract: contract}, nil
}

// NewRoundRobinLeaderElectionTransactor creates a new write-only instance of RoundRobinLeaderElection, bound to a specific deployed contract.
func NewRoundRobinLeaderElectionTransactor(address common.Address, transactor bind.ContractTransactor) (*RoundRobinLeaderElectionTransactor, error) {
	contract, err := bindRoundRobinLeaderElection(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RoundRobinLeaderElectionTransactor{contract: contract}, nil
}

// NewRoundRobinLeaderElectionFilterer creates a new log filterer instance of RoundRobinLeaderElection, bound to a specific deployed contract.
func NewRoundRobinLeaderElectionFilterer(address common.Address, filterer bind.ContractFilterer) (*RoundRobinLeaderElectionFilterer, error) {
	contract, err := bindRoundRobinLeaderElection(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RoundRobinLeaderElectionFilterer{contract: contract}, nil
}

// bindRoundRobinLeaderElection binds a generic wrapper to an already deployed contract.
func bindRoundRobinLeaderElection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RoundRobinLeaderElectionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoundRobinLeaderElection.Contract.RoundRobinLeaderElectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoundRobinLeaderElection.Contract.RoundRobinLeaderElectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoundRobinLeaderElection.Contract.RoundRobinLeaderElectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RoundRobinLeaderElection.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoundRobinLeaderElection.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RoundRobinLeaderElection.Contract.contract.Transact(opts, method, params...)
}

// CreationBlockNumber is a free data retrieval call binding the contract method 0x2b9a3fed.
//
// Solidity: function creation_block_number() view returns(uint256)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCaller) CreationBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoundRobinLeaderElection.contract.Call(opts, &out, "creation_block_number")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreationBlockNumber is a free data retrieval call binding the contract method 0x2b9a3fed.
//
// Solidity: function creation_block_number() view returns(uint256)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionSession) CreationBlockNumber() (*big.Int, error) {
	return _RoundRobinLeaderElection.Contract.CreationBlockNumber(&_RoundRobinLeaderElection.CallOpts)
}

// CreationBlockNumber is a free data retrieval call binding the contract method 0x2b9a3fed.
//
// Solidity: function creation_block_number() view returns(uint256)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCallerSession) CreationBlockNumber() (*big.Int, error) {
	return _RoundRobinLeaderElection.Contract.CreationBlockNumber(&_RoundRobinLeaderElection.CallOpts)
}

// IsCurrentLeader is a free data retrieval call binding the contract method 0xa912546e.
//
// Solidity: function isCurrentLeader(address _leaderId, uint256 _blockNumber) view returns(bool)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCaller) IsCurrentLeader(opts *bind.CallOpts, _leaderId common.Address, _blockNumber *big.Int) (bool, error) {
	var out []interface{}
	err := _RoundRobinLeaderElection.contract.Call(opts, &out, "isCurrentLeader", _leaderId, _blockNumber)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrentLeader is a free data retrieval call binding the contract method 0xa912546e.
//
// Solidity: function isCurrentLeader(address _leaderId, uint256 _blockNumber) view returns(bool)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionSession) IsCurrentLeader(_leaderId common.Address, _blockNumber *big.Int) (bool, error) {
	return _RoundRobinLeaderElection.Contract.IsCurrentLeader(&_RoundRobinLeaderElection.CallOpts, _leaderId, _blockNumber)
}

// IsCurrentLeader is a free data retrieval call binding the contract method 0xa912546e.
//
// Solidity: function isCurrentLeader(address _leaderId, uint256 _blockNumber) view returns(bool)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCallerSession) IsCurrentLeader(_leaderId common.Address, _blockNumber *big.Int) (bool, error) {
	return _RoundRobinLeaderElection.Contract.IsCurrentLeader(&_RoundRobinLeaderElection.CallOpts, _leaderId, _blockNumber)
}

// IsParticipant is a free data retrieval call binding the contract method 0xa7c87067.
//
// Solidity: function is_participant(address ) view returns(bool)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCaller) IsParticipant(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _RoundRobinLeaderElection.contract.Call(opts, &out, "is_participant", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsParticipant is a free data retrieval call binding the contract method 0xa7c87067.
//
// Solidity: function is_participant(address ) view returns(bool)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionSession) IsParticipant(arg0 common.Address) (bool, error) {
	return _RoundRobinLeaderElection.Contract.IsParticipant(&_RoundRobinLeaderElection.CallOpts, arg0)
}

// IsParticipant is a free data retrieval call binding the contract method 0xa7c87067.
//
// Solidity: function is_participant(address ) view returns(bool)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCallerSession) IsParticipant(arg0 common.Address) (bool, error) {
	return _RoundRobinLeaderElection.Contract.IsParticipant(&_RoundRobinLeaderElection.CallOpts, arg0)
}

// NextBlocksAsLeader is a free data retrieval call binding the contract method 0xd52f2d76.
//
// Solidity: function nextBlocksAsLeader(address _leaderId, uint256 _blockNumber) view returns(uint8, bool[])
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCaller) NextBlocksAsLeader(opts *bind.CallOpts, _leaderId common.Address, _blockNumber *big.Int) (uint8, []bool, error) {
	var out []interface{}
	err := _RoundRobinLeaderElection.contract.Call(opts, &out, "nextBlocksAsLeader", _leaderId, _blockNumber)

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
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionSession) NextBlocksAsLeader(_leaderId common.Address, _blockNumber *big.Int) (uint8, []bool, error) {
	return _RoundRobinLeaderElection.Contract.NextBlocksAsLeader(&_RoundRobinLeaderElection.CallOpts, _leaderId, _blockNumber)
}

// NextBlocksAsLeader is a free data retrieval call binding the contract method 0xd52f2d76.
//
// Solidity: function nextBlocksAsLeader(address _leaderId, uint256 _blockNumber) view returns(uint8, bool[])
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCallerSession) NextBlocksAsLeader(_leaderId common.Address, _blockNumber *big.Int) (uint8, []bool, error) {
	return _RoundRobinLeaderElection.Contract.NextBlocksAsLeader(&_RoundRobinLeaderElection.CallOpts, _leaderId, _blockNumber)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCaller) Participants(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RoundRobinLeaderElection.contract.Call(opts, &out, "participants", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _RoundRobinLeaderElection.Contract.Participants(&_RoundRobinLeaderElection.CallOpts, arg0)
}

// Participants is a free data retrieval call binding the contract method 0x35c1d349.
//
// Solidity: function participants(uint256 ) view returns(address)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCallerSession) Participants(arg0 *big.Int) (common.Address, error) {
	return _RoundRobinLeaderElection.Contract.Participants(&_RoundRobinLeaderElection.CallOpts, arg0)
}

// AddParticipant is a paid mutator transaction binding the contract method 0xdfafe10f.
//
// Solidity: function addParticipant(address _addr) returns()
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionTransactor) AddParticipant(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _RoundRobinLeaderElection.contract.Transact(opts, "addParticipant", _addr)
}

// AddParticipant is a paid mutator transaction binding the contract method 0xdfafe10f.
//
// Solidity: function addParticipant(address _addr) returns()
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionSession) AddParticipant(_addr common.Address) (*types.Transaction, error) {
	return _RoundRobinLeaderElection.Contract.AddParticipant(&_RoundRobinLeaderElection.TransactOpts, _addr)
}

// AddParticipant is a paid mutator transaction binding the contract method 0xdfafe10f.
//
// Solidity: function addParticipant(address _addr) returns()
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionTransactorSession) AddParticipant(_addr common.Address) (*types.Transaction, error) {
	return _RoundRobinLeaderElection.Contract.AddParticipant(&_RoundRobinLeaderElection.TransactOpts, _addr)
}
