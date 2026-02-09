package util_test

import (
	"testing"

	"github.com/binarysoupdev/cryptool/util"
	"github.com/binarysoupdev/tonsole/pipe"
	"github.com/binarysoupdev/tonsole/rand"
	"github.com/stretchr/testify/assert"
)

func TestReadPasswordFromTerminal(t *testing.T) {
	//--arrange
	r := rand.New(64)
	PROMPT := r.ASCII(10)
	PASSWORD := r.ASCII(30)

	in := pipe.Stdin()
	defer in.RestoreAndClose()
	in.WriteLines(PASSWORD)

	out := pipe.Stdout()
	defer out.Restore()

	//--act
	res := util.PromptPassword(PROMPT)
	out.CloseInput()

	//--assert
	assert.Equal(t, PASSWORD, res)
	assert.Contains(t, out.NextLine(t), PROMPT)
}
