package typesNew

import (
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/cypher"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/keys"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/signer"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/verifier"
)

// - Åddress Encoding
// - bech32 / bech32m
// - HRP
// - Address to Keyring reference+links
// - Vanity address?
// - Check address is valid?
type Wallet interface {
}

// - Keep REFERENCES to signer? entities
// - Record (pub/priv keys)  --> reference  URL ledger://    hsm://....      priv://ddd
// - A SignerRecord?? instance should knows how to store itself
// - Armoring
// - Not always possible (OpenPGP support)
// - Manage mTLS keys???
// https://github.com/hashicorp/go-plugin/blob/main/mtls.go
type Keyring interface {
	GetRecords() []KeyringRecord

	// FIXME: different storage solutions? 99 keyrings
}

type SecureItemMetadata interface {
	//ModificationTime time.Time
}

type SecureItem interface {
	Blob
}

type SecureStorage interface {
	Get(key string) (SecureItem, error)
	Set(key string, item SecureItem) error
	Delete(key string) error
	List() ([]string, error)

	// Get(key string) (Item, error)
	// // Returns the non-secret parts of an Item
	// GetMetadata(key string) (Metadata, error)
	// // Stores an Item on the keyring
	// Set(item Item) error
	// // Removes the item with matching key
	// Remove(key string) error
	// // Provides a slice of all keys stored on the keyring
	// Keys() ([]string, error)
}

type RandomnessSource interface {
}

//type Signature interface {
//	Blob
//	CanAggregate() bool
//
//	// TODO: this is more in terms of BLS or similar
//	Aggregate(other []Signature) (Signature, error)
//}

type Hasher interface {
	Hash(input Blob) Blob

	// ??? Add support for incremental hashing (keep internal state)
	CanHashIncrementally() bool
}

type Blob interface {
	// TODO: consider some proper zeroing approach. At least consolidate things here?
	Bytes() []byte
	NewBlob(data []byte) *Blob
}

// - Digest
//   - F(Blob, hashFunction)
type Digest interface {
	Blob
}

// Equivalent to record?
type KeyringRecord interface {
	CryptoProvider // ledger, localKP, remoteKP, etc.

	EncryptionProvider() // localKP, remoteKP, etc.
	Store() error
	Restore() error
}

type CryptoCipher interface {
	Encrypt(data Blob) (Blob, error)
	Decrypt(data Blob) (Blob, error)
}

// Equivalent to record?
// implement
// - LocalKeyPair
type CryptoProvider interface {
	CanProvidePubKey() bool
	CanProvidePrivKey() bool
	CanSign() bool
	CanVerify() bool
	CanCipher() bool
	CanGenerate() bool

	GetSigner() (signer.Signer, error)       //
	GetVerifier() (verifier.Verifier, error) //
	GetGenerator() (keys.KeyGenerator, error)
	GetCipher() (cypher.Cypher, error) // *
	GetHasher() (Hasher, error)        // / ?????
}

type LocalProvider interface {
}
type SecureElement interface {
	CryptoProvider
	Wipe() // TODO -> CLEAR FUNCTION
}

type LedgerDevice interface {
	SecureElement
}

type SecretKeypair interface {
	CryptoProvider
	PubKey() keys.PubKey
	PrivKey() keys.PrivKey
}

// --------------
