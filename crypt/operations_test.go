package crypt_test

import (
	"testing"

	"github.com/binarysoupdev/cryptool/crypt"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncryptDecryptCorrectKey(t *testing.T) {
	r := rand.New(1)
	password := r.ASCII(30)
	plaintext := r.Bytes(100)

	c, _ := crypt.New(password)
	res, err := c.Decrypt(c.Encrypt(plaintext))

	require.NoError(t, err)
	assert.Equal(t, plaintext, res, "plaintext do not match")
}

func TestEncryptDecryptWrongKey(t *testing.T) {
	r := rand.New(2)
	password := r.ASCII(30)
	plaintext := r.Bytes(100)

	c, salt := crypt.New(password)
	ciphertext := c.Encrypt(plaintext)

	c = crypt.Load("", salt)
	_, err := c.Decrypt(ciphertext)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "error decrypting ciphertext")
}
