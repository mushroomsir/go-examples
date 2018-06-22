package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"fmt"
	"log"
	"time"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Printf("rsa.GenerateKey: %v\n", err)
	}

	messageBytes := bytes.Repeat([]byte{1}, 87)
	sha256 := sha1.New()

	start := time.Now()
	encrypted, err := rsa.EncryptOAEP(sha256, rand.Reader, &privateKey.PublicKey, messageBytes, nil)
	if err != nil {
		fmt.Printf("EncryptOAEP: %s\n", err)
	}
	log.Println(time.Since(start))

	decrypted, err := rsa.DecryptOAEP(sha256, rand.Reader, privateKey, encrypted, nil)
	if err != nil {
		fmt.Printf("decrypt: %s\n", err)
	}

	if decrypted != nil {

	}

	// decryptedString := bytes.NewBuffer(decrypted).String()
	// fmt.Printf("message: %v\n", message)
	// fmt.Printf("encrypted: %v\n", encrypted)
	// fmt.Printf("decryptedString: %v\n", decryptedString)
}
