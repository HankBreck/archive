package cli

import (
	"strconv"

	"archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCda() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cda [id]",
		Short: "Query cda",
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

			params := &types.QueryCdaRequest{
				Id: id,
			}

			res, err := queryClient.Cda(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
