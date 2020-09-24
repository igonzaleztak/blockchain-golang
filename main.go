package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	adminap "./adminAP"
	clientap "./clientAP"
	accessControlContract "./contracts/accessContract"
	balanceContract "./contracts/balanceContract"
	dataContract "./contracts/dataContract"
	libs "./libs"
	logger "./logger"
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
		logger.Log.Fatal(err)
	}

	// Initialize the data contract
	dataContract, err := dataContract.NewDataLedgerContract(libs.DataContractAddress, client)
	if err != nil {
		logger.Log.Fatal(err)
	}

	// Get the private key of the admin user to set information in the contracts
	adminPrivateKey, err := libs.GetPrivateKey("", "")
	if err != nil {
		logger.Log.Fatal(err)
	}

	// Set the parameters of the transaction
	auth := bind.NewKeyedTransactor(adminPrivateKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	// Set the address of the access control contract in the data contract
	_, err = dataContract.SetAddress(auth, libs.AccessControlContractAddress)
	if err != nil {
		logger.Log.Fatal(err)
	}

	// Initialize the accessControlContract
	accessContract, err := accessControlContract.NewAccessControlContract(libs.AccessControlContractAddress, client)
	if err != nil {
		logger.Log.Fatal(err)
	}

	// Initialize the balanceContract
	balanceContract, err := balanceContract.NewBalanceContract(libs.BalanceContractAddress, client)
	if err != nil {
		logger.Log.Fatal(err)
	}

	// Set the data contract address in the balance contract
	_, err = balanceContract.SetAddress(auth, libs.DataContractAddress)
	if err != nil {
		logger.Log.Fatal(err)
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

		logger.Log.Printf("Request received from %s\n", req.Host)

		// Create a map with body of the message
		var auxBody map[string]interface{}

		// Create a map with the header of the message
		header := req.Header

		// Read the body of the message
		err := json.NewDecoder(req.Body).Decode(&auxBody)
		if err != nil {
			fmt.Println(err)
		}

		logger.Log.Printf("Header of the request: \n%s \n", fmt.Sprintln(header))
		logger.Log.Printf("Body of the request: \n%s \n", fmt.Sprintln(auxBody))

		// Insert all the measurements into a field
		// called "attributes"
		body := libs.PrepareInputData(auxBody)

		logger.Log.Printf("Formatted Body:\n%s \n", fmt.Sprintln(body))

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

		if hasAccess && err == nil {
			// Prepare the data that will be stored in the blockchain
			dataBlockchain := libs.PrepareData(header, body)
			libs.PrettyPint(dataBlockchain)
			logger.Log.Printf("Data that will be stored in the Blockchain: \n%s\n", fmt.Sprintln(dataBlockchain))
			// Prepare the data that will be sent to Orion
			err = libs.SendDataOrion(header, body, dataBlockchain["hash"].(string))
			if err != nil {
				logger.Log.Printf("%s\n", fmt.Sprintln(err))
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
			if err != nil {
				fmt.Printf("\n")
				fmt.Println(err)
				fmt.Printf("\n")
				http.Error(w, "401 Could not introduce the event in the blockchain due to: "+err.Error(), http.StatusBadRequest)
				return
			}
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

// AdminAP is an access point that the admin user has to introduce new producers
// in the blockchain
func (ethclient *EthereumLocal) AdminAP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/adminap" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	switch req.Method {
	case "POST":

		logger.Log.Printf("Request received from %s\n", req.Host)

		// check the authentication header value
		username, password, ok := req.BasicAuth()
		if !ok {
			http.Error(w, "There is no authentication header in the request", http.StatusBadRequest)
			logger.Log.Printf("Could not read the authentication header of the request\n")
			return
		}
		if username != libs.ADMINACCOUNT || password != libs.ADMINPASSWORD {
			logger.Log.Printf("Wrong user or pasword")
			http.Error(w, "Wrong password or user", http.StatusBadRequest)
		}

		// Process the request of the admin user
		ethClientArg := &libs.Ethereum{
			EthereumClient: ethclient.EthereumClient,
			DataCon:        ethclient.DataCon,
			AccessCon:      ethclient.AccessCon,
			BalanceCon:     ethclient.BalanceCon,
			AdminPrivKey:   ethclient.AdminPrivKey,
		}
		err := adminap.ProcessRequest(req, ethClientArg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "200 ok", http.StatusOK)

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

	// Log
	logger.Log.Printf("Starting HTTP server on port 5051\n")
	logger.Log.Printf("Starting HTTP server on port 8051\n")
	logger.Log.Printf("---------------------------------\n\n")

	// Start the server
	http.HandleFunc("/notify", myEthereumClient.EventListener)
	http.HandleFunc("/buydata", myEthereumClient.PurchaseData)
	http.HandleFunc("/adminap", myEthereumClient.AdminAP)

	// HTTP interface in a new subroutine
	go func() {
		err := http.ListenAndServe(":5051", nil)
		if err != nil {
			logger.Log.Fatal(err)
		}
	}()

	// HTTPS interface
	err := http.ListenAndServeTLS(":8051", "./certs/public-cert.pem", "./certs/private-key.pem", nil)
	if err != nil {
		logger.Log.Fatal(err)
	}

}
