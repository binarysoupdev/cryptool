package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/binarysoupdev/cryptool/util"
)

const CRYPT_EXT = ".crypt"

func Run(file string, remove bool) error {
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
		fmt.Printf("[-] %s\n", file)
	}
	return nil
}
