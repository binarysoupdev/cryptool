package crypt

import "crypto/rand"

func (c Crypt) Encrypt(plaintext []byte) (Ciphertext, error) {
	nonce := make([]byte, c.cipher.NonceSize())
	rand.Read(nonce)

	ciphertext := c.cipher.Seal(nil, nonce, plaintext, nil)
	return NewCiphertext(nonce, ciphertext), nil
}

func (c Crypt) Decrypt(ct Ciphertext) ([]byte, error) {
	plaintext, err := c.cipher.Open(nil, ct[:NONCE_SIZE], ct[NONCE_SIZE:], nil)
	if err != nil {
		return nil, chainError(err, "error decryting ciphertext")
	}

	return plaintext, nil
}
