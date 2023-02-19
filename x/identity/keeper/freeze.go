package keeper

import (
	"encoding/binary"

	"github.com/HankBreck/archive/x/identity/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Assumes:
//
//	msg_server has checked that the identity is not already frozen
//	msg_server has checked that the issuer is the sender
func (k Keeper) Freeze(ctx sdk.Context, certificateId uint64) error {
	// Ensure certificate exists
	if !k.HasCertificate(ctx, certificateId) {
		return types.ErrNonexistentCertificate.Wrapf("certificate not found for id %d", certificateId)
	}

	// Set ID as frozen in state
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, certificateId)
	store := k.getFreezeStore(ctx, certificateId)
	store.Set(bzKey, []byte{1})

	return nil
}

func (k Keeper) IsFrozen(ctx sdk.Context, certificateId uint64) bool {
	store := k.getFreezeStore(ctx, certificateId)
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, certificateId)
	return store.Has(bzKey)
}

func (k Keeper) getFreezeStore(ctx sdk.Context, id uint64) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.FreezeKey))
}
