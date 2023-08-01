package typesNew

import (
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/keys"
	"io"
)

// Generator
// TODO: Still needs to know how to generate from mnemonic
type Generator interface {
	Generate(rand io.Reader) (keys.PrivKey, error)
}
