package tools

import (
	"os"

	"github.com/dfuse-io/logging"
	"go.uber.org/zap"
)

var traceEnabled = os.Getenv("TRACE") == "true"
var zlog = zap.NewNop()

func init() {
	logging.Register("github.com/zhongshuwen/dfuse-eosio/tools", &zlog)
}
