package keyring

import (
	provider2 "cryptoV2/provider"
	"cryptoV2/storage"
)

type IKeyring interface {
	// RegisterCryptoProviderBuilder registers a new CryptoProviderBuilder
	RegisterCryptoProviderBuilder(builder provider2.ICryptoProviderBuilder)
	// RegisterStorageProvider registers a new StorageProvider
	RegisterStorageProvider(provider storage.IStorageProvider)

	// List returns a slice of all CryptoProviders stored on the keyring
	List() ([]storage.SecureItemMetadata, error)
	// Add adds a new CryptoProvider to the keyring
	Add(provider2.ICryptoProvider) error
	// Get returns a CryptoProvider from the keyring
	Get(string) (provider2.ICryptoProvider, error)
}
