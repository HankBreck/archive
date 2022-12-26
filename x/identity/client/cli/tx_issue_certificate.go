package cli

import (
	"encoding/json"
	"strconv"

	"archive/x/identity/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdIssueCertificate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "issue-certificate [recipient] [salt] [metadata-schema-uri] [hashes]",
		Short: "Broadcast message issue-certificate",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Capture expected command line args
			argRecipient := args[0]
			argSalt := args[1]
			argMetadataSchemaUri := args[2]

			// Parse hashes from string input
			var argHashes map[string]string
			err = json.Unmarshal([]byte(args[3]), &argHashes)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgIssueCertificate(
				clientCtx.GetFromAddress().String(),
				argRecipient,
				argSalt,
				argMetadataSchemaUri,
				argHashes,
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
