package cli

import (
	"strconv"

	"github.com/HankBreck/archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQueryContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contract [id]",
		Short: "Query contract",
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

			req := &types.QueryContractRequest{
				Id: id,
			}

			res, err := queryClient.Contract(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
