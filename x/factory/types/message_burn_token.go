package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgBurnToken = "burn_token"

var _ sdk.Msg = &MsgBurnToken{}

func NewMsgBurnToken(owner string, denom string, amount uint64) *MsgBurnToken {
	return &MsgBurnToken{
		Owner:  owner,
		Denom:  denom,
		Amount: amount,
	}
}

func (msg *MsgBurnToken) Route() string {
	return RouterKey
}

func (msg *MsgBurnToken) Type() string {
	return TypeMsgBurnToken
}

func (msg *MsgBurnToken) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

func (msg *MsgBurnToken) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnToken) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	return nil
}
