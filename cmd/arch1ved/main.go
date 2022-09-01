package main

import (
	"os"

	"archive/app"
	appparams "archive/app/params"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/ignite/cli/ignite/pkg/cosmoscmd"
)

func main() {
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		appparams.Bech32PrefixAccAddr,
		app.DefaultNodeHome,
		appparams.AppName,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)
	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
