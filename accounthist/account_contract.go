package accounthist

import (
	"fmt"

	"github.com/zhongshuwen/dfuse-eosio/accounthist/keyer"
	"github.com/zhongshuwen/zswchain-go"
)

type AccountContractKey struct {
	account  uint64
	contract uint64
}

func (a *AccountContractKey) Row(shard byte, seqData uint64) RowKey {
	return keyer.EncodeAccountContractKey(a.account, a.contract, shard, seqData)
}

func (a *AccountContractKey) String() string {
	return fmt.Sprintf("account (%s) contract (%s)", eos.NameToString(uint64(a.account)), eos.NameToString(uint64(a.contract)))
}

func (a *AccountContractKey) Account() uint64 {
	return a.account
}

func (a *AccountContractKey) Bytes() []byte {
	return keyer.EncodeAccountContractPrefixKey(a.account, a.contract)
}
