package ecdsa

type Signature struct {
	Sig []byte
}

func (s Signature) Bytes() []byte {
	return s.Sig
}
