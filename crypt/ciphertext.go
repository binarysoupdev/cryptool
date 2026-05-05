package crypt

const NONCE_SIZE = 12 // AES-GCM

// Ciphertext is a raw byte slice containing both the nonce and the encrypted text.
type Ciphertext []byte

// Return the nonce portion of the ciphertext.
func (ct Ciphertext) Nonce() []byte {
	return ct[:NONCE_SIZE]
}

// Return the text portion of the ciphertext.
func (ct Ciphertext) Text() []byte {
	return ct[NONCE_SIZE:]
}
