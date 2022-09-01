package keeper

import (
	"context"
	"encoding/binary"

	"archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
