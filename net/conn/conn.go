package conn

import (
	"net"

	"github.com/binarysoupdev/cryptool/crypt"
)

type Conn struct {
	net.Conn
	crypt     crypt.Crypt
	msgBuffer []byte
}

func New(conn net.Conn) *Conn {
	return &Conn{
		Conn: conn,
	}
}

func (c Conn) IsEncrypted() bool {
	return !c.crypt.IsNil()
}
