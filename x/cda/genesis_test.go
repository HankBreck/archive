package cda_test

import (
	"testing"

	keepertest "archive/testutil/keeper"
	"archive/testutil/nullify"
	"archive/x/cda"
	"archive/x/cda/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CdaKeeper(t)
	cda.InitGenesis(ctx, *k, genesisState)
	got := cda.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
