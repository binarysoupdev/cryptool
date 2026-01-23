package main

import (
	"errors"
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
	c, err := crypt.New(key)
	if err != nil {
		return util.ChainError(err, "error creating crypt object")
	}

	bytes, err := os.ReadFile(file)
	if err != nil {
		return util.ChainError(err, "error reading input file")
	}

	if filepath.Ext(file) == CRYPT_EXT {
		return errors.New("DECRYPT not implemented")
	} else {
		return encrypt(c, bytes, file+CRYPT_EXT)
	}
}

func encrypt(c crypt.Crypt, plaintext []byte, file string) error {
	ciphertext, err := c.Encrypt(plaintext)
	if err != nil {
		return util.ChainError(err, "error encrypting plaintext")
	}

	err = os.WriteFile(file, ciphertext, 0666)
	if err != nil {
		return util.ChainError(err, "error wrting encrypted file")
	}

	fmt.Printf("+ %s\n", file)
	return nil
}
