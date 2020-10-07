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
const AccessControlContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"newAddrRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"newAddrRemove\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addAccountToRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"pubKey\",\"type\":\"string\"}],\"name\":\"addPubKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAdminAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getPubKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"greet\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"removeAccountFromRegister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
var AccessControlContractBin = "0x6080604052600080546001600160a01b0319167321a018606490c031a8c02bb6b992d8ae44add37f17905534801561003657600080fd5b50610530806100466000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80636d645b441161005b5780636d645b441461018f5780638b7e87ce14610197578063cfae3217146101b4578063fce9512a146102315761007d565b806321f8a7211461008257806351c8fa97146100bb578063670d65ea146100e9575b600080fd5b61009f6004803603602081101561009857600080fd5b5035610257565b604080516001600160a01b039092168252519081900360200190f35b6100e7600480360360408110156100d157600080fd5b50803590602001356001600160a01b0316610272565b005b6100e7600480360360208110156100ff57600080fd5b81019060208101813564010000000081111561011a57600080fd5b82018360208201111561012c57600080fd5b8035906020019184600183028401116401000000008311171561014e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506102d8945050505050565b61009f6102fc565b6100e7600480360360208110156101ad57600080fd5b503561030c565b6101bc610366565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101f65781810151838201526020016101de565b50505050905090810190601f1680156102235780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101bc6004803603602081101561024757600080fd5b50356001600160a01b0316610386565b6000908152600160205260409020546001600160a01b031690565b6000546001600160a01b0316331461028657fe5b60008281526001602052604080822080546001600160a01b0319166001600160a01b0385161790555183917fe2888d2900e8be92bf075b2e7c635f9813c3e18afb476d80b6dd545ad34d717391a25050565b33600090815260026020908152604090912082516102f89284019061042f565b5050565b6000546001600160a01b03165b90565b6000546001600160a01b0316331461032057fe5b60008181526001602052604080822080546001600160a01b03191690555182917f761c7f5deea18480fd4c3286fb929a518afd9c1faa2b9a6f4669ba2f86995a4a91a250565b60606040518060600160405280603481526020016104c860349139905090565b6001600160a01b038116600090815260026020818152604092839020805484516001821615610100026000190190911693909304601f810183900483028401830190945283835260609390918301828280156104235780601f106103f857610100808354040283529160200191610423565b820191906000526020600020905b81548152906001019060200180831161040657829003601f168201915b50505050509050919050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061047057805160ff191683800117855561049d565b8280016001018555821561049d579182015b8281111561049d578251825591602001919060010190610482565b506104a99291506104ad565b5090565b61030991905b808211156104a957600081556001016104b356fe48656c6c6f20796f7520686176652063616c6c65642074686520636f6e747261637420616363657373436f6e74726f6c2e736f6ca265627a7a72315820795b03fcda99f1047c70144fa46939920e1993231742465bc30b79c9dcdb085264736f6c63430005100032"

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
