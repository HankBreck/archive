package keeper

import (
	"context"

	"archive/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RenounceIdentity(goCtx context.Context, msg *types.MsgRenounceIdentity) (*types.MsgRenounceIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRenounceIdentityResponse{}, nil
}
