// Copyright 2020 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eosws

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	zsw "github.com/zhongshuwen/zswchain-go"
	"go.opencensus.io/trace"
)

func NewTestContext() context.Context {
	return NewParameterizedTestContext("00000000000000000000000000000001", "test")
}

func NewParameterizedTestContext(hexTraceID string, spanName string) context.Context {
	traceID := fixedTraceID(hexTraceID)
	spanContext := trace.SpanContext{TraceID: traceID}
	ctx, _ := trace.StartSpanWithRemoteParent(context.Background(), "test", spanContext)

	return ctx
}

type TestABIGetter struct {
	abis map[zsw.AccountName]*zsw.ABI
}

func NewTestABIGetter() *TestABIGetter {
	return &TestABIGetter{
		abis: map[zsw.AccountName]*zsw.ABI{},
	}
}

func (g *TestABIGetter) SetABIForAccount(abiString string, account zsw.AccountName) {

	abi, err := zsw.NewABI(strings.NewReader(abiString))
	if err != nil {
		panic(err)
	}

	g.abis[account] = abi
}

func (g *TestABIGetter) GetABI(ctx context.Context, blockNum uint32, account zsw.AccountName) (*zsw.ABI, error) {
	return g.abis[account], nil
}

type TestAccountGetter struct {
	jsonData string
}

func NewTestAccountGetter() *TestAccountGetter {
	return &TestAccountGetter{}
}

func (g *TestAccountGetter) GetAccount(ctx context.Context, name string) (out *zsw.AccountResp, err error) {
	if g.jsonData == "" {
		return nil, fmt.Errorf("simulated error")
	}

	err = json.Unmarshal([]byte(g.jsonData), &out)
	if err != nil {
		panic(err)
	}
	return out, nil
}

func (g *TestAccountGetter) SetAccount(jsonData string) {
	g.jsonData = jsonData
}

func fixedTraceID(hexInput string) (out trace.TraceID) {
	rawTraceID, _ := hex.DecodeString(hexInput)
	copy(out[:], rawTraceID)

	return
}
