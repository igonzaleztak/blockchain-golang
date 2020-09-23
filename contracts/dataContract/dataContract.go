// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dataContract

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
const AccessControlContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AccessControlContractFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlContractFuncSigs = map[string]string{
	"21f8a721": "getAddress(bytes32)",
}

// AccessControlContractBin is the compiled bytecode used for deploying new contracts.
var AccessControlContractBin = "0x608060405234801561001057600080fd5b5060b38061001f6000396000f300608060405260043610603e5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166321f8a72181146043575b600080fd5b348015604e57600080fd5b5060586004356081565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b506000905600a165627a7a7230582053181999bac8a7a732f1aa3baf3d5be6b39eeafa5edbd8173c9f70af488a2a240029"

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

// GetAddress is a paid mutator transaction binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 ) returns(address)
func (_AccessControlContract *AccessControlContractTransactor) GetAddress(opts *bind.TransactOpts, arg0 [32]byte) (*types.Transaction, error) {
	return _AccessControlContract.contract.Transact(opts, "getAddress", arg0)
}

// GetAddress is a paid mutator transaction binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 ) returns(address)
func (_AccessControlContract *AccessControlContractSession) GetAddress(arg0 [32]byte) (*types.Transaction, error) {
	return _AccessControlContract.Contract.GetAddress(&_AccessControlContract.TransactOpts, arg0)
}

// GetAddress is a paid mutator transaction binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 ) returns(address)
func (_AccessControlContract *AccessControlContractTransactorSession) GetAddress(arg0 [32]byte) (*types.Transaction, error) {
	return _AccessControlContract.Contract.GetAddress(&_AccessControlContract.TransactOpts, arg0)
}

// DataLedgerContractABI is the input ABI used to generate the binding from.
const DataLedgerContractABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"deleteMeasurement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"uri\",\"type\":\"string\"},{\"name\":\"id\",\"type\":\"bytes32\"},{\"name\":\"description\",\"type\":\"string\"}],\"name\":\"storeInfo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"retrieveInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"greet\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_uri\",\"type\":\"string\"}],\"name\":\"evtStoreInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"deleteInfo\",\"type\":\"event\"}]"

// DataLedgerContractFuncSigs maps the 4-byte function signature to its string representation.
var DataLedgerContractFuncSigs = map[string]string{
	"77ad95ca": "deleteMeasurement(bytes32)",
	"21f8a721": "getAddress(bytes32)",
	"cfae3217": "greet()",
	"c9776a6d": "retrieveInfo(bytes32)",
	"e30081a0": "setAddress(address)",
	"b45a85a9": "storeInfo(bytes32,string,bytes32,string)",
}

