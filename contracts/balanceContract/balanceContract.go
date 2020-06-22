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

// AccessControlContractABI is the input ABI used to generate the binding from.
const AccessControlContractABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getPubKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AccessControlContractFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlContractFuncSigs = map[string]string{
	"fce9512a": "getPubKey(address)",
}

// AccessControlContractBin is the compiled bytecode used for deploying new contracts.
var AccessControlContractBin = "0x608060405234801561001057600080fd5b5060fd8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063fce9512a14602d575b600080fd5b605060048036036020811015604157600080fd5b50356001600160a01b031660c2565b6040805160208082528351818301528351919283929083019185019080838360005b8381101560885781810151838201526020016072565b50505050905090810190601f16801560b45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b5060609056fea265627a7a7231582058390e1554fe771fae130e0bdfe8a45ce12cb0a18756162443830d111da7cf7964736f6c63430005100032"

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

// GetPubKey is a free data retrieval call binding the contract method 0xfce9512a.
//
// Solidity: function getPubKey(address ) view returns(string)
func (_AccessControlContract *AccessControlContractCaller) GetPubKey(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AccessControlContract.contract.Call(opts, out, "getPubKey", arg0)
	return *ret0, err
}

// GetPubKey is a free data retrieval call binding the contract method 0xfce9512a.
//
// Solidity: function getPubKey(address ) view returns(string)
func (_AccessControlContract *AccessControlContractSession) GetPubKey(arg0 common.Address) (string, error) {
	return _AccessControlContract.Contract.GetPubKey(&_AccessControlContract.CallOpts, arg0)
}

// GetPubKey is a free data retrieval call binding the contract method 0xfce9512a.
//
// Solidity: function getPubKey(address ) view returns(string)
func (_AccessControlContract *AccessControlContractCallerSession) GetPubKey(arg0 common.Address) (string, error) {
	return _AccessControlContract.Contract.GetPubKey(&_AccessControlContract.CallOpts, arg0)
}

// BalanceContractABI is the input ABI used to generate the binding from.
const BalanceContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"ackPurchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"purchaseNotify\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_txHashOTP\",\"type\":\"bytes32\"}],\"name\":\"responseNotify\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"confirmData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getPriceData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"payData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"txHashOTP\",\"type\":\"bytes32\"}],\"name\":\"sendData\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddress2\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPriceData\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BalanceContractFuncSigs maps the 4-byte function signature to its string representation.
var BalanceContractFuncSigs = map[string]string{
	"2e3ad603": "confirmData(bytes32)",
	"43fa6211": "getPriceData(bytes32)",
	"e70996f2": "payData(bytes32)",
	"4b802d55": "sendData(address,bytes32,bytes32)",
	"e30081a0": "setAddress(address)",
	"13a5b50a": "setAddress2(address)",
	"63c6c2ef": "setPriceData(bytes32,uint256)",
}

