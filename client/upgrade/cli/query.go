package cli

import (
	"fmt"
	sdk "github.com/irisnet/irishub/types"
	"github.com/irisnet/irishub/codec"
	authcmd "github.com/irisnet/irishub/client/auth/cli"
	"github.com/irisnet/irishub/client/context"
	upgcli "github.com/irisnet/irishub/client/upgrade"
	"github.com/irisnet/irishub/modules/upgrade"
	"github.com/irisnet/irishub/modules/upgrade/params"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"github.com/irisnet/irishub/modules/params"
)

func GetInfoCmd(storeName string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "info",
		Short: "query the information of upgrade module",
		Example: "iriscli upgrade info",
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithLogger(os.Stdout).
				WithAccountDecoder(authcmd.GetAccountDecoder(cdc))

			res_height, _ := cliCtx.QueryStore(append([]byte(params.SignalParamspace + "/"), upgradeparams.ProposalAcceptHeightParameter.GetStoreKey()...), "params")
			res_proposalID, _ := cliCtx.QueryStore(append([]byte(params.SignalParamspace + "/"), upgradeparams.CurrentUpgradeProposalIdParameter.GetStoreKey()...), "params")
			var height int64
			var proposalID uint64
			cdc.UnmarshalJSON(res_height, &height)
			cdc.UnmarshalJSON(res_proposalID, &proposalID)

			res_versionID, _ := cliCtx.QueryStore(upgrade.GetCurrentVersionKey(), storeName)
			var versionID int64
			cdc.MustUnmarshalBinaryLengthPrefixed(res_versionID, &versionID)

			res_version, _ := cliCtx.QueryStore(upgrade.GetVersionIDKey(versionID), storeName)
			var version upgrade.Version
			cdc.MustUnmarshalBinaryLengthPrefixed(res_version, &version)

			upgradeInfoOutput := upgcli.ConvertUpgradeInfoToUpgradeOutput(version, proposalID, height)

			output, err := codec.MarshalJSONIndent(cdc, upgradeInfoOutput)
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}
	return cmd
}

// Command to Get a Switch Information
func GetCmdQuerySwitch(storeName string, cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-switch",
		Short: "query switch details",
		Example: "iriscli upgrade query-switch --proposal-id 1 --voter <voter address>",
		RunE: func(cmd *cobra.Command, args []string) error {
			proposalID := uint64(viper.GetInt64(flagProposalID))
			voterStr := viper.GetString(flagVoter)

			voter, err := sdk.AccAddressFromBech32(voterStr)
			if err != nil {
				return err
			}

			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithLogger(os.Stdout).
				WithAccountDecoder(authcmd.GetAccountDecoder(cdc))

			res, err := cliCtx.QueryStore(upgrade.GetSwitchKey(proposalID, voter), storeName)
			if len(res) == 0 || err != nil {
				return errors.Errorf("proposalID [%d] is not existed", proposalID)
			}

			var switchMsg upgrade.MsgSwitch
			cdc.MustUnmarshalBinaryLengthPrefixed(res, &switchMsg)
			output, err := codec.MarshalJSONIndent(cdc, switchMsg)
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	cmd.Flags().String(flagProposalID, "", "proposalID of upgrade swtich being queried")
	cmd.Flags().String(flagVoter, "", "Address sign the switch msg")

	return cmd
}
