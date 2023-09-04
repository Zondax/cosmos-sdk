package keyring

import (
	cryptoprovider "github.com/cosmos/cosmos-sdk/crypto/docs/crypto_provider"
	"github.com/cosmos/cosmos-sdk/crypto/docs/secure_storage"
)

type ConfigLoader interface {
	LoadConfig() (error, []secure_storage.SecureStorageSourceConfig)
}

type Keyring interface {
	Init(ConfigLoader)

	// SecureStorage management
	RegisterStorageSource(string, secure_storage.SecureStorageBuilder)
	DeleteSource(name string) error
	ListSources() ([]secure_storage.SecureStorageSourceMetadata, error)
	AddSource(secure_storage.SecureStorageSourceConfig) error
	GetSource(string) (secure_storage.SecureStorage, error)

	// CryptoProvider management
	RegisterProvider(string, cryptoprovider.ProviderBuilder)
	GetProvider(string) (cryptoprovider.CryptoProvider, error)

	// List keys on all available SecureStorage instances
	Keys() ([]string, error)
}
