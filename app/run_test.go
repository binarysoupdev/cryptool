package app_test

import (
	"testing"

	"github.com/binarysoupdev/cryptool/app"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const SEED = 64

func TestRunEmptyFilename(t *testing.T) {
	//-- act
	res := app.Run("", false)

	//-- assert
	require.Error(t, res)
	assert.Contains(t, res.Error(), "filepath cannot be empty")
}

func TestRunInvalidFilename(t *testing.T) {
	//-- arrange
	r := rand.New(SEED)
	FILE := r.ASCII(10)

	//-- act
	res := app.Run(FILE, false)

	//-- arrange
	require.Error(t, res)
	assert.Contains(t, res.Error(), "error reading input file")
}
