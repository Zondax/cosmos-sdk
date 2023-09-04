package verifier

import (
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/keys"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/signer"
)

// Verifier
// - Verify Signature + Digest
// - Validate pubkey
// - is on curve https://solanacookbook.com/references/keypairs-and-wallets.html#how-to-check-if-a-public-key-has-an-associated-private-key
// secp256k1
type Verifier interface {
	Verify(hash []byte, sig signer.Signature, pk keys.PubKey) (bool, error)
}
