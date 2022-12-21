package keeper

import (
	"archive/x/identity/types"
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	// Handle message and context
	if msg == nil {
		return nil, types.ErrInvalid.Wrap("Type MsgRegisterContract cannot be nil.")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Field validity check (this is likely a duplicate of ValidateBasic)
	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Create the issuer
	issuer := types.Issuer{
		Creator:     addr.String(),
		Name:        msg.Name,
		MoreInfoUri: msg.MoreInfoUri,
		Cost:        msg.Cost,
	}
	err = k.SetIssuer(ctx, issuer)
	if err != nil {
		return nil, err
	}

	// Emit Event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgRegisterIssuer,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(sdk.AttributeKeyAction, "RegisterIssuer"),
	))

	return &types.MsgRegisterIssuerResponse{}, nil
}

func (k Keeper) IssueCertificate(goCtx context.Context, msg *types.MsgIssueCertificate) (*types.MsgIssueCertificateResponse, error) {
	// Handle message and context
	if msg == nil {
		return nil, types.ErrInvalid.Wrap("Type MsgRegisterContract cannot be nil.")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Ensure msg.Creator is a registered Issuer (duplicate of ValidateBasic)
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	k.HasIssuer(ctx, creatorAddr.String())

	// Ensure msg.recipient is a valid account (duplicate of ValidateBasic)
	recipientAddr, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}
	if !k.accKeeper.HasAccount(ctx, recipientAddr) {
		return nil, sdkerrors.ErrNotFound.Wrapf("Recipient address does not exist")
	}

	// Create and store the certificate
	certificate := types.Certificate{
		// Id filled in AppendCertificate()
		IssuerAddress:     msg.Creator,
		Salt:              msg.Salt,
		MetadataSchemaUri: msg.MetadataSchemaUri,
		Hashes:            msg.Hashes,
	}
	id := k.AppendCertificate(ctx, certificate)

	// Add recipient to member store
	_ = msg.Recipient

	// emit events
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgIssueCertificate,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(sdk.AttributeKeyAction, "IssueCertificate"),
		sdk.NewAttribute("recipient", msg.Recipient),
		sdk.NewAttribute("certificate_id", strconv.FormatUint(id, 10)), // TODO: replace 0 with the cert ID
	))

	return nil, nil
}
