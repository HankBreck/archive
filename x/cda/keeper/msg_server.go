package keeper

import (
	"bytes"
	"context"
	"strconv"

	"github.com/HankBreck/archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	cda := types.CDA{
		// Id set inside k.AppendCDA
		Creator:          msg.Creator,
		SignerIdentities: msg.SignerIds,
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

	// TODO: Check if the contract specifies a code_id
	//	if so,
	//		instantiate the contract
	//		set signing data field in the contract_info
	//		clear the contract admin

	// witnessAddress, err := k.wasmKeeper.instantiate()

	// Store CDA & grab cda id
	id := k.AppendCDA(ctx, cda)
	for i := range cda.SignerIdentities {
		signer := cda.SignerIdentities[i]
		// Ensure signer exists
		if !k.identityKeeper.HasCertificate(ctx, signer) {
			return nil, types.ErrIdentityNotFound.Wrapf("No identity registered for ID (%d)", signer)
		}
		k.AppendSignerCDA(ctx, signer, id)
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

	// Ensure CDA exists
	cda, err := k.GetCDA(ctx, msg.CdaId)
	if err != nil {
		return nil, err
	} else if cda == nil {
		return nil, types.ErrNonExistentCdaId.Wrapf("no CDA found for ID %d", msg.CdaId)
	}

	// Only allow approvals when in the pending state
	if cda.Status != types.CDA_Pending {
		return nil, types.ErrInvalidCdaStatus.Wrap("The CDA must have a status of pending to be approved")
	}

	// Ensure signing data matches
	metadata, err := k.GetSigningData(ctx, msg.CdaId)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(metadata.Bytes(), msg.SigningData.Bytes()) {
		return nil, types.ErrInvalidSigningData
	}

	// Ensure the sender is authorized to sign for identity
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	hasSender, err := k.identityKeeper.HasMember(ctx, msg.SignerId, creator)
	if err != nil {
		return nil, err
	} else if !hasSender {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("account (%s) is not authorized to sign for identity %d", creator.String(), msg.SignerId)
	}

	// Set approval
	err = k.SetApproval(ctx, msg.CdaId, msg.SignerId)
	if err != nil {
		return nil, err
	}

	// TODO: emit events

	return &types.MsgApproveCdaResponse{}, nil
}

func (k msgServer) FinalizeCda(goCtx context.Context, msg *types.MsgFinalizeCda) (*types.MsgFinalizeCdaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// validator creator address (duplicate of VB)
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Ensure CDA exists and is pending
	cda, err := k.GetCDA(ctx, msg.CdaId)
	if err != nil {
		return nil, err
	} else if cda.Status != types.CDA_Pending {
		return nil, types.ErrInvalidCdaStatus.Wrapf("CDA must have a status of pending")
	}

	// Ensure each signer has approved the CDA
	for _, signerId := range cda.SignerIdentities {
		if !k.HasApproval(ctx, cda.Id, signerId) {
			return nil, types.ErrMissingApproval.Wrapf("missing approval for signer ID %d", signerId)
		}
	}

	// Update the CDA in storage
	cda.Status = types.CDA_Finalized
	err = k.UpdateCDA(ctx, cda.Id, cda)
	if err != nil {
		// TODO: abstract away internal error messages
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
		WitnessCodeId:     msg.WitnessCodeId,
	}

	// Store the Contract in state
	id := m.AppendContract(ctx, contract)

	// Store the schema in state
	err := m.SetSigningDataSchema(ctx, id, msg.SigningDataSchema)
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
