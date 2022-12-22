package keeper

import (
	"archive/x/identity/types"
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// CreateMembership stores the first member in the prefixed store for the given certificateId.
// It assumes that recipient is a valid address, so calling functions must ensure this.
// Panics if the certificate referenced by certificateId does not exist.
func (k Keeper) CreateMembership(ctx sdk.Context, certificateId uint64, recipient sdk.AccAddress) error {
	// ensure the certificate of ID exists
	if !k.HasCertificate(ctx, certificateId) {
		panic(fmt.Sprintf("No certificate found for ID %d... this should never happen", certificateId))
	}

	// ensure membership for this ID does not exist
	if k.HasMember(ctx, certificateId, recipient) {
		return types.ErrExistingMember.Wrapf("certificateId: %d, address: %s", certificateId, recipient.String())
	}

	k.uncheckedUpdateMembers(ctx, certificateId, []sdk.AccAddress{recipient}, []sdk.AccAddress{})
	return nil
}

// HasMember returns true if the member is a member of the certificate referenced by certificateId.
func (k Keeper) HasMember(ctx sdk.Context, certificateId uint64, member sdk.AccAddress) bool {
	store := k.getMembershipStoreForId(ctx, certificateId)
	return store.Has(member.Bytes())
}

// uncheckedUpdateMembers updates the membership list for the certificate referenced by id.
// Each address in the toAdd list is granted membership, whereas each address in
// toRemove's membership is revoked.
// All parameters are assumed to be valid (existing and correct), so calling functions must ensure this.
func (k Keeper) uncheckedUpdateMembers(ctx sdk.Context, id uint64, toAdd []sdk.AccAddress, toRemove []sdk.AccAddress) {
	store := k.getMembershipStoreForId(ctx, id)

	// Grant membership to each address
	for _, addr := range toAdd {
		store.Set(addr.Bytes(), []byte{0})
	}

	// Revoke membership from each address
	for _, addr := range toRemove {
		store.Delete(addr.Bytes())
	}
}

func (k Keeper) getMembershipStoreForId(ctx sdk.Context, id uint64) prefix.Store {
	keyPrefix := types.MembershipKeyPrefix(id)
	return prefix.NewStore(ctx.KVStore(k.storeKey), keyPrefix)
}
