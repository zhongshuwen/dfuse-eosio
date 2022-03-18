package main

import (
	"github.com/zhongshuwen/dfuse-eosio/cmd/dfuseeos/cli"
)

var version = "dev"
var commit = ""

func init() {
	cli.RootCmd.Version = version + "-" + commit
}

func main() {
	cli.Main()
}
