package main

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"log"

	ecies "github.com/ecies/go"

	libs "./libs"
)

var encrypted = "044feb639c52f3fc6852fc59224e96e47536ce17599b2e20f045a5625c8bbbf355bed7f923ac6bed903f7f6e3c51eb7e6ece30269b8da344d0f58e66a17348a843c18e392fec354df31a2188daaf233b15ba53d8a9ecf8db1e96fbd06620ce636bd0"

/*
func EncryptWithPublicKey(pubKeyBytes ecdsa.PublicKey, msg string) (encryptedText []byte, err error) {
	// Convert the public key to library format
	pubKey, err := secp256k1.ParsePubKey(elliptic.Marshal(pubKeyBytes.Curve, pubKeyBytes.X, pubKeyBytes.Y))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Encrypt the message with the public key
	encryptedMsg, err := secp256k1.Encrypt(pubKey, []byte(msg))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return encryptedMsg, nil
}
*/
/*
func DecryptWithPrivateKey(pkBytes *ecdsa.PrivateKey, ciphertext []byte) (plainText []byte, err error) {
	// Convert the private key to the format of the libary
	privKey := secp256k1.PrivKeyFromBytes(pkBytes.D.Bytes())

	// Decipher the cipherText
	plaintext, err := secp256k1.Decrypt(privKey, ciphertext)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return plaintext, nil
}
*/

func DecryptWithPrivateKey(pkBytes *ecdsa.PrivateKey, ct []byte) ([]byte, error) {

	privKey := ecies.NewPrivateKeyFromBytes(pkBytes.D.Bytes())
	plainText, err := ecies.Decrypt(privKey, ct)
	if err != nil {
		return nil, err
	}

	return plainText, nil
}

func main() {
	// Address of the account
	var address string = "c9585bb6fb2521f3630b2be104e43db394fd1db4"

	// Get the private key of the previous address
	pkBytes, err := libs.GetPrivateKey(address)

	// Message that will be ciphered
	//var msg string = "1"

	// Get the public key from the private key
	//pubKeyBytes := pkBytes.PublicKey

	/*
		// Encrypt the message with the public key
		encryptedMsg, err := encryptWithPublicKey(pubKeyBytes, msg)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hex.EncodeToString(encryptedMsg))
	*/

	ct, err := hex.DecodeString(encrypted)
	// Decrypt the message with the private key
	plainText, err := DecryptWithPrivateKey(pkBytes, ct)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(plainText))

}
