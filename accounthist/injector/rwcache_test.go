package injector

import (
	"context"
	"testing"

	"github.com/zhongshuwen/dfuse-eosio/accounthist"
	zsw "github.com/zhongshuwen/zswchain-go"

	"github.com/stretchr/testify/assert"
)

func TestRWCache(t *testing.T) {
	kvStore, cleanup := getKVTestFactory(t)
	defer cleanup()
	rwCache := NewRWCache(kvStore)
	ctx := context.Background()

	accountMama := zsw.MustStringToName("mama")
	accountPapa := zsw.MustStringToName("papa")
	accountDada := zsw.MustStringToName("dada")

	rwCache.Put(ctx, accounthist.AccountFacet(accountMama).Row(0, 0), []byte{0xaa})
	rwCache.Put(ctx, accounthist.AccountFacet(accountMama).Row(0, 2), []byte{0xaa})
	rwCache.Put(ctx, accounthist.AccountFacet(accountMama).Row(0, 1), []byte{0xaa})
	rwCache.Put(ctx, accounthist.AccountFacet(accountPapa).Row(0, 24), []byte{0xaa})
	rwCache.Put(ctx, accounthist.AccountFacet(accountPapa).Row(0, 23), []byte{0xaa})
	rwCache.Put(ctx, accounthist.AccountFacet(accountDada).Row(0, 25), []byte{0xaa})
	rwCache.Put(ctx, accounthist.AccountFacet(accountPapa).Row(0, 25), []byte{0xaa})

	rwCache.BatchDelete(ctx, [][]byte{
		accounthist.AccountFacet(accountMama).Row(0, 1),
		accounthist.AccountFacet(accountPapa).Row(0, 23),
	})

	expectedKeys := [][]byte{
		accounthist.AccountFacet(accountMama).Row(0, 0),
		accounthist.AccountFacet(accountMama).Row(0, 2),
		accounthist.AccountFacet(accountPapa).Row(0, 24),
		accounthist.AccountFacet(accountDada).Row(0, 25),
		accounthist.AccountFacet(accountPapa).Row(0, 25),
	}
	i := 0
	rwCache.OrderedPuts(func(sKey string, value []byte) error {
		assert.Equal(t, string(expectedKeys[i]), sKey)
		i += 1
		return nil
	})
}