// BalanceContractBin is the compiled bytecode used for deploying new contracts.
var BalanceContractBin = "0x6080604052600280546001600160a01b0319167321a018606490c031a8c02bb6b992d8ae44add37f179055600160045534801561003b57600080fd5b50610a9a8061004b6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80634b802d551161005b5780634b802d55146100f657806363c6c2ef14610128578063e30081a01461015f578063e70996f2146101855761007d565b806313a5b50a146100825780632e3ad603146100aa57806343fa6211146100c7575b600080fd5b6100a86004803603602081101561009857600080fd5b50356001600160a01b03166101a2565b005b6100a8600480360360208110156100c057600080fd5b50356101d8565b6100e4600480360360208110156100dd57600080fd5b5035610281565b60408051918252519081900360200190f35b6100a86004803603606081101561010c57600080fd5b506001600160a01b038135169060208101359060400135610293565b61014b6004803603604081101561013e57600080fd5b50803590602001356103b2565b604080519115158252519081900360200190f35b6100a86004803603602081101561017557600080fd5b50356001600160a01b03166105dd565b6100a86004803603602081101561019b57600080fd5b5035610613565b6002546001600160a01b031633146101b657fe5b600180546001600160a01b0319166001600160a01b0392909216919091179055565b33600090815260036020908152604080832084845290915290205460ff61010090910416151560011461023c5760405162461bcd60e51b815260040180806020018281038252602d815260200180610910602d913960400191505060405180910390fd5b336000818152600360205260408082206001018054600019019055518392917fe2e91035b1a52f7ba9a89d081b3b10b6ad4812f301e92a478e17cd9eb5641ea991a350565b60009081526006602052604090205490565b6002546001600160a01b031633146102dc5760405162461bcd60e51b81526004018080602001828103825260308152602001806109d16030913960400191505060405180910390fd5b6001600160a01b038316600090815260036020908152604080832085845290915290205460ff1615156001146103435760405162461bcd60e51b81526004018080602001828103825260458152602001806109656045913960600191505060405180910390fd5b6001600160a01b0383166000818152600360209081526040808320868452825291829020805461010061ff00199091161760ff19169055815184815291518593927fc845b00aa923b3cc1a6bf20401682a98918640788d078704e44ec3f19099439b92908290030190a3505050565b6002546000906001600160a01b031633146103c957fe5b600080546040805163c9776a6d60e01b815260048101879052905160609384936001600160a01b03169263c9776a6d9260248083019392829003018186803b15801561041457600080fd5b505afa158015610428573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604090815281101561045157600080fd5b810190808051604051939291908464010000000082111561047157600080fd5b90830190602082018581111561048657600080fd5b82516401000000008111828201881017156104a057600080fd5b82525081516020918201929091019080838360005b838110156104cd5781810151838201526020016104b5565b50505050905090810190601f1680156104fa5780820380516001836020036101000a031916815260200191505b506040526020018051604051939291908464010000000082111561051d57600080fd5b90830190602082018581111561053257600080fd5b825164010000000081118282018810171561054c57600080fd5b82525081516020918201929091019080838360005b83811015610579578181015183820152602001610561565b50505050905090810190601f1680156105a65780820380516001836020036101000a031916815260200191505b50604052505050915091508051600014156105bd57fe5b81516105c557fe5b50505060009182526006602052604090912055600190565b6002546001600160a01b031633146105f157fe5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b60015460408051637e74a89560e11b815233600482015290516001600160a01b039092169163fce9512a91602480820192600092909190829003018186803b15801561065e57600080fd5b505afa158015610672573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052602081101561069b57600080fd5b81019080805160405193929190846401000000008211156106bb57600080fd5b9083019060208201858111156106d057600080fd5b82516401000000008111828201881017156106ea57600080fd5b82525081516020918201929091019080838360005b838110156107175781810151838201526020016106ff565b50505050905090810190601f1680156107445780820380516001836020036101000a031916815260200191505b50604052505050516000141561078b5760405162461bcd60e51b81526004018080602001828103825260278152602001806109aa6027913960400191505060405180910390fd5b336000908152600360209081526040808320848452909152902054610100900460ff16156107ea5760405162461bcd60e51b815260040180806020018281038252602881526020018061093d6028913960400191505060405180910390fd5b33600090815260036020908152604080832084845290915290205460ff16156108445760405162461bcd60e51b815260040180806020018281038252602f815260200180610a01602f913960400191505060405180910390fd5b60045433600090815260036020526040902060010154106108965760405162461bcd60e51b8152600401808060200182810382526036815260200180610a306036913960400191505060405180910390fd5b336000818152600360208181526040808420868552808352908420805460ff191660019081179091559385905291905281018054909101905581907f58d967c9a9a54b98c710c372cfe363bb9466f443b58def6976bd0e2dc77af8416108fb83610281565b60408051918252519081900360200190a35056fe54686520636c69656e7420686173206e6f7420626f7567687420746865206d6561737572656d656e742079657454686520757365722068617320616c726561647920626f75676874207468697320656c656d656e74546865726520686173206e6f74206265656e20616e79207075726368617365206173736f63696174656420746f2074686973206861736820616e642074686973207573657254686520636c69656e74206d75737420696e64696361746520686973207075626c6963206b65795573657220686173206e6f7420656e6f7567682070726976696c6567657320746f20646f2074686973206f7074696f6e5468652070726f6363657373206f662074686973207075726368617365206973207374696c6c206f6e20676f696e67546865726520617265206d6561737572656d656e7473207468617420686173206e6f74206265656e20636f6e6669726d656420796574a265627a7a72315820c98145b03fc5cef4c6442c8485a36017b62ad77baf31de1a7f9a05c1e68e26e764736f6c63430005100032"

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

