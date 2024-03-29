package cli

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/HankBreck/archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

const SIGNING_PARTIES = "signing-parties"

func CmdCreateCda() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-cda [signer ids] [legal contract ID] [legal metadata URI] [signing data stringified] [expiration time UTC] [json-encoded witness init msg]",
		Short: "Broadcast message CreateCda",
		Args:  cobra.ExactArgs(6),
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
			var temp map[string]interface{}
			err = json.Unmarshal([]byte(args[3]), &temp)
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

			// Parse expiration time from argument string
			// TODO: Figure out how to limit this to UTC times
			utcExpireTime, err := time.Parse(time.RFC3339, args[4])
			if err != nil {
				return err
			}

			// Witness Initialization Message
			witnessInitMsg := []byte(args[5])

			msg := types.NewMsgCreateCda(
				clientCtx.GetFromAddress().String(),
				signerIds,
				contractId,
				legalMetadataUri,
				signingData,
				utcExpireTime,
				witnessInitMsg,
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
