package crypt

type Ciphertext []byte

const (
	NONCE_SIZE = 12 // AES-GCM
)

func NewCiphertext(nonce, text []byte) Ciphertext {
	ct := make(Ciphertext, NONCE_SIZE+len(text))
	copy(ct[:NONCE_SIZE], nonce)
	copy(ct[NONCE_SIZE:], text)

	return ct
}
