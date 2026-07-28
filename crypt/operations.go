package crypt

import (
	"crypto/rand"
	"fmt"
)

const NONCE_SIZE = 12 // AES-GCM

// Encrypt the plaintext and return the resulting ciphertext.
func (c Crypt) Encrypt(plaintext []byte) []byte {
	nonce := make([]byte, NONCE_SIZE)
	rand.Read(nonce)

	return c.cipher.Seal(nonce, nonce, plaintext, nil)
}

// Decrypt the ciphertext and return the resulting plaintext. Also returns any decryption errors.
func (c Crypt) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < NONCE_SIZE+1 {
		return nil, fmt.Errorf("invalid ciphertext length: %d", len(ciphertext))
	}

	return c.cipher.Open(nil, ciphertext[:NONCE_SIZE], ciphertext[NONCE_SIZE:], nil)
}
