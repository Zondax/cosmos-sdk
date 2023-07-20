package types_new

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
	GetRecords() KeyringRecord[]

	// FIXME: different storage solutions? 99 keyrings
}

type SecureItemMetadata interface {
	ModificationTime time.Time
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

// - Generate/Derive
// - From hardware
// - From pure entropy (KDF)
// - From previous key material + metadata (BIP44)
// - Retrieve instance from keyring
type KeyGenerator interface {
	// Can.... ??
	// TODO: comments
}

type Signature interface {
	Blob
	CanAggregate() bool

	// TODO: this is more in terms of BLS or similar
	Aggregate(other []Signature) (Signature, error)
}

// /- Signer  (persistence)
//   - From a keypair object
//   - From some external reference (remote, etc.)
//   - Retrieve instance from keyring
//   - Example: Ledger may keep a pubkey reference that is checked. Locally it can be imported
type Signer interface {
	Verifier
	New(reference CryptoProvider)
	Sign(input Digest) Signature
}

// - Verifier
// - Verify Signature + Digest
// - Validate pubkey
// - is on curve https://solanacookbook.com/references/keypairs-and-wallets.html#how-to-check-if-a-public-key-has-an-associated-private-key
// secp256k1
type Verifier interface {
	Sign(input Digest, signature Signature) (bool, error)
}

type Encryptor interface {
}

type Decryptor interface {
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

	GetSigner() Signer
	GetVerifier() Verifier
	GetGenerator() KeyGenerator
	GetCipher() Encryptor
	GetHasher() Hasher // / ?????
}

type SecureElement interface {
	CryptoProvider
}

type LedgerDevice interface {
	SecureElement
}

type SecretKeypair interface {
	CryptoProvider
	PubKey() PubKey
	PrivKey() PrivKey
}

type PubKey interface {
	// / amino
}

type PrivKey interface {
}

// --------------
