package keeper_test

import (
	"testing"

	testkeeper "archive/testutil/keeper"
	"archive/x/identity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// TODO:
//		Test IdentityMembers
//		Test Issuers

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.IdentityKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
