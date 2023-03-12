package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWitnessApproveCda = "witness_approve_cda"

var _ sdk.Msg = &MsgFinalizeCda{}

func NewMsgWitnessApproveCda(creator string, cdaId uint64, signingData RawSigningData) *MsgWitnessApproveCda {
	return &MsgWitnessApproveCda{
		Creator:     creator,
		CdaId:       cdaId,
		SigningData: signingData,
	}
}

func (msg *MsgWitnessApproveCda) Route() string {
	return RouterKey
}

func (msg *MsgWitnessApproveCda) Type() string {
	return TypeMsgWitnessApproveCda
}

func (msg *MsgWitnessApproveCda) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWitnessApproveCda) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWitnessApproveCda) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
