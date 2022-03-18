// Copyright 2019 dfuse Platform Inc.
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

package codec

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	pbcodec "github.com/zhongshuwen/dfuse-eosio/pb/dfuse/eosio/codec/v1"
	zsw "github.com/zhongshuwen/zswchain-go"
	"github.com/zhongshuwen/zswchain-go/ecc"
	"go.uber.org/zap"
)

type conversionOption interface{}

type actionConversionOption interface {
	apply(actionTrace *pbcodec.ActionTrace)
}

type actionConversionOptionFunc func(actionTrace *pbcodec.ActionTrace)

func (f actionConversionOptionFunc) apply(actionTrace *pbcodec.ActionTrace) {
	f(actionTrace)
}

func limitConsoleLengthConversionOption(maxByteCount int) conversionOption {
	return actionConversionOptionFunc(func(in *pbcodec.ActionTrace) {
		if maxByteCount == 0 {
			return
		}

		if len(in.Console) > maxByteCount {
			in.Console = in.Console[:maxByteCount]

			// Prior truncation, the string had only valid UTF-8 charaters, so at worst, we will need
			// 3 bytes (`utf8.UTFMax - 1`) to reach a valid UTF-8 sequence.
			for i := 0; i < utf8.UTFMax-1; i++ {
				lastRune, size := utf8.DecodeLastRuneInString(in.Console)
				if lastRune != utf8.RuneError {
					// Last element is a valid utf8 character, nothing more to do here
					return
				}

				// We have an invalid UTF-8 sequence, size 0 means empty string, size 1 means invalid character
				if size == 0 {
					// The actual string was empty, nothing more to do here
					return
				}

				in.Console = in.Console[:len(in.Console)-1]
			}
		}
	})
}

func ActivatedProtocolFeaturesToDEOS(in *zsw.ProtocolFeatureActivationSet) *pbcodec.ActivatedProtocolFeatures {
	out := &pbcodec.ActivatedProtocolFeatures{}
	out.ProtocolFeatures = checksumsToBytesSlices(in.ProtocolFeatures)
	return out
}

func PendingScheduleToDEOS(in *zsw.PendingSchedule) *pbcodec.PendingProducerSchedule {
	out := &pbcodec.PendingProducerSchedule{
		ScheduleLibNum: in.ScheduleLIBNum,
		ScheduleHash:   []byte(in.ScheduleHash),
	}

	/// Specific versions handling

	// Only in EOSIO 1.x
	if in.Schedule.V1 != nil {
		out.ScheduleV1 = ProducerScheduleToDEOS(in.Schedule.V1)
	}

	// Only in EOSIO 2.x
	if in.Schedule.V2 != nil {
		out.ScheduleV2 = ProducerAuthorityScheduleToDEOS(in.Schedule.V2)
	}

	// End (versions)

	return out
}

func ProducerToLastProducedToDEOS(in []zsw.PairAccountNameBlockNum) []*pbcodec.ProducerToLastProduced {
	out := make([]*pbcodec.ProducerToLastProduced, len(in))
	for i, elem := range in {
		out[i] = &pbcodec.ProducerToLastProduced{
			Name:                 string(elem.AccountName),
			LastBlockNumProduced: uint32(elem.BlockNum),
		}
	}
	return out
}

func ProducerToLastImpliedIrbToDEOS(in []zsw.PairAccountNameBlockNum) []*pbcodec.ProducerToLastImpliedIRB {
	out := make([]*pbcodec.ProducerToLastImpliedIRB, len(in))
	for i, elem := range in {
		out[i] = &pbcodec.ProducerToLastImpliedIRB{
			Name:                 string(elem.AccountName),
			LastBlockNumProduced: uint32(elem.BlockNum),
		}
	}
	return out
}

func BlockrootMerkleToDEOS(merkle *zsw.MerkleRoot) *pbcodec.BlockRootMerkle {
	return &pbcodec.BlockRootMerkle{
		NodeCount:   uint32(merkle.NodeCount),
		ActiveNodes: checksumsToBytesSlices(merkle.ActiveNodes),
	}
}

func checksumsToBytesSlices(in []zsw.Checksum256) [][]byte {
	out := make([][]byte, len(in))
	for i, s := range in {
		out[i] = s
	}
	return out
}

func hexBytesToBytesSlices(in []zsw.HexBytes) [][]byte {
	out := make([][]byte, len(in))
	for i, s := range in {
		out[i] = s
	}
	return out
}

func bytesSlicesToHexBytes(in [][]byte) []zsw.HexBytes {
	out := make([]zsw.HexBytes, len(in))
	for i, s := range in {
		out[i] = s
	}
	return out
}

func BlockHeaderToDEOS(blockHeader *zsw.BlockHeader) *pbcodec.BlockHeader {
	out := &pbcodec.BlockHeader{
		Timestamp:        mustProtoTimestamp(blockHeader.Timestamp.Time),
		Producer:         string(blockHeader.Producer),
		Confirmed:        uint32(blockHeader.Confirmed),
		Previous:         blockHeader.Previous.String(),
		TransactionMroot: blockHeader.TransactionMRoot,
		ActionMroot:      blockHeader.ActionMRoot,
		ScheduleVersion:  blockHeader.ScheduleVersion,
		HeaderExtensions: ExtensionsToDEOS(blockHeader.HeaderExtensions),
	}

	if blockHeader.NewProducersV1 != nil {
		out.NewProducersV1 = ProducerScheduleToDEOS(blockHeader.NewProducersV1)
	}

	return out
}

