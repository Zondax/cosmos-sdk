package secp256k1

import (
	crand "crypto/rand"
	"fmt"
	"github.com/cometbft/cometbft/crypto"
	ecdsa2 "github.com/cosmos/cosmos-sdk/crypto/types.new/secp256k1/ecdsa"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGenerator(t *testing.T) {
	priv, err := Secp256k1Generator.Generate(crand.Reader)
	require.NoError(t, err)
	fmt.Println(priv.Bytes())
	pub := priv.Pubkey()
	fmt.Println(pub.Bytes())

	msg := []byte("This is a message about to be encrypted.")
	hash := crypto.Sha256(msg)
	signature, err := ecdsa2.Signer.Sign(hash, priv)
	require.NoError(t, err)
	fmt.Println(signature.Bytes())

	is, err := ecdsa2.Verifier.Verify(msg, signature, pub)
	require.NoError(t, err)
	require.True(t, is)
}

func TestProvider(t *testing.T) {
	//p := NewSecp256k1Provider()

}
