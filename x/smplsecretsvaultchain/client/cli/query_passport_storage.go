package cli

import (
	"context"

	"github.com/SmplFinance/SmplSecretsVaultChain/x/smplsecretsvaultchain/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
)

func CmdShowPassportStorage() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-passport-storage",
		Short: "shows passportStorage",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetPassportStorageRequest{}

			res, err := queryClient.PassportStorage(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