func BlockHeaderToEOS(in *pbcodec.BlockHeader) *zsw.BlockHeader {
	stamp, _ := ptypes.Timestamp(in.Timestamp)
	prev, _ := hex.DecodeString(in.Previous)
	out := &zsw.BlockHeader{
		Timestamp:        zsw.BlockTimestamp{Time: stamp},
		Producer:         zsw.AccountName(in.Producer),
		Confirmed:        uint16(in.Confirmed),
		Previous:         prev,
		TransactionMRoot: in.TransactionMroot,
		ActionMRoot:      in.ActionMroot,
		ScheduleVersion:  in.ScheduleVersion,
		HeaderExtensions: ExtensionsToEOS(in.HeaderExtensions),
	}

	if in.NewProducersV1 != nil {
		out.NewProducersV1 = ProducerScheduleToEOS(in.NewProducersV1)
	}

	return out
}

func BlockSigningAuthorityToDEOS(authority *zsw.BlockSigningAuthority) *pbcodec.BlockSigningAuthority {
	out := &pbcodec.BlockSigningAuthority{}

	switch v := authority.Impl.(type) {
	case *zsw.BlockSigningAuthorityV0:
		out.Variant = &pbcodec.BlockSigningAuthority_V0{
			V0: &pbcodec.BlockSigningAuthorityV0{
				Threshold: v.Threshold,
				Keys:      KeyWeightsPToDEOS(v.Keys),
			},
		}
	default:
		panic(fmt.Errorf("unable to convert zsw.BlockSigningAuthority to deos: wrong type %T", authority.Impl))
	}

	return out
}

func BlockSigningAuthorityToEOS(in *pbcodec.BlockSigningAuthority) *zsw.BlockSigningAuthority {
	switch v := in.Variant.(type) {
	case *pbcodec.BlockSigningAuthority_V0:
		return &zsw.BlockSigningAuthority{
			BaseVariant: zsw.BaseVariant{
				TypeID: zsw.BlockSigningAuthorityVariant.TypeID("block_signing_authority_v0"),
				Impl: zsw.BlockSigningAuthorityV0{
					Threshold: v.V0.Threshold,
					Keys:      KeyWeightsPToEOS(v.V0.Keys),
				},
			},
		}
	default:
		panic(fmt.Errorf("unknown block signing authority variant %t", in.Variant))
	}
}

func ProducerScheduleToDEOS(e *zsw.ProducerSchedule) *pbcodec.ProducerSchedule {
	return &pbcodec.ProducerSchedule{
		Version:   uint32(e.Version),
		Producers: ProducerKeysToDEOS(e.Producers),
	}
}

func ProducerScheduleToEOS(in *pbcodec.ProducerSchedule) *zsw.ProducerSchedule {
	return &zsw.ProducerSchedule{
		Version:   in.Version,
		Producers: ProducerKeysToEOS(in.Producers),
	}
}

func ProducerAuthorityScheduleToDEOS(e *zsw.ProducerAuthoritySchedule) *pbcodec.ProducerAuthoritySchedule {
	return &pbcodec.ProducerAuthoritySchedule{
		Version:   uint32(e.Version),
		Producers: ProducerAuthoritiesToDEOS(e.Producers),
	}
}

func ProducerAuthorityScheduleToEOS(in *pbcodec.ProducerAuthoritySchedule) *zsw.ProducerAuthoritySchedule {
	return &zsw.ProducerAuthoritySchedule{
		Version:   in.Version,
		Producers: ProducerAuthoritiesToEOS(in.Producers),
	}
}

func ProducerKeysToDEOS(in []zsw.ProducerKey) (out []*pbcodec.ProducerKey) {
	out = make([]*pbcodec.ProducerKey, len(in))
	for i, key := range in {
		out[i] = &pbcodec.ProducerKey{
			AccountName:     string(key.AccountName),
			BlockSigningKey: key.BlockSigningKey.String(),
		}
	}
	return
}

func ProducerKeysToEOS(in []*pbcodec.ProducerKey) (out []zsw.ProducerKey) {
	out = make([]zsw.ProducerKey, len(in))
	for i, producer := range in {
		// panic on error instead?
		key, _ := ecc.NewPublicKey(producer.BlockSigningKey)

		out[i] = zsw.ProducerKey{
			AccountName:     zsw.AccountName(producer.AccountName),
			BlockSigningKey: key,
		}
	}
	return
}

func PublicKeysToEOS(in []string) (out []*ecc.PublicKey) {
	if len(in) <= 0 {
		return nil
	}
	out = make([]*ecc.PublicKey, len(in))
	for i, inkey := range in {
		// panic on error instead?
		key, _ := ecc.NewPublicKey(inkey)

		out[i] = &key
	}
	return
}

func ExtensionsToDEOS(in []*zsw.Extension) (out []*pbcodec.Extension) {
	out = make([]*pbcodec.Extension, len(in))
	for i, extension := range in {
		out[i] = &pbcodec.Extension{
			Type: uint32(extension.Type),
			Data: extension.Data,
		}
	}

	return
}

func ExtensionsToEOS(in []*pbcodec.Extension) (out []*zsw.Extension) {
	if len(in) <= 0 {
		return nil
	}

	out = make([]*zsw.Extension, len(in))
	for i, extension := range in {
		out[i] = &zsw.Extension{
			Type: uint16(extension.Type),
			Data: extension.Data,
		}
	}
	return
}

func ProducerAuthoritiesToDEOS(producerAuthorities []*zsw.ProducerAuthority) (out []*pbcodec.ProducerAuthority) {
	if len(producerAuthorities) <= 0 {
		return nil
	}

	out = make([]*pbcodec.ProducerAuthority, len(producerAuthorities))
	for i, authority := range producerAuthorities {
		out[i] = &pbcodec.ProducerAuthority{
			AccountName:           string(authority.AccountName),
			BlockSigningAuthority: BlockSigningAuthorityToDEOS(authority.BlockSigningAuthority),
		}
	}
	return
}

