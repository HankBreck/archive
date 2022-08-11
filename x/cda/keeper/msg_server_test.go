package keeper_test

import (
	"context"
	"testing"

	keepertest "arch1ve/testutil/keeper"
	"arch1ve/x/cda/keeper"
	"arch1ve/x/cda/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CdaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
