package secure_item

import "time"

type ItemId struct {
	UUID string
	Slot string
}

type SecureItemMetadata struct {
	ModificationTime time.Time
	ItemId           ItemId
}

type SecureItem struct {
	Metadata SecureItemMetadata
	// Blob format/encoding will be dependant of the CryptoProvider implementation
	Blob []byte
}
