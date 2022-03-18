package injector

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zhongshuwen/dfuse-eosio/accounthist"
	ct "github.com/zhongshuwen/dfuse-eosio/codec/testing"
	zsw "github.com/zhongshuwen/zswchain-go"
)

func Test_shardSummary(t *testing.T) {
	kvStore, cleanup := getKVTestFactory(t)
	defer cleanup()
	maxEntries := uint64(10)

	s := setupAccountInjector(NewRWCache(kvStore), 0, 1000)
	runShard(t, 0, maxEntries, kvStore,
		ct.Block(t, "00000002bb",
			ct.TrxTrace(t, ct.ActionTrace(t, "a:some:cthing1", ct.GlobalSequence(3))),
			ct.TrxTrace(t, ct.ActionTrace(t, "a:some:cthing2", ct.GlobalSequence(4))),
			ct.TrxTrace(t, ct.ActionTrace(t, "a:some:cthing3", ct.GlobalSequence(5))),
			ct.TrxTrace(t, ct.ActionTrace(t, "a:some:cthing4", ct.GlobalSequence(6))),
			ct.TrxTrace(t, ct.ActionTrace(t, "a:some:cthing5", ct.GlobalSequence(7))),
		),
	)

	runShard(t, 1, maxEntries, kvStore,
		ct.Block(t, "00000001aa",
			ct.TrxTrace(t, ct.ActionTrace(t, "a:some:athing1", ct.GlobalSequence(1))),
			ct.TrxTrace(t, ct.ActionTrace(t, "a:some:athing2", ct.GlobalSequence(2))),
		),
	)

	summary, err := s.FacetShardsSummary(context.Background(), accounthist.AccountFacet(zsw.MustStringToName("a")))
	require.NoError(t, err)
	assert.Equal(t, []*FacetShardSummary{
		{ShardNum: 0, SeqData: accounthist.SequenceData{CurrentOrdinal: 5, LastGlobalSeq: 7}},
		{ShardNum: 1, SeqData: accounthist.SequenceData{CurrentOrdinal: 2, LastGlobalSeq: 2}},
	}, summary)

}
