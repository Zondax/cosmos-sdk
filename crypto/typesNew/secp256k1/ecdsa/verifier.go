package ecdsa

import (
	"errors"
	"github.com/cometbft/cometbft/crypto"
	types_new "github.com/cosmos/cosmos-sdk/crypto/typesNew"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
)

var Verifier = VerifierEcdsa{}

type VerifierEcdsa struct{}

func (v VerifierEcdsa) Verify(msg []byte, signature types_new.Signature, pubKey types_new.PubKey) (bool, error) {
	if len(signature.Bytes()) != 64 {
		return false, errors.New("not required length")
	}
	pub, err := secp256k1.ParsePubKey(pubKey.Bytes())
	if err != nil {
		return false, err
	}
	// parse the signature, will return error if it is not in lower-S form
	sig, err := signatureFromBytes(signature.Bytes())
	if err != nil {
		return false, err
	}
	return sig.Verify(crypto.Sha256(msg), pub), nil
}

func signatureFromBytes(sigStr []byte) (*ecdsa.Signature, error) {
	var r secp256k1.ModNScalar
	r.SetByteSlice(sigStr[:32])
	var s secp256k1.ModNScalar
	s.SetByteSlice(sigStr[32:64])
	if s.IsOverHalfOrder() {
		return nil, errors.New("signature is not in lower-S form")
	}

	return ecdsa.NewSignature(&r, &s), nil
}
