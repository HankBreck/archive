package keeper

import (
	"archive/x/identity/types"
)

var _ types.QueryServer = Keeper{}
