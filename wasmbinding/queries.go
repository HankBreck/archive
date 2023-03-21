package wasmbinding

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/HankBreck/archive/wasmbinding/bindings"
	cdakeeper "github.com/HankBreck/archive/x/cda/keeper"
)

type QueryPlugin struct {
	cdaKeeper *cdakeeper.Keeper
}

// NewQueryPlugin returns a reference to a new QueryPlugin.
func NewQueryPlugin(cdaKeeper *cdakeeper.Keeper) *QueryPlugin {
	return &QueryPlugin{
		cdaKeeper: cdaKeeper,
	}
}

// GetSigningData is a query to get the signing data for a CDA.
func (qp QueryPlugin) GetSigningData(ctx sdk.Context, cdaId uint64) (*bindings.SigningDataResponse, error) {
	signingData, err := qp.cdaKeeper.GetSigningData(ctx, cdaId)
	if err != nil {
		return nil, fmt.Errorf("failed to get signing data for CDA: %d", cdaId)
	}

	return &bindings.SigningDataResponse{SigningData: signingData.Bytes()}, nil
}
