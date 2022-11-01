package keeper

import (
	"archive/x/contractregistry/types"
)

var _ types.QueryServer = Keeper{}
