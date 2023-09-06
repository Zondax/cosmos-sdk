package wallet

import (
	"github.com/cosmos/cosmos-sdk/crypto/docs/crypto_provider/verifier"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/signer"
)

type Wallet interface {
	Init(keyring keyring.Keyring)
	GetSigner(address string) signer.Signer
	GetVerifier(address string) verifier.Verifier
	Generate() string
}
