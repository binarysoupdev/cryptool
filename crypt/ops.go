package crypt

import (
	"crypto/rand"

	"github.com/binarysoupdev/cryptool/util"
)

func (c Crypt) Encrypt(plaintext []byte) Ciphertext {
	nonce := make([]byte, NONCE_SIZE)
	rand.Read(nonce)

	ciphertext := c.cipher.Seal(nil, nonce, plaintext, nil)
	return NewCiphertext(c.salt, nonce, ciphertext)
}

func (c Crypt) Decrypt(ct Ciphertext) ([]byte, error) {
	plaintext, err := c.cipher.Open(nil, ct.Nonce(), ct.Text(), nil)
	if err != nil {
		return nil, util.ChainError(err, "error decryting ciphertext")
	}

	return plaintext, nil
}
