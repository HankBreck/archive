package cli

import (
	"strconv"

	"archive/x/cda/types"
	crtypes "archive/x/contractregistry/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdApproveCda() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-cda [cda-id] [signing data stringified]",
		Short: "Broadcast message approve-cda",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// CDA ID
			cdaId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			// Signing Data
			var signingData crtypes.RawSigningData
			signingData.UnmarshalJSON([]byte(args[1]))

			msg := types.NewMsgApproveCda(
				clientCtx.GetFromAddress().String(),
				cdaId,
				signingData,
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
