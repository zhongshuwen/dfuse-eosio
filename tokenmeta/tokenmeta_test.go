package tokenmeta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	pbcodec "github.com/zhongshuwen/dfuse-eosio/pb/dfuse/eosio/codec/v1"
)

func TestTokenMeta_shouldProcessDbop(t *testing.T) {
	tests := []struct {
		name        string
		dbop        *pbcodec.DBOp
		expectValue bool
	}{
		{
			name: "accounts table",
			dbop: &pbcodec.DBOp{
				TableName: "accounts",
			},
			expectValue: true,
		},
		{
			name: "stat table",
			dbop: &pbcodec.DBOp{
				TableName: "stat",
			},
			expectValue: true,
		},
		{
			name: "invalid table",
			dbop: &pbcodec.DBOp{
				TableName: "stats",
			},
			expectValue: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expectValue, shouldProcessDbop(test.dbop, pbcodec.AlwaysIncludedFilteringActionMatcher))
		})
	}

}
