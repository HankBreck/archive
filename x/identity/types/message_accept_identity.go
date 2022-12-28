package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAcceptIdentity = "accept_identity"

var _ sdk.Msg = &MsgAcceptIdentity{}

func NewMsgAcceptIdentity(creator string, id uint64) *MsgAcceptIdentity {
	return &MsgAcceptIdentity{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgAcceptIdentity) Route() string {
	return RouterKey
}

func (msg *MsgAcceptIdentity) Type() string {
	return TypeMsgAcceptIdentity
}

func (msg *MsgAcceptIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAcceptIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAcceptIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
