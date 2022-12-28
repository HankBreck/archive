package keeper

import (
	"context"

	"archive/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AcceptIdentity(goCtx context.Context, msg *types.MsgAcceptIdentity) (*types.MsgAcceptIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAcceptIdentityResponse{}, nil
}
