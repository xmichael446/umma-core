package simulation

import (
	"math/rand"

	"github.com/umma-chain/umma-core/x/factory/keeper"
	"github.com/umma-chain/umma-core/x/factory/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgBurnTokens(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgBurnTokens{
			Owner: simAccount.Address.String(),
		}

		// TODO: Handling the BurnTokens simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "BurnTokens simulation not implemented"), nil, nil
	}
}
