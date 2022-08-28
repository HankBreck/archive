package keeper

import (
	"context"

	"archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateCDA(goCtx context.Context, msg *types.MsgCreateCDA) (*types.MsgCreateCDAResponse, error) {
	// Unwrap the context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create the CDA
	var cda = types.CDA{
		Creator:    msg.Creator,
		Cid:        msg.Cid,
		Ownership:  msg.Ownership,
		Expiration: msg.Expiration,
	}

	// Store CDA & grab cda id
	id := k.AppendCDA(ctx, cda)
	for owner := range cda.Ownership {
		err := k.AppendOwnerCDA(ctx, owner, id)
		// TODO: check if we need some sort of transaction/rollback option in case this fails
		if err != nil {
			return nil, err
		}
	}

	// Return the id to the user
	return &types.MsgCreateCDAResponse{Id: id}, nil
}
