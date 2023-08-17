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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"HORIZON\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"addParticipant\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creation_block_number\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_max_number_participants\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_leaderId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"isCurrentLeader\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"is_participant\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"max_number_participants\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_leaderId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockNumber\",\"type\":\"uint256\"}],\"name\":\"nextBlocksAsLeader\",\"outputs\":[{\"internalType\":\"enumLeaderElectionBatchInbox.LeaderStatusFlags\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"participants\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060006080819052600160a05260c0819052620000329061dead9062000038565b62000338565b600054600290610100900460ff161580156200005b575060005460ff8083169116105b620000c45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805461ffff191660ff831617610100179055620000e26200013c565b620000ed83620001a4565b6066829055436065556000805461ff001916905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b600054610100900460ff16620001985760405162461bcd60e51b815260206004820152602b60248201526000805160206200133e83398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000bb565b620001a262000223565b565b620001ae6200028a565b6001600160a01b038116620002155760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401620000bb565b6200022081620002e6565b50565b600054610100900460ff166200027f5760405162461bcd60e51b815260206004820152602b60248201526000805160206200133e83398151915260448201526a6e697469616c697a696e6760a81b6064820152608401620000bb565b620001a233620002e6565b6033546001600160a01b03163314620001a25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401620000bb565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60805160a05160c051610fd66200036860003960006102b40152600061028b015260006102620152610fd66000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c8063a7c870671161008c578063cd6dc68711610066578063cd6dc68714610201578063d52f2d7614610214578063dfafe10f14610235578063f2fde38b1461024857600080fd5b8063a7c87067146101a1578063a912546e146101d4578063b444ef9d146101e757600080fd5b806354fd4d50116100bd57806354fd4d5014610164578063715018a6146101795780638da5cb5b1461018357600080fd5b80632b9a3fed146100e457806335c1d3491461010057806352f705481461015b575b600080fd5b6100ed60655481565b6040519081526020015b60405180910390f35b61013661010e366004610c2d565b60686020526000908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f7565b6100ed60665481565b61016c61025b565b6040516100f79190610c76565b6101816102fe565b005b60335473ffffffffffffffffffffffffffffffffffffffff16610136565b6101c46101af366004610cf0565b60696020526000908152604090205460ff1681565b60405190151581526020016100f7565b6101c46101e2366004610d12565b610312565b6101ef600a81565b60405160ff90911681526020016100f7565b61018161020f366004610d12565b610388565b610227610222366004610d12565b6104e9565b6040516100f7929190610d3c565b610181610243366004610cf0565b6106a4565b610181610256366004610cf0565b610807565b60606102867f00000000000000000000000000000000000000000000000000000000000000006108be565b6102af7f00000000000000000000000000000000000000000000000000000000000000006108be565b6102d87f00000000000000000000000000000000000000000000000000000000000000006108be565b6040516020016102ea93929190610dc3565b604051602081830303815290604052905090565b6103066109fb565b6103106000610a7c565b565b600060655482101561032657506000610382565b60675460655463ffffffff9091169060009082906103449086610e68565b61034e9190610eae565b60009081526068602052604090205473ffffffffffffffffffffffffffffffffffffffff8681169116149250610382915050565b92915050565b600054600290610100900460ff161580156103aa575060005460ff8083169116105b61043b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff831617610100179055610474610af3565b61047d83610807565b606682905543606555600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a1505050565b6040517fa7c8706700000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83166004820152600090606090309063a7c8706790602401602060405180830381865afa158015610559573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061057d9190610ec2565b61058d575060029050606061069d565b60408051600a8082526101608201909252600091602082016101408036833701905050905060005b600a811015610695573063a912546e876105cf8489610f13565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e085901b16815273ffffffffffffffffffffffffffffffffffffffff90921660048301526024820152604401602060405180830381865afa15801561063d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106619190610ec2565b82828151811061067357610673610f2b565b911515602092830291909101909101528061068d81610f5a565b9150506105b5565b506000925090505b9250929050565b6106ac6109fb565b60665460675463ffffffff1610610745576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f526f756e64526f62696e4c6561646572456c656374696f6e3a206c697374206f60448201527f66207061727469636970616e74732069732066756c6c2e0000000000000000006064820152608401610432565b6067805463ffffffff908116600090815260686020908152604080832080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff881690811790915583526069909152812080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055825490911691906107e583610f92565b91906101000a81548163ffffffff021916908363ffffffff1602179055505050565b61080f6109fb565b73ffffffffffffffffffffffffffffffffffffffff81166108b2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610432565b6108bb81610a7c565b50565b60608160000361090157505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561092b578061091581610f5a565b91506109249050600a83610fb5565b9150610905565b60008167ffffffffffffffff81111561094657610946610ee4565b6040519080825280601f01601f191660200182016040528015610970576020820181803683370190505b5090505b84156109f357610985600183610e68565b9150610992600a86610eae565b61099d906030610f13565b60f81b8183815181106109b2576109b2610f2b565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506109ec600a86610fb5565b9450610974565b949350505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610310576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610432565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16610b8a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610432565b610310600054610100900460ff16610c24576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610432565b61031033610a7c565b600060208284031215610c3f57600080fd5b5035919050565b60005b83811015610c61578181015183820152602001610c49565b83811115610c70576000848401525b50505050565b6020815260008251806020840152610c95816040850160208701610c46565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610ceb57600080fd5b919050565b600060208284031215610d0257600080fd5b610d0b82610cc7565b9392505050565b60008060408385031215610d2557600080fd5b610d2e83610cc7565b946020939093013593505050565b60006040820160038510610d79577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8483526020604081850152818551808452606086019150828701935060005b81811015610db6578451151583529383019391830191600101610d98565b5090979650505050505050565b60008451610dd5818460208901610c46565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610e11816001850160208a01610c46565b60019201918201528351610e2c816002840160208801610c46565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015610e7a57610e7a610e39565b500390565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610ebd57610ebd610e7f565b500690565b600060208284031215610ed457600080fd5b81518015158114610d0b57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008219821115610f2657610f26610e39565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610f8b57610f8b610e39565b5060010190565b600063ffffffff808316818103610fab57610fab610e39565b6001019392505050565b600082610fc457610fc4610e7f565b50049056fea164736f6c634300080f000a496e697469616c697a61626c653a20636f6e7472616374206973206e6f742069",
}

