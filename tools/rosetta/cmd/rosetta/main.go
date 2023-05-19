package main

import (
	"fmt"
	"os"

	"cosmossdk.io/log"
	rosettaCmd "cosmossdk.io/tools/rosetta/cmd"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	ibcLightClient "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
	//ibcLightClientTypes "github.com/cosmos/ibc-go/v7/modules/core/exported"
	//sdk "github.com/cosmos/cosmos-sdk/types"
)

func main() {
	var (
		logger            = log.NewLogger(os.Stdout).With(log.ModuleKey, "rosetta")
		interfaceRegistry = codectypes.NewInterfaceRegistry()
		cdc               = codec.NewProtoCodec(interfaceRegistry)
	)

	ibcclienttypes.RegisterInterfaces(interfaceRegistry)
	ibcLightClient.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterInterface("/ibc.lightclients.tendermint.v1.Header", (*exported.ClientMessage)(nil))
	fmt.Println(interfaceRegistry.ListAllInterfaces())

	if err := rosettaCmd.RosettaCommand(interfaceRegistry, cdc).Execute(); err != nil {
		logger.Error("failed to run rosetta", "error", err)
		os.Exit(1)
	}
}
