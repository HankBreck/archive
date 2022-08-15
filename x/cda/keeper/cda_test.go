package keeper_test

import (
	testkeeper "archive/testutil/keeper"
	"archive/x/cda/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCDACount(t *testing.T) {
	keeper, ctx := testkeeper.CdaKeeper(t)

	require.Equal(t, uint64(0), keeper.GetCDACount(ctx))
}

func TestSetCDACount(t *testing.T) {
	keeper, ctx := testkeeper.CdaKeeper(t)

	// Get current CDA count (0)
	oldCount := keeper.GetCDACount(ctx)

	// Set current CDA count to 1
	keeper.SetCDACount(ctx, oldCount+1)

	// Get current CDA count (1)
	newCount := keeper.GetCDACount(ctx)

	// Require first CDA count count != second CDA count
	require.Greater(t, newCount, oldCount)

}

func TestAppendCDA(t *testing.T) {
	keeper, ctx := testkeeper.CdaKeeper(t)

	// it should return the next available id
	cda := types.CDA{
		Creator: "valid address",
		Cid:     "ipfscid",
	}
	expected := keeper.GetCDACount(ctx)
	actual := keeper.AppendCDA(ctx, cda)
	require.Equal(t, expected, actual)

	// it should increment the CDA count
	higher := keeper.GetCDACount(ctx)
	require.Greater(t, higher, expected)
}
