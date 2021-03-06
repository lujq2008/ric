// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msgx.proto

package cflm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// e2ap pdu
type GrpcMsg struct {
	Mtype                int32    `protobuf:"varint,1,opt,name=Mtype,proto3" json:"Mtype,omitempty"`
	Payload              []byte   `protobuf:"bytes,2,opt,name=Payload,proto3" json:"Payload,omitempty"`
	PayloadLen           int32    `protobuf:"varint,3,opt,name=PayloadLen,proto3" json:"PayloadLen,omitempty"`
	Callid               int32    `protobuf:"varint,4,opt,name=Callid,proto3" json:"Callid,omitempty"`
	Timeout              int32    `protobuf:"varint,5,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	Status               int32    `protobuf:"varint,6,opt,name=status,proto3" json:"status,omitempty"`
	Meid                 string   `protobuf:"bytes,7,opt,name=Meid,proto3" json:"Meid,omitempty"`
	Topic                string   `protobuf:"bytes,8,opt,name=topic,proto3" json:"topic,omitempty"`
	SubId                int32    `protobuf:"varint,9,opt,name=SubId,proto3" json:"SubId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcMsg) Reset()         { *m = GrpcMsg{} }
func (m *GrpcMsg) String() string { return proto.CompactTextString(m) }
func (*GrpcMsg) ProtoMessage()    {}
func (*GrpcMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_msgx_1550c6ecc65185e6, []int{0}
}
func (m *GrpcMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcMsg.Unmarshal(m, b)
}
func (m *GrpcMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcMsg.Marshal(b, m, deterministic)
}
func (dst *GrpcMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcMsg.Merge(dst, src)
}
func (m *GrpcMsg) XXX_Size() int {
	return xxx_messageInfo_GrpcMsg.Size(m)
}
func (m *GrpcMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcMsg.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcMsg proto.InternalMessageInfo

func (m *GrpcMsg) GetMtype() int32 {
	if m != nil {
		return m.Mtype
	}
	return 0
}

func (m *GrpcMsg) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *GrpcMsg) GetPayloadLen() int32 {
	if m != nil {
		return m.PayloadLen
	}
	return 0
}

func (m *GrpcMsg) GetCallid() int32 {
	if m != nil {
		return m.Callid
	}
	return 0
}

func (m *GrpcMsg) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *GrpcMsg) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *GrpcMsg) GetMeid() string {
	if m != nil {
		return m.Meid
	}
	return ""
}

func (m *GrpcMsg) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *GrpcMsg) GetSubId() int32 {
	if m != nil {
		return m.SubId
	}
	return 0
}

type GrpcReply struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Err                  string   `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcReply) Reset()         { *m = GrpcReply{} }
func (m *GrpcReply) String() string { return proto.CompactTextString(m) }
func (*GrpcReply) ProtoMessage()    {}
func (*GrpcReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_msgx_1550c6ecc65185e6, []int{1}
}
func (m *GrpcReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GrpcReply.Unmarshal(m, b)
}
func (m *GrpcReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GrpcReply.Marshal(b, m, deterministic)
}
func (dst *GrpcReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcReply.Merge(dst, src)
}
func (m *GrpcReply) XXX_Size() int {
	return xxx_messageInfo_GrpcReply.Size(m)
}
func (m *GrpcReply) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcReply.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcReply proto.InternalMessageInfo

func (m *GrpcReply) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *GrpcReply) GetErr() string {
	if m != nil {
		return m.Err
	}
	return ""
}

func init() {
	proto.RegisterType((*GrpcMsg)(nil), "GrpcMsg")
	proto.RegisterType((*GrpcReply)(nil), "GrpcReply")
}

func init() { proto.RegisterFile("msgx.proto", fileDescriptor_msgx_1550c6ecc65185e6) }

var fileDescriptor_msgx_1550c6ecc65185e6 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x18, 0x84, 0x89, 0xbb, 0x6d, 0xcd, 0x8f, 0x07, 0x09, 0x22, 0xff, 0x49, 0xca, 0x9e, 0x7a, 0x12,
	0xc4, 0x47, 0xf0, 0x20, 0x82, 0x05, 0x89, 0xbe, 0x40, 0xb6, 0x09, 0x4b, 0x20, 0x35, 0x21, 0x49,
	0xc1, 0x3e, 0xaf, 0x2f, 0x22, 0xf9, 0xd3, 0x82, 0xb7, 0xf9, 0xe6, 0x67, 0x86, 0x4c, 0x00, 0xe6,
	0x74, 0xf9, 0x79, 0x0c, 0xd1, 0x67, 0x7f, 0xfa, 0x65, 0xd0, 0xbd, 0xc6, 0x30, 0x8d, 0xe9, 0x22,
	0xee, 0xa0, 0x19, 0xf3, 0x1a, 0x0c, 0xb2, 0x9e, 0x0d, 0x8d, 0xac, 0x20, 0x10, 0xba, 0x0f, 0xb5,
	0x3a, 0xaf, 0x34, 0x5e, 0xf5, 0x6c, 0xb8, 0x91, 0x3b, 0x8a, 0x07, 0x80, 0x4d, 0xbe, 0x9b, 0x6f,
	0x3c, 0x50, 0xe8, 0x9f, 0x23, 0xee, 0xa1, 0x7d, 0x51, 0xce, 0x59, 0x8d, 0x47, 0xba, 0x6d, 0x54,
	0x1a, 0xbf, 0xec, 0x6c, 0xfc, 0x92, 0xb1, 0xa1, 0xc3, 0x8e, 0x25, 0x91, 0xb2, 0xca, 0x4b, 0xc2,
	0xb6, 0x26, 0x2a, 0x09, 0x01, 0xc7, 0xd1, 0x58, 0x8d, 0x5d, 0xcf, 0x06, 0x2e, 0x49, 0x97, 0xd7,
	0x66, 0x1f, 0xec, 0x84, 0xd7, 0x64, 0x56, 0x28, 0xee, 0xe7, 0x72, 0x7e, 0xd3, 0xc8, 0xeb, 0x06,
	0x82, 0xd3, 0x13, 0xf0, 0x32, 0x52, 0x9a, 0xe0, 0xd6, 0x52, 0x36, 0x79, 0xbd, 0xaf, 0x24, 0x2d,
	0x6e, 0xe1, 0x60, 0x62, 0xa4, 0x81, 0x5c, 0x16, 0x79, 0x6e, 0xe9, 0x7f, 0x9e, 0xff, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x34, 0x85, 0x21, 0x52, 0x2d, 0x01, 0x00, 0x00,
}
