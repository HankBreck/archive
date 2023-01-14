package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/HankBreck/archive/testutil/keeper"
	"github.com/HankBreck/archive/x/cda/keeper"
	"github.com/HankBreck/archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CdaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

// TODO:
// Write tests for msg server functions
// Create CDA
//		Test signing metadata valid json but not schema match
