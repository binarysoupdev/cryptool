package main

import (
	"fmt"

	"github.com/binarysoupdev/cryptool/crypt"
)

func main() {
	plaintext := "foobar"

	ciphertext, err := crypt.Encrypt([]byte(plaintext))
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
	fmt.Println(ciphertext)

	bytes, err := crypt.Decrypt(ciphertext)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
	fmt.Println(string(bytes))
}
