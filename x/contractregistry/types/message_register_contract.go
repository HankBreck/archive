package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegisterContract = "register_contract"

var _ sdk.Msg = &MsgRegisterContract{}

func NewMsgRegisterContract(creator string, description string, authors []string, contactInfo *ContactInfo, moreInfoUri string, signingDataSchema RawSigningData, templateUri string, templateSchemaUri string) *MsgRegisterContract {
	return &MsgRegisterContract{
		Creator:           creator,
		Description:       description,
		Authors:           authors,
		ContactInfo:       contactInfo,
		MoreInfoUri:       moreInfoUri,
		SigningDataSchema: signingDataSchema,
		TemplateUri:       templateUri,
		TemplateSchemaUri: templateSchemaUri,
	}
}

func (msg *MsgRegisterContract) Route() string {
	return RouterKey
}

func (msg *MsgRegisterContract) Type() string {
	return TypeMsgRegisterContract
}

func (msg *MsgRegisterContract) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterContract) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterContract) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid creator address (%s)", err)
	}

	// Validate that Schema is valid JSON
	// TODO: Use the JSON Schema library here
	err = msg.SigningDataSchema.ValidateBasic()
	if err != nil {
		return ErrInvalid.Wrapf("Signing data schema must be valid JSON")
	}
	if msg.SigningDataSchema.Bytes() == nil {
		return ErrEmpty.Wrapf("Signing data schema cannot be null")
	}

	// Should we allow no contact info?
	if msg.ContactInfo == nil {
		return ErrEmpty.Wrapf("Contact info cannot be null")
	}

	return nil
}
