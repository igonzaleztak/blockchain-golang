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
<<<<<<< HEAD
const BalanceContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"subName\",\"type\":\"bytes32\"}],\"name\":\"getSubsToType\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllTypes\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"subName\",\"type\":\"bytes32\"}],\"name\":\"addNewType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"subName\",\"type\":\"bytes32\"}],\"name\":\"deleteType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"txHash\",\"type\":\"bytes32\"}],\"name\":\"anwserToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"subName\",\"type\":\"bytes32\"},{\"name\":\"clientAddr\",\"type\":\"address\"}],\"name\":\"checkSubStatus\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"typeName\",\"type\":\"bytes32\"}],\"name\":\"checkType\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"txHash\",\"type\":\"bytes32\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"sendToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"subName\",\"type\":\"bytes32\"},{\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"subscribeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"greet\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"subName\",\"type\":\"bytes32\"}],\"name\":\"deleteSub\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_subID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_endTime\",\"type\":\"uint256\"}],\"name\":\"notifyNew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_subID\",\"type\":\"bytes32\"}],\"name\":\"notifyRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"notifyNewCategory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"notifyRemoveCategory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_txhahs\",\"type\":\"bytes32\"}],\"name\":\"sendTokenEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"_txhash\",\"type\":\"bytes32\"}],\"name\":\"anwserTokenEvent\",\"type\":\"event\"}]"
=======
const BalanceContractABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_subID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"}],\"name\":\"notifyNew\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"notifyNewCategory\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_subID\",\"type\":\"bytes32\"}],\"name\":\"notifyRemove\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"_name\",\"type\":\"bytes32\"}],\"name\":\"notifyRemoveCategory\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subName\",\"type\":\"bytes32\"}],\"name\":\"addNewType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subName\",\"type\":\"bytes32\"}],\"name\":\"deleteSub\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subName\",\"type\":\"bytes32\"}],\"name\":\"deleteType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllSubscriptions\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllTopics\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"greet\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subName\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"subscribeTo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653

// BalanceContractFuncSigs maps the 4-byte function signature to its string representation.
var BalanceContractFuncSigs = map[string]string{
	"149cde75": "addNewType(bytes32)",
<<<<<<< HEAD
	"2236067f": "anwserToken(bytes32,bytes32)",
	"7dfde040": "checkSubStatus(bytes32,address)",
	"99a50b42": "checkType(bytes32)",
	"f356cc79": "deleteSub(bytes32)",
	"1d9de32c": "deleteType(bytes32)",
	"09dddcf0": "getAllTypes()",
	"097796fd": "getSubsToType(bytes32)",
	"cfae3217": "greet()",
	"b074233a": "sendToken(bytes32,address,bytes32)",
=======
	"f356cc79": "deleteSub(bytes32)",
	"1d9de32c": "deleteType(bytes32)",
	"56eb00a1": "getAllSubscriptions()",
	"86416f14": "getAllTopics()",
	"cfae3217": "greet()",
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653
	"c990824d": "subscribeTo(bytes32,uint256)",
}

