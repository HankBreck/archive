package keeper_test

import (
	"context"
	"testing"

	keepertest "archive/testutil/keeper"
	"archive/x/identity/keeper"
	"archive/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IdentityKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}