package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/HankBreck/archive/testutil/keeper"
	"github.com/HankBreck/archive/x/identity/keeper"
	"github.com/HankBreck/archive/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IdentityKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
