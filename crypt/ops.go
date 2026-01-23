package crypt

import (
	"crypto/rand"

	"github.com/binarysoupdev/cryptool/util"
)

func (c Crypt) Encrypt(plaintext []byte) (Ciphertext, error) {
	nonce := make([]byte, c.cipher.NonceSize())
	rand.Read(nonce)

	ciphertext := c.cipher.Seal(nil, nonce, plaintext, nil)
	return NewCiphertext(nonce, ciphertext), nil
}

func (c Crypt) Decrypt(ct Ciphertext) ([]byte, error) {
	plaintext, err := c.cipher.Open(nil, ct.Nonce(), ct.Text(), nil)
	if err != nil {
		return nil, util.ChainError(err, "error decryting ciphertext")
	}

	return plaintext, nil
}
