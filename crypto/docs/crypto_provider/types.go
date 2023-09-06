package cryptoprovider

import (
	"github.com/cosmos/cosmos-sdk/crypto/docs/secure_item"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/cypher"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/keys"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/signer"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/verifier"
)

type ProviderOptions interface {
	CanProvidePubKey() bool
	CanProvidePrivateKey() bool
	CanExport() bool
	CanSign() bool
	CanVerify() bool
	CanCipher() bool
	CanGenerate() bool
}

type CryptoProvider interface {
	ProviderOptions

	Build(item secure_item.SecureItem) (*CryptoProvider, error) // Builds the corresponding provider

	GetSigner() (signer.Signer, error)
	GetVerifier() (verifier.Verifier, error)
	GetGenerator() (keys.KeyGenerator, error)
	GetCipher() (cypher.Cypher, error)
	GetHasher() (Hasher, error)
	Wipe()
}
