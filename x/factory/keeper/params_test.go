package keeper_test

import (
	"testing"

	testkeeper "github.com/umma-chain/core-umma/testutil/keeper"
	"github.com/umma-chain/core-umma/x/factory/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FactoryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
