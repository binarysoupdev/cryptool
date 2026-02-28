package crypt

func (c Crypt) Encrypt(plaintext []byte) []byte {
	nonce := make([]byte, c.cipher.NonceSize())

	return c.cipher.Seal(nil, nonce, plaintext, nil)
}

func (c Crypt) Decrypt(ciphertext []byte) ([]byte, error) {
	nonce := make([]byte, c.cipher.NonceSize())

	return c.cipher.Open(nil, nonce, ciphertext, nil)
}
