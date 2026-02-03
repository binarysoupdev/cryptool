package app

import (
	"errors"
	"os"

	"github.com/binarysoupdev/cryptool/crypt"
	"github.com/binarysoupdev/cryptool/util"
	"github.com/binarysoupdev/got-style/style"
)

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

	style.Create.PrintF("[+] %s\n", out)
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

	style.Create.PrintF("[+] %s\n", out)
	return nil
}
