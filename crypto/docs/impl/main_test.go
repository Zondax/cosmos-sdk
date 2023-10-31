package main

import (
	"cryptoV2/cli"
	"cryptoV2/keyring"
	"cryptoV2/provider/localSecp256k1"
	storage "cryptoV2/storage/filesystem"
	"fmt"
	"os"
	"testing"
)

func IntegrationTest(t *testing.T) {
	k := keyring.GetInstance()

	// Register CryptoProviderBuilder(s)
	k.RegisterCryptoProviderBuilder(localSecp256k1.Secp256k1Builder)

	// Register StorageProvider(s)
	lfs := storage.NewFileSystemStorageProvider("testing", os.TempDir()+"keyring-test")
	k.RegisterStorageProvider(lfs)

	cli.Execute()

	// Let's create a new SecureItem
	item1 := cli.CreateDummyCryptoProvider("myLocalSecpKey1")
	item2 := cli.CreateDummyCryptoProvider("myLocalSecpKey2")

	err := lfs.Set(item1)
	if err != nil {
		t.Fatalf("Set() error = %v", err)
	}

	err = lfs.Set(item2)
	if err != nil {
		t.Fatalf("Set() error = %v", err)
	}

	fmt.Println("List of items in FileSystemProvider:")
	list := lfs.List()
	for _, v := range list {
		fmt.Println(v)
	}

	fmt.Println("List of items in Keyring:")
	ids, err := k.List()
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	for _, v := range ids {
		fmt.Println(v)
	}

	myKey, err := k.GetCryptoProvider("myLocalSecpKey1_0")
	if err != nil {
		t.Fatalf("GetCryptoProvider() error = %v", err)
	}

	fmt.Println(myKey.GetKeys())

	signer, err := myKey.GetSigner()
	if err != nil {
		t.Fatalf("GetSigner() error = %v", err)
	}

	signature, err := signer.Sign([]byte("Hello"))
	if err != nil {
		t.Fatalf("Sign() error = %v", err)
	}

	verifier, err := myKey.GetVerifier()
	if err != nil {
		t.Fatalf("GetVerifier() error = %v", err)
	}

	ok, err := verifier.Verify([]byte("Hello"), signature)
	if err != nil {
		t.Fatalf("Verify() error = %v", err)
	}

	fmt.Println(ok)
}