func ProducerAuthoritiesToEOS(producerAuthorities []*pbcodec.ProducerAuthority) (out []*zsw.ProducerAuthority) {
	if len(producerAuthorities) <= 0 {
		return nil
	}

	out = make([]*zsw.ProducerAuthority, len(producerAuthorities))
	for i, authority := range producerAuthorities {
		out[i] = &zsw.ProducerAuthority{
			AccountName:           zsw.AccountName(authority.AccountName),
			BlockSigningAuthority: BlockSigningAuthorityToEOS(authority.BlockSigningAuthority),
		}
	}
	return
}

func TransactionReceiptToDEOS(txReceipt *zsw.TransactionReceipt) *pbcodec.TransactionReceipt {
	receipt := &pbcodec.TransactionReceipt{
		Status:               TransactionStatusToDEOS(txReceipt.Status),
		CpuUsageMicroSeconds: txReceipt.CPUUsageMicroSeconds,
		NetUsageWords:        uint32(txReceipt.NetUsageWords),
	}

	receipt.Id = txReceipt.Transaction.ID.String()
	if txReceipt.Transaction.Packed != nil {
		receipt.PackedTransaction = &pbcodec.PackedTransaction{
			Signatures:            SignaturesToDEOS(txReceipt.Transaction.Packed.Signatures),
			Compression:           uint32(txReceipt.Transaction.Packed.Compression),
			PackedContextFreeData: txReceipt.Transaction.Packed.PackedContextFreeData,
			PackedTransaction:     txReceipt.Transaction.Packed.PackedTransaction,
		}
	}

	return receipt
}

func TransactionReceiptHeaderToDEOS(in *zsw.TransactionReceiptHeader) *pbcodec.TransactionReceiptHeader {
	return &pbcodec.TransactionReceiptHeader{
		Status:               TransactionStatusToDEOS(in.Status),
		CpuUsageMicroSeconds: in.CPUUsageMicroSeconds,
		NetUsageWords:        uint32(in.NetUsageWords),
	}
}

func TransactionReceiptHeaderToEOS(in *pbcodec.TransactionReceiptHeader) *zsw.TransactionReceiptHeader {
	return &zsw.TransactionReceiptHeader{
		Status:               TransactionStatusToEOS(in.Status),
		CPUUsageMicroSeconds: in.CpuUsageMicroSeconds,
		NetUsageWords:        zsw.Varuint32(in.NetUsageWords),
	}
}

func SignaturesToDEOS(in []ecc.Signature) (out []string) {

	out = make([]string, len(in))
	for i, signature := range in {
		out[i] = signature.String()
	}
	return
}

func SignaturesToEOS(in []string) []ecc.Signature {
	out := make([]ecc.Signature, len(in))
	for i, signature := range in {
		sig, err := ecc.NewSignature(signature)
		if err != nil {
			panic(fmt.Sprintf("failed to read signature %q: %s", signature, err))
		}

		out[i] = sig
	}
	return out
}

func TransactionTraceToDEOS(in *zsw.TransactionTrace, opts ...conversionOption) *pbcodec.TransactionTrace {
	id := in.ID.String()

	out := &pbcodec.TransactionTrace{
		Id:              id,
		BlockNum:        uint64(in.BlockNum),
		BlockTime:       mustProtoTimestamp(in.BlockTime.Time),
		ProducerBlockId: in.ProducerBlockID.String(),
		Elapsed:         int64(in.Elapsed),
		NetUsage:        uint64(in.NetUsage),
		Scheduled:       in.Scheduled,
		Exception:       ExceptionToDEOS(in.Except),
		ErrorCode:       ErrorCodeToDEOS(in.ErrorCode),
	}

	var someConsoleTruncated bool
	out.ActionTraces, someConsoleTruncated = ActionTracesToDEOS(in.ActionTraces, opts...)
	if someConsoleTruncated {
		zlog.Info("transaction had some of its action trace's console entries truncated", zap.String("id", id))
	}

	if in.FailedDtrxTrace != nil {
		out.FailedDtrxTrace = TransactionTraceToDEOS(in.FailedDtrxTrace, opts...)
	}
	if in.Receipt != nil {
		out.Receipt = TransactionReceiptHeaderToDEOS(in.Receipt)
	}

	return out
}

func TransactionTraceToEOS(in *pbcodec.TransactionTrace) (out *zsw.TransactionTrace) {
	out = &zsw.TransactionTrace{
		ID:              ChecksumToEOS(in.Id),
		BlockNum:        uint32(in.BlockNum),
		BlockTime:       TimestampToBlockTimestamp(in.BlockTime),
		ProducerBlockID: ChecksumToEOS(in.ProducerBlockId),
		Elapsed:         zsw.Int64(in.Elapsed),
		NetUsage:        zsw.Uint64(in.NetUsage),
		Scheduled:       in.Scheduled,
		ActionTraces:    ActionTracesToEOS(in.ActionTraces),
		Except:          ExceptionToEOS(in.Exception),
		ErrorCode:       ErrorCodeToEOS(in.ErrorCode),
	}

	if in.FailedDtrxTrace != nil {
		out.FailedDtrxTrace = TransactionTraceToEOS(in.FailedDtrxTrace)
	}
	if in.Receipt != nil {
		out.Receipt = TransactionReceiptHeaderToEOS(in.Receipt)
	}

	return out
}

