package keeper

import (
	"context"
	"strconv"

	"github.com/HankBreck/archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateCda(goCtx context.Context, msg *types.MsgCreateCda) (*types.MsgCreateCdaResponse, error) {
	// Unwrap the context
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Create the CDA
	var cda = types.CDA{
		Creator:          msg.Creator,
		Id:               0, // Remove
		SigningParties:   msg.SigningParties,
		ContractId:       msg.ContractId,
		LegalMetadataUri: msg.LegalMetadataUri,
		UtcExpireTime:    msg.UtcExpireTime,
		Status:           types.CDA_Pending,
	}

	// Ensure the signing data matches the expected contract's schema
	matches, err := k.MatchesSigningDataSchema(ctx, msg.ContractId, msg.SigningData)
	if err != nil || !matches {
		return nil, types.ErrInvalidSigningData.Wrap("does not match schema")
	}

	// Store CDA & grab cda id
	id := k.AppendCDA(ctx, cda)
	for i := range cda.SigningParties {
		owner := cda.SigningParties[i]
		err := k.AppendOwnerCDA(ctx, owner, id)
		if err != nil {
			return nil, err
		}
	}

	// Store the signing metadata for the CDA
	k.SetSigningData(ctx, id, msg.SigningData)

	// Return the id to the user
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		sdk.EventTypeMessage,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(types.AttributeKeyCdaId, strconv.FormatUint(id, 10)),
	))

	return &types.MsgCreateCdaResponse{Id: id}, nil
}

func (k msgServer) ApproveCda(goCtx context.Context, msg *types.MsgApproveCda) (*types.MsgApproveCdaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.SetApproval(ctx, msg)

	if err != nil {
		return nil, err
	}

	return &types.MsgApproveCdaResponse{}, nil
}

func (k msgServer) FinalizeCda(goCtx context.Context, msg *types.MsgFinalizeCda) (*types.MsgFinalizeCdaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.Finalize(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &types.MsgFinalizeCdaResponse{}, nil
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
	err := m.SetSigningDataSchema(ctx, msg.SigningDataSchema, id)
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
