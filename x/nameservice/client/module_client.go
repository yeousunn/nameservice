package client

import amino "github.com/tendermint/go-amino"

type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

func(mc ModuleClient) GetQueryCmd() *cobra.Command {
	govQueryCmd := &cobra.Command{
		Use: "nameservice",
		Short: "Querying commands for the nameservice module"
	}

	govQueryCmd.AddCommand(client.GetCommands(
		nameservicecmd.GetCmdResolveName(mc.StoreKey, mc.cdc),
		nameservicecmd.GetCmdWhois(mc.storeKey, mc.cdc),
	)...)

	return govQueryCmd
}

func (mc ModuleClient) GetTxCmd() *cobra.Command {
	govTxCmd := &cobra.Command{
		Use:   "nameservice",
		Short: "Nameservice transactions subcommands",
	}

	govTxCmd.AddCommand(client.PostCommands(
		nameservicecmd.GetCmdBuyName(mc.cdc),
		nameservicecmd.GetCmdSetName(mc.cdc),
	)...)

	return govTxCmd
}