// DataLedgerContractBin is the compiled bytecode used for deploying new contracts.
var DataLedgerContractBin = "0x608060405260018054600160a060020a0319167321a018606490c031a8c02bb6b992d8ae44add37f17905534801561003657600080fd5b506109c1806100466000396000f3006080604052600436106100775763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166321f8a721811461007c57806377ad95ca146100b0578063b45a85a9146100ca578063c9776a6d1461016c578063cfae321714610262578063e30081a0146102ec575b600080fd5b34801561008857600080fd5b5061009460043561030d565b60408051600160a060020a039092168252519081900360200190f35b3480156100bc57600080fd5b506100c86004356103a6565b005b3480156100d657600080fd5b5060408051602060046024803582810135601f81018590048502860185019096528585526100c895833595369560449491939091019190819084018382808284375050604080516020601f818a01358b0180359182018390048302840183018552818452989b8a359b909a90999401975091955091820193509150819084018382808284375094975061049c9650505050505050565b34801561017857600080fd5b50610184600435610632565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b838110156101c55781810151838201526020016101ad565b50505050905090810190601f1680156101f25780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561022557818101518382015260200161020d565b50505050905090810190601f1680156102525780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b34801561026e57600080fd5b50610277610771565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102b1578181015183820152602001610299565b50505050905090810190601f1680156102de5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102f857600080fd5b506100c8600160a060020a03600435166107d1565b60008054604080517f21f8a721000000000000000000000000000000000000000000000000000000008152600481018590529051600160a060020a03909216916321f8a7219160248082019260209290919082900301818787803b15801561037457600080fd5b505af1158015610388573d6000803e3d6000fd5b505050506040513d602081101561039e57600080fd5b505192915050565b600154600160a060020a0316331461044557604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603360248201527f596f7520646f206e6f74206861766520656e6f7567682070726976696c65676560448201527f7320746f20646f207468697320616374696f6e00000000000000000000000000606482015290519081900360840190fd5b60008181526002602052604081209061045e828261089f565b61046c60018301600061089f565b505060405181907f072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe90600090a250565b6104a46108e6565b6104ad8361030d565b600160a060020a0316331461054957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f546865204944207468617420796f7520617265207573696e67206973206e6f7460448201527f2072656769737465726564000000000000000000000000000000000000000000606482015290519081900360840190fd5b8381526020808201839052600086815260028252604090208251805184936105759284929101906108fd565b50602082810151805161058e92600185019201906108fd565b505060408051602080825287518183015287518994507f06c4f2b8beb126621a4c74187a1573e7f48017b16d171df264507e2131b20f42938993928392918301919085019080838360005b838110156105f15781810151838201526020016105d9565b50505050905090810190601f16801561061e5780820380516001836020036101000a031916815260200191505b509250505060405180910390a25050505050565b60008181526002602081815260409283902080548451600019600180841615610100029190910190921694909404601f81018490048402850184019095528484526060948594929391840192918491908301828280156106d35780601f106106a8576101008083540402835291602001916106d3565b820191906000526020600020905b8154815290600101906020018083116106b657829003601f168201915b5050845460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152959750869450925084019050828280156107615780601f1061073657610100808354040283529160200191610761565b820191906000526020600020905b81548152906001019060200180831161074457829003601f168201915b5050505050905091509150915091565b60408051606081018252603181527f48656c6c6f20796f7520686176652063616c6c65642074686520636f6e74726160208201527f637420646174614c65646765722e736f6c000000000000000000000000000000918101919091525b90565b600154600160a060020a0316331461087057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f596f7520646f206e6f7420686176652070726976696c6567657320746f20646f60448201527f207468697320616374696f6e0000000000000000000000000000000000000000606482015290519081900360840190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b50805460018160011615610100020316600290046000825580601f106108c557506108e3565b601f0160209004906000526020600020908101906108e3919061097b565b50565b604080518082019091526060808252602082015290565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061093e57805160ff191683800117855561096b565b8280016001018555821561096b579182015b8281111561096b578251825591602001919060010190610950565b5061097792915061097b565b5090565b6107ce91905b8082111561097757600081556001016109815600a165627a7a72305820ef157ceb3abaea8864c5841a44250d219ca7108556b0f4fc1338326222cf07ca0029"

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

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() pure returns(string)
func (_DataLedgerContract *DataLedgerContractCaller) Greet(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DataLedgerContract.contract.Call(opts, out, "greet")
	return *ret0, err
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() pure returns(string)
func (_DataLedgerContract *DataLedgerContractSession) Greet() (string, error) {
	return _DataLedgerContract.Contract.Greet(&_DataLedgerContract.CallOpts)
}

// Greet is a free data retrieval call binding the contract method 0xcfae3217.
//
// Solidity: function greet() pure returns(string)
func (_DataLedgerContract *DataLedgerContractCallerSession) Greet() (string, error) {
	return _DataLedgerContract.Contract.Greet(&_DataLedgerContract.CallOpts)
}

// RetrieveInfo is a free data retrieval call binding the contract method 0xc9776a6d.
//
// Solidity: function retrieveInfo(bytes32 hash) view returns(string, string)
func (_DataLedgerContract *DataLedgerContractCaller) RetrieveInfo(opts *bind.CallOpts, hash [32]byte) (string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _DataLedgerContract.contract.Call(opts, out, "retrieveInfo", hash)
	return *ret0, *ret1, err
}

// RetrieveInfo is a free data retrieval call binding the contract method 0xc9776a6d.
//
// Solidity: function retrieveInfo(bytes32 hash) view returns(string, string)
func (_DataLedgerContract *DataLedgerContractSession) RetrieveInfo(hash [32]byte) (string, string, error) {
	return _DataLedgerContract.Contract.RetrieveInfo(&_DataLedgerContract.CallOpts, hash)
}

// RetrieveInfo is a free data retrieval call binding the contract method 0xc9776a6d.
//
// Solidity: function retrieveInfo(bytes32 hash) view returns(string, string)
func (_DataLedgerContract *DataLedgerContractCallerSession) RetrieveInfo(hash [32]byte) (string, string, error) {
	return _DataLedgerContract.Contract.RetrieveInfo(&_DataLedgerContract.CallOpts, hash)
}

// DeleteMeasurement is a paid mutator transaction binding the contract method 0x77ad95ca.
//
// Solidity: function deleteMeasurement(bytes32 hash) returns()
func (_DataLedgerContract *DataLedgerContractTransactor) DeleteMeasurement(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.contract.Transact(opts, "deleteMeasurement", hash)
}

// DeleteMeasurement is a paid mutator transaction binding the contract method 0x77ad95ca.
//
// Solidity: function deleteMeasurement(bytes32 hash) returns()
func (_DataLedgerContract *DataLedgerContractSession) DeleteMeasurement(hash [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DeleteMeasurement(&_DataLedgerContract.TransactOpts, hash)
}

// DeleteMeasurement is a paid mutator transaction binding the contract method 0x77ad95ca.
//
// Solidity: function deleteMeasurement(bytes32 hash) returns()
func (_DataLedgerContract *DataLedgerContractTransactorSession) DeleteMeasurement(hash [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.DeleteMeasurement(&_DataLedgerContract.TransactOpts, hash)
}

// GetAddress is a paid mutator transaction binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) returns(address)
func (_DataLedgerContract *DataLedgerContractTransactor) GetAddress(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.contract.Transact(opts, "getAddress", id)
}

// GetAddress is a paid mutator transaction binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) returns(address)
func (_DataLedgerContract *DataLedgerContractSession) GetAddress(id [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.GetAddress(&_DataLedgerContract.TransactOpts, id)
}

// GetAddress is a paid mutator transaction binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) returns(address)
func (_DataLedgerContract *DataLedgerContractTransactorSession) GetAddress(id [32]byte) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.GetAddress(&_DataLedgerContract.TransactOpts, id)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_DataLedgerContract *DataLedgerContractTransactor) SetAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _DataLedgerContract.contract.Transact(opts, "setAddress", _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_DataLedgerContract *DataLedgerContractSession) SetAddress(_address common.Address) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.SetAddress(&_DataLedgerContract.TransactOpts, _address)
}

// SetAddress is a paid mutator transaction binding the contract method 0xe30081a0.
//
// Solidity: function setAddress(address _address) returns()
func (_DataLedgerContract *DataLedgerContractTransactorSession) SetAddress(_address common.Address) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.SetAddress(&_DataLedgerContract.TransactOpts, _address)
}

// StoreInfo is a paid mutator transaction binding the contract method 0xb45a85a9.
//
// Solidity: function storeInfo(bytes32 hash, string uri, bytes32 id, string description) returns()
func (_DataLedgerContract *DataLedgerContractTransactor) StoreInfo(opts *bind.TransactOpts, hash [32]byte, uri string, id [32]byte, description string) (*types.Transaction, error) {
	return _DataLedgerContract.contract.Transact(opts, "storeInfo", hash, uri, id, description)
}

// StoreInfo is a paid mutator transaction binding the contract method 0xb45a85a9.
//
// Solidity: function storeInfo(bytes32 hash, string uri, bytes32 id, string description) returns()
func (_DataLedgerContract *DataLedgerContractSession) StoreInfo(hash [32]byte, uri string, id [32]byte, description string) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.StoreInfo(&_DataLedgerContract.TransactOpts, hash, uri, id, description)
}

