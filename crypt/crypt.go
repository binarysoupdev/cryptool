package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"
	"errors"
)

const (
	PASSWORD_KEY_SIZE = 32 // AES-256
	PBKDF2_ITERATIONS = 100_000
	SALT_SIZE         = 16
)

type Crypt struct {
	cipher cipher.AEAD
}

// Create a new AES-Crypt object from a key. Key must be 16, 24, or 32 bytes.
func New(key []byte) (Crypt, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return Crypt{}, errors.New("invalid key length")
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
	}, nil
}

// Create a new Crypt object from a password and a random salt. Returns the object and the salt.
func NewFromPassword(password string) (Crypt, []byte) {
	salt := make([]byte, SALT_SIZE)
	rand.Read(salt)

	return LoadFromPassword(password, salt), salt
}

// Load a Crypt object from a password and an existing salt. Returns the object.
func LoadFromPassword(password string, salt []byte) Crypt {
	key, err := pbkdf2.Key(sha256.New, password, salt, PBKDF2_ITERATIONS, PASSWORD_KEY_SIZE)
	if err != nil {
		panic(err)
	}

	c, _ := New(key)
	return c
}
