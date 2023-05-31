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

// PermImplABI is the input ABI used to generate the binding from.
const PermImplABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"id\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"iat\",\"type\":\"uint256\"}],\"name\":\"CredentialRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"}],\"name\":\"CredentialRevoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"exist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"registerCredential\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"credentialHash\",\"type\":\"bytes32\"}],\"name\":\"revokeCredential\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_credentialHash\",\"type\":\"bytes32\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

var PermImplParsedABI, _ = abi.JSON(strings.NewReader(PermImplABI))

// PermImpl is an auto generated Go binding around an Ethereum contract.
type PermImpl struct {
	PermImplCaller     // Read-only binding to the contract
	PermImplTransactor // Write-only binding to the contract
	PermImplFilterer   // Log filterer for contract events
}

// PermImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type PermImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PermImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PermImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PermImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PermImplSession struct {
	Contract     *PermImpl         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PermImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PermImplCallerSession struct {
	Contract *PermImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PermImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PermImplTransactorSession struct {
	Contract     *PermImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PermImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type PermImplRaw struct {
	Contract *PermImpl // Generic contract binding to access the raw methods on
}

// PermImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PermImplCallerRaw struct {
	Contract *PermImplCaller // Generic read-only contract binding to access the raw methods on
}

// PermImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PermImplTransactorRaw struct {
	Contract *PermImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPermImpl creates a new instance of PermImpl, bound to a specific deployed contract.
func NewPermImpl(address common.Address, backend bind.ContractBackend) (*PermImpl, error) {
	contract, err := bindPermImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PermImpl{PermImplCaller: PermImplCaller{contract: contract}, PermImplTransactor: PermImplTransactor{contract: contract}, PermImplFilterer: PermImplFilterer{contract: contract}}, nil
}

// NewPermImplCaller creates a new read-only instance of PermImpl, bound to a specific deployed contract.
func NewPermImplCaller(address common.Address, caller bind.ContractCaller) (*PermImplCaller, error) {
	contract, err := bindPermImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PermImplCaller{contract: contract}, nil
}

// NewPermImplTransactor creates a new write-only instance of PermImpl, bound to a specific deployed contract.
func NewPermImplTransactor(address common.Address, transactor bind.ContractTransactor) (*PermImplTransactor, error) {
	contract, err := bindPermImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PermImplTransactor{contract: contract}, nil
}

// NewPermImplFilterer creates a new log filterer instance of PermImpl, bound to a specific deployed contract.
func NewPermImplFilterer(address common.Address, filterer bind.ContractFilterer) (*PermImplFilterer, error) {
	contract, err := bindPermImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PermImplFilterer{contract: contract}, nil
}

// bindPermImpl binds a generic wrapper to an already deployed contract.
func bindPermImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PermImplABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermImpl *PermImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermImpl.Contract.PermImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermImpl *PermImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermImpl.Contract.PermImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermImpl *PermImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermImpl.Contract.PermImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PermImpl *PermImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PermImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PermImpl *PermImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PermImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PermImpl *PermImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PermImpl.Contract.contract.Transact(opts, method, params...)
}

