package types_new

type Signer interface {
	Sign(hash []byte, priv PrivKey) (Signature, error)
}

type Signature interface {
	Bytes() []byte
}
