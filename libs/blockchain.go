package libs

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	accessControlContract "../contracts/accessContract"
	balanceContract "../contracts/balanceContract"
	dataContract "../contracts/dataContract"
	"github.com/ethereum/go-ethereum/ethclient"
)

/**************************** Contract Addresses *********************************/

// DataContractAddress Address of the contract that holds the event
var DataContractAddress common.Address = common.HexToAddress("0x4b48bB7a81b5c01Df95F58B7C8504251eb04B670")

// AccessControlContractAddress address of the contract that controls the access to the blockchain
var AccessControlContractAddress common.Address = common.HexToAddress("0x5650D534eeC3A621b271DF1126E40772E4313cC4")

// BalanceContractAddress address of the contract that holds the purchases
var BalanceContractAddress common.Address = common.HexToAddress("0xB3Fd3051feA7caa87d400E29ba47e1218A0Ac19A")

/********************************************************************************/

// Ethereum is a struct that hodls the values attached to the blockchain
type Ethereum struct {
	EthereumClient *ethclient.Client
	DataCon        *dataContract.DataLedgerContract
	AccessCon      *accessControlContract.AccessControlContract
	BalanceCon     *balanceContract.BalanceContract
	AdminPrivKey   *ecdsa.PrivateKey
}

// PurchaseNotifyStruct stores the info of the events purchaseNotify
type PurchaseNotifyStruct struct {
	Addr   common.Address
	Hash   [32]byte
	Value  *big.Int
	TxHash []byte
}

// ResponseNotifyStruct store the event responseNotify
type ResponseNotifyStruct struct {
	Addr           common.Address
	Hash           [32]byte
	TxExchangeHash [32]byte
	TxData         [32]byte
	TxHash         []byte
}

// HexStringToBytes32 converts hex string to [32]byte
func HexStringToBytes32(str string) ([32]byte, error) {
	var bytes32 [32]byte

	// Converts the hex string to []byte
	bytes, err := hex.DecodeString(str)
	if err != nil {
		copy(bytes32[:], []byte("0"))
		return bytes32, err
	}

	copy(bytes32[:], bytes)
	return bytes32, nil
}

// ByteToByte32 converts []byte to [32]byte
func ByteToByte32(bytes []byte) [32]byte {
	var b [32]byte
	copy(b[:], bytes)
	return b
}

