package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	accessControlContract "./contracts/accessContract"
	balanceContract "./contracts/balanceContract"
	dataContract "./contracts/dataContract"
	libs "./libs"
)

/***************** Global variables *********************/

// GethPath stores the path to the IPC port of the ethereum node
const GethPath string = "/home/ivan/Desktop/demoPOA2/new-node/geth.ipc"

// Contract Addresses
var dataContractAddress common.Address = common.HexToAddress("0x9E2E646CAa6B948EA8246BA27464B326eb4A4F4c")
var accessControlContractAddress common.Address = common.HexToAddress("0xE3329BA616Dc4f9E37A64a5B834192FAB84ab41B")
var balanceContractAddress common.Address = common.HexToAddress("0x7e8b560e144e966F445e0044f7F489E6fdB12E8A")

// EthereumLocal is a struct that holds the values attached to the blockchain
type EthereumLocal libs.Ethereum

/******************** Functions ************************/

// Init initializes the ehtereum blockchain and its contracts
func Init() *EthereumLocal {

	// Connect to the IPC endpoint of the Ethereum node
	client, err := ethclient.Dial(GethPath)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the data contract
	dataContract, err := dataContract.NewDataLedgerContract(dataContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Get the private key of the admin user to set information in the contracts
	adminPrivateKey, err := libs.GetPrivateKey("")
	if err != nil {
		log.Fatal(err)
	}

	// Set the parameters of the transaction
	auth := bind.NewKeyedTransactor(adminPrivateKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	// Set the address of the access control contract in the data contract
	_, err = dataContract.SetAddress(auth, accessControlContractAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the accessControlContract
	accessContract, err := accessControlContract.NewAccessControlContract(accessControlContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the balanceContract
	balanceContract, err := balanceContract.NewBalanceContract(balanceContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Set the data contract address in the balance contract
	_, err = balanceContract.SetAddress(auth, dataContractAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Declaration of the struct that stores
	// the Ethereum client and the contracts
	myEthereumClient := &EthereumLocal{
		EthereumClient: client,
		DataCon:        dataContract,
		AccessCon:      accessContract,
		BalanceCon:     balanceContract,
		AdminPrivKey:   adminPrivateKey,
	}

	return myEthereumClient
}

// PurchaseData Handles the purchases of the client
func (ethclient *EthereumLocal) PurchaseData(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/buydata" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch req.Method {
	case "POST":
		// Read the body of the message
		var body map[string]interface{}
		err := json.NewDecoder(req.Body).Decode(&body)

		if err != nil {
			http.Error(w, "401 Could not process the purchase due to error: "+err.Error(), http.StatusBadRequest)
		}

	default:
		http.Error(w, "401 Only POST methods are supported", http.StatusBadRequest)
		fmt.Fprintf(w, "Only Post methods are supported")
	}

}

//EventListener listens for new events on /notify and process them
func (ethClient *EthereumLocal) EventListener(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/notify" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch req.Method {
	case "POST":

		// Create a map with body of the message
		var body map[string]interface{}

		// Create a map with the header of the message
		header := req.Header

		// Read the body of the message
		err := json.NewDecoder(req.Body).Decode(&body)
		if err != nil {
			fmt.Println(err)
		}

		// Get the ID of the producer from the body
		producerID := body["attributes"].(map[string]interface{})["value"].(map[string]interface{})["sensorID"].(map[string]interface{})["value"].(string)

		// Get the cipherParams from the body
		cipherParams := body["cipher"].(map[string]interface{})["value"].(string)

		// Check if the producer of the event has access to the blockchain
		hasAccess, producerPrivKey, err := libs.CheckAccess(
			ethClient.EthereumClient,
			ethClient.DataCon,
			ethClient.AccessCon,
			producerID,
			cipherParams)

		if err != nil {
			fmt.Println(err)
		}

		if hasAccess {
			// Prepare the data that will be stored in the blockchain
			dataBlockchain := libs.PrepareData(header, body)
			libs.PrettyPint(dataBlockchain)

			// Prepare the data that will be sent to Orion
			err = libs.SendDataOrion(header, body, dataBlockchain["hash"].(string))

			// Introduce the data in the blockchain
			ethClientArg := &libs.Ethereum{
				EthereumClient: ethClient.EthereumClient,
				DataCon:        ethClient.DataCon,
				AccessCon:      ethClient.AccessCon,
				BalanceCon:     ethClient.BalanceCon,
				AdminPrivKey:   ethClient.AdminPrivKey,
			}
			err = libs.InteractBlockchain(dataBlockchain, producerPrivKey, ethClientArg)
			fmt.Println("The event has been successfully stored in the blockchain")
		}

		// Set the producer private key to null
		producerPrivKey = nil

		if err != nil {
			http.Error(w, "401 Could not introduce the event in the blockchain due to: "+err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, "200 ok", http.StatusOK)
		}

	default:
		http.Error(w, "401 Only POST methods are supported", http.StatusBadRequest)
		fmt.Fprintf(w, "Only Post methods are supported")
	}

}

func main() {
	// Initialize the Ethreum client and its contracts
	myEthereumClient := Init()

	// Greeting line
	fmt.Println("Starting server on port 5051")

	// Start the server
	http.HandleFunc("/notify", myEthereumClient.EventListener)
	http.HandleFunc("/buydata", myEthereumClient.PurchaseData)
	err := http.ListenAndServe(":5051", nil)
	if err != nil {
		log.Fatal("ListenAndServe ", err)
	}
}
