package host

import (
	"fmt"

	"github.com/binarysoupdev/cryptool/net/conn"
	"github.com/binarysoupdev/go-extensions/errors"
)

func (h Host) Listen(hn Handler, maxClients int) {
	clients := make(chan struct{}, maxClients)
	clientID := 0

	for {
		c, err := h.waitForClient()
		if err != nil {
			hn.Log(S_ERROR, -1, err.Error())
			continue
		}

		if cap(clients) > 0 {
			clients <- struct{}{}
		}
		clientID++

		go h.handleClient(hn, c, clientID, clients)
	}
}

func (h Host) waitForClient() (*conn.Conn, error) {
	c, err := h.Listener.Accept()
	if err != nil {
		return nil, errors.Chain(err, "error accepting client")
	}

	return conn.New(c), nil
}

func (h Host) handleClient(hn Handler, c *conn.Conn, clientID int, clients chan struct{}) {
	defer func() {
		c.Close()
		if cap(clients) > 0 {
			<-clients
		}
	}()

	hn.Log(S_ACCEPT_CLIENT, clientID, fmt.Sprintf("Accepted Client: %s", c.RemoteAddr()))

	err := hn.Handle(clientID, c)
	if err != nil {
		hn.Log(S_ERROR, clientID, err)
	}

	hn.Log(S_LOST_CLIENT, clientID, fmt.Sprintf("Lost Client: %s", c.RemoteAddr()))
}
