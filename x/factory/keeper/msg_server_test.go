package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/umma-chain/umma-core/testutil/keeper"
	"github.com/umma-chain/umma-core/x/factory/keeper"
	"github.com/umma-chain/umma-core/x/factory/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.FactoryKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
