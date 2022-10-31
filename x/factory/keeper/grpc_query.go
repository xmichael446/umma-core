package keeper

import (
	"github.com/umma-chain/umma-core/x/factory/types"
)

var _ types.QueryServer = Keeper{}
