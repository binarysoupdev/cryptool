package main

import (
	"fmt"
	"log"

	"github.com/binarysoupdev/cryptool/crypt"
)

func main() {
	c := crypt.New()

	ciphertext := c.Encrypt([]byte("lorem ipsum"))
	fmt.Println(ciphertext)

	plaintext, err := c.Decrypt(ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(plaintext))
}
