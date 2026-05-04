package crypt

import (
	"crypto/rand"
)

// Encrypt the plaintext and return the resulting ciphertext.
func (c Crypt) Encrypt(plaintext []byte) Ciphertext {
	nonce := make([]byte, NONCE_SIZE)
	rand.Read(nonce)

	return c.cipher.Seal(nonce, nonce, plaintext, nil)
}

// Decrypt the ciphertext and return the resulting plaintext. Also returns any decryption errors.
func (c Crypt) Decrypt(ct Ciphertext) ([]byte, error) {
	return c.cipher.Open(nil, ct.Nonce(), ct.Text(), nil)
}
