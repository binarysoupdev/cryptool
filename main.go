package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/binarysoupdev/cryptool/crypt"
	"github.com/binarysoupdev/cryptool/util"
)

const CRYPT_EXT = ".crypt"

func main() {
	file := flag.String("i", "", "")
	flag.Parse()

	key := "foobar"

	err := run(key, *file)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}

func run(key, file string) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return util.ChainError(err, "error reading input file")
	}

	if filepath.Ext(file) == CRYPT_EXT {
		return decrypt(key, bytes, file[:len(file)-len(CRYPT_EXT)])
	} else {
		return encrypt(key, bytes, file+CRYPT_EXT)
	}
}

func encrypt(key string, plaintext []byte, file string) error {
	c, err := crypt.New(key)
	if err != nil {
		return util.ChainError(err, "error creating crypt object")
	}
	ciphertext := c.Encrypt(plaintext)

	err = os.WriteFile(file, ciphertext, 0666)
	if err != nil {
		return util.ChainError(err, "error wrting encrypted file")
	}

	fmt.Printf("> %s\n", file)
	return nil
}

func decrypt(key string, ct crypt.Ciphertext, file string) error {
	c, err := crypt.Load(key, ct.Salt())
	if err != nil {
		return util.ChainError(err, "error loading crypt object")
	}

	plaintext, err := c.Decrypt(ct)
	if err != nil {
		return err
	}

	err = os.WriteFile(file, plaintext, 0666)
	if err != nil {
		return util.ChainError(err, "error wrting decrypted file")
	}

	fmt.Printf("> %s\n", file)
	return nil
}
