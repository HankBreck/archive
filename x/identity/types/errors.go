package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/identity module sentinel errors
var (
	ErrEmpty                  = sdkerrors.Register(ModuleName, 2, "empty")
	ErrInvalid                = sdkerrors.Register(ModuleName, 3, "invalid")
	ErrExistingIssuer         = sdkerrors.Register(ModuleName, 4, "an issuer already exists for this address")
	ErrExistingMember         = sdkerrors.Register(ModuleName, 5, "a member already exists for this id, address combination")
	ErrNonexistentCertificate = sdkerrors.Register(ModuleName, 6, "the specified certificate does not exist")
)
