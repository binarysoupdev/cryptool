package util_test

import (
	"testing"

	"github.com/binarysoupdev/cryptool/util"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/binarysoupdev/tinsel/tinsel"
	"github.com/stretchr/testify/assert"
)

const SEED = 64

func TestReadPasswordFromTerminal(t *testing.T) {
	//-- arrange
	r := rand.New(SEED)
	PROMPT := r.ASCII(10)
	PASSWORD := r.ASCII(30)

	in := tinsel.OpenStdinPipe(1)
	defer in.Close()

	out := tinsel.OpenStdoutPipe()
	defer out.Close()

	//-- act
	in.Submit(PASSWORD)
	res := util.PromptPassword(PROMPT)
	out.EndLine()

	//-- assert
	assert.Equal(t, PASSWORD, res)
	assert.Contains(t, out.ReadLine(), PROMPT)
}
