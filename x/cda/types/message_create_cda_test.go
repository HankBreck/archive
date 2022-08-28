package types

import (
	"testing"
	"time"

	"archive/testutil/sample"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

// TODO: Test the new fields
// Expiration
// Ownership
func TestMsgCreateCDA_ValidateBasic(t *testing.T) {
	// Valid test values
	validOwnership := map[string]uint64{
		sample.AccAddress(): 60_000_000,
	}
	validCid := "QmSrnQXUtGqsVRcgY93CdWXf8GPE9Zjj7Tg3SZUgLKDN5W"
	validExpiration := uint64(time.Now().UnixMilli()) + 5000 // current time + 5 seconds

	tests := []struct {
		name string
		msg  MsgCreateCDA
		err  error
	}{
		{
			name: "valid message",
			msg: MsgCreateCDA{
				Creator:    sample.AccAddress(),
				Cid:        validCid,
				Ownership:  validOwnership,
				Expiration: validExpiration,
			},
		}, {
			name: "invalid address",
			msg: MsgCreateCDA{
				Creator:    "invalid address",
				Cid:        validCid,
				Ownership:  validOwnership,
				Expiration: validExpiration,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid cid",
			msg: MsgCreateCDA{
				Creator:    sample.AccAddress(),
				Cid:        "invalid cid",
				Ownership:  validOwnership,
				Expiration: validExpiration,
			},
			err: ErrInvalidCid,
		}, {
			name: "invalid ownership",
			msg: MsgCreateCDA{
				Creator:    sample.AccAddress(),
				Cid:        validCid,
				Ownership:  make(map[string]uint64), // empty map
				Expiration: validExpiration,
			},
			err: ErrInvalidOwnership,
		}, {
			name: "invalid ownership address",
			msg: MsgCreateCDA{
				Creator: sample.AccAddress(),
				Cid:     validCid,
				Ownership: map[string]uint64{
					"invalid address": 60_000_000,
				},
				Expiration: validExpiration,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid expiration",
			msg: MsgCreateCDA{
				Creator:    sample.AccAddress(),
				Cid:        validCid,
				Ownership:  validOwnership,
				Expiration: uint64(time.Now().UnixMilli()) - 50_000, // current time - 50 seconds
			},
			err: ErrInvalidExpiration,
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
