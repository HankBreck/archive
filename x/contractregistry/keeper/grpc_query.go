package keeper

import (
	"archive/x/contractregistry/types"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Params(goCtx context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}

func (k Keeper) Contract(goCtx context.Context, req *types.QueryContractRequest) (*types.QueryContractResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	contract, err := k.GetContract(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &types.QueryContractResponse{Contract: *contract}, nil
}

func (k Keeper) Contracts(goCtx context.Context, req *types.QueryContractsRequest) (*types.QueryContractsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	contracts, pageRes, err := k.GetContracts(ctx, req.Pagination)
	if err != nil {
		return nil, err
	}
	return &types.QueryContractsResponse{
		Contracts:  contracts,
		Pagination: pageRes,
	}, nil
}

func (k Keeper) SigningData(goCtx context.Context, req *types.QuerySigningDataRequest) (*types.QuerySigningDataResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	data, err := k.GetSigningData(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &types.QuerySigningDataResponse{SigningData: data}, nil
}
