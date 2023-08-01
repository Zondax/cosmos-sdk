package signer

import (
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/keys"
)

// Signer (persistence)
//   - From a keypair object
//   - From some external reference (remote, etc.)
//   - Retrieve instance from keyring
//   - Example: Ledger may keep a pubkey reference that is checked. Locally it can be imported
type Signer interface {
	//New(reference typesNew.CryptoProvider)
	Sign(
		hash []byte, priv keys.PrivKey) (Signature, error)
}

type Signature interface {
	Bytes() []byte
}
