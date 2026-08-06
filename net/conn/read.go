package conn

import (
	"encoding/binary"
	"io"

	"github.com/binarysoupdev/go-extensions/errors"
)

func (c Conn) ReadMessage() ([]byte, error) {
	header := make([]byte, 4)
	if _, err := io.ReadFull(c.Conn, header); err != nil {
		return nil, errors.Chain(err, "error reading header")
	}

	msg := make([]byte, int(binary.BigEndian.Uint32(header)))
	if _, err := io.ReadFull(c.Conn, msg); err != nil {
		return nil, errors.Chain(err, "error reading message")
	}

	if c.crypt.IsNil() {
		return msg, nil
	}

	plaintext, err := c.crypt.Decrypt(msg)
	if err != nil {
		return nil, errors.Chain(err, "error decrypting message")
	}

	return plaintext, nil
}

func (c *Conn) Read(b []byte) (int, error) {
	var err error

	if len(c.msgBuffer) == 0 {
		c.msgBuffer, err = c.ReadMessage()
		if err != nil {
			return 0, err
		}
	}

	n := copy(b, c.msgBuffer)

	if n >= len(c.msgBuffer) {
		c.msgBuffer = []byte{}
		return n, io.EOF
	}

	c.msgBuffer = c.msgBuffer[n:]
	return n, nil
}
