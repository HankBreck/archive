package keeper

import (
	"context"

	"archive/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) UpdateOperators(goCtx context.Context, msg *types.MsgUpdateOperators) (*types.MsgUpdateOperatorsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateOperatorsResponse{}, nil
}