func PermissionToDEOS(perm *zsw.Permission) *pbcodec.Permission {
	return &pbcodec.Permission{
		Name:         perm.PermName,
		Parent:       perm.Parent,
		RequiredAuth: AuthoritiesToDEOS(&perm.RequiredAuth),
	}
}

func AuthoritiesToDEOS(authority *zsw.Authority) *pbcodec.Authority {
	return &pbcodec.Authority{
		Threshold: authority.Threshold,
		Keys:      KeyWeightsToDEOS(authority.Keys),
		Accounts:  PermissionLevelWeightsToDEOS(authority.Accounts),
		Waits:     WaitWeightsToDEOS(authority.Waits),
	}
}

func AuthoritiesToEOS(authority *pbcodec.Authority) zsw.Authority {
	return zsw.Authority{
		Threshold: authority.Threshold,
		Keys:      KeyWeightsToEOS(authority.Keys),
		Accounts:  PermissionLevelWeightsToEOS(authority.Accounts),
		Waits:     WaitWeightsToEOS(authority.Waits),
	}
}

func WaitWeightsToDEOS(waits []zsw.WaitWeight) (out []*pbcodec.WaitWeight) {
	if len(waits) <= 0 {
		return nil
	}

	out = make([]*pbcodec.WaitWeight, len(waits))
	for i, o := range waits {
		out[i] = &pbcodec.WaitWeight{
			WaitSec: o.WaitSec,
			Weight:  uint32(o.Weight),
		}
	}
	return out
}

func WaitWeightsToEOS(waits []*pbcodec.WaitWeight) (out []zsw.WaitWeight) {
	if len(waits) <= 0 {
		return nil
	}

	out = make([]zsw.WaitWeight, len(waits))
	for i, o := range waits {
		out[i] = zsw.WaitWeight{
			WaitSec: o.WaitSec,
			Weight:  uint16(o.Weight),
		}
	}
	return out
}

func PermissionLevelWeightsToDEOS(weights []zsw.PermissionLevelWeight) (out []*pbcodec.PermissionLevelWeight) {
	if len(weights) <= 0 {
		return nil
	}

	out = make([]*pbcodec.PermissionLevelWeight, len(weights))
	for i, o := range weights {
		out[i] = &pbcodec.PermissionLevelWeight{
			Permission: PermissionLevelToDEOS(o.Permission),
			Weight:     uint32(o.Weight),
		}
	}
	return
}

func PermissionLevelWeightsToEOS(weights []*pbcodec.PermissionLevelWeight) (out []zsw.PermissionLevelWeight) {
	if len(weights) == 0 {
		return []zsw.PermissionLevelWeight{}
	}

	out = make([]zsw.PermissionLevelWeight, len(weights))
	for i, o := range weights {
		out[i] = zsw.PermissionLevelWeight{
			Permission: PermissionLevelToEOS(o.Permission),
			Weight:     uint16(o.Weight),
		}
	}
	return
}

func PermissionLevelToDEOS(perm zsw.PermissionLevel) *pbcodec.PermissionLevel {
	return &pbcodec.PermissionLevel{
		Actor:      string(perm.Actor),
		Permission: string(perm.Permission),
	}
}

func PermissionLevelToEOS(perm *pbcodec.PermissionLevel) zsw.PermissionLevel {
	return zsw.PermissionLevel{
		Actor:      zsw.AccountName(perm.Actor),
		Permission: zsw.PermissionName(perm.Permission),
	}
}

func KeyWeightsToDEOS(keys []zsw.KeyWeight) (out []*pbcodec.KeyWeight) {
	if len(keys) <= 0 {
		return nil
	}

	out = make([]*pbcodec.KeyWeight, len(keys))
	for i, o := range keys {
		out[i] = &pbcodec.KeyWeight{
			PublicKey: o.PublicKey.String(),
			Weight:    uint32(o.Weight),
		}
	}
	return
}

func KeyWeightsToEOS(keys []*pbcodec.KeyWeight) (out []zsw.KeyWeight) {
	if len(keys) <= 0 {
		return nil
	}

	out = make([]zsw.KeyWeight, len(keys))
	for i, o := range keys {
		out[i] = zsw.KeyWeight{
			PublicKey: ecc.MustNewPublicKey(o.PublicKey),
			Weight:    uint16(o.Weight),
		}
	}
	return

}

func KeyWeightsPToEOS(keys []*pbcodec.KeyWeight) (out []*zsw.KeyWeight) {
	if len(keys) <= 0 {
		return nil
	}

	out = make([]*zsw.KeyWeight, len(keys))
	for i, o := range keys {
		out[i] = &zsw.KeyWeight{
			PublicKey: ecc.MustNewPublicKey(o.PublicKey),
			Weight:    uint16(o.Weight),
		}
	}
	return

}

func KeyWeightsPToDEOS(keys []*zsw.KeyWeight) (out []*pbcodec.KeyWeight) {
	if len(keys) <= 0 {
		return nil
	}

	out = make([]*pbcodec.KeyWeight, len(keys))
	for i, o := range keys {
		out[i] = &pbcodec.KeyWeight{
			PublicKey: o.PublicKey.String(),
			Weight:    uint32(o.Weight),
		}
	}
	return
}

func TransactionToDEOS(trx *zsw.Transaction) *pbcodec.Transaction {
	var contextFreeActions []*pbcodec.Action
	if len(trx.ContextFreeActions) > 0 {
		contextFreeActions = make([]*pbcodec.Action, len(trx.ContextFreeActions))
		for i, act := range trx.ContextFreeActions {
			contextFreeActions[i] = ActionToDEOS(act)
		}
	}

	var actions []*pbcodec.Action
	if len(trx.Actions) > 0 {
		actions = make([]*pbcodec.Action, len(trx.Actions))
		for i, act := range trx.Actions {
			actions[i] = ActionToDEOS(act)
		}
	}

	return &pbcodec.Transaction{
		Header:             TransactionHeaderToDEOS(&trx.TransactionHeader),
		ContextFreeActions: contextFreeActions,
		Actions:            actions,
		Extensions:         ExtensionsToDEOS(trx.Extensions),
	}
}

