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
const BalanceContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"subName\",\"type\":\"bytes32\"}],\"name\":\"getSubsToType\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllTypes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"subName\",\"type\":\"bytes32\"}],\"name\":\"addNewType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"subName\",\"type\":\"bytes32\"}],\"name\":\"deleteType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"anwserToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"subName\",\"type\":\"bytes32\"},{\"name\":\"clientAddr\",\"type\":\"address\"}],\"name\":\"deleteSub\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"subName\",\"type\":\"bytes32\"},{\"name\":\"clientAddr\",\"type\":\"address\"}],\"name\":\"checkSubStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"typeName\",\"type\":\"bytes32\"}],\"name\":\"checkType\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txHash\",\"type\":\"bytes32\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"sendToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"subName\",\"type\":\"bytes32\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"subscribeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"greet\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_subID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_endTime\",\"type\":\"uint256\"}],\"name\":\"notifyNew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_subID\",\"type\":\"bytes32\"}],\"name\":\"notifyRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"notifyNewCategory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"notifyRemoveCategory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"sendTokenEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_txHash\",\"type\":\"bytes32\"}],\"name\":\"anwserTokenEvent\",\"type\":\"event\"}]"

// BalanceContractFuncSigs maps the 4-byte function signature to its string representation.
var BalanceContractFuncSigs = map[string]string{
	"149cde75": "addNewType(bytes32)",
	"2236067f": "anwserToken(bytes32,bytes32)",
	"7dfde040": "checkSubStatus(bytes32,address)",
	"99a50b42": "checkType(bytes32)",
	"42007c86": "deleteSub(bytes32,address)",
	"1d9de32c": "deleteType(bytes32)",
	"09dddcf0": "getAllTypes()",
	"097796fd": "getSubsToType(bytes32)",
	"cfae3217": "greet()",
	"b074233a": "sendToken(bytes32,address,bytes32)",
	"c990824d": "subscribeTo(bytes32,uint256)",
}

