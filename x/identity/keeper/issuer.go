package keeper

import (
	"archive/x/identity/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// SetIssuer stores the Issuer object under the creator's key.
// Returns an error if the creator address has already created an Issuer.
func (k Keeper) SetIssuer(ctx sdk.Context, issuer types.Issuer) error {
	if k.HasIssuer(ctx, issuer.Creator) {
		return types.ErrExistingIssuer
	}
	k.uncheckedSetIssuer(ctx, issuer)
	return nil
}

// GetIssuer returns the Issuer object created by creator.
// Returns an error if the creator has not created an issuer
func (k Keeper) GetIssuer(ctx sdk.Context, creator string) (*types.Issuer, error) {
	// Fetch Issuer from store
	store := k.getIssuerStore(ctx)
	bzIssuer := store.Get([]byte(creator))

	// Check if Issuer exists
	if len(bzIssuer) == 0 {
		return nil, sdkerrors.ErrNotFound.Wrapf("No Issuer found for address %s", creator)
	}

	// Unmarshal Issuer
	var issuer types.Issuer
	err := k.cdc.Unmarshal(bzIssuer, &issuer)
	if err != nil {
		return nil, err
	}

	return &issuer, nil
}

// HasIssuer returns true if a issuer is stored under the creator's address, else false
func (k Keeper) HasIssuer(ctx sdk.Context, creator string) bool {
	store := k.getIssuerStore(ctx)
	return store.Has([]byte(creator))
}

// Stores the issuer with a key of issuer.Creator.
// The issuer passed as an argument is assumed to be valid, so calling functions must assure this.
func (k Keeper) uncheckedSetIssuer(ctx sdk.Context, issuer types.Issuer) {
	store := k.getIssuerStore(ctx)
	bzContract := k.cdc.MustMarshal(&issuer)
	store.Set([]byte(issuer.Creator), bzContract)
}

func (k Keeper) getIssuerStore(ctx sdk.Context) prefix.Store {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.IssuerKey))
	return store
}
