package libs

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	accessControlContract "../contracts/accessContract"
	balanceContract "../contracts/balanceContract"
	dataContract "../contracts/dataContract"
	"github.com/ethereum/go-ethereum/ethclient"
)

/**************************** Contract Addresses *********************************/

// DataContractAddress Address of the contract that holds the event
var DataContractAddress common.Address = common.HexToAddress("0x77fd0aFF9117621d315D37b9966696fdDc49c007")

// AccessControlContractAddress address of the contract that controls the access to the blockchain
var AccessControlContractAddress common.Address = common.HexToAddress("0xDCDfe07b2a024B80637FD9FA2d23391d78880D23")

// BalanceContractAddress address of the contract that holds the purchases
var BalanceContractAddress common.Address = common.HexToAddress("0xB7124629A5f892a4E53D7C7399A881fb73aB3314")

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

// prepareTokenClient creates a token to those clients subscribed to it
// The token has the following fields:
// 		- SubscriptionTopic
//		- Client address
// 		- Expiration date
func sendTokenToClient(dataBlockchain map[string]interface{},
	ethclient *Ethereum, clientAddr common.Address,
) error {
	token := make(map[string]interface{})

	// Get the expiration date stored in the Balance Contract
	_, expirationDate, err := ethclient.BalanceCon.CheckSubStatus(&bind.CallOpts{From: common.HexToAddress(ADMINACCOUNT)},
		String2Bytes32(dataBlockchain["id"].(string)),
		clientAddr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Fill the fields of the Token
	token["topic"] = dataBlockchain["id"].(string)
	token["address"] = clientAddr.Hex()
	token["expiration"] = expirationDate
	token["url"] = dataBlockchain["url"].(string)

	// Create JSON token
	tokenJSON, err := json.Marshal(token)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Send a transaction to the client with the token
	auth := bind.NewKeyedTransactor(ethclient.AdminPrivKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	txHash, err := SendTransaction(common.HexToAddress(ADMINACCOUNT),
		clientAddr,
		ethclient.AdminPrivKey,
		ethclient,
		tokenJSON,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Emit event in the balance contract so client can anwser it
	_, err = ethclient.BalanceCon.SendToken(auth,
		ByteToByte32(txHash),
		clientAddr,
		String2Bytes32(dataBlockchain["id"].(string)))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// MangeSubscriptions checks if a subscription exists and
// who is subscribed to it.
func MangeSubscriptions(
	dataBlockchain map[string]interface{},
	ethclient *Ethereum,
) error {

	// Check whether the subscription exists
	isThere, err := ethclient.BalanceCon.CheckType(nil, String2Bytes32(dataBlockchain["id"].(string)))
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Set the parameters of the transaction to set the type as available
	// in the Balance SC in the Blockchain
	auth := bind.NewKeyedTransactor(ethclient.AdminPrivKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	// If it does not exist, create it
	if !isThere {
		_, err = ethclient.BalanceCon.AddNewType(auth, String2Bytes32(dataBlockchain["id"].(string)))
		if err != nil {
			fmt.Println(err)
			return err
		}

		// Check whether the type was stored successfully in the Blockchain
		// Wait until the value is received or the loop
		// is working for more than 15 seconds
		currentTime := time.Now()
		for {
			checkType, err := ethclient.BalanceCon.CheckType(nil, String2Bytes32(dataBlockchain["id"].(string)))
			if err != nil {
				fmt.Println(err)
				return err
			}

			if checkType {
				fmt.Printf("\nType %s was added to the blockchain\n\n", dataBlockchain["id"].(string))
				break
			}

			secondsPassed := time.Now().Sub(currentTime)
			if secondsPassed > 15*time.Second {
				fmt.Println("Could not check whether the data was introduced in the Blockchain")
				return errors.New("Could not check whether the data was introduced in the Blockchain")
			}
		}
		return nil
	}

	// Get all the account that are subscribed to the event
	addresses, err := ethclient.BalanceCon.GetSubsToType(&bind.CallOpts{From: common.HexToAddress(ADMINACCOUNT)}, String2Bytes32(dataBlockchain["id"].(string)))
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Check whether there is somenone subscribed
	if len(addresses) == 0 {
		fmt.Printf("\nNo customers subscribed to %s\n", dataBlockchain["id"].(string))
		return nil
	}

	fmt.Printf("\nAddresses subscribed to %s:", dataBlockchain["id"].(string))
	fmt.Println(addresses)

	// Send token to those clients that are subscribed to the event
	for _, address := range addresses {
		fmt.Printf("Sending %s token to %s", dataBlockchain["id"], fmt.Sprintln(address.Hex()))
		err = sendTokenToClient(dataBlockchain, ethclient, address)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("\nToken was sent successfully\n")
	}

	if err != nil {
		return err
	}

	return nil
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

	// Convert the hash to [32]byte
	hash32Byte, err := HexStringToBytes32(dataBlockchain["hash"].(string))
	if err != nil {
		return err
	}

	// Store the info of the measurement in the Blockchain
	_, err = ethclient.DataCon.StoreInfo(auth,
		hash32Byte,
		dataBlockchain["url"].(string),
		String2Bytes32(dataBlockchain["id"].(string)),
		dataBlockchain["description"].(string),
	)
	if err != nil {
		return err
	}

	// Check if the data has been stored in the contract
	// Wait until the value is received or the loop
	// works for more than 15 seconds
	currentTime := time.Now()
	for {
		var data1, data2 string
		data1, data2, err := ethclient.DataCon.RetrieveInfo(nil, hash32Byte)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if data1 != "" {
			fmt.Println()
			fmt.Println(data1)
			fmt.Println(data2)
			fmt.Println()
			break
		}

		secondsPassed := time.Now().Sub(currentTime)
		if secondsPassed > 15*time.Second {
			fmt.Println("Could not check whether the data was introduced in the Blockchain")
			return errors.New("Could not check whether the data was introduced in the Blockchain")
		}
	}

	// Process the subscriptions associated with the event
	err = MangeSubscriptions(dataBlockchain, ethclient)
	if err != nil {
		fmt.Println(err)
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
