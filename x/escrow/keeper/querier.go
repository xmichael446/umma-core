package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/umma-chain/umma-core/x/escrow/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new escrow Querier instance
func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryEscrow:
			return queryEscrow(ctx, req, k, legacyQuerierCdc)
		case types.QueryEscrows:
			return queryEscrows(ctx, req, k, legacyQuerierCdc)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query path: %s", types.ModuleName, path[0])
		}
	}
}

func queryEscrow(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryEscrowParams
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	if err := types.ValidateID(params.Id); err != nil {
		return nil, sdkerrors.Wrap(err, "Invalid provided ID")
	}

	escrow, found := k.GetEscrow(ctx, params.Id)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrEscrowNotFound, params.Id)
	}

	bz, err := legacyQuerierCdc.MarshalJSON(&types.QueryEscrowResponse{Escrow: &escrow})
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Cannot marshall the queried escrow")
	}
	return bz, nil
}

func queryEscrows(ctx sdk.Context, req abci.RequestQuery, k Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryEscrowsParams
	if err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	escrows, err := k.queryEscrowsByAttributes(
		ctx,
		params.Seller,
		params.State,
		params.ObjectKey,
		params.PaginationStart,
		params.PaginationLength,
	)
	if err != nil {
		return nil, err
	}

	bz, err := legacyQuerierCdc.MarshalJSON(&types.QueryEscrowsResponse{Escrows: escrows})
	if err != nil {
		return nil, sdkerrors.Wrap(err, "Cannot marshall the queried escrows")
	}
	return bz, nil
}
