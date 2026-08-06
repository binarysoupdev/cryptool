package host

import (
	"net"

	"github.com/binarysoupdev/go-extensions/errors"
)

type Host struct {
	net.Listener
}

func Open(port string) (Host, error) {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return Host{}, errors.ChainFormat(err, "error starting host on port \"%s\"", port)
	}

	return Host{
		Listener: ln,
	}, nil
}