// RoundRobinLeaderElectionABI is the input ABI used to generate the binding from.
// Deprecated: Use RoundRobinLeaderElectionMetaData.ABI instead.
var RoundRobinLeaderElectionABI = RoundRobinLeaderElectionMetaData.ABI

// RoundRobinLeaderElectionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RoundRobinLeaderElectionMetaData.Bin instead.
var RoundRobinLeaderElectionBin = RoundRobinLeaderElectionMetaData.Bin

// DeployRoundRobinLeaderElection deploys a new Ethereum contract, binding an instance of RoundRobinLeaderElection to it.
func DeployRoundRobinLeaderElection(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RoundRobinLeaderElection, error) {
	parsed, err := RoundRobinLeaderElectionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RoundRobinLeaderElectionBin), backend)
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

// HORIZON is a free data retrieval call binding the contract method 0xb444ef9d.
//
// Solidity: function HORIZON() view returns(uint8)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCaller) HORIZON(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _RoundRobinLeaderElection.contract.Call(opts, &out, "HORIZON")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// HORIZON is a free data retrieval call binding the contract method 0xb444ef9d.
//
// Solidity: function HORIZON() view returns(uint8)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionSession) HORIZON() (uint8, error) {
	return _RoundRobinLeaderElection.Contract.HORIZON(&_RoundRobinLeaderElection.CallOpts)
}

// HORIZON is a free data retrieval call binding the contract method 0xb444ef9d.
//
// Solidity: function HORIZON() view returns(uint8)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCallerSession) HORIZON() (uint8, error) {
	return _RoundRobinLeaderElection.Contract.HORIZON(&_RoundRobinLeaderElection.CallOpts)
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

// MaxNumberParticipants is a free data retrieval call binding the contract method 0x52f70548.
//
// Solidity: function max_number_participants() view returns(uint256)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCaller) MaxNumberParticipants(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RoundRobinLeaderElection.contract.Call(opts, &out, "max_number_participants")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxNumberParticipants is a free data retrieval call binding the contract method 0x52f70548.
//
// Solidity: function max_number_participants() view returns(uint256)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionSession) MaxNumberParticipants() (*big.Int, error) {
	return _RoundRobinLeaderElection.Contract.MaxNumberParticipants(&_RoundRobinLeaderElection.CallOpts)
}

