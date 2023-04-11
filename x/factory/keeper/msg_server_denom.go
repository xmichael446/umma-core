package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/umma-chain/umma-core/x/factory/types"
)

func (k msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "denom already exists, please change your denom name to another")
	}

	var denom = types.Denom{
		Owner:              msg.Owner,
		Denom:              msg.Denom,
		Description:        msg.Description,
		Symbol:             msg.Symbol,
		Decimal:            msg.Decimal,
		Url:                msg.Url,
		LogoUrl:            msg.LogoUrl,
		MaxSupply:          msg.MaxSupply,
		Supply:             0,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	k.SetDenom(
		ctx,
		denom,
	)
	return &types.MsgCreateDenomResponse{}, nil
}

func (k msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg owner is the same as the current owner
	if msg.Owner != valFound.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if !valFound.CanChangeMaxSupply && valFound.MaxSupply != msg.MaxSupply {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "cannot change maxsupply")
	}

	if !valFound.CanChangeMaxSupply && msg.CanChangeMaxSupply {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Cannot revert change maxsupply flag")
	}

	var denom = types.Denom{
		Owner:              msg.Owner,
		Denom:              valFound.Denom,
		Description:        msg.Description,
		Symbol:             valFound.Symbol,
		Decimal:            valFound.Decimal,
		Url:                msg.Url,
		LogoUrl:            msg.LogoUrl,
		MaxSupply:          msg.MaxSupply,
		Supply:             valFound.Supply,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
	}

	k.SetDenom(ctx, denom)

	return &types.MsgUpdateDenomResponse{}, nil
}

//func (k msgServer) DeleteDenom(goCtx context.Context, msg *types.MsgDeleteDenom) (*types.MsgDeleteDenomResponse, error) {
//	ctx := sdk.UnwrapSDKContext(goCtx)
//
//	// Check if the value exists
//	valFound, isFound := k.GetDenom(
//		ctx,
//		msg.Denom,
//	)
//	if !isFound {
//		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
//	}
//
//	// Checks if the the msg owner is the same as the current owner
//	if msg.Owner != valFound.Owner {
//		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
//	}
//
//	k.RemoveDenom(
//		ctx,
//		msg.Denom,
//	)
//
//	return &types.MsgDeleteDenomResponse{}, nil
//}
