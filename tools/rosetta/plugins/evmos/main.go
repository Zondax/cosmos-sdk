package main

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	evmosTypes "github.com/evmos/ethermint/x/evm/types"
)

func InitZone() {
	config := sdk.GetConfig()

	prefix := "evmos"
	config.SetBech32PrefixForAccount(prefix, prefix+"pub")
	config.SetBech32PrefixForValidator(prefix+"valoper", prefix+"valoperpub")
	config.SetBech32PrefixForConsensusNode(prefix+"valcons", prefix+"valconspub")
}

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	evmosTypes.RegisterInterfaces(registry)
}
