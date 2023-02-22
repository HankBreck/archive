package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/HankBreck/archive/x/cda/types"
)

// Stores cdaId in signer-prefixed storage.
// Assumes signer references an existing identity
func (k Keeper) AppendSignerCDA(ctx sdk.Context, signerId uint64, cdaId uint64) {
	// Get current index for this owner
	count := k.GetSignerCDACount(ctx, signerId)

	// Convert the index to bytes
	bzSignerIdx := make([]byte, 8)
	binary.BigEndian.PutUint64(bzSignerIdx, count)

	// Convert the CDA id to bytes
	byteCdaId := make([]byte, 8)
	binary.BigEndian.PutUint64(byteCdaId, cdaId)

	// Store the CDA id under the key CDA signer key (e.g. "CDA-signer-{signer ID}")
	store := k.getSignerCdaStore(ctx, signerId)
	store.Set(bzSignerIdx, byteCdaId)

	// Increment the index in storage
	k.SetSignerCDACount(ctx, signerId, count+1)
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

// GetCdasBySigner pages through all CDAs stored under the key of signerId.
//
// Returns a tuple of: the ids found, the page response, and an error.
func (k Keeper) GetCdasBySigner(ctx sdk.Context, signerId uint64, pageReq *query.PageRequest) ([]uint64, *query.PageResponse, error) {
	store := k.getSignerCdaStore(ctx, signerId)
	ids := []uint64{}

	// Unmarshal each key into the bech32 address
	pageRes, err := query.Paginate(store, pageReq, func(key []byte, _ []byte) error {
		// TODO: can this be invalid?
		id := binary.BigEndian.Uint64(key)
		ids = append(ids, id)
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return ids, pageRes, nil
}

// Increments the value of the owner's CDASignerCountKey ("CDA-signer-count-{signer identity ID}") in storage
func (k Keeper) SetSignerCDACount(ctx sdk.Context, signerId uint64, count uint64) {
	store := k.getSignerCdaStore(ctx, signerId)

	// Convert count to bytes
	bzCount := make([]byte, 8)
	binary.BigEndian.PutUint64(bzCount, count)

	// Set the count in storage
	bzSignerId := make([]byte, 8)
	binary.BigEndian.PutUint64(bzSignerId, signerId)
	store.Set(bzSignerId, bzCount)
}

func (k Keeper) getSignerCdaStore(ctx sdk.Context, signerId uint64) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.SignerCdaStoreKey(signerId))
}
