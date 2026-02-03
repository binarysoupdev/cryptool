package app

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/binarysoupdev/cryptool/util"
	"github.com/binarysoupdev/got-style/style"
)

const CRYPT_EXT = ".crypt"

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
