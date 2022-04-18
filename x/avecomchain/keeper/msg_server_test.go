package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/LoyalLabs/avecom-chain/testutil/keeper"
	"github.com/LoyalLabs/avecom-chain/x/avecomchain/keeper"
	"github.com/LoyalLabs/avecom-chain/x/avecomchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.AvecomchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
