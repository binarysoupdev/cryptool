package crypt

type Ciphertext []byte

const (
	NONCE_SIZE = 12 // AES-GCM
	SALT_SIZE  = 16
)

func NewCiphertext(salt, nonce, text []byte) Ciphertext {
	ct := make(Ciphertext, SALT_SIZE+NONCE_SIZE+len(text))
	copy(ct.Salt(), salt)
	copy(ct.Nonce(), nonce)
	copy(ct.Text(), text)

	return ct
}

func (ct Ciphertext) Salt() []byte {
	return ct[:SALT_SIZE]
}

func (ct Ciphertext) Nonce() []byte {
	return ct[SALT_SIZE : SALT_SIZE+NONCE_SIZE]
}

func (ct Ciphertext) Text() []byte {
	return ct[SALT_SIZE+NONCE_SIZE:]
}
