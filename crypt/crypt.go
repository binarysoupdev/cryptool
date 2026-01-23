package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/pbkdf2"
	"crypto/sha256"
)

const (
	KEY_SIZE          = 32 // AES-256
	PBKDF2_ITERATIONS = 100_000
)

type Crypt struct {
	cipher cipher.AEAD
}

func New(password string) (Crypt, error) {
	key, err := pbkdf2.Key(sha256.New, password, nil, PBKDF2_ITERATIONS, KEY_SIZE)
	if err != nil {
		return Crypt{}, chainError(err, "error generating key from password")
	}

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
