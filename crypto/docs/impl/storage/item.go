package storage

import (
	"fmt"
	"time"
)

type ItemId struct {
	Type string
	UUID string
	Slot string
}

func (t ItemId) Name() string {
	return fmt.Sprintf("%s_%s", t.UUID, t.Slot)
}

type SecureItemMetadata struct {
	ModificationTime time.Time
	ItemId           ItemId
}

type ISecureItem interface {
	Metadata() SecureItemMetadata

	// Blob format/encoding will be dependant of the CryptoProvider implementation
	Bytes() []byte
}

type SecureItem struct {
	Meta SecureItemMetadata
	// Blob format/encoding will be dependant of the CryptoProvider implementation
	Blob []byte
}

func NewSecureItem(itemId ItemId, blob []byte) *SecureItem {
	if itemId.UUID == "" {
		fmt.Println("Error: UUID cannot be empty")
		return nil
	}
	return &SecureItem{
		Meta: SecureItemMetadata{
			ModificationTime: time.Now().Round(time.Millisecond),
			ItemId:           itemId,
		},
		Blob: blob,
	}
}

func (s SecureItem) Metadata() SecureItemMetadata {
	return s.Meta
}

func (s SecureItem) UUID() string {
	return s.Meta.ItemId.UUID
}

func (s SecureItem) Slot() string {
	return s.Meta.ItemId.Slot
}

func (s SecureItem) Type() string {
	return s.Meta.ItemId.Type
}

func (s SecureItem) Bytes() []byte {
	return s.Blob
}
