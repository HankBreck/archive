package keeper

import (
	"archive/x/identity/types"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) RegisterIssuer(goCtx context.Context, msg *types.MsgRegisterIssuer) (*types.MsgRegisterIssuerResponse, error) {
	if msg == nil {
		return nil, types.ErrInvalid.Wrap("Type MsgRegisterContract cannot be nil.")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegisterIssuerResponse{}, nil
}
