package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFinalizeCda = "finalize_cda"

var _ sdk.Msg = &MsgFinalizeCda{}

func NewMsgFinalizeCda(creator string, cdaId uint64) *MsgFinalizeCda {
	return &MsgFinalizeCda{
		Creator: creator,
		CdaId:   cdaId,
	}
}

func (msg *MsgFinalizeCda) Route() string {
	return RouterKey
}

func (msg *MsgFinalizeCda) Type() string {
	return TypeMsgFinalizeCda
}

func (msg *MsgFinalizeCda) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFinalizeCda) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFinalizeCda) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
