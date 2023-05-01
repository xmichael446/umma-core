package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/umma-chain/umma-core/x/configuration"
	"github.com/umma-chain/umma-core/x/escrow/types"
)

// FIXME: this module should not have a dependency on the configuration module

// CollectFees collect the fees for the given message
func (k *Keeper) CollectFees(ctx sdk.Context, msg types.MsgWithFeePayer) error {
	fees := k.ComputeFees(ctx, msg)
	return k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		msg.GetFeePayer(),
		authtypes.FeeCollectorName,
		fees,
	)
}

func (k *Keeper) ComputeFees(ctx sdk.Context, msg sdk.Msg) sdk.Coins {
	feesConfiguration := k.configurationKeeper.GetFees(ctx)

	defaultFee := feesConfiguration.FeeDefault

	specificFee := getFee(feesConfiguration, msg)

	if specificFee.LT(defaultFee) {
		specificFee = defaultFee
	}

	finalFeeAmount := specificFee.Quo(feesConfiguration.FeeCoinPrice).TruncateInt()
	return sdk.NewCoins(sdk.NewCoin(feesConfiguration.FeeCoinDenom, finalFeeAmount))
}

func getFee(feesConfig *configuration.Fees, msg sdk.Msg) sdk.Dec {
	switch m := msg.(type) {
	case *types.MsgCreateEscrow:
		if m.Object != nil {
			if obj, hasCustomFees := m.Object.GetCachedValue().(types.ObjectWithCustomFees); hasCustomFees {
				return obj.GetCreationFees()
			}
		}
		return feesConfig.CreateEscrow
	case *types.MsgUpdateEscrow:
		return feesConfig.UpdateEscrow
	case *types.MsgTransferToEscrow:
		return feesConfig.TransferToEscrow
	case *types.MsgRefundEscrow:
		return feesConfig.RefundEscrow
	default:
		return feesConfig.FeeDefault
	}
}
