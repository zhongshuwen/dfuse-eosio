package tokenmeta

import (
	"github.com/dfuse-io/logging"
	"go.uber.org/zap"
)

var zlog = zap.NewNop()

func init() {
	logging.Register("github.com/zhongshuwen/dfuse-eosio/tokenmeta", &zlog)
}
