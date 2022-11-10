package keeper

import (
	"archive/x/cda/types"
	"context"
	"strconv"

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
