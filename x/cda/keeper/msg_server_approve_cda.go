package keeper

import (
	"context"

	"archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ApproveCda(goCtx context.Context, msg *types.MsgApproveCda) (*types.MsgApproveCdaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.SetApproval(ctx, msg)

	if err != nil {
		return nil, err
	}

	return &types.MsgApproveCdaResponse{}, nil
}
