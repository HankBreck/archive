package types

import (
	"github.com/HankBreck/archive/x/contractregistry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterIssuer = "register_issuer"

var _ sdk.Msg = &MsgRegisterIssuer{}

func NewMsgRegisterIssuer(creator string, name string, moreInfoUri string) *MsgRegisterIssuer {
	return &MsgRegisterIssuer{
		Creator:     creator,
		Name:        name,
		MoreInfoUri: moreInfoUri,
	}
}

func (msg *MsgRegisterIssuer) Route() string {
	return RouterKey
}

func (msg *MsgRegisterIssuer) Type() string {
	return TypeMsgRegisterIssuer
}

func (msg *MsgRegisterIssuer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterIssuer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterIssuer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Name == "" {
		return sdkerrors.Wrapf(types.ErrEmpty, "name cannot be empty")
	}

	return nil
}
