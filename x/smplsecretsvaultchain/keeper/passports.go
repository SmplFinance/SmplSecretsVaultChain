package keeper

import (
	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPassports set passports in the store
func (k Keeper) SetPassports(ctx sdk.Context, passports types.Passports) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PassportsKey))
	b := k.cdc.MustMarshal(&passports)
	store.Set([]byte{0}, b)
}

// GetPassports returns passports
func (k Keeper) GetPassports(ctx sdk.Context) (val types.Passports, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PassportsKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePassports removes passports from the store
func (k Keeper) RemovePassports(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PassportsKey))
	store.Delete([]byte{0})
}
