package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"
)

const (
	KEY_SIZE          = 32 // AES-256
	PBKDF2_ITERATIONS = 100_000
	SALT_SIZE         = 16
)

type Crypt struct {
	cipher cipher.AEAD
	salt   []byte
}

func New(password string) Crypt {
	salt := make([]byte, SALT_SIZE)
	rand.Read(salt)

	return Load(password, salt)
}

func Load(password string, salt []byte) Crypt {
	key, err := pbkdf2.Key(sha256.New, password, salt, PBKDF2_ITERATIONS, KEY_SIZE)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}

	return Crypt{
		cipher: gcm,
		salt:   salt,
	}
}
