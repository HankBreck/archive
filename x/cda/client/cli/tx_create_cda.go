package cli

import (
	"strconv"
	"time"

	"archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdktypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

const SIGNING_PARTIES = "signing-parties"

func CmdCreateCda() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-cda [legal contract ID] [legal metadata URI] [signing data type URI] [signing data stringified]",
		Short: "Broadcast message CreateCda",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Signing Parties
			signingParties, err := cmd.Flags().GetStringArray(SIGNING_PARTIES)
			if err != nil {
				return err
			}
			for _, addr := range signingParties {
				sdk.MustAccAddressFromBech32(addr)
			}

			// Contract ID
			contractId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			// Legal Metadata URI
			legalMetadataUri := args[1]

			// Signing Data
			signingData := &sdktypes.Any{
				TypeUrl: args[2],
				Value:   []byte(args[3]),
			}

			// Parse expiration time from argument string
			// TODO: Figure out how to limit this to UTC times
			utcExpireTime, err := time.Parse(time.RFC3339, args[3])
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateCda(
				clientCtx.GetFromAddress().String(),
				signingParties,
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

	// Signing parties flag
	cmd.Flags().StringArray(SIGNING_PARTIES, []string{}, "A list of account addresses that are signing parties in the CDA.")
	cmd.MarkFlagRequired(SIGNING_PARTIES)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
