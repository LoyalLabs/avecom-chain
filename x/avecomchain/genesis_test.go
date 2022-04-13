package avecomchain_test

import (
	"testing"

	keepertest "github.com/LoyalLabs/avecom-chain/testutil/keeper"
	"github.com/LoyalLabs/avecom-chain/testutil/nullify"
	"github.com/LoyalLabs/avecom-chain/x/avecomchain"
	"github.com/LoyalLabs/avecom-chain/x/avecomchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AvecomchainKeeper(t)
	avecomchain.InitGenesis(ctx, *k, genesisState)
	got := avecomchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
