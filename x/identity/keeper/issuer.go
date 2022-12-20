package keeper

import (
	"archive/x/identity/types"
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) AppendIssuer(ctx sdk.Context, issuer types.Issuer) uint64 {
	count := k.getIssuerCount(ctx)
	if k.HasIssuer(ctx, count) {
		panic("Duplicate Issuer ID found... this should never happen")
	}
	issuer.Id = count

	// TODO: Any more checks necessary?
	k.uncheckedSetIssuer(ctx, issuer)
	k.setIssuerCount(ctx, count+1)
	return count
}

// HasIssuer returns true if a issuer is stored under id, else false
func (k Keeper) HasIssuer(ctx sdk.Context, id uint64) bool {
	store := k.getIssuerStore(ctx)
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, id)
	return store.Has(bzKey)
}

// Stores the issuer with a key of issuer.Id. The issuer.Id field must be set by a calling function.
// The issuer passed as an argument is assumed to be valid, so calling functions must assure this.
func (k Keeper) uncheckedSetIssuer(ctx sdk.Context, issuer types.Issuer) {
	store := k.getIssuerStore(ctx)
	bzKey := make([]byte, 8)
	binary.BigEndian.PutUint64(bzKey, issuer.Id)
	bzContract := k.cdc.MustMarshal(&issuer)
	store.Set(bzKey, bzContract)
}

func (k Keeper) getIssuerStore(ctx sdk.Context) prefix.Store {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.IssuerKey))
	return store
}

func (k Keeper) getIssuerCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.IssuerCountKey))
	bzCount := store.Get([]byte{0})
	if bzCount == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bzCount)
}

func (k Keeper) setIssuerCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.IssuerCountKey))
	bzCount := make([]byte, 8)
	binary.BigEndian.PutUint64(bzCount, count)
	store.Set([]byte{0}, bzCount)
}
