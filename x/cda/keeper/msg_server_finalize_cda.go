package keeper

import (
	"context"

	"archive/x/cda/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) FinalizeCda(goCtx context.Context, msg *types.MsgFinalizeCda) (*types.MsgFinalizeCdaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgFinalizeCdaResponse{}, nil
}