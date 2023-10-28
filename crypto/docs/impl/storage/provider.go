package storage

type IStorageProvider interface {
	Type() string
	Name() string
	List() []ItemId
	Get(id ItemId) (*SecureItem, error)
	Set(item *SecureItem) error
	Remove(id ItemId) error
}
