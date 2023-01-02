package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRenounceIdentity = "renounce_identity"

var _ sdk.Msg = &MsgRenounceIdentity{}

func NewMsgRenounceIdentity(creator string, id uint64) *MsgRenounceIdentity {
	return &MsgRenounceIdentity{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgRenounceIdentity) Route() string {
	return RouterKey
}

func (msg *MsgRenounceIdentity) Type() string {
	return TypeMsgRenounceIdentity
}

func (msg *MsgRenounceIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRenounceIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRenounceIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
