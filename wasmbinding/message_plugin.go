package wasmbinding

import (
	"encoding/json"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	cdakeeper "github.com/HankBreck/archive/x/cda/keeper"
	cdatypes "github.com/HankBreck/archive/x/cda/types"

	"github.com/HankBreck/archive/wasmbinding/bindings"
)

var _ = bindings.ArchiveQuery{}

// CustomMessageDecorator returns decorator for custom CosmWasm bindings messages
func CustomMessageDecorator(cda *cdakeeper.Keeper) func(wasmkeeper.Messenger) wasmkeeper.Messenger {
	return func(old wasmkeeper.Messenger) wasmkeeper.Messenger {
		return &CustomMessenger{
			wrapped: old,
			cda:     cda,
		}
	}
}

type CustomMessenger struct {
	wrapped wasmkeeper.Messenger
	cda     *cdakeeper.Keeper
}

var _ wasmkeeper.Messenger = (*CustomMessenger)(nil)

// DispatchMsg executes on the contractMsg.
func (m *CustomMessenger) DispatchMsg(ctx sdk.Context, contractAddr sdk.AccAddress, contractIBCPortID string, msg wasmvmtypes.CosmosMsg) ([]sdk.Event, [][]byte, error) {
	if msg.Custom != nil {
		// only handle MsgWitnessApproveCda and MsgVoidCda
		// leave everything else for the wrapped version
		var contractMsg bindings.ArchiveMsg
		if err := json.Unmarshal(msg.Custom, &contractMsg); err != nil {
			return nil, nil, sdkerrors.Wrap(err, "archive msg")
		}
		if contractMsg.WitnessApproveCda != nil {
			return m.witnessApproveCda(ctx, contractAddr, contractMsg.WitnessApproveCda)
		}
		if contractMsg.FinalizeCda != nil {
			return m.finalizeCda(ctx, contractAddr, contractMsg.FinalizeCda)
		}
		if contractMsg.VoidCda != nil {
			return m.voidCda(ctx, contractAddr, contractMsg.VoidCda)
		}
	}
	return m.wrapped.DispatchMsg(ctx, contractAddr, contractIBCPortID, msg)
}

func (m *CustomMessenger) witnessApproveCda(ctx sdk.Context, contractAddr sdk.AccAddress, msg *bindings.WitnessApproveCda) ([]sdk.Event, [][]byte, error) {
	err := PerformWitnessApproveCda(m.cda, ctx, contractAddr, msg)
	if err != nil {
		return nil, nil, err
	}
	return nil, nil, nil
}

func PerformWitnessApproveCda(cdaKeeper *cdakeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, msg *bindings.WitnessApproveCda) error {
	if msg == nil {
		return wasmvmtypes.InvalidRequest{Err: "witnessApproveCda is nil"}
	}

	// Validate signing data
	var signingData cdatypes.RawSigningData
	err := signingData.UnmarshalJSON(msg.SigningData)
	if err != nil {
		return err
	}
	err = signingData.ValidateBasic()
	if err != nil {
		return err
	}

	// Validate the message
	msgWitnessApproveCda := cdatypes.NewMsgWitnessApproveCda(contractAddr.String(), msg.CdaId, signingData)
	err = msgWitnessApproveCda.ValidateBasic()
	if err != nil {
		return err
	}

	// Perform WitnessApproveCda through msgServer
	msgServer := cdakeeper.NewMsgServerImpl(*cdaKeeper)
	_, err = msgServer.WitnessApproveCda(sdk.WrapSDKContext(ctx), msgWitnessApproveCda)
	if err != nil {
		return err
	}

	return nil
}

func (m *CustomMessenger) finalizeCda(ctx sdk.Context, contractAddr sdk.AccAddress, msg *bindings.FinalizeCda) ([]sdk.Event, [][]byte, error) {
	err := PerformFinalizeCda(m.cda, ctx, contractAddr, msg)
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}

func PerformFinalizeCda(cdaKeeper *cdakeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, msg *bindings.FinalizeCda) error {
	return nil
}

func (m *CustomMessenger) voidCda(ctx sdk.Context, contractAddr sdk.AccAddress, msg *bindings.VoidCda) ([]sdk.Event, [][]byte, error) {
	err := PerformVoidCda(m.cda, ctx, contractAddr, msg)
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}

func PerformVoidCda(cdaKeeper *cdakeeper.Keeper, ctx sdk.Context, contractAddr sdk.AccAddress, msg *bindings.VoidCda) error {
	if msg == nil {
		return wasmvmtypes.InvalidRequest{Err: "voidCda is nil"}
	}

	// Build and validate the MsgVoidCda
	msgVoidCda := cdatypes.NewMsgVoidCda(contractAddr.String(), msg.CdaId)
	err := msgVoidCda.ValidateBasic()
	if err != nil {
		return err
	}

	// Perform the void through the msgServer
	msgServer := cdakeeper.NewMsgServerImpl(*cdaKeeper)
	// TODO: add event manager
	_, err = msgServer.VoidCda(sdk.WrapSDKContext(ctx), msgVoidCda)
	if err != nil {
		return err
	}

	return nil
}
