package cryptoprovider

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/docs/secure_item"
)

type ProviderBuilder func(item secure_item.SecureItem) (CryptoProvider, error)
type ProviderFactoryMap map[string]ProviderBuilder

type CryptoProviderFactory struct {
	factoryMap ProviderFactoryMap
}

func NewCryptoProviderFactory() *CryptoProviderFactory {
	return &CryptoProviderFactory{
		factoryMap: make(ProviderFactoryMap),
	}
}

func (f *CryptoProviderFactory) Build(item secure_item.SecureItem) (CryptoProvider, error) {
	builder, ok := f.factoryMap[item.Metadata.UUID]
	if ok {
		return builder(item)
	} else {
		return nil, fmt.Errorf("unknown provider UUID '%s'", item.Metadata.UUID)
	}
}

func (f *CryptoProviderFactory) RegisterBuilder(uuid string, builder ProviderBuilder) {
	f.factoryMap[uuid] = builder
}
