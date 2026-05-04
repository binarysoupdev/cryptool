package app_test

import (
	"os"
	"testing"

	"github.com/binarysoupdev/cryptool/app"
	"github.com/binarysoupdev/cryptool/crypt"
	"github.com/binarysoupdev/go-commando/test"
	"github.com/binarysoupdev/tinsel/file"
	"github.com/binarysoupdev/tinsel/pipe"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/suite"
)

type DecryptSuite struct {
	test.CommandSuite[*app.DecryptCommand]
	Password       string
	CiphertextFile string
	PlaintextFile  string
}

func (s *DecryptSuite) SetupTest() {
	r := rand.New(42)

	s.Password = r.ASCII(30)
	var f *os.File

	c, salt := crypt.New(s.Password)
	ciphertext := c.Encrypt(r.Bytes(50))

	f, s.CiphertextFile = file.Create(s.T(), r.ASCII(10))
	f.Write(salt)
	f.Write(ciphertext)
	f.Close()

	s.PlaintextFile = file.NewPath(s.T(), r.ASCII(10))
}

//==============================

func TestDecryptSuite(t *testing.T) {
	suite.Run(t, &DecryptSuite{
		CommandSuite: test.NewCommandSuite(app.NewDecryptCommand()),
	})
}

func (s *DecryptSuite) TestRunEmptyInputFilepath() {
	//-- act
	s.RunCommand()

	//-- assert
	s.RequireResultFail("input filepath cannot be empty")
}

func (s *DecryptSuite) TestRunEmptyOutputFilepath() {
	//-- act
	s.RunCommand("-i", s.CiphertextFile)

	//-- assert
	s.RequireResultFail("output filepath cannot be empty")
}

func (s *DecryptSuite) TestRunInvalidInputFile() {
	//-- arrange
	r := rand.New(42)
	FILE := r.ASCII(10)

	//-- act
	s.RunCommand("-i", FILE, "-o", s.PlaintextFile)

	//-- assert
	s.RequireResultFail("error reading ciphertext file")
}

func (s *DecryptSuite) TestRunWrongPassword() {
	//-- arrange
	io := pipe.OpenStdin(1)
	defer io.Close()

	//-- act
	io.Queue("PASSWORD: ", s.Password+"x")
	s.RunCommand("-i", s.CiphertextFile, "-o", s.PlaintextFile)

	//-- assert
	s.RequireResultFail("error decrypting ciphertext")
}

func (s *DecryptSuite) TestRunNoRemove() {
	//-- arrange
	io := pipe.OpenStdio(1, 2, false)
	defer io.Close()

	//-- act
	io.Queue("PASSWORD: ", s.Password)
	io.EndQueue()

	s.RunCommand("-i", s.CiphertextFile, "-o", s.PlaintextFile)

	//-- assert
	s.RequireResultPass()

	s.Assert().FileExists(s.CiphertextFile)
	s.Assert().FileExists(s.PlaintextFile)

	s.Assert().Equal("Enter PASSWORD: ", io.ReadLine())
	s.Assert().Contains(io.ReadLine(), "[+] "+s.PlaintextFile)
}

func (s *DecryptSuite) TestRunWithRemove() {
	//-- arrange
	io := pipe.OpenStdio(1, 3, false)
	defer io.Close()

	//-- act
	io.Queue("PASSWORD: ", s.Password)
	io.EndQueue()

	s.RunCommand("-i", s.CiphertextFile, "-o", s.PlaintextFile, "-rm")

	//-- assert
	s.RequireResultPass()

	s.Assert().NoFileExists(s.CiphertextFile)
	s.Assert().FileExists(s.PlaintextFile)

	s.Assert().Equal("Enter PASSWORD: ", io.ReadLine())
	s.Assert().Contains(io.ReadLine(), "[+] "+s.PlaintextFile)
	s.Assert().Contains(io.ReadLine(), "[-] "+s.CiphertextFile)
}
