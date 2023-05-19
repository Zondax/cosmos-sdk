package main

import (
	"os"

	"cosmossdk.io/log"
	rosettaCmd "cosmossdk.io/tools/rosetta/cmd"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"

	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibcLightClient "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"

	//ibcLightClientTypes "github.com/cosmos/ibc-go/v7/modules/core/exported"
	//sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/std"
	testCodec "github.com/cosmos/cosmos-sdk/types/module/testutil"
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
	aminoCodec := codec.NewLegacyAmino()
	encCfg := testCodec.TestEncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Codec:             cdc,
		TxConfig:          tx.NewTxConfig(cdc, tx.DefaultSignModes),
		Amino:             aminoCodec,
	}
	std.RegisterLegacyAminoCodec(encCfg.Amino)
	std.RegisterInterfaces(encCfg.InterfaceRegistry)

	if err := rosettaCmd.RosettaCommand(interfaceRegistry, cdc).Execute(); err != nil {
		logger.Error("failed to run rosetta", "error", err)
		os.Exit(1)
	}
}
