package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgAcceptIdentity   = "accept_identity"
	TypeMsgFreezeIdentity   = "freeze_identity"
	TypeMsgIssueCertificate = "issue_certificate"
	TypeMsgRegisterIssuer   = "register_issuer"
	TypeMsgRejectIdentity   = "reject_identity"
	TypeMsgRenounceIdentity = "renounce_identity"
	TypeMsgUpdateMembers    = "update_members"
	TypeMsgUpdateOperators  = "update_operators"
)

// MsgAcceptIdentity
var _ sdk.Msg = &MsgAcceptIdentity{}

func NewMsgAcceptIdentity(creator string, id uint64) *MsgAcceptIdentity {
	return &MsgAcceptIdentity{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgAcceptIdentity) Route() string {
	return RouterKey
}

func (msg *MsgAcceptIdentity) Type() string {
	return TypeMsgAcceptIdentity
}

func (msg *MsgAcceptIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAcceptIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAcceptIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// MsgFreezeIdentity
var _ sdk.Msg = &MsgFreezeIdentity{}

func NewMsgFreezeIdentity(creator string, id uint64) *MsgFreezeIdentity {
	return &MsgFreezeIdentity{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgFreezeIdentity) Route() string {
	return RouterKey
}

func (msg *MsgFreezeIdentity) Type() string {
	return TypeMsgFreezeIdentity
}

func (msg *MsgFreezeIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFreezeIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFreezeIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// MsgIssueCertificate
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
			return sdkerrors.Wrapf(ErrEmpty, "hash entry cannot be empty (index of %d)", i)
		}
		if entry.Field == "" {
			return sdkerrors.Wrapf(ErrEmpty, "hash entry field cannot be empty (index of %d)", i)
		}
		if entry.Hash == "" {
			return sdkerrors.Wrapf(ErrEmpty, "hash entry hash cannot be empty (index of %d)", i)
		}
	}
	return nil
}

// MsgRegisterIssuer
var _ sdk.Msg = &MsgRegisterIssuer{}

func NewMsgRegisterIssuer(creator string, name string, moreInfoUri string) *MsgRegisterIssuer {
	return &MsgRegisterIssuer{
		Creator:     creator,
		Name:        name,
		MoreInfoUri: moreInfoUri,
	}
}

func (msg *MsgRegisterIssuer) Route() string {
	return RouterKey
}

func (msg *MsgRegisterIssuer) Type() string {
	return TypeMsgRegisterIssuer
}

func (msg *MsgRegisterIssuer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegisterIssuer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegisterIssuer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Name == "" {
		return sdkerrors.Wrapf(ErrEmpty, "name cannot be empty")
	}

	return nil
}

// MsgRejectIdentity
var _ sdk.Msg = &MsgRejectIdentity{}

func NewMsgRejectIdentity(creator string, id uint64) *MsgRejectIdentity {
	return &MsgRejectIdentity{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgRejectIdentity) Route() string {
	return RouterKey
}

func (msg *MsgRejectIdentity) Type() string {
	return TypeMsgRejectIdentity
}

func (msg *MsgRejectIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRejectIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRejectIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// MsgRenounceIdentity
var _ sdk.Msg = &MsgRenounceIdentity{}

func NewMsgRenounceIdentity(creator string, id uint64) *MsgRenounceIdentity {
	return &MsgRenounceIdentity{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgRenounceIdentity) Route() string {
	return RouterKey
}

func (msg *MsgRenounceIdentity) Type() string {
	return TypeMsgRenounceIdentity
}

func (msg *MsgRenounceIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRenounceIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRenounceIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

// MsgUpdateMembers
var _ sdk.Msg = &MsgUpdateMembers{}

func NewMsgUpdateMembers(creator string, id uint64, toAdd []string, toRemove []string) *MsgUpdateMembers {
	return &MsgUpdateMembers{
		Creator:  creator,
		Id:       id,
		ToAdd:    toAdd,
		ToRemove: toRemove,
	}
}

func (msg *MsgUpdateMembers) Route() string {
	return RouterKey
}

func (msg *MsgUpdateMembers) Type() string {
	return TypeMsgUpdateMembers
}

func (msg *MsgUpdateMembers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateMembers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateMembers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	for i, addrStr := range msg.ToAdd {
		_, err := sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid ToAdd address at index %d (%s)", i, err)
		}
	}
	for i, addrStr := range msg.ToRemove {
		_, err := sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid ToRemove address at index %d (%s)", i, err)
		}
	}
	return nil
}

// MsgUpdateOperators
var _ sdk.Msg = &MsgUpdateOperators{}

func NewMsgUpdateOperators(creator string, id uint64, toAdd []string, toRemove []string) *MsgUpdateOperators {
	return &MsgUpdateOperators{
		Creator:  creator,
		Id:       id,
		ToAdd:    toAdd,
		ToRemove: toRemove,
	}
}

func (msg *MsgUpdateOperators) Route() string {
	return RouterKey
}

func (msg *MsgUpdateOperators) Type() string {
	return TypeMsgUpdateOperators
}

func (msg *MsgUpdateOperators) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateOperators) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateOperators) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	for i, addrStr := range msg.ToAdd {
		_, err := sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid ToAdd address at index %d (%s)", i, err)
		}
	}
	for i, addrStr := range msg.ToRemove {
		_, err := sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid ToRemove address at index %d (%s)", i, err)
		}
	}
	return nil
}
