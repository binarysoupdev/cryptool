package host

import (
	"github.com/binarysoupdev/cryptool/net/conn"
)

func (h Host) Accept() (*conn.Conn, error) {
	c, err := h.Listener.Accept()
	if err != nil {
		return nil, err
	}

	return conn.New(c), nil
}
