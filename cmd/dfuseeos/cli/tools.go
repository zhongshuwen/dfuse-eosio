package cli

import "github.com/zhongshuwen/dfuse-eosio/tools"

func init() {
	RootCmd.AddCommand(tools.Cmd)
}
