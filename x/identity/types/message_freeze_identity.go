package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFreezeIdentity = "freeze_identity"

var _ sdk.Msg = &MsgUpdateOperators{}

func NewMsgFreezeIdentity(creator string, id uint64) *MsgFreezeIdentity {
	return &MsgFreezeIdentity{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgFreezeIdentity) Route() string {
	return RouterKey
}

func (msg *MsgFreezeIdentity) Type() string {
	return TypeMsgUpdateOperators
}

func (msg *MsgFreezeIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFreezeIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFreezeIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
