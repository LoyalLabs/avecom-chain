package keeper

import (
	"github.com/LoyalLabs/avecom-chain/x/avecomchain/types"
)

var _ types.QueryServer = Keeper{}
