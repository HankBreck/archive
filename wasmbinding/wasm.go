package wasmbinding

import (
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	cdakeeper "github.com/HankBreck/archive/x/cda/keeper"
)

func RegisterCustomPlugins(cdaKeeper *cdakeeper.Keeper) []wasmkeeper.Option {
	wasmQueryPlugin := NewQueryPlugin(cdaKeeper)

	queryPluginOpt := wasmkeeper.WithQueryPlugins(&wasmkeeper.QueryPlugins{
		Custom: CustomQuerier(wasmQueryPlugin),
	})
	messageDecoratorOpt := wasmkeeper.WithMessageHandlerDecorator(
		CustomMessageDecorator(cdaKeeper),
	)

	return []wasmkeeper.Option{
		queryPluginOpt,
		messageDecoratorOpt,
	}
}
