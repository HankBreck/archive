package types

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateCda{}, "cda/CreateCda", nil)
	cdc.RegisterConcrete(&MsgApproveCda{}, "cda/ApproveCda", nil)
	cdc.RegisterConcrete(&MsgWitnessApproveCda{}, "cda/WitnessApproveCda", nil)
	cdc.RegisterConcrete(&MsgFinalizeCda{}, "cda/FinalizeCda", nil)
	cdc.RegisterConcrete(&MsgVoidCda{}, "cda/VoidCda", nil)
	cdc.RegisterConcrete(&SigningDataExtension{}, "cda/SigningDataExtension", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCda{},
		&MsgApproveCda{},
		&MsgWitnessApproveCda{},
		&MsgFinalizeCda{},
		&MsgVoidCda{},
		&MsgRegisterContract{},
	)
	registry.RegisterImplementations((*wasmtypes.ContractInfoExtension)(nil), &SigningDataExtension{})
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
