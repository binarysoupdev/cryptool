package app_test

import (
	"os"
	"testing"

	"github.com/binarysoupdev/cryptool/app"
	"github.com/binarysoupdev/cryptool/crypt"
	"github.com/binarysoupdev/tinsel/file"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/binarysoupdev/tinsel/tinsel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DecryptSuite struct {
	suite.Suite
	Password       string
	CiphertextFile string
	PlaintextFile  string
}

func (s *DecryptSuite) SetupTest() {
	r := rand.New(SEED)
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
	suite.Run(t, &DecryptSuite{})
}

func (s *DecryptSuite) TestRunDecryptWrongPassword() {
	//-- arrange
	in := tinsel.OpenStdinPipe(2)
	defer in.Close()

	out := tinsel.OpenStdoutPipe()
	defer out.Close()

	//-- act
	in.Submit(s.Password + "x")
	res := app.Run(s.CiphertextFile, false)

	//-- assert
	require.Error(s.T(), res)
	assert.Contains(s.T(), res.Error(), "error decrypting ciphertext")

	assert.Contains(s.T(), out.ReadLine(), "DECRYPT")
	assert.Contains(s.T(), out.ReadLine(), "Enter")
}

func (s *DecryptSuite) TestRunDecryptNoRemove() {
	//-- arrange
	in := tinsel.OpenStdinPipe(2)
	defer in.Close()

	out := tinsel.OpenStdoutPipe()
	defer out.Close()

	//-- act
	in.Submit(s.Password)
	res := app.Run(s.CiphertextFile, false)

	//-- assert
	require.NoError(s.T(), res)

	assert.FileExists(s.T(), s.CiphertextFile)
	assert.FileExists(s.T(), s.PlaintextFile)

	assert.Contains(s.T(), out.ReadLine(), "DECRYPT")
	assert.Contains(s.T(), out.ReadLine(), "Enter")
	assert.Contains(s.T(), out.ReadLine(), "[+] "+s.PlaintextFile)
}

func (s *DecryptSuite) TestRunDecryptWithRemove() {
	//-- arrange
	in := tinsel.OpenStdinPipe(2)
	defer in.Close()

	out := tinsel.OpenStdoutPipe()
	defer out.Close()

	//-- act
	in.Submit(s.Password)
	res := app.Run(s.CiphertextFile, true)

	//-- assert
	require.NoError(s.T(), res)

	assert.NoFileExists(s.T(), s.CiphertextFile)
	assert.FileExists(s.T(), s.PlaintextFile)

	assert.Contains(s.T(), out.ReadLine(), "DECRYPT")
	assert.Contains(s.T(), out.ReadLine(), "Enter")
	assert.Contains(s.T(), out.ReadLine(), "[+] "+s.PlaintextFile)
	assert.Contains(s.T(), out.ReadLine(), "[-] "+s.CiphertextFile)
}
