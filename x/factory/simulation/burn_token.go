package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/umma-chain/umma-core/x/factory/keeper"
	"github.com/umma-chain/umma-core/x/factory/types"
)

func SimulateMsgBurnToken(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgBurnToken{
			Owner: simAccount.Address.String(),
		}

		// TODO: Handling the BurnToken simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "BurnToken simulation not implemented"), nil, nil
	}
}
