# Deploy-contracts.go
This script deploys the smart contracts in the blockchain located in the folder *contracts*. 

To deploy the contracts, they have to be compiled first. To do so, use the following order:

```bash
./contracts/compile.sh
```

Once the contracts are deployed, the next step is to connect to the Ethereum's node. This can be done with the following lines of code:

```go
// GethPath stores the path to the IPC port of the ethereum node
const GethPath string = "/home/ivan/Desktop/demoPOA2/new-node/geth.ipc"

// Connect to the IPC endpoint of the Ethereum node
client, err := ethclient.Dial(GethPath)
if err != nil {
  log.Fatal(err)
}
```

The next step is to obtain the private key of the user who will deploy the contracts. This is done using the function *GetPrivateKey*.

```go
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
```

Then, using the go interfaces that were generated during the compilation, the contracts are deployed as follows:


```go
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
```

As a result of executing the code, the console displays the addresses where the contracts were deployed.

```bash
$ go run deploy-contracts.go 
DataContract: 0x0E9deb29A823585598188f3ef2dFA4E6047C314A
AccessContract: 0x0BA6aBE9a2F2EA1657848397Eda9076C1C6a36FA
Balance Contract: 0xF59c2624af8359F39020B4579504F81e9B2d835C
```