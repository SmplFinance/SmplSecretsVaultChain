package cli_test

import (
	"fmt"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/status"

	"github.com/SmplFinance/SmplSecretsVaultChain/testutil/network"
	"github.com/SmplFinance/SmplSecretsVaultChain/testutil/nullify"
	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain/client/cli"
	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
)

func networkWithPassportStorageObjects(t *testing.T) (*network.Network, types.PassportStorage) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	passportStorage := &types.PassportStorage{}
	nullify.Fill(&passportStorage)
	state.PassportStorage = passportStorage
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), *state.PassportStorage
}

func TestShowPassportStorage(t *testing.T) {
	net, obj := networkWithPassportStorageObjects(t)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc string
		args []string
		err  error
		obj  types.PassportStorage
	}{
		{
			desc: "get",
			args: common,
			obj:  obj,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			var args []string
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowPassportStorage(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetPassportStorageResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.PassportStorage)
				require.Equal(t,
					nullify.Fill(&tc.obj),
					nullify.Fill(&resp.PassportStorage),
				)
			}
		})
	}
}
