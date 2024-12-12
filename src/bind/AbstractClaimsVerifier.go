// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bind

import (
	"math/big"
	"strings"

	ethereum "github.com/MIRChain/MIR"
	"github.com/MIRChain/MIR/accounts/abi"
	"github.com/MIRChain/MIR/accounts/abi/bind"
	"github.com/MIRChain/MIR/common"
	"github.com/MIRChain/MIR/core/types"
	"github.com/MIRChain/MIR/crypto"
	"github.com/MIRChain/MIR/crypto/gost3410"
	"github.com/MIRChain/MIR/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = common.Big1
	_ = event.NewSubscription
)

// AbstractClaimsVerifierABI is the input ABI used to generate the binding from.
const AbstractClaimsVerifierABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_registryAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

var AbstractClaimsVerifierParsedABI, _ = abi.JSON[gost3410.PublicKey](strings.NewReader(AbstractClaimsVerifierABI))

// AbstractClaimsVerifierBin is the compiled bytecode used for deploying new contracts.
var AbstractClaimsVerifierBin = "0x608060405234801561001057600080fd5b506040516105043803806105048339818101604052810190610032919061034f565b61007760405180608001604052808781526020018681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff168152506100c760201b60201c565b600181905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050505061048c565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f82600001518051906020012083602001518051906020012084604001518560600151604051602001610120959493929190610439565b604051602081830303815290604052805190602001209050919050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f957a808100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6101a48261015b565b810181811067ffffffffffffffff821117156101c3576101c261016c565b5b80604052505050565b60006101d661013d565b90506101e2828261019b565b919050565b600067ffffffffffffffff8211156102025761020161016c565b5b61020b8261015b565b9050602081019050919050565b60005b8381101561023657808201518184015260208101905061021b565b83811115610245576000848401525b50505050565b600061025e610259846101e7565b6101cc565b90508281526020810184848401111561027a57610279610156565b5b610285848285610218565b509392505050565b600082601f8301126102a2576102a1610151565b5b81516102b284826020860161024b565b91505092915050565b6000819050919050565b6102ce816102bb565b81146102d957600080fd5b50565b6000815190506102eb816102c5565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061031c826102f1565b9050919050565b61032c81610311565b811461033757600080fd5b50565b60008151905061034981610323565b92915050565b600080600080600060a0868803121561036b5761036a610147565b5b600086015167ffffffffffffffff8111156103895761038861014c565b5b6103958882890161028d565b955050602086015167ffffffffffffffff8111156103b6576103b561014c565b5b6103c28882890161028d565b94505060406103d3888289016102dc565b93505060606103e48882890161033a565b92505060806103f58882890161033a565b9150509295509295909350565b6000819050919050565b61041581610402565b82525050565b610424816102bb565b82525050565b61043381610311565b82525050565b600060a08201905061044e600083018861040c565b61045b602083018761040c565b610468604083018661040c565b610475606083018561041b565b610482608083018461042a565b9695505050505050565b606a8061049a6000396000f3fe6080604052600080fdfea264697066735822122041952c64f45b376cadc0ae7e22dafdf588461465567dabdf8591c0009803536a64736f6c63782d302e382e31342d646576656c6f702e323032342e31322e31322b636f6d6d69742e62396362363034312e6d6f64005e"

// DeployAbstractClaimsVerifier deploys a new Ethereum contract, binding an instance of AbstractClaimsVerifier to it.
func DeployAbstractClaimsVerifier[P crypto.PublicKey](auth *bind.TransactOpts[P], backend bind.ContractBackend[P], name string, version string, chainId *big.Int, verifyingContract common.Address, _registryAddress common.Address) (common.Address, *types.Transaction[P], *AbstractClaimsVerifier[P], error) {
	parsed, err := abi.JSON[P](strings.NewReader(AbstractClaimsVerifierABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbstractClaimsVerifierBin), backend, name, version, chainId, verifyingContract, _registryAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AbstractClaimsVerifier[P]{AbstractClaimsVerifierCaller: AbstractClaimsVerifierCaller[P]{contract: contract}, AbstractClaimsVerifierTransactor: AbstractClaimsVerifierTransactor[P]{contract: contract}, AbstractClaimsVerifierFilterer: AbstractClaimsVerifierFilterer[P]{contract: contract}}, nil
}

// AbstractClaimsVerifier is an auto generated Go binding around an Ethereum contract.
type AbstractClaimsVerifier[P crypto.PublicKey] struct {
	AbstractClaimsVerifierCaller[P]     // Read-only binding to the contract
	AbstractClaimsVerifierTransactor[P] // Write-only binding to the contract
	AbstractClaimsVerifierFilterer[P]   // Log filterer for contract events
}

// AbstractClaimsVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbstractClaimsVerifierCaller[P crypto.PublicKey] struct {
	contract *bind.BoundContract[P] // Generic contract wrapper for the low level calls
}

// AbstractClaimsVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbstractClaimsVerifierTransactor[P crypto.PublicKey] struct {
	contract *bind.BoundContract[P] // Generic contract wrapper for the low level calls
}

// AbstractClaimsVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbstractClaimsVerifierFilterer[P crypto.PublicKey] struct {
	contract *bind.BoundContract[P] // Generic contract wrapper for the low level calls
}

// AbstractClaimsVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbstractClaimsVerifierSession[P crypto.PublicKey] struct {
	Contract     *AbstractClaimsVerifier[P] // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts[P]       // Transaction auth options to use throughout this session
}

// AbstractClaimsVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbstractClaimsVerifierCallerSession[P crypto.PublicKey] struct {
	Contract *AbstractClaimsVerifierCaller[P] // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// AbstractClaimsVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbstractClaimsVerifierTransactorSession[P crypto.PublicKey] struct {
	Contract     *AbstractClaimsVerifierTransactor[P] // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts[P]                 // Transaction auth options to use throughout this session
}

// AbstractClaimsVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbstractClaimsVerifierRaw[P crypto.PublicKey] struct {
	Contract *AbstractClaimsVerifier[P] // Generic contract binding to access the raw methods on
}

// AbstractClaimsVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbstractClaimsVerifierCallerRaw[P crypto.PublicKey] struct {
	Contract *AbstractClaimsVerifierCaller[P] // Generic read-only contract binding to access the raw methods on
}

// AbstractClaimsVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbstractClaimsVerifierTransactorRaw[P crypto.PublicKey] struct {
	Contract *AbstractClaimsVerifierTransactor[P] // Generic write-only contract binding to access the raw methods on
}

// NewAbstractClaimsVerifier creates a new instance of AbstractClaimsVerifier, bound to a specific deployed contract.
func NewAbstractClaimsVerifier[P crypto.PublicKey](address common.Address, backend bind.ContractBackend[P]) (*AbstractClaimsVerifier[P], error) {
	contract, err := bindAbstractClaimsVerifier[P](address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbstractClaimsVerifier[P]{AbstractClaimsVerifierCaller: AbstractClaimsVerifierCaller[P]{contract: contract}, AbstractClaimsVerifierTransactor: AbstractClaimsVerifierTransactor[P]{contract: contract}, AbstractClaimsVerifierFilterer: AbstractClaimsVerifierFilterer[P]{contract: contract}}, nil
}

// NewAbstractClaimsVerifierCaller creates a new read-only instance of AbstractClaimsVerifier, bound to a specific deployed contract.
func NewAbstractClaimsVerifierCaller[P crypto.PublicKey](address common.Address, caller bind.ContractCaller) (*AbstractClaimsVerifierCaller[P], error) {
	contract, err := bindAbstractClaimsVerifier[P](address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractClaimsVerifierCaller[P]{contract: contract}, nil
}

// NewAbstractClaimsVerifierTransactor creates a new write-only instance of AbstractClaimsVerifier, bound to a specific deployed contract.
func NewAbstractClaimsVerifierTransactor[P crypto.PublicKey](address common.Address, transactor bind.ContractTransactor[P]) (*AbstractClaimsVerifierTransactor[P], error) {
	contract, err := bindAbstractClaimsVerifier[P](address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractClaimsVerifierTransactor[P]{contract: contract}, nil
}

// NewAbstractClaimsVerifierFilterer creates a new log filterer instance of AbstractClaimsVerifier, bound to a specific deployed contract.
func NewAbstractClaimsVerifierFilterer[P crypto.PublicKey](address common.Address, filterer bind.ContractFilterer) (*AbstractClaimsVerifierFilterer[P], error) {
	contract, err := bindAbstractClaimsVerifier[P](address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbstractClaimsVerifierFilterer[P]{contract: contract}, nil
}

// bindAbstractClaimsVerifier binds a generic wrapper to an already deployed contract.
func bindAbstractClaimsVerifier[P crypto.PublicKey](address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor[P], filterer bind.ContractFilterer) (*bind.BoundContract[P], error) {
	parsed, err := abi.JSON[P](strings.NewReader(AbstractClaimsVerifierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractClaimsVerifier *AbstractClaimsVerifierRaw[P]) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractClaimsVerifier.Contract.AbstractClaimsVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractClaimsVerifier *AbstractClaimsVerifierRaw[P]) Transfer(opts *bind.TransactOpts[P]) (*types.Transaction[P], error) {
	return _AbstractClaimsVerifier.Contract.AbstractClaimsVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractClaimsVerifier *AbstractClaimsVerifierRaw[P]) Transact(opts *bind.TransactOpts[P], method string, params ...interface{}) (*types.Transaction[P], error) {
	return _AbstractClaimsVerifier.Contract.AbstractClaimsVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractClaimsVerifier *AbstractClaimsVerifierCallerRaw[P]) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractClaimsVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractClaimsVerifier *AbstractClaimsVerifierTransactorRaw[P]) Transfer(opts *bind.TransactOpts[P]) (*types.Transaction[P], error) {
	return _AbstractClaimsVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractClaimsVerifier *AbstractClaimsVerifierTransactorRaw[P]) Transact(opts *bind.TransactOpts[P], method string, params ...interface{}) (*types.Transaction[P], error) {
	return _AbstractClaimsVerifier.Contract.contract.Transact(opts, method, params...)
}
