package app_test

import (
	"os"
	"testing"

	"github.com/binarysoupdev/cryptool/app"
	"github.com/binarysoupdev/go-commando/test"
	"github.com/binarysoupdev/tinsel/file"
	"github.com/binarysoupdev/tinsel/pipe"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/suite"
)

type EncryptSuite struct {
	test.CommandSuite[*app.EncryptCommand]
	Password       string
	PlaintextFile  string
	CiphertextFile string
}

func (s *EncryptSuite) SetupTest() {
	r := rand.New(42)

	s.Password = r.ASCII(30)
	var f *os.File

	f, s.PlaintextFile = file.Create(s.T(), r.ASCII(10))
	f.Write(r.Bytes(50))
	f.Close()

	s.CiphertextFile = file.NewPath(s.T(), r.ASCII(10))
}

//==============================

func TestEncryptSuite(t *testing.T) {
	suite.Run(t, &EncryptSuite{
		CommandSuite: test.NewCommandSuite(app.NewEncryptCommand()),
	})
}

func (s *EncryptSuite) TestRunEmptyInputFilepath() {
	//-- act
	s.RunCommand()

	//-- assert
	s.RequireResultFail("input filepath cannot be empty")
}

func (s *EncryptSuite) TestRunEmptyOutputFilepath() {
	//-- act
	s.RunCommand("-i", s.PlaintextFile)

	//-- assert
	s.RequireResultFail("output filepath cannot be empty")
}

func (s *EncryptSuite) TestRunInvalidInputFile() {
	//-- arrange
	r := rand.New(42)
	FILE := r.ASCII(10)

	//-- act
	s.RunCommand("-i", FILE, "-o", s.CiphertextFile)

	//-- assert
	s.RequireResultFail("error reading plaintext file")
}

func (s *EncryptSuite) TestRunWrongVerify() {
	//-- arrange
	in := pipe.OpenStdin(2)
	defer in.Close()

	//-- act
	in.Queue("PASSWORD: ", s.Password)
	in.Queue("PASSWORD: ", s.Password+"x")

	s.RunCommand("-i", s.PlaintextFile, "-o", s.CiphertextFile)

	//-- assert
	s.RequireResultFail("passwords do not match")
}

func (s *EncryptSuite) TestRunNoRemove() {
	//-- arrange
	io := pipe.OpenStdio(2, 3, false)
	defer io.Close()

	//-- act
	io.Queue("PASSWORD: ", s.Password)
	io.Queue("PASSWORD: ", s.Password)
	io.EndQueue()

	s.RunCommand("-i", s.PlaintextFile, "-o", s.CiphertextFile)

	//-- assert
	s.RequireResultPass()

	s.Assert().FileExists(s.PlaintextFile)
	s.Assert().FileExists(s.CiphertextFile)

	io.SkipLines(2)
	s.Assert().Contains(io.ReadLine(), "[+] "+s.CiphertextFile)
}

func (s *EncryptSuite) TestRunWithRemove() {
	//-- arrange
	io := pipe.OpenStdio(2, 4, false)
	defer io.Close()

	//-- act
	io.Queue("PASSWORD: ", s.Password)
	io.Queue("PASSWORD: ", s.Password)
	io.EndQueue()

	s.RunCommand("-i", s.PlaintextFile, "-o", s.CiphertextFile, "-rm")

	//-- assert
	s.RequireResultPass()

	s.Assert().NoFileExists(s.PlaintextFile)
	s.Assert().FileExists(s.CiphertextFile)

	io.SkipLines(2)
	s.Assert().Contains(io.ReadLine(), "[+] "+s.CiphertextFile)
	s.Assert().Contains(io.ReadLine(), "[-] "+s.PlaintextFile)
}
