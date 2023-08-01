package secp256k1

import (
	"crypto/elliptic"
	types_new "github.com/cosmos/cosmos-sdk/crypto/typesNew"
	dsecp "github.com/decred/dcrd/dcrec/secp256k1/v4"
	"io"
)

var (
	_                  types_new.Generator = Generator{}
	Secp256k1Generator                     = Generator{c: dsecp.S256()}
)

type Generator struct {
	c elliptic.Curve
}

func (s Generator) Generate(rand io.Reader) (types_new.PrivKey, error) {
	pb, _, _, err := elliptic.GenerateKey(s.c, rand)
	return PrivateKey{key: pb}, err
}
