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

	// Build the contract
	contract := types.Contract{
		// Id is set in AppendContract
		Description:       msg.Description,
		Authors:           msg.Authors,
		ContactInfo:       msg.ContactInfo,
		MoreInfoUri:       msg.MoreInfoUri,
		TemplateUri:       msg.TemplateUri,
		TemplateSchemaUri: msg.TemplateSchemaUri,
	}

	// Store the Contract in state
	m.AppendContract(ctx, contract)

	// Store the schema in state

	// Emit Events

	return &types.MsgRegisterContractResponse{Id: uint64(1)}, nil
}
