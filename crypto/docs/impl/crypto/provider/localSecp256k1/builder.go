package localSecp256k1

import (
	"bytes"
	"cryptoImpl/crypto/provider"
	"cryptoImpl/storage"
	"errors"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"google.golang.org/protobuf/proto"
)

var Secp256k1Builder = Builder{}

type Builder struct{}

func (b Builder) GetTypeUUID() string {
	return uuid
}

func (b Builder) GetName() string {
	//TODO why this here?
	return ""
}

func (b Builder) FromSecureItem(item storage.ISecureItem) (provider.ICryptoProvider, error) {
	var prov LocalSecp256K1
	err := proto.Unmarshal(item.Bytes(), &prov)
	if err != nil {
		return nil, err
	}
	return &prov, nil
}

func (b Builder) FromSeed(seed []byte) (provider.ICryptoProvider, error) {
	sk, err := secp256k1.GeneratePrivateKeyFromRand(bytes.NewReader(seed))
	if err != nil {
		return nil, err
	}
	pk := sk.PubKey()
	prov := &LocalSecp256K1{
		PubKey:  pk.SerializeUncompressed(),
		PrivKey: sk.Serialize(),
	}
	return prov, nil
}

func (b Builder) FromMnemonic(mnemonic string) (provider.ICryptoProvider, error) {
	return nil, errors.New("not implemented")
}

func (b Builder) FromString(url string) (provider.ICryptoProvider, error) {
	return nil, errors.New("not implemented")
}
