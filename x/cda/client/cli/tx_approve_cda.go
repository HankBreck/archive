package cli

import (
	"encoding/json"
	"strconv"

	"github.com/HankBreck/archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdApproveCda() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-cda [cda-id] [signer-id] [signing data stringified]",
		Short: "Broadcast message approve-cda",
		Args:  cobra.ExactArgs(3),
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

			// Signer ID
			signerId, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			// Signing Data
			var temp map[string]interface{}
			err = json.Unmarshal([]byte(args[2]), &temp)
			if err != nil {
				return err
			}
			bzSigningData, err := json.Marshal(temp)
			if err != nil {
				return err
			}
			var signingData types.RawSigningData
			err = signingData.UnmarshalJSON(bzSigningData)
			if err != nil {
				return err
			}

			msg := types.NewMsgApproveCda(
				clientCtx.GetFromAddress().String(),
				cdaId,
				signerId,
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
