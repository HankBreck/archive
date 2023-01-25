package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/HankBreck/archive/x/cda/types"
)

// Stores cdaId in owner-prefixed storage
//
// Returns an error if the owner is not a valid bech32 address
func (k Keeper) AppendOwnerCDA(ctx sdk.Context, owner string, cdaId uint64) error {

	// // Validate address
	// ownerAddr, err := sdk.AccAddressFromBech32(owner)
	// if err != nil {
	// 	return err
	// }

	// Get current index for this owner
	count := k.GetOwnerCDACount(ctx, owner)

	// Convert the index to bytes
	byteOwnerIdx := make([]byte, 8)
	binary.BigEndian.PutUint64(byteOwnerIdx, count)

	// Convert the CDA id to bytes
	byteCdaId := make([]byte, 8)
	binary.BigEndian.PutUint64(byteCdaId, cdaId)

	// Store the CDA id under the key CDA owner key (e.g. "CDA-owner-{address}")
	storageKey := []byte(types.CDAOwnerKey + owner)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storageKey)
	store.Set(byteOwnerIdx, byteCdaId)

	// Increment the index in storage
	k.SetOwnerCDACount(ctx, owner, count+1)

	return nil
}

// Returns the next available index for storing the CDA id
func (k Keeper) GetOwnerCDACount(ctx sdk.Context, owner string) uint64 {

	// Covert the storage key to bytes
	storePrefix := []byte(types.CDAOwnerCountKey)

	// Load the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storePrefix)

	// Get the current count of CDAs
	byteCount := store.Get([]byte(owner))

	// Return 0 if the key is nil (first time accessing)
	if byteCount == nil {
		return 0
	}

	// Return count as uint64
	return binary.BigEndian.Uint64(byteCount)
}

// Increments the value of the owner's CDAOwnerCountKey ("CDA-owner-count-{owner address}") in storage
func (k Keeper) SetOwnerCDACount(ctx sdk.Context, owner string, count uint64) {

	// Build & convert the storage key to bytes
	storePrefix := []byte(types.CDAOwnerCountKey)

	// Load the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storePrefix)

	// Convert count to bytes
	byteCount := make([]byte, 8)
	binary.BigEndian.PutUint64(byteCount, count)

	// Set the count in storage
	store.Set([]byte(owner), byteCount)
}
