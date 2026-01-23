package main

import (
	"fmt"

	"github.com/binarysoupdev/cryptool/crypt"
)

func main() {
	plaintext := "foobar"

	c, err := crypt.New()

	ciphertext, err := c.Encrypt([]byte(plaintext))
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
	fmt.Println(ciphertext)

	bytes, err := c.Decrypt(ciphertext)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
	fmt.Println(string(bytes))
}
