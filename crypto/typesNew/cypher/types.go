package cypher

type Cypher interface {
	Encryptor
	Decryptor
}

type Encryptor interface {
	Encrypt(message []byte, secret []byte) ([]byte, error)
}

type Decryptor interface {
	Decrypt(message []byte, secret []byte) ([]byte, error)
}

type SampleCypher struct {
}

func (s SampleCypher) Encrypt(msg []byte, secret []byte) ([]byte, error) {
	return append(msg, secret...), nil
}

func (s SampleCypher) Decrypt(msg []byte, secret []byte) ([]byte, error) {
	return msg[:32], nil
}
