package cli

import (
	"strconv"

	"archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdCdas() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cdas",
		Short: "Query CDAs",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {

			// TODO: Figure out how to optionally parse flags

			// nextKey, err := cmd.Flags().GetString("next-key")
			// if err != nil {
			// 	return err
			// }

			// limit, err := cmd.Flags().GetUint64("limit")
			// if err != nil {
			// 	return err
			// }

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCdasRequest{
				// Pagination: &query.PageRequest{
				// 	Key:   []byte(nextKey),
				// 	Limit: limit,
				// },
			}

			res, err := queryClient.Cdas(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
