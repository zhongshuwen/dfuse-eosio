package cli

import (
	"strings"

	blockmetaApp "github.com/dfuse-io/blockmeta/app/blockmeta"
	"github.com/invisible-train-40/zsw-lishi-launcher/launcher"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	dblockmeta "github.com/zhongshuwen/dfuse-eosio/blockmeta"
	"github.com/zhongshuwen/dfuse-eosio/trxdb"
	zsw "github.com/zhongshuwen/zswchain-go"
)

func init() {
	// Blockmeta
	launcher.RegisterApp(&launcher.AppDef{
		ID:          "blockmeta",
		Title:       "Blockmeta",
		Description: "Serves information about blocks",
		MetricsID:   "blockmeta",
		Logger:      launcher.NewLoggingDef("github.com/dfuse-io/blockmeta.*", nil),
		RegisterFlags: func(cmd *cobra.Command) error {
			cmd.Flags().String("blockmeta-grpc-listen-addr", BlockmetaServingAddr, "Address to listen for incoming gRPC requests")
			cmd.Flags().Bool("blockmeta-live-source", true, "Whether we want to connect to a live block source or not.")
			cmd.Flags().Bool("blockmeta-enable-readiness-probe", true, "Enable blockmeta's app readiness probe")
			cmd.Flags().StringSlice("blockmeta-eos-api-upstream-addr", []string{NodeosAPIAddr}, "EOS API address to fetch info from running chain, must be in-sync")
			cmd.Flags().StringSlice("blockmeta-eos-api-extra-addr", []string{MindreaderNodeosAPIAddr}, "Additional EOS API address for ID lookups (valid even if it is out of sync or read-only)")
			return nil
		},
		FactoryFunc: func(runtime *launcher.Runtime) (launcher.App, error) {
			for _, addr := range viper.GetStringSlice("blockmeta-eos-api-upstream-addr") {
				if !strings.HasPrefix(addr, "http") {
					addr = "http://" + addr
				}
				dblockmeta.APIs = append(dblockmeta.APIs, zsw.New(addr))
			}
			for _, addr := range viper.GetStringSlice("blockmeta-eos-api-extra-addr") {
				if !strings.HasPrefix(addr, "http") {
					addr = "http://" + addr
				}
				dblockmeta.ExtraAPIs = append(dblockmeta.ExtraAPIs, zsw.New(addr))
			}

			dfuseDataDir := runtime.AbsDataDir

			// FIXME: Use the new WithLogger by getting the app logger!
			trxdbClient, err := trxdb.New(mustReplaceDataDir(dfuseDataDir, viper.GetString("common-trxdb-dsn")))
			if err != nil {
				return nil, err
			}

			//todo: add db to a modules struct in blockmeta
			db := &dblockmeta.EOSBlockmetaDB{
				Driver: trxdbClient,
			}

			return blockmetaApp.New(&blockmetaApp.Config{
				BlocksStoreURL:  mustReplaceDataDir(dfuseDataDir, viper.GetString("common-blocks-store-url")),
				BlockStreamAddr: viper.GetString("common-blockstream-addr"),
				GRPCListenAddr:  viper.GetString("blockmeta-grpc-listen-addr"),
				Protocol:        Protocol,
				LiveSource:      viper.GetBool("blockmeta-live-source"),
			}, db), nil
		},
	})
}
