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
		Use:   "cda [cda-id] [id] [cid]",
		Short: "Query cda",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqCdaId := args[0]
			reqId := args[1]
			reqCid := args[2]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryCdaRequest{

				CdaId: reqCdaId,
				Id:    reqId,
				Cid:   reqCid,
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
