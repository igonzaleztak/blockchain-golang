package libs

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	accessControlContract "../contracts/accessContract"
	balanceContract "../contracts/balanceContract"
	dataContract "../contracts/dataContract"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Ethereum is a struct that hodls the values attached to the blockchain
type Ethereum struct {
	EthereumClient *ethclient.Client
	DataCon        *dataContract.DataLedgerContract
	AccessCon      *accessControlContract.AccessControlContract
	BalanceCon     *balanceContract.BalanceContract
	AdminPrivKey   *ecdsa.PrivateKey
}

//InteractBlockchain stores the measurement in the blockchain
func InteractBlockchain(
	dataBlockchain map[string]interface{},
	producerPrivateKey *ecdsa.PrivateKey,
	ethclient *Ethereum,
) error {

	// Set the parameters of the transaction
	auth := bind.NewKeyedTransactor(producerPrivateKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	// Send the transaction
	hash32Byte := String2Bytes32(dataBlockchain["hash"].(string))
	_, err := ethclient.DataCon.StoreInfo(auth,
		hash32Byte,
		dataBlockchain["url"].(string),
		String2Bytes32(dataBlockchain["id"].(string)),
		dataBlockchain["description"].(string),
	)

	// Set the parameters of the new transaction to set the price of the product
	auth = bind.NewKeyedTransactor(ethclient.AdminPrivKey)
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(400000)
	auth.GasPrice = big.NewInt(0)

	// Set the price of the product
	price := big.NewInt(2)
	_, err = ethclient.BalanceCon.SetPriceData(auth, hash32Byte, price)

	if err != nil {
		return err
	}
	return nil
}
