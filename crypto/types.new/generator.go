package types_new

import (
	"io"
)

// Generator
// TODO: Still needs to know how to generate from mnemonic
type Generator interface {
	Generate(rand io.Reader) (PrivKey, error)
}
