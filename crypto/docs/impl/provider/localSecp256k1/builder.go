package localSecp256k1

import (
	"bytes"
	"cryptoV2/provider"
	"cryptoV2/storage"
	"errors"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"google.golang.org/protobuf/proto"
)

var Secp256k1Builder = Builder{Secp256k1}

type Builder struct {
	uuid string
}

func (b Builder) GetBuilderTypeUUID() string {
	return b.uuid
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
