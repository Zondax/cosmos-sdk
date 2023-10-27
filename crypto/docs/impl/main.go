package main

import (
	"crypto/rand"
	"cryptoImpl/keyring"
	"cryptoImpl/provider/localSecp256k1"
	storage2 "cryptoImpl/storage"
	storage "cryptoImpl/storage/filesystem"
	"fmt"
	"io"
	"os"

	"google.golang.org/protobuf/proto"
)

func main() {
	// TODO: keyring should a global var or singleton
	k := keyring.New()

	// Register CryptoProviderBuilder(s)
	k.RegisterCryptoProviderBuilder(localSecp256k1.Secp256k1Builder)

	// Register StorageProvider(s)
	lfs := storage.NewFileSystemStorageProvider("testing", os.TempDir()+"keyring")
	k.RegisterStorageProvider(lfs)

	// Let's create a new SecureItem
	item1 := createDummySecureItem("myLocalSecpKey1")
	item2 := createDummySecureItem("myLocalSecpKey2")

	err := lfs.Set(item1)
	if err != nil {
		panic(err)
	}

	err = lfs.Set(item2)
	if err != nil {
		panic(err)
	}

	fmt.Println("List of items in FileSystemProvider:")
	list := lfs.List()
	for _, v := range list {
		fmt.Println(v)
	}

	fmt.Println("List of items in Keyring:")
	ids, err := k.List()
	if err != nil {
		panic(err)
	}
	for _, v := range ids {
		fmt.Println(v)
	}

	myKey, err := k.GetCryptoProvider("myLocalSecpKey1_0")
	if err != nil {
		panic(err)
	}

	fmt.Println(myKey.GetKeys())

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

func createDummySecureItem(name string) *storage2.SecureItem {
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

	si := storage2.NewSecureItem(storage2.ItemId{
		Type: localSecp256k1.Secp256k1,
		UUID: name,
		Slot: "0",
	}, procMar)

	return si
}
