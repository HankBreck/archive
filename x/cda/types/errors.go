package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cda module sentinel errors
var (
	ErrInvalidCid        = sdkerrors.Register(ModuleName, 111, "Cid must be a valid string")
	ErrInvalidOwnership  = sdkerrors.Register(ModuleName, 112, "Invalid ownership map")
	ErrInvalidExpiration = sdkerrors.Register(ModuleName, 113, "Invalid value for expiration. Must be a valid UTC millisecond timestamp.")
	ErrExistingApproval  = sdkerrors.Register(ModuleName, 114, "The address has already given approval for this CDA.")
	ErrNonExistentCdaId  = sdkerrors.Register(ModuleName, 115, "Invalid CdaId. Please ensure the CDA exists for the given ID.")
)
