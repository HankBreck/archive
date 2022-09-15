package keeper

import (
	"context"
	"encoding/binary"

	"archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Returns the CDA for the supplied ID or an error if the id supplied in the request is invalid
func (k Keeper) Cda(goCtx context.Context, req *types.QueryCdaRequest) (*types.QueryCdaResponse, error) {
	// Throw an error if the request is nil
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// Unwrap context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Load the CDA store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.CDAKey))

	// Convert id to a uint64 and then to bytes
	byteKey := toUint64Bytes(req.Id)

	// Load the CDA bytes at key (panics if not found)
	byteCda := store.Get(byteKey)

	// Unmarshal the CDA bytes
	var cda types.CDA
	if err := k.cdc.Unmarshal(byteCda, &cda); err != nil {
		return nil, err
	}

	// Check if struct has zero-value fields (field not set)
	if cda.Creator == "" || cda.Cid == "" {
		return nil, sdkerrors.ErrNotFound
	}

	// Return query response with CDA data
	return &types.QueryCdaResponse{
		Cda: &cda,
	}, nil
}

func toUint64Bytes(id uint64) []byte {
	// Convert uint64 id to bytes
	byteId := make([]byte, 8)
	binary.BigEndian.PutUint64(byteId, id)
	return byteId
}
