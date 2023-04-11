package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/umma-chain/umma-core/x/factory/types"
)

func (k msgServer) BurnToken(goCtx context.Context, msg *types.MsgBurnToken) (*types.MsgBurnTokenResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	// Check if the value exists
	valFound, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "denom does not exist")
	}

	// Checks if the the msg owner is the same as the current owner
	if msg.Owner != valFound.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	tokens := sdk.NewCoin(msg.Denom, sdk.NewInt(int64(msg.Amount)))

	sender, err := sdk.AccAddressFromBech32(msg.Owner)

	if err != nil {
		return nil, err
	}
	if valFound.Supply < msg.Amount {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Your balance is simular")
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(tokens)); err != nil {
		return nil, err
	}

	if err := k.bankKeeper.BurnCoins(
		ctx, types.ModuleName, sdk.NewCoins(tokens),
	); err != nil {
		// NOTE: should not happen as the module account was
		// retrieved on the step above and it has enough balance
		// to burn.
		panic(fmt.Sprintf("cannot burn coins after a successful send to a module account: %v", err))
	}

	var denom = types.Denom{
		Owner:              valFound.Owner,
		Denom:              valFound.Denom,
		Description:        valFound.Description,
		MaxSupply:          valFound.MaxSupply,
		Supply:             valFound.Supply - msg.Amount,
		Decimal:            valFound.Decimal,
		Symbol:             valFound.Symbol,
		Url:                valFound.Url,
		LogoUrl:            valFound.LogoUrl,
		CanChangeMaxSupply: valFound.CanChangeMaxSupply,
	}

	k.SetDenom(
		ctx,
		denom,
	)

	return &types.MsgBurnTokenResponse{}, nil
}
