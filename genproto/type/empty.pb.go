// Code generated by protoc-gen-go. DO NOT EDIT.
// source: bolg/type/empty.proto

package _type

import (
	fmt "fmt"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_63afe88b2b5a7568, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Empty)(nil), "bolg.type.Empty")
}

func init() { proto.RegisterFile("bolg/type/empty.proto", fileDescriptor_63afe88b2b5a7568) }

var fileDescriptor_63afe88b2b5a7568 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xca, 0xcf, 0x49,
	0xd7, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0x4f, 0xcd, 0x2d, 0x28, 0xa9, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x04, 0x09, 0xeb, 0x81, 0x84, 0x95, 0xd8, 0xb9, 0x58, 0x5d, 0x41, 0x32, 0x4e,
	0xb1, 0x5c, 0x12, 0xc9, 0xf9, 0xb9, 0x7a, 0x20, 0x19, 0xdd, 0x94, 0xd4, 0xb2, 0xd4, 0x9c, 0xfc,
	0x82, 0xd4, 0xa2, 0x62, 0xb0, 0x22, 0x27, 0xce, 0x90, 0xca, 0x82, 0xd4, 0x00, 0x90, 0xd6, 0x00,
	0xc6, 0x28, 0x93, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x34, 0xd5,
	0xfa, 0x4e, 0xf9, 0x3e, 0xee, 0xba, 0xc1, 0xa9, 0x45, 0x65, 0xa9, 0x45, 0xfa, 0xe9, 0xa9, 0x79,
	0x60, 0xdb, 0xc0, 0xd6, 0x27, 0xb1, 0x81, 0xd9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee,
	0x4a, 0x72, 0x5d, 0x92, 0x00, 0x00, 0x00,
}