// Exist is a free data retrieval call binding the contract method 0x957839d9.
//
// Solidity: function exist(bytes32 credentialHash, address issuer) view returns(bool)
func (_PermImpl *PermImplCaller) Exist(opts *bind.CallOpts, credentialHash [32]byte, issuer common.Address) (bool, error) {
	var out []interface{}
	err := _PermImpl.contract.Call(opts, &out, "exist", credentialHash, issuer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exist is a free data retrieval call binding the contract method 0x957839d9.
//
// Solidity: function exist(bytes32 credentialHash, address issuer) view returns(bool)
func (_PermImpl *PermImplSession) Exist(credentialHash [32]byte, issuer common.Address) (bool, error) {
	return _PermImpl.Contract.Exist(&_PermImpl.CallOpts, credentialHash, issuer)
}

// Exist is a free data retrieval call binding the contract method 0x957839d9.
//
// Solidity: function exist(bytes32 credentialHash, address issuer) view returns(bool)
func (_PermImpl *PermImplCallerSession) Exist(credentialHash [32]byte, issuer common.Address) (bool, error) {
	return _PermImpl.Contract.Exist(&_PermImpl.CallOpts, credentialHash, issuer)
}

// Status is a free data retrieval call binding the contract method 0xd6d76ed5.
//
// Solidity: function status(address issuer, bytes32 _credentialHash) view returns(bool)
func (_PermImpl *PermImplCaller) Status(opts *bind.CallOpts, issuer common.Address, _credentialHash [32]byte) (bool, error) {
	var out []interface{}
	err := _PermImpl.contract.Call(opts, &out, "status", issuer, _credentialHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Status is a free data retrieval call binding the contract method 0xd6d76ed5.
//
// Solidity: function status(address issuer, bytes32 _credentialHash) view returns(bool)
func (_PermImpl *PermImplSession) Status(issuer common.Address, _credentialHash [32]byte) (bool, error) {
	return _PermImpl.Contract.Status(&_PermImpl.CallOpts, issuer, _credentialHash)
}

// Status is a free data retrieval call binding the contract method 0xd6d76ed5.
//
// Solidity: function status(address issuer, bytes32 _credentialHash) view returns(bool)
func (_PermImpl *PermImplCallerSession) Status(issuer common.Address, _credentialHash [32]byte) (bool, error) {
	return _PermImpl.Contract.Status(&_PermImpl.CallOpts, issuer, _credentialHash)
}

// RegisterCredential is a paid mutator transaction binding the contract method 0x75130747.
//
// Solidity: function registerCredential(address issuer, address subject, bytes32 credentialHash, uint256 from, uint256 exp, bytes signature) returns(bool)
func (_PermImpl *PermImplTransactor) RegisterCredential(opts *bind.TransactOpts, issuer common.Address, subject common.Address, credentialHash [32]byte, from *big.Int, exp *big.Int, signature []byte) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "registerCredential", issuer, subject, credentialHash, from, exp, signature)
}

// RegisterCredential is a paid mutator transaction binding the contract method 0x75130747.
//
// Solidity: function registerCredential(address issuer, address subject, bytes32 credentialHash, uint256 from, uint256 exp, bytes signature) returns(bool)
func (_PermImpl *PermImplSession) RegisterCredential(issuer common.Address, subject common.Address, credentialHash [32]byte, from *big.Int, exp *big.Int, signature []byte) (*types.Transaction, error) {
	return _PermImpl.Contract.RegisterCredential(&_PermImpl.TransactOpts, issuer, subject, credentialHash, from, exp, signature)
}

// RegisterCredential is a paid mutator transaction binding the contract method 0x75130747.
//
// Solidity: function registerCredential(address issuer, address subject, bytes32 credentialHash, uint256 from, uint256 exp, bytes signature) returns(bool)
func (_PermImpl *PermImplTransactorSession) RegisterCredential(issuer common.Address, subject common.Address, credentialHash [32]byte, from *big.Int, exp *big.Int, signature []byte) (*types.Transaction, error) {
	return _PermImpl.Contract.RegisterCredential(&_PermImpl.TransactOpts, issuer, subject, credentialHash, from, exp, signature)
}

// RevokeCredential is a paid mutator transaction binding the contract method 0xca6eec78.
//
// Solidity: function revokeCredential(bytes32 credentialHash) returns(bool)
func (_PermImpl *PermImplTransactor) RevokeCredential(opts *bind.TransactOpts, credentialHash [32]byte) (*types.Transaction, error) {
	return _PermImpl.contract.Transact(opts, "revokeCredential", credentialHash)
}

// RevokeCredential is a paid mutator transaction binding the contract method 0xca6eec78.
//
// Solidity: function revokeCredential(bytes32 credentialHash) returns(bool)
func (_PermImpl *PermImplSession) RevokeCredential(credentialHash [32]byte) (*types.Transaction, error) {
	return _PermImpl.Contract.RevokeCredential(&_PermImpl.TransactOpts, credentialHash)
}

// RevokeCredential is a paid mutator transaction binding the contract method 0xca6eec78.
//
// Solidity: function revokeCredential(bytes32 credentialHash) returns(bool)
func (_PermImpl *PermImplTransactorSession) RevokeCredential(credentialHash [32]byte) (*types.Transaction, error) {
	return _PermImpl.Contract.RevokeCredential(&_PermImpl.TransactOpts, credentialHash)
}

// PermImplCredentialRegisteredIterator is returned from FilterCredentialRegistered and is used to iterate over the raw logs and unpacked data for CredentialRegistered events raised by the PermImpl contract.
type PermImplCredentialRegisteredIterator struct {
	Event *PermImplCredentialRegistered // Event containing the contract specifics and raw log

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
func (it *PermImplCredentialRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermImplCredentialRegistered)
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
		it.Event = new(PermImplCredentialRegistered)
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
func (it *PermImplCredentialRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermImplCredentialRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermImplCredentialRegistered represents a CredentialRegistered event raised by the PermImpl contract.
type PermImplCredentialRegistered struct {
	CredentialHash [32]byte
	By             common.Address
	Id             common.Address
	Iat            *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCredentialRegistered is a free log retrieval operation binding the contract event 0x889569e4303664728fba29a472b24d3032707e6ebbee9a46e070c5c8f2ae4c25.
//
// Solidity: event CredentialRegistered(bytes32 indexed credentialHash, address by, address id, uint256 iat)
func (_PermImpl *PermImplFilterer) FilterCredentialRegistered(opts *bind.FilterOpts, credentialHash [][32]byte) (*PermImplCredentialRegisteredIterator, error) {

	var credentialHashRule []interface{}
	for _, credentialHashItem := range credentialHash {
		credentialHashRule = append(credentialHashRule, credentialHashItem)
	}

	logs, sub, err := _PermImpl.contract.FilterLogs(opts, "CredentialRegistered", credentialHashRule)
	if err != nil {
		return nil, err
	}
	return &PermImplCredentialRegisteredIterator{contract: _PermImpl.contract, event: "CredentialRegistered", logs: logs, sub: sub}, nil
}

var CredentialRegisteredTopicHash = "0x889569e4303664728fba29a472b24d3032707e6ebbee9a46e070c5c8f2ae4c25"

// WatchCredentialRegistered is a free log subscription operation binding the contract event 0x889569e4303664728fba29a472b24d3032707e6ebbee9a46e070c5c8f2ae4c25.
//
// Solidity: event CredentialRegistered(bytes32 indexed credentialHash, address by, address id, uint256 iat)
func (_PermImpl *PermImplFilterer) WatchCredentialRegistered(opts *bind.WatchOpts, sink chan<- *PermImplCredentialRegistered, credentialHash [][32]byte) (event.Subscription, error) {

	var credentialHashRule []interface{}
	for _, credentialHashItem := range credentialHash {
		credentialHashRule = append(credentialHashRule, credentialHashItem)
	}

	logs, sub, err := _PermImpl.contract.WatchLogs(opts, "CredentialRegistered", credentialHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermImplCredentialRegistered)
				if err := _PermImpl.contract.UnpackLog(event, "CredentialRegistered", log); err != nil {
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

// ParseCredentialRegistered is a log parse operation binding the contract event 0x889569e4303664728fba29a472b24d3032707e6ebbee9a46e070c5c8f2ae4c25.
//
// Solidity: event CredentialRegistered(bytes32 indexed credentialHash, address by, address id, uint256 iat)
func (_PermImpl *PermImplFilterer) ParseCredentialRegistered(log types.Log) (*PermImplCredentialRegistered, error) {
	event := new(PermImplCredentialRegistered)
	if err := _PermImpl.contract.UnpackLog(event, "CredentialRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PermImplCredentialRevokedIterator is returned from FilterCredentialRevoked and is used to iterate over the raw logs and unpacked data for CredentialRevoked events raised by the PermImpl contract.
type PermImplCredentialRevokedIterator struct {
	Event *PermImplCredentialRevoked // Event containing the contract specifics and raw log

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
func (it *PermImplCredentialRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PermImplCredentialRevoked)
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
		it.Event = new(PermImplCredentialRevoked)
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
func (it *PermImplCredentialRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PermImplCredentialRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PermImplCredentialRevoked represents a CredentialRevoked event raised by the PermImpl contract.
type PermImplCredentialRevoked struct {
	CredentialHash [32]byte
	By             common.Address
	Date           *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCredentialRevoked is a free log retrieval operation binding the contract event 0x80d4746828447d65eceb938c7f9d85fbeb232fe7819d5209ed452b3940f8904b.
//
// Solidity: event CredentialRevoked(bytes32 indexed credentialHash, address by, uint256 date)
func (_PermImpl *PermImplFilterer) FilterCredentialRevoked(opts *bind.FilterOpts, credentialHash [][32]byte) (*PermImplCredentialRevokedIterator, error) {

	var credentialHashRule []interface{}
	for _, credentialHashItem := range credentialHash {
		credentialHashRule = append(credentialHashRule, credentialHashItem)
	}

	logs, sub, err := _PermImpl.contract.FilterLogs(opts, "CredentialRevoked", credentialHashRule)
	if err != nil {
		return nil, err
	}
	return &PermImplCredentialRevokedIterator{contract: _PermImpl.contract, event: "CredentialRevoked", logs: logs, sub: sub}, nil
}

var CredentialRevokedTopicHash = "0x80d4746828447d65eceb938c7f9d85fbeb232fe7819d5209ed452b3940f8904b"

// WatchCredentialRevoked is a free log subscription operation binding the contract event 0x80d4746828447d65eceb938c7f9d85fbeb232fe7819d5209ed452b3940f8904b.
//
// Solidity: event CredentialRevoked(bytes32 indexed credentialHash, address by, uint256 date)
func (_PermImpl *PermImplFilterer) WatchCredentialRevoked(opts *bind.WatchOpts, sink chan<- *PermImplCredentialRevoked, credentialHash [][32]byte) (event.Subscription, error) {

	var credentialHashRule []interface{}
	for _, credentialHashItem := range credentialHash {
		credentialHashRule = append(credentialHashRule, credentialHashItem)
	}

	logs, sub, err := _PermImpl.contract.WatchLogs(opts, "CredentialRevoked", credentialHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PermImplCredentialRevoked)
				if err := _PermImpl.contract.UnpackLog(event, "CredentialRevoked", log); err != nil {
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

// ParseCredentialRevoked is a log parse operation binding the contract event 0x80d4746828447d65eceb938c7f9d85fbeb232fe7819d5209ed452b3940f8904b.
//
// Solidity: event CredentialRevoked(bytes32 indexed credentialHash, address by, uint256 date)
func (_PermImpl *PermImplFilterer) ParseCredentialRevoked(log types.Log) (*PermImplCredentialRevoked, error) {
	event := new(PermImplCredentialRevoked)
	if err := _PermImpl.contract.UnpackLog(event, "CredentialRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
