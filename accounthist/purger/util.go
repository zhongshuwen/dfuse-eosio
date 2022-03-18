package purger

import "github.com/zhongshuwen/zswchain-go"

type EOSName uint64

func (n EOSName) String() string {
	return eos.NameToString(uint64(n))
}
