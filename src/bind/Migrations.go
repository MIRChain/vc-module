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

// MigrationsABI is the input ABI used to generate the binding from.
const MigrationsABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"last_completed_migration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"completed\",\"type\":\"uint256\"}],\"name\":\"setCompleted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"new_address\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var MigrationsParsedABI, _ = abi.JSON[gost3410.PublicKey](strings.NewReader(MigrationsABI))

// MigrationsBin is the compiled bytecode used for deploying new contracts.
var MigrationsBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506102ec806100606000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630900f01014610051578063445df0ac1461006d5780638da5cb5b1461008b578063fdacd576146100a9575b600080fd5b61006b600480360381019061006691906101d2565b6100c5565b005b61007561013b565b6040516100829190610218565b60405180910390f35b610093610141565b6040516100a09190610242565b60405180910390f35b6100c360048036038101906100be9190610289565b610165565b005b60008190508073ffffffffffffffffffffffffffffffffffffffff1663fdacd5766001546040518263ffffffff1660e01b81526004016101059190610218565b600060405180830381600087803b15801561011f57600080fd5b505af1158015610133573d6000803e3d6000fd5b505050505050565b60015481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b8060018190555050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061019f82610174565b9050919050565b6101af81610194565b81146101ba57600080fd5b50565b6000813590506101cc816101a6565b92915050565b6000602082840312156101e8576101e761016f565b5b60006101f6848285016101bd565b91505092915050565b6000819050919050565b610212816101ff565b82525050565b600060208201905061022d6000830184610209565b92915050565b61023c81610194565b82525050565b60006020820190506102576000830184610233565b92915050565b610266816101ff565b811461027157600080fd5b50565b6000813590506102838161025d565b92915050565b60006020828403121561029f5761029e61016f565b5b60006102ad84828501610274565b9150509291505056fea264697066735822122069beafb288705ae0ed237cfc1538765ad62264f2d0f3684e08c70504aa48c64464736f6c63430008120033"

// DeployMigrations deploys a new Ethereum contract, binding an instance of Migrations to it.
func DeployMigrations[P crypto.PublicKey](auth *bind.TransactOpts[P], backend bind.ContractBackend[P]) (common.Address, *types.Transaction[P], *Migrations[P], error) {
	parsed, err := abi.JSON[P](strings.NewReader(MigrationsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MigrationsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Migrations[P]{MigrationsCaller: MigrationsCaller[P]{contract: contract}, MigrationsTransactor: MigrationsTransactor[P]{contract: contract}, MigrationsFilterer: MigrationsFilterer[P]{contract: contract}}, nil
}

// Migrations is an auto generated Go binding around an Ethereum contract.
type Migrations[P crypto.PublicKey] struct {
	MigrationsCaller[P]     // Read-only binding to the contract
	MigrationsTransactor[P] // Write-only binding to the contract
	MigrationsFilterer[P]   // Log filterer for contract events
}

// MigrationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MigrationsCaller[P crypto.PublicKey] struct {
	contract *bind.BoundContract[P] // Generic contract wrapper for the low level calls
}

// MigrationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MigrationsTransactor[P crypto.PublicKey] struct {
	contract *bind.BoundContract[P] // Generic contract wrapper for the low level calls
}

// MigrationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MigrationsFilterer[P crypto.PublicKey] struct {
	contract *bind.BoundContract[P] // Generic contract wrapper for the low level calls
}

// MigrationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MigrationsSession[P crypto.PublicKey] struct {
	Contract     *Migrations[P]       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts[P] // Transaction auth options to use throughout this session
}

// MigrationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MigrationsCallerSession[P crypto.PublicKey] struct {
	Contract *MigrationsCaller[P] // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MigrationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MigrationsTransactorSession[P crypto.PublicKey] struct {
	Contract     *MigrationsTransactor[P] // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts[P]     // Transaction auth options to use throughout this session
}

// MigrationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MigrationsRaw[P crypto.PublicKey] struct {
	Contract *Migrations[P] // Generic contract binding to access the raw methods on
}

// MigrationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MigrationsCallerRaw[P crypto.PublicKey] struct {
	Contract *MigrationsCaller[P] // Generic read-only contract binding to access the raw methods on
}

// MigrationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MigrationsTransactorRaw[P crypto.PublicKey] struct {
	Contract *MigrationsTransactor[P] // Generic write-only contract binding to access the raw methods on
}

