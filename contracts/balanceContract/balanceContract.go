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
const BalanceContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"purchaseNotify\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_txHashExchange\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_txHashData\",\"type\":\"bytes32\"}],\"name\":\"responseNotify\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"clientAccount\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"checkHasPaid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getPriceData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"payData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"clientAccount\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txHashExchange\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txHashData\",\"type\":\"bytes32\"}],\"name\":\"sendToClient\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPriceData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BalanceContractFuncSigs maps the 4-byte function signature to its string representation.
var BalanceContractFuncSigs = map[string]string{
	"c71daccb": "checkBalance()",
	"633d4c31": "checkHasPaid(address,bytes32)",
	"43fa6211": "getPriceData(bytes32)",
	"30d87a57": "payData(bytes32,uint256)",
	"4e635503": "sendToClient(address,bytes32,bytes32,bytes32)",
	"e30081a0": "setAddress(address)",
	"63c6c2ef": "setPriceData(bytes32,uint256)",
}

// BalanceContractBin is the compiled bytecode used for deploying new contracts.
var BalanceContractBin = "0x6080604052600480546001600160a01b0319167321a018606490c031a8c02bb6b992d8ae44add37f17905534801561003657600080fd5b506107cd806100466000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063633d4c311161005b578063633d4c311461010e57806363c6c2ef1461014e578063c71daccb14610171578063e30081a0146101795761007d565b806330d87a571461008257806343fa6211146100a75780634e635503146100d6575b600080fd5b6100a56004803603604081101561009857600080fd5b508035906020013561019f565b005b6100c4600480360360208110156100bd57600080fd5b5035610462565b60408051918252519081900360200190f35b6100a5600480360360808110156100ec57600080fd5b506001600160a01b038135169060208101359060408101359060600135610474565b61013a6004803603604081101561012457600080fd5b506001600160a01b0381351690602001356104f2565b604080519115158252519081900360200190f35b61013a6004803603604081101561016457600080fd5b508035906020013561051d565b6100c4610744565b6100a56004803603602081101561018f57600080fd5b50356001600160a01b0316610762565b6000828152600360205260409020548110156101b757fe5b3360009081526001602081815260408084208685529091529091205460ff16151514156101e057fe5b600080546040805163c9776a6d60e01b815260048101869052905160609384936001600160a01b03169263c9776a6d9260248083019392829003018186803b15801561022b57600080fd5b505afa15801561023f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604090815281101561026857600080fd5b8101908080516040519392919084600160201b82111561028757600080fd5b90830190602082018581111561029c57600080fd5b8251600160201b8111828201881017156102b557600080fd5b82525081516020918201929091019080838360005b838110156102e25781810151838201526020016102ca565b50505050905090810190601f16801561030f5780820380516001836020036101000a031916815260200191505b5060405260200180516040519392919084600160201b82111561033157600080fd5b90830190602082018581111561034657600080fd5b8251600160201b81118282018810171561035f57600080fd5b82525081516020918201929091019080838360005b8381101561038c578181015183820152602001610374565b50505050905090810190601f1680156103b95780820380516001836020036101000a031916815260200191505b50604052505050915091508051600014156103d057fe5b81516103d857fe5b6103e184610462565b6103e757fe5b600084815260036020818152604080842080546002805490910190553380865260018085528387208b8852855295839020805460ff19169096179095559282529154825190815291518793927f58d967c9a9a54b98c710c372cfe363bb9466f443b58def6976bd0e2dc77af84192908290030190a350505050565b60009081526003602052604090205490565b6004546001600160a01b0316331461048857fe5b6001600160a01b0384166000818152600160209081526040808320878452825291829020805460ff19169055815185815290810184905281518693927f5561f506afbfc9a46c0e3398234971c86c54a89b00fca057c93dde013fb63198928290030190a350505050565b6001600160a01b03919091166000908152600160209081526040808320938352929052205460ff1690565b6004546000906001600160a01b0316331461053457fe5b600080546040805163c9776a6d60e01b815260048101879052905160609384936001600160a01b03169263c9776a6d9260248083019392829003018186803b15801561057f57600080fd5b505afa158015610593573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160409081528110156105bc57600080fd5b8101908080516040519392919084600160201b8211156105db57600080fd5b9083019060208201858111156105f057600080fd5b8251600160201b81118282018810171561060957600080fd5b82525081516020918201929091019080838360005b8381101561063657818101518382015260200161061e565b50505050905090810190601f1680156106635780820380516001836020036101000a031916815260200191505b5060405260200180516040519392919084600160201b82111561068557600080fd5b90830190602082018581111561069a57600080fd5b8251600160201b8111828201881017156106b357600080fd5b82525081516020918201929091019080838360005b838110156106e05781810151838201526020016106c8565b50505050905090810190601f16801561070d5780820380516001836020036101000a031916815260200191505b506040525050509150915080516000141561072457fe5b815161072c57fe5b50505060009182526003602052604090912055600190565b6004546000906001600160a01b0316331461075b57fe5b5060025490565b6004546001600160a01b0316331461077657fe5b600080546001600160a01b0319166001600160a01b039290921691909117905556fea265627a7a72315820065633782d29618fd7c50731c6ca3015bb189cf0d867c6d995de3946ef3c38e864736f6c63430005100032"

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

