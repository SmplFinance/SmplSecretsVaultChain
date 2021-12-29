package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "github.com/SmplFinance/SmplSecretsVaultChain/testutil/keeper"
	"github.com/SmplFinance/SmplSecretsVaultChain/testutil/nullify"
	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain/keeper"
	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
)

func createTestPassports(keeper *keeper.Keeper, ctx sdk.Context) types.Passports {
	item := types.Passports{}
	keeper.SetPassports(ctx, item)
	return item
}

func TestPassportsGet(t *testing.T) {
	keeper, ctx := keepertest.SmplsecretsvaultchainKeeper(t)
	item := createTestPassports(keeper, ctx)
	rst, found := keeper.GetPassports(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestPassportsRemove(t *testing.T) {
	keeper, ctx := keepertest.SmplsecretsvaultchainKeeper(t)
	createTestPassports(keeper, ctx)
	keeper.RemovePassports(ctx)
	_, found := keeper.GetPassports(ctx)
	require.False(t, found)
}
