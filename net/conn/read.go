package conn

import (
	"encoding/binary"
	"io"

	"github.com/binarysoupdev/go-extensions/errors"
	io_exts "github.com/binarysoupdev/go-extensions/io"
)

func (c Conn) ReadMessage() ([]byte, error) {
	header, err := io_exts.ReadBytes(c.Conn, 4)
	if err != nil {
		return nil, errors.Chain(err, "error reading header")
	}

	msg, err := io_exts.ReadBytes(c.Conn, int(binary.BigEndian.Uint32(header)))
	if err != nil {
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
