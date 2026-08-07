package host

import (
	"fmt"

	"github.com/binarysoupdev/cryptool/net/conn"
	"github.com/binarysoupdev/go-extensions/errors"
)

func (h Host) Listen(hn Handler) {
	for {
		c, err := h.Listener.Accept()
		if err != nil {
			hn.Log(S_ERROR, errors.Chain(err, "error accepting client").Error())
			continue
		}

		h.accept(conn.New(c), hn)
	}
}

func (h Host) accept(c *conn.Conn, hn Handler) {
	hn.Log(S_ACCEPT_CLIENT, fmt.Sprintf("Accepted Client: %s", c.RemoteAddr()))
	defer func() {
		hn.Log(S_LOST_CLIENT, fmt.Sprintf("Lost Client: %s", c.RemoteAddr()))
	}()

	err := hn.Handle(c)
	if err != nil {
		hn.Log(S_ERROR, err.Error())
	}
}
