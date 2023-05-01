package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	crud "github.com/iov-one/cosmos-sdk-crud"
	"github.com/stretchr/testify/suite"

	"github.com/umma-chain/umma-core/x/escrow/keeper"
	"github.com/umma-chain/umma-core/x/escrow/test"
	"github.com/umma-chain/umma-core/x/escrow/types"
)

type BaseKeeperSuite struct {
	suite.Suite
	keeper       keeper.Keeper
	msgServer    types.MsgServer
	ctx          sdk.Context
	generator    *test.EscrowGenerator
	store        crud.Store
	storeKey     sdk.StoreKey
	balances     map[string]sdk.Coins
	configKeeper types.ConfigurationKeeper
}

func (s *BaseKeeperSuite) Setup(coinHolders []sdk.AccAddress, isModuleEnabled bool) {
	test.SetConfig()
	s.keeper, s.ctx, s.store, s.balances, s.storeKey, s.configKeeper = test.NewTestKeeper(coinHolders, isModuleEnabled)
	s.keeper.ImportNextID(s.ctx, 1)
	s.msgServer = keeper.NewMsgServerImpl(s.keeper)
	s.generator = test.NewEscrowGenerator(uint64(s.ctx.BlockTime().Unix()))
}
