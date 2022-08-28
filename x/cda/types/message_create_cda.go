package types

import (
	"time"

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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid creator address (%s)", err)
	}

	// kind of a dumb check, but an IPFS cid should not be less than 30 chars long
	// v0 lengths seem to be consistently 46 chars in length
	// v1 is longer
	if len(msg.Cid) < 40 {
		return sdkerrors.Wrapf(ErrInvalidCid, "Invalid cid length")
	}

	// Ensure Ownership contains objects
	if len(msg.Ownership) < 1 {
		return sdkerrors.Wrapf(ErrInvalidOwnership, "Invalid ownership length")
	}
	// Ensure Ownership addresses are valid
	for owner := range msg.Ownership {
		_, err := sdk.AccAddressFromBech32(owner)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid owner address (%s)", err)
		}
	}

	// Do not allow expirations from the past
	if msg.Expiration < uint64(time.Now().UnixMilli()) {
		return sdkerrors.Wrapf(ErrInvalidExpiration, "Invalid expiration timestamp. Make sure the time is represented in milliseconds")
	}
	return nil
}
