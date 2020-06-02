package test

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"log"

	ecies "github.com/ecies/go"

	libs "./libs"
)

//var encrypted = "044feb639c52f3fc6852fc59224e96e47536ce17599b2e20f045a5625c8bbbf355bed7f923ac6bed903f7f6e3c51eb7e6ece30269b8da344d0f58e66a17348a843c18e392fec354df31a2188daaf233b15ba53d8a9ecf8db1e96fbd06620ce636bd0"

func decryptWithPrivateKey(pkBytes *ecdsa.PrivateKey, ct []byte) ([]byte, error) {

	privKey := ecies.NewPrivateKeyFromBytes(pkBytes.D.Bytes())
	plainText, err := ecies.Decrypt(privKey, ct)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

// EncryptWithPublicKey encrypts a message using ECIES
func encryptWithPublicKey(pubKeyECDSA ecdsa.PublicKey, msg []byte) ([]byte, error) {
	// Convert the public key to the appropiate format
	var pubKeyBytes []byte = elliptic.Marshal(pubKeyECDSA.Curve, pubKeyECDSA.X, pubKeyECDSA.Y)
	pubKey, err := ecies.NewPublicKeyFromBytes(pubKeyBytes)
	if err != nil {
		return nil, err
	}

	// Encrypt the message with the public key
	cipherText, err := ecies.Encrypt(pubKey, msg)
	if err != nil {
		return nil, err
	}

	return cipherText, nil
}

func main() {
	// Address of the account
	//var address string = "21A018606490C031A8c02Bb6b992D8AE44ADD37f"

	// Get the private key of the previous address
	pkBytes, err := libs.GetPrivateKey("", "")

	// Message that will be ciphered
	var msg string = "1"

	// Get the public key from the private key
	pubKeyBytes := pkBytes.PublicKey

	// Encrypt the message with the public key
	encryptedMsg, err := encryptWithPublicKey(pubKeyBytes, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(encryptedMsg))

	// Decrypt the message with the private key
	plainText, err := decryptWithPrivateKey(pkBytes, encryptedMsg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(plainText))

}
