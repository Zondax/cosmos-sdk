package localSecp256k1

import (
	"cryptoV2/provider"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
)

const Secp256k1 = "localSecp256k1"

func (x *LocalSecp256K1) Sign(hash []byte) ([]byte, error) {
	sk := secp256k1.PrivKeyFromBytes(x.PrivKey)
	sig := ecdsa.Sign(sk, hash)
	return sig.Serialize(), nil
}

func (x *LocalSecp256K1) Verify(hash []byte, signature []byte) (bool, error) {
	sig, err := ecdsa.ParseDERSignature(signature)
	if err != nil {
		return false, err
	}
	pk, err := secp256k1.ParsePubKey(x.GetPubKey())
	if err != nil {
		return false, err
	}
	return sig.Verify(hash, pk), nil
}

func (x *LocalSecp256K1) GetMetadata() provider.ICryptoProviderMetadata {
	return x
}

func (x *LocalSecp256K1) GetKeys() ([]byte, []byte, error) {
	return x.PrivKey, x.PubKey, nil
}

func (x *LocalSecp256K1) GetSigner() (provider.ISigner, error) {
	return x, nil
}

func (x *LocalSecp256K1) GetVerifier() (provider.IVerifier, error) {
	return x, nil
}

func (x *LocalSecp256K1) GetCipher() (provider.ICipher, error) {
	//TODO implement me
	panic("implement me")
}

func (x *LocalSecp256K1) GetHasher() (provider.IHasher, error) {
	//TODO implement me
	panic("implement me")
}
