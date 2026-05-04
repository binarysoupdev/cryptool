package app_test

import (
	"testing"

	"github.com/binarysoupdev/cryptool/app"
	"github.com/binarysoupdev/go-commando/test"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/suite"
)

type AppTestSuite struct {
	test.CommandSuite[*app.AppCommand]
}

func TestAppCommandSuite(t *testing.T) {
	suite.Run(t, &AppTestSuite{
		CommandSuite: test.NewCommandSuite(app.NewAppCommand()),
	})
}

func (s *AppTestSuite) TestRunEmptyFilename() {
	//-- act
	s.RunCommand()

	//-- assert
	s.RequireResultFail("filepath cannot be empty")
}

func (s *AppTestSuite) TestRunInvalidFilename() {
	//-- arrange
	r := rand.New(42)
	FILE := r.ASCII(10)

	//-- act
	s.RunCommand("-i", FILE)

	//-- assert
	s.RequireResultFail("error reading input file")
}
