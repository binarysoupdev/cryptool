package crypt

type Ciphertext []byte

const NONCE_SIZE = 12 // AES-GCM

func (ct Ciphertext) Nonce() []byte {
	return ct[:NONCE_SIZE]
}

func (ct Ciphertext) Text() []byte {
	return ct[NONCE_SIZE:]
}
