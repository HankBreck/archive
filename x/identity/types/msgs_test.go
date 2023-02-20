package types_test

import (
	"testing"

	"github.com/HankBreck/archive/testutil/sample"
	"github.com/HankBreck/archive/x/identity/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgAcceptIdentity_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgAcceptIdentity
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgAcceptIdentity{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgAcceptIdentity{
				Creator: sample.AccAddress(),
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

func TestMsgFreezeIdentity_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgFreezeIdentity
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgFreezeIdentity{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgFreezeIdentity{
				Creator: sample.AccAddress(),
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

func TestMsgIssueCertificate_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgIssueCertificate
		err  error
	}{
		{
			name: "invalid creator address",
			msg: types.MsgIssueCertificate{
				Creator:           "invalid_address",
				Recipient:         sample.AccAddress(),
				MetadataSchemaUri: "google.com",
				Hashes:            []*types.HashEntry{{Field: "foo", Hash: "bar"}},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid recipient address",
			msg: types.MsgIssueCertificate{
				Creator:           sample.AccAddress(),
				Recipient:         "invalid_address",
				MetadataSchemaUri: "google.com",
				Hashes:            []*types.HashEntry{{Field: "foo", Hash: "bar"}},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "empty hash entry field",
			msg: types.MsgIssueCertificate{
				Creator:           sample.AccAddress(),
				Recipient:         sample.AccAddress(),
				Salt:              "",
				MetadataSchemaUri: "",
				Hashes:            []*types.HashEntry{{Field: "", Hash: "bar"}},
			},
			err: types.ErrEmpty,
		}, {
			name: "empty hash entry hash",
			msg: types.MsgIssueCertificate{
				Creator:           sample.AccAddress(),
				Recipient:         sample.AccAddress(),
				Salt:              "",
				MetadataSchemaUri: "",
				Hashes:            []*types.HashEntry{{Field: "foo", Hash: ""}},
			},
			err: types.ErrEmpty,
		}, {
			name: "nil hash entry",
			msg: types.MsgIssueCertificate{
				Creator:           sample.AccAddress(),
				Recipient:         sample.AccAddress(),
				Salt:              "",
				MetadataSchemaUri: "",
				Hashes:            []*types.HashEntry{{Field: "foo", Hash: "bar"}, nil},
			},
			err: types.ErrEmpty,
		}, {
			name: "valid address",
			msg: types.MsgIssueCertificate{
				Creator:           sample.AccAddress(),
				Recipient:         sample.AccAddress(),
				MetadataSchemaUri: "google.com",
				Hashes:            []*types.HashEntry{{Field: "foo", Hash: "bar"}},
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

func TestMsgRegisterIssuer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgRegisterIssuer
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgRegisterIssuer{
				Creator:     "invalid_address",
				Name:        "test",
				MoreInfoUri: "google.com",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "empty name",
			msg: types.MsgRegisterIssuer{
				Creator:     sample.AccAddress(),
				Name:        "",
				MoreInfoUri: "google.com",
			},
			err: types.ErrEmpty,
		}, {
			name: "empty more info URI",
			msg: types.MsgRegisterIssuer{
				Creator:     sample.AccAddress(),
				Name:        "test",
				MoreInfoUri: "",
			},
		}, {
			name: "valid address",
			msg: types.MsgRegisterIssuer{
				Creator:     sample.AccAddress(),
				Name:        "test",
				MoreInfoUri: "google.com",
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

func TestMsgRejectIdentity_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgRejectIdentity
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgRejectIdentity{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgRejectIdentity{
				Creator: sample.AccAddress(),
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

func TestMsgRenounceIdentity_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgRenounceIdentity
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgRenounceIdentity{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgRenounceIdentity{
				Creator: sample.AccAddress(),
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

func TestMsgUpdateMembers_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgUpdateMembers
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgUpdateMembers{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid toAdd address",
			msg: types.MsgUpdateMembers{
				Creator: sample.AccAddress(),
				ToAdd:   []string{"invalid_address"},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid toRemove address",
			msg: types.MsgUpdateMembers{
				Creator:  sample.AccAddress(),
				ToRemove: []string{"invalid_address"},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid message",
			msg: types.MsgUpdateMembers{
				Creator:  sample.AccAddress(),
				Id:       10,
				ToAdd:    []string{sample.AccAddress()},
				ToRemove: []string{sample.AccAddress()},
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

func TestMsgUpdateOperators_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgUpdateOperators
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgUpdateOperators{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid toAdd address",
			msg: types.MsgUpdateOperators{
				Creator: sample.AccAddress(),
				ToAdd:   []string{"invalid_address"},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid toRemove address",
			msg: types.MsgUpdateOperators{
				Creator:  sample.AccAddress(),
				ToRemove: []string{"invalid_address"},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid message",
			msg: types.MsgUpdateOperators{
				Creator:  sample.AccAddress(),
				Id:       10,
				ToAdd:    []string{sample.AccAddress()},
				ToRemove: []string{sample.AccAddress()},
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
