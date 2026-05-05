// Package crypt streamlines the process of encrypting and decrypting data.
// It leverages the AES encryption algorithm with PBKDF2 password-based key derivation to provide a user friendly yet secure means of data protection.
//
// Basic Usage:
//
//	c, salt := crypt.NewFromPassword("Pass123!")
//
//	ciphertext := c.Encrypt([]byte("lorem ipsum"))
//	fmt.Println(ciphertext)
//
//	plaintext, err := c.Decrypt(ciphertext)
//	fmt.Println(string(plaintext))
package crypt
