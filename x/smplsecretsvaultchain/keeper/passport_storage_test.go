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

func createTestPassportStorage(keeper *keeper.Keeper, ctx sdk.Context) types.PassportStorage {
	item := types.PassportStorage{}
	keeper.SetPassportStorage(ctx, item)
	return item
}

func TestPassportStorageGet(t *testing.T) {
	keeper, ctx := keepertest.SmplsecretsvaultchainKeeper(t)
	item := createTestPassportStorage(keeper, ctx)
	rst, found := keeper.GetPassportStorage(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestPassportStorageRemove(t *testing.T) {
	keeper, ctx := keepertest.SmplsecretsvaultchainKeeper(t)
	createTestPassportStorage(keeper, ctx)
	keeper.RemovePassportStorage(ctx)
	_, found := keeper.GetPassportStorage(ctx)
	require.False(t, found)
}
