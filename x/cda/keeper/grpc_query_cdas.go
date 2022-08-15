package keeper

import (
	"context"

	"archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Cdas(goCtx context.Context, req *types.QueryCdasRequest) (*types.QueryCdasResponse, error) {
	// Throw an error if the request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Unwrap context
	ctx := sdk.UnwrapSDKContext(goCtx)

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
