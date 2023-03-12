package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgVoidCda = "witness_approve_cda"

var _ sdk.Msg = &MsgFinalizeCda{}

func NewMsgVoidCda(creator string, cdaId uint64) *MsgVoidCda {
	return &MsgVoidCda{
		Creator: creator,
		CdaId:   cdaId,
	}
}

func (msg *MsgVoidCda) Route() string {
	return RouterKey
}

func (msg *MsgVoidCda) Type() string {
	return TypeMsgVoidCda
}

func (msg *MsgVoidCda) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVoidCda) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVoidCda) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
