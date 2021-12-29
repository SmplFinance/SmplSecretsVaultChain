package smplsecretsvaultchain_test

import (
	"testing"

	keepertest "github.com/SmplFinance/SmplSecretsVaultChain/testutil/keeper"
	"github.com/SmplFinance/SmplSecretsVaultChain/testutil/nullify"
	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain"
	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SmplsecretsvaultchainKeeper(t)
	smplsecretsvaultchain.InitGenesis(ctx, *k, genesisState)
	got := smplsecretsvaultchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
