package keeper_test

import (
	"testing"

	testkeeper "github.com/LoyalLabs/avecom-chain/testutil/keeper"
	"github.com/LoyalLabs/avecom-chain/x/avecomchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.AvecomchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
