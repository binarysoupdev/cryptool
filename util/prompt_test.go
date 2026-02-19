package util_test

import (
	"fmt"
	"testing"

	"github.com/binarysoupdev/cryptool/util"
	"github.com/binarysoupdev/tinsel/pipe"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/assert"
)

func TestReadPasswordFromTerminal(t *testing.T) {
	//-- arrange
	const SEED = 42
	r := rand.New(SEED)

	PROMPT := r.ASCII(10)
	PASSWORD := r.ASCII(30)

	io := pipe.OpenStdio(1, 1, false)
	defer io.Close()

	//-- act
	io.Queue("PASSWORD: ", PASSWORD)
	io.EndQueue()

	res := util.PromptPassword(PROMPT)

	//-- assert
	assert.Equal(t, PASSWORD, res)
	assert.Equal(t, fmt.Sprintf("%s PASSWORD: ", PROMPT), io.ReadLine())
}
