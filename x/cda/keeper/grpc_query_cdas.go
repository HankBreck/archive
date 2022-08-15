package keeper

import (
	"context"

	"archive/x/cda/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Cdas(goCtx context.Context, req *types.QueryCdasRequest) (*types.QueryCdasResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryCdasResponse{}, nil
}