// BalanceContractBin is the compiled bytecode used for deploying new contracts.
<<<<<<< HEAD
var BalanceContractBin = "0x608060405260008054600160a060020a0319167321a018606490c031a8c02bb6b992d8ae44add37f17905534801561003657600080fd5b50610de5806100466000396000f3006080604052600436106100ae5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663097796fd81146100b357806309dddcf01461011b578063149cde75146101305780631d9de32c1461014a5780632236067f146101625780637dfde0401461017d57806399a50b42146101bc578063b074233a146101e8578063c990824d1461020f578063cfae32171461022a578063f356cc79146102b4575b600080fd5b3480156100bf57600080fd5b506100cb6004356102cc565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101075781810151838201526020016100ef565b505050509050019250505060405180910390f35b34801561012757600080fd5b506100cb6103c1565b34801561013c57600080fd5b5061014860043561041d565b005b34801561015657600080fd5b506101486004356105ce565b34801561016e57600080fd5b50610148600435602435610702565b34801561018957600080fd5b506101a1600435600160a060020a0360243516610733565b60408051921515835260208301919091528051918290030190f35b3480156101c857600080fd5b506101d46004356107b9565b604080519115158252519081900360200190f35b3480156101f457600080fd5b50610148600435600160a060020a03602435166044356107ce565b34801561021b57600080fd5b50610148600435602435610893565b34801561023657600080fd5b5061023f610b2b565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610279578181015183820152602001610261565b50505050905090810190601f1680156102a65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102c057600080fd5b50610148600435610b8a565b600054606090600160a060020a03163314610357576040805160e560020a62461bcd02815260206004820152603360248201527f596f7520646f206e6f74206861766520656e6f7567682070726976696c65676560448201527f7320746f20646f207468697320616374696f6e00000000000000000000000000606482015290519081900360840190fd5b600082815260036020908152604091829020805483518184028101840190945280845290918301828280156103b557602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311610397575b50505050509050919050565b6060600260040180548060200260200160405190810160405280929190818152602001828054801561041357602002820191906000526020600020905b815481526001909101906020018083116103fe575b5050505050905090565b60008054600160a060020a031633146104a5576040805160e560020a62461bcd028152602060048201526024808201527f5573657220646f6573206e6f74206861766520656e6f7567682070726976696c60448201527f6567657300000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008281526002602052604090205460ff1615610531576040805160e560020a62461bcd028152602060048201526024808201527f546865207265717565737465642063617465676f727920616c7261647920657860448201527f6973747300000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000828152600260205260409020805460ff19166001179055600754151561059757600680546001810182557ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f018390555460008381526008602052604090205561059d565b60078054fe5b604051829033907fde3cd713f10754ba21eabefe59d58e9251b68f9ff9016fc20324b4e261f1d59890600090a35050565b600054600160a060020a03163314610655576040805160e560020a62461bcd028152602060048201526024808201527f5573657220646f6573206e6f74206861766520656e6f7567682070726976696c60448201527f6567657300000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b60008181526002602052604090205460ff1615156001146106e6576040805160e560020a62461bcd02815260206004820152602560248201527f546865207265717565737465642063617465676f727920646f6573206e6f742060448201527f6578697374000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b6000818152600260205260409020805460ff1916905560068054fe5b604051819083907f48585d03016d8be31bb329d396ec74ded946f7e152eaccce1445939fb7e4214790600090a35050565b600080548190600160a060020a0316331415610784575050600160a060020a03811660009081526001602081815260408084208685528083528185205493019091529091205460ff909116906107b2565b50503360009081526001602081815260408084208685528083528185205493019091529091205460ff909116905b9250929050565b60009081526002602052604090205460ff1690565b600054600160a060020a03163314610856576040805160e560020a62461bcd02815260206004820152603360248201527f596f7520646f206e6f74206861766520656e6f7567682070726976696c65676560448201527f7320746f20646f207468697320616374696f6e00000000000000000000000000606482015290519081900360840190fd5b60405183908290600160a060020a038516907f026171e27f75693d971e6674100c15eae4e974c1cf862fe68f053af6ea9ff50b90600090a4505050565b600082815260026020526040812054819060ff161515600114610926576040805160e560020a62461bcd02815260206004820152602560248201527f546865207265717565737465642063617465676f727920646f6573206e6f742060448201527f6578697374000000000000000000000000000000000000000000000000000000606482015290519081900360840190fd5b33600090815260016020908152604080832087845290915290205460ff16156109bf576040805160e560020a62461bcd02815260206004820152602a60248201527f546865207573657220697320616c72656164792073756273637269626520746f60448201527f2074686174207479706500000000000000000000000000000000000000000000606482015290519081900360840190fd5b336000818152600160208181526040808420898552808352818520805460ff19168517905580840183529084204289019055939092529052600401541515610a405733600090815260016020818152604080842060028101805494850181558086528386209094018990559254888552600390930190915290912055610a56565b3360009081526001602052604090206004018054fe5b6000848152600260209081526040808320805460ff1916600117905560059091529020541515610ada576000848152600360209081526040808320805460018101825581855283852001805473ffffffffffffffffffffffffffffffffffffffff191633908117909155888552905460048452828520918552925290912055610aec565b60008481526005602052604090208054fe5b6040805142850181529051859133917fc7c18d7fd8f2b08630c0d4f0bcf252583ff50a5cf2c5d8811f8a4347191679d49181900360200190a350505050565b60408051606081018252603381527f48656c6c6f20796f7520686176652063616c6c65642074686520636f6e74726160208201527f63742062616c616e63652d737562732e736f6c000000000000000000000000009181019190915290565b336000908152600160208181526040808420858552909152822054829160ff909116151514610c29576040805160e560020a62461bcd02815260206004820152602e60248201527f5468652075736572206973206e6f742073757363726962656420746f2074686560448201527f2072657175657374652074797065000000000000000000000000000000000000606482015290519081900360840190fd5b33600090815260016020818152604080842087855290920190529020544210610c9c576040805160e560020a62461bcd02815260206004820152601a60248201527f537562736372697074696f6e206973206e6f7420616374697665000000000000604482015290519081900360640190fd5b33600081815260016020818152604080842088855260038101835281852054818452918520805460ff191690559490935252600290910180549193509083908110610ce357fe5b60009182526020808320909101829055848252600281526040808320805460ff191690556004825280832033845282528083205486845260058352818420805460018101825590855283852001819055868452600390925290912080549192509082908110610d4e57fe5b60009182526020808320909101805473ffffffffffffffffffffffffffffffffffffffff191690558482526004815260408083203380855292528083208390555185927fab7ba7d234105b15e41e97f548ae644ed346089e33e01f8f03a862bfaebacb9191a35050505600a165627a7a723058204eb2e2db28f0a94cd37b913442cda734be56bf6b4a65a47de96bfd21199bd9ed0029"
=======
var BalanceContractBin = "0x6080604052600080546001600160a01b0319167321a018606490c031a8c02bb6b992d8ae44add37f17905534801561003657600080fd5b5061086f806100466000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806386416f141161005b57806386416f1414610116578063c990824d1461011e578063cfae321714610141578063f356cc79146101be5761007d565b8063149cde75146100825780631d9de32c146100a157806356eb00a1146100be575b600080fd5b61009f6004803603602081101561009857600080fd5b50356101db565b005b61009f600480360360208110156100b757600080fd5b50356102fa565b6100c66103b1565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101025781810151838201526020016100ea565b505050509050019250505060405180910390f35b6100c6610415565b61009f6004803603604081101561013457600080fd5b508035906020013561046e565b6101496105c1565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561018357818101518382015260200161016b565b50505050905090810190601f1680156101b05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b61009f600480360360208110156101d457600080fd5b50356105e1565b6000546001600160a01b031633146102245760405162461bcd60e51b81526004018080602001828103825260248152602001806108176024913960400191505060405180910390fd5b60008181526002602052604090205460ff16156102725760405162461bcd60e51b81526004018080602001828103825260248152602001806107f36024913960400191505060405180910390fd5b6000818152600260209081526040808320805460ff191660019081179091556003805491820181557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b9091018590555460049092528083209190915551829133917fde3cd713f10754ba21eabefe59d58e9251b68f9ff9016fc20324b4e261f1d5989190a350565b6000546001600160a01b031633146103435760405162461bcd60e51b81526004018080602001828103825260248152602001806108176024913960400191505060405180910390fd5b60008181526002602052604090205460ff1615156001146103955760405162461bcd60e51b81526004018080602001828103825260258152602001806107716025913960400191505060405180910390fd5b6000818152600260205260409020805460ff1916905560038054fe5b3360009081526001602090815260409182902060020180548351818402810184019094528084526060939283018282801561040b57602002820191906000526020600020905b8154815260200190600101908083116103f7575b5050505050905090565b6060600260010180548060200260200160405190810160405280929190818152602001828054801561040b57602002820191906000526020600020908154815260200190600101908083116103f7575050505050905090565b60008281526002602052604090205460ff1615156001146104c05760405162461bcd60e51b81526004018080602001828103825260258152602001806107716025913960400191505060405180910390fd5b33600090815260016020908152604080832085845290915290205460ff161561051a5760405162461bcd60e51b815260040180806020018281038252602a8152602001806107c9602a913960400191505060405180910390fd5b336000818152600160208181526040808420878552808352818520805460ff199081168617909155818501845282862042890190819055600280840180548089018255818a52878a20018c9055548b89526003909401865284882093909355918452948290208054909516909317909355825191825291518593927fc7c18d7fd8f2b08630c0d4f0bcf252583ff50a5cf2c5d8811f8a4347191679d4928290030190a35050565b606060405180606001604052806033815260200161079660339139905090565b3360009081526001602081815260408084208585529091529091205460ff1615151461063e5760405162461bcd60e51b815260040180806020018281038252602e815260200180610743602e913960400191505060405180910390fd5b336000908152600160208181526040808420858552909201905290205442106106ae576040805162461bcd60e51b815260206004820152601a60248201527f537562736372697074696f6e206973206e6f7420616374697665000000000000604482015290519081900360640190fd5b33600081815260016020818152604080842086855260038101835281852054818452918520805460ff19169055949093525260029091018054829081106106f157fe5b60009182526020808320909101829055838252600290526040808220805460ff1916905551839133917fab7ba7d234105b15e41e97f548ae644ed346089e33e01f8f03a862bfaebacb919190a3505056fe5468652075736572206973206e6f742073757363726962656420746f207468652072657175657374652074797065546865207265717565737465642063617465676f727920646f6573206e6f7420657869737448656c6c6f20796f7520686176652063616c6c65642074686520636f6e74726163742062616c616e63652d737562732e736f6c546865207573657220697320616c72656164792073756273637269626520746f20746861742074797065546865207265717565737465642063617465676f727920616c72616479206578697374735573657220646f6573206e6f74206861766520656e6f7567682070726976696c65676573a265627a7a7231582023a2062e4b7787f79dc73da5f311fe6dd3f50f84ddebdd0c21314e9bd7cbc99f64736f6c63430005100032"
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653

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

