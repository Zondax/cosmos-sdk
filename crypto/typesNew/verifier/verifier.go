package verifier

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/keys"
	"github.com/cosmos/cosmos-sdk/crypto/typesNew/signer"
)

type SampleVerifier struct {
}

func (v SampleVerifier) Verify(msg []byte, sig signer.Signature, pk keys.PubKey) (bool, error) {

	switch sig.(type) {
	case *signer.MultiSig:
		fmt.Println("Some multisig")
	case *signer.SampleSignature:
		sig.Bytes()
		return true, nil
	}

	return true, nil
}
