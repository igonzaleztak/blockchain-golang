package clientap

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	cipher "../cipherLibs"
	libs "../libs"
)

// SecretStruct Struct that holds the encrypted keys that
// will be send in a transaction
type SecretStruct struct {
	AdminSecret  []byte
	ClientSecret []byte
}

// ResponseClient struct that contains the response to the client
type ResponseClient struct {
	TxSecretHash []byte
	TxDataHash   []byte
}

// getSignedMessage Gets the message that was signed
func getSignedMessage(account, hash string) ([]byte, error) {
	// Make up the message that was signed
	var sigMsgStr string = account[2:] + hash[2:]

	// Do the hash
	h := sha256.New()
	h.Write([]byte(sigMsgStr))

	return h.Sum(nil), nil
}

// notifyBlockchain notifies the blokchain that the payment has been accepted
func notifyBlockchain(ethClient *libs.Ethereum, hash [32]byte, client common.Address) error {
	// Set the parameters of the new transaction to set the price of the product
	auth := bind.NewKeyedTransactor(ethClient.AdminPrivKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	// Legacy value. It does not do anything. It might be use in the future.
	var a [32]byte
	copy(a[:], []byte("0"))

	_, err := ethClient.BalanceCon.SendData(auth, client, hash, a)
	if err != nil {
		return err
	}
	return nil
}

// ProccessClientPurchase Processes the purchases of the client
// and returns the measurements to him
func ProccessClientPurchase(ethClient *libs.Ethereum, body map[string]interface{}) ([]byte, error) {

	// Get the hash that the client wants to purchase
	var hash string = body["_hash"].(string)

	// Get the account that made the purchase
	var account string = body["_account"].(string)

	fmt.Printf("\nThe client %s  wants to buy %s\n", account, hash)

	// address formatted
	addrFormatted := common.HexToAddress(account[2:])

	// Get the signature
	var signature string = body["_signature"].(string)

	// Get the hash that was signed
	signedMessage, err := getSignedMessage(account, hash)
	if err != nil {
		return nil, err
	}

	// Convert the signature to bytes
	signatureBytes, err := hex.DecodeString(signature[2:])
	if err != nil {
		return nil, err
	}

	// Get the publicKey of the client
	clientPubKey, err := ethClient.AccessCon.GetPubKey(nil, addrFormatted)
	if err != nil {
		return nil, err
	}

	// Convert the public key to []byte
	clientPubKeyBytes, err := hex.DecodeString("04" + clientPubKey)
	if err != nil {
		return nil, err
	}

	// Verify the signature
	var isProperlySigned bool = cipher.VerifySignature(clientPubKeyBytes, signedMessage, signatureBytes)

	if !isProperlySigned {
		return nil, errors.New("The signature is not valid")
	}

	// Check if there is an event showing that the client
	// has purchased the information.
	filter := map[string]interface{}{"Addr": account, "Hash": hash[2:]}
	_, err = libs.ReadEventsFromBalanceContract(ethClient, "purchaseNotify", filter)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// error == nil means that there is at least one event that matched with the filter
	// Convert the hash (hex string) to [32]byte
	hashBytes32, _ := libs.HexStringToBytes32(hash[2:])

	// Check if the purchase has already been acknowledge in the blockcain
	_, err = libs.ReadEventsFromBalanceContract(ethClient, "responseNotify", filter)

	// error == nil means that there is at least one event that matched with the filter
	// error != nil means that there is not any event that matched the filter.
	// Therefore, the server has to notify the blockchain that the payment of the client
	// has been accepted
	if err != nil {
		if err.Error() != "Not match found" {
			fmt.Println(err)
			return nil, err
		}

		// Notify the blockchain of the pruchase of the client
		err = notifyBlockchain(ethClient, hashBytes32, addrFormatted)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	// Get the url where the data is stored
	url, _, err := ethClient.DataCon.RetrieveInfo(nil, hashBytes32)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Get the data from the REST server
	data, err := libs.GetDataFromRestServer(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}
