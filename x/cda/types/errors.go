package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cda module sentinel errors
var (
	ErrEmpty                  = sdkerrors.Register(ModuleName, 2, "empty")
	ErrInvalid                = sdkerrors.Register(ModuleName, 3, "invalid")
	ErrNonExistentContract    = sdkerrors.Register(ModuleName, 4, "Invalid Contract ID. Please ensure the Contract exists for the given ID.")
	ErrNonExistentSigningData = sdkerrors.Register(ModuleName, 5, "Signing data was not found.")
	ErrExistingEntry          = sdkerrors.Register(ModuleName, 6, "Existing entry found.")

	ErrInvalidExpiration  = sdkerrors.Register(ModuleName, 113, "Invalid value for expiration. Must be a valid UTC millisecond timestamp.")
	ErrExistingApproval   = sdkerrors.Register(ModuleName, 114, "The address has already given approval for this CDA.")
	ErrNonExistentCdaId   = sdkerrors.Register(ModuleName, 115, "Invalid CdaId. Please ensure the CDA exists for the given ID.")
	ErrMissingApproval    = sdkerrors.Register(ModuleName, 116, "Missing CDA approvals.")
	ErrAlreadyFinalized   = sdkerrors.Register(ModuleName, 117, "CDA has already been finalized.")
	ErrInvalidSigningData = sdkerrors.Register(ModuleName, 118, "Signing data provided does not match the signing data stored in the CDA.")
	ErrInvalidCdaStatus   = sdkerrors.Register(ModuleName, 119, "The CDA's status did not match the expected status.")
)
