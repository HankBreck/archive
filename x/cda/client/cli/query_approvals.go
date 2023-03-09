package cli

import (
	"strconv"

	"github.com/HankBreck/archive/x/cda/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdApproval() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approval [cda-id] [signer id]",
		Short: "Responds with true if the owner has approved the CDA",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			cdaId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			signerId, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryApprovalRequest{
				CdaId:    cdaId,
				SignerId: signerId,
			}

			res, err := queryClient.Approval(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
