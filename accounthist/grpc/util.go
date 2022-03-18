package grpc

import zsw "github.com/zhongshuwen/zswchain-go"

type EOSName uint64

func (n EOSName) String() string {
	return zsw.NameToString(uint64(n))
}
