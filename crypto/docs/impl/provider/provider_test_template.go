package provider

import (
	"testing"
)

type Expected = map[string]any

type ICryptoProviderTest interface {
	TestGetKeys(*testing.T, ICryptoProvider, Expected)
	TestGetSigner(*testing.T, ICryptoProvider)
}

func TestCriptoProvider() {
	builder := Secp256k1Builder
	pk := [32]byte{}
	_, err := rand.Read(pk[:])
	p, err := builder.FromSeed(pk[:])
	if err != nil {
		t.Fatalf(err.Error())
	}
}