// MaxNumberParticipants is a free data retrieval call binding the contract method 0x52f70548.
//
// Solidity: function max_number_participants() view returns(uint256)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCallerSession) MaxNumberParticipants() (*big.Int, error) {
	return _RoundRobinLeaderElection.Contract.MaxNumberParticipants(&_RoundRobinLeaderElection.CallOpts)
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RoundRobinLeaderElection.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionSession) Owner() (common.Address, error) {
	return _RoundRobinLeaderElection.Contract.Owner(&_RoundRobinLeaderElection.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCallerSession) Owner() (common.Address, error) {
	return _RoundRobinLeaderElection.Contract.Owner(&_RoundRobinLeaderElection.CallOpts)
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

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RoundRobinLeaderElection.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionSession) Version() (string, error) {
	return _RoundRobinLeaderElection.Contract.Version(&_RoundRobinLeaderElection.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionCallerSession) Version() (string, error) {
	return _RoundRobinLeaderElection.Contract.Version(&_RoundRobinLeaderElection.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _owner, uint256 _max_number_participants) returns()
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _max_number_participants *big.Int) (*types.Transaction, error) {
	return _RoundRobinLeaderElection.contract.Transact(opts, "initialize", _owner, _max_number_participants)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _owner, uint256 _max_number_participants) returns()
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionSession) Initialize(_owner common.Address, _max_number_participants *big.Int) (*types.Transaction, error) {
	return _RoundRobinLeaderElection.Contract.Initialize(&_RoundRobinLeaderElection.TransactOpts, _owner, _max_number_participants)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _owner, uint256 _max_number_participants) returns()
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionTransactorSession) Initialize(_owner common.Address, _max_number_participants *big.Int) (*types.Transaction, error) {
	return _RoundRobinLeaderElection.Contract.Initialize(&_RoundRobinLeaderElection.TransactOpts, _owner, _max_number_participants)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RoundRobinLeaderElection.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionSession) RenounceOwnership() (*types.Transaction, error) {
	return _RoundRobinLeaderElection.Contract.RenounceOwnership(&_RoundRobinLeaderElection.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RoundRobinLeaderElection.Contract.RenounceOwnership(&_RoundRobinLeaderElection.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RoundRobinLeaderElection.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RoundRobinLeaderElection.Contract.TransferOwnership(&_RoundRobinLeaderElection.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RoundRobinLeaderElection.Contract.TransferOwnership(&_RoundRobinLeaderElection.TransactOpts, newOwner)
}

// RoundRobinLeaderElectionInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RoundRobinLeaderElection contract.
type RoundRobinLeaderElectionInitializedIterator struct {
	Event *RoundRobinLeaderElectionInitialized // Event containing the contract specifics and raw log

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
func (it *RoundRobinLeaderElectionInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundRobinLeaderElectionInitialized)
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
		it.Event = new(RoundRobinLeaderElectionInitialized)
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
func (it *RoundRobinLeaderElectionInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundRobinLeaderElectionInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundRobinLeaderElectionInitialized represents a Initialized event raised by the RoundRobinLeaderElection contract.
type RoundRobinLeaderElectionInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionFilterer) FilterInitialized(opts *bind.FilterOpts) (*RoundRobinLeaderElectionInitializedIterator, error) {

	logs, sub, err := _RoundRobinLeaderElection.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RoundRobinLeaderElectionInitializedIterator{contract: _RoundRobinLeaderElection.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RoundRobinLeaderElectionInitialized) (event.Subscription, error) {

	logs, sub, err := _RoundRobinLeaderElection.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundRobinLeaderElectionInitialized)
				if err := _RoundRobinLeaderElection.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionFilterer) ParseInitialized(log types.Log) (*RoundRobinLeaderElectionInitialized, error) {
	event := new(RoundRobinLeaderElectionInitialized)
	if err := _RoundRobinLeaderElection.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RoundRobinLeaderElectionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RoundRobinLeaderElection contract.
type RoundRobinLeaderElectionOwnershipTransferredIterator struct {
	Event *RoundRobinLeaderElectionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RoundRobinLeaderElectionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RoundRobinLeaderElectionOwnershipTransferred)
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
		it.Event = new(RoundRobinLeaderElectionOwnershipTransferred)
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
func (it *RoundRobinLeaderElectionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RoundRobinLeaderElectionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RoundRobinLeaderElectionOwnershipTransferred represents a OwnershipTransferred event raised by the RoundRobinLeaderElection contract.
type RoundRobinLeaderElectionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RoundRobinLeaderElectionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RoundRobinLeaderElection.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RoundRobinLeaderElectionOwnershipTransferredIterator{contract: _RoundRobinLeaderElection.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RoundRobinLeaderElectionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RoundRobinLeaderElection.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RoundRobinLeaderElectionOwnershipTransferred)
				if err := _RoundRobinLeaderElection.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RoundRobinLeaderElection *RoundRobinLeaderElectionFilterer) ParseOwnershipTransferred(log types.Log) (*RoundRobinLeaderElectionOwnershipTransferred, error) {
	event := new(RoundRobinLeaderElectionOwnershipTransferred)
	if err := _RoundRobinLeaderElection.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
