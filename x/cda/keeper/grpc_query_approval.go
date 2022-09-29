package keeper

import (
	"context"
	"encoding/binary"
	"strconv"

	"archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
