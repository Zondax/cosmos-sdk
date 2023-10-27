package keyring

import (
	provider2 "cryptoImpl/provider"
	"cryptoImpl/storage"
	"fmt"
	"go.uber.org/zap"
)

type Keyring struct {
	cp     map[string]provider2.ICryptoProviderBuilder
	sp     map[string]storage.IStorageProvider
	buffer map[string]map[string]storage.ItemId // map[providerName]map[SecureItemName]ItemId
}

func New() *Keyring {
	return &Keyring{
		cp:     make(map[string]provider2.ICryptoProviderBuilder),
		sp:     make(map[string]storage.IStorageProvider),
		buffer: make(map[string]map[string]storage.ItemId),
	}
}

func (k *Keyring) RegisterCryptoProviderBuilder(builder provider2.ICryptoProviderBuilder) {
	t := builder.GetBuilderTypeUUID()
	k.cp[t] = builder
	zap.S().Infof("Registered crypto provider builder %s", t)
}

func (k *Keyring) RegisterStorageProvider(provider storage.IStorageProvider) {
	k.sp[provider.Name()] = provider
	k.syncProviderWithBuffer(provider)
}

func (k *Keyring) ListStorageProviders() error {
	fmt.Println("Registered storage providers:")
	for _, v := range k.sp {
		fmt.Println("- Provider: ", v.Name())
	}
	return nil
}

func (k *Keyring) ListCryptoProviderBuilders() error {
	fmt.Println("Registered crypto provider builders:")
	for _, v := range k.cp {
		fmt.Println("- Builder: ", v.GetBuilderTypeUUID())
	}
	return nil
}

func (k *Keyring) GetCryptoProvider(name string) (provider2.ICryptoProvider, error) {
	// search the item in buffer
	for provider, v := range k.buffer {
		if _, ok := v[name]; ok {
			item := v[name]
			secureItem, err := k.sp[provider].Get(item)
			if err != nil {
				return nil, err
			}

			// get the proper builder
			builder, ok := k.cp[item.Type]
			if !ok {
				return nil, fmt.Errorf("no builder found for type: %s", item.UUID)
			}

			cp, err := builder.FromSecureItem(secureItem)
			if err != nil {
				return nil, err
			}

			return cp, nil
		}
	}

	return nil, fmt.Errorf("no crypto provider found with name: %s", name)
}

func (k *Keyring) List() ([]storage.SecureItemMetadata, error) {
	var metadataList []storage.SecureItemMetadata
	for _, v := range k.sp {
		items := v.List()
		for _, item := range items {
			si, err := v.Get(item)
			if err != nil {
				return nil, err
			}
			metadataList = append(metadataList, si.Metadata())
		}
	}
	return metadataList, nil
}

func (k *Keyring) syncProviderWithBuffer(provider storage.IStorageProvider) error {
	items := provider.List()
	k.buffer[provider.Name()] = make(map[string]storage.ItemId, len(items))
	for i, v := range items {
		k.buffer[provider.Name()][v.Name()] = items[i]
	}
	return nil
}