func TransactionToEOS(trx *pbcodec.Transaction) *zsw.Transaction {
	var contextFreeActions []*zsw.Action
	if len(trx.ContextFreeActions) > 0 {
		contextFreeActions = make([]*zsw.Action, len(trx.ContextFreeActions))
		for i, act := range trx.ContextFreeActions {
			contextFreeActions[i] = ActionToEOS(act)
		}
	}

	var actions []*zsw.Action
	if len(trx.Actions) > 0 {
		actions = make([]*zsw.Action, len(trx.Actions))
		for i, act := range trx.Actions {
			actions[i] = ActionToEOS(act)
		}
	}

	return &zsw.Transaction{
		TransactionHeader:  *(TransactionHeaderToEOS(trx.Header)),
		ContextFreeActions: contextFreeActions,
		Actions:            actions,
		Extensions:         ExtensionsToEOS(trx.Extensions),
	}
}

func TransactionHeaderToDEOS(trx *zsw.TransactionHeader) *pbcodec.TransactionHeader {
	out := &pbcodec.TransactionHeader{
		Expiration:       mustProtoTimestamp(trx.Expiration.Time),
		RefBlockNum:      uint32(trx.RefBlockNum),
		RefBlockPrefix:   trx.RefBlockPrefix,
		MaxNetUsageWords: uint32(trx.MaxNetUsageWords),
		MaxCpuUsageMs:    uint32(trx.MaxCPUUsageMS),
		DelaySec:         uint32(trx.DelaySec),
	}

	return out
}

func TransactionHeaderToEOS(trx *pbcodec.TransactionHeader) *zsw.TransactionHeader {
	out := &zsw.TransactionHeader{
		Expiration:       TimestampToJSONTime(trx.Expiration),
		RefBlockNum:      uint16(trx.RefBlockNum),
		RefBlockPrefix:   uint32(trx.RefBlockPrefix),
		MaxNetUsageWords: zsw.Varuint32(trx.MaxNetUsageWords),
		MaxCPUUsageMS:    uint8(trx.MaxCpuUsageMs),
		DelaySec:         zsw.Varuint32(trx.DelaySec),
	}

	return out
}

func SignedTransactionToDEOS(trx *zsw.SignedTransaction) *pbcodec.SignedTransaction {
	return &pbcodec.SignedTransaction{
		Transaction:     TransactionToDEOS(trx.Transaction),
		Signatures:      SignaturesToDEOS(trx.Signatures),
		ContextFreeData: hexBytesToBytesSlices(trx.ContextFreeData),
	}
}

func SignedTransactionToEOS(trx *pbcodec.SignedTransaction) *zsw.SignedTransaction {
	return &zsw.SignedTransaction{
		Transaction:     TransactionToEOS(trx.Transaction),
		Signatures:      SignaturesToEOS(trx.Signatures),
		ContextFreeData: bytesSlicesToHexBytes(trx.ContextFreeData),
	}
}

func CreationTreeToDEOS(tree CreationFlatTree) []*pbcodec.CreationFlatNode {
	if len(tree) <= 0 {
		return nil
	}

	out := make([]*pbcodec.CreationFlatNode, len(tree))
	for i, node := range tree {
		out[i] = &pbcodec.CreationFlatNode{
			CreatorActionIndex:   int32(node[1]),
			ExecutionActionIndex: uint32(node[2]),
		}
	}
	return out
}

func ActionTracesToDEOS(actionTraces []zsw.ActionTrace, opts ...conversionOption) (out []*pbcodec.ActionTrace, someConsoleTruncated bool) {
	if len(actionTraces) <= 0 {
		return nil, false
	}

	sort.Slice(actionTraces, func(i, j int) bool {
		leftSeq := uint64(math.MaxUint64)
		rightSeq := uint64(math.MaxUint64)

		if leftReceipt := actionTraces[i].Receipt; leftReceipt != nil {
			if seq := leftReceipt.GlobalSequence; seq != 0 {
				leftSeq = uint64(seq)
			}
		}
		if rightReceipt := actionTraces[j].Receipt; rightReceipt != nil {
			if seq := rightReceipt.GlobalSequence; seq != 0 {
				rightSeq = uint64(seq)
			}
		}

		return leftSeq < rightSeq
	})

	out = make([]*pbcodec.ActionTrace, len(actionTraces))
	var consoleTruncated bool
	for idx, actionTrace := range actionTraces {
		out[idx], consoleTruncated = ActionTraceToDEOS(actionTrace, uint32(idx), opts...)
		if consoleTruncated {
			someConsoleTruncated = true
		}
	}

	return
}

func ActionTracesToEOS(actionTraces []*pbcodec.ActionTrace) (out []zsw.ActionTrace) {
	if len(actionTraces) <= 0 {
		return nil
	}

	out = make([]zsw.ActionTrace, len(actionTraces))
	for i, actionTrace := range actionTraces {
		out[i] = ActionTraceToEOS(actionTrace)
	}

	sort.Slice(out, func(i, j int) bool { return out[i].ActionOrdinal < out[j].ActionOrdinal })

	return
}

func AuthSequenceToDEOS(in zsw.TransactionTraceAuthSequence) *pbcodec.AuthSequence {
	return &pbcodec.AuthSequence{
		AccountName: string(in.Account),
		Sequence:    uint64(in.Sequence),
	}
}

