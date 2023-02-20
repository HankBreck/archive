package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateMembers = "update_members"

var _ sdk.Msg = &MsgUpdateMembers{}

func NewMsgUpdateMembers(creator string, id uint64, toAdd []string, toRemove []string) *MsgUpdateMembers {
	return &MsgUpdateMembers{
		Creator:  creator,
		Id:       id,
		ToAdd:    toAdd,
		ToRemove: toRemove,
	}
}

func (msg *MsgUpdateMembers) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMembers) Type() string {
	return TypeMsgUpdateMembers
}

func (msg *MsgUpdateMembers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMembers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMembers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	for i, addrStr := range msg.ToAdd {
		_, err := sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid ToAdd address at index %d (%s)", i, err)
		}
	}
	for i, addrStr := range msg.ToRemove {
		_, err := sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid ToRemove address at index %d (%s)", i, err)
		}
	}
	return nil
}
