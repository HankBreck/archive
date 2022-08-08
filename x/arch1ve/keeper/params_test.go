package keeper_test

import (
	"testing"

	testkeeper "arch1ve/testutil/keeper"
	"arch1ve/x/arch1ve/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.Arch1veKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
