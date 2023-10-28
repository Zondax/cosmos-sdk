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
}
