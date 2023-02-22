package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/HankBreck/archive/x/cda/types"
)

// Stores cdaId in signer-prefixed storage.
// Assumes signer references an existing identity
func (k Keeper) AppendSignerCDA(ctx sdk.Context, signer uint64, cdaId uint64) {
	// Get current index for this owner
	count := k.GetSignerCDACount(ctx, signer)

	// Convert the index to bytes
	bzSignerIdx := make([]byte, 8)
	binary.BigEndian.PutUint64(bzSignerIdx, count)

	// Convert the CDA id to bytes
	byteCdaId := make([]byte, 8)
	binary.BigEndian.PutUint64(byteCdaId, cdaId)

	// Store the CDA id under the key CDA signer key (e.g. "CDA-signer-{signer ID}")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.SignerCdaStoreKey(signer))
	store.Set(bzSignerIdx, byteCdaId)

	// Increment the index in storage
	k.SetSignerCDACount(ctx, signer, count+1)
}

// Returns the next available index for storing the CDA id
func (k Keeper) GetSignerCDACount(ctx sdk.Context, signer uint64) uint64 {

	// Covert the storage key to bytes
	storePrefix := []byte(types.CDASignerCountKey)

	// Load the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storePrefix)

	// Get the current count of CDAs
	bzSignerId := make([]byte, 8)
	binary.BigEndian.PutUint64(bzSignerId, signer)
	bzCount := store.Get(bzSignerId)

	// Return 0 if the key is nil (first time accessing)
	if bzCount == nil {
		return 0
	}

	// Return count as uint64
	return binary.BigEndian.Uint64(bzCount)
}

// Increments the value of the owner's CDASignerCountKey ("CDA-signer-count-{signer identity ID}") in storage
func (k Keeper) SetSignerCDACount(ctx sdk.Context, signer uint64, count uint64) {

	// Build & convert the storage key to bytes
	storePrefix := []byte(types.CDASignerCountKey)

	// Load the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), storePrefix)

	// Convert count to bytes
	bzCount := make([]byte, 8)
	binary.BigEndian.PutUint64(bzCount, count)

	// Set the count in storage
	bzSignerId := make([]byte, 8)
	binary.BigEndian.PutUint64(bzSignerId, signer)
	store.Set(bzSignerId, bzCount)
}
