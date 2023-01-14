package cli

import (
	"strconv"

	"github.com/HankBreck/archive/x/contractregistry/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQuerySigningData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signing-data [id]",
		Short: "Query the signing data for a contract",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqId := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			// Parse id from arguments
			id, err := strconv.ParseUint(reqId, 10, 64)
			if err != nil {
				return err
			}

			req := &types.QuerySigningDataRequest{
				Id: id,
			}

			res, err := queryClient.SigningData(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
