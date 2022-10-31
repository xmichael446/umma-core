package keeper

import (
	"github.com/umma-chain/core-umma/x/factory/types"
)

var _ types.QueryServer = Keeper{}
