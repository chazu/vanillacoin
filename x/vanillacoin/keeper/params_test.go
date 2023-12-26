package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "vanillacoin/testutil/keeper"
	"vanillacoin/x/vanillacoin/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.VanillacoinKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
