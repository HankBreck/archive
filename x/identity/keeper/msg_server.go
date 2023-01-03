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

	// Ensure msg.Creator is a registered Issuer
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Creator) // (duplicate of ValidateBasic)
	if err != nil {
		return nil, err
	}
	if !k.HasIssuer(ctx, creatorAddr.String()) {
		return nil, sdkerrors.ErrNotFound.Wrapf("Sender (%s) is not a registered Issuer", creatorAddr.String())
	}

	// Ensure msg.recipient is a valid account (duplicate of ValidateBasic)
	recipientAddr, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
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
	k.CreateMembership(ctx, id, recipientAddr)

	// Emit events
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgIssueCertificate,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(sdk.AttributeKeyAction, "IssueCertificate"),
		sdk.NewAttribute("recipient", msg.Recipient),
		sdk.NewAttribute("certificate_id", strconv.FormatUint(id, 10)),
	))

	return &types.MsgIssueCertificateResponse{Id: id}, nil
}

func (k msgServer) AcceptIdentity(goCtx context.Context, msg *types.MsgAcceptIdentity) (*types.MsgAcceptIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Ensure msg.Creator is a registered Issuer (duplicate of ValidateBasic)
	senderAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Add sender to the accepted membership list
	err = k.UpdateMembershipStatus(ctx, msg.Id, senderAddr, true)
	if err != nil {
		return nil, err
	}

	// Emit events
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgAcceptIdentity,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(sdk.AttributeKeyAction, "AcceptIdentity"),
		sdk.NewAttribute("certificate_id", strconv.FormatUint(msg.Id, 10)),
	))

	return &types.MsgAcceptIdentityResponse{}, nil
}

func (k msgServer) RejectIdentity(goCtx context.Context, msg *types.MsgRejectIdentity) (*types.MsgRejectIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Ensure msg.Creator is a registered Issuer (duplicate of ValidateBasic)
	senderAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Add sender to the accepted membership list
	err = k.UpdateMembershipStatus(ctx, msg.Id, senderAddr, false)
	if err != nil {
		return nil, err
	}

	// Emit events
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgRejectIdentity,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(sdk.AttributeKeyAction, "RejectIdentity"),
		sdk.NewAttribute("certificate_id", strconv.FormatUint(msg.Id, 10)),
	))

	return &types.MsgRejectIdentityResponse{}, nil
}

func (k msgServer) RevokeIdentity(goCtx context.Context, msg *types.MsgRevokeIdentity) (*types.MsgRevokeIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Ensure msg.Creator is a registered Issuer (duplicate of ValidateBasic)
	senderAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Ensure Id is valid & msg.Creator issued the Certificate
	memberAddr, err := sdk.AccAddressFromBech32(msg.Member)
	if err != nil {
		return nil, err
	}

	// Ensure issuer created the certificate
	found, err := k.HasIssuerForId(ctx, msg.Id, senderAddr)
	if err != nil {
		return nil, err
	}
	if !found {
		sdkerrors.ErrUnauthorized.Wrapf("Sender is not an issuer for id %d", msg.Id)
	}

	// Remove member from both pending and accepted lists
	toAdd := []sdk.AccAddress{}
	toRemove := []sdk.AccAddress{memberAddr}
	err = k.UpdateAcceptedMembers(ctx, msg.Id, toAdd, toRemove)
	if err != nil {
		return nil, err
	}
	err = k.UpdatePendingMembers(ctx, msg.Id, toAdd, toRemove)
	if err != nil {
		return nil, err
	}

	// Emit Events
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgRevokeIdentity,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(sdk.AttributeKeyAction, "RevokeIdentity"),
		sdk.NewAttribute("certificate_id", strconv.FormatUint(msg.Id, 10)),
		sdk.NewAttribute("member", memberAddr.String()),
	))

	return &types.MsgRevokeIdentityResponse{}, nil
}

func (k msgServer) RenounceIdentity(goCtx context.Context, msg *types.MsgRenounceIdentity) (*types.MsgRenounceIdentityResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Ensure the certificate exists
	if !k.HasCertificate(ctx, msg.Id) {
		return nil, types.ErrNonexistentCertificate.Wrapf("No identity found for ID %d", msg.Id)
	}

	// Ensure sender is a member of the identity
	if !k.HasMember(ctx, msg.Id, addr) {
		return nil, sdkerrors.ErrNotFound.Wrapf("The sender is not a member identity %d", msg.Id)
	}

	// Remove from pending and accepted membership lists
	toAdd := []sdk.AccAddress{}
	toRemove := []sdk.AccAddress{addr}
	err = k.UpdateAcceptedMembers(ctx, msg.Id, toAdd, toRemove)
	if err != nil {
		return nil, err
	}
	err = k.UpdatePendingMembers(ctx, msg.Id, toAdd, toRemove)
	if err != nil {
		return nil, err
	}

	// Emit events
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgRenounceIdentity,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(sdk.AttributeKeyAction, "RenounceIdentity"),
		sdk.NewAttribute("certificate_id", strconv.FormatUint(msg.Id, 10)),
	))

	return &types.MsgRenounceIdentityResponse{}, nil
}

func (k msgServer) AddIdentityMember(goCtx context.Context, msg *types.MsgAddIdentityMember) (*types.MsgAddIdentityMemberResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	senderAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}
	// TODO: constrain this to operators
	if !k.HasMember(ctx, msg.Id, senderAddr) {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("Sender must be an operator of identity %d", msg.Id)
	}

	// Cannot overwrite existing members
	// Operator status could be revoked by saving the value 0x0 to the membership key instead of 0x1
	memberAddr, err := sdk.AccAddressFromBech32(msg.Member)
	if err != nil {
		return nil, err
	}
	if k.HasMember(ctx, msg.Id, memberAddr) || k.HasPendingMember(ctx, msg.Id, memberAddr) {
		return nil, types.ErrExistingMember
	}

	// Add member as pending
	toAdd := []sdk.AccAddress{memberAddr}
	toRemove := []sdk.AccAddress{}
	err = k.UpdatePendingMembers(ctx, msg.Id, toAdd, toRemove)
	if err != nil {
		return nil, err
	}

	// Emit events
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgAddIdentityMember,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(sdk.AttributeKeyAction, "AddIdentityMember"),
		sdk.NewAttribute("certificate_id", strconv.FormatUint(msg.Id, 10)),
		sdk.NewAttribute("member_address", memberAddr.String()),
		// sdk.NewAttribute("role", "operator" OR "member"),
	))

	return &types.MsgAddIdentityMemberResponse{}, nil
}
