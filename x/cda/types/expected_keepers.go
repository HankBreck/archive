package types

import (
	crtypes "github.com/HankBreck/archive/x/contractregistry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ContractregistryKeeper interface {
	MatchesSigningDataSchema(ctx sdk.Context, targetContractId uint64, rawInputData crtypes.RawSigningData) (bool, error)
}
