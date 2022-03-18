package tokenhist

import (
	"github.com/dfuse-io/logging"
	"go.uber.org/zap"
)

var zlog *zap.Logger

func init() {
	logging.Register("github.com/zhongshuwen/dfuse-eosio/accounthist/app/accounthist", &zlog)
}
