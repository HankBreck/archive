package keeper

import (
	"context"

	"archive/x/archive/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateCDA(goCtx context.Context, msg *types.MsgCreateCDA) (*types.MsgCreateCDAResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateCDAResponse{}, nil
}
