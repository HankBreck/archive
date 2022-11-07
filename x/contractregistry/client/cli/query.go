package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"archive/x/contractregistry/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group contractregistry queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdQueryContract())
	cmd.AddCommand(CmdQueryContracts())
	// this line is used by starport scaffolding # 1

	return cmd
}
