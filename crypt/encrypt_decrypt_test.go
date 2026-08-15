package crypt_test

import (
	"testing"

	"github.com/binarysoupdev/cryptool/crypt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncryptDecryptCorrectKey(t *testing.T) {
	//-- arrange
	const PASSWORD = "Password123!"
	PLAINTEXT := []byte("plaintext")

	//-- act
	c, salt := crypt.NewFromPassword(PASSWORD)
	ciphertext := c.Encrypt(PLAINTEXT)

	c = crypt.LoadFromPassword(PASSWORD, salt)
	res, err := c.Decrypt(ciphertext)

	//-- assert
	require.NoError(t, err)
	assert.Equal(t, PLAINTEXT, res)
}

func TestEncryptDecryptWrongKey(t *testing.T) {
	//-- arrange
	const PASSWORD = "Password123!"
	PLAINTEXT := []byte("plaintext")

	//-- act
	c, salt := crypt.NewFromPassword(PASSWORD)
	ciphertext := c.Encrypt(PLAINTEXT)

	c = crypt.LoadFromPassword(PASSWORD+"x", salt)
	_, err := c.Decrypt(ciphertext)

	//-- assert
	require.Error(t, err)
	assert.Contains(t, err.Error(), "message authentication failed")
}
