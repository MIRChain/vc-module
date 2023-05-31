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

// ClaimTypesABI is the input ABI used to generate the binding from.
const ClaimTypesABI = "[]"

var ClaimTypesParsedABI, _ = abi.JSON(strings.NewReader(ClaimTypesABI))

// ClaimTypesBin is the compiled bytecode used for deploying new contracts.
var ClaimTypesBin = "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212201ed2aa6ca8aef9a5f4beb09ab6ecf5ebc2336ff850d03d627837ad5d7e58e8bb64736f6c634300080e0033"

// DeployClaimTypes deploys a new Ethereum contract, binding an instance of ClaimTypes to it.
func DeployClaimTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ClaimTypes, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimTypesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClaimTypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClaimTypes{ClaimTypesCaller: ClaimTypesCaller{contract: contract}, ClaimTypesTransactor: ClaimTypesTransactor{contract: contract}, ClaimTypesFilterer: ClaimTypesFilterer{contract: contract}}, nil
}

// ClaimTypes is an auto generated Go binding around an Ethereum contract.
type ClaimTypes struct {
	ClaimTypesCaller     // Read-only binding to the contract
	ClaimTypesTransactor // Write-only binding to the contract
	ClaimTypesFilterer   // Log filterer for contract events
}

// ClaimTypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimTypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimTypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimTypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimTypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimTypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimTypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimTypesSession struct {
	Contract     *ClaimTypes       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimTypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimTypesCallerSession struct {
	Contract *ClaimTypesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ClaimTypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimTypesTransactorSession struct {
	Contract     *ClaimTypesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ClaimTypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimTypesRaw struct {
	Contract *ClaimTypes // Generic contract binding to access the raw methods on
}

// ClaimTypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimTypesCallerRaw struct {
	Contract *ClaimTypesCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimTypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimTypesTransactorRaw struct {
	Contract *ClaimTypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimTypes creates a new instance of ClaimTypes, bound to a specific deployed contract.
func NewClaimTypes(address common.Address, backend bind.ContractBackend) (*ClaimTypes, error) {
	contract, err := bindClaimTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimTypes{ClaimTypesCaller: ClaimTypesCaller{contract: contract}, ClaimTypesTransactor: ClaimTypesTransactor{contract: contract}, ClaimTypesFilterer: ClaimTypesFilterer{contract: contract}}, nil
}

// NewClaimTypesCaller creates a new read-only instance of ClaimTypes, bound to a specific deployed contract.
func NewClaimTypesCaller(address common.Address, caller bind.ContractCaller) (*ClaimTypesCaller, error) {
	contract, err := bindClaimTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimTypesCaller{contract: contract}, nil
}

// NewClaimTypesTransactor creates a new write-only instance of ClaimTypes, bound to a specific deployed contract.
func NewClaimTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimTypesTransactor, error) {
	contract, err := bindClaimTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimTypesTransactor{contract: contract}, nil
}

// NewClaimTypesFilterer creates a new log filterer instance of ClaimTypes, bound to a specific deployed contract.
func NewClaimTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimTypesFilterer, error) {
	contract, err := bindClaimTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimTypesFilterer{contract: contract}, nil
}

// bindClaimTypes binds a generic wrapper to an already deployed contract.
func bindClaimTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimTypes *ClaimTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimTypes.Contract.ClaimTypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimTypes *ClaimTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimTypes.Contract.ClaimTypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimTypes *ClaimTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimTypes.Contract.ClaimTypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimTypes *ClaimTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimTypes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimTypes *ClaimTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimTypes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimTypes *ClaimTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimTypes.Contract.contract.Transact(opts, method, params...)
}
