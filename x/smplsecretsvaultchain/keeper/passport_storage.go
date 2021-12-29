package keeper

import (
	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetPassportStorage set passportStorage in the store
func (k Keeper) SetPassportStorage(ctx sdk.Context, passportStorage types.PassportStorage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PassportStorageKey))
	b := k.cdc.MustMarshal(&passportStorage)
	store.Set([]byte{0}, b)
}

// GetPassportStorage returns passportStorage
func (k Keeper) GetPassportStorage(ctx sdk.Context) (val types.PassportStorage, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PassportStorageKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemovePassportStorage removes passportStorage from the store
func (k Keeper) RemovePassportStorage(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.PassportStorageKey))
	store.Delete([]byte{0})
}
