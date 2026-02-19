package app_test

import (
	"testing"

	"github.com/binarysoupdev/cryptool/app"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRunEmptyFilename(t *testing.T) {
	//-- act
	res := app.Run("", false)

	//-- assert
	require.Error(t, res)
	assert.Contains(t, res.Error(), "filepath cannot be empty")
}

func TestRunInvalidFilename(t *testing.T) {
	//-- arrange
	const SEED = 42
	r := rand.New(SEED)

	FILE := r.ASCII(10)

	//-- act
	res := app.Run(FILE, false)

	//-- arrange
	require.Error(t, res)
	assert.Contains(t, res.Error(), "error reading input file")
}
