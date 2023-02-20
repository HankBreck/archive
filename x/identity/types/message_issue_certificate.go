package types

import (
	"github.com/HankBreck/archive/x/contractregistry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgIssueCertificate = "issue_certificate"

var _ sdk.Msg = &MsgIssueCertificate{}

func NewMsgIssueCertificate(creator string, recipient string, salt string, metadataSchemaUri string, hashes []*HashEntry) *MsgIssueCertificate {
	return &MsgIssueCertificate{
		Creator:           creator,
		Recipient:         recipient,
		Salt:              salt,
		MetadataSchemaUri: metadataSchemaUri,
		Hashes:            hashes,
	}
}

func (msg *MsgIssueCertificate) Route() string {
	return RouterKey
}

func (msg *MsgIssueCertificate) Type() string {
	return TypeMsgIssueCertificate
}

func (msg *MsgIssueCertificate) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgIssueCertificate) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgIssueCertificate) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
	}
	for i, entry := range msg.Hashes {
		if entry == nil {
			return sdkerrors.Wrapf(types.ErrEmpty, "hash entry cannot be empty (index of %d)", i)
		}
		if entry.Field == "" {
			return sdkerrors.Wrapf(types.ErrEmpty, "hash entry field cannot be empty (index of %d)", i)
		}
		if entry.Hash == "" {
			return sdkerrors.Wrapf(types.ErrEmpty, "hash entry hash cannot be empty (index of %d)", i)
		}
	}
	return nil
}
