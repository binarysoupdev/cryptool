package app

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/binarysoupdev/cryptool/util"
	"github.com/binarysoupdev/got-style/style"
)

const CRYPT_EXT = ".crypt"

// Runs the app. This tool will encrypt/decrypt a file on disk using a password collected from stdin.
//
// 'file' is the path to the plaintext or ciphertext file.
// If the extension is ".crypt", the app runs decryption; else encryption is used.
//
// 'remove' indicates if the original file should be removed after completion.
func Run(file string, remove bool) error {
	if file == "" {
		return errors.New("filepath cannot be empty")
	}

	bytes, err := os.ReadFile(file)
	if err != nil {
		return util.ChainError(err, "error reading input file")
	}

	if filepath.Ext(file) != CRYPT_EXT {
		err = encrypt(bytes, file+CRYPT_EXT)
	} else {
		err = decrypt(bytes, file[:len(file)-len(CRYPT_EXT)])
	}
	if err != nil {
		return err
	}

	if remove {
		os.Remove(file)
		style.Delete.PrintF("[-] %s\n", file)
	}
	return nil
}
