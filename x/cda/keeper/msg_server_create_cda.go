package keeper

import (
	"context"
	"strconv"

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
		Approved:   false,
	}

	// Store CDA & grab cda id
	id := k.AppendCDA(ctx, cda)
	for i := range cda.Ownership {
		owner := cda.Ownership[i]
		err := k.AppendOwnerCDA(ctx, owner.Owner, id)
		if err != nil {
			return nil, err
		}
	}

	// Return the id to the user
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(types.AttributeKeyCdaId, strconv.FormatUint(id, 10)),
	))

	return &types.MsgCreateCDAResponse{Id: id}, nil
}
