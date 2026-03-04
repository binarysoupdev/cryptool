package main

import (
	"bufio"
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
	line, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	c := crypt.NewFromPassword(line[:len(line)-1])

	ciphertext := c.Encrypt([]byte("lorem ipsum"))
	fmt.Println(ciphertext)

	plaintext, err := c.Decrypt(ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(plaintext))
}
