package keeper

import (
	"bytes"

	"github.com/HankBreck/archive/x/identity/types"

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

	k.uncheckedUpdateMembers(ctx, certificateId, []sdk.AccAddress{recipient}, []sdk.AccAddress{}, true)
}

// GetMembers pages through the members for a given identity. It separates the member lists
// into those that are pending and those that have accepted their membership.
//
// Returns a tuple of: the members found, the page response, and an error.
func (k Keeper) GetMembers(ctx sdk.Context, certificateId uint64, isPending bool, includeRemoved bool, pageReq *query.PageRequest) ([]string, *query.PageResponse, error) {
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
		if includeRemoved || bytes.Equal([]byte{1}, value) {
			members = append(members, memberAddr.String())
		}
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return members, pageRes, nil
}

// HasMember returns true if the member is an "accepted" member of the
// certificate referenced by certificateId.
//
// Returns an error if the certificate does not exist
func (k Keeper) HasMember(ctx sdk.Context, certificateId uint64, member sdk.AccAddress) (bool, error) {
	// Ensure certificate exists
	if !k.HasCertificate(ctx, certificateId) {
		return false, types.ErrNonexistentCertificate.Wrapf("no certificate found for ID: %d", certificateId)
	}

	// Check store for member
	store := k.getMembershipStoreForId(ctx, certificateId, false)
	status := store.Get(member.Bytes())
	return bytes.Equal([]byte{1}, status), nil
}

// HasMember returns true if the member is a "pending" member of the
// certificate referenced by certificateId.
//
// Returns an error if the certificate does not exist
func (k Keeper) HasPendingMember(ctx sdk.Context, certificateId uint64, member sdk.AccAddress) (bool, error) {
	// Ensure certificate exists
	if !k.HasCertificate(ctx, certificateId) {
		return false, types.ErrNonexistentCertificate.Wrapf("no certificate found for ID: %d", certificateId)
	}

	// Check store for member
	store := k.getMembershipStoreForId(ctx, certificateId, true)
	status := store.Get(member.Bytes())
	return bytes.Equal([]byte{1}, status), nil
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

// UpdateMembers updates the accepted membership list for the certificate referenced by id.
// Each address in the toAdd list is granted "accepted" membership, whereas each address in
// toRemove's list is removed from the "accepted" list.
//
// Returns an error if no certificate exists for the given certificateId or if an address in
// toRemove is an operator
func (k Keeper) UpdateAcceptedMembers(ctx sdk.Context, certificateId uint64, toAdd []sdk.AccAddress, toRemove []sdk.AccAddress) error {
	// Ensure certId exists
	if !k.HasCertificate(ctx, certificateId) {
		return types.ErrNonexistentCertificate.Wrapf("no certificate found for ID: %d", certificateId)
	}

	// Ensure no members in toRemove are operators
	for _, addr := range toRemove {
		hasOp, err := k.HasOperator(ctx, certificateId, addr)
		if err != nil {
			return err
		}
		if hasOp {
			return types.ErrExistingOperator.Wrapf("address (%s) must be demoted from operator before it can be removed as a member", addr.String())
		}
	}

	// Perform update
	k.uncheckedUpdateMembers(ctx, certificateId, toAdd, toRemove, false)
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
		store.Set(addr.Bytes(), []byte{1})
	}

	// Revoke membership from each address
	for _, addr := range toRemove {
		store.Set(addr.Bytes(), []byte{0})
	}
}

func (k Keeper) getMembershipStoreForId(ctx sdk.Context, id uint64, isPending bool) prefix.Store {
	keyPrefix := types.MembershipKeyPrefix(id, isPending)
	return prefix.NewStore(ctx.KVStore(k.storeKey), keyPrefix)
}