// NewMigrations creates a new instance of Migrations, bound to a specific deployed contract.
func NewMigrations[P crypto.PublicKey](address common.Address, backend bind.ContractBackend[P]) (*Migrations[P], error) {
	contract, err := bindMigrations[P](address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Migrations[P]{MigrationsCaller: MigrationsCaller[P]{contract: contract}, MigrationsTransactor: MigrationsTransactor[P]{contract: contract}, MigrationsFilterer: MigrationsFilterer[P]{contract: contract}}, nil
}

// NewMigrationsCaller creates a new read-only instance of Migrations, bound to a specific deployed contract.
func NewMigrationsCaller[P crypto.PublicKey](address common.Address, caller bind.ContractCaller) (*MigrationsCaller[P], error) {
	contract, err := bindMigrations[P](address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MigrationsCaller[P]{contract: contract}, nil
}

// NewMigrationsTransactor creates a new write-only instance of Migrations, bound to a specific deployed contract.
func NewMigrationsTransactor[P crypto.PublicKey](address common.Address, transactor bind.ContractTransactor[P]) (*MigrationsTransactor[P], error) {
	contract, err := bindMigrations[P](address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MigrationsTransactor[P]{contract: contract}, nil
}

// NewMigrationsFilterer creates a new log filterer instance of Migrations, bound to a specific deployed contract.
func NewMigrationsFilterer[P crypto.PublicKey](address common.Address, filterer bind.ContractFilterer) (*MigrationsFilterer[P], error) {
	contract, err := bindMigrations[P](address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MigrationsFilterer[P]{contract: contract}, nil
}

// bindMigrations binds a generic wrapper to an already deployed contract.
func bindMigrations[P crypto.PublicKey](address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor[P], filterer bind.ContractFilterer) (*bind.BoundContract[P], error) {
	parsed, err := abi.JSON[P](strings.NewReader(MigrationsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Migrations *MigrationsRaw[P]) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Migrations.Contract.MigrationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Migrations *MigrationsRaw[P]) Transfer(opts *bind.TransactOpts[P]) (*types.Transaction[P], error) {
	return _Migrations.Contract.MigrationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Migrations *MigrationsRaw[P]) Transact(opts *bind.TransactOpts[P], method string, params ...interface{}) (*types.Transaction[P], error) {
	return _Migrations.Contract.MigrationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Migrations *MigrationsCallerRaw[P]) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Migrations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Migrations *MigrationsTransactorRaw[P]) Transfer(opts *bind.TransactOpts[P]) (*types.Transaction[P], error) {
	return _Migrations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Migrations *MigrationsTransactorRaw[P]) Transact(opts *bind.TransactOpts[P], method string, params ...interface{}) (*types.Transaction[P], error) {
	return _Migrations.Contract.contract.Transact(opts, method, params...)
}

// LastCompletedMigration is a free data retrieval call binding the contract method 0x445df0ac.
//
// Solidity: function last_completed_migration() view returns(uint256)
func (_Migrations *MigrationsCaller[P]) LastCompletedMigration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Migrations.contract.Call(opts, &out, "last_completed_migration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastCompletedMigration is a free data retrieval call binding the contract method 0x445df0ac.
//
// Solidity: function last_completed_migration() view returns(uint256)
func (_Migrations *MigrationsSession[P]) LastCompletedMigration() (*big.Int, error) {
	return _Migrations.Contract.LastCompletedMigration(&_Migrations.CallOpts)
}

// LastCompletedMigration is a free data retrieval call binding the contract method 0x445df0ac.
//
// Solidity: function last_completed_migration() view returns(uint256)
func (_Migrations *MigrationsCallerSession[P]) LastCompletedMigration() (*big.Int, error) {
	return _Migrations.Contract.LastCompletedMigration(&_Migrations.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Migrations *MigrationsCaller[P]) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Migrations.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Migrations *MigrationsSession[P]) Owner() (common.Address, error) {
	return _Migrations.Contract.Owner(&_Migrations.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Migrations *MigrationsCallerSession[P]) Owner() (common.Address, error) {
	return _Migrations.Contract.Owner(&_Migrations.CallOpts)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
//
// Solidity: function setCompleted(uint256 completed) returns()
func (_Migrations *MigrationsTransactor[P]) SetCompleted(opts *bind.TransactOpts[P], completed *big.Int) (*types.Transaction[P], error) {
	return _Migrations.contract.Transact(opts, "setCompleted", completed)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
//
// Solidity: function setCompleted(uint256 completed) returns()
func (_Migrations *MigrationsSession[P]) SetCompleted(completed *big.Int) (*types.Transaction[P], error) {
	return _Migrations.Contract.SetCompleted(&_Migrations.TransactOpts, completed)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
//
// Solidity: function setCompleted(uint256 completed) returns()
func (_Migrations *MigrationsTransactorSession[P]) SetCompleted(completed *big.Int) (*types.Transaction[P], error) {
	return _Migrations.Contract.SetCompleted(&_Migrations.TransactOpts, completed)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address new_address) returns()
func (_Migrations *MigrationsTransactor[P]) Upgrade(opts *bind.TransactOpts[P], new_address common.Address) (*types.Transaction[P], error) {
	return _Migrations.contract.Transact(opts, "upgrade", new_address)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address new_address) returns()
func (_Migrations *MigrationsSession[P]) Upgrade(new_address common.Address) (*types.Transaction[P], error) {
	return _Migrations.Contract.Upgrade(&_Migrations.TransactOpts, new_address)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(address new_address) returns()
func (_Migrations *MigrationsTransactorSession[P]) Upgrade(new_address common.Address) (*types.Transaction[P], error) {
	return _Migrations.Contract.Upgrade(&_Migrations.TransactOpts, new_address)
}
