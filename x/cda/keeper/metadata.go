package keeper

import (
	"encoding/binary"

	"github.com/HankBreck/archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: Check that the contract id exists
func (k Keeper) SetSigningData(ctx sdk.Context, cdaId uint64, metadata types.RawSigningData) error {
	// Ensure cdaId references an existing CDA
	if !k.HasCDA(ctx, cdaId) {
		return types.ErrNonExistentCdaId
	}

	// TODO: Should we throw an error if the metadata already exists?
	err := k.uncheckedSetMetadata(ctx, cdaId, metadata)
	return err

}

func (k Keeper) GetSigningData(ctx sdk.Context, cdaId uint64) (types.RawSigningData, error) {
	store := k.getMetadataStore(ctx)

	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, cdaId)

	bzMetadata := store.Get(bzKey)
	if len(bzMetadata) == 0 {
		return nil, types.ErrNonExistentSigningData.Wrapf("id %d not found", cdaId)
	}
	var metadata types.RawSigningData
	err := metadata.UnmarshalJSON(bzMetadata)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}

func (k Keeper) uncheckedSetMetadata(ctx sdk.Context, cdaId uint64, metadata types.RawSigningData) error {
	store := k.getMetadataStore(ctx)

	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, cdaId)

	bzValue, err := metadata.MarshalJSON()
	if err != nil {
		return err
	}

	store.Set(bzKey, bzValue)

	return nil
}

func (k Keeper) getMetadataStore(ctx sdk.Context) prefix.Store {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CDAMetadataKey))
	return store
}
