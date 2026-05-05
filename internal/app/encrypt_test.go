package app_test

import (
	"os"
	"testing"

	"github.com/binarysoupdev/cryptool/internal/app"
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
	Keyfile        string
}

func (s *EncryptSuite) SetupTest() {
	r := rand.New(42)

	s.Password = r.ASCII(30)
	var f *os.File

	f, s.PlaintextFile = file.Create(s.T(), r.ASCII(10))
	f.Write(r.Bytes(50))
	f.Close()

	s.CiphertextFile = file.NewPath(s.T(), r.ASCII(10))

	f, s.Keyfile = file.Create(s.T(), r.ASCII(10))
	f.Write(r.Bytes(32))
	f.Close()
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

func (s *EncryptSuite) TestRunPasswordWrongVerify() {
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

func (s *EncryptSuite) TestRunPasswordCorrectVerify() {
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

func (s *EncryptSuite) TestRunInvalidKeyfile() {
	//-- arrange
	r := rand.New(42)
	os.WriteFile(s.Keyfile, r.Bytes(50), 0666)

	//-- act
	s.RunCommand("-i", s.PlaintextFile, "-o", s.CiphertextFile, "-key", s.Keyfile)

	//-- assert
	s.RequireResultFail("invalid key length")
}

func (s *EncryptSuite) TestRunValidKeyfile() {
	//-- arrange
	out := pipe.OpenStdout(1)
	defer out.Close()

	//-- act
	s.RunCommand("-i", s.PlaintextFile, "-o", s.CiphertextFile, "-key", s.Keyfile)

	//-- assert
	s.RequireResultPass()

	s.Assert().FileExists(s.PlaintextFile)
	s.Assert().FileExists(s.CiphertextFile)

	s.Assert().Contains(out.ReadLine(), "[+] "+s.CiphertextFile)
}

func (s *EncryptSuite) TestRunWithRemove() {
	//-- arrange
	out := pipe.OpenStdout(2)
	defer out.Close()

	//-- act
	s.RunCommand("-i", s.PlaintextFile, "-o", s.CiphertextFile, "-key", s.Keyfile, "-rm")

	//-- assert
	s.RequireResultPass()

	s.Assert().NoFileExists(s.PlaintextFile)
	s.Assert().FileExists(s.CiphertextFile)

	s.Assert().Contains(out.ReadLine(), "[+] "+s.CiphertextFile)
	s.Assert().Contains(out.ReadLine(), "[-] "+s.PlaintextFile)
}
