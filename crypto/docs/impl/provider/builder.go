package provider

import "cryptoImpl/storage"

type ICryptoProviderBuilder interface {
	GetBuilderTypeUUID() string
	FromSecureItem(item storage.ISecureItem) (ICryptoProvider, error)
	FromSeed(seed []byte) (ICryptoProvider, error)
	FromMnemonic(mnemonic string) (ICryptoProvider, error)
	FromString(url string) (ICryptoProvider, error)
}