// CheckBalance is a free data retrieval call binding the contract method 0xc71daccb.
//
// Solidity: function checkBalance() view returns(uint256)
func (_BalanceContract *BalanceContractCaller) CheckBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalanceContract.contract.Call(opts, out, "checkBalance")
	return *ret0, err
}

// CheckBalance is a free data retrieval call binding the contract method 0xc71daccb.
//
// Solidity: function checkBalance() view returns(uint256)
func (_BalanceContract *BalanceContractSession) CheckBalance() (*big.Int, error) {
	return _BalanceContract.Contract.CheckBalance(&_BalanceContract.CallOpts)
}

// CheckBalance is a free data retrieval call binding the contract method 0xc71daccb.
//
// Solidity: function checkBalance() view returns(uint256)
func (_BalanceContract *BalanceContractCallerSession) CheckBalance() (*big.Int, error) {
	return _BalanceContract.Contract.CheckBalance(&_BalanceContract.CallOpts)
}

// CheckHasPaid is a free data retrieval call binding the contract method 0x633d4c31.
//
// Solidity: function checkHasPaid(address clientAccount, bytes32 hash) view returns(bool)
func (_BalanceContract *BalanceContractCaller) CheckHasPaid(opts *bind.CallOpts, clientAccount common.Address, hash [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BalanceContract.contract.Call(opts, out, "checkHasPaid", clientAccount, hash)
	return *ret0, err
}

// CheckHasPaid is a free data retrieval call binding the contract method 0x633d4c31.
//
// Solidity: function checkHasPaid(address clientAccount, bytes32 hash) view returns(bool)
func (_BalanceContract *BalanceContractSession) CheckHasPaid(clientAccount common.Address, hash [32]byte) (bool, error) {
	return _BalanceContract.Contract.CheckHasPaid(&_BalanceContract.CallOpts, clientAccount, hash)
}

// CheckHasPaid is a free data retrieval call binding the contract method 0x633d4c31.
//
// Solidity: function checkHasPaid(address clientAccount, bytes32 hash) view returns(bool)
func (_BalanceContract *BalanceContractCallerSession) CheckHasPaid(clientAccount common.Address, hash [32]byte) (bool, error) {
	return _BalanceContract.Contract.CheckHasPaid(&_BalanceContract.CallOpts, clientAccount, hash)
}

// GetPriceData is a free data retrieval call binding the contract method 0x43fa6211.
//
// Solidity: function getPriceData(bytes32 hash) view returns(uint256)
func (_BalanceContract *BalanceContractCaller) GetPriceData(opts *bind.CallOpts, hash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalanceContract.contract.Call(opts, out, "getPriceData", hash)
	return *ret0, err
}

// GetPriceData is a free data retrieval call binding the contract method 0x43fa6211.
//
// Solidity: function getPriceData(bytes32 hash) view returns(uint256)
func (_BalanceContract *BalanceContractSession) GetPriceData(hash [32]byte) (*big.Int, error) {
	return _BalanceContract.Contract.GetPriceData(&_BalanceContract.CallOpts, hash)
}

// GetPriceData is a free data retrieval call binding the contract method 0x43fa6211.
//
// Solidity: function getPriceData(bytes32 hash) view returns(uint256)
func (_BalanceContract *BalanceContractCallerSession) GetPriceData(hash [32]byte) (*big.Int, error) {
	return _BalanceContract.Contract.GetPriceData(&_BalanceContract.CallOpts, hash)
}

// PayData is a paid mutator transaction binding the contract method 0x30d87a57.
//
// Solidity: function payData(bytes32 hash, uint256 tokens) returns()
func (_BalanceContract *BalanceContractTransactor) PayData(opts *bind.TransactOpts, hash [32]byte, tokens *big.Int) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "payData", hash, tokens)
}

