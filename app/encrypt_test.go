package app_test

import (
	"os"
	"testing"

	"github.com/binarysoupdev/cryptool/app"
	"github.com/binarysoupdev/go-commando/test"
	"github.com/binarysoupdev/tinsel/file"
	"github.com/binarysoupdev/tinsel/pipe"
	"github.com/binarysoupdev/tinsel/rand"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EncryptSuite struct {
	test.CommandSuite[*app.AppCommand]
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

	s.CiphertextFile = s.PlaintextFile + app.CRYPT_EXT
}

//==============================

func TestEncryptSuite(t *testing.T) {
	suite.Run(t, &EncryptSuite{
		CommandSuite: test.NewCommandSuite(app.NewAppCommand()),
	})
}

func (s *EncryptSuite) TestRunEncryptWrongVerify() {
	//-- arrange
	in := pipe.OpenStdin(2)
	defer in.Close()

	//-- act
	in.Queue("PASSWORD: ", s.Password)
	in.Queue("PASSWORD: ", s.Password+"x")

	s.RunCommand("-i", s.PlaintextFile)

	//-- assert
	s.RequireResultFail("passwords do not match")
}

func (s *EncryptSuite) TestRunEncryptNoRemove() {
	//-- arrange
	io := pipe.OpenStdio(2, 3, false)
	defer io.Close()

	//-- act
	io.Queue("PASSWORD: ", s.Password)
	io.Queue("PASSWORD: ", s.Password)
	io.EndQueue()

	s.RunCommand("-i", s.PlaintextFile)

	//-- assert
	s.RequireResultPass()

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

	s.RunCommand("-i", s.PlaintextFile, "-rm")

	//-- assert
	s.RequireResultPass()

	assert.NoFileExists(s.T(), s.PlaintextFile)
	assert.FileExists(s.T(), s.CiphertextFile)

	io.SkipLines(2)
	assert.Contains(s.T(), io.ReadLine(), "[+] "+s.CiphertextFile)
	assert.Contains(s.T(), io.ReadLine(), "[-] "+s.PlaintextFile)
}
