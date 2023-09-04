package secure_storage

import "github.com/cosmos/cosmos-sdk/crypto/docs/secure_item"

type SecureStorageSourceMetadata struct {
	Type string
	Name string
}

type SecureStorageSourceConfig struct {
	Metadata SecureStorageSourceMetadata
	Config   any // specific config for the desired backend, if necessary
}

type SecureStorageBuilder func(SecureStorageSourceConfig) (SecureStorage, error)

type SecureStorage interface {
	Build(SecureStorageSourceConfig) (SecureStorage, error)

	Get(string) (secure_item.SecureItem, error)
	Store(string, secure_item.SecureItem) error
	Delete(string) error
	List() ([]string, error)
}
