package keeper

import (
	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPassport set passport in the store
func (k Keeper) SetPassport(ctx sdk.Context, passport types.Passport) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PassportKey))
	b := k.cdc.MustMarshal(&passport)
	store.Set([]byte{0}, b)
}

// GetPassport returns passport
func (k Keeper) GetPassport(ctx sdk.Context) (val types.Passport, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PassportKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePassport removes passport from the store
func (k Keeper) RemovePassport(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PassportKey))
	store.Delete([]byte{0})
}
