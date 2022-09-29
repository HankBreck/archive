package keeper

import (
	"context"

	"archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Approval(goCtx context.Context, req *types.QueryApprovalRequest) (*types.QueryApprovalResponse, error) {
	// Respond with an error if the request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryApprovalResponse{}, nil
}