package generator

type BaseKey interface {
	String() string
	Bytes() []byte
}

type PubKey interface {
	BaseKey
	Address() []byte
}

type PrivKey interface {
	BaseKey
	Pubkey() PubKey
}

// KeyGenerator Generate/Derive
// - From hardware
// - From pure entropy (KDF)
// - From previous key material + metadata (BIP44)
// - Retrieve instance from keyring
type KeyGenerator interface {
	// TODO: comments
	New() KeyPair
	Derive(kp KeyPair) KeyPair
}

type KeyPair interface {
	PubKey
	PrivKey
}
