package secp256k1

import (
	types_new "github.com/cosmos/cosmos-sdk/crypto/types.new"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
)

type PrivateKey struct {
	key []byte
}

func (p PrivateKey) Bytes() []byte {
	return p.key
}

func (p PrivateKey) Pubkey() types_new.PubKey {
	pubkeyObject := secp256k1.PrivKeyFromBytes(p.key).PubKey()
	pk := pubkeyObject.SerializeCompressed()
	return PubKey{key: pk}
}
