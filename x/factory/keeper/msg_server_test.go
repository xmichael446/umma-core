package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/umma-chain/core-umma/testutil/keeper"
	"github.com/umma-chain/core-umma/x/factory/keeper"
	"github.com/umma-chain/core-umma/x/factory/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FactoryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
