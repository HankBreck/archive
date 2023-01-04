package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateOperators = "update_operators"

var _ sdk.Msg = &MsgUpdateOperators{}

func NewMsgUpdateOperators(creator string, id uint64, toAdd []string, toRemove []string) *MsgUpdateOperators {
	return &MsgUpdateOperators{
		Creator:  creator,
		Id:       id,
		ToAdd:    toAdd,
		ToRemove: toRemove,
	}
}

func (msg *MsgUpdateOperators) Route() string {
	return RouterKey
}

func (msg *MsgUpdateOperators) Type() string {
	return TypeMsgUpdateOperators
}

func (msg *MsgUpdateOperators) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateOperators) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateOperators) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
