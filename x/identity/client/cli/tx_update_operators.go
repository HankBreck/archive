package cli

import (
	"strconv"

	"strings"

	"github.com/HankBreck/archive/x/identity/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdUpdateOperators() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-operators [id] [to-add] [to-remove]",
		Short: "Broadcast message update-operators",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			// Parse addresses from user input
			toAdd := []string{}
			argToAdd := strings.Split(args[1], listSeparator)
			for _, addr := range argToAdd {
				if strings.TrimSpace(addr) != "" {
					toAdd = append(toAdd, addr)
				}
			}
			toRemove := []string{}
			argToRemove := strings.Split(args[2], listSeparator)
			for _, addr := range argToRemove {
				if strings.TrimSpace(addr) != "" {
					toRemove = append(toRemove, addr)
				}
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateOperators(
				clientCtx.GetFromAddress().String(),
				argId,
				toAdd,
				toRemove,
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
