// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balanceContract

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

// BalanceContractABI is the input ABI used to generate the binding from.
const BalanceContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_subID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"}],\"name\":\"notifyNew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"notifyNewCategory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_subID\",\"type\":\"bytes32\"}],\"name\":\"notifyRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"notifyRemoveCategory\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subName\",\"type\":\"bytes32\"}],\"name\":\"addNewType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subName\",\"type\":\"bytes32\"}],\"name\":\"deleteSub\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subName\",\"type\":\"bytes32\"}],\"name\":\"deleteType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllSubscriptions\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllTopics\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"greet\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subName\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"subscribeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BalanceContractFuncSigs maps the 4-byte function signature to its string representation.
var BalanceContractFuncSigs = map[string]string{
	"149cde75": "addNewType(bytes32)",
	"f356cc79": "deleteSub(bytes32)",
	"1d9de32c": "deleteType(bytes32)",
	"56eb00a1": "getAllSubscriptions()",
	"86416f14": "getAllTopics()",
	"cfae3217": "greet()",
	"c990824d": "subscribeTo(bytes32,uint256)",
}

// BalanceContractBin is the compiled bytecode used for deploying new contracts.
var BalanceContractBin = "0x6080604052600080546001600160a01b0319167321a018606490c031a8c02bb6b992d8ae44add37f17905534801561003657600080fd5b5061086f806100466000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806386416f141161005b57806386416f1414610116578063c990824d1461011e578063cfae321714610141578063f356cc79146101be5761007d565b8063149cde75146100825780631d9de32c146100a157806356eb00a1146100be575b600080fd5b61009f6004803603602081101561009857600080fd5b50356101db565b005b61009f600480360360208110156100b757600080fd5b50356102fa565b6100c66103b1565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101025781810151838201526020016100ea565b505050509050019250505060405180910390f35b6100c6610415565b61009f6004803603604081101561013457600080fd5b508035906020013561046e565b6101496105c1565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561018357818101518382015260200161016b565b50505050905090810190601f1680156101b05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61009f600480360360208110156101d457600080fd5b50356105e1565b6000546001600160a01b031633146102245760405162461bcd60e51b81526004018080602001828103825260248152602001806108176024913960400191505060405180910390fd5b60008181526002602052604090205460ff16156102725760405162461bcd60e51b81526004018080602001828103825260248152602001806107f36024913960400191505060405180910390fd5b6000818152600260209081526040808320805460ff191660019081179091556003805491820181557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9091018590555460049092528083209190915551829133917fde3cd713f10754ba21eabefe59d58e9251b68f9ff9016fc20324b4e261f1d5989190a350565b6000546001600160a01b031633146103435760405162461bcd60e51b81526004018080602001828103825260248152602001806108176024913960400191505060405180910390fd5b60008181526002602052604090205460ff1615156001146103955760405162461bcd60e51b81526004018080602001828103825260258152602001806107716025913960400191505060405180910390fd5b6000818152600260205260409020805460ff1916905560038054fe5b3360009081526001602090815260409182902060020180548351818402810184019094528084526060939283018282801561040b57602002820191906000526020600020905b8154815260200190600101908083116103f7575b5050505050905090565b6060600260010180548060200260200160405190810160405280929190818152602001828054801561040b57602002820191906000526020600020908154815260200190600101908083116103f7575050505050905090565b60008281526002602052604090205460ff1615156001146104c05760405162461bcd60e51b81526004018080602001828103825260258152602001806107716025913960400191505060405180910390fd5b33600090815260016020908152604080832085845290915290205460ff161561051a5760405162461bcd60e51b815260040180806020018281038252602a8152602001806107c9602a913960400191505060405180910390fd5b336000818152600160208181526040808420878552808352818520805460ff199081168617909155818501845282862042890190819055600280840180548089018255818a52878a20018c9055548b89526003909401865284882093909355918452948290208054909516909317909355825191825291518593927fc7c18d7fd8f2b08630c0d4f0bcf252583ff50a5cf2c5d8811f8a4347191679d4928290030190a35050565b606060405180606001604052806033815260200161079660339139905090565b3360009081526001602081815260408084208585529091529091205460ff1615151461063e5760405162461bcd60e51b815260040180806020018281038252602e815260200180610743602e913960400191505060405180910390fd5b336000908152600160208181526040808420858552909201905290205442106106ae576040805162461bcd60e51b815260206004820152601a60248201527f537562736372697074696f6e206973206e6f7420616374697665000000000000604482015290519081900360640190fd5b33600081815260016020818152604080842086855260038101835281852054818452918520805460ff19169055949093525260029091018054829081106106f157fe5b60009182526020808320909101829055838252600290526040808220805460ff1916905551839133917fab7ba7d234105b15e41e97f548ae644ed346089e33e01f8f03a862bfaebacb919190a3505056fe5468652075736572206973206e6f742073757363726962656420746f207468652072657175657374652074797065546865207265717565737465642063617465676f727920646f6573206e6f7420657869737448656c6c6f20796f7520686176652063616c6c65642074686520636f6e74726163742062616c616e63652d737562732e736f6c546865207573657220697320616c72656164792073756273637269626520746f20746861742074797065546865207265717565737465642063617465676f727920616c72616479206578697374735573657220646f6573206e6f74206861766520656e6f7567682070726976696c65676573a265627a7a7231582023a2062e4b7787f79dc73da5f311fe6dd3f50f84ddebdd0c21314e9bd7cbc99f64736f6c63430005100032"

