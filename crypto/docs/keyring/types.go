package keyring

import "time"

type SecureStorageSourceMetadata struct {
	Type string
	Name string
}

type SecureStorageSourceConfig struct {
	Metadata SecureStorageSourceMetadata
	Config   any // specific config for the desired backend, if necessary
}

type SecureItemMetadata struct {
	ModificationTime time.Time
	UUID             string
}

type SecureItem struct {
	Metadata SecureItemMetadata
	Blob     []byte // versioned protobuf blob
}

type SecureStorage interface {
	Get(key string) (SecureItem, error)
	Set(key string, item SecureItem) error
	Delete(key string) error
	List() ([]string, error)
}

type Keyring interface {
	Init() //  registers all available secure backends (keychain, aws, 1pass, etc)

	DeleteSecureStorageSource(name string) error
	ListSecureStorageSources() ([]SecureStorageSourceMetadata, error)
	AddSecureStorageSource(config SecureStorageSourceConfig) error
	GetSecureStorageSource(name string) (SecureStorage, error)
	Keys() ([]string, error) // calls List() for every SecureStorage
}
