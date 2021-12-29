package types_test

import (
	"testing"

	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				Passport: &types.Passport{
					Mnemonic:    []string{"49"},
					Prefix:      "53",
					Hdpath:      "2",
					ChainSymbol: "55",
				},
				PassportStorage: &types.PassportStorage{
					Passport:            new(types.Passport),
					TestPhrase:          "38",
					EncryptedTestPhrase: "32",
				},
				Passports: &types.Passports{
					PassportStorageMap:  "91",
					TestPhrase:          "92",
					EncryptedTestPhrase: "81",
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
