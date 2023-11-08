package provider

import "cryptoV2/storage"

type ICryptoProviderBuilder interface {
	GetBuilderTypeUUID() string
	FromSecureItem(item storage.ISecureItem) (ICryptoProvider, error)
	FromSeed(seed []byte) (ICryptoProvider, error)
	FromMnemonic(mnemonic string) (ICryptoProvider, error)
	FromString(url string) (ICryptoProvider, error)
	FromHDPath(hdPath string, coinType uint32, index uint32, hrp string) (ICryptoProvider, error)
}
