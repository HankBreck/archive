package keeper_test

import (
	"context"
	"testing"

	keepertest "archive/testutil/keeper"
	"archive/x/contractregistry/keeper"
	"archive/x/contractregistry/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ContractregistryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
