package keeper

import (
	"archive/x/identity/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

// CreateMembership stores the first member in the "pending" prefixed store for the given certificateId.
// Pending memberships need to be approved by the recipient.
// It assumes that recipient is a valid address, so calling functions must ensure this.
// Panics if the certificate referenced by certificateId does not exist.
func (k Keeper) CreateMembership(ctx sdk.Context, certificateId uint64, recipient sdk.AccAddress) {
	// Ensure the certificate of ID exists
	if !k.HasCertificate(ctx, certificateId) {
		panic(types.ErrNonexistentCertificate.Wrapf("no certificate found for ID: %d", certificateId))
	}

	// Ensure membership for this ID does not exist (i.e. this is the initialization of the membership)
	if k.HasMember(ctx, certificateId, recipient) || k.HasPendingMember(ctx, certificateId, recipient) {
		panic(types.ErrExistingMember.Wrapf("certificateId: %d, address: %s", certificateId, recipient.String()))
	}

	k.uncheckedUpdateMembers(ctx, certificateId, []sdk.AccAddress{recipient}, []sdk.AccAddress{}, true)
}

// GetMembers pages through the members for a given identity. It separates the member lists
// into those that are pending and those that have accepted their membership.
//
// Returns a tuple of: the members found, the page response, and an error.
func (k Keeper) GetMembers(ctx sdk.Context, certificateId uint64, isPending bool, pageReq *query.PageRequest) ([]string, *query.PageResponse, error) {
	// Ensure certificateId exists in storage
	if !k.HasCertificate(ctx, certificateId) {
		return nil, nil, sdkerrors.ErrNotFound.Wrapf("A certificate with an ID of %d was not found", certificateId)
	}
	store := k.getMembershipStoreForId(ctx, certificateId, isPending)
	members := []string{}

	// Unmarshal each key into the bech32 address
	pageRes, err := query.Paginate(store, pageReq, func(key []byte, value []byte) error {
		var memberAddr sdk.AccAddress
		err := memberAddr.Unmarshal(key)
		if err != nil {
			return err
		}
		members = append(members, memberAddr.String())
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return members, pageRes, nil
}

// HasMember returns true if the member is an "accepted" member of the
// certificate referenced by certificateId.
func (k Keeper) HasMember(ctx sdk.Context, certificateId uint64, member sdk.AccAddress) bool {
	store := k.getMembershipStoreForId(ctx, certificateId, false)
	return store.Has(member.Bytes())
}

// HasMember returns true if the member is a "pending" member of the
// certificate referenced by certificateId.
func (k Keeper) HasPendingMember(ctx sdk.Context, certificateId uint64, member sdk.AccAddress) bool {
	store := k.getMembershipStoreForId(ctx, certificateId, true)
	return store.Has(member.Bytes())
}

// UpdateMembershipStatus transitions the state of the identity to accept or reject membership invitations.
// Returns an error if the certificate does not exist or the address is not a pending member.
func (k Keeper) UpdateMembershipStatus(ctx sdk.Context, certificateId uint64, member sdk.AccAddress, isAccept bool) error {
	// Ensure the certificate of ID exists
	if !k.HasCertificate(ctx, certificateId) {
		return types.ErrNonexistentCertificate.Wrapf("no certificate found for ID: %d", certificateId)
	}

	// Ensure membership is in the pending state
	if !k.HasPendingMember(ctx, certificateId, member) {
		return sdkerrors.ErrNotFound.Wrapf("member %s is not in the pending state", member.String())
	}

	if isAccept {
		// Add to accepted
		k.uncheckedUpdateMembers(ctx, certificateId, []sdk.AccAddress{member}, []sdk.AccAddress{}, false)
	}

	// Remove from pending
	k.uncheckedUpdateMembers(ctx, certificateId, []sdk.AccAddress{}, []sdk.AccAddress{member}, true)

	return nil
}

// UpdateMembers updates the pending membership list for the certificate referenced by id.
// Each address in the toAdd list is granted "pending" membership, whereas each address in
// toRemove's list is removed from the "pending" list.
//
// Returns an error if no certificate exists for the given certificateId.
func (k Keeper) UpdatePendingMembers(ctx sdk.Context, certificateId uint64, toAdd []sdk.AccAddress, toRemove []sdk.AccAddress) error {
	// Ensure certId exists
	if !k.HasCertificate(ctx, certificateId) {
		return types.ErrNonexistentCertificate.Wrapf("no certificate found for ID: %d", certificateId)
	}

	// Perform update
	k.uncheckedUpdateMembers(ctx, certificateId, toAdd, toRemove, true)
	return nil
}

// uncheckedUpdateMembers updates the membership list for the certificate referenced by id.
// isPending determines which membership list should be updated.
// Each address in the toAdd list is granted membership, whereas each address in
// toRemove's list is removed from the list.
//
// This function does not perform any sort of validity checks, so calling functions must
// perform checks before calling.
func (k Keeper) uncheckedUpdateMembers(ctx sdk.Context, id uint64, toAdd []sdk.AccAddress, toRemove []sdk.AccAddress, isPending bool) {
	store := k.getMembershipStoreForId(ctx, id, isPending)

	// Grant membership to each address
	for _, addr := range toAdd {
		store.Set(addr.Bytes(), []byte{0})
	}

	// Revoke membership from each address
	for _, addr := range toRemove {
		store.Delete(addr.Bytes())
	}
}

func (k Keeper) getMembershipStoreForId(ctx sdk.Context, id uint64, isPending bool) prefix.Store {
	keyPrefix := types.MembershipKeyPrefix(id, isPending)
	return prefix.NewStore(ctx.KVStore(k.storeKey), keyPrefix)
}
