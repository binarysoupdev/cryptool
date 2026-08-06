package client

import (
	"net"

	"github.com/binarysoupdev/cryptool/net/conn"
	"github.com/binarysoupdev/go-extensions/errors"
)

type Client struct {
	*conn.Conn
}

func Connect(addr string) (Client, error) {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		return Client{}, errors.ChainFormat(err, "error dialing host at \"%s\"", addr)
	}

	return Client{
		conn.New(c),
	}, nil
}
