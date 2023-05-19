package rosetta

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	authcodec "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankcodec "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/cosmos/cosmos-sdk/x/auth/tx"

	ibcclienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	"github.com/cosmos/ibc-go/v7/modules/core/exported"
	ibcLightClient "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"

	//ibcLightClientTypes "github.com/cosmos/ibc-go/v7/modules/core/exported"
	//sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/std"
	testCodec "github.com/cosmos/cosmos-sdk/types/module/testutil"
)

// MakeCodec generates the codec required to interact
// with the cosmos APIs used by the rosetta gateway
func MakeCodec() (*codec.ProtoCodec, codectypes.InterfaceRegistry) {
	ir := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(ir)

	authcodec.RegisterInterfaces(ir)
	bankcodec.RegisterInterfaces(ir)
	cryptocodec.RegisterInterfaces(ir)

	return cdc, ir
}

// RegisterInterfaces registers rosetta related implementations and interfaces.
func RegisterInterfaces(cdc *codec.ProtoCodec, registry codectypes.InterfaceRegistry) {
	ibcclienttypes.RegisterInterfaces(registry)
	ibcLightClient.RegisterInterfaces(registry)
	registry.RegisterInterface("/ibc.lightclients.tendermint.v1.Header", (*exported.ClientMessage)(nil))
	aminoCodec := codec.NewLegacyAmino()
	encCfg := testCodec.TestEncodingConfig{
		InterfaceRegistry: registry,
		Codec:             cdc,
		TxConfig:          tx.NewTxConfig(cdc, tx.DefaultSignModes),
		Amino:             aminoCodec,
	}
	std.RegisterLegacyAminoCodec(encCfg.Amino)
	std.RegisterInterfaces(encCfg.InterfaceRegistry)
}
