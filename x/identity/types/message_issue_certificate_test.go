package types

import (
	"testing"

	"github.com/HankBreck/archive/testutil/sample"
	"github.com/HankBreck/archive/x/contractregistry/types"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgIssueCertificate_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgIssueCertificate
		err  error
	}{
		{
			name: "invalid creator address",
			msg: MsgIssueCertificate{
				Creator:           "invalid_address",
				Recipient:         sample.AccAddress(),
				MetadataSchemaUri: "google.com",
				Hashes:            []*HashEntry{{Field: "foo", Hash: "bar"}},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid recipient address",
			msg: MsgIssueCertificate{
				Creator:           sample.AccAddress(),
				Recipient:         "invalid_address",
				MetadataSchemaUri: "google.com",
				Hashes:            []*HashEntry{{Field: "foo", Hash: "bar"}},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "empty hash entry field",
			msg: MsgIssueCertificate{
				Creator:           sample.AccAddress(),
				Recipient:         sample.AccAddress(),
				Salt:              "",
				MetadataSchemaUri: "",
				Hashes:            []*HashEntry{{Field: "", Hash: "bar"}},
			},
			err: types.ErrEmpty,
		}, {
			name: "empty hash entry hash",
			msg: MsgIssueCertificate{
				Creator:           sample.AccAddress(),
				Recipient:         sample.AccAddress(),
				Salt:              "",
				MetadataSchemaUri: "",
				Hashes:            []*HashEntry{{Field: "foo", Hash: ""}},
			},
			err: types.ErrEmpty,
		}, {
			name: "nil hash entry",
			msg: MsgIssueCertificate{
				Creator:           sample.AccAddress(),
				Recipient:         sample.AccAddress(),
				Salt:              "",
				MetadataSchemaUri: "",
				Hashes:            []*HashEntry{{Field: "foo", Hash: "bar"}, nil},
			},
			err: types.ErrEmpty,
		}, {
			name: "valid address",
			msg: MsgIssueCertificate{
				Creator:           sample.AccAddress(),
				Recipient:         sample.AccAddress(),
				MetadataSchemaUri: "google.com",
				Hashes:            []*HashEntry{{Field: "foo", Hash: "bar"}},
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
