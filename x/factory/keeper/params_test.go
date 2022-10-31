package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/umma-chain/umma-core/testutil/keeper"
	"github.com/umma-chain/umma-core/x/factory/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FactoryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
