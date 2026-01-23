package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/binarysoupdev/cryptool/crypt"
	"github.com/binarysoupdev/cryptool/util"
	"golang.org/x/term"
)

const CRYPT_EXT = ".crypt"

func main() {
	file := flag.String("i", "", "")
	flag.Parse()

	err := run(*file)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}

func run(file string) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return util.ChainError(err, "error reading input file")
	}

	fmt.Println("Enter PASSWORD:")
	key, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return util.ChainError(err, "error reading password from terminal")
	}

	if filepath.Ext(file) == CRYPT_EXT {
		return decrypt(string(key), bytes, file[:len(file)-len(CRYPT_EXT)])
	} else {
		return encrypt(string(key), bytes, file+CRYPT_EXT)
	}
}

func encrypt(key string, plaintext []byte, file string) error {
	fmt.Println("Verify PASSWORD:")
	verify, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return util.ChainError(err, "error reading password from terminal")
	}

	if string(verify) != key {
		return errors.New("passwords do not match")
	}

	c, err := crypt.New(key)
	if err != nil {
		return util.ChainError(err, "error creating crypt object")
	}
	ciphertext := c.Encrypt(plaintext)

	err = os.WriteFile(file, ciphertext, 0666)
	if err != nil {
		return util.ChainError(err, "error writing encrypted file")
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
		return util.ChainError(err, "error writing decrypted file")
	}

	fmt.Printf("> %s\n", file)
	return nil
}
