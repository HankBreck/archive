package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/HankBreck/archive/x/cda/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group cda queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(CmdCda())
	cmd.AddCommand(CmdCdas())
	cmd.AddCommand(CmdCdasBySigner())
	cmd.AddCommand(CmdApproval())
	cmd.AddCommand(CmdQueryContract())
	cmd.AddCommand(CmdQueryContracts())
	cmd.AddCommand(CmdQuerySigningData())
	cmd.AddCommand(CmdQuerySigningDataSchema())

	return cmd
}
