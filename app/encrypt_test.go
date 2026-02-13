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
	PlaintextFile  string
	CiphertextFile string
}

func (s *EncryptSuite) SetupTest() {
	r := rand.New(SEED)
	var f *os.File

	f, s.PlaintextFile = file.Create(s.T(), r.ASCII(10))
	f.Write(r.Bytes(50))
	f.Close()

	s.CiphertextFile = s.PlaintextFile + app.CRYPT_EXT
}

//==============================

func TestEncryptSuite(t *testing.T) {
	suite.Run(t, &EncryptSuite{})
}

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
	res := app.Run(s.PlaintextFile, false)

	//-- assert
	require.Error(s.T(), res)
	assert.Contains(s.T(), res.Error(), "passwords do not match")

	assert.Contains(s.T(), out.ReadLine(), "ENCRYPT")
	assert.Contains(s.T(), out.ReadLine(), "New")
	assert.Contains(s.T(), out.ReadLine(), "Verify")
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
	res := app.Run(s.PlaintextFile, false)

	//-- assert
	require.NoError(s.T(), res)

	assert.FileExists(s.T(), s.PlaintextFile)
	assert.FileExists(s.T(), s.CiphertextFile)

	assert.Contains(s.T(), out.ReadLine(), "ENCRYPT")
	assert.Contains(s.T(), out.ReadLine(), "New")
	assert.Contains(s.T(), out.ReadLine(), "Verify")
	assert.Contains(s.T(), out.ReadLine(), "[+] "+s.CiphertextFile)
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
	res := app.Run(s.PlaintextFile, true)

	//-- assert
	require.NoError(s.T(), res)

	assert.NoFileExists(s.T(), s.PlaintextFile)
	assert.FileExists(s.T(), s.CiphertextFile)

	assert.Contains(s.T(), out.ReadLine(), "ENCRYPT")
	assert.Contains(s.T(), out.ReadLine(), "New")
	assert.Contains(s.T(), out.ReadLine(), "Verify")
	assert.Contains(s.T(), out.ReadLine(), "[+] "+s.CiphertextFile)
	assert.Contains(s.T(), out.ReadLine(), "[-] "+s.PlaintextFile)
}
