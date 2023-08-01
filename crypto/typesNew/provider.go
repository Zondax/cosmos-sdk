package typesNew

import (
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/cypher"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/keys"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/signer"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/verifier"
)

type SampleCryptoProvider struct {
	generator Generator
	signer2   signer.Signer
	verifier2 verifier.Verifier
	cypher2   cypher.Cypher
}

func (p SampleCryptoProvider) CanProvidePubKey() bool {
	return true
}

func (p SampleCryptoProvider) CanProvidePrivKey() bool {
	return true
}

func (p SampleCryptoProvider) CanSign() bool {
	return true
}

func (p SampleCryptoProvider) CanVerify() bool {
	return true
}

func (p SampleCryptoProvider) CanCipher() bool {
	return true
}

func (p SampleCryptoProvider) CanGenerate() bool {
	return true
}

func (p SampleCryptoProvider) GetSigner() (signer.Signer, error) {
	return nil, nil
}

func (p SampleCryptoProvider) GetVerifier() (verifier.Verifier, error) {
	return nil, nil
}

func (p SampleCryptoProvider) GetGenerator() (keys.KeyGenerator, error) {
	return nil, nil
}

func (p SampleCryptoProvider) GetCipher() (cypher.Cypher, error) {
	return nil, nil
}

func (p SampleCryptoProvider) GetHasher() (Hasher, error) {
	return nil, nil
}