// ReadEventsFromBalanceContract reads the events associated to the balanceContract
func ReadEventsFromBalanceContract(ethclient *Ethereum, nameEvent string, filter map[string]interface{}) (*types.Log, error) {
	// Prepare the query to look for events
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Addresses: []common.Address{
			BalanceContractAddress,
		},
	}

	// Get the logs that hold the events
	logs, err := ethclient.EthereumClient.FilterLogs(context.Background(), query)
	if err != nil {
		fmt.Println("Error while creating the filter")
		return nil, err
	}

	// Get the appropiate interface of the contract to decode the logs
	contractAbi, err := abi.JSON(strings.NewReader(string(balanceContract.BalanceContractABI)))
	if err != nil {
		fmt.Println("Error while generating the abi")
		return nil, err
	}

	_ = contractAbi

	var logIndex common.Hash
	switch nameEvent {
	case "purchaseNotify":
		// Get the index associated to 'purchaseNotify'
		logPurchaseNotify := []byte("purchaseNotify(address,bytes32,uint256)")
		logIndex = crypto.Keccak256Hash(logPurchaseNotify)
	case "responseNotify":
		logResponseNotify := []byte("purchaseNotify(address,bytes32,bytes32,bytes32)")
		logIndex = crypto.Keccak256Hash(logResponseNotify)
	default:
		return nil, errors.New("Wrong name of the event")
	}

	switch nameEvent {
	case "purchaseNotify":
		// Iterate through the log to obtain the events
		for _, vLog := range logs {

			// Compare the index of the log with the one that it should be
			if vLog.Topics[0].Hex() == logIndex.Hex() {

				// Get the elements of the event which are indexed.
				// In this case there are to elements: _addr and _hash.
				// The other elements are stored in vLog.Data.
				addrLog := common.HexToAddress(vLog.Topics[1].Hex())
				hashLog := vLog.Topics[2]

				// Filter the events by Hash and address
				hashToCompare, _ := HexStringToBytes32(filter["Hash"].(string))
				if strings.ToLower(addrLog.Hex()) == filter["Addr"] && hashLog == hashToCompare {
					return &vLog, nil
				}
			}
		}
		return nil, errors.New("Not match found")

	case "responseNotify":
	default:
		return nil, errors.New("Wrong name of the event")
	}

	/*
		switch nameEvent {
		case "purchaseNotify":
			// Iterate through the log to obtain the events
			for _, vLog := range logs {
				// Declare the struct that will hold the fields of the event
				var event PurchaseNotifyStruct

				// Use the ABI decode the log
				err := contractAbi.Unpack(&event, nameEvent, vLog.Data)
				if err != nil {
					fmt.Println("Failed to unpack")
					return nil, nil, err
				}

				t := new(big.Int)
				t.SetBytes(vLog.Data)
				fmt.Printf("%d\n", t)

				// Unpack will not parse indexed event types because those are stored under topics
				// so they have to be unpacked manually
				event.Addr = common.HexToAddress(vLog.Topics[1].Hex())
				event.Hash = vLog.Topics[2]

				// Store the transaction hash
				event.TxHash = vLog.TxHash.Bytes()
				event.Addr = common.BytesToAddress(vLog.Topics[1].Bytes())

				// Filter the events by Hash and address
				hashToCompare, _ := HexStringToBytes32(filter["Hash"].(string))
				if strings.ToLower(event.Addr.Hex()) == filter["Addr"] && event.Hash == hashToCompare {
					fmt.Printf("TxHash: %x\n", event.TxHash)
					fmt.Printf("Addr: %s\n", event.Addr.Hex())
					fmt.Printf("Hash: %x\n", event.Hash[:])
					fmt.Printf("Value: %d\n\n", event.Value)

					return &event, nil, nil
				}
			}
			return nil, nil, errors.New("Not match found")

		case "responseNotify":
		default:
			return nil, nil, errors.New("Wrong name of the event")
		}
	*/

	// Decode the logs importing the abi of the contract
	return nil, nil
}

//InteractBlockchain stores the measurement in the blockchain
func InteractBlockchain(
	dataBlockchain map[string]interface{},
	producerPrivateKey *ecdsa.PrivateKey,
	ethclient *Ethereum,
) error {

	// Set the parameters of the transaction
	auth := bind.NewKeyedTransactor(producerPrivateKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	// Send the transaction
	hash32Byte, err := HexStringToBytes32(dataBlockchain["hash"].(string))
	if err != nil {
		return err
	}
	_, err = ethclient.DataCon.StoreInfo(auth,
		hash32Byte,
		dataBlockchain["url"].(string),
		String2Bytes32(dataBlockchain["id"].(string)),
		dataBlockchain["description"].(string),
	)
	if err != nil {
		return err
	}

	// Set the parameters of the new transaction to set the price of the product
	auth = bind.NewKeyedTransactor(ethclient.AdminPrivKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	// Set the price of the product
	price := big.NewInt(2)
	_, err = ethclient.BalanceCon.SetPriceData(auth, hash32Byte, price)
	if err != nil {
		return err
	}
	return nil
}

// SendTransaction sends transaction with data to the blockchain
func SendTransaction(from, to common.Address,
	privKey *ecdsa.PrivateKey,
	ethclient *Ethereum,
	data []byte) ([]byte, error) {

	// Set the parameters of the transaction
	value := big.NewInt(0)
	gasLimit := uint64(400000)
	gasPrice := big.NewInt(0)

	nonce, err := ethclient.EthereumClient.PendingNonceAt(context.Background(), from)
	if err != nil {
		log.Fatal(err)
	}

	// Create the transaction
	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, data)

	// Sign the transaction with the private key of the sender
	chainID, err := ethclient.EthereumClient.NetworkID(context.Background())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Send the transaction
	err = ethclient.EthereumClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return signedTx.Hash().Bytes(), nil
}
