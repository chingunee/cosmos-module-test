package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "module-test/testutil/keeper"
	"module-test/x/moduletest/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ModuletestKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
