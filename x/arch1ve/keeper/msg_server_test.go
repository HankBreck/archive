package keeper_test

import (
	"context"
	"testing"

	keepertest "arch1ve/testutil/keeper"
	"arch1ve/x/arch1ve/keeper"
	"arch1ve/x/arch1ve/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.Arch1veKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