// ConfirmData is a paid mutator transaction binding the contract method 0x2e3ad603.
//
// Solidity: function confirmData(bytes32 hash) returns()
func (_BalanceContract *BalanceContractTransactor) ConfirmData(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "confirmData", hash)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x2e3ad603.
//
// Solidity: function confirmData(bytes32 hash) returns()
func (_BalanceContract *BalanceContractSession) ConfirmData(hash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.ConfirmData(&_BalanceContract.TransactOpts, hash)
}

// ConfirmData is a paid mutator transaction binding the contract method 0x2e3ad603.
//
// Solidity: function confirmData(bytes32 hash) returns()
func (_BalanceContract *BalanceContractTransactorSession) ConfirmData(hash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.ConfirmData(&_BalanceContract.TransactOpts, hash)
}

// PayData is a paid mutator transaction binding the contract method 0xe70996f2.
//
// Solidity: function payData(bytes32 hash) returns()
func (_BalanceContract *BalanceContractTransactor) PayData(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "payData", hash)
}

// PayData is a paid mutator transaction binding the contract method 0xe70996f2.
//
// Solidity: function payData(bytes32 hash) returns()
func (_BalanceContract *BalanceContractSession) PayData(hash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.PayData(&_BalanceContract.TransactOpts, hash)
}

// PayData is a paid mutator transaction binding the contract method 0xe70996f2.
//
// Solidity: function payData(bytes32 hash) returns()
func (_BalanceContract *BalanceContractTransactorSession) PayData(hash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.PayData(&_BalanceContract.TransactOpts, hash)
}

// SendData is a paid mutator transaction binding the contract method 0x4b802d55.
//
// Solidity: function sendData(address client, bytes32 hash, bytes32 txHashOTP) returns()
func (_BalanceContract *BalanceContractTransactor) SendData(opts *bind.TransactOpts, client common.Address, hash [32]byte, txHashOTP [32]byte) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "sendData", client, hash, txHashOTP)
}

