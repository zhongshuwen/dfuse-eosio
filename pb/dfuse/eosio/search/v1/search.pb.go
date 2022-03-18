// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dfuse/eosio/search/v1/search.proto

package pbsearcheos

import (
	fmt "fmt"
	v1 "github.com/zhongshuwen/dfuse-eosio/pb/dfuse/eosio/codec/v1"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DocumentID struct {
	BlockNum             uint64   `protobuf:"varint,1,opt,name=blockNum,proto3" json:"blockNum,omitempty"`
	ActionIndex          uint64   `protobuf:"varint,2,opt,name=actionIndex,proto3" json:"actionIndex,omitempty"`
	TransactionIndex     uint64   `protobuf:"varint,3,opt,name=transactionIndex,proto3" json:"transactionIndex,omitempty"`
	TransactionIDPrefix  []byte   `protobuf:"bytes,4,opt,name=transactionIDPrefix,proto3" json:"transactionIDPrefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DocumentID) Reset()         { *m = DocumentID{} }
func (m *DocumentID) String() string { return proto.CompactTextString(m) }
func (*DocumentID) ProtoMessage()    {}
func (*DocumentID) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6416b04c85aeead, []int{0}
}

func (m *DocumentID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DocumentID.Unmarshal(m, b)
}
func (m *DocumentID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DocumentID.Marshal(b, m, deterministic)
}
func (m *DocumentID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DocumentID.Merge(m, src)
}
func (m *DocumentID) XXX_Size() int {
	return xxx_messageInfo_DocumentID.Size(m)
}
func (m *DocumentID) XXX_DiscardUnknown() {
	xxx_messageInfo_DocumentID.DiscardUnknown(m)
}

var xxx_messageInfo_DocumentID proto.InternalMessageInfo

func (m *DocumentID) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *DocumentID) GetActionIndex() uint64 {
	if m != nil {
		return m.ActionIndex
	}
	return 0
}

func (m *DocumentID) GetTransactionIndex() uint64 {
	if m != nil {
		return m.TransactionIndex
	}
	return 0
}

func (m *DocumentID) GetTransactionIDPrefix() []byte {
	if m != nil {
		return m.TransactionIDPrefix
	}
	return nil
}

type Match struct {
	ActionIndexes        []uint32         `protobuf:"varint,1,rep,packed,name=actionIndexes,proto3" json:"actionIndexes,omitempty"`
	Block                *BlockTrxPayload `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Match) Reset()         { *m = Match{} }
func (m *Match) String() string { return proto.CompactTextString(m) }
func (*Match) ProtoMessage()    {}
func (*Match) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6416b04c85aeead, []int{1}
}

func (m *Match) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Match.Unmarshal(m, b)
}
func (m *Match) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Match.Marshal(b, m, deterministic)
}
func (m *Match) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Match.Merge(m, src)
}
func (m *Match) XXX_Size() int {
	return xxx_messageInfo_Match.Size(m)
}
func (m *Match) XXX_DiscardUnknown() {
	xxx_messageInfo_Match.DiscardUnknown(m)
}

var xxx_messageInfo_Match proto.InternalMessageInfo

func (m *Match) GetActionIndexes() []uint32 {
	if m != nil {
		return m.ActionIndexes
	}
	return nil
}

func (m *Match) GetBlock() *BlockTrxPayload {
	if m != nil {
		return m.Block
	}
	return nil
}

