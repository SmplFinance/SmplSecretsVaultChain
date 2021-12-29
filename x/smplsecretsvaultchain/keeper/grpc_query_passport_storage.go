package keeper

import (
	"context"

	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) PassportStorage(c context.Context, req *types.QueryGetPassportStorageRequest) (*types.QueryGetPassportStorageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetPassportStorage(ctx)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetPassportStorageResponse{PassportStorage: val}, nil
}
