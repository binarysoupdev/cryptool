package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"

	"github.com/binarysoupdev/cryptool/crypt"
)

func main() {
	cryptLoop()
	//keyGen()
}

func keyGen() {
	key := make([]byte, 32) // 256-bit
	rand.Read(key)

	os.Stdout.Write(key)
}

func cryptLoop() {
	key := make([]byte, 32)
	os.Stdin.Read(key)

	c := crypt.New(key)

	ciphertext := c.Encrypt([]byte("lorem ipsum"))
	fmt.Println(ciphertext)

	plaintext, err := c.Decrypt(ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(plaintext))
}
