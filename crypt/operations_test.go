package crypt_test

import (
	"testing"

	"github.com/binarysoupdev/cryptool/crypt"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncryptDecryptCorrectKey(t *testing.T) {
	//-- arrange
	const SEED = 42
	r := rand.New(SEED)

	PASSWORD := r.ASCII(30)
	PLAINTEXT := r.Bytes(100)

	//-- act
	c, salt := crypt.New(PASSWORD)
	ciphertext := c.Encrypt(PLAINTEXT)

	c = crypt.Load(PASSWORD, salt)
	res, err := c.Decrypt(ciphertext)

	//-- assert
	require.NoError(t, err)
	assert.Equal(t, PLAINTEXT, res)
}

func TestEncryptDecryptWrongKey(t *testing.T) {
	//-- arrange
	const SEED = 42
	r := rand.New(SEED)

	PASSWORD := r.ASCII(30)
	PLAINTEXT := r.Bytes(100)

	//-- act
	c, salt := crypt.New(PASSWORD)
	ciphertext := c.Encrypt(PLAINTEXT)

	c = crypt.Load(PASSWORD+"x", salt)
	_, err := c.Decrypt(ciphertext)

	//-- assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "message authentication failed")
}
