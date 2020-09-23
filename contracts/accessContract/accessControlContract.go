// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package accessControlContract

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

// AccessControlContractABI is the input ABI used to generate the binding from.
const AccessControlContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAccountToRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"pubKey\",\"type\":\"string\"}],\"name\":\"addPubKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdminAddr\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"removeAccountFromRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"greet\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getPubKey\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"newAddrRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"newAddrRemove\",\"type\":\"event\"}]"

// AccessControlContractFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlContractFuncSigs = map[string]string{
	"51c8fa97": "addAccountToRegister(bytes32,address)",
	"670d65ea": "addPubKey(string)",
	"21f8a721": "getAddress(bytes32)",
	"6d645b44": "getAdminAddr()",
	"fce9512a": "getPubKey(address)",
	"cfae3217": "greet()",
	"8b7e87ce": "removeAccountFromRegister(bytes32)",
}

// AccessControlContractBin is the compiled bytecode used for deploying new contracts.
var AccessControlContractBin = "0x608060405260008054600160a060020a0319167321a018606490c031a8c02bb6b992d8ae44add37f17905534801561003657600080fd5b50610507806100466000396000f3006080604052600436106100825763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166321f8a721811461008757806351c8fa97146100bb578063670d65ea146100e15780636d645b441461013a5780638b7e87ce1461014f578063cfae321714610167578063fce9512a146101f1575b600080fd5b34801561009357600080fd5b5061009f600435610212565b60408051600160a060020a039092168252519081900360200190f35b3480156100c757600080fd5b506100df600435600160a060020a036024351661022d565b005b3480156100ed57600080fd5b506040805160206004803580820135601f81018490048402850184019095528484526100df9436949293602493928401919081908401838280828437509497506102a09650505050505050565b34801561014657600080fd5b5061009f6102c4565b34801561015b57600080fd5b506100df6004356102d4565b34801561017357600080fd5b5061017c61033b565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101b657818101518382015260200161019e565b50505050905090810190601f1680156101e35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101fd57600080fd5b5061017c600160a060020a036004351661039a565b600090815260016020526040902054600160a060020a031690565b600054600160a060020a0316331461024157fe5b600082815260016020526040808220805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0385161790555183917fe2888d2900e8be92bf075b2e7c635f9813c3e18afb476d80b6dd545ad34d717391a25050565b33600090815260026020908152604090912082516102c092840190610443565b5050565b600054600160a060020a03165b90565b600054600160a060020a031633146102e857fe5b600081815260016020526040808220805473ffffffffffffffffffffffffffffffffffffffff191690555182917f761c7f5deea18480fd4c3286fb929a518afd9c1faa2b9a6f4669ba2f86995a4a91a250565b60408051606081018252603481527f48656c6c6f20796f7520686176652063616c6c65642074686520636f6e74726160208201527f637420616363657373436f6e74726f6c2e736f6c0000000000000000000000009181019190915290565b600160a060020a038116600090815260026020818152604092839020805484516001821615610100026000190190911693909304601f810183900483028401830190945283835260609390918301828280156104375780601f1061040c57610100808354040283529160200191610437565b820191906000526020600020905b81548152906001019060200180831161041a57829003601f168201915b50505050509050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061048457805160ff19168380011785556104b1565b828001600101855582156104b1579182015b828111156104b1578251825591602001919060010190610496565b506104bd9291506104c1565b5090565b6102d191905b808211156104bd57600081556001016104c75600a165627a7a723058207c16b075425a71dcd86b2db7dcd87a273c765f0750a33837fce0e2f234f0eb0a0029"

// DeployAccessControlContract deploys a new Ethereum contract, binding an instance of AccessControlContract to it.
func DeployAccessControlContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AccessControlContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AccessControlContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AccessControlContract{AccessControlContractCaller: AccessControlContractCaller{contract: contract}, AccessControlContractTransactor: AccessControlContractTransactor{contract: contract}, AccessControlContractFilterer: AccessControlContractFilterer{contract: contract}}, nil
}

// AccessControlContract is an auto generated Go binding around an Ethereum contract.
type AccessControlContract struct {
	AccessControlContractCaller     // Read-only binding to the contract
	AccessControlContractTransactor // Write-only binding to the contract
	AccessControlContractFilterer   // Log filterer for contract events
}

// AccessControlContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlContractSession struct {
	Contract     *AccessControlContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AccessControlContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlContractCallerSession struct {
	Contract *AccessControlContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// AccessControlContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlContractTransactorSession struct {
	Contract     *AccessControlContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// AccessControlContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlContractRaw struct {
	Contract *AccessControlContract // Generic contract binding to access the raw methods on
}

// AccessControlContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlContractCallerRaw struct {
	Contract *AccessControlContractCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlContractTransactorRaw struct {
	Contract *AccessControlContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControlContract creates a new instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContract(address common.Address, backend bind.ContractBackend) (*AccessControlContract, error) {
	contract, err := bindAccessControlContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControlContract{AccessControlContractCaller: AccessControlContractCaller{contract: contract}, AccessControlContractTransactor: AccessControlContractTransactor{contract: contract}, AccessControlContractFilterer: AccessControlContractFilterer{contract: contract}}, nil
}

// NewAccessControlContractCaller creates a new read-only instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContractCaller(address common.Address, caller bind.ContractCaller) (*AccessControlContractCaller, error) {
	contract, err := bindAccessControlContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractCaller{contract: contract}, nil
}

// NewAccessControlContractTransactor creates a new write-only instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContractTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlContractTransactor, error) {
	contract, err := bindAccessControlContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractTransactor{contract: contract}, nil
}

// NewAccessControlContractFilterer creates a new log filterer instance of AccessControlContract, bound to a specific deployed contract.
func NewAccessControlContractFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlContractFilterer, error) {
	contract, err := bindAccessControlContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractFilterer{contract: contract}, nil
}

// bindAccessControlContract binds a generic wrapper to an already deployed contract.
func bindAccessControlContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlContract *AccessControlContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccessControlContract.Contract.AccessControlContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlContract *AccessControlContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AccessControlContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlContract *AccessControlContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AccessControlContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlContract *AccessControlContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AccessControlContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlContract *AccessControlContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlContract *AccessControlContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlContract.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_AccessControlContract *AccessControlContractCaller) GetAddress(opts *bind.CallOpts, id [32]byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccessControlContract.contract.Call(opts, out, "getAddress", id)
	return *ret0, err
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_AccessControlContract *AccessControlContractSession) GetAddress(id [32]byte) (common.Address, error) {
	return _AccessControlContract.Contract.GetAddress(&_AccessControlContract.CallOpts, id)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_AccessControlContract *AccessControlContractCallerSession) GetAddress(id [32]byte) (common.Address, error) {
	return _AccessControlContract.Contract.GetAddress(&_AccessControlContract.CallOpts, id)
}

// GetAdminAddr is a free data retrieval call binding the contract method 0x6d645b44.
//
// Solidity: function getAdminAddr() view returns(address)
func (_AccessControlContract *AccessControlContractCaller) GetAdminAddr(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AccessControlContract.contract.Call(opts, out, "getAdminAddr")
	return *ret0, err
}

// GetAdminAddr is a free data retrieval call binding the contract method 0x6d645b44.
//
// Solidity: function getAdminAddr() view returns(address)
func (_AccessControlContract *AccessControlContractSession) GetAdminAddr() (common.Address, error) {
	return _AccessControlContract.Contract.GetAdminAddr(&_AccessControlContract.CallOpts)
}

// GetAdminAddr is a free data retrieval call binding the contract method 0x6d645b44.
//
// Solidity: function getAdminAddr() view returns(address)
func (_AccessControlContract *AccessControlContractCallerSession) GetAdminAddr() (common.Address, error) {
	return _AccessControlContract.Contract.GetAdminAddr(&_AccessControlContract.CallOpts)
}

// GetPubKey is a free data retrieval call binding the contract method 0xfce9512a.
//
// Solidity: function getPubKey(address addr) view returns(string)
func (_AccessControlContract *AccessControlContractCaller) GetPubKey(opts *bind.CallOpts, addr common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AccessControlContract.contract.Call(opts, out, "getPubKey", addr)
	return *ret0, err
}

// GetPubKey is a free data retrieval call binding the contract method 0xfce9512a.
//
// Solidity: function getPubKey(address addr) view returns(string)
func (_AccessControlContract *AccessControlContractSession) GetPubKey(addr common.Address) (string, error) {
	return _AccessControlContract.Contract.GetPubKey(&_AccessControlContract.CallOpts, addr)
}

