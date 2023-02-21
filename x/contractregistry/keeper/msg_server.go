package keeper

import (
	"context"
	"strconv"

	"github.com/HankBreck/archive/x/contractregistry/types"

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

	// Validate Basic checks:
	//		signingDataSchema != nil
	//		signingDataSchema is valid JSON
	if msg == nil {
		return nil, types.ErrInvalid.Wrap("Type MsgRegisterContract cannot be nil.")
	}

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
	id := m.AppendContract(ctx, contract)

	// Store the schema in state
	err := m.SetSigningData(ctx, msg.SigningDataSchema, id)
	if err != nil {
		return nil, err
	}

	// Emit Event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgRegisterContract,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(sdk.AttributeKeyAction, "RegisterContract"),
		sdk.NewAttribute("contract_id", strconv.FormatUint(id, 10)),
	))

	return &types.MsgRegisterContractResponse{Id: id}, nil
}