// DeployBalanceContract deploys a new Ethereum contract, binding an instance of BalanceContract to it.
func DeployBalanceContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BalanceContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BalanceContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BalanceContract{BalanceContractCaller: BalanceContractCaller{contract: contract}, BalanceContractTransactor: BalanceContractTransactor{contract: contract}, BalanceContractFilterer: BalanceContractFilterer{contract: contract}}, nil
}

// BalanceContract is an auto generated Go binding around an Ethereum contract.
type BalanceContract struct {
	BalanceContractCaller     // Read-only binding to the contract
	BalanceContractTransactor // Write-only binding to the contract
	BalanceContractFilterer   // Log filterer for contract events
}

// BalanceContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalanceContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalanceContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalanceContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalanceContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalanceContractSession struct {
	Contract     *BalanceContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalanceContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalanceContractCallerSession struct {
	Contract *BalanceContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// BalanceContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalanceContractTransactorSession struct {
	Contract     *BalanceContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// BalanceContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalanceContractRaw struct {
	Contract *BalanceContract // Generic contract binding to access the raw methods on
}

// BalanceContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalanceContractCallerRaw struct {
	Contract *BalanceContractCaller // Generic read-only contract binding to access the raw methods on
}

// BalanceContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalanceContractTransactorRaw struct {
	Contract *BalanceContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalanceContract creates a new instance of BalanceContract, bound to a specific deployed contract.
func NewBalanceContract(address common.Address, backend bind.ContractBackend) (*BalanceContract, error) {
	contract, err := bindBalanceContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalanceContract{BalanceContractCaller: BalanceContractCaller{contract: contract}, BalanceContractTransactor: BalanceContractTransactor{contract: contract}, BalanceContractFilterer: BalanceContractFilterer{contract: contract}}, nil
}

// NewBalanceContractCaller creates a new read-only instance of BalanceContract, bound to a specific deployed contract.
func NewBalanceContractCaller(address common.Address, caller bind.ContractCaller) (*BalanceContractCaller, error) {
	contract, err := bindBalanceContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceContractCaller{contract: contract}, nil
}

// NewBalanceContractTransactor creates a new write-only instance of BalanceContract, bound to a specific deployed contract.
func NewBalanceContractTransactor(address common.Address, transactor bind.ContractTransactor) (*BalanceContractTransactor, error) {
	contract, err := bindBalanceContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalanceContractTransactor{contract: contract}, nil
}

// NewBalanceContractFilterer creates a new log filterer instance of BalanceContract, bound to a specific deployed contract.
func NewBalanceContractFilterer(address common.Address, filterer bind.ContractFilterer) (*BalanceContractFilterer, error) {
	contract, err := bindBalanceContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalanceContractFilterer{contract: contract}, nil
}

// bindBalanceContract binds a generic wrapper to an already deployed contract.
func bindBalanceContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalanceContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceContract *BalanceContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BalanceContract.Contract.BalanceContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceContract *BalanceContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceContract.Contract.BalanceContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceContract *BalanceContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceContract.Contract.BalanceContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalanceContract *BalanceContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BalanceContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalanceContract *BalanceContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalanceContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalanceContract *BalanceContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalanceContract.Contract.contract.Transact(opts, method, params...)
}

