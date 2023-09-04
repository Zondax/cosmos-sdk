package keyring

import (
	cryptoprovider "github.com/cosmos/cosmos-sdk/crypto/docs/crypto_provider"
	"github.com/cosmos/cosmos-sdk/crypto/docs/secure_storage"
)

type CosmosKeyring struct {
	// here we store the available SecureStorage instances where SecureItems are stored
	secureStorageSources map[string]*secure_storage.SecureStorage

	// factory to build CryptoProvider instances
	cryptoProviderFactory *cryptoprovider.CryptoProviderFactory
}

func (k CosmosKeyring) Init(loader ConfigLoader) {
	// Load configuration
	err, storageConfigs := loader.LoadConfig()
	if err != nil {
		return
	}

	secureStorageFactory := secure_storage.NewSecureStorageFactory()
	// Register SecureStorage implementations
	//secureStorageFactory.RegisterBuilder("file", secure_storage.NewFileStorageBuilder())
	//secureStorageFactory.RegisterBuilder("memory", secure_storage.NewMemoryStorageBuilder())

	// Load SecureStorage instances according to configuration
	for _, config := range storageConfigs {
		storage, err := secureStorageFactory.Build(config)
		if err != nil {
			return
		}

		k.secureStorageSources[config.Metadata.Name] = &storage
	}

	// Create CryptoProviderFactory, each implementation will register its own builder at init time
	k.cryptoProviderFactory = cryptoprovider.NewCryptoProviderFactory()
}

func (k CosmosKeyring) RegisterProvider(uuid string, builder cryptoprovider.ProviderBuilder) {
	k.cryptoProviderFactory.RegisterBuilder(uuid, builder)
}
