package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRevokeIdentity = "revoke_identity"

var _ sdk.Msg = &MsgRevokeIdentity{}

func NewMsgRevokeIdentity(creator string, id uint64, member string) *MsgRevokeIdentity {
	return &MsgRevokeIdentity{
		Creator: creator,
		Id:      id,
		Member:  member,
	}
}

func (msg *MsgRevokeIdentity) Route() string {
	return RouterKey
}

func (msg *MsgRevokeIdentity) Type() string {
	return TypeMsgRevokeIdentity
}

func (msg *MsgRevokeIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRevokeIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRevokeIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Member)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid member address (%s)", err)
	}
	return nil
}
