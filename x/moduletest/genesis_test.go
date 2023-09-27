package moduletest_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "module-test/testutil/keeper"
	"module-test/testutil/nullify"
	"module-test/x/moduletest"
	"module-test/x/moduletest/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ModuletestKeeper(t)
	moduletest.InitGenesis(ctx, *k, genesisState)
	got := moduletest.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
