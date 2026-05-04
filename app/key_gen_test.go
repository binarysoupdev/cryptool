package app_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/binarysoupdev/cryptool/app"
	"github.com/binarysoupdev/go-commando/test"
	"github.com/binarysoupdev/tinsel/file"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/suite"
)

type KeyGenSuite struct {
	test.CommandSuite[*app.KeyGenCommand]
}

func TestKeyGenSuite(t *testing.T) {
	suite.Run(t, &KeyGenSuite{
		CommandSuite: test.NewCommandSuite(app.NewKeyGenCommand()),
	})
}

func (s *KeyGenSuite) TestRunEmptyOutputFilepath() {
	//-- act
	s.RunCommand()

	//-- assert
	s.RequireResultFail("output filepath cannot be empty")
}

func (s *KeyGenSuite) TestRunInvalidKeyLength() {
	//-- arrange
	r := rand.New(42)
	OUT := file.NewPath(s.T(), r.ASCII(10))

	//-- act
	s.RunCommand("-o", OUT, "-l", "50")

	//-- assert
	s.RequireResultFail("invalid key length")
}

func (s *KeyGenSuite) TestValidKeyLengths() {
	//-- arrange
	r := rand.New(42)
	OUT := file.NewPath(s.T(), r.ASCII(10))
	LENGTHS := []int{16, 24, 32}

	for _, length := range LENGTHS {
		//-- act
		s.RunCommand("-o", OUT, "-l", fmt.Sprintf("%d", length))

		//-- assert
		s.RequireResultPass()
	}
}

func (s *KeyGenSuite) TestRunCreateKeyFile() {
	//-- arrange
	r := rand.New(42)
	OUT := file.NewPath(s.T(), r.ASCII(10))

	//-- act
	s.RunCommand("-o", OUT)

	//-- assert
	s.RequireResultPass()
	s.Require().FileExists(OUT)

	bytes, err := os.ReadFile(OUT)
	s.Require().NoError(err)

	s.Assert().Len(bytes, 32) //default 32 bytes
}
