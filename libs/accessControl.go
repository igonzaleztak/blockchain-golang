package libs

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	cipherlib "../cipherLibs"

	accessControlContract "../contracts/accessContract"
	dataContract "../contracts/dataContract"
)

/***************** Global variables *********************/

// ADMINACCOUNT address of the admin account
const ADMINACCOUNT string = "21A018606490C031A8c02Bb6b992D8AE44ADD37f"

// ADMINPASSWORD Ethereum passoword of the admin account
const ADMINPASSWORD string = "1"

// FOLDER where the private keys of the accounts are stored
var FOLDER string = FolderPath

/******************** Functions ************************/

//String2Bytes32 gets a string and converts it to [32]byte
func String2Bytes32(str string) [32]byte {
	aux := [32]byte{}
	copy(aux[:], []byte(str))
	return aux
}

// GetAddrAdminAccount returns the address of the admin account
// so it can be used in othre packages
func GetAddrAdminAccount() common.Address {
	return common.HexToAddress("0x" + ADMINACCOUNT)
}

// GetUTCFile search the UTC file associated to the address
func GetUTCFile(address string) (string, error) {
	// Compile the regex expresion
	libRegEx, err := regexp.Compile("(?i).*" + address)
	if err != nil {
		log.Fatal(err)
	}

	// Get all the files of the folder
	files, err := ioutil.ReadDir(FOLDER)
	if err != nil {
		log.Fatal(err)
	}

	// Get the name of the file that matches the expression
	for _, f := range files {
		if libRegEx.MatchString(f.Name()) {
			return FOLDER + f.Name(), nil
		}
	}

	return "", errors.New("UTC File not found")
}

// GetPrivateKey gets the private key of the address
func GetPrivateKey(address, password string) (*ecdsa.PrivateKey, error) {
	// Check if an address was passed to the function
	if address == "" {
		address = ADMINACCOUNT
		password = ADMINPASSWORD
	}

	// Get the UTC folder that contains the private key
	file, err := GetUTCFile(address)
	if err != nil {
		return nil, err
	}

	// Read the file
	jsonBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	// Get the private key
	keyWrapper, err := keystore.DecryptKey(jsonBytes, password)
	if err != nil {
		return nil, err
	}

	return keyWrapper.PrivateKey, nil

}

// CheckAccess checks whether the sensor that produced
// the event has access to the blockchain or not
func CheckAccess(
	client *ethclient.Client,
	adminPrivKey *ecdsa.PrivateKey,
	dataContract *dataContract.DataLedgerContract,
	accessContract *accessControlContract.AccessControlContract,
	producerID string,
	cipherText string) (bool, *ecdsa.PrivateKey, error) {

	// Get the address of the producer from the blockchain
	address, err := accessContract.GetAddress(nil, String2Bytes32(producerID))
	if err != nil {
		return false, nil, err
	}

	// Check if the address is 0
	if address == common.HexToAddress("0x0") {
		return false, nil, errors.New("There is not a producer with that address")
	}

	// Convert the address to string
	addressStr := address.Hex()

	fmt.Printf("\nEvent received from: %s\n", addressStr)

	// Decrypt the passphrase using the private key of the admin
	cipherTextBytes, err := hex.DecodeString(cipherText[2:])
	if err != nil {
		return false, nil, err
	}
	passphraseBytes, err := cipherlib.DecryptWithPrivateKey(adminPrivKey, cipherTextBytes)
	if err != nil {
		return false, nil, err
	}

	// Get the private Key of the producer account
	producerPrivKey, err := GetPrivateKey(addressStr[2:], string(passphraseBytes))
	if err != nil {
		return false, nil, err
	}

	fmt.Printf("\nThe account %s logged in the blockchain and introduced the following data\n", addressStr)
	return true, producerPrivKey, nil
}
