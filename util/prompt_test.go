package util_test

import (
	"testing"

	"github.com/binarysoupdev/cryptool/util"
	"github.com/binarysoupdev/tonsole/rand"
	"github.com/binarysoupdev/tonsole/testio"
	"github.com/stretchr/testify/assert"
)

func TestReadPasswordFromTerminal(t *testing.T) {
	//-- arrange
	r := rand.New(64)
	PROMPT := r.ASCII(10)
	PASSWORD := r.ASCII(30)

	out := testio.OpenStdoutPipe()
	defer out.Restore()

	in := testio.OpenStdinPipe()
	defer in.Restore()

	//-- act
	in.Submit(PASSWORD)
	res := util.PromptPassword(PROMPT)
	out.CloseInput()

	//-- assert
	assert.Equal(t, PASSWORD, res)
	assert.Contains(t, out.NextLine(t), PROMPT)
}
