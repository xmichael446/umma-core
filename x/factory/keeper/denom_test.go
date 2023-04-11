package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "github.com/umma-chain/umma-core/testutil/keeper"
	"github.com/umma-chain/umma-core/testutil/nullify"
	"github.com/umma-chain/umma-core/x/factory/keeper"
	"github.com/umma-chain/umma-core/x/factory/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDenom(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Denom {
	items := make([]types.Denom, n)
	for i := range items {
		items[i].Denom = strconv.Itoa(i)

		keeper.SetDenom(ctx, items[i])
	}
	return items
}

func TestDenomGet(t *testing.T) {
	keeper, ctx := keepertest.FactoryKeeper(t)
	items := createNDenom(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDenom(ctx,
			item.Denom,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

//func TestDenomRemove(t *testing.T) {
//	keeper, ctx := keepertest.FactoryKeeper(t)
//	items := createNDenom(keeper, ctx, 10)
//	for _, item := range items {
//		keeper.RemoveDenom(ctx,
//			item.Denom,
//		)
//		_, found := keeper.GetDenom(ctx,
//			item.Denom,
//		)
//		require.False(t, found)
//	}
//}

func TestDenomGetAll(t *testing.T) {
	keeper, ctx := keepertest.FactoryKeeper(t)
	items := createNDenom(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDenom(ctx)),
	)
}
