package keeper

import (
	"archive/x/cda/types"
	"encoding/binary"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Adds the approval for the (CDA, owner) pair
//
// Returns an error if:
//
// (1) The CDA does not exist
// (2) The message's Ownership object does not match the stored Ownership object
// (3) The Creator has already approved the CDA
func (k Keeper) SetApproval(ctx sdk.Context, msg *types.MsgApproveCda) error {
	// Validate Creator address
	signer := sdk.MustAccAddressFromBech32(msg.Creator)

	// Ensure the CDA exists
	cdaStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CDAKey))
	bzCdaId := make([]byte, 8)
	binary.BigEndian.PutUint64(bzCdaId, msg.CdaId)
	bzCda := cdaStore.Get(bzCdaId)
	if bzCda == nil {
		return types.ErrNonExistentCdaId
	}
	var cda types.CDA
	if err := k.cdc.Unmarshal(bzCda, &cda); err != nil {
		return err
	}

	// Ensure msg.Ownership matches cda.Ownership
	if len(msg.Ownership) != len(cda.Ownership) {
		// TODO: make a more informative error type / message
		return sdkerrors.Wrap(types.ErrInvalidOwnership, "wrong ownership list length")
	}
	for i := range msg.Ownership {
		if *msg.Ownership[i] != *cda.Ownership[i] {
			return sdkerrors.Wrap(types.ErrInvalidOwnership, "wrong ownership list order")
		}
	}

	// Check if Creator has already approved the CDA
	keySuffix := strconv.FormatUint(msg.CdaId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CDAApprovalKey+keySuffix))
	bzApproval := store.Get(signer.Bytes())
	if bzApproval != nil {
		return types.ErrExistingApproval
	}

	// If not, update the store to include their address
	store.Set(signer.Bytes(), signer.Bytes())
	return nil
}
