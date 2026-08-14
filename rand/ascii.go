package rand

import (
	"crypto/rand"
	"math/big"
)

// Fill bytes with a cryptographically random string of printable ASCII characters.
// Could be used as a password string.
func ASCII(bytes []byte) {
	const START = 32 // SPACE
	const END = 126  // ~
	max := big.NewInt(END - START)

	for i := range bytes {
		num, err := rand.Int(rand.Reader, max)
		if err != nil {
			panic(err) // rand.Reader should never fail
		}

		bytes[i] = byte(num.Int64() + START)
	}
}
