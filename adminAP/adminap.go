package adminap

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	libs "../libs"
)

// BodyStruct struct used to store the body of the message that the admin sends
type BodyStruct struct {
	Action       string
	ProducerAddr string
	ProducerID   string
	Hash         string
	Price        uint
}

// Set the parameters of the admin's transaction
func setParamsTransactions(privKey *ecdsa.PrivateKey) *bind.TransactOpts {

	// Set the parameters of the transaction
	auth := bind.NewKeyedTransactor(privKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	return auth
}

// addProducer grants access to a prodcuer, so It can access the blockchain
func addProducer(b BodyStruct, ethClient *libs.Ethereum) error {
	// Convert the arguments from string to the proper format
	id := libs.String2Bytes32(b.ProducerID)
	addr := common.HexToAddress(b.ProducerAddr)

	// Set the producer in the blockchain
	auth := setParamsTransactions(ethClient.AdminPrivKey)
	_, err := ethClient.AccessCon.AddAccountToRegister(auth, id, addr)
	if err != nil {
		return err
	}

	return nil
}

// Removes a producer from the blockchain
func removeProducer(b BodyStruct, ethClient *libs.Ethereum) error {
	// Convert the arguments from string to the proper format
	id := libs.String2Bytes32(b.ProducerID)

	// Remove the producer from the blockchain
	auth := setParamsTransactions(ethClient.AdminPrivKey)
	_, err := ethClient.AccessCon.RemoveAccountFromRegister(auth, id)
	if err != nil {
		return err
	}
	return nil
}

// setPrice sets the price of a measurement
func setPrice(b BodyStruct, ethClient *libs.Ethereum) error {
	// Convert the arguments to the proper format
	hash := libs.ByteToByte32(common.Hex2Bytes(b.Hash[2:]))
	price := big.NewInt(int64(b.Price))

	// Set the price of the measurement in the blockchain
	auth := setParamsTransactions(ethClient.AdminPrivKey)
	_, err := ethClient.BalanceCon.SetPriceData(auth, hash, price)
	if err != nil {
		return err
	}
	return nil
}

// ProcessRequest processes the request of the admin
func ProcessRequest(req *http.Request, ethClient *libs.Ethereum) error {

	// Parse the JSON body
	var body BodyStruct
	err := json.NewDecoder(req.Body).Decode(&body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Check the action field to see what to do with the request
	switch body.Action {
	case "addProducer":
		// Add the producer to the blockchain
		err := addProducer(body, ethClient)
		if err != nil {
			fmt.Println(err)
			return errors.New("Could not store the producer in the blockchain")
		}
	case "removeProducer":
		// Remove the producer from the blockchain
		err := removeProducer(body, ethClient)
		if err != nil {
			fmt.Println(err)
			return errors.New("Could not remove the producer from the blockchain")
		}
	case "setPrice":
		// Set the price of a product
		err := setPrice(body, ethClient)
		if err != nil {
			fmt.Println(err)
			return errors.New("Could not set the price of the measurement")
		}
	default:
		return errors.New("Wrong action")
	}

	return nil
}