// GetAllSubscriptions is a free data retrieval call binding the contract method 0x56eb00a1.
//
// Solidity: function getAllSubscriptions() view returns(bytes32[])
func (_BalanceContract *BalanceContractCaller) GetAllSubscriptions(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _BalanceContract.contract.Call(opts, out, "getAllSubscriptions")
	return *ret0, err
}

// GetAllSubscriptions is a free data retrieval call binding the contract method 0x56eb00a1.
//
// Solidity: function getAllSubscriptions() view returns(bytes32[])
func (_BalanceContract *BalanceContractSession) GetAllSubscriptions() ([][32]byte, error) {
	return _BalanceContract.Contract.GetAllSubscriptions(&_BalanceContract.CallOpts)
}

// GetAllSubscriptions is a free data retrieval call binding the contract method 0x56eb00a1.
//
// Solidity: function getAllSubscriptions() view returns(bytes32[])
func (_BalanceContract *BalanceContractCallerSession) GetAllSubscriptions() ([][32]byte, error) {
	return _BalanceContract.Contract.GetAllSubscriptions(&_BalanceContract.CallOpts)
}

// GetAllTopics is a free data retrieval call binding the contract method 0x86416f14.
//
// Solidity: function getAllTopics() view returns(bytes32[])
func (_BalanceContract *BalanceContractCaller) GetAllTopics(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _BalanceContract.contract.Call(opts, out, "getAllTopics")
	return *ret0, err
}

// GetAllTopics is a free data retrieval call binding the contract method 0x86416f14.
//
// Solidity: function getAllTopics() view returns(bytes32[])
func (_BalanceContract *BalanceContractSession) GetAllTopics() ([][32]byte, error) {
	return _BalanceContract.Contract.GetAllTopics(&_BalanceContract.CallOpts)
}

