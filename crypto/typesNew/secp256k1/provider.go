package secp256k1

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/secp256k1/ecdsa"
)

type Secp256k1Provider struct {
	signer    typesNew.Signer
	verifier  typesNew.Verifier
	generator typesNew.Generator
}

func NewSecp256k1Provider() *Secp256k1Provider {
	return &Secp256k1Provider{
		signer:    ecdsa.Signer,
		verifier:  ecdsa.Verifier,
		generator: Secp256k1Generator,
	}
}

func (s Secp256k1Provider) CanProvidePubKey() bool {
	return true
}

func (s Secp256k1Provider) CanProvidePrivKey() bool {
	return true
}

func (s Secp256k1Provider) CanSign() bool {
	return true
}

func (s Secp256k1Provider) CanVerify() bool {
	return true
}

func (s Secp256k1Provider) CanCipher() bool {
	return false
}

func (s Secp256k1Provider) CanGenerate() bool {
	return true
}

func (s Secp256k1Provider) GetSigner() (typesNew.Signer, error) {
	return s.signer, nil
}

func (s Secp256k1Provider) GetVerifier() (typesNew.Verifier, error) {
	return s.verifier, nil
}

func (s Secp256k1Provider) GetGenerator() (typesNew.KeyGenerator, error) {
	return s.generator, nil
}

func (s Secp256k1Provider) GetCipher() (typesNew.Encryptor, error) {
	return nil, errors.New("cant provide cypher")
}

func (s Secp256k1Provider) GetHasher() (typesNew.Hasher, error) {
	return nil, errors.New("cant provide hasher")
}
