package keeper

import (
	"context"

	"archive/x/identity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Keeper) IdentityMembers(goCtx context.Context, req *types.QueryIdentityMembersRequest) (*types.QueryIdentityMembersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	members, pageRes, err := k.GetMembers(ctx, req.Id, req.IsPending, req.Pagination)
	if err != nil {
		return nil, err
	}

	return &types.QueryIdentityMembersResponse{
		Members:    members,
		Pagination: pageRes,
	}, nil
}

func (k Keeper) Issuers(goCtx context.Context, req *types.QueryIssuersRequest) (*types.QueryIssuersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	issuers, pageRes, err := k.GetIssuers(ctx, req.Pagination)
	if err != nil {
		return nil, err
	}

	return &types.QueryIssuersResponse{
		Issuers:    issuers,
		Pagination: pageRes,
	}, nil
}
