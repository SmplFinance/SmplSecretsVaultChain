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

		Passport: &types.Passport{
			Mnemonic:    []string{"38"},
			Prefix:      "56",
			Hdpath:      "22",
			ChainSymbol: "21",
		},
		PassportStorage: &types.PassportStorage{
			Passport:            new(types.Passport),
			TestPhrase:          "60",
			EncryptedTestPhrase: "51",
		},
		Passports: &types.Passports{
			PassportStorageMap:  "82",
			TestPhrase:          "3",
			EncryptedTestPhrase: "43",
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SmplsecretsvaultchainKeeper(t)
	smplsecretsvaultchain.InitGenesis(ctx, *k, genesisState)
	got := smplsecretsvaultchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.Passport, got.Passport)
	require.Equal(t, genesisState.PassportStorage, got.PassportStorage)
	require.Equal(t, genesisState.Passports, got.Passports)
	// this line is used by starport scaffolding # genesis/test/assert
}
