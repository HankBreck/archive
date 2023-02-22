package keeper

import (
	"context"

	"github.com/HankBreck/archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
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

// Returns the CDA for the supplied ID or an error if the id supplied in the request is invalid
func (k Keeper) Cda(goCtx context.Context, req *types.QueryCdaRequest) (*types.QueryCdaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Throw an error if the request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	cda, err := k.GetCDA(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	// Return query response with CDA data
	return &types.QueryCdaResponse{
		Cda: cda,
	}, nil
}

func (k Keeper) Cdas(goCtx context.Context, req *types.QueryCdasRequest) (*types.QueryCdasResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Throw an error if the request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var cdas []*types.CDA

	// Load the KV store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CDAKey))

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		// Unmarshal the CDA's bytes into the a new variable
		var cda types.CDA
		if err := k.cdc.Unmarshal(value, &cda); err != nil {
			return err
		}

		// If we successfully unmarshal the bytes, add to result
		cdas = append(cdas, &cda)

		return nil

	})

	// Throw error if pagination fails
	if err != nil {
		return nil, err
	}

	// Return list of CDA objects and new pagination information
	return &types.QueryCdasResponse{CDAs: cdas, Pagination: pageRes}, nil
}

func (k Keeper) CdasBySigner(goCtx context.Context, req *types.QueryCdasBySignerRequest) (*types.QueryCdasBySignerResponse, error) {
	// Return an error if the request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	ids, pageRes, err := k.GetCdasBySigner(ctx, req.Signer, req.Pagination)
	if err != nil {
		return nil, err
	}

	return &types.QueryCdasBySignerResponse{Ids: ids, Pagination: pageRes}, nil
}

func (k Keeper) Approval(goCtx context.Context, req *types.QueryApprovalRequest) (*types.QueryApprovalResponse, error) {
	// Respond with an error if the request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Ensure CdaId is valid
	if !k.HasCDA(ctx, req.CdaId) {
		return nil, types.ErrNonExistentCdaId.Wrapf("CDA with ID %d not found", req.CdaId)
	}

	// Check state for approval and respond
	hasApproval := k.HasApproval(ctx, req.CdaId, req.SignerId)
	return &types.QueryApprovalResponse{Approved: hasApproval}, nil
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

func (k Keeper) SigningDataSchema(goCtx context.Context, req *types.QuerySigningDataSchemaRequest) (*types.QuerySigningDataSchemaResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	data, err := k.GetSigningDataSchema(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &types.QuerySigningDataSchemaResponse{Schema: data}, nil
}
