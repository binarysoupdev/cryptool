package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func Encrypt(plaintext []byte) ([]byte, error) {
	key := make([]byte, 32) // AES-256

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, chainError(err, "error creating AES cipher")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, chainError(err, "error creating GCM mode")
	}

	nonce := make([]byte, gcm.NonceSize())

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nil
}

func Decrypt(ciphertext []byte) ([]byte, error) {
	key := make([]byte, 32) // AES-256

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, chainError(err, "error creating AES cipher")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, chainError(err, "error creating GCM mode")
	}

	nonce := make([]byte, gcm.NonceSize())

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, chainError(err, "error decryting ciphertext")
	}

	return plaintext, nil
}

func chainError(err error, msg string) error {
	return fmt.Errorf("%s\n  %s", msg, err.Error())
}