func AuthSequenceListToEOS(in []*pbcodec.AuthSequence) (out []zsw.TransactionTraceAuthSequence) {
	if len(in) == 0 {
		return []zsw.TransactionTraceAuthSequence{}
	}

	out = make([]zsw.TransactionTraceAuthSequence, len(in))
	for i, seq := range in {
		out[i] = AuthSequenceToEOS(seq)
	}

	return
}

func AuthSequenceToEOS(in *pbcodec.AuthSequence) zsw.TransactionTraceAuthSequence {
	return zsw.TransactionTraceAuthSequence{
		Account:  zsw.AccountName(in.AccountName),
		Sequence: zsw.Uint64(in.Sequence),
	}
}

func ActionTraceToDEOS(in zsw.ActionTrace, execIndex uint32, opts ...conversionOption) (out *pbcodec.ActionTrace, consoleTruncated bool) {
	out = &pbcodec.ActionTrace{
		Receiver:             string(in.Receiver),
		Action:               ActionToDEOS(in.Action),
		Elapsed:              int64(in.Elapsed),
		Console:              string(in.Console),
		TransactionId:        in.TransactionID.String(),
		ContextFree:          in.ContextFree,
		ProducerBlockId:      in.ProducerBlockID.String(),
		BlockNum:             uint64(in.BlockNum),
		BlockTime:            mustProtoTimestamp(in.BlockTime.Time),
		AccountRamDeltas:     AccountRAMDeltasToDEOS(in.AccountRAMDeltas),
		Exception:            ExceptionToDEOS(in.Except),
		ActionOrdinal:        uint32(in.ActionOrdinal),
		CreatorActionOrdinal: uint32(in.CreatorActionOrdinal),
		ExecutionIndex:       execIndex,
		ErrorCode:            ErrorCodeToDEOS(in.ErrorCode),
	}
	out.ClosestUnnotifiedAncestorActionOrdinal = uint32(in.ClosestUnnotifiedAncestorActionOrdinal) // freaking long line, stay away from me

	if in.Receipt != nil {
		authSequences := in.Receipt.AuthSequence

		var deosAuthSequence []*pbcodec.AuthSequence
		if len(authSequences) > 0 {
			deosAuthSequence = make([]*pbcodec.AuthSequence, len(authSequences))
			for i, seq := range authSequences {
				deosAuthSequence[i] = AuthSequenceToDEOS(seq)
			}
		}

		out.Receipt = &pbcodec.ActionReceipt{
			Receiver:       string(in.Receipt.Receiver),
			Digest:         in.Receipt.ActionDigest.String(),
			GlobalSequence: uint64(in.Receipt.GlobalSequence),
			AuthSequence:   deosAuthSequence,
			RecvSequence:   uint64(in.Receipt.ReceiveSequence),
			CodeSequence:   uint64(in.Receipt.CodeSequence),
			AbiSequence:    uint64(in.Receipt.ABISequence),
		}
	}

	initialConsoleLength := len(in.Console)
	for _, opt := range opts {
		if v, ok := opt.(actionConversionOption); ok {
			v.apply(out)
		}
	}

	return out, initialConsoleLength != len(out.Console)
}

func ErrorCodeToDEOS(in *zsw.Uint64) uint64 {
	if in != nil {
		return uint64(*in)
	}
	return 0
}

func ErrorCodeToEOS(in uint64) *zsw.Uint64 {
	if in != 0 {
		val := zsw.Uint64(in)
		return &val
	}
	return nil
}

func ActionTraceToEOS(in *pbcodec.ActionTrace) (out zsw.ActionTrace) {
	out = zsw.ActionTrace{
		Receiver:             zsw.AccountName(in.Receiver),
		Action:               ActionToEOS(in.Action),
		Elapsed:              zsw.Int64(in.Elapsed),
		Console:              zsw.SafeString(in.Console),
		TransactionID:        ChecksumToEOS(in.TransactionId),
		ContextFree:          in.ContextFree,
		ProducerBlockID:      ChecksumToEOS(in.ProducerBlockId),
		BlockNum:             uint32(in.BlockNum),
		BlockTime:            TimestampToBlockTimestamp(in.BlockTime),
		AccountRAMDeltas:     AccountRAMDeltasToEOS(in.AccountRamDeltas),
		Except:               ExceptionToEOS(in.Exception),
		ActionOrdinal:        zsw.Varuint32(in.ActionOrdinal),
		CreatorActionOrdinal: zsw.Varuint32(in.CreatorActionOrdinal),
		ErrorCode:            ErrorCodeToEOS(in.ErrorCode),
	}
	out.ClosestUnnotifiedAncestorActionOrdinal = zsw.Varuint32(in.ClosestUnnotifiedAncestorActionOrdinal) // freaking long line, stay away from me

	if in.Receipt != nil {
		receipt := in.Receipt

		out.Receipt = &zsw.ActionTraceReceipt{
			Receiver:        zsw.AccountName(receipt.Receiver),
			ActionDigest:    ChecksumToEOS(receipt.Digest),
			GlobalSequence:  zsw.Uint64(receipt.GlobalSequence),
			AuthSequence:    AuthSequenceListToEOS(receipt.AuthSequence),
			ReceiveSequence: zsw.Uint64(receipt.RecvSequence),
			CodeSequence:    zsw.Varuint32(receipt.CodeSequence),
			ABISequence:     zsw.Varuint32(receipt.AbiSequence),
		}
	}

	return
}

func ChecksumToEOS(in string) zsw.Checksum256 {
	out, err := hex.DecodeString(in)
	if err != nil {
		panic(fmt.Sprintf("failed decoding checksum %q: %s", in, err))
	}

	return zsw.Checksum256(out)
}