type BlockTrxPayload struct {
	BlockID              string               `protobuf:"bytes,1,opt,name=blockID,proto3" json:"blockID,omitempty"`
	BlockHeader          *v1.BlockHeader      `protobuf:"bytes,2,opt,name=blockHeader,proto3" json:"blockHeader,omitempty"`
	Trace                *v1.TransactionTrace `protobuf:"bytes,3,opt,name=trace,proto3" json:"trace,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *BlockTrxPayload) Reset()         { *m = BlockTrxPayload{} }
func (m *BlockTrxPayload) String() string { return proto.CompactTextString(m) }
func (*BlockTrxPayload) ProtoMessage()    {}
func (*BlockTrxPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6416b04c85aeead, []int{2}
}

func (m *BlockTrxPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockTrxPayload.Unmarshal(m, b)
}
func (m *BlockTrxPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockTrxPayload.Marshal(b, m, deterministic)
}
func (m *BlockTrxPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockTrxPayload.Merge(m, src)
}
func (m *BlockTrxPayload) XXX_Size() int {
	return xxx_messageInfo_BlockTrxPayload.Size(m)
}
func (m *BlockTrxPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockTrxPayload.DiscardUnknown(m)
}

var xxx_messageInfo_BlockTrxPayload proto.InternalMessageInfo

func (m *BlockTrxPayload) GetBlockID() string {
	if m != nil {
		return m.BlockID
	}
	return ""
}

func (m *BlockTrxPayload) GetBlockHeader() *v1.BlockHeader {
	if m != nil {
		return m.BlockHeader
	}
	return nil
}

func (m *BlockTrxPayload) GetTrace() *v1.TransactionTrace {
	if m != nil {
		return m.Trace
	}
	return nil
}

func init() {
	proto.RegisterType((*DocumentID)(nil), "dfuse.eosio.search.v1.DocumentID")
	proto.RegisterType((*Match)(nil), "dfuse.eosio.search.v1.Match")
	proto.RegisterType((*BlockTrxPayload)(nil), "dfuse.eosio.search.v1.BlockTrxPayload")
}

func init() { proto.RegisterFile("dfuse/eosio/search/v1/search.proto", fileDescriptor_f6416b04c85aeead) }

var fileDescriptor_f6416b04c85aeead = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x41, 0x4f, 0xf2, 0x40,
	0x14, 0x4c, 0x3f, 0xe0, 0x53, 0x1f, 0x12, 0xcd, 0x1a, 0x93, 0x86, 0x53, 0x6d, 0x8c, 0x21, 0x26,
	0xb6, 0xa2, 0x47, 0x3d, 0x61, 0x63, 0xe4, 0xa0, 0x21, 0x1b, 0x4e, 0xde, 0xb6, 0xdb, 0x45, 0x1a,
	0xa0, 0x4b, 0xb6, 0x5b, 0x82, 0xff, 0xc8, 0x93, 0xbf, 0xd1, 0xf4, 0x6d, 0xc1, 0x05, 0x7b, 0x9b,
	0x7d, 0x33, 0xb3, 0x6f, 0x32, 0x79, 0xe0, 0x27, 0x93, 0x22, 0x17, 0xa1, 0x90, 0x79, 0x2a, 0xc3,
	0x5c, 0x30, 0xc5, 0xa7, 0xe1, 0xaa, 0x5f, 0xa1, 0x60, 0xa9, 0xa4, 0x96, 0xe4, 0x1c, 0x35, 0x01,
	0x6a, 0x82, 0x8a, 0x59, 0xf5, 0xbb, 0x9e, 0x6d, 0xe5, 0x32, 0x11, 0xbc, 0x74, 0x22, 0x30, 0x46,
	0xff, 0xcb, 0x01, 0x88, 0x24, 0x2f, 0x16, 0x22, 0xd3, 0xc3, 0x88, 0x74, 0xe1, 0x30, 0x9e, 0x4b,
	0x3e, 0x7b, 0x2b, 0x16, 0xae, 0xe3, 0x39, 0xbd, 0x26, 0xdd, 0xbe, 0x89, 0x07, 0x6d, 0xc6, 0x75,
	0x2a, 0xb3, 0x61, 0x96, 0x88, 0xb5, 0xfb, 0x0f, 0x69, 0x7b, 0x44, 0xae, 0xe1, 0x54, 0x2b, 0x96,
	0xe5, 0xb6, 0xac, 0x81, 0xb2, 0x3f, 0x73, 0x72, 0x0b, 0x67, 0xf6, 0x2c, 0x1a, 0x29, 0x31, 0x49,
	0xd7, 0x6e, 0xd3, 0x73, 0x7a, 0xc7, 0xb4, 0x8e, 0xf2, 0x67, 0xd0, 0x7a, 0x65, 0x9a, 0x4f, 0xc9,
	0x25, 0x74, 0xac, 0x9f, 0x44, 0xee, 0x3a, 0x5e, 0xa3, 0xd7, 0xa1, 0xbb, 0x43, 0xf2, 0x08, 0x2d,
	0x8c, 0x8e, 0x41, 0xdb, 0x77, 0x57, 0x41, 0x6d, 0x45, 0xc1, 0xa0, 0xd4, 0x8c, 0xd5, 0x7a, 0xc4,
	0x3e, 0xe7, 0x92, 0x25, 0xd4, 0x98, 0xfc, 0x6f, 0x07, 0x4e, 0xf6, 0x28, 0xe2, 0xc2, 0x01, 0x92,
	0xc3, 0x08, 0xbb, 0x39, 0xa2, 0x9b, 0x27, 0x79, 0x82, 0x36, 0xc2, 0x17, 0xc1, 0x12, 0xa1, 0xaa,
	0x8d, 0x17, 0x3b, 0x1b, 0x4d, 0xe9, 0x9b, 0x85, 0x46, 0x48, 0x6d, 0x57, 0x19, 0x58, 0x2b, 0xc6,
	0x05, 0x56, 0xb6, 0x1f, 0x78, 0x6b, 0x1f, 0xff, 0x36, 0x33, 0x2e, 0xd5, 0xd4, 0x98, 0x06, 0xcf,
	0xef, 0xd1, 0x47, 0xaa, 0xa7, 0x45, 0x1c, 0x70, 0xb9, 0x08, 0xd1, 0x7a, 0x93, 0xca, 0x0a, 0x98,
	0x03, 0x58, 0xc6, 0x61, 0xed, 0x29, 0x3d, 0x2c, 0x63, 0x83, 0x85, 0xcc, 0xe3, 0xff, 0x78, 0x17,
	0xf7, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xa2, 0xc7, 0x16, 0x76, 0x02, 0x00, 0x00,
}
