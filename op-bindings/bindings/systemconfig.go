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

// ResourceMeteringResourceConfig is an auto generated low-level Go binding around an user-defined struct.
type ResourceMeteringResourceConfig struct {
	MaxResourceLimit            uint32
	ElasticityMultiplier        uint8
	BaseFeeMaxChangeDenominator uint8
	MinimumBaseFee              uint32
	SystemTxMaxGas              uint32
	MaximumBaseFee              *big.Int
}

// SystemConfigMetaData contains all meta data concerning the SystemConfig contract.
var SystemConfigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_overhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_scalar\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_batcherHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_unsafeBlockSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_batchInboxAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxResourceLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"elasticityMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"baseFeeMaxChangeDenominator\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"minimumBaseFee\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"systemTxMaxGas\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"maximumBaseFee\",\"type\":\"uint128\"}],\"internalType\":\"structResourceMetering.ResourceConfig\",\"name\":\"_config\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"enumSystemConfig.UpdateType\",\"name\":\"updateType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ConfigUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"UNSAFE_BLOCK_SIGNER_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batchInboxAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"batcherHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_overhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_scalar\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_batcherHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"_gasLimit\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"_unsafeBlockSigner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_batchInboxAddr\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxResourceLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"elasticityMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"baseFeeMaxChangeDenominator\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"minimumBaseFee\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"systemTxMaxGas\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"maximumBaseFee\",\"type\":\"uint128\"}],\"internalType\":\"structResourceMetering.ResourceConfig\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumGasLimit\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"overhead\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resourceConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxResourceLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"elasticityMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"baseFeeMaxChangeDenominator\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"minimumBaseFee\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"systemTxMaxGas\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"maximumBaseFee\",\"type\":\"uint128\"}],\"internalType\":\"structResourceMetering.ResourceConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scalar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_batchInboxAddr\",\"type\":\"address\"}],\"name\":\"setBatchInboxAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_batcherHash\",\"type\":\"bytes32\"}],\"name\":\"setBatcherHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_overhead\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_scalar\",\"type\":\"uint256\"}],\"name\":\"setGasConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_gasLimit\",\"type\":\"uint64\"}],\"name\":\"setGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"maxResourceLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"elasticityMultiplier\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"baseFeeMaxChangeDenominator\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"minimumBaseFee\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"systemTxMaxGas\",\"type\":\"uint32\"},{\"internalType\":\"uint128\",\"name\":\"maximumBaseFee\",\"type\":\"uint128\"}],\"internalType\":\"structResourceMetering.ResourceConfig\",\"name\":\"_config\",\"type\":\"tuple\"}],\"name\":\"setResourceConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_unsafeBlockSigner\",\"type\":\"address\"}],\"name\":\"setUnsafeBlockSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unsafeBlockSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b506040516200245f3803806200245f833981016040819052620000349162000885565b6001608052600460a052600060c05262000055888888888888888862000063565b505050505050505062000a98565b600054610100900460ff1615808015620000845750600054600160ff909116105b80620000b45750620000a1306200029c60201b62000c181760201c565b158015620000b4575060005460ff166001145b6200011d5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff19166001179055801562000141576000805461ff0019166101001790555b6200014b620002ab565b620001568962000313565b606588905560668790556067869055606880546001600160401b0319166001600160401b038716179055620001a9847f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0855565b60688054600160401b600160e01b031916680100000000000000006001600160a01b03861602179055620001dd8262000392565b620001e7620006e7565b6001600160401b0316856001600160401b031610156200024a5760405162461bcd60e51b815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f7700604482015260640162000114565b801562000291576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050505050565b6001600160a01b03163b151590565b600054610100900460ff16620003075760405162461bcd60e51b815260206004820152602b60248201526000805160206200243f83398151915260448201526a6e697469616c697a696e6760a81b606482015260840162000114565b6200031162000714565b565b6200031d6200077b565b6001600160a01b038116620003845760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b606482015260840162000114565b6200038f81620007d7565b50565b8060a001516001600160801b0316816060015163ffffffff161115620004215760405162461bcd60e51b815260206004820152603560248201527f53797374656d436f6e6669673a206d696e206261736520666565206d7573742060448201527f6265206c657373207468616e206d617820626173650000000000000000000000606482015260840162000114565b6001816040015160ff1611620004925760405162461bcd60e51b815260206004820152602f60248201527f53797374656d436f6e6669673a2064656e6f6d696e61746f72206d757374206260448201526e65206c6172676572207468616e203160881b606482015260840162000114565b606854608082015182516001600160401b0390921691620004b49190620009e7565b63ffffffff1611156200050a5760405162461bcd60e51b815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f7700604482015260640162000114565b6000816020015160ff16116200057b5760405162461bcd60e51b815260206004820152602f60248201527f53797374656d436f6e6669673a20656c6173746963697479206d756c7469706c60448201526e06965722063616e6e6f74206265203608c1b606482015260840162000114565b8051602082015163ffffffff82169160ff909116906200059d90829062000a12565b620005a9919062000a44565b63ffffffff1614620006245760405162461bcd60e51b815260206004820152603760248201527f53797374656d436f6e6669673a20707265636973696f6e206c6f73732077697460448201527f6820746172676574207265736f75726365206c696d6974000000000000000000606482015260840162000114565b805160698054602084015160408501516060860151608087015160a09097015163ffffffff96871664ffffffffff199095169490941764010000000060ff948516021764ffffffffff60281b191665010000000000939092169290920263ffffffff60301b19161766010000000000009185169190910217600160501b600160f01b0319166a01000000000000000000009390941692909202600160701b600160f01b03191692909217600160701b6001600160801b0390921691909102179055565b6069546000906200070f9063ffffffff6a010000000000000000000082048116911662000a73565b905090565b600054610100900460ff16620007705760405162461bcd60e51b815260206004820152602b60248201526000805160206200243f83398151915260448201526a6e697469616c697a696e6760a81b606482015260840162000114565b6200031133620007d7565b6033546001600160a01b03163314620003115760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640162000114565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b80516001600160a01b03811681146200084157600080fd5b919050565b805163ffffffff811681146200084157600080fd5b805160ff811681146200084157600080fd5b80516001600160801b03811681146200084157600080fd5b600080600080600080600080888a036101a0811215620008a457600080fd5b620008af8a62000829565b60208b015160408c015160608d015160808e0151939c50919a50985096506001600160401b038082168214620008e457600080fd5b819650620008f560a08d0162000829565b95506200090560c08d0162000829565b945060c060df19840112156200091a57600080fd5b604051925060c08301915082821081831117156200094857634e487b7160e01b600052604160045260246000fd5b506040526200095a60e08b0162000846565b81526200096b6101008b016200085b565b60208201526200097f6101208b016200085b565b6040820152620009936101408b0162000846565b6060820152620009a76101608b0162000846565b6080820152620009bb6101808b016200086d565b60a0820152809150509295985092959890939650565b634e487b7160e01b600052601160045260246000fd5b600063ffffffff80831681851680830382111562000a095762000a09620009d1565b01949350505050565b600063ffffffff8084168062000a3857634e487b7160e01b600052601260045260246000fd5b92169190910492915050565b600063ffffffff8083168185168183048111821515161562000a6a5762000a6a620009d1565b02949350505050565b60006001600160401b0382811684821680830382111562000a095762000a09620009d1565b60805160a05160c05161197762000ac860003960006105d3015260006105aa0152600061058101526119776000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c8063a0ca2cfb116100d8578063cc731b021161008c578063f45e65d811610066578063f45e65d814610466578063f68016b71461046f578063ffa1ad741461048357600080fd5b8063cc731b0214610316578063e81b2c6d1461044a578063f2fde38b1461045357600080fd5b8063c495de62116100bd578063c495de62146102c4578063c71973f6146102f0578063c9b26f611461030357600080fd5b8063a0ca2cfb1461029e578063b40a817c146102b157600080fd5b806354fd4d501161012f578063715018a611610114578063715018a6146102655780638da5cb5b1461026d578063935f029e1461028b57600080fd5b806354fd4d501461023d57806358619b901461025257600080fd5b80631fd19ee1116101605780631fd19ee1146101ad5780634add321d146101f55780634f16540b1461021657600080fd5b80630c18c1621461017c57806318d1391814610198575b600080fd5b61018560655481565b6040519081526020015b60405180910390f35b6101ab6101a6366004611445565b61048b565b005b7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c08545b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161018f565b6101fd61054f565b60405167ffffffffffffffff909116815260200161018f565b6101857f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0881565b61024561057a565b60405161018f91906114e1565b6101ab610260366004611649565b61061d565b6101ab6108df565b60335473ffffffffffffffffffffffffffffffffffffffff166101d0565b6101ab6102993660046116cc565b6108f3565b6101ab6102ac366004611445565b61098c565b6101ab6102bf3660046116ee565b610a3a565b6068546101d09068010000000000000000900473ffffffffffffffffffffffffffffffffffffffff1681565b6101ab6102fe366004611709565b610b20565b6101ab610311366004611725565b610b34565b6103da6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260695463ffffffff8082168352640100000000820460ff9081166020850152650100000000008304169383019390935266010000000000008104831660608301526a0100000000000000000000810490921660808201526e0100000000000000000000000000009091046fffffffffffffffffffffffffffffffff1660a082015290565b60405161018f9190600060c08201905063ffffffff80845116835260ff602085015116602084015260ff6040850151166040840152806060850151166060840152806080850151166080840152506fffffffffffffffffffffffffffffffff60a08401511660a083015292915050565b61018560675481565b6101ab610461366004611445565b610b64565b61018560665481565b6068546101fd9067ffffffffffffffff1681565b610185600081565b610493610c34565b6104bb817f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0855565b6040805173ffffffffffffffffffffffffffffffffffffffff8316602082015260009101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052905060035b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be8360405161054391906114e1565b60405180910390a35050565b6069546000906105759063ffffffff6a010000000000000000000082048116911661176d565b905090565b60606105a57f0000000000000000000000000000000000000000000000000000000000000000610cb5565b6105ce7f0000000000000000000000000000000000000000000000000000000000000000610cb5565b6105f77f0000000000000000000000000000000000000000000000000000000000000000610cb5565b60405160200161060993929190611799565b604051602081830303815290604052905090565b600054610100900460ff161580801561063d5750600054600160ff909116105b806106575750303b158015610657575060005460ff166001145b6106e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561074657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b61074e610df2565b61075789610b64565b606588905560668790556067869055606880547f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0886905573ffffffffffffffffffffffffffffffffffffffff851668010000000000000000027fffffffff0000000000000000000000000000000000000000000000000000000090911667ffffffffffffffff8816171790556107ec82610e91565b6107f461054f565b67ffffffffffffffff168567ffffffffffffffff161015610871576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016106df565b80156108d457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050505050505050565b6108e7610c34565b6108f16000611305565b565b6108fb610c34565b606582905560668190556040805160208101849052908101829052600090606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be8360405161097f91906114e1565b60405180910390a3505050565b610994610c34565b606880547fffffffff0000000000000000000000000000000000000000ffffffffffffffff166801000000000000000073ffffffffffffffffffffffffffffffffffffffff8416021790556040805173ffffffffffffffffffffffffffffffffffffffff8316602082015260009101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081840301815291905290506004610512565b610a42610c34565b610a4a61054f565b67ffffffffffffffff168167ffffffffffffffff161015610ac7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016106df565b606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83169081179091556040805160208082019390935281518082039093018352810190526002610512565b610b28610c34565b610b3181610e91565b50565b610b3c610c34565b6067819055604080516020808201849052825180830390910181529082019091526000610512565b610b6c610c34565b73ffffffffffffffffffffffffffffffffffffffff8116610c0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016106df565b610b3181611305565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b60335473ffffffffffffffffffffffffffffffffffffffff1633146108f1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016106df565b606081600003610cf857505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610d225780610d0c8161180f565b9150610d1b9050600a83611876565b9150610cfc565b60008167ffffffffffffffff811115610d3d57610d3d61150c565b6040519080825280601f01601f191660200182016040528015610d67576020820181803683370190505b5090505b8415610dea57610d7c60018361188a565b9150610d89600a866118a1565b610d949060306118b5565b60f81b818381518110610da957610da96118cd565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610de3600a86611876565b9450610d6b565b949350505050565b600054610100900460ff16610e89576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106df565b6108f161137c565b8060a001516fffffffffffffffffffffffffffffffff16816060015163ffffffff161115610f41576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f53797374656d436f6e6669673a206d696e206261736520666565206d7573742060448201527f6265206c657373207468616e206d61782062617365000000000000000000000060648201526084016106df565b6001816040015160ff1611610fd8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a2064656e6f6d696e61746f72206d757374206260448201527f65206c6172676572207468616e2031000000000000000000000000000000000060648201526084016106df565b6068546080820151825167ffffffffffffffff90921691610ff991906118fc565b63ffffffff161115611067576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016106df565b6000816020015160ff16116110fe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a20656c6173746963697479206d756c7469706c60448201527f6965722063616e6e6f742062652030000000000000000000000000000000000060648201526084016106df565b8051602082015163ffffffff82169160ff9091169061111e90829061191b565b611128919061193e565b63ffffffff16146111bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f53797374656d436f6e6669673a20707265636973696f6e206c6f73732077697460448201527f6820746172676574207265736f75726365206c696d697400000000000000000060648201526084016106df565b805160698054602084015160408501516060860151608087015160a09097015163ffffffff9687167fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009095169490941764010000000060ff94851602177fffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffff166501000000000093909216929092027fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1617660100000000000091851691909102177fffff0000000000000000000000000000000000000000ffffffffffffffffffff166a010000000000000000000093909416929092027fffff00000000000000000000000000000000ffffffffffffffffffffffffffff16929092176e0100000000000000000000000000006fffffffffffffffffffffffffffffffff90921691909102179055565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16611413576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016106df565b6108f133611305565b803573ffffffffffffffffffffffffffffffffffffffff8116811461144057600080fd5b919050565b60006020828403121561145757600080fd5b6114608261141c565b9392505050565b60005b8381101561148257818101518382015260200161146a565b83811115611491576000848401525b50505050565b600081518084526114af816020860160208601611467565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006114606020830184611497565b803567ffffffffffffffff8116811461144057600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b803563ffffffff8116811461144057600080fd5b803560ff8116811461144057600080fd5b80356fffffffffffffffffffffffffffffffff8116811461144057600080fd5b600060c0828403121561159257600080fd5b60405160c0810181811067ffffffffffffffff821117156115dc577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040529050806115eb8361153b565b81526115f96020840161154f565b602082015261160a6040840161154f565b604082015261161b6060840161153b565b606082015261162c6080840161153b565b608082015261163d60a08401611560565b60a08201525092915050565b6000806000806000806000806101a0898b03121561166657600080fd5b61166f8961141c565b975060208901359650604089013595506060890135945061169260808a016114f4565b93506116a060a08a0161141c565b92506116ae60c08a0161141c565b91506116bd8a60e08b01611580565b90509295985092959890939650565b600080604083850312156116df57600080fd5b50508035926020909101359150565b60006020828403121561170057600080fd5b611460826114f4565b600060c0828403121561171b57600080fd5b6114608383611580565b60006020828403121561173757600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff8083168185168083038211156117905761179061173e565b01949350505050565b600084516117ab818460208901611467565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516117e7816001850160208a01611467565b60019201918201528351611802816002840160208801611467565b0160020195945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036118405761184061173e565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261188557611885611847565b500490565b60008282101561189c5761189c61173e565b500390565b6000826118b0576118b0611847565b500690565b600082198211156118c8576118c861173e565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600063ffffffff8083168185168083038211156117905761179061173e565b600063ffffffff8084168061193257611932611847565b92169190910492915050565b600063ffffffff808316818516818304811182151516156119615761196161173e565b0294935050505056fea164736f6c634300080f000a496e697469616c697a61626c653a20636f6e7472616374206973206e6f742069",
}

// SystemConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemConfigMetaData.ABI instead.
var SystemConfigABI = SystemConfigMetaData.ABI

// SystemConfigBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SystemConfigMetaData.Bin instead.
var SystemConfigBin = SystemConfigMetaData.Bin

// DeploySystemConfig deploys a new Ethereum contract, binding an instance of SystemConfig to it.
func DeploySystemConfig(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _overhead *big.Int, _scalar *big.Int, _batcherHash [32]byte, _gasLimit uint64, _unsafeBlockSigner common.Address, _batchInboxAddr common.Address, _config ResourceMeteringResourceConfig) (common.Address, *types.Transaction, *SystemConfig, error) {
	parsed, err := SystemConfigMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SystemConfigBin), backend, _owner, _overhead, _scalar, _batcherHash, _gasLimit, _unsafeBlockSigner, _batchInboxAddr, _config)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SystemConfig{SystemConfigCaller: SystemConfigCaller{contract: contract}, SystemConfigTransactor: SystemConfigTransactor{contract: contract}, SystemConfigFilterer: SystemConfigFilterer{contract: contract}}, nil
}

// SystemConfig is an auto generated Go binding around an Ethereum contract.
type SystemConfig struct {
	SystemConfigCaller     // Read-only binding to the contract
	SystemConfigTransactor // Write-only binding to the contract
	SystemConfigFilterer   // Log filterer for contract events
}

// SystemConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemConfigSession struct {
	Contract     *SystemConfig     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SystemConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemConfigCallerSession struct {
	Contract *SystemConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SystemConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemConfigTransactorSession struct {
	Contract     *SystemConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SystemConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemConfigRaw struct {
	Contract *SystemConfig // Generic contract binding to access the raw methods on
}

// SystemConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemConfigCallerRaw struct {
	Contract *SystemConfigCaller // Generic read-only contract binding to access the raw methods on
}

// SystemConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemConfigTransactorRaw struct {
	Contract *SystemConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemConfig creates a new instance of SystemConfig, bound to a specific deployed contract.
func NewSystemConfig(address common.Address, backend bind.ContractBackend) (*SystemConfig, error) {
	contract, err := bindSystemConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemConfig{SystemConfigCaller: SystemConfigCaller{contract: contract}, SystemConfigTransactor: SystemConfigTransactor{contract: contract}, SystemConfigFilterer: SystemConfigFilterer{contract: contract}}, nil
}

// NewSystemConfigCaller creates a new read-only instance of SystemConfig, bound to a specific deployed contract.
func NewSystemConfigCaller(address common.Address, caller bind.ContractCaller) (*SystemConfigCaller, error) {
	contract, err := bindSystemConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemConfigCaller{contract: contract}, nil
}

// NewSystemConfigTransactor creates a new write-only instance of SystemConfig, bound to a specific deployed contract.
func NewSystemConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemConfigTransactor, error) {
	contract, err := bindSystemConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemConfigTransactor{contract: contract}, nil
}

