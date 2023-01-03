package keeper

import (
	"context"

	"archive/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddIdentityMember(goCtx context.Context, msg *types.MsgAddIdentityMember) (*types.MsgAddIdentityMemberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgAddIdentityMemberResponse{}, nil
}
