package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegisterIssuer{}, "identity/RegisterIssuer", nil)
	cdc.RegisterConcrete(&MsgAcceptIdentity{}, "identity/AcceptIdentity", nil)
	cdc.RegisterConcrete(&MsgRejectIdentity{}, "identity/RejectIdentity", nil)
	cdc.RegisterConcrete(&MsgRenounceIdentity{}, "identity/RenounceIdentity", nil)
	cdc.RegisterConcrete(&MsgUpdateOperators{}, "identity/UpdateOperators", nil)
	cdc.RegisterConcrete(&MsgUpdateMembers{}, "identity/UpdateMembers", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegisterIssuer{},
		&MsgAcceptIdentity{},
		&MsgRejectIdentity{},
		&MsgRenounceIdentity{},
		&MsgUpdateOperators{},
		&MsgUpdateMembers{},
		&MsgFreezeIdentity{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
