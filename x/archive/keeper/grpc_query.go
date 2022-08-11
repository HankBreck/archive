package keeper

import (
	"archive/x/archive/types"
)

var _ types.QueryServer = Keeper{}
