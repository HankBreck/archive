package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"archive/x/cda/types"
)

// Stores cda in with the CDAKey ("CDA-value-") and increments the count
// by one with the CDACountKey ("CDA-count-")
//
// Returns the id of cda
func (k Keeper) AppendCDA(ctx sdk.Context, cda types.CDA) uint64 {
	// Get current CDA ID
	count := k.GetCDACount(ctx)

	// Set the id of the CDA
	cda.Id = count

	// Convert the id to bytes
	byteId := make([]byte, 8)
	binary.BigEndian.PutUint64(byteId, cda.Id)

	// Marshal the CDA to bytes
	byteCda := k.cdc.MustMarshal(&cda)

	// Store the CDA under the key of id
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CDAKey))
	store.Set(byteId, byteCda)

	// Increment the CDA count in storage
	k.SetCDACount(ctx, count+1)

	// Return the stored CDA's id
	return cda.Id
}

// Returns the next available id for a CDA to use
func (k Keeper) GetCDACount(ctx sdk.Context) uint64 {
	// Convert the key to bytes
	byteKey := []byte(types.CDACountKey)

	// Load the store using the module's storage key ("cda") and CDACountKey ("CDA-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)

	// Get the current count of CDAs
	byteCount := store.Get(byteKey)

	// Return 0 if the key is nil (first time accessing)
	if byteCount == nil {
		return 0
	}

	// Return count as uint64
	return binary.BigEndian.Uint64(byteCount)
}

// Increments the value of CDACountKey ("CDA-count-") in storage
func (k Keeper) SetCDACount(ctx sdk.Context, count uint64) {
	// Convert the key to bytes
	byteKey := []byte(types.CDACountKey)

	// Load the store using the module's storage key ("cda") and CDACountKey ("CDA-count-")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)

	// Convert count to bytes
	byteCount := make([]byte, 8)
	binary.BigEndian.PutUint64(byteCount, count)

	// Set the count in storage
	store.Set(byteKey, byteCount)
}
