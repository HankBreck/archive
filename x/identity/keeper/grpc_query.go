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

func (k Keeper) IssuerInfo(goCtx context.Context, req *types.QueryIssuerInfoRequest) (*types.QueryIssuerInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Validate issuer address
	_, err := sdk.AccAddressFromBech32(req.Issuer)
	if err != nil {
		return nil, err
	}

	// Fetch issuer from storage
	issuer, err := k.GetIssuer(ctx, req.Issuer)
	if err != nil {
		return nil, err
	}

	return &types.QueryIssuerInfoResponse{IssuerInfo: issuer}, nil
}

func (k Keeper) Identity(goCtx context.Context, req *types.QueryIdentityRequest) (*types.QueryIdentityResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Fetch certificate from storage
	cert, err := k.GetCertificate(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryIdentityResponse{Certificate: cert}, nil
}

func (k Keeper) Operators(goCtx context.Context, req *types.QueryOperatorsRequest) (*types.QueryOperatorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	operators, pageRes, err := k.GetOperators(ctx, req.Id, req.Pagination)
	if err != nil {
		return nil, err
	}

	return &types.QueryOperatorsResponse{
		Operators:  operators,
		Pagination: pageRes,
	}, nil
}

func (k Keeper) MemberRole(goCtx context.Context, req *types.QueryMemberRoleRequest) (*types.QueryMemberRoleResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryMemberRoleResponse{}, nil
}
