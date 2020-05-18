package cipherlib

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"

	"github.com/decred/dcrd/dcrec/secp256k1"
)

// EncryptWithPublicKey encrypts a message using ECDH
func EncryptWithPublicKey(pubKeyBytes ecdsa.PublicKey, msg string) ([]byte, error) {
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

// DecryptWithPrivateKey Decrypts a text encrypted using ECDH
func DecryptWithPrivateKey(pkBytes *ecdsa.PrivateKey, ciphertext []byte) ([]byte, error) {
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
