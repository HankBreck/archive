package types_test

import (
	"testing"

	"github.com/HankBreck/archive/testutil/sample"
	"github.com/HankBreck/archive/x/cda/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

// TODO: test passing an empty message as input (should fail)

// MsgApproveCda
func TestMsgApproveCda_ValidateBasic(t *testing.T) {
	defaultSigningData := types.RawSigningData("{}")
	tests := []struct {
		name string
		msg  types.MsgApproveCda
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgApproveCda{
				Creator:     "invalid_address",
				CdaId:       1,
				SignerId:    1,
				SigningData: defaultSigningData,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid cda id",
			msg: types.MsgApproveCda{
				Creator:     sample.AccAddress(),
				CdaId:       0,
				SignerId:    1,
				SigningData: defaultSigningData,
			},
			err: types.ErrInvalid,
		}, {
			name: "invalid signer id",
			msg: types.MsgApproveCda{
				Creator:     sample.AccAddress(),
				CdaId:       1,
				SignerId:    0,
				SigningData: defaultSigningData,
			},
			err: types.ErrInvalid,
		}, {
			name: "invalid signing data",
			msg: types.MsgApproveCda{
				Creator:     sample.AccAddress(),
				CdaId:       1,
				SignerId:    1,
				SigningData: []byte{0x0},
			},
			err: types.ErrInvalid,
		}, {
			name: "valid address, cda id, signer id, and signing data",
			msg: types.MsgApproveCda{
				Creator:     sample.AccAddress(),
				CdaId:       1,
				SignerId:    1,
				SigningData: defaultSigningData,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

// MsgCreateCda
func TestMsgCreateCda_ValidateBasic(t *testing.T) {
	defaultSigningData := types.RawSigningData([]byte("{}"))

	tests := []struct {
		name string
		msg  types.MsgCreateCda
		err  error
	}{
		{
			name: "invalid creator",
			msg: types.MsgCreateCda{
				Creator:          "invalid_address",
				SignerIds:        []uint64{1},
				ContractId:       1,
				LegalMetadataUri: "google.com",
				SigningData:      defaultSigningData,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid signer id length",
			msg: types.MsgCreateCda{
				Creator:          sample.AccAddress(),
				SignerIds:        []uint64{},
				ContractId:       1,
				LegalMetadataUri: "google.com",
				SigningData:      defaultSigningData,
			},
			err: sdkerrors.Wrapf(types.ErrEmpty, "signerIds cannot be empty"),
		}, {
			name: "zero signer id",
			msg: types.MsgCreateCda{
				Creator:          sample.AccAddress(),
				SignerIds:        []uint64{0},
				ContractId:       1,
				LegalMetadataUri: "google.com",
				SigningData:      defaultSigningData,
			},
			err: sdkerrors.Wrapf(types.ErrInvalid, "signer ID must be greater than 0"),
		}, {
			name: "zero contract id",
			msg: types.MsgCreateCda{
				Creator:          sample.AccAddress(),
				SignerIds:        []uint64{1},
				ContractId:       0,
				LegalMetadataUri: "google.com",
				SigningData:      defaultSigningData,
			},
			err: sdkerrors.Wrapf(types.ErrInvalid, "contract ID must be greater than 0"),
		}, {
			name: "empty legal metadata URI",
			msg: types.MsgCreateCda{
				Creator:          sample.AccAddress(),
				SignerIds:        []uint64{1},
				ContractId:       1,
				LegalMetadataUri: "",
				SigningData:      defaultSigningData,
			},
			err: sdkerrors.Wrapf(types.ErrEmpty, "legalMetadataUri cannot be empty"),
		}, {
			name: "invalid signing data",
			msg: types.MsgCreateCda{
				Creator:          sample.AccAddress(),
				SignerIds:        []uint64{1},
				ContractId:       1,
				LegalMetadataUri: "google.com",
				SigningData:      types.RawSigningData{0x0},
			},
			err: sdkerrors.Wrapf(types.ErrInvalid, "signing data must be valid json"),
		}, {
			name: "valid message",
			msg: types.MsgCreateCda{
				Creator:          sample.AccAddress(),
				SignerIds:        []uint64{1},
				ContractId:       1,
				LegalMetadataUri: "google.com",
				SigningData:      defaultSigningData,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

// MsgFinalizeCda
func TestMsgFinalizeCda_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgFinalizeCda
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgFinalizeCda{
				Creator: "invalid_address",
				CdaId:   1,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "zero cda id",
			msg: types.MsgFinalizeCda{
				Creator: sample.AccAddress(),
				CdaId:   0,
			},
			err: sdkerrors.Wrapf(types.ErrInvalid, "CDA ID must be greater than 0"),
		}, {
			name: "valid address and cda id",
			msg: types.MsgFinalizeCda{
				Creator: sample.AccAddress(),
				CdaId:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

// MsgRegisterContract
func TestMsgRegisterContract_ValidateBasic(t *testing.T) {
	defaultSigningDataSchema := types.RawSigningData("{}")
	tests := []struct {
		name string
		msg  types.MsgRegisterContract
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgRegisterContract{
				Creator:           "invalid_address",
				SigningDataSchema: defaultSigningDataSchema,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid signing data schema",
			msg: types.MsgRegisterContract{
				Creator:           sample.AccAddress(),
				SigningDataSchema: types.RawSigningData{0x0},
			},
			err: types.ErrInvalid.Wrapf("signing data schema must be valid JSON"),
		}, {
			name: "valid address and signing data schema",
			msg: types.MsgRegisterContract{
				Creator:           sample.AccAddress(),
				SigningDataSchema: defaultSigningDataSchema,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

// MsgVoidCda
func TestMsgVoidCda_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgVoidCda
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgVoidCda{
				Creator: "invalid_address",
				CdaId:   1,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "zero cda id",
			msg: types.MsgVoidCda{
				Creator: sample.AccAddress(),
				CdaId:   0,
			},
			err: sdkerrors.Wrapf(types.ErrInvalid, "CDA ID must be greater than 0"),
		}, {
			name: "valid address and cda id",
			msg: types.MsgVoidCda{
				Creator: sample.AccAddress(),
				CdaId:   1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

// MsgWitnessApproveCda
func TestMsgWitnessApproveCda_ValidateBasic(t *testing.T) {
	// requires valid JSON
	defaultSigningDataSchema := types.RawSigningData("{}")
	tests := []struct {
		name string
		msg  types.MsgWitnessApproveCda
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgWitnessApproveCda{
				Creator:     "invalid_address",
				CdaId:       1,
				SigningData: defaultSigningDataSchema,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid cda id",
			msg: types.MsgWitnessApproveCda{
				Creator:     sample.AccAddress(),
				CdaId:       0,
				SigningData: defaultSigningDataSchema,
			},
			err: sdkerrors.Wrapf(types.ErrInvalid, "CDA ID must be greater than 0"),
		}, {
			name: "invalid signing data",
			msg: types.MsgWitnessApproveCda{
				Creator:     sample.AccAddress(),
				CdaId:       1,
				SigningData: types.RawSigningData{0x0},
			},
			err: types.ErrInvalid.Wrapf("signing data must be valid JSON"),
		}, {
			name: "valid address, cda id, and signing data",
			msg: types.MsgWitnessApproveCda{
				Creator:     sample.AccAddress(),
				CdaId:       1,
				SigningData: defaultSigningDataSchema,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
