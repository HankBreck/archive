package keeper

import (
	"encoding/binary"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/HankBreck/archive/x/cda/types"
)

// AppendCDA stores cda in with the CDAKey ("CDA-value-") and increments the count
// by one with the CDACountKey ("CDA-count-")
//
// Returns the id of cda
func (k Keeper) AppendCDA(ctx sdk.Context, cda types.CDA) uint64 {
	count := k.getCDACount(ctx)
	if k.HasCDA(ctx, count) {
		panic("Duplicate CDA id found" + strconv.FormatUint(count, 10))
	}
	cda.Id = count
	k.uncheckedSetCda(ctx, cda)
	k.setCDACount(ctx, count+1)
	return cda.Id
}

// GetCDA fetches the CDA stored under the key of cdaId. If no CDA is found, an error is thrown
func (k Keeper) GetCDA(ctx sdk.Context, cdaId uint64) (*types.CDA, error) {
	store := k.getCdaStore(ctx)
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, cdaId)

	var cda types.CDA
	bzCda := store.Get(bzKey)
	if len(bzCda) == 0 {
		return nil, types.ErrNonExistentCdaId
	}

	err := k.cdc.Unmarshal(bzCda, &cda)
	if err != nil {
		return nil, err
	}

	return &cda, nil
}

// HasCDA returns true if a CDA is stored under cdaId, else false
func (k Keeper) HasCDA(ctx sdk.Context, cdaId uint64) bool {
	store := k.getCdaStore(ctx)

	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, cdaId)

	return store.Has(bzKey)
}

func (k Keeper) uncheckedSetCda(ctx sdk.Context, cda types.CDA) {
	// Convert the id to bytes
	byteId := make([]byte, 8)
	binary.BigEndian.PutUint64(byteId, cda.Id)

	// Marshal the CDA to bytes
	byteCda := k.cdc.MustMarshal(&cda)

	// Store the CDA under the key of id
	store := k.getCdaStore(ctx)
	store.Set(byteId, byteCda)
}

func (k Keeper) getCdaStore(ctx sdk.Context) prefix.Store {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CDAKey))
	return store
}

// getCDACount returns the next available id for a CDA to use
func (k Keeper) getCDACount(ctx sdk.Context) uint64 {
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

// setCDACount increments the value of CDACountKey ("CDA-count-") in storage
func (k Keeper) setCDACount(ctx sdk.Context, count uint64) {
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
