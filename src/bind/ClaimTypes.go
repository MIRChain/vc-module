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

// ClaimTypesABI is the input ABI used to generate the binding from.
const ClaimTypesABI = "[]"

var ClaimTypesParsedABI, _ = abi.JSON[gost3410.PublicKey](strings.NewReader(ClaimTypesABI))

// ClaimTypesBin is the compiled bytecode used for deploying new contracts.
var ClaimTypesBin = "0x6080604052348015600f57600080fd5b50606a80601d6000396000f3fe6080604052600080fdfea26469706673582212201c7569083210b54165821efb7a5bf0227d3bdb1c4627df6aacb6f9feaab581fc64736f6c63782d302e382e31342d646576656c6f702e323032342e31322e31322b636f6d6d69742e62396362363034312e6d6f64005e"

// DeployClaimTypes deploys a new Ethereum contract, binding an instance of ClaimTypes to it.
func DeployClaimTypes[P crypto.PublicKey](auth *bind.TransactOpts[P], backend bind.ContractBackend[P]) (common.Address, *types.Transaction[P], *ClaimTypes[P], error) {
	parsed, err := abi.JSON[P](strings.NewReader(ClaimTypesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClaimTypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClaimTypes[P]{ClaimTypesCaller: ClaimTypesCaller[P]{contract: contract}, ClaimTypesTransactor: ClaimTypesTransactor[P]{contract: contract}, ClaimTypesFilterer: ClaimTypesFilterer[P]{contract: contract}}, nil
}

// ClaimTypes is an auto generated Go binding around an Ethereum contract.
type ClaimTypes[P crypto.PublicKey] struct {
	ClaimTypesCaller[P]     // Read-only binding to the contract
	ClaimTypesTransactor[P] // Write-only binding to the contract
	ClaimTypesFilterer[P]   // Log filterer for contract events
}

// ClaimTypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimTypesCaller[P crypto.PublicKey] struct {
	contract *bind.BoundContract[P] // Generic contract wrapper for the low level calls
}

// ClaimTypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimTypesTransactor[P crypto.PublicKey] struct {
	contract *bind.BoundContract[P] // Generic contract wrapper for the low level calls
}

// ClaimTypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimTypesFilterer[P crypto.PublicKey] struct {
	contract *bind.BoundContract[P] // Generic contract wrapper for the low level calls
}

// ClaimTypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimTypesSession[P crypto.PublicKey] struct {
	Contract     *ClaimTypes[P]       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts[P] // Transaction auth options to use throughout this session
}

// ClaimTypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimTypesCallerSession[P crypto.PublicKey] struct {
	Contract *ClaimTypesCaller[P] // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ClaimTypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimTypesTransactorSession[P crypto.PublicKey] struct {
	Contract     *ClaimTypesTransactor[P] // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts[P]     // Transaction auth options to use throughout this session
}

// ClaimTypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimTypesRaw[P crypto.PublicKey] struct {
	Contract *ClaimTypes[P] // Generic contract binding to access the raw methods on
}

// ClaimTypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimTypesCallerRaw[P crypto.PublicKey] struct {
	Contract *ClaimTypesCaller[P] // Generic read-only contract binding to access the raw methods on
}

// ClaimTypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimTypesTransactorRaw[P crypto.PublicKey] struct {
	Contract *ClaimTypesTransactor[P] // Generic write-only contract binding to access the raw methods on
}

// NewClaimTypes creates a new instance of ClaimTypes, bound to a specific deployed contract.
func NewClaimTypes[P crypto.PublicKey](address common.Address, backend bind.ContractBackend[P]) (*ClaimTypes[P], error) {
	contract, err := bindClaimTypes[P](address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimTypes[P]{ClaimTypesCaller: ClaimTypesCaller[P]{contract: contract}, ClaimTypesTransactor: ClaimTypesTransactor[P]{contract: contract}, ClaimTypesFilterer: ClaimTypesFilterer[P]{contract: contract}}, nil
}

// NewClaimTypesCaller creates a new read-only instance of ClaimTypes, bound to a specific deployed contract.
func NewClaimTypesCaller[P crypto.PublicKey](address common.Address, caller bind.ContractCaller) (*ClaimTypesCaller[P], error) {
	contract, err := bindClaimTypes[P](address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimTypesCaller[P]{contract: contract}, nil
}

// NewClaimTypesTransactor creates a new write-only instance of ClaimTypes, bound to a specific deployed contract.
func NewClaimTypesTransactor[P crypto.PublicKey](address common.Address, transactor bind.ContractTransactor[P]) (*ClaimTypesTransactor[P], error) {
	contract, err := bindClaimTypes[P](address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimTypesTransactor[P]{contract: contract}, nil
}

// NewClaimTypesFilterer creates a new log filterer instance of ClaimTypes, bound to a specific deployed contract.
func NewClaimTypesFilterer[P crypto.PublicKey](address common.Address, filterer bind.ContractFilterer) (*ClaimTypesFilterer[P], error) {
	contract, err := bindClaimTypes[P](address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimTypesFilterer[P]{contract: contract}, nil
}

// bindClaimTypes binds a generic wrapper to an already deployed contract.
func bindClaimTypes[P crypto.PublicKey](address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor[P], filterer bind.ContractFilterer) (*bind.BoundContract[P], error) {
	parsed, err := abi.JSON[P](strings.NewReader(ClaimTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimTypes *ClaimTypesRaw[P]) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimTypes.Contract.ClaimTypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimTypes *ClaimTypesRaw[P]) Transfer(opts *bind.TransactOpts[P]) (*types.Transaction[P], error) {
	return _ClaimTypes.Contract.ClaimTypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimTypes *ClaimTypesRaw[P]) Transact(opts *bind.TransactOpts[P], method string, params ...interface{}) (*types.Transaction[P], error) {
	return _ClaimTypes.Contract.ClaimTypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimTypes *ClaimTypesCallerRaw[P]) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimTypes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimTypes *ClaimTypesTransactorRaw[P]) Transfer(opts *bind.TransactOpts[P]) (*types.Transaction[P], error) {
	return _ClaimTypes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimTypes *ClaimTypesTransactorRaw[P]) Transact(opts *bind.TransactOpts[P], method string, params ...interface{}) (*types.Transaction[P], error) {
	return _ClaimTypes.Contract.contract.Transact(opts, method, params...)
}
