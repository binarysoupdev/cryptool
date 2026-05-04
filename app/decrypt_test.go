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
	Password               string
	PasswordCiphertextFile string
	PlaintextFile          string
	Keyfile                string
	KeyCiphertextFile      string
}

func (s *DecryptSuite) SetupTest() {
	r := rand.New(42)

	s.Password = r.ASCII(30)
	var f *os.File

	c, salt := crypt.NewFromPassword(s.Password)
	f, s.PasswordCiphertextFile = file.Create(s.T(), r.ASCII(10))
	f.Write(salt)
	f.Write(c.Encrypt(r.Bytes(50)))
	f.Close()

	s.PlaintextFile = file.NewPath(s.T(), r.ASCII(10))

	key := r.Bytes(32)
	f, s.Keyfile = file.Create(s.T(), r.ASCII(10))
	f.Write(key)
	f.Close()

	c, _ = crypt.New(key)
	f, s.KeyCiphertextFile = file.Create(s.T(), r.ASCII(10))
	f.Write(c.Encrypt(r.Bytes(50)))
	f.Close()
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
	s.RunCommand("-i", s.PasswordCiphertextFile)

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
	s.RunCommand("-i", s.PasswordCiphertextFile, "-o", s.PlaintextFile)

	//-- assert
	s.RequireResultFail("error decrypting ciphertext")
}

func (s *DecryptSuite) TestRunCorrectPassword() {
	//-- arrange
	io := pipe.OpenStdio(1, 2, false)
	defer io.Close()

	//-- act
	io.Queue("PASSWORD: ", s.Password)
	io.EndQueue()

	s.RunCommand("-i", s.PasswordCiphertextFile, "-o", s.PlaintextFile)

	//-- assert
	s.RequireResultPass()

	s.Assert().FileExists(s.PasswordCiphertextFile)
	s.Assert().FileExists(s.PlaintextFile)

	s.Assert().Equal("Enter PASSWORD: ", io.ReadLine())
	s.Assert().Contains(io.ReadLine(), "[+] "+s.PlaintextFile)
}

func (s *DecryptSuite) TestRunInvalidKeyfile() {
	//-- arrange
	r := rand.New(42)
	os.WriteFile(s.Keyfile, r.Bytes(50), 0666)

	//-- act
	s.RunCommand("-i", s.KeyCiphertextFile, "-o", s.PlaintextFile, "-key", s.Keyfile)

	//-- assert
	s.RequireResultFail("invalid key length")
}

func (s *DecryptSuite) TestRunValidKeyfile() {
	//-- arrange
	out := pipe.OpenStdout(1)
	defer out.Close()

	//-- act
	s.RunCommand("-i", s.KeyCiphertextFile, "-o", s.PlaintextFile, "-key", s.Keyfile)

	//-- assert
	s.RequireResultPass()

	s.Assert().FileExists(s.KeyCiphertextFile)
	s.Assert().FileExists(s.PlaintextFile)

	s.Assert().Contains(out.ReadLine(), "[+] "+s.PlaintextFile)
}

func (s *DecryptSuite) TestRunWithRemove() {
	//-- arrange
	out := pipe.OpenStdout(2)
	defer out.Close()

	//-- act
	s.RunCommand("-i", s.KeyCiphertextFile, "-o", s.PlaintextFile, "-key", s.Keyfile, "-rm")

	//-- assert
	s.RequireResultPass()

	s.Assert().NoFileExists(s.KeyCiphertextFile)
	s.Assert().FileExists(s.PlaintextFile)

	s.Assert().Contains(out.ReadLine(), "[+] "+s.PlaintextFile)
	s.Assert().Contains(out.ReadLine(), "[-] "+s.KeyCiphertextFile)
}
