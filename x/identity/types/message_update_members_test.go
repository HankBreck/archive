package types

import (
	"testing"

	"github.com/HankBreck/archive/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateMembers_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateMembers
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateMembers{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid toAdd address",
			msg: MsgUpdateMembers{
				Creator: sample.AccAddress(),
				ToAdd:   []string{"invalid_address"},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid toRemove address",
			msg: MsgUpdateMembers{
				Creator:  sample.AccAddress(),
				ToRemove: []string{"invalid_address"},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid message",
			msg: MsgUpdateMembers{
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
