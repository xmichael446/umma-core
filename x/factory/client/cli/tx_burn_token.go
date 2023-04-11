package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/umma-chain/umma-core/x/factory/types"
)

var _ = strconv.Itoa(0)

func CmdBurnToken() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burn-token [denom] [amount]",
		Short: "Broadcast message BurnToken",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argDenom := args[0]
			argAmount, err := cast.ToUint64E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgBurnToken(
				clientCtx.GetFromAddress().String(),
				argDenom,
				argAmount,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
