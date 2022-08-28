package keeper_test

import (
	testkeeper "archive/testutil/keeper"
	sample "archive/testutil/sample"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetOwnerCDACount(t *testing.T) {
	keeper, ctx := testkeeper.CdaKeeper(t)
	account, _ := sdk.AccAddressFromBech32(sample.AccAddress())

	require.Equal(t, uint64(0), keeper.GetOwnerCDACount(ctx, account))
}

func TestSetOwnerCDACount(t *testing.T) {
	keeper, ctx := testkeeper.CdaKeeper(t)
	account, _ := sdk.AccAddressFromBech32(sample.AccAddress())

	// Get current CDA count (0)
	oldCount := keeper.GetOwnerCDACount(ctx, account)

	// Set current CDA count to 1
	keeper.SetOwnerCDACount(ctx, account, oldCount+1)

	// Get current CDA count (1)
	newCount := keeper.GetOwnerCDACount(ctx, account)

	// Require first CDA count count != second CDA count
	require.Greater(t, newCount, oldCount)

}

func TestAppendOwnerCDA(t *testing.T) {
	keeper, ctx := testkeeper.CdaKeeper(t)
	account1, _ := sdk.AccAddressFromBech32(sample.AccAddress())

	err := keeper.AppendOwnerCDA(ctx, "invalid account", 0)
	require.Error(t, err)

	// it should return the next available id
	initial := keeper.GetOwnerCDACount(ctx, account1)
	keeper.AppendOwnerCDA(ctx, account1.String(), 0)
	final := keeper.GetOwnerCDACount(ctx, account1)

	require.Greater(t, final, initial)
}
