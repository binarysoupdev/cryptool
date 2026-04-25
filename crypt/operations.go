package crypt

import "crypto/rand"

func (c Crypt) Encrypt(plaintext []byte) []byte {
	nonce := make([]byte, c.cipher.NonceSize())
	rand.Read(nonce)

	return c.cipher.Seal(nonce, nonce, plaintext, nil)
}

func (c Crypt) Decrypt(ciphertext []byte) ([]byte, error) {
	return c.cipher.Open(nil, ciphertext[:c.cipher.NonceSize()], ciphertext[c.cipher.NonceSize():], nil)
}
