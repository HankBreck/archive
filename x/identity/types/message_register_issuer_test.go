package types

import (
	"testing"

	"github.com/HankBreck/archive/testutil/sample"
	"github.com/HankBreck/archive/x/contractregistry/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgRegisterIssuer_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRegisterIssuer
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgRegisterIssuer{
				Creator:     "invalid_address",
				Name:        "test",
				MoreInfoUri: "google.com",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "empty name",
			msg: MsgRegisterIssuer{
				Creator:     sample.AccAddress(),
				Name:        "",
				MoreInfoUri: "google.com",
			},
			err: types.ErrEmpty,
		}, {
			name: "empty more info URI",
			msg: MsgRegisterIssuer{
				Creator:     sample.AccAddress(),
				Name:        "test",
				MoreInfoUri: "",
			},
		}, {
			name: "valid address",
			msg: MsgRegisterIssuer{
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
