package cli

import (
	"bytes"
	"encoding/base64"
	"strconv"

	"archive/x/contractregistry/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQueryContracts() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "contracts",
		Short: "Query contracts",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			err = cmd.ParseFlags(args)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			nextKey := cmd.Flag("nextKey")
			limit := cmd.Flag("limit")
			reverse := cmd.Flag("reverse")

			// Parse user input
			nextKeyVal, err := base64.StdEncoding.DecodeString(nextKey.Value.String())
			if err != nil {
				return err
			}
			limitVal, err := strconv.ParseUint(limit.Value.String(), 10, 64)
			if err != nil {
				return err
			}
			reverseVal := reverse.Changed

			pagination := &query.PageRequest{
				Reverse: reverseVal,
			}

			// Build pagination
			if !bytes.Equal(nextKeyVal, []byte("")) {
				pagination = &query.PageRequest{
					Reverse: pagination.Reverse,
					Key:     []byte(nextKeyVal),
				}
			}
			if limitVal > uint64(0) {
				if pagination.Key != nil {
					pagination = &query.PageRequest{
						Reverse: pagination.Reverse,
						Key:     pagination.Key,
						Limit:   limitVal,
					}
				} else {
					pagination = &query.PageRequest{
						Reverse: pagination.Reverse,
						Limit:   limitVal,
					}
				}
			}

			req := &types.QueryContractsRequest{Pagination: pagination}
			res, err := queryClient.Contracts(cmd.Context(), req)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	// Pagination flags
	cmd.Flags().BytesBase64("nextKey", []byte(""), "the next key string used in pagination")
	cmd.Flags().Uint64("limit", uint64(0), "the maximum number of contracts fetched")
	cmd.Flags().Bool("reverse", false, "if included, the query will be performed in the reverse order")

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
