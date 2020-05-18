/*
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/decred/dcrd/dcrec/secp256k1"

	libs "./libs"
)

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

func main() {
	// Address of the account
	var address string = "c9585bb6fb2521f3630b2be104e43db394fd1db4"

	// Get the private key of the previous address
	var pkBytes *ecdsa.PrivateKey = libs.GetPrivateKey(address)

	// Message that will be ciphered
	var msg string = "1"

	// Get the public key from the private key
	pubKeyBytes := pkBytes.PublicKey

	// Encrypt the message with the public key
	encryptedMsg, err := encryptWithPublicKey(pubKeyBytes, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(encryptedMsg))

	// Decrypt the message with the private key
	plainText, err := DecryptWithPrivateKey(pkBytes, encryptedMsg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(plainText))

}
*/

package main2
