package keeper_test

import (
	"testing"

	testkeeper "github.com/HankBreck/archive/testutil/keeper"
	sample "github.com/HankBreck/archive/testutil/sample"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetOwnerCDACount(t *testing.T) {
	keeper, ctx := testkeeper.CdaKeeper(t)
	account, _ := sdk.AccAddressFromBech32(sample.AccAddress())

	require.Equal(t, uint64(0), keeper.GetOwnerCDACount(ctx, account.String()))
}

func TestSetOwnerCDACount(t *testing.T) {
	keeper, ctx := testkeeper.CdaKeeper(t)
	account, _ := sdk.AccAddressFromBech32(sample.AccAddress())

	// Get current CDA count (0)
	oldCount := keeper.GetOwnerCDACount(ctx, account.String())

	// Set current CDA count to 1
	keeper.SetOwnerCDACount(ctx, account.String(), oldCount+1)

	// Get current CDA count (1)
	newCount := keeper.GetOwnerCDACount(ctx, account.String())

	// Require first CDA count count != second CDA count
	require.Greater(t, newCount, oldCount)

}

func TestAppendOwnerCDA(t *testing.T) {
	keeper, ctx := testkeeper.CdaKeeper(t)
	account1, _ := sdk.AccAddressFromBech32(sample.AccAddress())

	err := keeper.AppendOwnerCDA(ctx, "invalid account", 0)
	require.Error(t, err)

	// it should return the next available id
	initial := keeper.GetOwnerCDACount(ctx, account1.String())
	keeper.AppendOwnerCDA(ctx, account1.String(), 0)
	final := keeper.GetOwnerCDACount(ctx, account1.String())

	require.Greater(t, final, initial)
}