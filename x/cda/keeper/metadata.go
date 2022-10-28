package keeper

import (
	"archive/x/cda/types"
	"encoding/binary"

	sdktypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: Check that the contract id exists
func (k Keeper) SetMetadata(ctx sdk.Context, cdaId uint64, metadata *sdktypes.Any) error {
	// Ensure cdaId references an existing CDA
	if !k.HasCDA(ctx, cdaId) {
		return types.ErrNonExistentCdaId
	}

	// TODO: Should we throw an error if the metadata already exists?
	err := k.uncheckedSetMetadata(ctx, cdaId, metadata)
	return err

}

func (k Keeper) GetMetadata(ctx sdk.Context, cdaId uint64) (*sdktypes.Any, error) {
	store := k.getMetadataStore(ctx)

	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, cdaId)

	bzMetadata := store.Get(bzKey)
	metadata := &sdktypes.Any{}
	err := k.cdc.Unmarshal(bzMetadata, metadata)
	if err != nil {
		return nil, err
	}

	return metadata, nil
}

func (k Keeper) uncheckedSetMetadata(ctx sdk.Context, cdaId uint64, metadata *sdktypes.Any) error {
	store := k.getMetadataStore(ctx)

	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, cdaId)

	bzValue, err := k.cdc.Marshal(metadata)
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