// SendData is a paid mutator transaction binding the contract method 0x4b802d55.
//
// Solidity: function sendData(address client, bytes32 hash, bytes32 txHashOTP) returns()
func (_BalanceContract *BalanceContractSession) SendData(client common.Address, hash [32]byte, txHashOTP [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.SendData(&_BalanceContract.TransactOpts, client, hash, txHashOTP)
}

// SendData is a paid mutator transaction binding the contract method 0x4b802d55.
//
// Solidity: function sendData(address client, bytes32 hash, bytes32 txHashOTP) returns()
func (_BalanceContract *BalanceContractTransactorSession) SendData(client common.Address, hash [32]byte, txHashOTP [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.SendData(&_BalanceContract.TransactOpts, client, hash, txHashOTP)
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

// SetAddress2 is a paid mutator transaction binding the contract method 0x13a5b50a.
//
// Solidity: function setAddress2(address _address) returns()
func (_BalanceContract *BalanceContractTransactor) SetAddress2(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "setAddress2", _address)
}

// SetAddress2 is a paid mutator transaction binding the contract method 0x13a5b50a.
//
// Solidity: function setAddress2(address _address) returns()
func (_BalanceContract *BalanceContractSession) SetAddress2(_address common.Address) (*types.Transaction, error) {
	return _BalanceContract.Contract.SetAddress2(&_BalanceContract.TransactOpts, _address)
}

// SetAddress2 is a paid mutator transaction binding the contract method 0x13a5b50a.
//
// Solidity: function setAddress2(address _address) returns()
func (_BalanceContract *BalanceContractTransactorSession) SetAddress2(_address common.Address) (*types.Transaction, error) {
	return _BalanceContract.Contract.SetAddress2(&_BalanceContract.TransactOpts, _address)
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

// BalanceContractAckPurchaseIterator is returned from FilterAckPurchase and is used to iterate over the raw logs and unpacked data for AckPurchase events raised by the BalanceContract contract.
type BalanceContractAckPurchaseIterator struct {
	Event *BalanceContractAckPurchase // Event containing the contract specifics and raw log

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
func (it *BalanceContractAckPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractAckPurchase)
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
		it.Event = new(BalanceContractAckPurchase)
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
func (it *BalanceContractAckPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractAckPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractAckPurchase represents a AckPurchase event raised by the BalanceContract contract.
type BalanceContractAckPurchase struct {
	Addr common.Address
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAckPurchase is a free log retrieval operation binding the contract event 0xe2e91035b1a52f7ba9a89d081b3b10b6ad4812f301e92a478e17cd9eb5641ea9.
//
// Solidity: event ackPurchase(address indexed _addr, bytes32 indexed _hash)
func (_BalanceContract *BalanceContractFilterer) FilterAckPurchase(opts *bind.FilterOpts, _addr []common.Address, _hash [][32]byte) (*BalanceContractAckPurchaseIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "ackPurchase", _addrRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractAckPurchaseIterator{contract: _BalanceContract.contract, event: "ackPurchase", logs: logs, sub: sub}, nil
}

// WatchAckPurchase is a free log subscription operation binding the contract event 0xe2e91035b1a52f7ba9a89d081b3b10b6ad4812f301e92a478e17cd9eb5641ea9.
//
// Solidity: event ackPurchase(address indexed _addr, bytes32 indexed _hash)
func (_BalanceContract *BalanceContractFilterer) WatchAckPurchase(opts *bind.WatchOpts, sink chan<- *BalanceContractAckPurchase, _addr []common.Address, _hash [][32]byte) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "ackPurchase", _addrRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractAckPurchase)
				if err := _BalanceContract.contract.UnpackLog(event, "ackPurchase", log); err != nil {
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

// ParseAckPurchase is a log parse operation binding the contract event 0xe2e91035b1a52f7ba9a89d081b3b10b6ad4812f301e92a478e17cd9eb5641ea9.
//
// Solidity: event ackPurchase(address indexed _addr, bytes32 indexed _hash)
func (_BalanceContract *BalanceContractFilterer) ParseAckPurchase(log types.Log) (*BalanceContractAckPurchase, error) {
	event := new(BalanceContractAckPurchase)
	if err := _BalanceContract.contract.UnpackLog(event, "ackPurchase", log); err != nil {
		return nil, err
	}
	return event, nil
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
	Addr      common.Address
	Hash      [32]byte
	TxHashOTP [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResponseNotify is a free log retrieval operation binding the contract event 0xc845b00aa923b3cc1a6bf20401682a98918640788d078704e44ec3f19099439b.
//
// Solidity: event responseNotify(address indexed _addr, bytes32 indexed _hash, bytes32 _txHashOTP)
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

// WatchResponseNotify is a free log subscription operation binding the contract event 0xc845b00aa923b3cc1a6bf20401682a98918640788d078704e44ec3f19099439b.
//
// Solidity: event responseNotify(address indexed _addr, bytes32 indexed _hash, bytes32 _txHashOTP)
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

// ParseResponseNotify is a log parse operation binding the contract event 0xc845b00aa923b3cc1a6bf20401682a98918640788d078704e44ec3f19099439b.
//
// Solidity: event responseNotify(address indexed _addr, bytes32 indexed _hash, bytes32 _txHashOTP)
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
var DataLedgerContractBin = "0x608060405234801561001057600080fd5b50610168806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063c9776a6d14610030575b600080fd5b61004d6004803603602081101561004657600080fd5b503561012b565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561008e578181015183820152602001610076565b50505050905090810190601f1680156100bb5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156100ee5781810151838201526020016100d6565b50505050905090810190601f16801561011b5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b60608091509156fea265627a7a723158207a3246663b5862cea3b240c55e964eb4d5b65fd8dbc98f99f7b709995122880964736f6c63430005100032"

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
