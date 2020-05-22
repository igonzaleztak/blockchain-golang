package libs

import (
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

/***************** Global variables *********************/

// RestHost URL of the rest server that queries the database
const RestHost string = "https://localhost:8052/"

// Authentication params
const user string = "cygnus"
const pass string = "telematica"

/******************** Functions ************************/

// GetDataFromRestServer Gets the data from the REST server
func GetDataFromRestServer(url string) ([]byte, error) {

	// Get the authentication params in base 64
	authParams := base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
	fmt.Println(authParams)

	// Create new Transport that ignores self-signed SSL
	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := http.Client{Transport: customTransport}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Add the basic authentication header
	req.Header.Add("Authorization", "Basic "+authParams)

	// Send the request to the server
	resp, err := client.Do(req)
	if (err) != nil {
		fmt.Println(err)
		return nil, errors.New("Something went wrong while sending the request to the server")
	}

	defer resp.Body.Close()

	// Decode the body
	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respBody))

	// Get the status code of the request
	if resp.StatusCode != 200 {
		fmt.Println(string(respBody))
		return nil, errors.New("Status code is not 200 Ok")
	}

	return respBody, nil
}
