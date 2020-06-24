package main

import (
	"fmt"
	"log"
	"math/big"

	libs "./libs"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	accessControlContract "./contracts/accessContract"
	balanceContract "./contracts/balanceContract"
	dataContract "./contracts/dataContract"
)

/***************** Global variables *********************/

// GethPath stores the path to the IPC port of the ethereum node
const GethPath string = "/home/ivan/Desktop/demoPOA2/new-node/geth.ipc"

/******************** Functions ************************/
func main() {
	// Connect to the IPC endpoint of the Ethereum node
	client, err := ethclient.Dial(GethPath)
	if err != nil {
		log.Fatal(err)
	}

	// Get the private key of the admin user to set information in the contracts
	adminPrivateKey, err := libs.GetPrivateKey("", "")
	if err != nil {
		log.Fatal(err)
	}

	// Set the parameters of the transaction
	auth := bind.NewKeyedTransactor(adminPrivateKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	// Deploy the access control contract
	address, _, _, err := dataContract.DeployDataLedgerContract(auth, client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("DataContract: %s\n", address.Hex())

	// Deploy the access control contract
	address, _, _, err = accessControlContract.DeployAccessControlContract(auth, client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("AccessContract: %s\n", address.Hex())

	// Deploy the access control contract
	address, _, _, err = balanceContract.DeployBalanceContract(auth, client)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Balance Contract: %s\n", address.Hex())
}
