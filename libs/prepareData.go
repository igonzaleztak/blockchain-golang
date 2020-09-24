package libs

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"strings"
)

/***************** Global variables *********************/
var restServerHost string = "https://localhost:8052/"

/******************** Functions ************************/

// ConvertStringArrayToString converts []string to string
func ConvertStringArrayToString(stringArray []string) string {
	return strings.Join(stringArray, " ")
}

// HashData hashes the measurement without the fields cipher and description
// The hash is SHA256(service, servicePath, body)
func HashData(body map[string]interface{}, fiwareService []string, fiwareServicePath []string) string {
	// Copy the body into another structure
	bodyAux := body

	// Delete the fields cipher and description from this structure
	delete(bodyAux, "cipher")
	delete(bodyAux, "description")

	// Include the heades Fiware-Service and Fiware-Servicepath
	bodyAux["Fiware-Service"] = ConvertStringArrayToString(fiwareService)
	bodyAux["Fiware-Servicepath"] = ConvertStringArrayToString(fiwareServicePath)

	// Parse the map to []bytes
	data, _ := json.Marshal(bodyAux)

	// Do the hash of the data
	hashBytes := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hashBytes[:])
}

// PrepareData conforms the data that will be stored in the blockchain
func PrepareData(header http.Header, body map[string]interface{}) map[string]interface{} {

	// Get the description of the event
	description := body["description"].(map[string]interface{})["value"].(string)

	// Get the id of the sensor
	id := body["attributes"].(map[string]interface{})["value"].(map[string]interface{})["sensorID"].(map[string]interface{})["value"].(string)

	// Get the Fiware-Service from the header
	fiwareService := header["Fiware-Service"]

	// Get the Fiware-Servicepath
	fiwareServicePath := header["Fiware-Servicepath"]

	// Get the hash associated to the event
	hash := HashData(body, fiwareService, fiwareServicePath)

	// Get the URL used to access the data stored in the REST server
	url := strings.ToLower(restServerHost +
		ConvertStringArrayToString(fiwareService) + ConvertStringArrayToString(fiwareServicePath) + "/" + hash)

	// Create a map that holds the data that will be stored in the blockchain
	var dataToBlockchain = make(map[string]interface{})
	dataToBlockchain["hash"] = hash
	dataToBlockchain["url"] = url
	dataToBlockchain["description"] = description
	dataToBlockchain["id"] = id

	return dataToBlockchain
}