// GetPubKey is a free data retrieval call binding the contract method 0xfce9512a.
//
// Solidity: function getPubKey(address addr) view returns(string)
func (_AccessControlContract *AccessControlContractCallerSession) GetPubKey(addr common.Address) (string, error) {
	return _AccessControlContract.Contract.GetPubKey(&_AccessControlContract.CallOpts, addr)
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() pure returns(string)
func (_AccessControlContract *AccessControlContractCaller) Greet(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AccessControlContract.contract.Call(opts, out, "greet")
	return *ret0, err
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() pure returns(string)
func (_AccessControlContract *AccessControlContractSession) Greet() (string, error) {
	return _AccessControlContract.Contract.Greet(&_AccessControlContract.CallOpts)
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() pure returns(string)
func (_AccessControlContract *AccessControlContractCallerSession) Greet() (string, error) {
	return _AccessControlContract.Contract.Greet(&_AccessControlContract.CallOpts)
}

// AddAccountToRegister is a paid mutator transaction binding the contract method 0x51c8fa97.
//
// Solidity: function addAccountToRegister(bytes32 id, address account) returns()
func (_AccessControlContract *AccessControlContractTransactor) AddAccountToRegister(opts *bind.TransactOpts, id [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlContract.contract.Transact(opts, "addAccountToRegister", id, account)
}

// AddAccountToRegister is a paid mutator transaction binding the contract method 0x51c8fa97.
//
// Solidity: function addAccountToRegister(bytes32 id, address account) returns()
func (_AccessControlContract *AccessControlContractSession) AddAccountToRegister(id [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AddAccountToRegister(&_AccessControlContract.TransactOpts, id, account)
}

// AddAccountToRegister is a paid mutator transaction binding the contract method 0x51c8fa97.
//
// Solidity: function addAccountToRegister(bytes32 id, address account) returns()
func (_AccessControlContract *AccessControlContractTransactorSession) AddAccountToRegister(id [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AddAccountToRegister(&_AccessControlContract.TransactOpts, id, account)
}

// AddPubKey is a paid mutator transaction binding the contract method 0x670d65ea.
//
// Solidity: function addPubKey(string pubKey) returns()
func (_AccessControlContract *AccessControlContractTransactor) AddPubKey(opts *bind.TransactOpts, pubKey string) (*types.Transaction, error) {
	return _AccessControlContract.contract.Transact(opts, "addPubKey", pubKey)
}

// AddPubKey is a paid mutator transaction binding the contract method 0x670d65ea.
//
// Solidity: function addPubKey(string pubKey) returns()
func (_AccessControlContract *AccessControlContractSession) AddPubKey(pubKey string) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AddPubKey(&_AccessControlContract.TransactOpts, pubKey)
}

// AddPubKey is a paid mutator transaction binding the contract method 0x670d65ea.
//
// Solidity: function addPubKey(string pubKey) returns()
func (_AccessControlContract *AccessControlContractTransactorSession) AddPubKey(pubKey string) (*types.Transaction, error) {
	return _AccessControlContract.Contract.AddPubKey(&_AccessControlContract.TransactOpts, pubKey)
}

// RemoveAccountFromRegister is a paid mutator transaction binding the contract method 0x8b7e87ce.
//
// Solidity: function removeAccountFromRegister(bytes32 id) returns()
func (_AccessControlContract *AccessControlContractTransactor) RemoveAccountFromRegister(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _AccessControlContract.contract.Transact(opts, "removeAccountFromRegister", id)
}

// RemoveAccountFromRegister is a paid mutator transaction binding the contract method 0x8b7e87ce.
//
// Solidity: function removeAccountFromRegister(bytes32 id) returns()
func (_AccessControlContract *AccessControlContractSession) RemoveAccountFromRegister(id [32]byte) (*types.Transaction, error) {
	return _AccessControlContract.Contract.RemoveAccountFromRegister(&_AccessControlContract.TransactOpts, id)
}

// RemoveAccountFromRegister is a paid mutator transaction binding the contract method 0x8b7e87ce.
//
// Solidity: function removeAccountFromRegister(bytes32 id) returns()
func (_AccessControlContract *AccessControlContractTransactorSession) RemoveAccountFromRegister(id [32]byte) (*types.Transaction, error) {
	return _AccessControlContract.Contract.RemoveAccountFromRegister(&_AccessControlContract.TransactOpts, id)
}

// AccessControlContractNewAddrRegisteredIterator is returned from FilterNewAddrRegistered and is used to iterate over the raw logs and unpacked data for NewAddrRegistered events raised by the AccessControlContract contract.
type AccessControlContractNewAddrRegisteredIterator struct {
	Event *AccessControlContractNewAddrRegistered // Event containing the contract specifics and raw log

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
func (it *AccessControlContractNewAddrRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlContractNewAddrRegistered)
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
		it.Event = new(AccessControlContractNewAddrRegistered)
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
func (it *AccessControlContractNewAddrRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlContractNewAddrRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlContractNewAddrRegistered represents a NewAddrRegistered event raised by the AccessControlContract contract.
type AccessControlContractNewAddrRegistered struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewAddrRegistered is a free log retrieval operation binding the contract event 0xe2888d2900e8be92bf075b2e7c635f9813c3e18afb476d80b6dd545ad34d7173.
//
// Solidity: event newAddrRegistered(bytes32 indexed _id)
func (_AccessControlContract *AccessControlContractFilterer) FilterNewAddrRegistered(opts *bind.FilterOpts, _id [][32]byte) (*AccessControlContractNewAddrRegisteredIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _AccessControlContract.contract.FilterLogs(opts, "newAddrRegistered", _idRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractNewAddrRegisteredIterator{contract: _AccessControlContract.contract, event: "newAddrRegistered", logs: logs, sub: sub}, nil
}

// WatchNewAddrRegistered is a free log subscription operation binding the contract event 0xe2888d2900e8be92bf075b2e7c635f9813c3e18afb476d80b6dd545ad34d7173.
//
// Solidity: event newAddrRegistered(bytes32 indexed _id)
func (_AccessControlContract *AccessControlContractFilterer) WatchNewAddrRegistered(opts *bind.WatchOpts, sink chan<- *AccessControlContractNewAddrRegistered, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _AccessControlContract.contract.WatchLogs(opts, "newAddrRegistered", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlContractNewAddrRegistered)
				if err := _AccessControlContract.contract.UnpackLog(event, "newAddrRegistered", log); err != nil {
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

// ParseNewAddrRegistered is a log parse operation binding the contract event 0xe2888d2900e8be92bf075b2e7c635f9813c3e18afb476d80b6dd545ad34d7173.
//
// Solidity: event newAddrRegistered(bytes32 indexed _id)
func (_AccessControlContract *AccessControlContractFilterer) ParseNewAddrRegistered(log types.Log) (*AccessControlContractNewAddrRegistered, error) {
	event := new(AccessControlContractNewAddrRegistered)
	if err := _AccessControlContract.contract.UnpackLog(event, "newAddrRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccessControlContractNewAddrRemoveIterator is returned from FilterNewAddrRemove and is used to iterate over the raw logs and unpacked data for NewAddrRemove events raised by the AccessControlContract contract.
type AccessControlContractNewAddrRemoveIterator struct {
	Event *AccessControlContractNewAddrRemove // Event containing the contract specifics and raw log

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
func (it *AccessControlContractNewAddrRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlContractNewAddrRemove)
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
		it.Event = new(AccessControlContractNewAddrRemove)
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
func (it *AccessControlContractNewAddrRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlContractNewAddrRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlContractNewAddrRemove represents a NewAddrRemove event raised by the AccessControlContract contract.
type AccessControlContractNewAddrRemove struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterNewAddrRemove is a free log retrieval operation binding the contract event 0x761c7f5deea18480fd4c3286fb929a518afd9c1faa2b9a6f4669ba2f86995a4a.
//
// Solidity: event newAddrRemove(bytes32 indexed _id)
func (_AccessControlContract *AccessControlContractFilterer) FilterNewAddrRemove(opts *bind.FilterOpts, _id [][32]byte) (*AccessControlContractNewAddrRemoveIterator, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _AccessControlContract.contract.FilterLogs(opts, "newAddrRemove", _idRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlContractNewAddrRemoveIterator{contract: _AccessControlContract.contract, event: "newAddrRemove", logs: logs, sub: sub}, nil
}

// WatchNewAddrRemove is a free log subscription operation binding the contract event 0x761c7f5deea18480fd4c3286fb929a518afd9c1faa2b9a6f4669ba2f86995a4a.
//
// Solidity: event newAddrRemove(bytes32 indexed _id)
func (_AccessControlContract *AccessControlContractFilterer) WatchNewAddrRemove(opts *bind.WatchOpts, sink chan<- *AccessControlContractNewAddrRemove, _id [][32]byte) (event.Subscription, error) {

	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _AccessControlContract.contract.WatchLogs(opts, "newAddrRemove", _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlContractNewAddrRemove)
				if err := _AccessControlContract.contract.UnpackLog(event, "newAddrRemove", log); err != nil {
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

// ParseNewAddrRemove is a log parse operation binding the contract event 0x761c7f5deea18480fd4c3286fb929a518afd9c1faa2b9a6f4669ba2f86995a4a.
//
// Solidity: event newAddrRemove(bytes32 indexed _id)
func (_AccessControlContract *AccessControlContractFilterer) ParseNewAddrRemove(log types.Log) (*AccessControlContractNewAddrRemove, error) {
	event := new(AccessControlContractNewAddrRemove)
	if err := _AccessControlContract.contract.UnpackLog(event, "newAddrRemove", log); err != nil {
		return nil, err
	}
	return event, nil
}
