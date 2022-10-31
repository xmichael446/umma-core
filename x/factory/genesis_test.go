package factory_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/umma-chain/core-umma/testutil/keeper"
	"github.com/umma-chain/core-umma/testutil/nullify"
	"github.com/umma-chain/core-umma/x/factory"
	"github.com/umma-chain/core-umma/x/factory/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DenomList: []types.Denom{
			{
				Denom: "0",
			},
			{
				Denom: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FactoryKeeper(t)
	factory.InitGenesis(ctx, *k, genesisState)
	got := factory.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DenomList, got.DenomList)
	// this line is used by starport scaffolding # genesis/test/assert
}
