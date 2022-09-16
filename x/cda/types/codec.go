package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateCDA{}, "cda/CreateCDA", nil)
	cdc.RegisterConcrete(&MsgApproveCda{}, "cda/ApproveCda", nil)
	cdc.RegisterConcrete(&MsgFinalizeCda{}, "cda/FinalizeCda", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCDA{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgApproveCda{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFinalizeCda{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
