package storage

type ISecureItemMetadata interface {
	Type() string // creates typeUUid   // Relates to the corresponding crypto
	Name() string
}

type ISecureItem interface {
	ISecureItemMetadata

	// Blob format/encoding will be dependant of the CryptoProvider implementation
	Bytes() []byte
}

type SecureItem struct {
	TypeUuid string
	NameId   string
	Blob     []byte
}

func NewSecureItem(uuid, name string, blob []byte) *SecureItem {
	return &SecureItem{
		TypeUuid: uuid,
		NameId:   name,
		Blob:     blob,
	}
}

func (s SecureItem) Type() string {
	return s.TypeUuid
}

func (s SecureItem) Name() string {
	return s.NameId
}

func (s SecureItem) Bytes() []byte {
	return s.Blob
}
