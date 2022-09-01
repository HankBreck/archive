package main

import (
	"io"
	"os"

	"archive/app"
	appparams "archive/app/params"

	"github.com/cosmos/cosmos-sdk/baseapp"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/ignite/cli/ignite/pkg/cosmoscmd"
	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"
)

func main() {
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		appparams.Bech32PrefixAccAddr,
		app.DefaultNodeHome,
		appparams.AppName,
		app.ModuleBasics,
		cosmosCmdNewApp,
		// this line is used by starport scaffolding # root/arguments
	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}

func cosmosCmdNewApp(
	logger log.Logger,
	db dbm.DB,
	traceStore io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invCheckPeriod uint,
	encodingConfig cosmoscmd.EncodingConfig,
	appOpts servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) cosmoscmd.App {
	var apple cosmoscmd.App
	apple = app.New(
		logger,
		db,
		traceStore,
		loadLatest,
		skipUpgradeHeights,
		homePath,
		invCheckPeriod,
		encodingConfig,
		appOpts,
		baseAppOptions...,
	)
	return apple
}
