package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/umma-chain/umma-core/testutil/sample"
)

func TestMsgBurnToken_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgBurnToken
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgBurnToken{
				Owner: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgBurnToken{
				Owner: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
