package crypt

import (
	"encoding/json"
	"io"
)

func Decode[T any](c Crypt, r io.Reader) (T, error) {
	var obj T
	ciphertext, err := io.ReadAll(r)
	if err != nil {
		return obj, err
	}

	return Unmarshal[T](c, ciphertext)
}

func Unmarshal[T any](c Crypt, ciphertext Ciphertext) (T, error) {
	var obj T

	plaintext, err := c.Decrypt(ciphertext)
	if err != nil {
		return obj, err
	}

	err = json.Unmarshal(plaintext, &obj)
	if err != nil {
		return obj, err
	}

	return obj, nil
}
