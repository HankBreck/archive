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

	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Ensure the signing data matches the expected contract's schema
	matches, err := k.MatchesSigningDataSchema(ctx, msg.ContractId, msg.SigningData)
	if err != nil || !matches {
		return nil, types.ErrInvalidSigningData.Wrap("does not match schema")
	}

	// Fetch registered contract from storage
	contract, err := k.GetContract(ctx, msg.ContractId)
	if err != nil {
		return nil, err
	}

	// Setup witness contract if WitnessCodeId was set
	var witnessAddress sdk.AccAddress
	var witnessResponse []byte
	if contract.WitnessCodeId != 0 {
		// Instantiate the witness contract
		witnessAddress, witnessResponse, err = k.wasmKeeper.Instantiate(ctx, contract.WitnessCodeId, creatorAddr, creatorAddr, msg.WitnessInitMsg, "", sdk.Coins{})
		if err != nil {
			return nil, err
		}

		// Set signing data to witness ContractInfo
		err = k.wasmKeeper.SetContractInfoExtension(ctx, witnessAddress, &types.SigningDataExtension{SigningData: msg.SigningData})
		if err != nil {
			return nil, err
		}
		// Remove contract admin to ensure the witness cannot be corrupted
		err = k.wasmKeeper.ClearContractAdmin(ctx, witnessAddress, creatorAddr)
		if err != nil {
			return nil, err
		}
	}

	// Create the CDA
	cda := types.CDA{
		// Id set inside k.AppendCDA
		Creator:          msg.Creator,
		SignerIdentities: msg.SignerIds,
		ContractId:       msg.ContractId,
		LegalMetadataUri: msg.LegalMetadataUri,
		UtcExpireTime:    msg.UtcExpireTime,
		Status:           types.CDA_Pending,
		WitnessAddress:   witnessAddress.String(),
	}

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
		sdk.NewAttribute(types.AttributeKeyContractId, strconv.FormatUint(msg.ContractId, 10)),
	))

	return &types.MsgCreateCdaResponse{
		Id:                  id,
		WitnessAddress:      witnessAddress.String(),
		WitnessInitResponse: witnessResponse,
	}, nil
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

	// Emit Event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgApproveCda,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(types.AttributeKeyCdaId, strconv.FormatUint(cda.Id, 10)),
		sdk.NewAttribute("signer-id", strconv.FormatUint(msg.SignerId, 10)),
	))

	return &types.MsgApproveCdaResponse{}, nil
}

func (k msgServer) WitnessApproveCda(goCtx context.Context, msg *types.MsgWitnessApproveCda) (*types.MsgWitnessApproveCdaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Ensure CDA exists
	cda, err := k.GetCDA(ctx, msg.CdaId)
	if err != nil {
		return nil, err
	} else if cda == nil {
		return nil, types.ErrNonExistentCdaId.Wrapf("no CDA found for ID %d", msg.CdaId)
	}

	// Only allow approval when in the pending state
	if cda.Status != types.CDA_Pending {
		return nil, types.ErrInvalidCdaStatus.Wrap("The CDA must have a status of pending to be approved")
	}

	// Ensure the sender is the CDA's witness
	if cda.WitnessAddress != msg.Creator {
		return nil, types.ErrInvalid.Wrapf("sender (%s) must be the CDA's witness (%s)", msg.Creator, cda.WitnessAddress)
	}

	// Ensure signing data matches
	// TODO: is this necessary?
	metadata, err := k.GetSigningData(ctx, msg.CdaId)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(metadata.Bytes(), msg.SigningData.Bytes()) {
		return nil, types.ErrInvalidSigningData
	}

	// Set approval
	err = k.SetWitnessApproval(ctx, cda.Id)
	if err != nil {
		return nil, err
	}

	// Emit Event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgWitnessApproveCda,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(types.AttributeKeyCdaId, strconv.FormatUint(cda.Id, 10)),
		sdk.NewAttribute("witness-address", cda.WitnessAddress),
	))

	return &types.MsgWitnessApproveCdaResponse{}, nil
}

func (k msgServer) FinalizeCda(goCtx context.Context, msg *types.MsgFinalizeCda) (*types.MsgFinalizeCdaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

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

	// Ensure witness approval if required
	if cda.WitnessAddress != "" {
		if !k.HasWitnessApproval(ctx, cda.Id) {
			return nil, types.ErrMissingApproval.Wrapf("missing approval from witness (address: %s)", cda.WitnessAddress)
		}
	}

	// Update the CDA in storage
	cda.Status = types.CDA_Finalized
	err = k.UpdateCDA(ctx, cda.Id, cda)
	if err != nil {
		return nil, err
	}

	// Emit event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgFinalizeCda,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(types.AttributeKeyCdaId, strconv.FormatUint(cda.Id, 10)),
	))

	return &types.MsgFinalizeCdaResponse{}, nil
}

func (k msgServer) VoidCda(goCtx context.Context, msg *types.MsgVoidCda) (*types.MsgVoidCdaResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Ensure CDA exists
	cda, err := k.GetCDA(ctx, msg.CdaId)
	if err != nil {
		return nil, err
	} else if cda == nil {
		panic("CDA should never be nil")
	}

	if cda.WitnessAddress != msg.Creator {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("sender (%s) must be the witness (%s)", msg.Creator, cda.WitnessAddress)
	}

	// TODO: Should we prevent voiding for certain statuses?
	// TODO: Should we have a "reason" field under voids?
	// 		Could help with tracking if terms were violated

	// Ensure a CDA may only be voided once
	if cda.Status == types.CDA_Voided {
		return nil, types.ErrInvalidCdaStatus.Wrap("CDA already in voided status")
	}

	// Update in state
	cda.Status = types.CDA_Voided
	err = k.UpdateCDA(ctx, cda.Id, cda)
	if err != nil {
		return nil, err
	}

	// Emit event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgVoidCda,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(types.AttributeKeyCdaId, strconv.FormatUint(cda.Id, 10)),
	))

	return &types.MsgVoidCdaResponse{}, nil
}

func (k msgServer) RegisterContract(goCtx context.Context, msg *types.MsgRegisterContract) (*types.MsgRegisterContractResponse, error) {
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
	id := k.AppendContract(ctx, contract)

	// Store the schema in state
	err := k.SetSigningDataSchema(ctx, id, msg.SigningDataSchema)
	if err != nil {
		return nil, err
	}

	// Emit Event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgRegisterContract,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(types.AttributeKeyContractId, strconv.FormatUint(id, 10)),
	))

	return &types.MsgRegisterContractResponse{Id: id}, nil
}
