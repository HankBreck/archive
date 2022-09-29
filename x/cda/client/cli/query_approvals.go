package cli

import (
	"strconv"

	"archive/x/cda/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdApprovals() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approvals [cda-id]",
		Short: "Responds with the account address of the owners that have approved the CDA",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqCdaId := args[0]

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryApprovalsRequest{

				CdaId: reqCdaId,
			}

			res, err := queryClient.Approvals(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