// StoreInfo is a paid mutator transaction binding the contract method 0xb45a85a9.
//
// Solidity: function storeInfo(bytes32 hash, string uri, bytes32 id, string description) returns()
func (_DataLedgerContract *DataLedgerContractTransactorSession) StoreInfo(hash [32]byte, uri string, id [32]byte, description string) (*types.Transaction, error) {
	return _DataLedgerContract.Contract.StoreInfo(&_DataLedgerContract.TransactOpts, hash, uri, id, description)
}

// DataLedgerContractDeleteInfoIterator is returned from FilterDeleteInfo and is used to iterate over the raw logs and unpacked data for DeleteInfo events raised by the DataLedgerContract contract.
type DataLedgerContractDeleteInfoIterator struct {
	Event *DataLedgerContractDeleteInfo // Event containing the contract specifics and raw log

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
func (it *DataLedgerContractDeleteInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLedgerContractDeleteInfo)
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
		it.Event = new(DataLedgerContractDeleteInfo)
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
func (it *DataLedgerContractDeleteInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLedgerContractDeleteInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLedgerContractDeleteInfo represents a DeleteInfo event raised by the DataLedgerContract contract.
type DataLedgerContractDeleteInfo struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDeleteInfo is a free log retrieval operation binding the contract event 0x072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe.
//
// Solidity: event deleteInfo(bytes32 indexed _hash)
func (_DataLedgerContract *DataLedgerContractFilterer) FilterDeleteInfo(opts *bind.FilterOpts, _hash [][32]byte) (*DataLedgerContractDeleteInfoIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.FilterLogs(opts, "deleteInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractDeleteInfoIterator{contract: _DataLedgerContract.contract, event: "deleteInfo", logs: logs, sub: sub}, nil
}

// WatchDeleteInfo is a free log subscription operation binding the contract event 0x072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe.
//
// Solidity: event deleteInfo(bytes32 indexed _hash)
func (_DataLedgerContract *DataLedgerContractFilterer) WatchDeleteInfo(opts *bind.WatchOpts, sink chan<- *DataLedgerContractDeleteInfo, _hash [][32]byte) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.WatchLogs(opts, "deleteInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLedgerContractDeleteInfo)
				if err := _DataLedgerContract.contract.UnpackLog(event, "deleteInfo", log); err != nil {
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

// ParseDeleteInfo is a log parse operation binding the contract event 0x072007d551e16de6c1b8938fdd0559f70033d87037e5dffa28631256df69f9fe.
//
// Solidity: event deleteInfo(bytes32 indexed _hash)
func (_DataLedgerContract *DataLedgerContractFilterer) ParseDeleteInfo(log types.Log) (*DataLedgerContractDeleteInfo, error) {
	event := new(DataLedgerContractDeleteInfo)
	if err := _DataLedgerContract.contract.UnpackLog(event, "deleteInfo", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DataLedgerContractEvtStoreInfoIterator is returned from FilterEvtStoreInfo and is used to iterate over the raw logs and unpacked data for EvtStoreInfo events raised by the DataLedgerContract contract.
type DataLedgerContractEvtStoreInfoIterator struct {
	Event *DataLedgerContractEvtStoreInfo // Event containing the contract specifics and raw log

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
func (it *DataLedgerContractEvtStoreInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataLedgerContractEvtStoreInfo)
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
		it.Event = new(DataLedgerContractEvtStoreInfo)
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
func (it *DataLedgerContractEvtStoreInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataLedgerContractEvtStoreInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataLedgerContractEvtStoreInfo represents a EvtStoreInfo event raised by the DataLedgerContract contract.
type DataLedgerContractEvtStoreInfo struct {
	Hash [32]byte
	Uri  string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEvtStoreInfo is a free log retrieval operation binding the contract event 0x06c4f2b8beb126621a4c74187a1573e7f48017b16d171df264507e2131b20f42.
//
// Solidity: event evtStoreInfo(bytes32 indexed _hash, string _uri)
func (_DataLedgerContract *DataLedgerContractFilterer) FilterEvtStoreInfo(opts *bind.FilterOpts, _hash [][32]byte) (*DataLedgerContractEvtStoreInfoIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.FilterLogs(opts, "evtStoreInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return &DataLedgerContractEvtStoreInfoIterator{contract: _DataLedgerContract.contract, event: "evtStoreInfo", logs: logs, sub: sub}, nil
}

// WatchEvtStoreInfo is a free log subscription operation binding the contract event 0x06c4f2b8beb126621a4c74187a1573e7f48017b16d171df264507e2131b20f42.
//
// Solidity: event evtStoreInfo(bytes32 indexed _hash, string _uri)
func (_DataLedgerContract *DataLedgerContractFilterer) WatchEvtStoreInfo(opts *bind.WatchOpts, sink chan<- *DataLedgerContractEvtStoreInfo, _hash [][32]byte) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _DataLedgerContract.contract.WatchLogs(opts, "evtStoreInfo", _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataLedgerContractEvtStoreInfo)
				if err := _DataLedgerContract.contract.UnpackLog(event, "evtStoreInfo", log); err != nil {
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

// ParseEvtStoreInfo is a log parse operation binding the contract event 0x06c4f2b8beb126621a4c74187a1573e7f48017b16d171df264507e2131b20f42.
//
// Solidity: event evtStoreInfo(bytes32 indexed _hash, string _uri)
func (_DataLedgerContract *DataLedgerContractFilterer) ParseEvtStoreInfo(log types.Log) (*DataLedgerContractEvtStoreInfo, error) {
	event := new(DataLedgerContractEvtStoreInfo)
	if err := _DataLedgerContract.contract.UnpackLog(event, "evtStoreInfo", log); err != nil {
		return nil, err
	}
	return event, nil
}
