package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBurnTokens = "burn_tokens"

var _ sdk.Msg = &MsgBurnTokens{}

func NewMsgBurnTokens(owner string, denom string, amount int32) *MsgBurnTokens {
	return &MsgBurnTokens{
		Owner:  owner,
		Denom:  denom,
		Amount: amount,
	}
}

func (msg *MsgBurnTokens) Route() string {
	return RouterKey
}

func (msg *MsgBurnTokens) Type() string {
	return TypeMsgBurnTokens
}

func (msg *MsgBurnTokens) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgBurnTokens) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
