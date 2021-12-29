package smplsecretsvaultchain

import (
	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain/keeper"
	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.Passports != nil {
		k.SetPassports(ctx, *genState.Passports)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all passports
	passports, found := k.GetPassports(ctx)
	if found {
		genesis.Passports = &passports
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
