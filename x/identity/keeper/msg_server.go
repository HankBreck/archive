package keeper

import (
	"archive/x/identity/types"
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
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

	// Field validity check (is duplicate of validatebasic?)
	addr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// Create the issuer
	issuer := types.Issuer{
		// ID set inside of k.AppendIssuer()
		Creator:     addr.String(),
		Name:        msg.Name,
		MoreInfoUri: msg.MoreInfoUri,
		Cost:        msg.Cost,
	}
	id := k.AppendIssuer(ctx, issuer)

	// Emit Event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.TypeMsgRegisterIssuer,
		sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		sdk.NewAttribute(sdk.AttributeKeyAction, "RegisterIssuer"),
		sdk.NewAttribute("issuer_id", strconv.FormatUint(id, 10)),
	))

	return &types.MsgRegisterIssuerResponse{Id: id}, nil
}
