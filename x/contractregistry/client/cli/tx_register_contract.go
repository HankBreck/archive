package cli

import (
	"encoding/json"
	"fmt"
	"strconv"

	"archive/x/contractregistry/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

const AUTHORS = "authors"

func CmdRegisterContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-contract [description] [contact method (phone=0, email=1)] [contact value] [more info URI] [signing data schema stringified] [contract template URI] [contract template schema URI] --authors [author 1],[author 2]",
		Short: "Broadcast message RegisterContract",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Description
			description := args[0]

			// Authors
			authors, err := cmd.Flags().GetStringArray(AUTHORS)
			if err != nil {
				return err
			}

			// Contact Info
			method, err := strconv.ParseInt(args[1], 10, 32)
			if err != nil {
				panic(err)
			}
			contactInfo := types.ContactInfo{
				Method: types.ContactMethod(method),
				Value:  args[2],
			}

			// More Info URI
			moreInfoUri := args[3]

			// Signing Data Schema
			bzSchema := []byte(args[4])
			fmt.Println("Valid JSON?", json.Valid(bzSchema))
			if err != nil {
				panic(err)
			}

			// Template URI
			templateUri := args[5]

			// Temlate Schema URI
			templateSchemaUri := args[6]

			msg := types.NewMsgRegisterContract(
				clientCtx.GetFromAddress().String(),
				description,
				authors,
				&contactInfo,
				moreInfoUri,
				bzSchema,
				templateUri,
				templateSchemaUri,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	// Signing parties flag
	cmd.Flags().StringArray(AUTHORS, []string{}, "A list the contract authors.")
	cmd.MarkFlagRequired(AUTHORS)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
