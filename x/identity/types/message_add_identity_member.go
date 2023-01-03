package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddIdentityMember = "add_identity_member"

var _ sdk.Msg = &MsgAddIdentityMember{}

func NewMsgAddIdentityMember(creator string, id uint64, member string) *MsgAddIdentityMember {
	return &MsgAddIdentityMember{
		Creator: creator,
		Id:      id,
		Member:  member,
	}
}

func (msg *MsgAddIdentityMember) Route() string {
	return RouterKey
}

func (msg *MsgAddIdentityMember) Type() string {
	return TypeMsgAddIdentityMember
}

func (msg *MsgAddIdentityMember) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddIdentityMember) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddIdentityMember) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
