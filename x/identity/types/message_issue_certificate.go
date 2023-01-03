package types

import (
	"archive/x/contractregistry/types"

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
	if msg.MetadataSchemaUri == "" {
		return sdkerrors.Wrapf(types.ErrEmpty, "invalid metadata schema uri")
	}
	if len(msg.Hashes) == 0 {
		return sdkerrors.Wrapf(types.ErrEmpty, "must contain at least one hash")
	}
	return nil
}
