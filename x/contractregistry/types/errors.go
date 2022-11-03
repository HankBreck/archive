package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/contractregistry module sentinel errors
var (
	ErrEmpty                  = sdkerrors.Register(ModuleName, 2, "empty")
	ErrInvalid                = sdkerrors.Register(ModuleName, 3, "invalid")
	ErrNonExistentContract    = sdkerrors.Register(ModuleName, 4, "Invalid Contract ID. Please ensure the Contract exists for the given ID.")
	ErrNonExistentSigningData = sdkerrors.Register(ModuleName, 4, "Signing data was not found.")
)
