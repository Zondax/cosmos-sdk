package typesNew

import (
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/cypher"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/keys"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/signer"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/verifier"
	"github.com/google/go-cmp/cmp/internal/function"
)

// - Åddress Encoding
// - bech32 / bech32m
// - HRP
// - Address to Keyring reference+links
// - Vanity address?
// - Check address is valid?
type Wallet interface {
	--- Keyring
	--- Keyring
}

-------------

// - Keep REFERENCES to cryptoProvider entities

// - Record (pub/priv keys)  --> reference  URL ledger://    hsm://....      priv://ddd
// - A SignerRecord?? instance should knows how to store itself
// - Armoring - Not always possible (OpenPGP support)
// - Manage mTLS keys???
// https://github.com/hashicorp/go-plugin/blob/main/mtls.go

type Keyring interface {
	Initialize(SecureStorageConfigsFilePath string) error
	ListAvailableStorage() ([]string, error)
	Keys() ([]string, error)
}


type Keyring interface {
	GetProviders() []Providers

	---- Initialize
		-- Register CryptoProvider Types    
				UUID -> CryptoProviderType		
					Ledger				....
					AWS Amazon			....
					1Pass				....
					Local				....
					RemoteProvider      ....

		-- Register ISecureStorage(....)
					1PassSecureStorage(....)
					S3SecureStorage(....)

		SecureStorage.List() --> keys
			SecureStorage.Get(key) --> SecureItem
				SecureItem.UUID --> type -> CryptoProvider.Restore(blob)

	// FIXME: different storage solutions? 99 keyrings
}

							type SecureItemMetadata interface {
								//ModificationTime time.Time
							}

							type SecureItem interface {
								SecureItemMetadata
									UUID			// modification, restrictions, version, etc.
								Blob						// << protobuf / encryption?
							}

							// Can be encrypted or not (using a provider?)
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

/////////////////////////////////////////

type RandomnessSource interface {
}

type Signature interface {
	Blob
	CanAggregate() bool

	// TODO: this is more in terms of BLS or similar
	Aggregate(other []Signature) (Signature, error)
}

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

type CryptoCipher interface {
	Encrypt(data Blob) (Blob, error)
	Decrypt(data Blob) (Blob, error)
}



	// TODO: Define how capabilities can be expressed
	CanProvidePubKey() bool
	CanProvidePrivKey() bool
	CanExport() bool
	CanSign() bool
	CanVerify() bool
	CanCipher() bool
	CanGenerate() bool



init()
	keyring.add(self.UUID, self)
	Restore(blob)...

	// Equivalent to record?
// implement
// - LocalKeyPair
type CryptoProvider interface {
	Options....

	GetSigner() (signer.Signer, error)       //			{ return self }
	GetVerifier() (verifier.Verifier, error) //
	GetGenerator() (keys.KeyGenerator, error)
	GetCipher() (cypher.Cypher, error) // *
	GetHasher() (Hasher, error)        // / ?????
	Wipe()                             // TODO -> CLEAR FUNCTION
}


/////////////////

LocalSecp256k1Provider : CryptoProvider {
	... build (SecureItem / Record /              protobuf -- privatekey)
}

LedgerProvider : CryptoProvider {
	... build (protobuf -- USB0 - pubkey)

	Signer				-- hardware
	Private NO
}}


/////////////////////////////


type LocalProvider interface {
}

type LedgerDevice struct {
	// complies with CryptoProvider
}

// We need to find a way to link metadata to key material to track compatibility
