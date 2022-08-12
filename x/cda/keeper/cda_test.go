package keeper_test

import (
	testkeeper "archive/testutil/keeper"
	"archive/x/cda/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCDACount(t *testing.T) {
	// keeper, ctx := testkeeper.CdaKeeper(t)

	//
}

func TestSetCDACount(t *testing.T) {
	// TODO: test the SetCDACount functionality
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
