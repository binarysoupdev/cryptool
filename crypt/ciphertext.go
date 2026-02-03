package crypt

type Ciphertext []byte

const NONCE_SIZE = 12 // AES-GCM

func NewCiphertext(nonce, text []byte) Ciphertext {
	ct := make(Ciphertext, NONCE_SIZE+len(text))
	copy(ct.Nonce(), nonce)
	copy(ct.Text(), text)

	return ct
}

func (ct Ciphertext) Nonce() []byte {
	return ct[:NONCE_SIZE]
}

func (ct Ciphertext) Text() []byte {
	return ct[NONCE_SIZE:]
}
