package keyring

import (
	"cryptoImpl/crypto/provider"
	"cryptoImpl/storage"
	"errors"
)

type IKeyring interface {
	// TODO: review
	RegisterCryptoProvider(builder provider.ICryptoProviderBuilder)
	RegisterStorageProvider(provider storage.IStorageProvider)

	ListStorageProviders() ([]storage.IStorageProvider, error)
	ListCryptoProviders() ([]provider.ICryptoProviderBuilder, error)

	List() ([]storage.ISecureItemMetadata, error)

	GetCryptoProvider(string) (provider.ICryptoProvider, error)
}

type Keyring struct {
	cp map[string]provider.ICryptoProviderBuilder
	sp map[string]storage.IStorageProvider
}

func New() *Keyring {
	return &Keyring{
		cp: make(map[string]provider.ICryptoProviderBuilder),
		sp: make(map[string]storage.IStorageProvider),
	}
}

func (k Keyring) RegisterCryptoProvider(builder provider.ICryptoProviderBuilder) {
	k.cp[builder.GetTypeUUID()] = builder
}

func (k Keyring) RegisterStorageProvider(provider storage.IStorageProvider) {
	k.sp[provider.Name()] = provider
}

func (k Keyring) ListStorageProviders() ([]storage.IStorageProvider, error) {
	r := make([]storage.IStorageProvider, len(k.sp))
	i := 0
	for _, v := range k.sp {
		r[i] = v
		i++
	}
	return r, nil
}

func (k Keyring) ListCryptoProviders() ([]provider.ICryptoProviderBuilder, error) {
	r := make([]provider.ICryptoProviderBuilder, len(k.cp))
	i := 0
	for _, v := range k.cp {
		r[i] = v
		i++
	}
	return r, nil
}

func (k Keyring) List() ([]storage.ISecureItemMetadata, error) {
	return nil, nil
}

func (k Keyring) GetCryptoProvider(s string) (provider.ICryptoProvider, error) {
	for _, v := range k.sp {
		for _, i := range v.List() {
			if i == s {
				si, err := v.Get(s)
				if err != nil {
					return nil, err
				}
				return k.cp[si.TypeUuid].FromSecureItem(si)
			}
		}
	}
	return nil, errors.New("provider not found")
}