<<<<<<< HEAD
// CheckSubStatus is a free data retrieval call binding the contract method 0x7dfde040.
//
// Solidity: function checkSubStatus(bytes32 subName, address clientAddr) view returns(bool, uint256)
func (_BalanceContract *BalanceContractCaller) CheckSubStatus(opts *bind.CallOpts, subName [32]byte, clientAddr common.Address) (bool, *big.Int, error) {
	var (
		ret0 = new(bool)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _BalanceContract.contract.Call(opts, out, "checkSubStatus", subName, clientAddr)
	return *ret0, *ret1, err
}

// CheckSubStatus is a free data retrieval call binding the contract method 0x7dfde040.
//
// Solidity: function checkSubStatus(bytes32 subName, address clientAddr) view returns(bool, uint256)
func (_BalanceContract *BalanceContractSession) CheckSubStatus(subName [32]byte, clientAddr common.Address) (bool, *big.Int, error) {
	return _BalanceContract.Contract.CheckSubStatus(&_BalanceContract.CallOpts, subName, clientAddr)
}

// CheckSubStatus is a free data retrieval call binding the contract method 0x7dfde040.
//
// Solidity: function checkSubStatus(bytes32 subName, address clientAddr) view returns(bool, uint256)
func (_BalanceContract *BalanceContractCallerSession) CheckSubStatus(subName [32]byte, clientAddr common.Address) (bool, *big.Int, error) {
	return _BalanceContract.Contract.CheckSubStatus(&_BalanceContract.CallOpts, subName, clientAddr)
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
=======
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
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653
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
<<<<<<< HEAD
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
=======
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
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653
}

// DeleteSub is a paid mutator transaction binding the contract method 0xf356cc79.
//
// Solidity: function deleteSub(bytes32 subName) returns()
<<<<<<< HEAD
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
=======
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
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653
}

