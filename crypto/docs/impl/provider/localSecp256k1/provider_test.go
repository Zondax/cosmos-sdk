package localSecp256k1

import (
	"cryptoV2/provider"
	"github.com/stretchr/testify/assert"
	"testing"
)

type LocalSecp256k1Test struct{}

func (LocalSecp256k1Test) TestMyProviderGetKeys(t *testing.T, p provider.ICryptoProvider, e provider.Expected) {
	puk, pik, err := p.GetKeys()
	if err != nil {
		t.Fatalf(err.Error())
	}

	expectedPubkey := e["pubkey"]
	expectedPrivkey := e["privkey"]
	assert.Equalf(t, expectedPubkey, puk, "expectedPubkey = %v, pubKey = %v", expectedPubkey, puk)
	assert.Equalf(t, expectedPrivkey, pik, "expectedPrivateKey = %v, privKey = %v", expectedPrivkey, pik)
}

func (LocalSecp256k1Test) TestGetSigner(t *testing.T, p provider.ICryptoProvider) {

}
