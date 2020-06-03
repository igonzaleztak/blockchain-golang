package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	clientap "./clientAP"
	accessControlContract "./contracts/accessContract"
	balanceContract "./contracts/balanceContract"
	dataContract "./contracts/dataContract"
	libs "./libs"
)

/***************** Global variables *********************/

// GethPath stores the path to the IPC port of the ethereum node
const GethPath string = "/home/ivan/Desktop/demoPOA2/new-node/geth.ipc"

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
	dataContract, err := dataContract.NewDataLedgerContract(libs.DataContractAddress, client)
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

	// Set the address of the access control contract in the data contract
	_, err = dataContract.SetAddress(auth, libs.AccessControlContractAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Deploy contract example
	//address, _, accessContract, err := accessControlContract.DeployAccessControlContract(auth, client)
	//fmt.Println(address.Hex())

	// Initialize the accessControlContract
	accessContract, err := accessControlContract.NewAccessControlContract(libs.AccessControlContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize the balanceContract
	balanceContract, err := balanceContract.NewBalanceContract(libs.BalanceContractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// Set the data contract address in the balance contract
	_, err = balanceContract.SetAddress(auth, libs.DataContractAddress)
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
			http.Error(w, "Could not parse the message", http.StatusBadRequest)
		}

		// Process the purchase
		ethClientArg := &libs.Ethereum{
			EthereumClient: ethclient.EthereumClient,
			DataCon:        ethclient.DataCon,
			AccessCon:      ethclient.AccessCon,
			BalanceCon:     ethclient.BalanceCon,
			AdminPrivKey:   ethclient.AdminPrivKey,
		}
		response, err := clientap.ProccessClientPurchase(ethClientArg, body)
		if err != nil {
			if err.Error() == "The measurement has already been given" {
				http.Error(w, err.Error(), 415)
			} else {
				http.Error(w, "401 Could not process the purchase due to error: "+err.Error(), http.StatusBadRequest)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)

	default:
		http.Error(w, "401 Only POST methods are supported", http.StatusBadRequest)
		fmt.Fprintf(w, "Only Post methods are supported")
	}

}

//EventListener listens for new events on /notify and process them
func (ethclient *EthereumLocal) EventListener(w http.ResponseWriter, req *http.Request) {
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
			ethclient.EthereumClient,
			ethclient.AdminPrivKey,
			ethclient.DataCon,
			ethclient.AccessCon,
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
			if err != nil {
				fmt.Println(err)
				http.Error(w, "401 Could not introduce the event in the blockchain due to: "+err.Error(), http.StatusBadRequest)
			}

			// Introduce the data in the blockchain
			ethClientArg := &libs.Ethereum{
				EthereumClient: ethclient.EthereumClient,
				DataCon:        ethclient.DataCon,
				AccessCon:      ethclient.AccessCon,
				BalanceCon:     ethclient.BalanceCon,
				AdminPrivKey:   ethclient.AdminPrivKey,
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
	fmt.Println("Starting HTTP server on port 5051")
	fmt.Println("Starting HTTPS server on port 8051")

	// Start the server
	http.HandleFunc("/notify", myEthereumClient.EventListener)
	http.HandleFunc("/buydata", myEthereumClient.PurchaseData)

	// HTTP interface in a new subroutine
	go func() {
		err := http.ListenAndServe(":5051", nil)
		if err != nil {
			log.Fatal("ListenAndServe ", err)
		}
	}()

	// HTTPS interface
	err := http.ListenAndServeTLS(":8051", "./certs/public-cert.pem", "./certs/private-key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe ", err)
	}

}
