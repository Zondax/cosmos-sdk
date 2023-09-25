package main

import (
	"crypto/rand"
	"cryptoImpl/crypto/provider/localSecp256k1"
	"cryptoImpl/keyring"
	"cryptoImpl/storage"
	"fmt"
	"google.golang.org/protobuf/proto"
	"io"
	"os"
)

func main() {
	k := keyring.New()
	k.RegisterCryptoProvider(localSecp256k1.Secp256k1Builder)
	lfs := storage.NewLFSystem("testing", os.TempDir()+"mainADR")
	k.RegisterStorageProvider(lfs)

	privKeyBytes := [32]byte{}
	r := rand.Reader
	_, err := io.ReadFull(r, privKeyBytes[:])
	if err != nil {
		panic(err)
	}

	proc, err := localSecp256k1.Secp256k1Builder.FromSeed(privKeyBytes[:])
	if err != nil {
		panic(err)
	}
	procMar, err := proto.Marshal(proc)
	if err != nil {
		panic(err)
	}

	si := storage.NewSecureItem(proc.GetTypeUUID(), "myKey", procMar)
	err = lfs.Set(si)
	if err != nil {
		panic(err)
	}

	myKey, err := k.GetCryptoProvider("myKey")
	if err != nil {
		panic(err)
	}

	signer, err := myKey.GetSigner()
	if err != nil {
		panic(err)
	}

	signature, err := signer.Sign([]byte("Hello"))
	if err != nil {
		panic(err)
	}

	verifier, err := myKey.GetVerifier()
	if err != nil {
		panic(err)
	}

	ok, err := verifier.Verify([]byte("Hello"), signature)
	if err != nil {
		panic(err)
	}
	fmt.Println(ok)
}
