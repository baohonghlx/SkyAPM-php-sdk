// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package v2

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

// In most cases, detect point should be `server` or `client`.
// Even in service mesh, this means `server`/`client` side sidecar
// `proxy` is reserved only.
type DetectPointV2 int32

const (
	DetectPointV2_client DetectPointV2 = 0
	DetectPointV2_server DetectPointV2 = 1
	DetectPointV2_proxy  DetectPointV2 = 2
)

var DetectPointV2_name = map[int32]string{
	0: "client",
	1: "server",
	2: "proxy",
}

var DetectPointV2_value = map[string]int32{
	"client": 0,
	"server": 1,
	"proxy":  2,
}

func (x DetectPointV2) String() string {
	return proto.EnumName(DetectPointV2_name, int32(x))
}

func (DetectPointV2) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

type KeyStringValuePairV2 struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyStringValuePairV2) Reset()         { *m = KeyStringValuePairV2{} }
func (m *KeyStringValuePairV2) String() string { return proto.CompactTextString(m) }
func (*KeyStringValuePairV2) ProtoMessage()    {}
func (*KeyStringValuePairV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *KeyStringValuePairV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyStringValuePairV2.Unmarshal(m, b)
}
func (m *KeyStringValuePairV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyStringValuePairV2.Marshal(b, m, deterministic)
}
func (m *KeyStringValuePairV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyStringValuePairV2.Merge(m, src)
}
func (m *KeyStringValuePairV2) XXX_Size() int {
	return xxx_messageInfo_KeyStringValuePairV2.Size(m)
}
func (m *KeyStringValuePairV2) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyStringValuePairV2.DiscardUnknown(m)
}

var xxx_messageInfo_KeyStringValuePairV2 proto.InternalMessageInfo

func (m *KeyStringValuePairV2) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyStringValuePairV2) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KeyIntValuePairV2 struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyIntValuePairV2) Reset()         { *m = KeyIntValuePairV2{} }
func (m *KeyIntValuePairV2) String() string { return proto.CompactTextString(m) }
func (*KeyIntValuePairV2) ProtoMessage()    {}
func (*KeyIntValuePairV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *KeyIntValuePairV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyIntValuePairV2.Unmarshal(m, b)
}
func (m *KeyIntValuePairV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyIntValuePairV2.Marshal(b, m, deterministic)
}
func (m *KeyIntValuePairV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyIntValuePairV2.Merge(m, src)
}
func (m *KeyIntValuePairV2) XXX_Size() int {
	return xxx_messageInfo_KeyIntValuePairV2.Size(m)
}
func (m *KeyIntValuePairV2) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyIntValuePairV2.DiscardUnknown(m)
}

var xxx_messageInfo_KeyIntValuePairV2 proto.InternalMessageInfo

func (m *KeyIntValuePairV2) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyIntValuePairV2) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type CPUV2 struct {
	UsagePercent         float64  `protobuf:"fixed64,2,opt,name=usagePercent,proto3" json:"usagePercent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPUV2) Reset()         { *m = CPUV2{} }
func (m *CPUV2) String() string { return proto.CompactTextString(m) }
func (*CPUV2) ProtoMessage()    {}
func (*CPUV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

func (m *CPUV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPUV2.Unmarshal(m, b)
}
func (m *CPUV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPUV2.Marshal(b, m, deterministic)
}
func (m *CPUV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPUV2.Merge(m, src)
}
func (m *CPUV2) XXX_Size() int {
	return xxx_messageInfo_CPUV2.Size(m)
}
func (m *CPUV2) XXX_DiscardUnknown() {
	xxx_messageInfo_CPUV2.DiscardUnknown(m)
}

var xxx_messageInfo_CPUV2 proto.InternalMessageInfo

func (m *CPUV2) GetUsagePercent() float64 {
	if m != nil {
		return m.UsagePercent
	}
	return 0
}

type CommandsV2 struct {
	Commands             []*CommandV2 `protobuf:"bytes,1,rep,name=commands,proto3" json:"commands,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CommandsV2) Reset()         { *m = CommandsV2{} }
func (m *CommandsV2) String() string { return proto.CompactTextString(m) }
func (*CommandsV2) ProtoMessage()    {}
func (*CommandsV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{3}
}

func (m *CommandsV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandsV2.Unmarshal(m, b)
}
func (m *CommandsV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandsV2.Marshal(b, m, deterministic)
}
func (m *CommandsV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandsV2.Merge(m, src)
}
func (m *CommandsV2) XXX_Size() int {
	return xxx_messageInfo_CommandsV2.Size(m)
}
func (m *CommandsV2) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandsV2.DiscardUnknown(m)
}

var xxx_messageInfo_CommandsV2 proto.InternalMessageInfo

func (m *CommandsV2) GetCommands() []*CommandV2 {
	if m != nil {
		return m.Commands
	}
	return nil
}

type CommandV2 struct {
	Command              string                  `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Args                 []*KeyStringValuePairV2 `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CommandV2) Reset()         { *m = CommandV2{} }
func (m *CommandV2) String() string { return proto.CompactTextString(m) }
func (*CommandV2) ProtoMessage()    {}
func (*CommandV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{4}
}

