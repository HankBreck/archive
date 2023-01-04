package cli

import (
	"strconv"

	"archive/x/identity/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdOperators() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "operators [id]",
		Short: "Query operators",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryOperatorsRequest{

				Id: reqId,
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}
			params.Pagination = pageReq

			res, err := queryClient.Operators(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
