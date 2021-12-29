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

func createTestPassport(keeper *keeper.Keeper, ctx sdk.Context) types.Passport {
	item := types.Passport{}
	keeper.SetPassport(ctx, item)
	return item
}

func TestPassportGet(t *testing.T) {
	keeper, ctx := keepertest.SmplsecretsvaultchainKeeper(t)
	item := createTestPassport(keeper, ctx)
	rst, found := keeper.GetPassport(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestPassportRemove(t *testing.T) {
	keeper, ctx := keepertest.SmplsecretsvaultchainKeeper(t)
	createTestPassport(keeper, ctx)
	keeper.RemovePassport(ctx)
	_, found := keeper.GetPassport(ctx)
	require.False(t, found)
}
