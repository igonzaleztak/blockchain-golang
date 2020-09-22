package libs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

/***************** Global variables *********************/

// OrionHost is the url of Orion
const OrionHost string = "http://localhost:1026/v2/entities"

/******************** Functions ************************/

// PrettyPint prints prettily a map/struct variable
func PrettyPint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", " ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

// DoRequestOrion Sends the http request to Orion
func DoRequestOrion(url string, method string, body []byte, header http.Header) (int, error) {
	// Create a new http client
	client := http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return 0, err
	}

	// Add the proper headers to the request
	req.Header.Add("Fiware-Service", header.Get("Fiware-Service"))
	req.Header.Add("Fiware-ServicePath", header.Get("Fiware-ServicePath"))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	// Send the request to the server
	resp, err := client.Do(req)
	if (err) != nil {
		return 0, errors.New("Something went wrong while sending the request to the server")
	}

	defer resp.Body.Close()

	// Get the status code of the request
	return resp.StatusCode, nil

}

// PostData sends the post request to the REST server
func PostData(header http.Header, bodyOrion map[string]interface{}) error {
	// Marshal the mapping to JSON
	jsonBody, err := json.Marshal(bodyOrion)
	if err != nil {
		return errors.New("Could not parse the request to JSON")
	}

	// Do the request to
	var statusCode int
	statusCode, err = DoRequestOrion(OrionHost, "POST", jsonBody, header)
	if err != nil {
		return err
	}

	// 422 status code indicates that the entity has already been created.
	// Therefore, the url and the method must change to append the new measurement
	if statusCode == 422 {
		// Get the new url
		var url = OrionHost + "/" + bodyOrion["id"].(string) + "/attrs"

		// Delete the id and the type from the message's body
		delete(bodyOrion, "id")
		delete(bodyOrion, "type")

		// Marshal the new mapping to JSON
		jsonBody, err := json.Marshal(bodyOrion)
		if err != nil {
			return errors.New("Could not parse the request to JSON")
		}

		// Send the PATCH method that updates the value of the measurement
		statusCode, err = DoRequestOrion(url, "PATCH", jsonBody, header)
		if err != nil {
			return err
		}
	}

	if float32(statusCode)/200 < 1 && float32(statusCode)/200 >= 1.5 {
		fmt.Println(statusCode)
		return errors.New("The status code received is not the expected one: " + string(statusCode))
	}

	return nil
}

//SendDataOrion sends the data to Orion
func SendDataOrion(header http.Header, body map[string]interface{}, hashString string) error {
	// Copy the body into a new variable.
	// This will be the body of the http request to Orion
	var bodyOrion = body

	// Add the field hash to the metadata that will be send to Orion
	var hash = make(map[string]interface{})
	var hashObjects = make(map[string]interface{})
	hashObjects["type"] = "String"
	hashObjects["value"] = hashString
	hash["hash"] = hashObjects

	bodyOrion["attributes"].(map[string]interface{})["metadata"] = hash

	// Delte the fields that are not neccessary
	delete(bodyOrion, "cipher")
	delete(bodyOrion, "description")
	delete(bodyOrion, "Fiware-Service")
	delete(bodyOrion, "Fiware-Servicepath")

	// Post the data to Orion
	var err = PostData(header, bodyOrion)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
