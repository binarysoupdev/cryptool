package crypt

import (
	"crypto/aes"
	"crypto/cipher"
)

type Crypt struct {
	cipher cipher.AEAD
}

func New() Crypt {
	key := make([]byte, 32) //256-bit

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
