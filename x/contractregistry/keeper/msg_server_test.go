package keeper_test

import (
	"context"
	"testing"

	keepertest "archive/testutil/keeper"
	"archive/x/contractregistry/keeper"
	"archive/x/contractregistry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: Test msgServer like this https://github.com/osmosis-labs/osmosis/blob/main/x/tokenfactory/keeper/msg_server_test.go

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ContractregistryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