func ActionToDEOS(action *zsw.Action) *pbcodec.Action {
	deosAction := &pbcodec.Action{
		Account:       string(action.Account),
		Name:          string(action.Name),
		Authorization: AuthorizationToDEOS(action.Authorization),
		RawData:       action.HexData,
	}

	if action.Data != nil {
		v, dataIsString := action.Data.(string)
		if dataIsString && len(action.HexData) == 0 {
			// When the action.Data is actually a string, and the HexData field is not set, we assume data sould be rawData instead
			rawData, err := hex.DecodeString(v)
			if err != nil {
				panic(fmt.Errorf("unable to unmarshal action data %q as hex: %s", v, err))
			}

			deosAction.RawData = rawData
		} else {
			serializedData, err := json.Marshal(action.Data)
			if err != nil {
				panic(fmt.Errorf("unable to unmarshal action data JSON: %s", err))
			}

			deosAction.JsonData = string(serializedData)
		}
	}

	return deosAction
}

func ActionToEOS(action *pbcodec.Action) (out *zsw.Action) {
	d := zsw.ActionData{}
	d.SetToServer(false) // rather, what we expect FROM `nodeos` servers

	d.HexData = zsw.HexBytes(action.RawData)
	if len(action.JsonData) != 0 {
		err := json.Unmarshal([]byte(action.JsonData), &d.Data)
		if err != nil {
			panic(fmt.Sprintf("unmarshaling action json data %q: %s", action.JsonData, err))
		}
	}

	out = &zsw.Action{
		Account:       zsw.AccountName(action.Account),
		Name:          zsw.ActionName(action.Name),
		Authorization: AuthorizationToEOS(action.Authorization),
		ActionData:    d,
	}

	return out
}

func AuthorizationToDEOS(authorization []zsw.PermissionLevel) (out []*pbcodec.PermissionLevel) {
	if len(authorization) <= 0 {
		return nil
	}

	out = make([]*pbcodec.PermissionLevel, len(authorization))
	for i, permission := range authorization {
		out[i] = PermissionLevelToDEOS(permission)
	}
	return
}

func AuthorizationToEOS(authorization []*pbcodec.PermissionLevel) (out []zsw.PermissionLevel) {
	if len(authorization) == 0 {
		return []zsw.PermissionLevel{}
	}

	out = make([]zsw.PermissionLevel, len(authorization))
	for i, permission := range authorization {
		out[i] = PermissionLevelToEOS(permission)
	}
	return
}

func AccountRAMDeltasToDEOS(deltas []*zsw.AccountRAMDelta) (out []*pbcodec.AccountRAMDelta) {
	if len(deltas) <= 0 {
		return nil
	}

	out = make([]*pbcodec.AccountRAMDelta, len(deltas))
	for i, delta := range deltas {
		out[i] = &pbcodec.AccountRAMDelta{
			Account: string(delta.Account),
			Delta:   int64(delta.Delta),
		}
	}
	return
}

func AccountRAMDeltasToEOS(deltas []*pbcodec.AccountRAMDelta) (out []*zsw.AccountRAMDelta) {
	if len(deltas) == 0 {
		return []*zsw.AccountRAMDelta{}
	}

	out = make([]*zsw.AccountRAMDelta, len(deltas))
	for i, delta := range deltas {
		out[i] = &zsw.AccountRAMDelta{
			Account: zsw.AccountName(delta.Account),
			Delta:   zsw.Int64(delta.Delta),
		}
	}
	return
}

func ExceptionToDEOS(in *zsw.Except) *pbcodec.Exception {
	if in == nil {
		return nil
	}
	out := &pbcodec.Exception{
		Code:    int32(in.Code),
		Name:    in.Name,
		Message: in.Message,
	}

	if len(in.Stack) > 0 {
		out.Stack = make([]*pbcodec.Exception_LogMessage, len(in.Stack))
		for i, el := range in.Stack {
			out.Stack[i] = &pbcodec.Exception_LogMessage{
				Context: LogContextToDEOS(el.Context),
				Format:  el.Format,
				Data:    el.Data,
			}
		}
	}

	return out
}

func ExceptionToEOS(in *pbcodec.Exception) *zsw.Except {
	if in == nil {
		return nil
	}
	out := &zsw.Except{
		Code:    zsw.Int64(in.Code),
		Name:    in.Name,
		Message: in.Message,
	}

	if len(in.Stack) > 0 {
		out.Stack = make([]*zsw.ExceptLogMessage, len(in.Stack))
		for i, el := range in.Stack {
			msg := &zsw.ExceptLogMessage{
				Format: el.Format,
			}

			ctx := LogContextToEOS(el.Context)
			if ctx != nil {
				msg.Context = *ctx
			}

			if len(el.Data) > 0 {
				msg.Data = json.RawMessage(el.Data)
			}

			out.Stack[i] = msg
		}
	}

	return out
}

func LogContextToDEOS(in zsw.ExceptLogContext) *pbcodec.Exception_LogContext {
	out := &pbcodec.Exception_LogContext{
		Level:      in.Level.String(),
		File:       in.File,
		Line:       int32(in.Line),
		Method:     in.Method,
		Hostname:   in.Hostname,
		ThreadName: in.ThreadName,
		Timestamp:  mustProtoTimestamp(in.Timestamp.Time),
	}
	if in.Context != nil {
		out.Context = LogContextToDEOS(*in.Context)
	}
	return out
}

