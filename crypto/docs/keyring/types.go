package keyring

import (
	cryptoprovider "github.com/cosmos/cosmos-sdk/crypto/docs/crypto_provider"
	"github.com/cosmos/cosmos-sdk/crypto/docs/secure_storage"
)

type Keyring interface {
	RegisterCryptoProvider(string, cryptoprovider.ProviderBuilder)
	RegisterSecureStorage(string, secure_storage.SecureStorageBuilder)

	GetCryptoProvider(key string) (cryptoprovider.CryptoProvider, error)
	Keys() ([]string, error)
}
