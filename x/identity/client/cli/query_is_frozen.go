package cli

import (
	"strconv"

	"github.com/HankBreck/archive/x/identity/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdIsFrozen() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "is-frozen [id]",
		Short: "Query identity for frozen status",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			argId, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			req := types.QueryIsFrozenRequest{Id: argId}
			res, err := queryClient.IsFrozen(cmd.Context(), &req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
