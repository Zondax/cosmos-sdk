package keyring

import (
	provider2 "cryptoV2/provider"
	"cryptoV2/storage"
)

type IKeyring interface {
	RegisterCryptoProviderBuilder(builder provider2.ICryptoProviderBuilder)
	RegisterStorageProvider(provider storage.IStorageProvider)

	ListItems() ([]storage.SecureItemMetadata, error)

	AddCryptoProvider() // TODO
	GetCryptoProvider(string) (provider2.ICryptoProvider, error)
}
