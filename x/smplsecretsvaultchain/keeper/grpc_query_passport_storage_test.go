package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/SmplFinance/SmplSecretsVaultChain/testutil/keeper"
	"github.com/SmplFinance/SmplSecretsVaultChain/testutil/nullify"
	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
)

func TestPassportStorageQuery(t *testing.T) {
	keeper, ctx := keepertest.SmplsecretsvaultchainKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestPassportStorage(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetPassportStorageRequest
		response *types.QueryGetPassportStorageResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetPassportStorageRequest{},
			response: &types.QueryGetPassportStorageResponse{PassportStorage: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.PassportStorage(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
