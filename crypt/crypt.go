package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"

	"github.com/binarysoupdev/cryptool/util"
)

const (
	KEY_SIZE          = 32 // AES-256
	PBKDF2_ITERATIONS = 100_000
	SALT_SIZE         = 16
)

type Crypt struct {
	cipher cipher.AEAD
}

func New(password string) (Crypt, error) {
	salt := make([]byte, SALT_SIZE)
	rand.Read(salt)

	key, err := pbkdf2.Key(sha256.New, password, salt, PBKDF2_ITERATIONS, KEY_SIZE)
	if err != nil {
		return Crypt{}, util.ChainError(err, "error generating key from password")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return Crypt{}, util.ChainError(err, "error creating AES cipher")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return Crypt{}, util.ChainError(err, "error creating GCM mode")
	}

	return Crypt{
		cipher: gcm,
	}, nil
}
