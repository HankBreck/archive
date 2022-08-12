package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cda module sentinel errors
var (
	ErrInvalidCid = sdkerrors.Register(ModuleName, 111, "Cid must be a valid string")
)
