package keeper

import (
	"archive/x/identity/types"
	"context"
	"fmt"
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

	// Add recipient to operator store
	k.SetInitialOperator(ctx, id, recipientAddr)

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
	senderAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Remove sender from the pending membership list
	err = k.UpdateMembershipStatus(ctx, msg.Id, senderAddr, false)
	if err != nil {
		return nil, err
	}

	// Remove operator status if sender was the initial recipient
	hasOp, err := k.HasOperator(ctx, msg.Id, senderAddr)
	if err != nil {
		return nil, err
	} else if hasOp {
		k.RemoveOperators(ctx, msg.Id, []sdk.AccAddress{senderAddr})
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
	hasAccepted, err := k.HasMember(ctx, msg.Id, addr)
	if err != nil {
		return nil, err
	} else if !hasAccepted {
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

func (k msgServer) UpdateMembers(goCtx context.Context, msg *types.MsgUpdateMembers) (*types.MsgUpdateMembersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	senderAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Ensure sender is an operator/issuer (and certificate exists)
	valid, err := k.hasOperatorOrIssuer(ctx, msg.Id, senderAddr)
	if err != nil {
		return nil, err
	} else if !valid {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("account (%s) is not an operator of identity (%d)", senderAddr.String(), msg.Id)
	}

	// Convert msg.ToAdd from []string to []sdk.AccAddress
	toAdd := []sdk.AccAddress{}
	for _, addrStr := range msg.ToAdd {
		addr, err := sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return nil, err
		}
		toAdd = append(toAdd, addr)
	}
	// Convert msg.ToRemove from []string to []sdk.AccAddress
	toRemove := []sdk.AccAddress{}
	for _, addrStr := range msg.ToRemove {
		addr, err := sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return nil, err
		}

		toRemove = append(toRemove, addr)
	}

	// Update pending members
	err = k.UpdatePendingMembers(ctx, msg.Id, toAdd, toRemove)
	if err != nil {
		return nil, err
	}
	// Update accepted members (do not directly add to AcceptedMembers)
	err = k.UpdateAcceptedMembers(ctx, msg.Id, []sdk.AccAddress{}, toRemove)
	if err != nil {
		return nil, err
	}

	// Emit Events
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgUpdateMembers,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(sdk.AttributeKeyAction, "UpdateMembers"),
		sdk.NewAttribute("certificate_id", strconv.FormatUint(msg.Id, 10)),
		sdk.NewAttribute("num_added", fmt.Sprint(len(toAdd))),
		sdk.NewAttribute("num_removed", fmt.Sprint(len(toRemove))),
	))

	return &types.MsgUpdateMembersResponse{}, nil
}

func (k msgServer) UpdateOperators(goCtx context.Context, msg *types.MsgUpdateOperators) (*types.MsgUpdateOperatorsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	senderAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Ensure sender is an operator/issuer (and certificate exists)
	valid, err := k.hasOperatorOrIssuer(ctx, msg.Id, senderAddr)
	if err != nil {
		return nil, err
	} else if !valid {
		return nil, sdkerrors.ErrUnauthorized.Wrapf("account (%s) is not an operator of identity (%d)", senderAddr.String(), msg.Id)
	}

	// Add new operators
	toAdd := []sdk.AccAddress{}
	for _, addrStr := range msg.ToAdd {
		addr, err := sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return nil, err
		}
		toAdd = append(toAdd, addr)
	}
	err = k.SetOperators(ctx, msg.Id, toAdd)
	if err != nil {
		return nil, err
	}

	// Remove old operators
	toRemove := []sdk.AccAddress{}
	for _, addrStr := range msg.ToRemove {
		addr, err := sdk.AccAddressFromBech32(addrStr)
		if err != nil {
			return nil, err
		}
		toRemove = append(toRemove, addr)
	}
	err = k.RemoveOperators(ctx, msg.Id, toRemove)
	if err != nil {
		return nil, err
	}

	// Emit events
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgUpdateOperators,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(sdk.AttributeKeyAction, "UpdateOperators"),
		sdk.NewAttribute("certificate_id", strconv.FormatUint(msg.Id, 10)),
		sdk.NewAttribute("num_added", fmt.Sprint(len(toAdd))),
		sdk.NewAttribute("num_removed", fmt.Sprint(len(toRemove))),
	))

	return &types.MsgUpdateOperatorsResponse{}, nil
}
