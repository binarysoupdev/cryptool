package crypt

import (
	"fmt"
)

// Decrypt the ciphertext and return the resulting plaintext. Also returns any decryption errors.
func (c Crypt) Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < NONCE_SIZE+1 {
		return nil, fmt.Errorf("invalid ciphertext length: %d", len(ciphertext))
	}

	return c.cipher.Open(nil, ciphertext[:NONCE_SIZE], ciphertext[NONCE_SIZE:], nil)
}