// DeleteType is a paid mutator transaction binding the contract method 0x1d9de32c.
//
// Solidity: function deleteType(bytes32 subName) returns()
<<<<<<< HEAD
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
=======
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
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653
}

// SubscribeTo is a paid mutator transaction binding the contract method 0xc990824d.
//
// Solidity: function subscribeTo(bytes32 subName, uint256 time) returns()
<<<<<<< HEAD
func (_BalanceContract *BalanceContractTransactor) SubscribeTo(opts *bind.TransactOpts, subName [32]byte, time *big.Int) (*types.Transaction, error) {
	return _BalanceContract.contract.Transact(opts, "subscribeTo", subName, time)
=======
func (_BalanceContract *BalanceContractSession) SubscribeTo(subName [32]byte, time *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.SubscribeTo(&_BalanceContract.TransactOpts, subName, time)
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653
}

// SubscribeTo is a paid mutator transaction binding the contract method 0xc990824d.
//
// Solidity: function subscribeTo(bytes32 subName, uint256 time) returns()
<<<<<<< HEAD
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
=======
func (_BalanceContract *BalanceContractTransactorSession) SubscribeTo(subName [32]byte, time *big.Int) (*types.Transaction, error) {
	return _BalanceContract.Contract.SubscribeTo(&_BalanceContract.TransactOpts, subName, time)
}

// BalanceContractNotifyNewIterator is returned from FilterNotifyNew and is used to iterate over the raw logs and unpacked data for NotifyNew events raised by the BalanceContract contract.
type BalanceContractNotifyNewIterator struct {
	Event *BalanceContractNotifyNew // Event containing the contract specifics and raw log
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653

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
<<<<<<< HEAD
func (it *BalanceContractAnwserTokenEventIterator) Next() bool {
=======
func (it *BalanceContractNotifyNewIterator) Next() bool {
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
<<<<<<< HEAD
			it.Event = new(BalanceContractAnwserTokenEvent)
=======
			it.Event = new(BalanceContractNotifyNew)
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653
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
<<<<<<< HEAD
		it.Event = new(BalanceContractAnwserTokenEvent)
=======
		it.Event = new(BalanceContractNotifyNew)
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653
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
<<<<<<< HEAD
func (it *BalanceContractAnwserTokenEventIterator) Error() error {
=======
func (it *BalanceContractNotifyNewIterator) Error() error {
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
<<<<<<< HEAD
func (it *BalanceContractAnwserTokenEventIterator) Close() error {
=======
func (it *BalanceContractNotifyNewIterator) Close() error {
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653
	it.sub.Unsubscribe()
	return nil
}

<<<<<<< HEAD
// BalanceContractAnwserTokenEvent represents a AnwserTokenEvent event raised by the BalanceContract contract.
type BalanceContractAnwserTokenEvent struct {
	Hash   [32]byte
	Txhash [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAnwserTokenEvent is a free log retrieval operation binding the contract event 0x48585d03016d8be31bb329d396ec74ded946f7e152eaccce1445939fb7e42147.
//
// Solidity: event anwserTokenEvent(bytes32 indexed _hash, bytes32 indexed _txhash)
func (_BalanceContract *BalanceContractFilterer) FilterAnwserTokenEvent(opts *bind.FilterOpts, _hash [][32]byte, _txhash [][32]byte) (*BalanceContractAnwserTokenEventIterator, error) {
=======
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
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}
	var _txhashRule []interface{}
	for _, _txhashItem := range _txhash {
		_txhashRule = append(_txhashRule, _txhashItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "anwserTokenEvent", _hashRule, _txhashRule)
	if err != nil {
		return nil, err
	}
<<<<<<< HEAD
	return &BalanceContractAnwserTokenEventIterator{contract: _BalanceContract.contract, event: "anwserTokenEvent", logs: logs, sub: sub}, nil
}

// WatchAnwserTokenEvent is a free log subscription operation binding the contract event 0x48585d03016d8be31bb329d396ec74ded946f7e152eaccce1445939fb7e42147.
//
// Solidity: event anwserTokenEvent(bytes32 indexed _hash, bytes32 indexed _txhash)
func (_BalanceContract *BalanceContractFilterer) WatchAnwserTokenEvent(opts *bind.WatchOpts, sink chan<- *BalanceContractAnwserTokenEvent, _hash [][32]byte, _txhash [][32]byte) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
=======
	var _subIDRule []interface{}
	for _, _subIDItem := range _subID {
		_subIDRule = append(_subIDRule, _subIDItem)
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653
	}
	var _txhashRule []interface{}
	for _, _txhashItem := range _txhash {
		_txhashRule = append(_txhashRule, _txhashItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "anwserTokenEvent", _hashRule, _txhashRule)
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
// Solidity: event anwserTokenEvent(bytes32 indexed _hash, bytes32 indexed _txhash)
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
<<<<<<< HEAD
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
	Txhahs [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSendTokenEvent is a free log retrieval operation binding the contract event 0x026171e27f75693d971e6674100c15eae4e974c1cf862fe68f053af6ea9ff50b.
//
// Solidity: event sendTokenEvent(address indexed _addr, bytes32 indexed _hash, bytes32 indexed _txhahs)
func (_BalanceContract *BalanceContractFilterer) FilterSendTokenEvent(opts *bind.FilterOpts, _addr []common.Address, _hash [][32]byte, _txhahs [][32]byte) (*BalanceContractSendTokenEventIterator, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}
	var _txhahsRule []interface{}
	for _, _txhahsItem := range _txhahs {
		_txhahsRule = append(_txhahsRule, _txhahsItem)
	}

	logs, sub, err := _BalanceContract.contract.FilterLogs(opts, "sendTokenEvent", _addrRule, _hashRule, _txhahsRule)
	if err != nil {
		return nil, err
	}
	return &BalanceContractSendTokenEventIterator{contract: _BalanceContract.contract, event: "sendTokenEvent", logs: logs, sub: sub}, nil
}

// WatchSendTokenEvent is a free log subscription operation binding the contract event 0x026171e27f75693d971e6674100c15eae4e974c1cf862fe68f053af6ea9ff50b.
//
// Solidity: event sendTokenEvent(address indexed _addr, bytes32 indexed _hash, bytes32 indexed _txhahs)
func (_BalanceContract *BalanceContractFilterer) WatchSendTokenEvent(opts *bind.WatchOpts, sink chan<- *BalanceContractSendTokenEvent, _addr []common.Address, _hash [][32]byte, _txhahs [][32]byte) (event.Subscription, error) {

	var _addrRule []interface{}
	for _, _addrItem := range _addr {
		_addrRule = append(_addrRule, _addrItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}
	var _txhahsRule []interface{}
	for _, _txhahsItem := range _txhahs {
		_txhahsRule = append(_txhahsRule, _txhahsItem)
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "sendTokenEvent", _addrRule, _hashRule, _txhahsRule)
=======
	}

	logs, sub, err := _BalanceContract.contract.WatchLogs(opts, "notifyNewCategory", _addrRule, _nameRule)
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
<<<<<<< HEAD
				event := new(BalanceContractSendTokenEvent)
				if err := _BalanceContract.contract.UnpackLog(event, "sendTokenEvent", log); err != nil {
=======
				event := new(BalanceContractNotifyNewCategory)
				if err := _BalanceContract.contract.UnpackLog(event, "notifyNewCategory", log); err != nil {
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653
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

<<<<<<< HEAD
// ParseSendTokenEvent is a log parse operation binding the contract event 0x026171e27f75693d971e6674100c15eae4e974c1cf862fe68f053af6ea9ff50b.
//
// Solidity: event sendTokenEvent(address indexed _addr, bytes32 indexed _hash, bytes32 indexed _txhahs)
func (_BalanceContract *BalanceContractFilterer) ParseSendTokenEvent(log types.Log) (*BalanceContractSendTokenEvent, error) {
	event := new(BalanceContractSendTokenEvent)
	if err := _BalanceContract.contract.UnpackLog(event, "sendTokenEvent", log); err != nil {
=======
// ParseNotifyNewCategory is a log parse operation binding the contract event 0xde3cd713f10754ba21eabefe59d58e9251b68f9ff9016fc20324b4e261f1d598.
//
// Solidity: event notifyNewCategory(address indexed _addr, bytes32 indexed _name)
func (_BalanceContract *BalanceContractFilterer) ParseNotifyNewCategory(log types.Log) (*BalanceContractNotifyNewCategory, error) {
	event := new(BalanceContractNotifyNewCategory)
	if err := _BalanceContract.contract.UnpackLog(event, "notifyNewCategory", log); err != nil {
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653
		return nil, err
	}
	return event, nil
}

// BalanceContractNotifyRemoveIterator is returned from FilterNotifyRemove and is used to iterate over the raw logs and unpacked data for NotifyRemove events raised by the BalanceContract contract.
type BalanceContractNotifyRemoveIterator struct {
	Event *BalanceContractNotifyRemove // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

<<<<<<< HEAD
// DataLedgerContractBin is the compiled bytecode used for deploying new contracts.
var DataLedgerContractBin = "0x608060405234801561001057600080fd5b5061016f806100206000396000f3006080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663c9776a6d8114610045575b600080fd5b34801561005157600080fd5b5061005d60043561013b565b604051808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561009e578181015183820152602001610086565b50505050905090810190601f1680156100cb5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b838110156100fe5781810151838201526020016100e6565b50505050905090810190601f16801561012b5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b6060809150915600a165627a7a72305820ee08d9796005b5a0e6b88e9d98f455e4404196df989d5fad9152c27fd5cb52170029"
=======
	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}
>>>>>>> c596b8087f983583ad67605a3306c2f1dd772653

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
