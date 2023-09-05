package keyring

import (
	cryptoprovider "github.com/cosmos/cosmos-sdk/crypto/docs/crypto_provider"
	"github.com/cosmos/cosmos-sdk/crypto/docs/secure_storage"
)

type ConfigLoader interface {
	LoadConfig() (error, []secure_storage.SecureStorageSourceConfig)
}

type KeyStore interface {
	RegisterStorageSource(string, secure_storage.SecureStorageBuilder)
	Init([]secure_storage.SecureStorageSourceConfig)

	DeleteSource(name string) error
	ListSources() ([]secure_storage.SecureStorageSourceMetadata, error)
	AddSource(secure_storage.SecureStorageSourceConfig) error
	GetSource(string) (secure_storage.SecureStorage, error)
}

type Keyring interface {
	Init(ConfigLoader, KeyStore)

	// CryptoProvider management
	RegisterProvider(string, cryptoprovider.ProviderBuilder)
	RegisterStorage(string, KeyStore)

	// NewMnemonic generates a new mnemonic, derives a hierarchical deterministic key from it, and
	// persists the key to storage. Returns the generated mnemonic and the key Info.
	// It returns an error if it fails to generate a key for the given algo type, or if
	// another key is already stored under the same name or address.
	NewMnemonic()

	// NewAccount converts a mnemonic to a private key and HD Path and persists it.
	// It fails if there is an existing key Info with the same address.
	NewAccount()

	Keys() ([]string, error)

	Delete(string) error
}
