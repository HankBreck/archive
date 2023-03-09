package keeper

import (
	"encoding/binary"
	"strconv"

	"github.com/HankBreck/archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Adds the approval for the (CDA, owner) pair
//
// Returns an error if signer has already approved the CDA
func (k Keeper) SetApproval(ctx sdk.Context, cdaId uint64, signerId uint64) error {
	// Check if msgSigner has already approved the CDA
	if k.HasApproval(ctx, cdaId, signerId) {
		return types.ErrExistingApproval
	}

	// If not, update the store to include their address
	k.uncheckedSetApproval(ctx, cdaId, signerId)
	return nil
}

// Checks if the store contains an entry for signer.
// Returns true if an entry is found
func (k Keeper) HasApproval(ctx sdk.Context, cdaId uint64, signerId uint64) bool {
	store := k.getApprovalStore(ctx, cdaId)
	bzSignerId := make([]byte, 8)
	binary.BigEndian.PutUint64(bzSignerId, signerId)
	return store.Has(bzSignerId)
}

func (k Keeper) uncheckedSetApproval(ctx sdk.Context, cdaId uint64, signerId uint64) {
	store := k.getApprovalStore(ctx, cdaId)
	bzSignerId := make([]byte, 8)
	binary.BigEndian.PutUint64(bzSignerId, signerId)
	store.Set(bzSignerId, []byte{0x01})
}

func (k Keeper) getApprovalStore(ctx sdk.Context, cdaId uint64) prefix.Store {
	keySuffix := strconv.FormatUint(cdaId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CDAApprovalKey+keySuffix))
	return store
}
