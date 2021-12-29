package keeper

import (
	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
)

var _ types.QueryServer = Keeper{}
