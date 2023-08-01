package ecdsa

import (
	"github.com/cosmos/cosmos-sdk/crypto/typesNew"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
)

var Signer = Secp256k1Ecdsa{}

type Secp256k1Ecdsa struct{}

func (s Secp256k1Ecdsa) Sign(hash []byte, priv typesNew.PrivKey) (typesNew.Signature, error) {
	p := secp256k1.PrivKeyFromBytes(priv.Bytes())
	sig := ecdsa.SignCompact(p, hash, false)

	// remove the first byte which is compactSigRecoveryCode
	return Signature{Sig: sig[1:]}, nil
}
