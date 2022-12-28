package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRejectIdentity = "reject_identity"

var _ sdk.Msg = &MsgRejectIdentity{}

func NewMsgRejectIdentity(creator string, id uint64) *MsgRejectIdentity {
	return &MsgRejectIdentity{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgRejectIdentity) Route() string {
	return RouterKey
}

func (msg *MsgRejectIdentity) Type() string {
	return TypeMsgRejectIdentity
}

func (msg *MsgRejectIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRejectIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRejectIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
