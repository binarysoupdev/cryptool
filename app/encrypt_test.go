package app_test

import (
	"os"
	"testing"

	"github.com/binarysoupdev/cryptool/app"
	"github.com/binarysoupdev/tinsel/file"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/binarysoupdev/tinsel/tinsel"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type EncryptSuite struct {
	suite.Suite
	Filepath string
}

func TestEncryptSuite(t *testing.T) {
	suite.Run(t, &EncryptSuite{})
}

func (s *EncryptSuite) SetupTest() {
	r := rand.New(SEED)
	var f *os.File

	f, s.Filepath = file.Create(s.T(), r.ASCII(10))
	f.Write(r.Bytes(50))
	f.Close()
}

//==============================

func (s *EncryptSuite) TestRunEncryptWrongVerify() {
	//-- arrange
	r := rand.New(SEED)
	PASSWORD := r.ASCII(30)

	in := tinsel.OpenStdinPipe(2)
	defer in.Close()

	out := tinsel.OpenStdoutPipe()
	defer out.Close()

	//-- act
	in.Submit(PASSWORD, PASSWORD+"x")
	res := app.Run(s.Filepath, false)

	//-- assert
	require.Error(s.T(), res)
	assert.Contains(s.T(), out.ReadLine(), "ENCRYPT")

	assert.Contains(s.T(), res.Error(), "passwords do not match")
}

func (s *EncryptSuite) TestRunEncryptNoRemove() {
	//-- arrange
	r := rand.New(SEED)
	PASSWORD := r.ASCII(30)

	in := tinsel.OpenStdinPipe(2)
	defer in.Close()

	out := tinsel.OpenStdoutPipe()
	defer out.Close()

	//-- act
	in.Submit(PASSWORD, PASSWORD)
	res := app.Run(s.Filepath, false)

	//-- assert
	require.NoError(s.T(), res)
	assert.Contains(s.T(), out.ReadLine(), "ENCRYPT")

	assert.FileExists(s.T(), s.Filepath)
	assert.FileExists(s.T(), s.Filepath+app.CRYPT_EXT)

	assert.Contains(s.T(), out.ReadLine(), "[+] "+s.Filepath+app.CRYPT_EXT)
}

func (s *EncryptSuite) TestRunEncryptWithRemove() {
	//-- arrange
	r := rand.New(SEED)
	PASSWORD := r.ASCII(30)

	in := tinsel.OpenStdinPipe(2)
	defer in.Close()

	out := tinsel.OpenStdoutPipe()
	defer out.Close()

	//-- act
	in.Submit(PASSWORD, PASSWORD)
	res := app.Run(s.Filepath, true)

	//-- assert
	require.NoError(s.T(), res)
	assert.Contains(s.T(), out.ReadLine(), "ENCRYPT")

	assert.NoFileExists(s.T(), s.Filepath)
	assert.FileExists(s.T(), s.Filepath+app.CRYPT_EXT)

	assert.Contains(s.T(), out.ReadLine(), "[+] "+s.Filepath+app.CRYPT_EXT)
	assert.Contains(s.T(), out.ReadLine(), "[-] "+s.Filepath)
}
