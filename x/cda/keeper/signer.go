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
// Assumes cdaId references an existing CDA
func (k Keeper) AppendSignerCDA(ctx sdk.Context, signerId uint64, cdaId uint64) {
	// Convert the CDA id to bytes
	byteCdaId := make([]byte, 8)
	binary.BigEndian.PutUint64(byteCdaId, cdaId)

	// Store the CDA id under the key CDA signer key (e.g. "CDA-signer-{signer ID}-{CDA ID}")
	store := k.getSignerCdaStore(ctx, signerId)
	store.Set(byteCdaId, []byte{0x1})
}

func (k Keeper) SignerHasCda(ctx sdk.Context, signerId uint64, cdaId uint64) bool {
	// Convert the CDA id to bytes
	bzCdaId := make([]byte, 8)
	binary.BigEndian.PutUint64(bzCdaId, cdaId)

	store := k.getSignerCdaStore(ctx, signerId)
	return store.Has(bzCdaId)
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

func (k Keeper) getSignerCdaStore(ctx sdk.Context, signerId uint64) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), types.SignerCdaStoreKey(signerId))
}