// BalanceContractBin is the compiled bytecode used for deploying new contracts.
var BalanceContractBin = "0x608060405260008054600160a060020a0319167321a018606490c031a8c02bb6b992d8ae44add37f17905534801561003657600080fd5b506110ea806100466000396000f3006080604052600436106100ae5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663097796fd81146100b357806309dddcf01461011b578063149cde75146101305780631d9de32c1461014a5780632236067f1461016257806342007c861461017d5780637dfde040146101a157806399a50b42146101e0578063b074233a1461020c578063c990824d14610233578063cfae32171461024e575b600080fd5b3480156100bf57600080fd5b506100cb6004356102d8565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101075781810151838201526020016100ef565b505050509050019250505060405180910390f35b34801561012757600080fd5b506100cb6103cd565b34801561013c57600080fd5b50610148600435610429565b005b34801561015657600080fd5b506101486004356105da565b34801561016e57600080fd5b5061014860043560243561070e565b34801561018957600080fd5b50610148600435600160a060020a036024351661073f565b3480156101ad57600080fd5b506101c5600435600160a060020a0360243516610bd4565b60408051921515835260208301919091528051918290030190f35b3480156101ec57600080fd5b506101f8600435610cbf565b604080519115158252519081900360200190f35b34801561021857600080fd5b50610148600435600160a060020a0360243516604435610cd4565b34801561023f57600080fd5b50610148600435602435610d99565b34801561025a57600080fd5b5061026361105f565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561029d578181015183820152602001610285565b50505050905090810190601f1680156102ca5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600054606090600160a060020a03163314610363576040805160e560020a62461bcd02815260206004820152603360248201527f596f7520646f206e6f74206861766520656e6f7567682070726976696c65676560448201527f7320746f20646f207468697320616374696f6e00000000000000000000000000606482015290519081900360840190fd5b600082815260036020908152604091829020805483518184028101840190945280845290918301828280156103c157602002820191906000526020600020905b8154600160a060020a031681526001909101906020018083116103a3575b50505050509050919050565b6060600260040180548060200260200160405190810160405280929190818152602001828054801561041f57602002820191906000526020600020905b8154815260019091019060200180831161040a575b5050505050905090565b60008054600160a060020a031633146104b1576040805160e560020a62461bcd028152602060048201526024808201527f5573657220646f6573206e6f74206861766520656e6f7567682070726976696c60448201527f6567657300000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008281526002602052604090205460ff161561053d576040805160e560020a62461bcd028152602060048201526024808201527f546865207265717565737465642063617465676f727920616c7261647920657860448201527f6973747300000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000828152600260205260409020805460ff1916600117905560075415156105a357600680546001810182557ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f01839055546000838152600860205260409020556105a9565b60078054fe5b604051829033907fde3cd713f10754ba21eabefe59d58e9251b68f9ff9016fc20324b4e261f1d59890600090a35050565b600054600160a060020a03163314610661576040805160e560020a62461bcd028152602060048201526024808201527f5573657220646f6573206e6f74206861766520656e6f7567682070726976696c60448201527f6567657300000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008181526002602052604090205460ff1615156001146106f2576040805160e560020a62461bcd02815260206004820152602560248201527f546865207265717565737465642063617465676f727920646f6573206e6f742060448201527f6578697374000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000818152600260205260409020805460ff1916905560068054fe5b604051819083907f48585d03016d8be31bb329d396ec74ded946f7e152eaccce1445939fb7e4214790600090a35050565b60008054819081908190600160a060020a03163314156109a557600160a060020a03851660009081526001602081815260408084208a85529091529091205460ff161515146107fe576040805160e560020a62461bcd02815260206004820152602e60248201527f5468652075736572206973206e6f742073757363726962656420746f2074686560448201527f2072657175657374652074797065000000000000000000000000000000000000606482015290519081900360840190fd5b3360009081526001602081815260408084208a855290920190529020544211610871576040805160e560020a62461bcd02815260206004820152601c60248201527f537562736372697074696f6e206973207374696c6c2061637469766500000000604482015290519081900360640190fd5b600160a060020a03851660008181526001602081815260408084208b855260038101835281852054818452918520805460ff1916905594909352526002909101805491955090859081106108c157fe5b60009182526020808320909101829055878252600281526040808320805460ff1916905560048252808320600160a060020a03891684528252808320548984526005835281842080546001810182559085528385200181905589845260039092529091208054919450908490811061093557fe5b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff19169055878252600481526040808320600160a060020a03891680855292528083208390555188927fab7ba7d234105b15e41e97f548ae644ed346089e33e01f8f03a862bfaebacb9191a35b3360009081526001602081815260408084208a85529091529091205460ff16151514610a41576040805160e560020a62461bcd02815260206004820152602e60248201527f5468652075736572206973206e6f742073757363726962656420746f2074686560448201527f2072657175657374652074797065000000000000000000000000000000000000606482015290519081900360840190fd5b3360009081526001602081815260408084208a855290920190529020544210610ab4576040805160e560020a62461bcd02815260206004820152601a60248201527f537562736372697074696f6e206973206e6f7420616374697665000000000000604482015290519081900360640190fd5b3360008181526001602081815260408084208b855260038101835281852054818452918520805460ff191690559490935252600290910180549193509083908110610afb57fe5b60009182526020808320909101829055878252600281526040808320805460ff191690556004825280832033845282528083205489845260058352818420805460018101825590855283852001819055898452600390925290912080549192509082908110610b6657fe5b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff191690558782526004815260408083203380855292528083208390555188927fab7ba7d234105b15e41e97f548ae644ed346089e33e01f8f03a862bfaebacb9191a3505050505050565b600080548190600160a060020a0316331415610c5c57600160a060020a03831660009081526001602081815260408084208885529092019052902054421015610c2157610c21848461073f565b5050600160a060020a03811660009081526001602081815260408084208685528083528185205493019091529091205460ff90911690610cb8565b3360009081526001602081815260408084208885529092019052902054421015610c8a57610c8a843361073f565b50503360009081526001602081815260408084208685528083528185205493019091529091205460ff909116905b9250929050565b60009081526002602052604090205460ff1690565b600054600160a060020a03163314610d5c576040805160e560020a62461bcd02815260206004820152603360248201527f596f7520646f206e6f74206861766520656e6f7567682070726976696c65676560448201527f7320746f20646f207468697320616374696f6e00000000000000000000000000606482015290519081900360840190fd5b60405183908290600160a060020a038516907f026171e27f75693d971e6674100c15eae4e974c1cf862fe68f053af6ea9ff50b90600090a4505050565b600082815260026020526040812054819060ff161515600114610e2c576040805160e560020a62461bcd02815260206004820152602560248201527f546865207265717565737465642063617465676f727920646f6573206e6f742060448201527f6578697374000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b3360009081526001602081815260408084208885529092019052902054421015610e5a57610e5a843361073f565b33600090815260016020908152604080832087845290915290205460ff1615610ef3576040805160e560020a62461bcd02815260206004820152602a60248201527f546865207573657220697320616c72656164792073756273637269626520746f60448201527f2074686174207479706500000000000000000000000000000000000000000000606482015290519081900360840190fd5b336000818152600160208181526040808420898552808352818520805460ff19168517905580840183529084204289019055939092529052600401541515610f745733600090815260016020818152604080842060028101805494850181558086528386209094018990559254888552600390930190915290912055610f8a565b3360009081526001602052604090206004018054fe5b6000848152600260209081526040808320805460ff191660011790556005909152902054151561100e576000848152600360209081526040808320805460018101825581855283852001805473ffffffffffffffffffffffffffffffffffffffff191633908117909155888552905460048452828520918552925290912055611020565b60008481526005602052604090208054fe5b6040805142850181529051859133917fc7c18d7fd8f2b08630c0d4f0bcf252583ff50a5cf2c5d8811f8a4347191679d49181900360200190a350505050565b60408051606081018252603381527f48656c6c6f20796f7520686176652063616c6c65642074686520636f6e74726160208201527f63742062616c616e63652d737562732e736f6c0000000000000000000000000091810191909152905600a165627a7a723058208d4acb2164b3a87a792515a7a7831987e32b87ed7fabe11944c8f0f4d69569800029"

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

