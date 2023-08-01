package cypher

type TestCypher struct {
	Cypher
}

func (cypher *TestCypher) Encrypt(message []byte, secret []byte) (ciphertext []byte, err error) {
	return message, nil
}

// secret must be 32 bytes long. Use something like Sha256(Bcrypt(passphrase))
// The ciphertext is (secretbox.Overhead + 24) bytes longer than the plaintext.
func (cypher *TestCypher) Decrypt(message []byte, secret []byte) (plaintext []byte, err error) {
	return plaintext, nil
}
