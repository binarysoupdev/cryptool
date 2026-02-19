package app_test

import (
	"os"
	"testing"

	"github.com/binarysoupdev/cryptool/app"
	"github.com/binarysoupdev/tinsel/file"
	"github.com/binarysoupdev/tinsel/pipe"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type EncryptSuite struct {
	suite.Suite
	Password       string
	PlaintextFile  string
	CiphertextFile string
}

func (s *EncryptSuite) SetupTest() {
	const SEED = 42
	r := rand.New(SEED)

	s.Password = r.ASCII(30)
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
	in := pipe.OpenStdin(2)
	defer in.Close()

	//-- act
	in.Queue("PASSWORD: ", s.Password)
	in.Queue("PASSWORD: ", s.Password+"x")

	res := app.Run(s.PlaintextFile, false)

	//-- assert
	require.Error(s.T(), res)
	assert.Contains(s.T(), res.Error(), "passwords do not match")
}

func (s *EncryptSuite) TestRunEncryptNoRemove() {
	//-- arrange
	io := pipe.OpenStdio(2, 3, false)
	defer io.Close()

	//-- act
	io.Queue("PASSWORD: ", s.Password)
	io.Queue("PASSWORD: ", s.Password)
	io.EndQueue()

	res := app.Run(s.PlaintextFile, false)

	//-- assert
	require.NoError(s.T(), res)

	assert.FileExists(s.T(), s.PlaintextFile)
	assert.FileExists(s.T(), s.CiphertextFile)

	io.SkipLines(2)
	assert.Contains(s.T(), io.ReadLine(), "[+] "+s.CiphertextFile)
}

func (s *EncryptSuite) TestRunEncryptWithRemove() {
	//-- arrange
	io := pipe.OpenStdio(2, 4, false)
	defer io.Close()

	//-- act
	io.Queue("PASSWORD: ", s.Password)
	io.Queue("PASSWORD: ", s.Password)
	io.EndQueue()

	res := app.Run(s.PlaintextFile, true)

	//-- assert
	require.NoError(s.T(), res)

	assert.NoFileExists(s.T(), s.PlaintextFile)
	assert.FileExists(s.T(), s.CiphertextFile)

	io.SkipLines(2)
	assert.Contains(s.T(), io.ReadLine(), "[+] "+s.CiphertextFile)
	assert.Contains(s.T(), io.ReadLine(), "[-] "+s.PlaintextFile)
}
