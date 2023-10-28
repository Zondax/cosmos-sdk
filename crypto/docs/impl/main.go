package main

import (
	"cryptoV2/cli"
	"cryptoV2/keyring"
	"cryptoV2/provider/localSecp256k1"
	storage "cryptoV2/storage/filesystem"
	"os"
)

func main() {
	// TODO: keyring should a global var or singleton
	k := keyring.GetInstance()

	// Register CryptoProviderBuilder(s)
	k.RegisterCryptoProviderBuilder(localSecp256k1.Secp256k1Builder)

	// Register StorageProvider(s)
	lfs := storage.NewFileSystemStorageProvider("testing", os.TempDir()+"keyring")
	k.RegisterStorageProvider(lfs)

	cli.Execute()

	// // Let's create a new SecureItem
	// item1 := createDummySecureItem("myLocalSecpKey1")
	// item2 := createDummySecureItem("myLocalSecpKey2")

	// err := lfs.Set(item1)
	// if err != nil {
	// 	panic(err)
	// }

	// err = lfs.Set(item2)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("List of items in FileSystemProvider:")
	// list := lfs.List()
	// for _, v := range list {
	// 	fmt.Println(v)
	// }

	// fmt.Println("List of items in Keyring:")
	// ids, err := k.List()
	// if err != nil {
	// 	panic(err)
	// }
	// for _, v := range ids {
	// 	fmt.Println(v)
	// }

	// myKey, err := k.GetCryptoProvider("myLocalSecpKey1_0")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(myKey.GetKeys())

	// signer, err := myKey.GetSigner()
	// if err != nil {
	// 	panic(err)
	// }

	// signature, err := signer.Sign([]byte("Hello"))
	// if err != nil {
	// 	panic(err)
	// }

	// verifier, err := myKey.GetVerifier()
	// if err != nil {
	// 	panic(err)
	// }

	// ok, err := verifier.Verify([]byte("Hello"), signature)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(ok)
}