func (m *CommandV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandV2.Unmarshal(m, b)
}
func (m *CommandV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandV2.Marshal(b, m, deterministic)
}
func (m *CommandV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandV2.Merge(m, src)
}
func (m *CommandV2) XXX_Size() int {
	return xxx_messageInfo_CommandV2.Size(m)
}
func (m *CommandV2) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandV2.DiscardUnknown(m)
}

var xxx_messageInfo_CommandV2 proto.InternalMessageInfo

func (m *CommandV2) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *CommandV2) GetArgs() []*KeyStringValuePairV2 {
	if m != nil {
		return m.Args
	}
	return nil
}

func init() {
	proto.RegisterEnum("DetectPointV2", DetectPointV2_name, DetectPointV2_value)
	proto.RegisterType((*KeyStringValuePairV2)(nil), "KeyStringValuePairV2")
	proto.RegisterType((*KeyIntValuePairV2)(nil), "KeyIntValuePairV2")
	proto.RegisterType((*CPUV2)(nil), "CPUV2")
	proto.RegisterType((*CommandsV2)(nil), "CommandsV2")
	proto.RegisterType((*CommandV2)(nil), "CommandV2")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xff, 0xed, 0xfe, 0x9d, 0xee, 0x55, 0xa1, 0xc6, 0x09, 0x45, 0x3c, 0x8c, 0x1e, 0x64,
	0x2a, 0x74, 0x10, 0xbd, 0x09, 0x1e, 0x9c, 0x17, 0x19, 0x48, 0xe8, 0xb0, 0x82, 0xb7, 0x18, 0x5f,
	0x6a, 0x69, 0x9b, 0x8c, 0x34, 0xdb, 0xec, 0x57, 0xf2, 0x53, 0x4a, 0xdb, 0x6c, 0x22, 0x78, 0xf1,
	0x94, 0xf7, 0x7d, 0x7e, 0xcf, 0x13, 0x42, 0x1e, 0x38, 0x12, 0xaa, 0x2c, 0x95, 0x9c, 0x74, 0x47,
	0xb4, 0xd0, 0xca, 0xa8, 0xf0, 0x16, 0x86, 0x33, 0xac, 0xe7, 0x46, 0x67, 0x32, 0x4d, 0x78, 0xb1,
	0x44, 0xc6, 0x33, 0x9d, 0x50, 0xe2, 0x43, 0x2f, 0xc7, 0x3a, 0x70, 0x46, 0xce, 0x78, 0x10, 0x37,
	0x23, 0x19, 0x82, 0xb7, 0x6a, 0x0c, 0x81, 0xdb, 0x6a, 0xdd, 0x12, 0xde, 0xc0, 0xe1, 0x0c, 0xeb,
	0x07, 0x69, 0xfe, 0x10, 0xf6, 0x36, 0xe1, 0x4b, 0xf0, 0xa6, 0xec, 0x29, 0xa1, 0x24, 0x84, 0xfd,
	0x65, 0xc5, 0x53, 0x64, 0xa8, 0x05, 0x4a, 0xd3, 0xba, 0x9c, 0xf8, 0x87, 0x16, 0x5e, 0x03, 0x4c,
	0x55, 0x59, 0x72, 0xf9, 0x56, 0x25, 0x94, 0x9c, 0xc1, 0xae, 0xb0, 0x5b, 0xe0, 0x8c, 0x7a, 0xe3,
	0x3d, 0x0a, 0x91, 0xc5, 0x09, 0x8d, 0xb7, 0x2c, 0x64, 0x30, 0xd8, 0xca, 0x24, 0x80, 0x1d, 0x0b,
	0xec, 0xdb, 0x36, 0x2b, 0x39, 0x87, 0xff, 0x5c, 0xa7, 0x55, 0xe0, 0xb6, 0x57, 0x1d, 0x47, 0xbf,
	0xfd, 0x49, 0xdc, 0x5a, 0x2e, 0x28, 0x1c, 0xdc, 0xa3, 0x41, 0x61, 0x98, 0xca, 0xa4, 0x49, 0x28,
	0x01, 0xe8, 0x8b, 0x22, 0x43, 0x69, 0xfc, 0x7f, 0xcd, 0x5c, 0xa1, 0x5e, 0xa1, 0xf6, 0x1d, 0x32,
	0x00, 0x6f, 0xa1, 0xd5, 0x47, 0xed, 0xbb, 0x77, 0x29, 0x8c, 0x95, 0x4e, 0x23, 0xbe, 0xe0, 0xe2,
	0x1d, 0xa3, 0x2a, 0xaf, 0xd7, 0xbc, 0xc8, 0x33, 0xd9, 0x28, 0x65, 0x24, 0xd1, 0xac, 0x95, 0xce,
	0xa3, 0xae, 0x17, 0xe6, 0xbc, 0x9c, 0x7e, 0x1b, 0x26, 0x16, 0xda, 0xd2, 0x26, 0x2b, 0xfa, 0xe9,
	0x9e, 0xcc, 0xf3, 0xfa, 0xd9, 0xe6, 0x1f, 0x3b, 0xcc, 0x9a, 0x2e, 0x85, 0x2a, 0x5e, 0xfb, 0x6d,
	0xab, 0x57, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x84, 0x2a, 0x9c, 0x20, 0xec, 0x01, 0x00, 0x00,
}
