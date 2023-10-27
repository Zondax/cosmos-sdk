package keyring

import (
	provider2 "cryptoImpl/provider"
	"cryptoImpl/storage"
)

type IKeyring interface {
	RegisterCryptoProviderBuilder(builder provider2.ICryptoProviderBuilder)
	RegisterStorageProvider(provider storage.IStorageProvider)

	ListStorageProviders() error
	ListCryptoProviders() error

	AddCryptoProvider()

	ListItems() ([]storage.SecureItemMetadata, error)

	GetCryptoProvider(string) (provider2.ICryptoProvider, error)
}
