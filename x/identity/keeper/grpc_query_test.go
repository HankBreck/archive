package keeper_test

import (
	"testing"

	testkeeper "github.com/HankBreck/archive/testutil/keeper"
	"github.com/HankBreck/archive/x/identity/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// TODO:
//		Test IdentityMembers
//		Test Issuers
//		Test IssuerInfo
//		Test Identity
//		Test Operators
// 		Test MemberRole
// 		Test IsFrozen

func TestParamsQuery(t *testing.T) {
	keeper, ctx := testkeeper.IdentityKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
