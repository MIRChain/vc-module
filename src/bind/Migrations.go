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
var MigrationsBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610317806100606000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806323ccd125146100515780637800dc651461006f578063805b60611461008d578063d8d9f218146100a9575b600080fd5b6100596100c5565b60405161006691906101b0565b60405180910390f35b6100776100e9565b60405161008491906101e4565b60405180910390f35b6100a760048036038101906100a29190610230565b6100ef565b005b6100c360048036038101906100be9190610289565b610165565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b60008190508073ffffffffffffffffffffffffffffffffffffffff1663d8d9f2186001546040518263ffffffff1660e01b815260040161012f91906101e4565b600060405180830381600087803b15801561014957600080fd5b505af115801561015d573d6000803e3d6000fd5b505050505050565b8060018190555050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061019a8261016f565b9050919050565b6101aa8161018f565b82525050565b60006020820190506101c560008301846101a1565b92915050565b6000819050919050565b6101de816101cb565b82525050565b60006020820190506101f960008301846101d5565b92915050565b600080fd5b61020d8161018f565b811461021857600080fd5b50565b60008135905061022a81610204565b92915050565b600060208284031215610246576102456101ff565b5b60006102548482850161021b565b91505092915050565b610266816101cb565b811461027157600080fd5b50565b6000813590506102838161025d565b92915050565b60006020828403121561029f5761029e6101ff565b5b60006102ad84828501610274565b9150509291505056fea26469706673582212204911660f17b639b40e38920a9643ddcc09a9094a0ec51d2306fd5fa65c38a9c564736f6c63782d302e382e31342d646576656c6f702e323032342e31322e31322b636f6d6d69742e62396362363034312e6d6f64005e"

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

// LastCompletedMigration is a free data retrieval call binding the contract method 0x7800dc65.
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

// LastCompletedMigration is a free data retrieval call binding the contract method 0x7800dc65.
//
// Solidity: function last_completed_migration() view returns(uint256)
func (_Migrations *MigrationsSession[P]) LastCompletedMigration() (*big.Int, error) {
	return _Migrations.Contract.LastCompletedMigration(&_Migrations.CallOpts)
}

// LastCompletedMigration is a free data retrieval call binding the contract method 0x7800dc65.
//
// Solidity: function last_completed_migration() view returns(uint256)
func (_Migrations *MigrationsCallerSession[P]) LastCompletedMigration() (*big.Int, error) {
	return _Migrations.Contract.LastCompletedMigration(&_Migrations.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x23ccd125.
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

// Owner is a free data retrieval call binding the contract method 0x23ccd125.
//
// Solidity: function owner() view returns(address)
func (_Migrations *MigrationsSession[P]) Owner() (common.Address, error) {
	return _Migrations.Contract.Owner(&_Migrations.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x23ccd125.
//
// Solidity: function owner() view returns(address)
func (_Migrations *MigrationsCallerSession[P]) Owner() (common.Address, error) {
	return _Migrations.Contract.Owner(&_Migrations.CallOpts)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xd8d9f218.
//
// Solidity: function setCompleted(uint256 completed) returns()
func (_Migrations *MigrationsTransactor[P]) SetCompleted(opts *bind.TransactOpts[P], completed *big.Int) (*types.Transaction[P], error) {
	return _Migrations.contract.Transact(opts, "setCompleted", completed)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xd8d9f218.
//
// Solidity: function setCompleted(uint256 completed) returns()
func (_Migrations *MigrationsSession[P]) SetCompleted(completed *big.Int) (*types.Transaction[P], error) {
	return _Migrations.Contract.SetCompleted(&_Migrations.TransactOpts, completed)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xd8d9f218.
//
// Solidity: function setCompleted(uint256 completed) returns()
func (_Migrations *MigrationsTransactorSession[P]) SetCompleted(completed *big.Int) (*types.Transaction[P], error) {
	return _Migrations.Contract.SetCompleted(&_Migrations.TransactOpts, completed)
}

// Upgrade is a paid mutator transaction binding the contract method 0x805b6061.
//
// Solidity: function upgrade(address new_address) returns()
func (_Migrations *MigrationsTransactor[P]) Upgrade(opts *bind.TransactOpts[P], new_address common.Address) (*types.Transaction[P], error) {
	return _Migrations.contract.Transact(opts, "upgrade", new_address)
}

// Upgrade is a paid mutator transaction binding the contract method 0x805b6061.
//
// Solidity: function upgrade(address new_address) returns()
func (_Migrations *MigrationsSession[P]) Upgrade(new_address common.Address) (*types.Transaction[P], error) {
	return _Migrations.Contract.Upgrade(&_Migrations.TransactOpts, new_address)
}

// Upgrade is a paid mutator transaction binding the contract method 0x805b6061.
//
// Solidity: function upgrade(address new_address) returns()
func (_Migrations *MigrationsTransactorSession[P]) Upgrade(new_address common.Address) (*types.Transaction[P], error) {
	return _Migrations.Contract.Upgrade(&_Migrations.TransactOpts, new_address)
}