func LogContextToEOS(in *pbcodec.Exception_LogContext) *zsw.ExceptLogContext {
	if in == nil {
		return nil
	}

	var exceptLevel zsw.ExceptLogLevel
	exceptLevel.FromString(in.Level)

	return &zsw.ExceptLogContext{
		Level:      exceptLevel,
		File:       in.File,
		Line:       uint64(in.Line),
		Method:     in.Method,
		Hostname:   in.Hostname,
		ThreadName: in.ThreadName,
		Timestamp:  TimestampToJSONTime(in.Timestamp),
		Context:    LogContextToEOS(in.Context),
	}
}

func TimestampToJSONTime(in *timestamp.Timestamp) zsw.JSONTime {
	out, _ := ptypes.Timestamp(in)
	return zsw.JSONTime{Time: out}
}

func TimestampToBlockTimestamp(in *timestamp.Timestamp) zsw.BlockTimestamp {
	out, _ := ptypes.Timestamp(in)
	return zsw.BlockTimestamp{Time: out}
}

func dbOpPathQuad(path string) (code string, scope string, table string, primaryKey string) {
	chunks := strings.Split(path, "/")
	if len(chunks) != 4 {
		panic("received db operation with a path with less than 4 '/'-separated chunks")
	}

	return chunks[0], chunks[1], chunks[2], chunks[3]
}

func tableOpPathQuad(path string) (code string, scope string, table string) {
	chunks := strings.Split(path, "/")
	if len(chunks) != 3 {
		panic("received db operation with a path with less than 3 '/'-separated chunks")
	}

	return chunks[0], chunks[1], chunks[2]
}

func mustProtoTimestamp(in time.Time) *timestamp.Timestamp {
	out, err := ptypes.TimestampProto(in)
	if err != nil {
		panic(fmt.Sprintf("invalid timestamp conversion %q: %s", in, err))
	}
	return out
}

func mustTimestamp(in *timestamp.Timestamp) time.Time {
	out, err := ptypes.Timestamp(in)
	if err != nil {
		panic(fmt.Sprintf("invalid timestamp conversion %q: %s", in, err))
	}
	return out
}

func TransactionStatusToDEOS(in zsw.TransactionStatus) pbcodec.TransactionStatus {
	switch in {
	case zsw.TransactionStatusExecuted:
		return pbcodec.TransactionStatus_TRANSACTIONSTATUS_EXECUTED
	case zsw.TransactionStatusSoftFail:
		return pbcodec.TransactionStatus_TRANSACTIONSTATUS_SOFTFAIL
	case zsw.TransactionStatusHardFail:
		return pbcodec.TransactionStatus_TRANSACTIONSTATUS_HARDFAIL
	case zsw.TransactionStatusDelayed:
		return pbcodec.TransactionStatus_TRANSACTIONSTATUS_DELAYED
	case zsw.TransactionStatusExpired:
		return pbcodec.TransactionStatus_TRANSACTIONSTATUS_EXPIRED
	default:
		return pbcodec.TransactionStatus_TRANSACTIONSTATUS_UNKNOWN
	}
}

func TransactionStatusToEOS(in pbcodec.TransactionStatus) zsw.TransactionStatus {
	switch in {
	case pbcodec.TransactionStatus_TRANSACTIONSTATUS_EXECUTED:
		return zsw.TransactionStatusExecuted
	case pbcodec.TransactionStatus_TRANSACTIONSTATUS_SOFTFAIL:
		return zsw.TransactionStatusSoftFail
	case pbcodec.TransactionStatus_TRANSACTIONSTATUS_HARDFAIL:
		return zsw.TransactionStatusHardFail
	case pbcodec.TransactionStatus_TRANSACTIONSTATUS_DELAYED:
		return zsw.TransactionStatusDelayed
	case pbcodec.TransactionStatus_TRANSACTIONSTATUS_EXPIRED:
		return zsw.TransactionStatusExpired
	default:
		return zsw.TransactionStatusUnknown
	}
}

func ExtractEOSSignedTransactionFromReceipt(trxReceipt *pbcodec.TransactionReceipt) (*zsw.SignedTransaction, error) {
	eosPackedTx, err := pbcodecPackedTransactionToEOS(trxReceipt.PackedTransaction)
	if err != nil {
		return nil, fmt.Errorf("pbcodec.PackedTransaction to EOS conversion failed: %s", err)
	}

	signedTransaction, err := eosPackedTx.UnpackBare()
	if err != nil {
		return nil, fmt.Errorf("unable to unpack packed transaction: %s", err)
	}

	return signedTransaction, nil
}

// Best effort to extract public keys from a signed transaction
func GetPublicKeysFromSignedTransaction(chainID zsw.Checksum256, signedTransaction *zsw.SignedTransaction) []string {
	eccPublicKeys, err := signedTransaction.SignedByKeys(chainID)
	if err != nil {
		// We discard any errors and simply return an empty array
		return nil
	}

	publicKeys := make([]string, len(eccPublicKeys))
	for i, eccPublicKey := range eccPublicKeys {
		publicKeys[i] = eccPublicKey.String()
	}

	return publicKeys
}

func pbcodecPackedTransactionToEOS(packedTrx *pbcodec.PackedTransaction) (*zsw.PackedTransaction, error) {
	signatures := make([]ecc.Signature, len(packedTrx.Signatures))
	for i, signature := range packedTrx.Signatures {
		eccSignature, err := ecc.NewSignature(signature)
		if err != nil {
			return nil, err
		}

		signatures[i] = eccSignature
	}

	return &zsw.PackedTransaction{
		Signatures:            signatures,
		Compression:           zsw.CompressionType(packedTrx.Compression),
		PackedContextFreeData: packedTrx.PackedContextFreeData,
		PackedTransaction:     packedTrx.PackedTransaction,
	}, nil
}
