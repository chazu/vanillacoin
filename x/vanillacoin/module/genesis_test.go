package vanillacoin_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "vanillacoin/testutil/keeper"
	"vanillacoin/testutil/nullify"
	"vanillacoin/x/vanillacoin/module"
	"vanillacoin/x/vanillacoin/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VanillacoinKeeper(t)
	vanillacoin.InitGenesis(ctx, k, genesisState)
	got := vanillacoin.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
