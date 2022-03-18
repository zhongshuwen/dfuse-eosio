package injector

import (
	"github.com/zhongshuwen/dfuse-eosio/accounthist"
)

func (i *Injector) UpdateSeqData(key accounthist.Facet, seqData accounthist.SequenceData) {
	i.cacheSeqData[key.String()] = seqData
}
