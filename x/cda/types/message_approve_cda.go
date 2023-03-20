package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgApproveCda = "approve_cda"

var _ sdk.Msg = &MsgApproveCda{}

func NewMsgApproveCda(creator string, cdaId uint64, signerId uint64, signingData RawSigningData) *MsgApproveCda {
	return &MsgApproveCda{
		Creator:     creator,
		CdaId:       cdaId,
		SignerId:    signerId,
		SigningData: signingData,
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

	return nil
}
