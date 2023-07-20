package types_new

type DerivationMetadata interface {
}

type BIP32Path interface {
	DerivationMetadata
}

type BIP44Path interface {
	BIP32Path
}
