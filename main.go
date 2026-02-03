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
	flag.Usage = func() {
		fmt.Println("Simple cryptography tool to encrypt/decrypt a file with a password.")
		flag.PrintDefaults()
	}

	file := flag.String("i", "", "the file to encrypt/decrypt")
	rm := flag.Bool("rm", false, "remove the old file")
	flag.Parse()

	err := run(*file, *rm)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}

func run(file string, remove bool) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return util.ChainError(err, "error reading input file")
	}

	if filepath.Ext(file) == CRYPT_EXT {
		err = decrypt(bytes, file[:len(file)-len(CRYPT_EXT)])
	} else {
		err = encrypt(bytes, file+CRYPT_EXT)
	}
	if err != nil {
		return err
	}

	if remove {
		os.Remove(file)
		fmt.Printf("[-] %s\n", file)
	}
	return nil
}

func encrypt(in []byte, out string) error {
	password := promptPassword("NEW")
	verify := promptPassword("VERIFY")

	if verify != password {
		return errors.New("passwords do not match")
	}

	salt, ciphertext := crypt.New(password).Encrypt(in)

	err := os.WriteFile(out, append(salt, ciphertext...), 0666)
	if err != nil {
		return util.ChainError(err, "error writing encrypted file")
	}

	fmt.Printf("[+] %s\n", out)
	return nil
}

func decrypt(in []byte, out string) error {
	password := promptPassword("ENTER")

	plaintext, err := crypt.Load(password, in[:crypt.SALT_SIZE]).Decrypt(in[crypt.SALT_SIZE:])
	if err != nil {
		return err
	}

	err = os.WriteFile(out, plaintext, 0666)
	if err != nil {
		return util.ChainError(err, "error writing decrypted file")
	}

	fmt.Printf("[+] %s\n", out)
	return nil
}

func promptPassword(prompt string) string {
	fmt.Printf("%s PASSWORD:\n", prompt)

	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	return string(password)
}
