package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgApproveCda        = "approve_cda"
	TypeMsgCreateCda         = "create_cda"
	TypeMsgFinalizeCda       = "finalize_cda"
	TypeMsgRegisterContract  = "register_contract"
	TypeMsgVoidCda           = "void_cda"
	TypeMsgWitnessApproveCda = "witness_approve_cda"
)

// MsgApproveCda
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
	if msg.CdaId < 1 {
		return sdkerrors.Wrapf(ErrInvalid, "CDA ID must be greater than 0")
	}
	if msg.SignerId < 1 {
		return sdkerrors.Wrapf(ErrInvalid, "signer ID must be greater than 0")
	}
	err = msg.SigningData.ValidateBasic()
	if err != nil {
		return err
	}

	return nil
}

// MsgCreateCda
var _ sdk.Msg = &MsgCreateCda{}

func NewMsgCreateCda(creator string, signerIds []uint64, contractId uint64, legalMetadataUri string, signingData RawSigningData, utcExpireTime time.Time, witnessInitMsg RawSigningData) *MsgCreateCda {
	return &MsgCreateCda{
		Creator:          creator,
		SignerIds:        signerIds,
		ContractId:       contractId,
		LegalMetadataUri: legalMetadataUri,
		SigningData:      signingData,
		UtcExpireTime:    utcExpireTime,
		WitnessInitMsg:   witnessInitMsg,
	}
}

func (msg *MsgCreateCda) Route() string {
	return RouterKey
}

func (msg *MsgCreateCda) Type() string {
	return TypeMsgCreateCda
}

func (msg *MsgCreateCda) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCda) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCda) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if len(msg.SignerIds) < 1 {
		return sdkerrors.Wrapf(ErrEmpty, "signerIds cannot be empty")
	}
	for _, signerId := range msg.SignerIds {
		if signerId < 1 {
			return sdkerrors.Wrapf(ErrInvalid, "signer ID must be greater than 0")
		}
	}
	if msg.ContractId < 1 {
		return sdkerrors.Wrapf(ErrInvalid, "contract ID must be greater than 0")
	}
	if len(msg.LegalMetadataUri) < 1 {
		return sdkerrors.Wrapf(ErrEmpty, "legalMetadataUri cannot be empty")
	}
	err = msg.SigningData.ValidateBasic()
	if err != nil {
		return sdkerrors.Wrapf(ErrInvalid, "signing data must be valid json")
	}
	// TODO: should we require signing data to exist?

	return nil
}

// MsgFinalizeCda
var _ sdk.Msg = &MsgFinalizeCda{}

func NewMsgFinalizeCda(creator string, cdaId uint64) *MsgFinalizeCda {
	return &MsgFinalizeCda{
		Creator: creator,
		CdaId:   cdaId,
	}
}

func (msg *MsgFinalizeCda) Route() string {
	return RouterKey
}

func (msg *MsgFinalizeCda) Type() string {
	return TypeMsgFinalizeCda
}

func (msg *MsgFinalizeCda) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFinalizeCda) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFinalizeCda) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.CdaId < 1 {
		return sdkerrors.Wrapf(ErrInvalid, "CDA ID must be greater than 0")
	}
	return nil
}

// MsgRegisterContract
var _ sdk.Msg = &MsgRegisterContract{}

func NewMsgRegisterContract(creator string, description string, authors []string, contactInfo *ContactInfo, moreInfoUri string, signingDataSchema RawSigningData, templateUri string, templateSchemaUri string, witnessCodeId uint64) *MsgRegisterContract {
	return &MsgRegisterContract{
		Creator:           creator,
		Description:       description,
		Authors:           authors,
		ContactInfo:       contactInfo,
		MoreInfoUri:       moreInfoUri,
		SigningDataSchema: signingDataSchema,
		TemplateUri:       templateUri,
		TemplateSchemaUri: templateSchemaUri,
		WitnessCodeId:     witnessCodeId,
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	err = msg.SigningDataSchema.ValidateBasic()
	if err != nil {
		return ErrInvalid.Wrapf("signing data schema must be valid JSON")
	}

	return nil
}

// MsgVoidCda
var _ sdk.Msg = &MsgVoidCda{}

func NewMsgVoidCda(creator string, cdaId uint64) *MsgVoidCda {
	return &MsgVoidCda{
		Creator: creator,
		CdaId:   cdaId,
	}
}

func (msg *MsgVoidCda) Route() string {
	return RouterKey
}

func (msg *MsgVoidCda) Type() string {
	return TypeMsgVoidCda
}

func (msg *MsgVoidCda) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVoidCda) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVoidCda) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.CdaId < 1 {
		return sdkerrors.Wrapf(ErrInvalid, "CDA ID must be greater than 0")
	}
	return nil
}

// MsgWitnessApproveCda
var _ sdk.Msg = &MsgWitnessApproveCda{}

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
	// TODO: should we enforce address length here?
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.CdaId < 1 {
		return sdkerrors.Wrapf(ErrInvalid, "CDA ID must be greater than 0")
	}
	err = msg.SigningData.ValidateBasic()
	if err != nil {
		return ErrInvalid.Wrapf("signing data must be valid JSON")
	}
	return nil
}
