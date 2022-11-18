package keeper

import (
	"github.com/umma-chain/umma-core/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
