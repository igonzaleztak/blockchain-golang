package clientap

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

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

// ResponseSubroutine handles the responses of the goroutines
type ResponseSubroutine struct {
	txHash []byte
	Error  error
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

// sendTransactionSecretRoutine ciphers  the secret with the public key of
// the admin and the public key of the client. Then, it sends the
// encrypted secret in a transaction
func sendTransactionSecretRoutine(response chan ResponseSubroutine,
	ethClient *libs.Ethereum,
	randomKey []byte,
	clientPubKeyBytes []byte,
	addrFormatted common.Address) {

	// Cipher the key with the public key of the admin
	encryptedAdmin, err := cipher.EncryptWithPublicKey(ethClient.AdminPrivKey.PublicKey, randomKey)
	if err != nil {
		fmt.Println(err)
		response <- ResponseSubroutine{nil, err}
		return
	}

	// Cipher with the public key of the client
	// Convert the public key of the client to ECDSA format
	clientPubKeyECDSA, err := crypto.UnmarshalPubkey(clientPubKeyBytes)
	if err != nil {
		fmt.Println(err)
		response <- ResponseSubroutine{nil, err}
		return
	}
	encryptedClient, err := cipher.EncryptWithPublicKey(*clientPubKeyECDSA, randomKey)

	// Prepare the body of the transaction in JSON format
	transactionKeysData := &SecretStruct{AdminSecret: encryptedAdmin, ClientSecret: encryptedClient}
	transactionKeysDataJSON, err := json.Marshal(transactionKeysData)
	if err != nil {
		fmt.Println(err)
		response <- ResponseSubroutine{nil, err}
		return
	}

	// Send the transaction with the secret
	txSecretHash, err := libs.SendTransaction(common.HexToAddress(libs.ADMINACCOUNT),
		addrFormatted,
		ethClient.AdminPrivKey,
		ethClient,
		transactionKeysDataJSON)
	if err != nil {
		fmt.Println(err)
		response <- ResponseSubroutine{nil, err}
		return
	}

	// Prepare the response
	response <- ResponseSubroutine{txSecretHash, nil}
}

// ProccessClientPurchase Processes the purchases of the client
// and returns the measurements to him
func ProccessClientPurchase(ethClient *libs.Ethereum, body map[string]interface{}) ([]byte, error) {
	// Get the hash that the client wants to purchase
	var hash string = body["_hash"].(string)

	// Get the account that made the purchase
	var account string = body["_account"].(string)

	// address formatted
	addrFormatted := common.HexToAddress(account[2:])

	// Get the signature
	var signature string = body["_signature"].(string)

	// Get the hash that was signed
	signedMessage, err := getSignedMessage(account, hash)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Convert the signature to bytes
	signatureBytes, err := hex.DecodeString(signature[2:])
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Get the publicKey of the client
	clientPubKey, err := ethClient.AccessCon.GetPubKey(nil, addrFormatted)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Convert the public key to []byte
	clientPubKeyBytes, err := hex.DecodeString("04" + clientPubKey)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Verify the signature
	var isProperlySigned bool = cipher.VerifySignature(clientPubKeyBytes, signedMessage, signatureBytes)
	if !isProperlySigned {
		return nil, errors.New("Message is not properly signed")
	}
	// Check if there is an event showing that the client
	// has purchased the information.
	filter := map[string]interface{}{"Addr": account, "Hash": hash[2:]}
	_, err = libs.ReadEventsFromBalanceContract(ethClient, "purchaseNotify", filter)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// error != nil means that there is at least one event that matched with the filter
	// Check if the data has already been given to the client:
	// 	- True:	 If the data has not been given to the client
	// 	- False: If the data has already been given to the client
	hashBytes32, _ := libs.HexStringToBytes32(hash[2:])
	HasToBePaid, err := ethClient.BalanceCon.CheckHasPaid(nil, addrFormatted, hashBytes32)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !HasToBePaid {
		return nil, errors.New("The measurement has already been given")
	}

	// Generate a new key that will be used in the symmetric encryption
	randomKey := make([]byte, 32)
	rand.Read(randomKey)

	// Declare the channels used in the goroutines
	chSecret := make(chan ResponseSubroutine)

	// Init the the subroutine that ciphers the secret key
	go sendTransactionSecretRoutine(chSecret, ethClient, randomKey, clientPubKeyBytes, addrFormatted)

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

	// Encrypt the measurement with the symmetric key
	ciphertext, err := cipher.SymmetricEncryption(randomKey, data)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Check if the encryption works
	//_ = cipher.DecryptSymmetricEncryption(randomKey, ciphertext)

	// Send the transaction with the encrypted measurement
	txCiphertextHash, err := libs.SendTransaction(common.HexToAddress(libs.ADMINACCOUNT),
		addrFormatted,
		ethClient.AdminPrivKey,
		ethClient,
		ciphertext)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Read the responses from the channels
	responseSecret := <-chSecret
	txSecretHash := responseSecret.txHash
	if responseSecret.Error != nil {
		fmt.Println(responseSecret.Error)
		return nil, responseSecret.Error
	}

	// Close the channel
	close(chSecret)

	// Indicate in the contract balance that the data has been sent
	// Set the parameters of the new transaction to set the price of the product
	auth := bind.NewKeyedTransactor(ethClient.AdminPrivKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	_, err = ethClient.BalanceCon.SendToClient(auth,
		addrFormatted,
		hashBytes32,
		libs.ByteToByte32(txSecretHash),
		libs.ByteToByte32(txCiphertextHash))

	// Prepare the response that will be send to the client
	response := &ResponseClient{TxSecretHash: txSecretHash, TxDataHash: txCiphertextHash}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return responseJSON, nil

}