// CheckType is a free data retrieval call binding the contract method 0x99a50b42.
//
// Solidity: function checkType(bytes32 typeName) view returns(bool)
func (_BalanceContract *BalanceContractCaller) CheckType(opts *bind.CallOpts, typeName [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BalanceContract.contract.Call(opts, out, "checkType", typeName)
	return *ret0, err
}

// CheckType is a free data retrieval call binding the contract method 0x99a50b42.
//
// Solidity: function checkType(bytes32 typeName) view returns(bool)
func (_BalanceContract *BalanceContractSession) CheckType(typeName [32]byte) (bool, error) {
	return _BalanceContract.Contract.CheckType(&_BalanceContract.CallOpts, typeName)
}

// CheckType is a free data retrieval call binding the contract method 0x99a50b42.
//
// Solidity: function checkType(bytes32 typeName) view returns(bool)
func (_BalanceContract *BalanceContractCallerSession) CheckType(typeName [32]byte) (bool, error) {
	return _BalanceContract.Contract.CheckType(&_BalanceContract.CallOpts, typeName)
}

// GetAllTypes is a free data retrieval call binding the contract method 0x09dddcf0.
//
// Solidity: function getAllTypes() view returns(bytes32[])
func (_BalanceContract *BalanceContractCaller) GetAllTypes(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _BalanceContract.contract.Call(opts, out, "getAllTypes")
	return *ret0, err
}

// GetAllTypes is a free data retrieval call binding the contract method 0x09dddcf0.
//
// Solidity: function getAllTypes() view returns(bytes32[])
func (_BalanceContract *BalanceContractSession) GetAllTypes() ([][32]byte, error) {
	return _BalanceContract.Contract.GetAllTypes(&_BalanceContract.CallOpts)
}

// GetAllTypes is a free data retrieval call binding the contract method 0x09dddcf0.
//
// Solidity: function getAllTypes() view returns(bytes32[])
func (_BalanceContract *BalanceContractCallerSession) GetAllTypes() ([][32]byte, error) {
	return _BalanceContract.Contract.GetAllTypes(&_BalanceContract.CallOpts)
}

// GetSubsToType is a free data retrieval call binding the contract method 0x097796fd.
//
// Solidity: function getSubsToType(bytes32 subName) view returns(address[])
func (_BalanceContract *BalanceContractCaller) GetSubsToType(opts *bind.CallOpts, subName [32]byte) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BalanceContract.contract.Call(opts, out, "getSubsToType", subName)
	return *ret0, err
}

