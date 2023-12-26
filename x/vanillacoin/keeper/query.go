package keeper

import (
	"vanillacoin/x/vanillacoin/types"
)

var _ types.QueryServer = Keeper{}
