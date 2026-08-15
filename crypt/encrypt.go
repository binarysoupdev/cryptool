package crypt

import (
	"crypto/rand"
)

const NONCE_SIZE = 12 // AES-GCM

// Encrypt the plaintext and return the resulting ciphertext.
func (c Crypt) Encrypt(plaintext []byte) []byte {
	nonce := make([]byte, NONCE_SIZE)
	rand.Read(nonce)

	return c.cipher.Seal(nonce, nonce, plaintext, nil)
}