// GetSubsToType is a free data retrieval call binding the contract method 0x097796fd.
//
// Solidity: function getSubsToType(bytes32 subName) view returns(address[])
func (_BalanceContract *BalanceContractSession) GetSubsToType(subName [32]byte) ([]common.Address, error) {
	return _BalanceContract.Contract.GetSubsToType(&_BalanceContract.CallOpts, subName)
}

// GetSubsToType is a free data retrieval call binding the contract method 0x097796fd.
//
// Solidity: function getSubsToType(bytes32 subName) view returns(address[])
func (_BalanceContract *BalanceContractCallerSession) GetSubsToType(subName [32]byte) ([]common.Address, error) {
	return _BalanceContract.Contract.GetSubsToType(&_BalanceContract.CallOpts, subName)
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

// AnwserToken is a paid mutator transaction binding the contract method 0x2236067f.
//
// Solidity: function anwserToken(bytes32 hash, bytes32 txHash) returns()
func (_BalanceContract *BalanceContractTransactor) AnwserToken(opts *bind.TransactOpts, hash [32]byte, txHash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "anwserToken", hash, txHash)
}

// AnwserToken is a paid mutator transaction binding the contract method 0x2236067f.
//
// Solidity: function anwserToken(bytes32 hash, bytes32 txHash) returns()
func (_BalanceContract *BalanceContractSession) AnwserToken(hash [32]byte, txHash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.AnwserToken(&_BalanceContract.TransactOpts, hash, txHash)
}

// AnwserToken is a paid mutator transaction binding the contract method 0x2236067f.
//
// Solidity: function anwserToken(bytes32 hash, bytes32 txHash) returns()
func (_BalanceContract *BalanceContractTransactorSession) AnwserToken(hash [32]byte, txHash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.AnwserToken(&_BalanceContract.TransactOpts, hash, txHash)
}

// CheckSubStatus is a paid mutator transaction binding the contract method 0x7dfde040.
//
// Solidity: function checkSubStatus(bytes32 subName, address clientAddr) returns(bool, uint256)
func (_BalanceContract *BalanceContractTransactor) CheckSubStatus(opts *bind.TransactOpts, subName [32]byte, clientAddr common.Address) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "checkSubStatus", subName, clientAddr)
}

// CheckSubStatus is a paid mutator transaction binding the contract method 0x7dfde040.
//
// Solidity: function checkSubStatus(bytes32 subName, address clientAddr) returns(bool, uint256)
func (_BalanceContract *BalanceContractSession) CheckSubStatus(subName [32]byte, clientAddr common.Address) (*types.Transaction, error) {
	return _BalanceContract.Contract.CheckSubStatus(&_BalanceContract.TransactOpts, subName, clientAddr)
}

// CheckSubStatus is a paid mutator transaction binding the contract method 0x7dfde040.
//
// Solidity: function checkSubStatus(bytes32 subName, address clientAddr) returns(bool, uint256)
func (_BalanceContract *BalanceContractTransactorSession) CheckSubStatus(subName [32]byte, clientAddr common.Address) (*types.Transaction, error) {
	return _BalanceContract.Contract.CheckSubStatus(&_BalanceContract.TransactOpts, subName, clientAddr)
}

// DeleteSub is a paid mutator transaction binding the contract method 0x42007c86.
//
// Solidity: function deleteSub(bytes32 subName, address clientAddr) returns()
func (_BalanceContract *BalanceContractTransactor) DeleteSub(opts *bind.TransactOpts, subName [32]byte, clientAddr common.Address) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "deleteSub", subName, clientAddr)
}

// DeleteSub is a paid mutator transaction binding the contract method 0x42007c86.
//
// Solidity: function deleteSub(bytes32 subName, address clientAddr) returns()
func (_BalanceContract *BalanceContractSession) DeleteSub(subName [32]byte, clientAddr common.Address) (*types.Transaction, error) {
	return _BalanceContract.Contract.DeleteSub(&_BalanceContract.TransactOpts, subName, clientAddr)
}