// GetAllTopics is a free data retrieval call binding the contract method 0x86416f14.
//
// Solidity: function getAllTopics() view returns(bytes32[])
func (_BalanceContract *BalanceContractCallerSession) GetAllTopics() ([][32]byte, error) {
	return _BalanceContract.Contract.GetAllTopics(&_BalanceContract.CallOpts)
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() pure returns(string)
func (_BalanceContract *BalanceContractCaller) Greet(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BalanceContract.contract.Call(opts, out, "greet")
	return *ret0, err
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() pure returns(string)
func (_BalanceContract *BalanceContractSession) Greet() (string, error) {
	return _BalanceContract.Contract.Greet(&_BalanceContract.CallOpts)
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() pure returns(string)
func (_BalanceContract *BalanceContractCallerSession) Greet() (string, error) {
	return _BalanceContract.Contract.Greet(&_BalanceContract.CallOpts)
}

// AddNewType is a paid mutator transaction binding the contract method 0x149cde75.
//
// Solidity: function addNewType(bytes32 subName) returns()
func (_BalanceContract *BalanceContractTransactor) AddNewType(opts *bind.TransactOpts, subName [32]byte) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "addNewType", subName)
}

// AddNewType is a paid mutator transaction binding the contract method 0x149cde75.
//
// Solidity: function addNewType(bytes32 subName) returns()
func (_BalanceContract *BalanceContractSession) AddNewType(subName [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.AddNewType(&_BalanceContract.TransactOpts, subName)
}

// AddNewType is a paid mutator transaction binding the contract method 0x149cde75.
//
// Solidity: function addNewType(bytes32 subName) returns()
func (_BalanceContract *BalanceContractTransactorSession) AddNewType(subName [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.AddNewType(&_BalanceContract.TransactOpts, subName)
}

// DeleteSub is a paid mutator transaction binding the contract method 0xf356cc79.
//
// Solidity: function deleteSub(bytes32 subName) returns()
func (_BalanceContract *BalanceContractTransactor) DeleteSub(opts *bind.TransactOpts, subName [32]byte) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "deleteSub", subName)
}

// DeleteSub is a paid mutator transaction binding the contract method 0xf356cc79.
//
// Solidity: function deleteSub(bytes32 subName) returns()
func (_BalanceContract *BalanceContractSession) DeleteSub(subName [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.DeleteSub(&_BalanceContract.TransactOpts, subName)
}

// DeleteSub is a paid mutator transaction binding the contract method 0xf356cc79.
//
// Solidity: function deleteSub(bytes32 subName) returns()
func (_BalanceContract *BalanceContractTransactorSession) DeleteSub(subName [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.DeleteSub(&_BalanceContract.TransactOpts, subName)
}

// DeleteType is a paid mutator transaction binding the contract method 0x1d9de32c.
//
// Solidity: function deleteType(bytes32 subName) returns()
func (_BalanceContract *BalanceContractTransactor) DeleteType(opts *bind.TransactOpts, subName [32]byte) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "deleteType", subName)
}

// DeleteType is a paid mutator transaction binding the contract method 0x1d9de32c.
//
// Solidity: function deleteType(bytes32 subName) returns()
func (_BalanceContract *BalanceContractSession) DeleteType(subName [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.DeleteType(&_BalanceContract.TransactOpts, subName)
}

// DeleteType is a paid mutator transaction binding the contract method 0x1d9de32c.
//
// Solidity: function deleteType(bytes32 subName) returns()
func (_BalanceContract *BalanceContractTransactorSession) DeleteType(subName [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.DeleteType(&_BalanceContract.TransactOpts, subName)
}

// SubscribeTo is a paid mutator transaction binding the contract method 0xc990824d.
//
// Solidity: function subscribeTo(bytes32 subName, uint256 time) returns()
func (_BalanceContract *BalanceContractTransactor) SubscribeTo(opts *bind.TransactOpts, subName [32]byte, time *big.Int) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "subscribeTo", subName, time)
}

// SubscribeTo is a paid mutator transaction binding the contract method 0xc990824d.
//
// Solidity: function subscribeTo(bytes32 subName, uint256 time) returns()
func (_BalanceContract *BalanceContractSession) SubscribeTo(subName [32]byte, time *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.SubscribeTo(&_BalanceContract.TransactOpts, subName, time)
}

// SubscribeTo is a paid mutator transaction binding the contract method 0xc990824d.
//
// Solidity: function subscribeTo(bytes32 subName, uint256 time) returns()
func (_BalanceContract *BalanceContractTransactorSession) SubscribeTo(subName [32]byte, time *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.SubscribeTo(&_BalanceContract.TransactOpts, subName, time)
}

// BalanceContractNotifyNewIterator is returned from FilterNotifyNew and is used to iterate over the raw logs and unpacked data for NotifyNew events raised by the BalanceContract contract.
type BalanceContractNotifyNewIterator struct {
	Event *BalanceContractNotifyNew // Event containing the contract specifics and raw log

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
func (it *BalanceContractNotifyNewIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractNotifyNew)
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
		it.Event = new(BalanceContractNotifyNew)
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
func (it *BalanceContractNotifyNewIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractNotifyNewIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractNotifyNew represents a NotifyNew event raised by the BalanceContract contract.
type BalanceContractNotifyNew struct {
	Addr    common.Address
	SubID   [32]byte
	EndTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNotifyNew is a free log retrieval operation binding the contract event 0xc7c18d7fd8f2b08630c0d4f0bcf252583ff50a5cf2c5d8811f8a4347191679d4.
//
// Solidity: event notifyNew(address indexed _addr, bytes32 indexed _subID, uint256 _endTime)
func (_BalanceContract *BalanceContractFilterer) FilterNotifyNew(opts *bind.FilterOpts, _addr []common.Address, _subID [][32]byte) (*BalanceContractNotifyNewIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _subIDRule []interface{}
	for _, _subIDItem := range _subID {
		_subIDRule = append(_subIDRule, _subIDItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "notifyNew", _addrRule, _subIDRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractNotifyNewIterator{contract: _BalanceContract.contract, event: "notifyNew", logs: logs, sub: sub}, nil
}

// WatchNotifyNew is a free log subscription operation binding the contract event 0xc7c18d7fd8f2b08630c0d4f0bcf252583ff50a5cf2c5d8811f8a4347191679d4.
//
// Solidity: event notifyNew(address indexed _addr, bytes32 indexed _subID, uint256 _endTime)
func (_BalanceContract *BalanceContractFilterer) WatchNotifyNew(opts *bind.WatchOpts, sink chan<- *BalanceContractNotifyNew, _addr []common.Address, _subID [][32]byte) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _subIDRule []interface{}
	for _, _subIDItem := range _subID {
		_subIDRule = append(_subIDRule, _subIDItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "notifyNew", _addrRule, _subIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractNotifyNew)
				if err := _BalanceContract.contract.UnpackLog(event, "notifyNew", log); err != nil {
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

// ParseNotifyNew is a log parse operation binding the contract event 0xc7c18d7fd8f2b08630c0d4f0bcf252583ff50a5cf2c5d8811f8a4347191679d4.
//
// Solidity: event notifyNew(address indexed _addr, bytes32 indexed _subID, uint256 _endTime)
func (_BalanceContract *BalanceContractFilterer) ParseNotifyNew(log types.Log) (*BalanceContractNotifyNew, error) {
	event := new(BalanceContractNotifyNew)
	if err := _BalanceContract.contract.UnpackLog(event, "notifyNew", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalanceContractNotifyNewCategoryIterator is returned from FilterNotifyNewCategory and is used to iterate over the raw logs and unpacked data for NotifyNewCategory events raised by the BalanceContract contract.
type BalanceContractNotifyNewCategoryIterator struct {
	Event *BalanceContractNotifyNewCategory // Event containing the contract specifics and raw log

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
func (it *BalanceContractNotifyNewCategoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractNotifyNewCategory)
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
		it.Event = new(BalanceContractNotifyNewCategory)
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
func (it *BalanceContractNotifyNewCategoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractNotifyNewCategoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractNotifyNewCategory represents a NotifyNewCategory event raised by the BalanceContract contract.
type BalanceContractNotifyNewCategory struct {
	Addr common.Address
	Name [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNotifyNewCategory is a free log retrieval operation binding the contract event 0xde3cd713f10754ba21eabefe59d58e9251b68f9ff9016fc20324b4e261f1d598.
//
// Solidity: event notifyNewCategory(address indexed _addr, bytes32 indexed _name)
func (_BalanceContract *BalanceContractFilterer) FilterNotifyNewCategory(opts *bind.FilterOpts, _addr []common.Address, _name [][32]byte) (*BalanceContractNotifyNewCategoryIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "notifyNewCategory", _addrRule, _nameRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractNotifyNewCategoryIterator{contract: _BalanceContract.contract, event: "notifyNewCategory", logs: logs, sub: sub}, nil
}

// WatchNotifyNewCategory is a free log subscription operation binding the contract event 0xde3cd713f10754ba21eabefe59d58e9251b68f9ff9016fc20324b4e261f1d598.
//
// Solidity: event notifyNewCategory(address indexed _addr, bytes32 indexed _name)
func (_BalanceContract *BalanceContractFilterer) WatchNotifyNewCategory(opts *bind.WatchOpts, sink chan<- *BalanceContractNotifyNewCategory, _addr []common.Address, _name [][32]byte) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "notifyNewCategory", _addrRule, _nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractNotifyNewCategory)
				if err := _BalanceContract.contract.UnpackLog(event, "notifyNewCategory", log); err != nil {
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

// ParseNotifyNewCategory is a log parse operation binding the contract event 0xde3cd713f10754ba21eabefe59d58e9251b68f9ff9016fc20324b4e261f1d598.
//
// Solidity: event notifyNewCategory(address indexed _addr, bytes32 indexed _name)
func (_BalanceContract *BalanceContractFilterer) ParseNotifyNewCategory(log types.Log) (*BalanceContractNotifyNewCategory, error) {
	event := new(BalanceContractNotifyNewCategory)
	if err := _BalanceContract.contract.UnpackLog(event, "notifyNewCategory", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalanceContractNotifyRemoveIterator is returned from FilterNotifyRemove and is used to iterate over the raw logs and unpacked data for NotifyRemove events raised by the BalanceContract contract.
type BalanceContractNotifyRemoveIterator struct {
	Event *BalanceContractNotifyRemove // Event containing the contract specifics and raw log

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
func (it *BalanceContractNotifyRemoveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractNotifyRemove)
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
		it.Event = new(BalanceContractNotifyRemove)
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
func (it *BalanceContractNotifyRemoveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractNotifyRemoveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractNotifyRemove represents a NotifyRemove event raised by the BalanceContract contract.
type BalanceContractNotifyRemove struct {
	Addr  common.Address
	SubID [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNotifyRemove is a free log retrieval operation binding the contract event 0xab7ba7d234105b15e41e97f548ae644ed346089e33e01f8f03a862bfaebacb91.
//
// Solidity: event notifyRemove(address indexed _addr, bytes32 indexed _subID)
func (_BalanceContract *BalanceContractFilterer) FilterNotifyRemove(opts *bind.FilterOpts, _addr []common.Address, _subID [][32]byte) (*BalanceContractNotifyRemoveIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _subIDRule []interface{}
	for _, _subIDItem := range _subID {
		_subIDRule = append(_subIDRule, _subIDItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "notifyRemove", _addrRule, _subIDRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractNotifyRemoveIterator{contract: _BalanceContract.contract, event: "notifyRemove", logs: logs, sub: sub}, nil
}

// WatchNotifyRemove is a free log subscription operation binding the contract event 0xab7ba7d234105b15e41e97f548ae644ed346089e33e01f8f03a862bfaebacb91.
//
// Solidity: event notifyRemove(address indexed _addr, bytes32 indexed _subID)
func (_BalanceContract *BalanceContractFilterer) WatchNotifyRemove(opts *bind.WatchOpts, sink chan<- *BalanceContractNotifyRemove, _addr []common.Address, _subID [][32]byte) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _subIDRule []interface{}
	for _, _subIDItem := range _subID {
		_subIDRule = append(_subIDRule, _subIDItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "notifyRemove", _addrRule, _subIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractNotifyRemove)
				if err := _BalanceContract.contract.UnpackLog(event, "notifyRemove", log); err != nil {
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

// ParseNotifyRemove is a log parse operation binding the contract event 0xab7ba7d234105b15e41e97f548ae644ed346089e33e01f8f03a862bfaebacb91.
//
// Solidity: event notifyRemove(address indexed _addr, bytes32 indexed _subID)
func (_BalanceContract *BalanceContractFilterer) ParseNotifyRemove(log types.Log) (*BalanceContractNotifyRemove, error) {
	event := new(BalanceContractNotifyRemove)
	if err := _BalanceContract.contract.UnpackLog(event, "notifyRemove", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalanceContractNotifyRemoveCategoryIterator is returned from FilterNotifyRemoveCategory and is used to iterate over the raw logs and unpacked data for NotifyRemoveCategory events raised by the BalanceContract contract.
type BalanceContractNotifyRemoveCategoryIterator struct {
	Event *BalanceContractNotifyRemoveCategory // Event containing the contract specifics and raw log

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
func (it *BalanceContractNotifyRemoveCategoryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractNotifyRemoveCategory)
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
		it.Event = new(BalanceContractNotifyRemoveCategory)
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
func (it *BalanceContractNotifyRemoveCategoryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractNotifyRemoveCategoryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractNotifyRemoveCategory represents a NotifyRemoveCategory event raised by the BalanceContract contract.
type BalanceContractNotifyRemoveCategory struct {
	Addr common.Address
	Name [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterNotifyRemoveCategory is a free log retrieval operation binding the contract event 0xf4d63e41754a74d0227558f224c4d7d0beb73170d73b23db38f9317e650da562.
//
// Solidity: event notifyRemoveCategory(address indexed _addr, bytes32 indexed _name)
func (_BalanceContract *BalanceContractFilterer) FilterNotifyRemoveCategory(opts *bind.FilterOpts, _addr []common.Address, _name [][32]byte) (*BalanceContractNotifyRemoveCategoryIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "notifyRemoveCategory", _addrRule, _nameRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractNotifyRemoveCategoryIterator{contract: _BalanceContract.contract, event: "notifyRemoveCategory", logs: logs, sub: sub}, nil
}

// WatchNotifyRemoveCategory is a free log subscription operation binding the contract event 0xf4d63e41754a74d0227558f224c4d7d0beb73170d73b23db38f9317e650da562.
//
// Solidity: event notifyRemoveCategory(address indexed _addr, bytes32 indexed _name)
func (_BalanceContract *BalanceContractFilterer) WatchNotifyRemoveCategory(opts *bind.WatchOpts, sink chan<- *BalanceContractNotifyRemoveCategory, _addr []common.Address, _name [][32]byte) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _nameRule []interface{}
	for _, _nameItem := range _name {
		_nameRule = append(_nameRule, _nameItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "notifyRemoveCategory", _addrRule, _nameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractNotifyRemoveCategory)
				if err := _BalanceContract.contract.UnpackLog(event, "notifyRemoveCategory", log); err != nil {
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

// ParseNotifyRemoveCategory is a log parse operation binding the contract event 0xf4d63e41754a74d0227558f224c4d7d0beb73170d73b23db38f9317e650da562.
//
// Solidity: event notifyRemoveCategory(address indexed _addr, bytes32 indexed _name)
func (_BalanceContract *BalanceContractFilterer) ParseNotifyRemoveCategory(log types.Log) (*BalanceContractNotifyRemoveCategory, error) {
	event := new(BalanceContractNotifyRemoveCategory)
	if err := _BalanceContract.contract.UnpackLog(event, "notifyRemoveCategory", log); err != nil {
		return nil, err
	}
	return event, nil
}
