package cli

import (
	"encoding/json"
	"strconv"

	"archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdApproveCda() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-cda [cda-id] [ownership]",
		Short: "Broadcast message approve-cda",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCdaId := args[0]
			argOwnership := args[1]

			// Extract CdaId from string
			cdaId, err := cast.ToUint64E(argCdaId)
			if err != nil {
				return err
			}

			// Unmarshal ownership from JSON string
			var ownership []*types.Ownership
			if err := json.Unmarshal([]byte(argOwnership), &ownership); err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgApproveCda(
				clientCtx.GetFromAddress().String(),
				cdaId,
				ownership,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