// DeleteSub is a paid mutator transaction binding the contract method 0x42007c86.
//
// Solidity: function deleteSub(bytes32 subName, address clientAddr) returns()
func (_BalanceContract *BalanceContractTransactorSession) DeleteSub(subName [32]byte, clientAddr common.Address) (*types.Transaction, error) {
	return _BalanceContract.Contract.DeleteSub(&_BalanceContract.TransactOpts, subName, clientAddr)
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

// SendToken is a paid mutator transaction binding the contract method 0xb074233a.
//
// Solidity: function sendToken(bytes32 txHash, address to, bytes32 hash) returns()
func (_BalanceContract *BalanceContractTransactor) SendToken(opts *bind.TransactOpts, txHash [32]byte, to common.Address, hash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "sendToken", txHash, to, hash)
}

// SendToken is a paid mutator transaction binding the contract method 0xb074233a.
//
// Solidity: function sendToken(bytes32 txHash, address to, bytes32 hash) returns()
func (_BalanceContract *BalanceContractSession) SendToken(txHash [32]byte, to common.Address, hash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.SendToken(&_BalanceContract.TransactOpts, txHash, to, hash)
}

// SendToken is a paid mutator transaction binding the contract method 0xb074233a.
//
// Solidity: function sendToken(bytes32 txHash, address to, bytes32 hash) returns()
func (_BalanceContract *BalanceContractTransactorSession) SendToken(txHash [32]byte, to common.Address, hash [32]byte) (*types.Transaction, error) {
	return _BalanceContract.Contract.SendToken(&_BalanceContract.TransactOpts, txHash, to, hash)
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

// BalanceContractAnwserTokenEventIterator is returned from FilterAnwserTokenEvent and is used to iterate over the raw logs and unpacked data for AnwserTokenEvent events raised by the BalanceContract contract.
type BalanceContractAnwserTokenEventIterator struct {
	Event *BalanceContractAnwserTokenEvent // Event containing the contract specifics and raw log

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
func (it *BalanceContractAnwserTokenEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractAnwserTokenEvent)
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
		it.Event = new(BalanceContractAnwserTokenEvent)
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
func (it *BalanceContractAnwserTokenEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractAnwserTokenEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractAnwserTokenEvent represents a AnwserTokenEvent event raised by the BalanceContract contract.
type BalanceContractAnwserTokenEvent struct {
	Hash   [32]byte
	TxHash [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAnwserTokenEvent is a free log retrieval operation binding the contract event 0x48585d03016d8be31bb329d396ec74ded946f7e152eaccce1445939fb7e42147.
//
// Solidity: event anwserTokenEvent(bytes32 indexed _hash, bytes32 indexed _txHash)
func (_BalanceContract *BalanceContractFilterer) FilterAnwserTokenEvent(opts *bind.FilterOpts, _hash [][32]byte, _txHash [][32]byte) (*BalanceContractAnwserTokenEventIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}
	var _txHashRule []interface{}
	for _, _txHashItem := range _txHash {
		_txHashRule = append(_txHashRule, _txHashItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "anwserTokenEvent", _hashRule, _txHashRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractAnwserTokenEventIterator{contract: _BalanceContract.contract, event: "anwserTokenEvent", logs: logs, sub: sub}, nil
}

// WatchAnwserTokenEvent is a free log subscription operation binding the contract event 0x48585d03016d8be31bb329d396ec74ded946f7e152eaccce1445939fb7e42147.
//
// Solidity: event anwserTokenEvent(bytes32 indexed _hash, bytes32 indexed _txHash)
func (_BalanceContract *BalanceContractFilterer) WatchAnwserTokenEvent(opts *bind.WatchOpts, sink chan<- *BalanceContractAnwserTokenEvent, _hash [][32]byte, _txHash [][32]byte) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}
	var _txHashRule []interface{}
	for _, _txHashItem := range _txHash {
		_txHashRule = append(_txHashRule, _txHashItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "anwserTokenEvent", _hashRule, _txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractAnwserTokenEvent)
				if err := _BalanceContract.contract.UnpackLog(event, "anwserTokenEvent", log); err != nil {
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

// ParseAnwserTokenEvent is a log parse operation binding the contract event 0x48585d03016d8be31bb329d396ec74ded946f7e152eaccce1445939fb7e42147.
//
// Solidity: event anwserTokenEvent(bytes32 indexed _hash, bytes32 indexed _txHash)
func (_BalanceContract *BalanceContractFilterer) ParseAnwserTokenEvent(log types.Log) (*BalanceContractAnwserTokenEvent, error) {
	event := new(BalanceContractAnwserTokenEvent)
	if err := _BalanceContract.contract.UnpackLog(event, "anwserTokenEvent", log); err != nil {
		return nil, err
	}
	return event, nil
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

// BalanceContractSendTokenEventIterator is returned from FilterSendTokenEvent and is used to iterate over the raw logs and unpacked data for SendTokenEvent events raised by the BalanceContract contract.
type BalanceContractSendTokenEventIterator struct {
	Event *BalanceContractSendTokenEvent // Event containing the contract specifics and raw log

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
func (it *BalanceContractSendTokenEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalanceContractSendTokenEvent)
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
		it.Event = new(BalanceContractSendTokenEvent)
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
func (it *BalanceContractSendTokenEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalanceContractSendTokenEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalanceContractSendTokenEvent represents a SendTokenEvent event raised by the BalanceContract contract.
type BalanceContractSendTokenEvent struct {
	Addr   common.Address
	Hash   [32]byte
	TxHash [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSendTokenEvent is a free log retrieval operation binding the contract event 0x026171e27f75693d971e6674100c15eae4e974c1cf862fe68f053af6ea9ff50b.
//
// Solidity: event sendTokenEvent(address indexed _addr, bytes32 indexed _hash, bytes32 indexed _txHash)
func (_BalanceContract *BalanceContractFilterer) FilterSendTokenEvent(opts *bind.FilterOpts, _addr []common.Address, _hash [][32]byte, _txHash [][32]byte) (*BalanceContractSendTokenEventIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}
	var _txHashRule []interface{}
	for _, _txHashItem := range _txHash {
		_txHashRule = append(_txHashRule, _txHashItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "sendTokenEvent", _addrRule, _hashRule, _txHashRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractSendTokenEventIterator{contract: _BalanceContract.contract, event: "sendTokenEvent", logs: logs, sub: sub}, nil
}

// WatchSendTokenEvent is a free log subscription operation binding the contract event 0x026171e27f75693d971e6674100c15eae4e974c1cf862fe68f053af6ea9ff50b.
//
// Solidity: event sendTokenEvent(address indexed _addr, bytes32 indexed _hash, bytes32 indexed _txHash)
func (_BalanceContract *BalanceContractFilterer) WatchSendTokenEvent(opts *bind.WatchOpts, sink chan<- *BalanceContractSendTokenEvent, _addr []common.Address, _hash [][32]byte, _txHash [][32]byte) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}
	var _txHashRule []interface{}
	for _, _txHashItem := range _txHash {
		_txHashRule = append(_txHashRule, _txHashItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "sendTokenEvent", _addrRule, _hashRule, _txHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalanceContractSendTokenEvent)
				if err := _BalanceContract.contract.UnpackLog(event, "sendTokenEvent", log); err != nil {
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

// ParseSendTokenEvent is a log parse operation binding the contract event 0x026171e27f75693d971e6674100c15eae4e974c1cf862fe68f053af6ea9ff50b.
//
// Solidity: event sendTokenEvent(address indexed _addr, bytes32 indexed _hash, bytes32 indexed _txHash)
func (_BalanceContract *BalanceContractFilterer) ParseSendTokenEvent(log types.Log) (*BalanceContractSendTokenEvent, error) {
	event := new(BalanceContractSendTokenEvent)
	if err := _BalanceContract.contract.UnpackLog(event, "sendTokenEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DataLedgerContractABI is the input ABI used to generate the binding from.
const DataLedgerContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"retrieveInfo\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// DataLedgerContractFuncSigs maps the 4-byte function signature to its string representation.
var DataLedgerContractFuncSigs = map[string]string{
	"c9776a6d": "retrieveInfo(bytes32)",
}

// DataLedgerContractBin is the compiled bytecode used for deploying new contracts.
var DataLedgerContractBin = "0x608060405234801561001057600080fd5b5061016f806100206000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663c9776a6d8114610045575b600080fd5b34801561005157600080fd5b5061005d60043561013b565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561009e578181015183820152602001610086565b50505050905090810190601f1680156100cb5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156100fe5781810151838201526020016100e6565b50505050905090810190601f16801561012b5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b6060809150915600a165627a7a723058203d3790d1f7672e532c6c67b3f52453aeaeef1305b01e18559497b548f82c25600029"

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
