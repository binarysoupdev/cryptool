package crypt

func (c Crypt) Encrypt(plaintext []byte) ([]byte, error) {
	nonce := make([]byte, c.cipher.NonceSize())

	ciphertext := c.cipher.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nil
}

func (c Crypt) Decrypt(ciphertext []byte) ([]byte, error) {
	nonce := make([]byte, c.cipher.NonceSize())

	plaintext, err := c.cipher.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, chainError(err, "error decryting ciphertext")
	}

	return plaintext, nil
}
