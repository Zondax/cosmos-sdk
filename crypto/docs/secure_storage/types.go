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
	// Build builds the corresponding secure storage backend
	Build(SecureStorageSourceConfig) (SecureStorage, error)

	// Get returns the SecureItem matching the key or ErrKeyNotFound
	Get(string) (secure_item.SecureItem, error)
	// GetMetadata returns the metadata field of the SecureItem
	GetMetadata(string) (secure_item.SecureItemMetadata, error)
	// Set stores the SecureItem on the backend
	Set(string, secure_item.SecureItem) error
	// Remove removes the SecureItem matching the key
	Remove(string) error
	// Keys returns a slice of all keys stored on the backend
	Keys() ([]string, error)
}
