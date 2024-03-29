package keeper

import (
	"github.com/HankBreck/archive/x/identity/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
)

//	Add operator store
//		- SetOperator
//			- operator (or issuer) can add a new operator
//		- GetOperators (paginated)
// 		- RemoveOperator
//			- operator (or issuer) can remove an operator

// SetOperators stores the operator entries under each account's address. Operators are stored
// in their own prefixed storage.
// Operators must be accepted members, and membership may not be revoked for an account that is currently an operator.
// This means operators must be demoted before their membership status can be revoked.
//
// Returns an error if the certificate doesn't exist or an account is not an accepted member.
func (k Keeper) SetOperators(ctx sdk.Context, certificateId uint64, operators []sdk.AccAddress) error {
	// Ensure certificate exists
	if !k.HasCertificate(ctx, certificateId) {
		return types.ErrNonexistentCertificate.Wrapf("could not find identity %d", certificateId)
	}
	store := k.getOperatorStoreForId(ctx, certificateId)
	for _, op := range operators {
		hasOp, err := k.HasOperator(ctx, certificateId, op)
		if err != nil {
			return err
		} else if hasOp {
			return types.ErrExistingOperator.Wrapf("%s is already an operator", op.String())
		}
		// Ensure new operator is an accepted member of the identity
		hasMember, err := k.HasMember(ctx, certificateId, op)
		if err != nil {
			return err
		} else if !hasMember {
			return sdkerrors.ErrNotFound.Wrapf("new operator is not a member of identity %d", certificateId)
		}
		// Set value in operator store
		store.Set(op.Bytes(), []byte{0})
	}
	return nil
}

// SetInitialOperator stored the operator entry under the recipient's account address. This method assumes
// that the certificate exists and does not enforce the accepted membership requirement of the SetOperators
// method. It should only be called from the IssueCertificate msgServer method.
func (k Keeper) SetInitialOperator(ctx sdk.Context, certificateId uint64, recipient sdk.AccAddress) {
	store := k.getOperatorStoreForId(ctx, certificateId)
	store.Set(recipient.Bytes(), []byte{0})
}

// GetMembers pages through the members for a given identity.
//
// Returns a tuple of: the operators found, the page response, and an error.
func (k Keeper) GetOperators(ctx sdk.Context, certificateId uint64, pageReq *query.PageRequest) ([]string, *query.PageResponse, error) {
	if !k.HasCertificate(ctx, certificateId) {
		return nil, nil, types.ErrNonexistentCertificate
	}

	// Paginate over next operators
	store := k.getOperatorStoreForId(ctx, certificateId)
	operators := []string{}
	pageRes, err := query.Paginate(store, pageReq, func(key []byte, _ []byte) error {
		var operAddr sdk.AccAddress
		err := operAddr.Unmarshal(key)
		if err != nil {
			return err
		}
		operators = append(operators, operAddr.String())
		return nil
	})
	if err != nil {
		return nil, nil, err
	}

	return operators, pageRes, nil
}

// HasMember returns true if the account is an operator of the
// certificate referenced by certificateId.
//
// Returns an error if the certificate does not exist
func (k Keeper) HasOperator(ctx sdk.Context, certificateId uint64, operator sdk.AccAddress) (bool, error) {
	if !k.HasCertificate(ctx, certificateId) {
		return false, types.ErrNonexistentCertificate.Wrapf("no certificate found for ID: %d", certificateId)
	}
	store := k.getOperatorStoreForId(ctx, certificateId)
	return store.Has(operator.Bytes()), nil
}

// RemoveOperators removes each operator's address from the address-prefixed operator store.
//
// Returns an error if the certificate doesn't exist
func (k Keeper) RemoveOperators(ctx sdk.Context, certificateId uint64, operators []sdk.AccAddress) error {
	// Ensure certificate exists
	if !k.HasCertificate(ctx, certificateId) {
		return types.ErrNonexistentCertificate.Wrapf("could not find identity %d", certificateId)
	}
	store := k.getOperatorStoreForId(ctx, certificateId)
	for _, op := range operators {
		// Remove entry from operator store
		store.Delete(op.Bytes())
	}
	return nil
}

func (k Keeper) getOperatorStoreForId(ctx sdk.Context, id uint64) prefix.Store {
	keyPrefix := types.OperatorKeyPrefix(id)
	return prefix.NewStore(ctx.KVStore(k.storeKey), keyPrefix)
}
