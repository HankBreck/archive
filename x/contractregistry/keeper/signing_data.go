package keeper

import (
	"archive/x/contractregistry/types"
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Stores signingData with a key of id. The most recent contract stored must have a contract.Id == id.
// This function should only be called from RegisterContract.
//
// Panics on any id that is not equal to the current contract count - 1.
//
// Returns an error if the signingData is nil.
func (k Keeper) SetSigningData(ctx sdk.Context, signingData types.RawSigningData, id uint64) error {
	// Ensure we are setting the metadata for the most recent Contract ID
	if count := k.getContractCount(ctx); count != id+1 {
		panic(fmt.Sprintf("Unexpected value for contract ID in uncheckedSetSigningData! Expected %d but got %d", id+1, count))
	}
	err := k.uncheckedSetSigningData(ctx, signingData, id)
	if err != nil {
		return err
	}
	return nil
}

// GetSigningData fetches the contract stored under the key of id. If no contract is found, an error is returned
func (k Keeper) GetSigningData(ctx sdk.Context, id uint64) (types.RawSigningData, error) {
	store := k.getSigningDataStore(ctx)
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, id)

	bzData := store.Get(bzKey)
	if len(bzData) == 0 {
		return nil, types.ErrNonExistentSigningData.Wrapf("Signing data not found for id %d", id)
	}
	var result types.RawSigningData
	result.UnmarshalJSON(bzData)

	return result, nil
}

func (k Keeper) HasSigningData(ctx sdk.Context, id uint64) bool {
	store := k.getSigningDataStore(ctx)
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, id)

	return store.Has(bzKey)
}

// Stores the contract's signing data with the contract's ID as the key. The contract.Id field must be set by a calling function.
// The signing data and ID passed as arguments are assumed to be valid, so calling functions must assure this.
func (k Keeper) uncheckedSetSigningData(ctx sdk.Context, signingData types.RawSigningData, id uint64) error {
	store := k.getSigningDataStore(ctx)

	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, id)

	// NOTE: this implementation of MarshalJSON does not seem to return any errors.
	// Might want to add an error if bzData == []byte("null")
	bzData, err := signingData.MarshalJSON()
	if err != nil {
		return err
	}

	store.Set(bzKey, bzData)
	return nil
}

func (k Keeper) getSigningDataStore(ctx sdk.Context) prefix.Store {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SigningDataKey))
	return store
}
