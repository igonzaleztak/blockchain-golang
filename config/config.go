package config

import (
	"encoding/json"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

// Path of the configuration file
const configPath string = "config/config.json"

// StructConfig stores the configuration of the file config.json
type StructConfig struct {
	EthereumPath           string
	DataContractAddress    common.Address
	AccessContractAddress  common.Address
	BalanceContractAddress common.Address
}

// ReadConfiguration reads the configuration file
func ReadConfiguration() StructConfig {

	// Declaration of the struct
	var configStruct StructConfig

	// Open the configuration file
	file, err := os.Open(configPath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(file).Decode(&configStruct)

	return configStruct
}
