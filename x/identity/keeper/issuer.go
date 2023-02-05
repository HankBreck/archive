package keeper

import (
	"github.com/HankBreck/archive/x/identity/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

// SetIssuer stores the Issuer object under the creator's account address.
// Returns an error if the creator address has already created an Issuer.
func (k Keeper) SetIssuer(ctx sdk.Context, issuer types.Issuer) error {
	creator, err := sdk.AccAddressFromBech32(issuer.Creator)
	if err != nil {
		return err
	}
	if k.HasIssuer(ctx, creator) {
		return types.ErrExistingIssuer
	}
	k.uncheckedSetIssuer(ctx, issuer, creator)
	return nil
}

// GetIssuer returns the Issuer object created by creator.
// Returns an error if the creator has not created an issuer
func (k Keeper) GetIssuer(ctx sdk.Context, creator sdk.AccAddress) (*types.Issuer, error) {
	// Fetch Issuer from store
	store := k.getIssuerStore(ctx)
	bzIssuer := store.Get(creator.Bytes())

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

// GetIssuers pages through all registered issuers.
//
// Returns a tuple of: the issuers found, the page response, and an error.
func (k Keeper) GetIssuers(ctx sdk.Context, pageReq *query.PageRequest) ([]sdk.AccAddress, *query.PageResponse, error) {
	store := k.getIssuerStore(ctx)
	var issuers []sdk.AccAddress

	// Unmarshal each key into the bech32 address
	pageRes, err := query.Paginate(store, pageReq, func(key []byte, _ []byte) error {
		var issuerAddr sdk.AccAddress
		err := issuerAddr.Unmarshal(key)
		if err != nil {
			return err
		}
		issuers = append(issuers, issuerAddr)
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return issuers, pageRes, nil
}

// HasIssuer returns true if a issuer is stored under the creator's address, else false
func (k Keeper) HasIssuer(ctx sdk.Context, creator sdk.AccAddress) bool {
	store := k.getIssuerStore(ctx)
	return store.Has(creator.Bytes())
}

// Stores the issuer with a key of issuer.Creator.
// The issuer passed as an argument is assumed to be valid, so calling functions must assure this.
func (k Keeper) uncheckedSetIssuer(ctx sdk.Context, issuer types.Issuer, creatorAddr sdk.AccAddress) {
	store := k.getIssuerStore(ctx)
	bzContract := k.cdc.MustMarshal(&issuer)
	store.Set(creatorAddr.Bytes(), bzContract)
}

func (k Keeper) getIssuerStore(ctx sdk.Context) prefix.Store {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.IssuerKey))
	return store
}
