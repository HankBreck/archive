package keeper

import (
	"archive/x/contractregistry/types"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = msgServer{}

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (m msgServer) RegisterContract(goCtx context.Context, msg *types.MsgRegisterContract) (*types.MsgRegisterContractResponse, error) {

	// Unwrap the context
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	// Store the Contract in state

	// Store the schema in state

	// Emit Events

	return &types.MsgRegisterContractResponse{Id: uint64(1)}, nil
}
