package cli

import (
	"crypto/rand"
	"cryptoV2/keyring"
	"cryptoV2/provider/localSecp256k1"
	"cryptoV2/storage"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

var rootCmd = &cobra.Command{
	Use:   "keyring",
	Short: "Keyring CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use the create or list commands.")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all items",
	Run: func(cmd *cobra.Command, args []string) {
		ids, err := keyring.GetInstance().List()
		if err != nil {
			panic(err)
		}
		for _, v := range ids {
			fmt.Println(v)
		}
	},
}

var listStorageProvidersCmd = &cobra.Command{
	Use:   "list-storage-providers",
	Short: "List all storage providers",
	Run: func(cmd *cobra.Command, args []string) {
		err := keyring.GetInstance().ListStorageProviders()
		if err != nil {
			panic(err)
		}
	},
}

var listCryptoProvidersCmd = &cobra.Command{
	Use:   "list-crypto-providers",
	Short: "List all crypto providers",
	Run: func(cmd *cobra.Command, args []string) {
		err := keyring.GetInstance().ListCryptoProviderBuilders()
		if err != nil {
			panic(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(listStorageProvidersCmd)
	rootCmd.AddCommand(listCryptoProvidersCmd)
}

func CreateDummySecureItem(name string) *storage.SecureItem {
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

	si := storage.NewSecureItem(storage.ItemId{
		Type: localSecp256k1.Secp256k1,
		UUID: name,
		Slot: "0",
	}, procMar)

	return si
}
