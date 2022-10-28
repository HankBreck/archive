package keeper_test

import (
	"context"
	"testing"

	keepertest "archive/testutil/keeper"
	"archive/x/cda/keeper"
	"archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CdaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

// TODO:
// Write tests for msg server functions
