package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) hasOperatorOrIssuer(ctx sdk.Context, id uint64, address sdk.AccAddress) (bool, error) {
	// Check if address is an operator of the id
	hasOper, err := k.HasOperator(ctx, id, address)
	if err != nil {
		return false, err
	} else if hasOper {
		return true, nil
	}

	// Check if address is the issuer of the id
	hasIssuer, err := k.HasIssuerForId(ctx, id, address)
	if err != nil {
		return false, err
	}

	return hasIssuer, nil
}
