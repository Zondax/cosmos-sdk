package secp256k1

type PubKey struct {
	key []byte
}

func (p PubKey) Address() []byte {
	return nil
}

func (p PubKey) Bytes() []byte {
	return p.key
}
