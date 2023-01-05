package keeper

import (
	"context"

	"archive/x/identity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	memberAddr, err := sdk.AccAddressFromBech32(req.Member)
	if err != nil {
		return nil, err
	}

	// Check if member is an operator
	hasOp, err := k.HasOperator(ctx, req.Id, memberAddr)
	if err != nil {
		return nil, err
	} else if hasOp {
		return &types.QueryMemberRoleResponse{Role: "operator"}, nil
	}

	// Check if member is an accepted member
	hasAccepted, err := k.HasMember(ctx, req.Id, memberAddr)
	if err != nil {
		return nil, err
	} else if hasAccepted {
		return &types.QueryMemberRoleResponse{Role: "member"}, nil
	}

	// Check if member is a pending member
	hasPending, err := k.HasMember(ctx, req.Id, memberAddr)
	if err != nil {
		return nil, err
	} else if hasPending {
		return &types.QueryMemberRoleResponse{Role: "pending-member"}, nil
	}

	return nil, sdkerrors.ErrNotFound.Wrapf("account (%s) is not a member of identity %d", memberAddr.String(), req.Id)
}
