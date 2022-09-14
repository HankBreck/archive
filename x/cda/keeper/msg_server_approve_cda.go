package keeper

import (
	"context"

	"archive/x/cda/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ApproveCda(goCtx context.Context, msg *types.MsgApproveCda) (*types.MsgApproveCdaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgApproveCdaResponse{}, nil
}
