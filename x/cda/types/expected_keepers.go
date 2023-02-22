package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type IdentityKeeper interface {
	HasMember(ctx sdk.Context, certificateId uint64, member sdk.AccAddress) (bool, error)
	HasCertificate(ctx sdk.Context, id uint64) bool
}