// PayData is a paid mutator transaction binding the contract method 0x30d87a57.
//
// Solidity: function payData(bytes32 hash, uint256 tokens) returns()
func (_BalanceContract *BalanceContractSession) PayData(hash [32]byte, tokens *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.PayData(&_BalanceContract.TransactOpts, hash, tokens)
}

// PayData is a paid mutator transaction binding the contract method 0x30d87a57.
//
// Solidity: function payData(bytes32 hash, uint256 tokens) returns()
func (_BalanceContract *BalanceContractTransactorSession) PayData(hash [32]byte, tokens *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.PayData(&_BalanceContract.TransactOpts, hash, tokens)
}

// SendToClient is a paid mutator transaction binding the contract method 0x4e635503.
//
// Solidity: function sendToClient(address clientAccount, bytes32 hash, bytes32 txHashExchange, bytes32 txHashData) returns()
func (_BalanceContract *BalanceContractTransactor) SendToClient(opts *bind.TransactOpts, clientAccount common.Address, hash [32]byte, txHashExchange [32]byte, txHashData [32]byte) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "sendToClient", clientAccount, hash, txHashExchange, txHashData)
}

// SendToClient is a paid mutator transaction binding the contract method 0x4e635503.
//
// Solidity: function sendToClient(address clientAccount, bytes32 hash, bytes32 txHashExchange, bytes32 txHashData) returns()
func (_BalanceContract *BalanceContractSession) SendToClient(clientAccount common.Address, hash [32]byte, txHashExchange [32]byte, txHashData [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.SendToClient(&_BalanceContract.TransactOpts, clientAccount, hash, txHashExchange, txHashData)
}

// SendToClient is a paid mutator transaction binding the contract method 0x4e635503.
//
// Solidity: function sendToClient(address clientAccount, bytes32 hash, bytes32 txHashExchange, bytes32 txHashData) returns()
func (_BalanceContract *BalanceContractTransactorSession) SendToClient(clientAccount common.Address, hash [32]byte, txHashExchange [32]byte, txHashData [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.SendToClient(&_BalanceContract.TransactOpts, clientAccount, hash, txHashExchange, txHashData)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_BalanceContract *BalanceContractTransactor) SetAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "setAddress", _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_BalanceContract *BalanceContractSession) SetAddress(_address common.Address) (*types.Transaction, error) {
	return _BalanceContract.Contract.SetAddress(&_BalanceContract.TransactOpts, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_BalanceContract *BalanceContractTransactorSession) SetAddress(_address common.Address) (*types.Transaction, error) {
	return _BalanceContract.Contract.SetAddress(&_BalanceContract.TransactOpts, _address)
}

// SetPriceData is a paid mutator transaction binding the contract method 0x63c6c2ef.
//
// Solidity: function setPriceData(bytes32 hash, uint256 price) returns(bool)
func (_BalanceContract *BalanceContractTransactor) SetPriceData(opts *bind.TransactOpts, hash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "setPriceData", hash, price)
}

// SetPriceData is a paid mutator transaction binding the contract method 0x63c6c2ef.
//
// Solidity: function setPriceData(bytes32 hash, uint256 price) returns(bool)
func (_BalanceContract *BalanceContractSession) SetPriceData(hash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.SetPriceData(&_BalanceContract.TransactOpts, hash, price)
}

// SetPriceData is a paid mutator transaction binding the contract method 0x63c6c2ef.
//
// Solidity: function setPriceData(bytes32 hash, uint256 price) returns(bool)
func (_BalanceContract *BalanceContractTransactorSession) SetPriceData(hash [32]byte, price *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.SetPriceData(&_BalanceContract.TransactOpts, hash, price)
}

// BalanceContractPurchaseNotifyIterator is returned from FilterPurchaseNotify and is used to iterate over the raw logs and unpacked data for PurchaseNotify events raised by the BalanceContract contract.
type BalanceContractPurchaseNotifyIterator struct {
	Event *BalanceContractPurchaseNotify // Event containing the contract specifics and raw log

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
func (it *BalanceContractPurchaseNotifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractPurchaseNotify)
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
		it.Event = new(BalanceContractPurchaseNotify)
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
func (it *BalanceContractPurchaseNotifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractPurchaseNotifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractPurchaseNotify represents a PurchaseNotify event raised by the BalanceContract contract.
type BalanceContractPurchaseNotify struct {
	Addr  common.Address
	Hash  [32]byte
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPurchaseNotify is a free log retrieval operation binding the contract event 0x58d967c9a9a54b98c710c372cfe363bb9466f443b58def6976bd0e2dc77af841.
//
// Solidity: event purchaseNotify(address indexed _addr, bytes32 indexed _hash, uint256 _value)
func (_BalanceContract *BalanceContractFilterer) FilterPurchaseNotify(opts *bind.FilterOpts, _addr []common.Address, _hash [][32]byte) (*BalanceContractPurchaseNotifyIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "purchaseNotify", _addrRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractPurchaseNotifyIterator{contract: _BalanceContract.contract, event: "purchaseNotify", logs: logs, sub: sub}, nil
}

// WatchPurchaseNotify is a free log subscription operation binding the contract event 0x58d967c9a9a54b98c710c372cfe363bb9466f443b58def6976bd0e2dc77af841.
//
// Solidity: event purchaseNotify(address indexed _addr, bytes32 indexed _hash, uint256 _value)
func (_BalanceContract *BalanceContractFilterer) WatchPurchaseNotify(opts *bind.WatchOpts, sink chan<- *BalanceContractPurchaseNotify, _addr []common.Address, _hash [][32]byte) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "purchaseNotify", _addrRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractPurchaseNotify)
				if err := _BalanceContract.contract.UnpackLog(event, "purchaseNotify", log); err != nil {
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

// ParsePurchaseNotify is a log parse operation binding the contract event 0x58d967c9a9a54b98c710c372cfe363bb9466f443b58def6976bd0e2dc77af841.
//
// Solidity: event purchaseNotify(address indexed _addr, bytes32 indexed _hash, uint256 _value)
func (_BalanceContract *BalanceContractFilterer) ParsePurchaseNotify(log types.Log) (*BalanceContractPurchaseNotify, error) {
	event := new(BalanceContractPurchaseNotify)
	if err := _BalanceContract.contract.UnpackLog(event, "purchaseNotify", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalanceContractResponseNotifyIterator is returned from FilterResponseNotify and is used to iterate over the raw logs and unpacked data for ResponseNotify events raised by the BalanceContract contract.
type BalanceContractResponseNotifyIterator struct {
	Event *BalanceContractResponseNotify // Event containing the contract specifics and raw log

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
func (it *BalanceContractResponseNotifyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractResponseNotify)
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
		it.Event = new(BalanceContractResponseNotify)
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
func (it *BalanceContractResponseNotifyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractResponseNotifyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractResponseNotify represents a ResponseNotify event raised by the BalanceContract contract.
type BalanceContractResponseNotify struct {
	Addr           common.Address
	Hash           [32]byte
	TxHashExchange [32]byte
	TxHashData     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterResponseNotify is a free log retrieval operation binding the contract event 0x5561f506afbfc9a46c0e3398234971c86c54a89b00fca057c93dde013fb63198.
//
// Solidity: event responseNotify(address indexed _addr, bytes32 indexed _hash, bytes32 _txHashExchange, bytes32 _txHashData)
func (_BalanceContract *BalanceContractFilterer) FilterResponseNotify(opts *bind.FilterOpts, _addr []common.Address, _hash [][32]byte) (*BalanceContractResponseNotifyIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "responseNotify", _addrRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractResponseNotifyIterator{contract: _BalanceContract.contract, event: "responseNotify", logs: logs, sub: sub}, nil
}

// WatchResponseNotify is a free log subscription operation binding the contract event 0x5561f506afbfc9a46c0e3398234971c86c54a89b00fca057c93dde013fb63198.
//
// Solidity: event responseNotify(address indexed _addr, bytes32 indexed _hash, bytes32 _txHashExchange, bytes32 _txHashData)
func (_BalanceContract *BalanceContractFilterer) WatchResponseNotify(opts *bind.WatchOpts, sink chan<- *BalanceContractResponseNotify, _addr []common.Address, _hash [][32]byte) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "responseNotify", _addrRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractResponseNotify)
				if err := _BalanceContract.contract.UnpackLog(event, "responseNotify", log); err != nil {
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

// ParseResponseNotify is a log parse operation binding the contract event 0x5561f506afbfc9a46c0e3398234971c86c54a89b00fca057c93dde013fb63198.
//
// Solidity: event responseNotify(address indexed _addr, bytes32 indexed _hash, bytes32 _txHashExchange, bytes32 _txHashData)
func (_BalanceContract *BalanceContractFilterer) ParseResponseNotify(log types.Log) (*BalanceContractResponseNotify, error) {
	event := new(BalanceContractResponseNotify)
	if err := _BalanceContract.contract.UnpackLog(event, "responseNotify", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DataLedgerContractABI is the input ABI used to generate the binding from.
const DataLedgerContractABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"retrieveInfo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DataLedgerContractFuncSigs maps the 4-byte function signature to its string representation.
var DataLedgerContractFuncSigs = map[string]string{
	"c9776a6d": "retrieveInfo(bytes32)",
}

// DataLedgerContractBin is the compiled bytecode used for deploying new contracts.
var DataLedgerContractBin = "0x608060405234801561001057600080fd5b50610168806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063c9776a6d14610030575b600080fd5b61004d6004803603602081101561004657600080fd5b503561012b565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561008e578181015183820152602001610076565b50505050905090810190601f1680156100bb5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156100ee5781810151838201526020016100d6565b50505050905090810190601f16801561011b5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b60608091509156fea265627a7a723158200c87d5552b453564d7bb2db13f81e5034877709748a57138c3a2b3c19556e55c64736f6c63430005100032"

// DeployDataLedgerContract deploys a new Ethereum contract, binding an instance of DataLedgerContract to it.
func DeployDataLedgerContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataLedgerContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataLedgerContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DataLedgerContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataLedgerContract{DataLedgerContractCaller: DataLedgerContractCaller{contract: contract}, DataLedgerContractTransactor: DataLedgerContractTransactor{contract: contract}, DataLedgerContractFilterer: DataLedgerContractFilterer{contract: contract}}, nil
}

// DataLedgerContract is an auto generated Go binding around an Ethereum contract.
type DataLedgerContract struct {
	DataLedgerContractCaller     // Read-only binding to the contract
	DataLedgerContractTransactor // Write-only binding to the contract
	DataLedgerContractFilterer   // Log filterer for contract events
}

// DataLedgerContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataLedgerContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLedgerContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataLedgerContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLedgerContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataLedgerContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataLedgerContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataLedgerContractSession struct {
	Contract     *DataLedgerContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DataLedgerContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataLedgerContractCallerSession struct {
	Contract *DataLedgerContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DataLedgerContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataLedgerContractTransactorSession struct {
	Contract     *DataLedgerContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DataLedgerContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataLedgerContractRaw struct {
	Contract *DataLedgerContract // Generic contract binding to access the raw methods on
}

// DataLedgerContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataLedgerContractCallerRaw struct {
	Contract *DataLedgerContractCaller // Generic read-only contract binding to access the raw methods on
}

// DataLedgerContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataLedgerContractTransactorRaw struct {
	Contract *DataLedgerContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataLedgerContract creates a new instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContract(address common.Address, backend bind.ContractBackend) (*DataLedgerContract, error) {
	contract, err := bindDataLedgerContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContract{DataLedgerContractCaller: DataLedgerContractCaller{contract: contract}, DataLedgerContractTransactor: DataLedgerContractTransactor{contract: contract}, DataLedgerContractFilterer: DataLedgerContractFilterer{contract: contract}}, nil
}

// NewDataLedgerContractCaller creates a new read-only instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContractCaller(address common.Address, caller bind.ContractCaller) (*DataLedgerContractCaller, error) {
	contract, err := bindDataLedgerContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractCaller{contract: contract}, nil
}

// NewDataLedgerContractTransactor creates a new write-only instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DataLedgerContractTransactor, error) {
	contract, err := bindDataLedgerContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractTransactor{contract: contract}, nil
}

// NewDataLedgerContractFilterer creates a new log filterer instance of DataLedgerContract, bound to a specific deployed contract.
func NewDataLedgerContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DataLedgerContractFilterer, error) {
	contract, err := bindDataLedgerContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractFilterer{contract: contract}, nil
}

// bindDataLedgerContract binds a generic wrapper to an already deployed contract.
func bindDataLedgerContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataLedgerContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataLedgerContract *DataLedgerContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DataLedgerContract.Contract.DataLedgerContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataLedgerContract *DataLedgerContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DataLedgerContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataLedgerContract *DataLedgerContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DataLedgerContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataLedgerContract *DataLedgerContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DataLedgerContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataLedgerContract *DataLedgerContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataLedgerContract *DataLedgerContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.contract.Transact(opts, method, params...)
}

// RetrieveInfo is a free data retrieval call binding the contract method 0xc9776a6d.
//
// Solidity: function retrieveInfo(bytes32 ) view returns(string, string)
func (_DataLedgerContract *DataLedgerContractCaller) RetrieveInfo(opts *bind.CallOpts, arg0 [32]byte) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DataLedgerContract.contract.Call(opts, out, "retrieveInfo", arg0)
	return *ret0, *ret1, err
}

// RetrieveInfo is a free data retrieval call binding the contract method 0xc9776a6d.
//
// Solidity: function retrieveInfo(bytes32 ) view returns(string, string)
func (_DataLedgerContract *DataLedgerContractSession) RetrieveInfo(arg0 [32]byte) (string, string, error) {
	return _DataLedgerContract.Contract.RetrieveInfo(&_DataLedgerContract.CallOpts, arg0)
}

// RetrieveInfo is a free data retrieval call binding the contract method 0xc9776a6d.
//
// Solidity: function retrieveInfo(bytes32 ) view returns(string, string)
func (_DataLedgerContract *DataLedgerContractCallerSession) RetrieveInfo(arg0 [32]byte) (string, string, error) {
	return _DataLedgerContract.Contract.RetrieveInfo(&_DataLedgerContract.CallOpts, arg0)
}
