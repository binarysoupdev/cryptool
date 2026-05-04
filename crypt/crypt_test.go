package crypt_test

import (
	"testing"

	"github.com/binarysoupdev/cryptool/crypt"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewInvalidKeySize(t *testing.T) {
	//-- arrange
	r := rand.New(42)
	KEY := r.Bytes(50)

	//-- act
	_, res := crypt.New(KEY)

	//-- assert
	require.Error(t, res)
	assert.ErrorContains(t, res, "invalid key length")
}

func TestNewValidKeySize(t *testing.T) {
	//-- arrange
	r := rand.New(42)
	KEYS := [][]byte{r.Bytes(16), r.Bytes(24), r.Bytes(32)}

	for _, key := range KEYS {
		//-- act
		_, res := crypt.New(key)

		//-- assert
		require.NoError(t, res)
	}
}
