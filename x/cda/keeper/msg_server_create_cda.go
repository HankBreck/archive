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
		Creator: msg.Creator,
		Cid:     msg.Cid,
	}

	ctx.Logger().Info("Creator: " + cda.Creator)
	ctx.Logger().Info("Cid: " + cda.Cid)

	// Store CDA & grab cda id
	id := k.AppendCDA(ctx, cda)

	// Return the id to the user
	return &types.MsgCreateCDAResponse{Id: id}, nil
}
