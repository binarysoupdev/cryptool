package crypt_test

import (
	"testing"

	"github.com/binarysoupdev/cryptool/crypt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewInvalidKeySize(t *testing.T) {
	//-- arrange
	KEY := make([]byte, 50)

	//-- act
	_, res := crypt.New(KEY)

	//-- assert
	require.Error(t, res)
	assert.ErrorContains(t, res, "invalid key length")
}

func TestNewValidKeySize(t *testing.T) {
	//-- arrange
	KEYS := [][]byte{make([]byte, 16), make([]byte, 24), make([]byte, 32)}

	for _, key := range KEYS {
		//-- act
		_, res := crypt.New(key)

		//-- assert
		require.NoError(t, res)
	}
}
