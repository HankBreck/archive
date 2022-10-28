package keeper

import (
	"archive/x/cda/types"
	"context"
	"encoding/binary"
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

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

func (k Keeper) CdasOwned(goCtx context.Context, req *types.QueryCdasOwnedRequest) (*types.QueryCdasOwnedResponse, error) {
	// Return an error if the request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Return an error if the owner is an invalid address
	_, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	var ids []uint64

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CDAOwnerKey+req.Owner))

	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		// Unmarshal the id
		id := binary.BigEndian.Uint64(value)
		// append to our result
		ids = append(ids, id)
		// return no error
		return nil
	})

	if err != nil {
		return nil, err
	}

	return &types.QueryCdasOwnedResponse{Ids: ids, Pagination: pageRes}, nil
}

func (k Keeper) Approval(goCtx context.Context, req *types.QueryApprovalRequest) (*types.QueryApprovalResponse, error) {
	// Respond with an error if the request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Ensure owner field is a valid address
	owner, err := sdk.AccAddressFromBech32(req.Owner)
	if err != nil {
		return nil, err
	}

	// Ensure CdaId is valid
	cdaStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CDAKey))
	bzId := make([]byte, 8)
	binary.BigEndian.PutUint64(bzId, req.CdaId)
	valid := cdaStore.Has(bzId)
	if !valid {
		return nil, sdkerrors.ErrKeyNotFound.Wrapf("Could not find the cda with an id of %d", req.CdaId)
	}

	keySuffix := strconv.FormatUint(req.CdaId, 10)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CDAApprovalKey+keySuffix))

	entry := store.Get(owner.Bytes())
	return &types.QueryApprovalResponse{Approved: entry != nil}, nil
}

func (k Keeper) Params(c context.Context, req *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	return &types.QueryParamsResponse{Params: k.GetParams(ctx)}, nil
}
