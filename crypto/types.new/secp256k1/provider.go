package secp256k1

import (
	"errors"
	"github.com/cosmos/cosmos-sdk/crypto/types.new"
	"github.com/cosmos/cosmos-sdk/crypto/types.new/secp256k1/ecdsa"
)

type Secp256k1Provider struct {
	signer    types_new.Signer
	verifier  types_new.Verifier
	generator types_new.Generator
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

func (s Secp256k1Provider) GetSigner() (types_new.Signer, error) {
	return s.signer, nil
}

func (s Secp256k1Provider) GetVerifier() (types_new.Verifier, error) {
	return s.verifier, nil
}

func (s Secp256k1Provider) GetGenerator() (types_new.KeyGenerator, error) {
	return s.generator, nil
}

func (s Secp256k1Provider) GetCipher() (types_new.Encryptor, error) {
	return nil, errors.New("cant provide encryptor")
}

func (s Secp256k1Provider) GetHasher() (types_new.Hasher, error) {
	return nil, errors.New("cant provide hasher")
}
