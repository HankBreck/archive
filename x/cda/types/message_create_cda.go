package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateCDA = "create_cda"

var _ sdk.Msg = &MsgCreateCDA{}

func NewMsgCreateCDA(creator string, cid string) *MsgCreateCDA {
	return &MsgCreateCDA{
		Creator: creator,
		Cid:     cid,
	}
}

func (msg *MsgCreateCDA) Route() string {
	return RouterKey
}

func (msg *MsgCreateCDA) Type() string {
	return TypeMsgCreateCDA
}

func (msg *MsgCreateCDA) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCDA) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCDA) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.Cid) == 0 {
		return sdkerrors.Wrapf(ErrInvalidCid, "invalid cid length (0)")
	}
	return nil
}
