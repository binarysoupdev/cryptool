package crypt

import "crypto/rand"

func (c Crypt) Encrypt(plaintext []byte) ([]byte, error) {
	ciphertext := c.cipher.Seal(nil, c.nonce(), plaintext, nil)
	return ciphertext, nil
}

func (c Crypt) Decrypt(ciphertext []byte) ([]byte, error) {
	plaintext, err := c.cipher.Open(nil, c.nonce(), ciphertext, nil)
	if err != nil {
		return nil, chainError(err, "error decryting ciphertext")
	}

	return plaintext, nil
}

func (c Crypt) nonce() []byte {
	nonce := make([]byte, c.cipher.NonceSize())
	rand.Read(nonce)

	return nonce
}
