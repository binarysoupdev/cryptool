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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DecryptSuite struct {
	test.CommandSuite[*app.AppCommand]
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

	f, s.CiphertextFile = file.Create(s.T(), r.ASCII(10)+app.CRYPT_EXT)
	f.Write(salt)
	f.Write(ciphertext)
	f.Close()

	s.PlaintextFile = s.CiphertextFile[:len(s.CiphertextFile)-len(app.CRYPT_EXT)]
}

//==============================

func TestDecryptSuite(t *testing.T) {
	suite.Run(t, &DecryptSuite{
		CommandSuite: test.NewCommandSuite(app.NewAppCommand()),
	})
}

func (s *DecryptSuite) TestRunDecryptWrongPassword() {
	//-- arrange
	io := pipe.OpenStdin(1)
	defer io.Close()

	//-- act
	io.Queue("PASSWORD: ", s.Password+"x")
	s.RunCommand("-i", s.CiphertextFile)

	//-- assert
	s.RequireResultFail("error decrypting ciphertext")
}

func (s *DecryptSuite) TestRunDecryptNoRemove() {
	//-- arrange
	io := pipe.OpenStdio(1, 2, false)
	defer io.Close()

	//-- act
	io.Queue("PASSWORD: ", s.Password)
	io.EndQueue()

	s.RunCommand("-i", s.CiphertextFile)

	//-- assert
	s.RequireResultPass()

	assert.FileExists(s.T(), s.CiphertextFile)
	assert.FileExists(s.T(), s.PlaintextFile)

	assert.Equal(s.T(), "Enter PASSWORD: ", io.ReadLine())
	assert.Contains(s.T(), io.ReadLine(), "[+] "+s.PlaintextFile)
}

func (s *DecryptSuite) TestRunDecryptWithRemove() {
	//-- arrange
	io := pipe.OpenStdio(1, 3, false)
	defer io.Close()

	//-- act
	io.Queue("PASSWORD: ", s.Password)
	io.EndQueue()

	s.RunCommand("-i", s.CiphertextFile, "-rm")

	//-- assert
	s.RequireResultPass()

	assert.NoFileExists(s.T(), s.CiphertextFile)
	assert.FileExists(s.T(), s.PlaintextFile)

	assert.Equal(s.T(), "Enter PASSWORD: ", io.ReadLine())
	assert.Contains(s.T(), io.ReadLine(), "[+] "+s.PlaintextFile)
	assert.Contains(s.T(), io.ReadLine(), "[-] "+s.CiphertextFile)
}
