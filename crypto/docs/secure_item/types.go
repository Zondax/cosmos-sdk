package secure_item

import "time"

type SecureItemMetadata struct {
	ModificationTime time.Time
	UUID             string
}

type SecureItem struct {
	Metadata SecureItemMetadata
	// Blob format/encoding will be dependant of the CryptoProvider implementation
	Blob []byte
}
