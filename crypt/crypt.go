package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"
)

const (
	KEY_SIZE   = 32
	ITERATIONS = 100_000
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

func NewFromPassword(password string) Crypt {
	salt := make([]byte, 16)
	rand.Read(salt)

	key, err := pbkdf2.Key(sha256.New, password, salt, ITERATIONS, KEY_SIZE)
	if err != nil {
		panic(err)
	}

	return New(key)
}
