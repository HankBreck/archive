package types

import (
	time "time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateCda = "create_cda"

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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid creator address (%s)", err)
	}

	return nil
}
