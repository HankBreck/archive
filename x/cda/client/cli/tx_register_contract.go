package cli

import (
	"fmt"
	"strconv"

	"github.com/HankBreck/archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

const AUTHORS = "authors"

func CmdRegisterContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-contract [description] [contact method (phone=0, email=1)] [contact value] [more info URI] [signing data schema stringified] [contract template URI] [contract template schema URI] [author 1] [author 2]...",
		Short: "Broadcast message RegisterContract",
		Args:  cobra.MinimumNArgs(7),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Description
			description := args[0]

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
			if err != nil {
				panic(err)
			}

			// Template URI
			templateUri := args[5]

			// Temlate Schema URI
			templateSchemaUri := args[6]

			// Authors
			authors := args[7:]
			for i, author := range authors {
				fmt.Println("Author", i, ":", author)
			}

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

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
