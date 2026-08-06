package conn

import (
	"net"

	"github.com/binarysoupdev/cryptool/crypt"
)

type Connection struct {
	net.Conn
	crypt     crypt.Crypt
	msgBuffer []byte
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{
		Conn: conn,
	}
}

func (c Connection) IsEncrypted() bool {
	return !c.crypt.IsNil()
}