// NewSystemConfigFilterer creates a new log filterer instance of SystemConfig, bound to a specific deployed contract.
func NewSystemConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemConfigFilterer, error) {
	contract, err := bindSystemConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemConfigFilterer{contract: contract}, nil
}

// bindSystemConfig binds a generic wrapper to an already deployed contract.
func bindSystemConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemConfigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemConfig *SystemConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemConfig.Contract.SystemConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemConfig *SystemConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemConfig.Contract.SystemConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemConfig *SystemConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemConfig.Contract.SystemConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemConfig *SystemConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemConfig *SystemConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemConfig *SystemConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemConfig.Contract.contract.Transact(opts, method, params...)
}

// UNSAFEBLOCKSIGNERSLOT is a free data retrieval call binding the contract method 0x4f16540b.
//
// Solidity: function UNSAFE_BLOCK_SIGNER_SLOT() view returns(bytes32)
func (_SystemConfig *SystemConfigCaller) UNSAFEBLOCKSIGNERSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SystemConfig.contract.Call(opts, &out, "UNSAFE_BLOCK_SIGNER_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNSAFEBLOCKSIGNERSLOT is a free data retrieval call binding the contract method 0x4f16540b.
//
// Solidity: function UNSAFE_BLOCK_SIGNER_SLOT() view returns(bytes32)
func (_SystemConfig *SystemConfigSession) UNSAFEBLOCKSIGNERSLOT() ([32]byte, error) {
	return _SystemConfig.Contract.UNSAFEBLOCKSIGNERSLOT(&_SystemConfig.CallOpts)
}

// UNSAFEBLOCKSIGNERSLOT is a free data retrieval call binding the contract method 0x4f16540b.
//
// Solidity: function UNSAFE_BLOCK_SIGNER_SLOT() view returns(bytes32)
func (_SystemConfig *SystemConfigCallerSession) UNSAFEBLOCKSIGNERSLOT() ([32]byte, error) {
	return _SystemConfig.Contract.UNSAFEBLOCKSIGNERSLOT(&_SystemConfig.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_SystemConfig *SystemConfigCaller) VERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemConfig.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_SystemConfig *SystemConfigSession) VERSION() (*big.Int, error) {
	return _SystemConfig.Contract.VERSION(&_SystemConfig.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_SystemConfig *SystemConfigCallerSession) VERSION() (*big.Int, error) {
	return _SystemConfig.Contract.VERSION(&_SystemConfig.CallOpts)
}

// BatchInboxAddr is a free data retrieval call binding the contract method 0xc495de62.
//
// Solidity: function batchInboxAddr() view returns(address)
func (_SystemConfig *SystemConfigCaller) BatchInboxAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemConfig.contract.Call(opts, &out, "batchInboxAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BatchInboxAddr is a free data retrieval call binding the contract method 0xc495de62.
//
// Solidity: function batchInboxAddr() view returns(address)
func (_SystemConfig *SystemConfigSession) BatchInboxAddr() (common.Address, error) {
	return _SystemConfig.Contract.BatchInboxAddr(&_SystemConfig.CallOpts)
}

// BatchInboxAddr is a free data retrieval call binding the contract method 0xc495de62.
//
// Solidity: function batchInboxAddr() view returns(address)
func (_SystemConfig *SystemConfigCallerSession) BatchInboxAddr() (common.Address, error) {
	return _SystemConfig.Contract.BatchInboxAddr(&_SystemConfig.CallOpts)
}

// BatcherHash is a free data retrieval call binding the contract method 0xe81b2c6d.
//
// Solidity: function batcherHash() view returns(bytes32)
func (_SystemConfig *SystemConfigCaller) BatcherHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SystemConfig.contract.Call(opts, &out, "batcherHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatcherHash is a free data retrieval call binding the contract method 0xe81b2c6d.
//
// Solidity: function batcherHash() view returns(bytes32)
func (_SystemConfig *SystemConfigSession) BatcherHash() ([32]byte, error) {
	return _SystemConfig.Contract.BatcherHash(&_SystemConfig.CallOpts)
}

// BatcherHash is a free data retrieval call binding the contract method 0xe81b2c6d.
//
// Solidity: function batcherHash() view returns(bytes32)
func (_SystemConfig *SystemConfigCallerSession) BatcherHash() ([32]byte, error) {
	return _SystemConfig.Contract.BatcherHash(&_SystemConfig.CallOpts)
}

// GasLimit is a free data retrieval call binding the contract method 0xf68016b7.
//
// Solidity: function gasLimit() view returns(uint64)
func (_SystemConfig *SystemConfigCaller) GasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SystemConfig.contract.Call(opts, &out, "gasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GasLimit is a free data retrieval call binding the contract method 0xf68016b7.
//
// Solidity: function gasLimit() view returns(uint64)
func (_SystemConfig *SystemConfigSession) GasLimit() (uint64, error) {
	return _SystemConfig.Contract.GasLimit(&_SystemConfig.CallOpts)
}

// GasLimit is a free data retrieval call binding the contract method 0xf68016b7.
//
// Solidity: function gasLimit() view returns(uint64)
func (_SystemConfig *SystemConfigCallerSession) GasLimit() (uint64, error) {
	return _SystemConfig.Contract.GasLimit(&_SystemConfig.CallOpts)
}

// MinimumGasLimit is a free data retrieval call binding the contract method 0x4add321d.
//
// Solidity: function minimumGasLimit() view returns(uint64)
func (_SystemConfig *SystemConfigCaller) MinimumGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SystemConfig.contract.Call(opts, &out, "minimumGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinimumGasLimit is a free data retrieval call binding the contract method 0x4add321d.
//
// Solidity: function minimumGasLimit() view returns(uint64)
func (_SystemConfig *SystemConfigSession) MinimumGasLimit() (uint64, error) {
	return _SystemConfig.Contract.MinimumGasLimit(&_SystemConfig.CallOpts)
}

// MinimumGasLimit is a free data retrieval call binding the contract method 0x4add321d.
//
// Solidity: function minimumGasLimit() view returns(uint64)
func (_SystemConfig *SystemConfigCallerSession) MinimumGasLimit() (uint64, error) {
	return _SystemConfig.Contract.MinimumGasLimit(&_SystemConfig.CallOpts)
}

// Overhead is a free data retrieval call binding the contract method 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (_SystemConfig *SystemConfigCaller) Overhead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemConfig.contract.Call(opts, &out, "overhead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Overhead is a free data retrieval call binding the contract method 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (_SystemConfig *SystemConfigSession) Overhead() (*big.Int, error) {
	return _SystemConfig.Contract.Overhead(&_SystemConfig.CallOpts)
}

// Overhead is a free data retrieval call binding the contract method 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (_SystemConfig *SystemConfigCallerSession) Overhead() (*big.Int, error) {
	return _SystemConfig.Contract.Overhead(&_SystemConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemConfig *SystemConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemConfig.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemConfig *SystemConfigSession) Owner() (common.Address, error) {
	return _SystemConfig.Contract.Owner(&_SystemConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemConfig *SystemConfigCallerSession) Owner() (common.Address, error) {
	return _SystemConfig.Contract.Owner(&_SystemConfig.CallOpts)
}

// ResourceConfig is a free data retrieval call binding the contract method 0xcc731b02.
//
// Solidity: function resourceConfig() view returns((uint32,uint8,uint8,uint32,uint32,uint128))
func (_SystemConfig *SystemConfigCaller) ResourceConfig(opts *bind.CallOpts) (ResourceMeteringResourceConfig, error) {
	var out []interface{}
	err := _SystemConfig.contract.Call(opts, &out, "resourceConfig")

	if err != nil {
		return *new(ResourceMeteringResourceConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ResourceMeteringResourceConfig)).(*ResourceMeteringResourceConfig)

	return out0, err

}

// ResourceConfig is a free data retrieval call binding the contract method 0xcc731b02.
//
// Solidity: function resourceConfig() view returns((uint32,uint8,uint8,uint32,uint32,uint128))
func (_SystemConfig *SystemConfigSession) ResourceConfig() (ResourceMeteringResourceConfig, error) {
	return _SystemConfig.Contract.ResourceConfig(&_SystemConfig.CallOpts)
}

// ResourceConfig is a free data retrieval call binding the contract method 0xcc731b02.
//
// Solidity: function resourceConfig() view returns((uint32,uint8,uint8,uint32,uint32,uint128))
func (_SystemConfig *SystemConfigCallerSession) ResourceConfig() (ResourceMeteringResourceConfig, error) {
	return _SystemConfig.Contract.ResourceConfig(&_SystemConfig.CallOpts)
}

// Scalar is a free data retrieval call binding the contract method 0xf45e65d8.
//
// Solidity: function scalar() view returns(uint256)
func (_SystemConfig *SystemConfigCaller) Scalar(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemConfig.contract.Call(opts, &out, "scalar")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Scalar is a free data retrieval call binding the contract method 0xf45e65d8.
//
// Solidity: function scalar() view returns(uint256)
func (_SystemConfig *SystemConfigSession) Scalar() (*big.Int, error) {
	return _SystemConfig.Contract.Scalar(&_SystemConfig.CallOpts)
}

// Scalar is a free data retrieval call binding the contract method 0xf45e65d8.
//
// Solidity: function scalar() view returns(uint256)
func (_SystemConfig *SystemConfigCallerSession) Scalar() (*big.Int, error) {
	return _SystemConfig.Contract.Scalar(&_SystemConfig.CallOpts)
}

// UnsafeBlockSigner is a free data retrieval call binding the contract method 0x1fd19ee1.
//
// Solidity: function unsafeBlockSigner() view returns(address)
func (_SystemConfig *SystemConfigCaller) UnsafeBlockSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemConfig.contract.Call(opts, &out, "unsafeBlockSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnsafeBlockSigner is a free data retrieval call binding the contract method 0x1fd19ee1.
//
// Solidity: function unsafeBlockSigner() view returns(address)
func (_SystemConfig *SystemConfigSession) UnsafeBlockSigner() (common.Address, error) {
	return _SystemConfig.Contract.UnsafeBlockSigner(&_SystemConfig.CallOpts)
}

// UnsafeBlockSigner is a free data retrieval call binding the contract method 0x1fd19ee1.
//
// Solidity: function unsafeBlockSigner() view returns(address)
func (_SystemConfig *SystemConfigCallerSession) UnsafeBlockSigner() (common.Address, error) {
	return _SystemConfig.Contract.UnsafeBlockSigner(&_SystemConfig.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SystemConfig *SystemConfigCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SystemConfig.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SystemConfig *SystemConfigSession) Version() (string, error) {
	return _SystemConfig.Contract.Version(&_SystemConfig.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_SystemConfig *SystemConfigCallerSession) Version() (string, error) {
	return _SystemConfig.Contract.Version(&_SystemConfig.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x58619b90.
//
// Solidity: function initialize(address _owner, uint256 _overhead, uint256 _scalar, bytes32 _batcherHash, uint64 _gasLimit, address _unsafeBlockSigner, address _batchInboxAddr, (uint32,uint8,uint8,uint32,uint32,uint128) _config) returns()
func (_SystemConfig *SystemConfigTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _overhead *big.Int, _scalar *big.Int, _batcherHash [32]byte, _gasLimit uint64, _unsafeBlockSigner common.Address, _batchInboxAddr common.Address, _config ResourceMeteringResourceConfig) (*types.Transaction, error) {
	return _SystemConfig.contract.Transact(opts, "initialize", _owner, _overhead, _scalar, _batcherHash, _gasLimit, _unsafeBlockSigner, _batchInboxAddr, _config)
}

// Initialize is a paid mutator transaction binding the contract method 0x58619b90.
//
// Solidity: function initialize(address _owner, uint256 _overhead, uint256 _scalar, bytes32 _batcherHash, uint64 _gasLimit, address _unsafeBlockSigner, address _batchInboxAddr, (uint32,uint8,uint8,uint32,uint32,uint128) _config) returns()
func (_SystemConfig *SystemConfigSession) Initialize(_owner common.Address, _overhead *big.Int, _scalar *big.Int, _batcherHash [32]byte, _gasLimit uint64, _unsafeBlockSigner common.Address, _batchInboxAddr common.Address, _config ResourceMeteringResourceConfig) (*types.Transaction, error) {
	return _SystemConfig.Contract.Initialize(&_SystemConfig.TransactOpts, _owner, _overhead, _scalar, _batcherHash, _gasLimit, _unsafeBlockSigner, _batchInboxAddr, _config)
}

// Initialize is a paid mutator transaction binding the contract method 0x58619b90.
//
// Solidity: function initialize(address _owner, uint256 _overhead, uint256 _scalar, bytes32 _batcherHash, uint64 _gasLimit, address _unsafeBlockSigner, address _batchInboxAddr, (uint32,uint8,uint8,uint32,uint32,uint128) _config) returns()
func (_SystemConfig *SystemConfigTransactorSession) Initialize(_owner common.Address, _overhead *big.Int, _scalar *big.Int, _batcherHash [32]byte, _gasLimit uint64, _unsafeBlockSigner common.Address, _batchInboxAddr common.Address, _config ResourceMeteringResourceConfig) (*types.Transaction, error) {
	return _SystemConfig.Contract.Initialize(&_SystemConfig.TransactOpts, _owner, _overhead, _scalar, _batcherHash, _gasLimit, _unsafeBlockSigner, _batchInboxAddr, _config)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemConfig *SystemConfigTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemConfig.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemConfig *SystemConfigSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemConfig.Contract.RenounceOwnership(&_SystemConfig.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SystemConfig *SystemConfigTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SystemConfig.Contract.RenounceOwnership(&_SystemConfig.TransactOpts)
}

// SetBatchInboxAddr is a paid mutator transaction binding the contract method 0xa0ca2cfb.
//
// Solidity: function setBatchInboxAddr(address _batchInboxAddr) returns()
func (_SystemConfig *SystemConfigTransactor) SetBatchInboxAddr(opts *bind.TransactOpts, _batchInboxAddr common.Address) (*types.Transaction, error) {
	return _SystemConfig.contract.Transact(opts, "setBatchInboxAddr", _batchInboxAddr)
}

// SetBatchInboxAddr is a paid mutator transaction binding the contract method 0xa0ca2cfb.
//
// Solidity: function setBatchInboxAddr(address _batchInboxAddr) returns()
func (_SystemConfig *SystemConfigSession) SetBatchInboxAddr(_batchInboxAddr common.Address) (*types.Transaction, error) {
	return _SystemConfig.Contract.SetBatchInboxAddr(&_SystemConfig.TransactOpts, _batchInboxAddr)
}

// SetBatchInboxAddr is a paid mutator transaction binding the contract method 0xa0ca2cfb.
//
// Solidity: function setBatchInboxAddr(address _batchInboxAddr) returns()
func (_SystemConfig *SystemConfigTransactorSession) SetBatchInboxAddr(_batchInboxAddr common.Address) (*types.Transaction, error) {
	return _SystemConfig.Contract.SetBatchInboxAddr(&_SystemConfig.TransactOpts, _batchInboxAddr)
}

// SetBatcherHash is a paid mutator transaction binding the contract method 0xc9b26f61.
//
// Solidity: function setBatcherHash(bytes32 _batcherHash) returns()
func (_SystemConfig *SystemConfigTransactor) SetBatcherHash(opts *bind.TransactOpts, _batcherHash [32]byte) (*types.Transaction, error) {
	return _SystemConfig.contract.Transact(opts, "setBatcherHash", _batcherHash)
}

// SetBatcherHash is a paid mutator transaction binding the contract method 0xc9b26f61.
//
// Solidity: function setBatcherHash(bytes32 _batcherHash) returns()
func (_SystemConfig *SystemConfigSession) SetBatcherHash(_batcherHash [32]byte) (*types.Transaction, error) {
	return _SystemConfig.Contract.SetBatcherHash(&_SystemConfig.TransactOpts, _batcherHash)
}

// SetBatcherHash is a paid mutator transaction binding the contract method 0xc9b26f61.
//
// Solidity: function setBatcherHash(bytes32 _batcherHash) returns()
func (_SystemConfig *SystemConfigTransactorSession) SetBatcherHash(_batcherHash [32]byte) (*types.Transaction, error) {
	return _SystemConfig.Contract.SetBatcherHash(&_SystemConfig.TransactOpts, _batcherHash)
}

// SetGasConfig is a paid mutator transaction binding the contract method 0x935f029e.
//
// Solidity: function setGasConfig(uint256 _overhead, uint256 _scalar) returns()
func (_SystemConfig *SystemConfigTransactor) SetGasConfig(opts *bind.TransactOpts, _overhead *big.Int, _scalar *big.Int) (*types.Transaction, error) {
	return _SystemConfig.contract.Transact(opts, "setGasConfig", _overhead, _scalar)
}

// SetGasConfig is a paid mutator transaction binding the contract method 0x935f029e.
//
// Solidity: function setGasConfig(uint256 _overhead, uint256 _scalar) returns()
func (_SystemConfig *SystemConfigSession) SetGasConfig(_overhead *big.Int, _scalar *big.Int) (*types.Transaction, error) {
	return _SystemConfig.Contract.SetGasConfig(&_SystemConfig.TransactOpts, _overhead, _scalar)
}

// SetGasConfig is a paid mutator transaction binding the contract method 0x935f029e.
//
// Solidity: function setGasConfig(uint256 _overhead, uint256 _scalar) returns()
func (_SystemConfig *SystemConfigTransactorSession) SetGasConfig(_overhead *big.Int, _scalar *big.Int) (*types.Transaction, error) {
	return _SystemConfig.Contract.SetGasConfig(&_SystemConfig.TransactOpts, _overhead, _scalar)
}

// SetGasLimit is a paid mutator transaction binding the contract method 0xb40a817c.
//
// Solidity: function setGasLimit(uint64 _gasLimit) returns()
func (_SystemConfig *SystemConfigTransactor) SetGasLimit(opts *bind.TransactOpts, _gasLimit uint64) (*types.Transaction, error) {
	return _SystemConfig.contract.Transact(opts, "setGasLimit", _gasLimit)
}

// SetGasLimit is a paid mutator transaction binding the contract method 0xb40a817c.
//
// Solidity: function setGasLimit(uint64 _gasLimit) returns()
func (_SystemConfig *SystemConfigSession) SetGasLimit(_gasLimit uint64) (*types.Transaction, error) {
	return _SystemConfig.Contract.SetGasLimit(&_SystemConfig.TransactOpts, _gasLimit)
}

// SetGasLimit is a paid mutator transaction binding the contract method 0xb40a817c.
//
// Solidity: function setGasLimit(uint64 _gasLimit) returns()
func (_SystemConfig *SystemConfigTransactorSession) SetGasLimit(_gasLimit uint64) (*types.Transaction, error) {
	return _SystemConfig.Contract.SetGasLimit(&_SystemConfig.TransactOpts, _gasLimit)
}

// SetResourceConfig is a paid mutator transaction binding the contract method 0xc71973f6.
//
// Solidity: function setResourceConfig((uint32,uint8,uint8,uint32,uint32,uint128) _config) returns()
func (_SystemConfig *SystemConfigTransactor) SetResourceConfig(opts *bind.TransactOpts, _config ResourceMeteringResourceConfig) (*types.Transaction, error) {
	return _SystemConfig.contract.Transact(opts, "setResourceConfig", _config)
}

// SetResourceConfig is a paid mutator transaction binding the contract method 0xc71973f6.
//
// Solidity: function setResourceConfig((uint32,uint8,uint8,uint32,uint32,uint128) _config) returns()
func (_SystemConfig *SystemConfigSession) SetResourceConfig(_config ResourceMeteringResourceConfig) (*types.Transaction, error) {
	return _SystemConfig.Contract.SetResourceConfig(&_SystemConfig.TransactOpts, _config)
}

// SetResourceConfig is a paid mutator transaction binding the contract method 0xc71973f6.
//
// Solidity: function setResourceConfig((uint32,uint8,uint8,uint32,uint32,uint128) _config) returns()
func (_SystemConfig *SystemConfigTransactorSession) SetResourceConfig(_config ResourceMeteringResourceConfig) (*types.Transaction, error) {
	return _SystemConfig.Contract.SetResourceConfig(&_SystemConfig.TransactOpts, _config)
}

// SetUnsafeBlockSigner is a paid mutator transaction binding the contract method 0x18d13918.
//
// Solidity: function setUnsafeBlockSigner(address _unsafeBlockSigner) returns()
func (_SystemConfig *SystemConfigTransactor) SetUnsafeBlockSigner(opts *bind.TransactOpts, _unsafeBlockSigner common.Address) (*types.Transaction, error) {
	return _SystemConfig.contract.Transact(opts, "setUnsafeBlockSigner", _unsafeBlockSigner)
}

// SetUnsafeBlockSigner is a paid mutator transaction binding the contract method 0x18d13918.
//
// Solidity: function setUnsafeBlockSigner(address _unsafeBlockSigner) returns()
func (_SystemConfig *SystemConfigSession) SetUnsafeBlockSigner(_unsafeBlockSigner common.Address) (*types.Transaction, error) {
	return _SystemConfig.Contract.SetUnsafeBlockSigner(&_SystemConfig.TransactOpts, _unsafeBlockSigner)
}

// SetUnsafeBlockSigner is a paid mutator transaction binding the contract method 0x18d13918.
//
// Solidity: function setUnsafeBlockSigner(address _unsafeBlockSigner) returns()
func (_SystemConfig *SystemConfigTransactorSession) SetUnsafeBlockSigner(_unsafeBlockSigner common.Address) (*types.Transaction, error) {
	return _SystemConfig.Contract.SetUnsafeBlockSigner(&_SystemConfig.TransactOpts, _unsafeBlockSigner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemConfig *SystemConfigTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SystemConfig.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemConfig *SystemConfigSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemConfig.Contract.TransferOwnership(&_SystemConfig.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SystemConfig *SystemConfigTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SystemConfig.Contract.TransferOwnership(&_SystemConfig.TransactOpts, newOwner)
}

// SystemConfigConfigUpdateIterator is returned from FilterConfigUpdate and is used to iterate over the raw logs and unpacked data for ConfigUpdate events raised by the SystemConfig contract.
type SystemConfigConfigUpdateIterator struct {
	Event *SystemConfigConfigUpdate // Event containing the contract specifics and raw log

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
func (it *SystemConfigConfigUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemConfigConfigUpdate)
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
		it.Event = new(SystemConfigConfigUpdate)
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
func (it *SystemConfigConfigUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemConfigConfigUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemConfigConfigUpdate represents a ConfigUpdate event raised by the SystemConfig contract.
type SystemConfigConfigUpdate struct {
	Version    *big.Int
	UpdateType uint8
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfigUpdate is a free log retrieval operation binding the contract event 0x1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be.
//
// Solidity: event ConfigUpdate(uint256 indexed version, uint8 indexed updateType, bytes data)
func (_SystemConfig *SystemConfigFilterer) FilterConfigUpdate(opts *bind.FilterOpts, version []*big.Int, updateType []uint8) (*SystemConfigConfigUpdateIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var updateTypeRule []interface{}
	for _, updateTypeItem := range updateType {
		updateTypeRule = append(updateTypeRule, updateTypeItem)
	}

	logs, sub, err := _SystemConfig.contract.FilterLogs(opts, "ConfigUpdate", versionRule, updateTypeRule)
	if err != nil {
		return nil, err
	}
	return &SystemConfigConfigUpdateIterator{contract: _SystemConfig.contract, event: "ConfigUpdate", logs: logs, sub: sub}, nil
}

// WatchConfigUpdate is a free log subscription operation binding the contract event 0x1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be.
//
// Solidity: event ConfigUpdate(uint256 indexed version, uint8 indexed updateType, bytes data)
func (_SystemConfig *SystemConfigFilterer) WatchConfigUpdate(opts *bind.WatchOpts, sink chan<- *SystemConfigConfigUpdate, version []*big.Int, updateType []uint8) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var updateTypeRule []interface{}
	for _, updateTypeItem := range updateType {
		updateTypeRule = append(updateTypeRule, updateTypeItem)
	}

	logs, sub, err := _SystemConfig.contract.WatchLogs(opts, "ConfigUpdate", versionRule, updateTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemConfigConfigUpdate)
				if err := _SystemConfig.contract.UnpackLog(event, "ConfigUpdate", log); err != nil {
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

// ParseConfigUpdate is a log parse operation binding the contract event 0x1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be.
//
// Solidity: event ConfigUpdate(uint256 indexed version, uint8 indexed updateType, bytes data)
func (_SystemConfig *SystemConfigFilterer) ParseConfigUpdate(log types.Log) (*SystemConfigConfigUpdate, error) {
	event := new(SystemConfigConfigUpdate)
	if err := _SystemConfig.contract.UnpackLog(event, "ConfigUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemConfigInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SystemConfig contract.
type SystemConfigInitializedIterator struct {
	Event *SystemConfigInitialized // Event containing the contract specifics and raw log

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
func (it *SystemConfigInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemConfigInitialized)
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
		it.Event = new(SystemConfigInitialized)
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
func (it *SystemConfigInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemConfigInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemConfigInitialized represents a Initialized event raised by the SystemConfig contract.
type SystemConfigInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemConfig *SystemConfigFilterer) FilterInitialized(opts *bind.FilterOpts) (*SystemConfigInitializedIterator, error) {

	logs, sub, err := _SystemConfig.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SystemConfigInitializedIterator{contract: _SystemConfig.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemConfig *SystemConfigFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SystemConfigInitialized) (event.Subscription, error) {

	logs, sub, err := _SystemConfig.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemConfigInitialized)
				if err := _SystemConfig.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SystemConfig *SystemConfigFilterer) ParseInitialized(log types.Log) (*SystemConfigInitialized, error) {
	event := new(SystemConfigInitialized)
	if err := _SystemConfig.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemConfigOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SystemConfig contract.
type SystemConfigOwnershipTransferredIterator struct {
	Event *SystemConfigOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SystemConfigOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemConfigOwnershipTransferred)
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
		it.Event = new(SystemConfigOwnershipTransferred)
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
func (it *SystemConfigOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemConfigOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemConfigOwnershipTransferred represents a OwnershipTransferred event raised by the SystemConfig contract.
type SystemConfigOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemConfig *SystemConfigFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SystemConfigOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemConfig.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SystemConfigOwnershipTransferredIterator{contract: _SystemConfig.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SystemConfig *SystemConfigFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SystemConfigOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SystemConfig.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemConfigOwnershipTransferred)
				if err := _SystemConfig.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SystemConfig *SystemConfigFilterer) ParseOwnershipTransferred(log types.Log) (*SystemConfigOwnershipTransferred, error) {
	event := new(SystemConfigOwnershipTransferred)
	if err := _SystemConfig.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
