package cryptoprovider

import (
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/cypher"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/keys"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/signer"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/verifier"
)

type CryptoProviderOptions interface {
	CanProvidePubKey() bool
	CanProvidePrivKey() bool
	CanExport() bool
	CanSign() bool
	CanVerify() bool
	CanCipher() bool
	CanGenerate() bool
}

type CryptoProvider interface {
	CryptoProviderOptions

	Load(SecureItem) error // Builds the corresponding provider

	GetSigner() (signer.Signer, error)
	GetVerifier() (verifier.Verifier, error)
	GetGenerator() (keys.KeyGenerator, error)
	GetCipher() (cypher.Cypher, error)
	GetHasher() (Hasher, error)
	Wipe()
}

type CryptoProviderBuilder struct {
}

func (c CryptoProviderBuilder) Build(providerItem SecureItem) CryptoProvider {
	switch providerItem.Metadata.Type {
	case "Ledger":
		ledger := NewLedgerProvider() // returns a CryptoProvider
		ledger.Load(providerItem)
		return ledger

	case "Keychain":
		keychain := NewKeychainProvider() // returns a CryptoProvider
		keychain.Load(providerItem)
		return keychain
	}

	return nil
}
