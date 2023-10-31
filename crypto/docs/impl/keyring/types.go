package keyring

import (
	provider2 "cryptoV2/provider"
	"cryptoV2/storage"
)

type IKeyring interface {
	RegisterCryptoProviderBuilder(builder provider2.ICryptoProviderBuilder)
	RegisterStorageProvider(provider storage.IStorageProvider)

	ListCryptoProviders() ([]storage.SecureItemMetadata, error)
	AddCryptoProvider(provider2.ICryptoProvider) error
	GetCryptoProvider(string) (provider2.ICryptoProvider, error)
}
