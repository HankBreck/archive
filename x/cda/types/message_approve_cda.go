package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveCda = "approve_cda"

var _ sdk.Msg = &MsgApproveCda{}

func NewMsgApproveCda(creator string, cdaId uint64, ownership []*Ownership) *MsgApproveCda {
	return &MsgApproveCda{
		Creator:   creator,
		CdaId:     cdaId,
		Ownership: ownership,
	}
}

func (msg *MsgApproveCda) Route() string {
	return RouterKey
}

func (msg *MsgApproveCda) Type() string {
	return TypeMsgApproveCda
}

func (msg *MsgApproveCda) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgApproveCda) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgApproveCda) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	// Ensure Ownership contains objects
	if len(msg.Ownership) < 1 {
		return sdkerrors.Wrapf(ErrInvalidOwnership, "Invalid ownership length")
	}
	// Ensure Ownership addresses are valid
	for _, owner := range msg.Ownership {
		_, err := sdk.AccAddressFromBech32(owner.Owner)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
		}
	}

	return nil
}
