package main

import "testing"

func TestXOR(t *testing.T) {
	// Params
	data := []byte("0")
	key := byte(1)

	// Encryption
	ciphertext := make([]byte, len(data))
	for i, b := range data {
		ciphertext[i] = b ^ key
	}
	// Decryption
	for i, b := range ciphertext {
		ciphertext[i] = b ^ key
	}

	// Output result
	t.Log(string(ciphertext))
}
