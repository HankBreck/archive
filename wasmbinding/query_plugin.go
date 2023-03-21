package wasmbinding

import (
	"encoding/json"

	"github.com/HankBreck/archive/wasmbinding/bindings"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Add stargate querier
// https://github.com/osmosis-labs/osmosis/blob/main/wasmbinding/query_plugin.go#L18

// CustomQuerier dispatches custom CosmWasm bindings queries.
func CustomQuerier(qp *QueryPlugin) func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var contractQuery bindings.ArchiveQuery
		if err := json.Unmarshal(request, &contractQuery); err != nil {
			return nil, sdkerrors.Wrap(err, "archive query")
		}

		switch {
		case contractQuery.SigningData != nil:

			signingDataRes, err := qp.GetSigningData(ctx, contractQuery.SigningData.CdaId)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "archive signing data query")
			}

			res := bindings.SigningDataResponse{
				SigningData: signingDataRes.SigningData,
			}

			bz, err := json.Marshal(res)
			if err != nil {
				return nil, sdkerrors.Wrap(err, "archive signing data query response")
			}

			return bz, nil

		default:
			return nil, wasmvmtypes.UnsupportedRequest{Kind: "unknown archive query variant"}
		}
	}
}
