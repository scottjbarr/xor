package xor

// EncryptDecrypt performs XOR encryption on the input.
//
// If the input is plaintext, the result is an encrypted payload. If the
// input is encrypted, the plaintext is returned.
//
// The key can be any length, with larger keys obviously being preferred.
func EncryptDecrypt(input, key []byte) []byte {
	output := make([]byte, len(input), len(input))

	for i := 0; i < len(input); i++ {
		output[i] = input[i] ^ key[i%len(key)]
	}

	return output
}
