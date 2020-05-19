package clientap

import (
	"errors"

	libs "../libs"
)

// EthereumLocal is a struct that holds the values attached to the blockchain
type EthereumLocal libs.Ethereum

// ProccessClientPurchase Processes the purchases of the client
// and returns the measurements to him
func ProccessClientPurchase(ethClient *EthereumLocal, bodyReq map[string]interface{}) error {

	return errors.New("ALGO")
}
