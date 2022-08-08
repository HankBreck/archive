package arch1ve_test

import (
	"testing"

	keepertest "arch1ve/testutil/keeper"
	"arch1ve/testutil/nullify"
	"arch1ve/x/arch1ve"
	"arch1ve/x/arch1ve/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.Arch1veKeeper(t)
	arch1ve.InitGenesis(ctx, *k, genesisState)
	got := arch1ve.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
