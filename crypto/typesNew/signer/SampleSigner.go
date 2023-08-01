package signer

import (
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/keys"
)

type SampleSigner struct {
	Signer
}

//func (s SampleSigner) New(provider typesNew.CryptoProvider) {
//
//}

func (s SampleSigner) Sign(hash []byte, priv keys.PrivKey) (Signature, error) {
	signatureBytes := append(hash, priv.Bytes()...)
	sampleSignature := SampleSignature{signatureBytes}

	return sampleSignature, nil
}

type SampleSignature struct {
	someInfo []byte
}

func (s SampleSignature) Bytes() []byte {
	return s.someInfo
}

type SampleSignerB struct {
	Signer
}

//func (s SampleSigner) New(provider typesNew.CryptoProvider) {
//
//}

func (s SampleSignerB) Sign(hash []byte, priv keys.PrivKey) (Signature, error) {
	signatureBytes := append(priv.Bytes(), priv.Bytes()...)
	sampleSignature := SampleSignature{signatureBytes}

	return sampleSignature, nil
}

type MultiSig struct {
	Signer
}

func (s MultiSig) Bytes() []byte {
	return nil
}
