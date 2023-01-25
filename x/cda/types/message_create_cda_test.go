package types

import (
	"testing"

	"github.com/HankBreck/archive/testutil/sample"
)

// TODO: Test the new fields
// Expiration
// Ownership
func TestMsgCreateCDA_ValidateBasic(t *testing.T) {
	// Valid test values
	var validOwnerships = make([]*Ownership, 1)
	validOwnerships = append(validOwnerships, &Ownership{
		Owner:     sample.AccAddress(),
		Ownership: 100_000_000,
	})
	// validCid := "QmSrnQXUtGqsVRcgY93CdWXf8GPE9Zjj7Tg3SZUgLKDN5W"
	// validExpiration := uint64(time.Now().UnixMilli()) + 5000 // current time + 5 seconds

	var invalidOwnerships = make([]*Ownership, 1)
	invalidOwnerships = append(invalidOwnerships, &Ownership{
		Owner:     "invalid address",
		Ownership: 100_000_000,
	})

	// tests := []struct {
	// 	name string
	// 	msg  MsgCreateCDA
	// 	err  error
	// }{
	// 	{
	// 		name: "valid message",
	// 		msg: MsgCreateCDA{
	// 			Creator:    sample.AccAddress(),
	// 			Cid:        validCid,
	// 			Ownership:  validOwnerships,
	// 			Expiration: validExpiration,
	// 		},
	// 	}, {
	// 		name: "invalid address",
	// 		msg: MsgCreateCDA{
	// 			Creator:    "invalid address",
	// 			Cid:        validCid,
	// 			Ownership:  validOwnerships,
	// 			Expiration: validExpiration,
	// 		},
	// 		err: sdkerrors.ErrInvalidAddress,
	// 	}, {
	// 		name: "invalid cid",
	// 		msg: MsgCreateCDA{
	// 			Creator:    sample.AccAddress(),
	// 			Cid:        "invalid cid",
	// 			Ownership:  validOwnerships,
	// 			Expiration: validExpiration,
	// 		},
	// 		err: ErrInvalidCid,
	// 	}, {
	// 		name: "invalid ownership",
	// 		msg: MsgCreateCDA{
	// 			Creator:    sample.AccAddress(),
	// 			Cid:        validCid,
	// 			Ownership:  make([]*Ownership, 1), // empty map
	// 			Expiration: validExpiration,
	// 		},
	// 		err: ErrInvalidOwnership,
	// 	}, {
	// 		name: "invalid ownership address",
	// 		msg: MsgCreateCDA{
	// 			Creator:    sample.AccAddress(),
	// 			Cid:        validCid,
	// 			Ownership:  invalidOwnerships,
	// 			Expiration: validExpiration,
	// 		},
	// 		err: sdkerrors.ErrInvalidAddress,
	// 	}, {
	// 		name: "invalid expiration",
	// 		msg: MsgCreateCDA{
	// 			Creator:    sample.AccAddress(),
	// 			Cid:        validCid,
	// 			Ownership:  validOwnerships,
	// 			Expiration: uint64(time.Now().UnixMilli()) - 50_000, // current time - 50 seconds
	// 		},
	// 		err: ErrInvalidExpiration,
	// 	},
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		err := tt.msg.ValidateBasic()
	// 		if tt.err != nil {
	// 			require.ErrorIs(t, err, tt.err)
	// 			return
	// 		}
	// 		require.NoError(t, err)
	// 	})
	// }
}
