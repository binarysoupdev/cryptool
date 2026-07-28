package crypt

import (
	"encoding/json"
	"io"
)

func Encode[T any](c Crypt, w io.Writer, obj T) (int, error) {
	ciphertext, err := Marshal(c, obj)
	if err != nil {
		return 0, err
	}

	return w.Write(ciphertext)
}

func Marshal[T any](c Crypt, obj T) ([]byte, error) {
	plaintext, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	return c.Encrypt(plaintext), nil
}
