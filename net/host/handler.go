package host

import (
	"github.com/binarysoupdev/cryptool/net/conn"
)

const (
	S_ERROR = iota
	S_ACCEPT_CLIENT
	S_LOST_CLIENT
)

type Handler interface {
	Log(status int, msg string)
	Handle(c *conn.Conn) error
}
