package cli

import (
	"encoding/json"
	"strconv"

	"archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCreateCDA() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-cda [cid] [ownership JSON string] [expiration date UNIX timestamp (ms)]",
		Short: "Broadcast message createCDA",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argCid := args[0]
			argOwnership := args[1]
			argExpiration := args[2]

			// Unmarshal ownership from JSON string
			var ownership []*types.Ownership
			if err := json.Unmarshal([]byte(argOwnership), &ownership); err != nil {
				return err
			}

			// Parse uint64 from argument string
			expiration, err := strconv.ParseUint(argExpiration, 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCDA(
				clientCtx.GetFromAddress().String(),
				argCid,
				ownership,
				expiration,
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
