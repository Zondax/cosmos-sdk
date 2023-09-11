package secure_storage

import "fmt"

type SecureStorageFactoryMap map[string]SecureStorageBuilder

type SecureStorageFactory struct {
	factoryMap SecureStorageFactoryMap
}

func NewSecureStorageFactory() *SecureStorageFactory {
	f := &SecureStorageFactory{
		factoryMap: make(SecureStorageFactoryMap),
	}

	return f
}

func (s *SecureStorageFactory) Build(config SecureStorageSourceConfig) (SecureStorage, error) {
	builder, ok := s.factoryMap[config.Metadata.Type]
	if ok {
		return builder(config)
	} else {
		return nil, fmt.Errorf("unknown secure storage implementation '%s'", config.Metadata.Type)
	}
}

func (s *SecureStorageFactory) RegisterBuilder(storageType string, builder SecureStorageBuilder) {
	s.factoryMap[storageType] = builder
}
