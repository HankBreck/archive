package cli

import (
	"strconv"
	"strings"
	"time"

	"github.com/HankBreck/archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

const SIGNING_PARTIES = "signing-parties"

func CmdCreateCda() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-cda [signer ids] [legal contract ID] [legal metadata URI] [signing data stringified] [expiration time UTC]",
		Short: "Broadcast message CreateCda",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Signing Parties
			signerIdStrs := strings.Split(args[0], ",")
			signerIds := make([]uint64, len(signerIdStrs))
			for i, idStr := range signerIdStrs {
				id, err := strconv.ParseUint(idStr, 10, 64)
				if err != nil {
					return err
				}
				signerIds[i] = id
			}

			// Contract ID
			contractId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			// Legal Metadata URI
			legalMetadataUri := args[2]

			// Signing Data
			var signingData types.RawSigningData
			signingData.UnmarshalJSON([]byte(args[3]))

			// Parse expiration time from argument string
			// TODO: Figure out how to limit this to UTC times
			utcExpireTime, err := time.Parse(time.RFC3339, args[4])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCda(
				clientCtx.GetFromAddress().String(),
				signerIds,
				contractId,
				legalMetadataUri,
				signingData,
				utcExpireTime,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	return cmd
}
