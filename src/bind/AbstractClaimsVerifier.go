// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bind

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AbstractClaimsVerifierABI is the input ABI used to generate the binding from.
const AbstractClaimsVerifierABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_registryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

var AbstractClaimsVerifierParsedABI, _ = abi.JSON(strings.NewReader(AbstractClaimsVerifierABI))

// AbstractClaimsVerifierBin is the compiled bytecode used for deploying new contracts.
var AbstractClaimsVerifierBin = "0x608060405234801561001057600080fd5b506040516104d93803806104d98339818101604052810190610032919061034f565b61007760405180608001604052808781526020018681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff168152506100c760201b60201c565b600181905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050505061048c565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f82600001518051906020012083602001518051906020012084604001518560600151604051602001610120959493929190610439565b604051602081830303815290604052805190602001209050919050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6101a48261015b565b810181811067ffffffffffffffff821117156101c3576101c261016c565b5b80604052505050565b60006101d661013d565b90506101e2828261019b565b919050565b600067ffffffffffffffff8211156102025761020161016c565b5b61020b8261015b565b9050602081019050919050565b60005b8381101561023657808201518184015260208101905061021b565b83811115610245576000848401525b50505050565b600061025e610259846101e7565b6101cc565b90508281526020810184848401111561027a57610279610156565b5b610285848285610218565b509392505050565b600082601f8301126102a2576102a1610151565b5b81516102b284826020860161024b565b91505092915050565b6000819050919050565b6102ce816102bb565b81146102d957600080fd5b50565b6000815190506102eb816102c5565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061031c826102f1565b9050919050565b61032c81610311565b811461033757600080fd5b50565b60008151905061034981610323565b92915050565b600080600080600060a0868803121561036b5761036a610147565b5b600086015167ffffffffffffffff8111156103895761038861014c565b5b6103958882890161028d565b955050602086015167ffffffffffffffff8111156103b6576103b561014c565b5b6103c28882890161028d565b94505060406103d3888289016102dc565b93505060606103e48882890161033a565b92505060806103f58882890161033a565b9150509295509295909350565b6000819050919050565b61041581610402565b82525050565b610424816102bb565b82525050565b61043381610311565b82525050565b600060a08201905061044e600083018861040c565b61045b602083018761040c565b610468604083018661040c565b610475606083018561041b565b610482608083018461042a565b9695505050505050565b603f8061049a6000396000f3fe6080604052600080fdfea264697066735822122066a4c44b456811b1f4ce971ecab3d993caf3a9bf2dbfbf36450f1ec9ae68a4d564736f6c634300080e0033"

// DeployAbstractClaimsVerifier deploys a new Ethereum contract, binding an instance of AbstractClaimsVerifier to it.
func DeployAbstractClaimsVerifier(auth *bind.TransactOpts, backend bind.ContractBackend, name string, version string, chainId *big.Int, verifyingContract common.Address, _registryAddress common.Address) (common.Address, *types.Transaction, *AbstractClaimsVerifier, error) {
	parsed, err := abi.JSON(strings.NewReader(AbstractClaimsVerifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbstractClaimsVerifierBin), backend, name, version, chainId, verifyingContract, _registryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AbstractClaimsVerifier{AbstractClaimsVerifierCaller: AbstractClaimsVerifierCaller{contract: contract}, AbstractClaimsVerifierTransactor: AbstractClaimsVerifierTransactor{contract: contract}, AbstractClaimsVerifierFilterer: AbstractClaimsVerifierFilterer{contract: contract}}, nil
}

// AbstractClaimsVerifier is an auto generated Go binding around an Ethereum contract.
type AbstractClaimsVerifier struct {
	AbstractClaimsVerifierCaller     // Read-only binding to the contract
	AbstractClaimsVerifierTransactor // Write-only binding to the contract
	AbstractClaimsVerifierFilterer   // Log filterer for contract events
}

// AbstractClaimsVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbstractClaimsVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractClaimsVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbstractClaimsVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractClaimsVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbstractClaimsVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractClaimsVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbstractClaimsVerifierSession struct {
	Contract     *AbstractClaimsVerifier // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AbstractClaimsVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbstractClaimsVerifierCallerSession struct {
	Contract *AbstractClaimsVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// AbstractClaimsVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbstractClaimsVerifierTransactorSession struct {
	Contract     *AbstractClaimsVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// AbstractClaimsVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbstractClaimsVerifierRaw struct {
	Contract *AbstractClaimsVerifier // Generic contract binding to access the raw methods on
}

// AbstractClaimsVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbstractClaimsVerifierCallerRaw struct {
	Contract *AbstractClaimsVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// AbstractClaimsVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbstractClaimsVerifierTransactorRaw struct {
	Contract *AbstractClaimsVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbstractClaimsVerifier creates a new instance of AbstractClaimsVerifier, bound to a specific deployed contract.
func NewAbstractClaimsVerifier(address common.Address, backend bind.ContractBackend) (*AbstractClaimsVerifier, error) {
	contract, err := bindAbstractClaimsVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbstractClaimsVerifier{AbstractClaimsVerifierCaller: AbstractClaimsVerifierCaller{contract: contract}, AbstractClaimsVerifierTransactor: AbstractClaimsVerifierTransactor{contract: contract}, AbstractClaimsVerifierFilterer: AbstractClaimsVerifierFilterer{contract: contract}}, nil
}

// NewAbstractClaimsVerifierCaller creates a new read-only instance of AbstractClaimsVerifier, bound to a specific deployed contract.
func NewAbstractClaimsVerifierCaller(address common.Address, caller bind.ContractCaller) (*AbstractClaimsVerifierCaller, error) {
	contract, err := bindAbstractClaimsVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractClaimsVerifierCaller{contract: contract}, nil
}

// NewAbstractClaimsVerifierTransactor creates a new write-only instance of AbstractClaimsVerifier, bound to a specific deployed contract.
func NewAbstractClaimsVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*AbstractClaimsVerifierTransactor, error) {
	contract, err := bindAbstractClaimsVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractClaimsVerifierTransactor{contract: contract}, nil
}

// NewAbstractClaimsVerifierFilterer creates a new log filterer instance of AbstractClaimsVerifier, bound to a specific deployed contract.
func NewAbstractClaimsVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*AbstractClaimsVerifierFilterer, error) {
	contract, err := bindAbstractClaimsVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbstractClaimsVerifierFilterer{contract: contract}, nil
}

// bindAbstractClaimsVerifier binds a generic wrapper to an already deployed contract.
func bindAbstractClaimsVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbstractClaimsVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractClaimsVerifier *AbstractClaimsVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractClaimsVerifier.Contract.AbstractClaimsVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractClaimsVerifier *AbstractClaimsVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractClaimsVerifier.Contract.AbstractClaimsVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractClaimsVerifier *AbstractClaimsVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractClaimsVerifier.Contract.AbstractClaimsVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractClaimsVerifier *AbstractClaimsVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractClaimsVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractClaimsVerifier *AbstractClaimsVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractClaimsVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractClaimsVerifier *AbstractClaimsVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractClaimsVerifier.Contract.contract.Transact(opts, method, params...)
}
