package keeper

import (
	"context"

	"archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Cda(goCtx context.Context, req *types.QueryCdaRequest) (*types.QueryCdaResponse, error) {
	// Throw an error if the request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Unwrap context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Load the CDA store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CDAKey))

	//
	store.Get([]byte(req.Id))

	return &types.QueryCdaResponse{}, nil
}
