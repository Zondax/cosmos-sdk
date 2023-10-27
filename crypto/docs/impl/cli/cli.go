package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "keyring",
	Short: "Keyring CLI",
	Long:  `Keyring CLI allows you to create and list items.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use the create or list commands.")
	},
}

var createCmd = &cobra.Command{
	Use:   "create [name]",
	Short: "Create a new item",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		item := createDummySecureItem(name)
		err := lfs.Set(item)
		if err != nil {
			panic(err)
		}
		fmt.Println("Item created successfully.")
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all items",
	Run: func(cmd *cobra.Command, args []string) {
		ids, err := k.List()
		if err != nil {
			panic(err)
		}
		for _, v := range ids {
			fmt.Println(v)
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
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(listCmd)
}
