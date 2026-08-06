package conn

import (
	"encoding/binary"

	"github.com/binarysoupdev/go-extensions/errors"
)

func (c Connection) WriteMessage(msg []byte) error {
	_, err := c.Write(msg)
	return err
}

func (c Connection) Write(b []byte) (int, error) {
	if !c.crypt.IsNil() {
		b = c.crypt.Encrypt(b)
	}

	header := make([]byte, 4)
	binary.BigEndian.PutUint32(header, uint32(len(b)))

	_, err := c.Conn.Write(header)
	if err != nil {
		return 0, errors.Chain(err, "error writing header")
	}

	n, err := c.Conn.Write(b)
	if err != nil {
		return 0, errors.Chain(err, "error writing message")
	}

	return n, nil
}
