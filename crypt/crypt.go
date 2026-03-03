package crypt

import (
	"crypto/aes"
	"crypto/cipher"
)

type Crypt struct {
	cipher cipher.AEAD
}

func New(key []byte) Crypt {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	cipher, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	return Crypt{
		cipher: cipher,
	}
}
