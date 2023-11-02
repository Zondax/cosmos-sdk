package cli

import (
	"crypto/rand"
	"cryptoV2/keyring"
	"cryptoV2/provider"
	"cryptoV2/provider/localSecp256k1"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
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

func CreateDummyCryptoProvider(name string) provider.ICryptoProvider {
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

	return proc
}
