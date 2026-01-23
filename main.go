package main

import (
	"fmt"

	"github.com/binarysoupdev/cryptool/crypt"
)

func main() {
	key := "password123"

	c, err := crypt.New(key)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}

	plaintext := "foobar"
	run(c, plaintext)
	run(c, plaintext)
}

func run(c crypt.Crypt, plaintext string) {
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
