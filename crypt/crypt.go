package crypt

import (
	"crypto/aes"
	"crypto/cipher"
)

type Crypt struct {
	cipher cipher.AEAD
}

func New() (Crypt, error) {
	key := make([]byte, 32) // AES-256

	block, err := aes.NewCipher(key)
	if err != nil {
		return Crypt{}, chainError(err, "error creating AES cipher")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return Crypt{}, chainError(err, "error creating GCM mode")
	}

	return Crypt{
		cipher: gcm,
	}, nil
}
