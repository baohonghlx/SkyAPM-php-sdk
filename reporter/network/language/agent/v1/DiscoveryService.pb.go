// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language-agent/DiscoveryService.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ApplicationInstance struct {
	ApplicationId        int32    `protobuf:"varint,1,opt,name=applicationId,proto3" json:"applicationId,omitempty"`
	AgentUUID            string   `protobuf:"bytes,2,opt,name=agentUUID,proto3" json:"agentUUID,omitempty"`
	RegisterTime         int64    `protobuf:"varint,3,opt,name=registerTime,proto3" json:"registerTime,omitempty"`
	Osinfo               *OSInfo  `protobuf:"bytes,4,opt,name=osinfo,proto3" json:"osinfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplicationInstance) Reset()         { *m = ApplicationInstance{} }
func (m *ApplicationInstance) String() string { return proto.CompactTextString(m) }
func (*ApplicationInstance) ProtoMessage()    {}
func (*ApplicationInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_59659d1d84d9b7a2, []int{0}
}

func (m *ApplicationInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationInstance.Unmarshal(m, b)
}
func (m *ApplicationInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationInstance.Marshal(b, m, deterministic)
}
func (m *ApplicationInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationInstance.Merge(m, src)
}
func (m *ApplicationInstance) XXX_Size() int {
	return xxx_messageInfo_ApplicationInstance.Size(m)
}
func (m *ApplicationInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationInstance.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationInstance proto.InternalMessageInfo

func (m *ApplicationInstance) GetApplicationId() int32 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *ApplicationInstance) GetAgentUUID() string {
	if m != nil {
		return m.AgentUUID
	}
	return ""
}

func (m *ApplicationInstance) GetRegisterTime() int64 {
	if m != nil {
		return m.RegisterTime
	}
	return 0
}

func (m *ApplicationInstance) GetOsinfo() *OSInfo {
	if m != nil {
		return m.Osinfo
	}
	return nil
}

type ApplicationInstanceMapping struct {
	ApplicationId         int32    `protobuf:"varint,1,opt,name=applicationId,proto3" json:"applicationId,omitempty"`
	ApplicationInstanceId int32    `protobuf:"varint,2,opt,name=applicationInstanceId,proto3" json:"applicationInstanceId,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ApplicationInstanceMapping) Reset()         { *m = ApplicationInstanceMapping{} }
func (m *ApplicationInstanceMapping) String() string { return proto.CompactTextString(m) }
func (*ApplicationInstanceMapping) ProtoMessage()    {}
func (*ApplicationInstanceMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_59659d1d84d9b7a2, []int{1}
}

func (m *ApplicationInstanceMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationInstanceMapping.Unmarshal(m, b)
}
func (m *ApplicationInstanceMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationInstanceMapping.Marshal(b, m, deterministic)
}
func (m *ApplicationInstanceMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationInstanceMapping.Merge(m, src)
}
func (m *ApplicationInstanceMapping) XXX_Size() int {
	return xxx_messageInfo_ApplicationInstanceMapping.Size(m)
}
func (m *ApplicationInstanceMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationInstanceMapping.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationInstanceMapping proto.InternalMessageInfo

func (m *ApplicationInstanceMapping) GetApplicationId() int32 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *ApplicationInstanceMapping) GetApplicationInstanceId() int32 {
	if m != nil {
		return m.ApplicationInstanceId
	}
	return 0
}

type ApplicationInstanceRecover struct {
	ApplicationId         int32    `protobuf:"varint,1,opt,name=applicationId,proto3" json:"applicationId,omitempty"`
	ApplicationInstanceId int32    `protobuf:"varint,2,opt,name=applicationInstanceId,proto3" json:"applicationInstanceId,omitempty"`
	RegisterTime          int64    `protobuf:"varint,3,opt,name=registerTime,proto3" json:"registerTime,omitempty"`
	Osinfo                *OSInfo  `protobuf:"bytes,4,opt,name=osinfo,proto3" json:"osinfo,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ApplicationInstanceRecover) Reset()         { *m = ApplicationInstanceRecover{} }
func (m *ApplicationInstanceRecover) String() string { return proto.CompactTextString(m) }
func (*ApplicationInstanceRecover) ProtoMessage()    {}
func (*ApplicationInstanceRecover) Descriptor() ([]byte, []int) {
	return fileDescriptor_59659d1d84d9b7a2, []int{2}
}

func (m *ApplicationInstanceRecover) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationInstanceRecover.Unmarshal(m, b)
}
func (m *ApplicationInstanceRecover) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationInstanceRecover.Marshal(b, m, deterministic)
}
func (m *ApplicationInstanceRecover) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationInstanceRecover.Merge(m, src)
}
func (m *ApplicationInstanceRecover) XXX_Size() int {
	return xxx_messageInfo_ApplicationInstanceRecover.Size(m)
}
func (m *ApplicationInstanceRecover) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationInstanceRecover.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationInstanceRecover proto.InternalMessageInfo

func (m *ApplicationInstanceRecover) GetApplicationId() int32 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *ApplicationInstanceRecover) GetApplicationInstanceId() int32 {
	if m != nil {
		return m.ApplicationInstanceId
	}
	return 0
}

func (m *ApplicationInstanceRecover) GetRegisterTime() int64 {
	if m != nil {
		return m.RegisterTime
	}
	return 0
}

func (m *ApplicationInstanceRecover) GetOsinfo() *OSInfo {
	if m != nil {
		return m.Osinfo
	}
	return nil
}

type ApplicationInstanceHeartbeat struct {
	ApplicationInstanceId int32    `protobuf:"varint,1,opt,name=applicationInstanceId,proto3" json:"applicationInstanceId,omitempty"`
	HeartbeatTime         int64    `protobuf:"varint,2,opt,name=heartbeatTime,proto3" json:"heartbeatTime,omitempty"`
	XXX_NoUnkeyedLiteral  struct{} `json:"-"`
	XXX_unrecognized      []byte   `json:"-"`
	XXX_sizecache         int32    `json:"-"`
}

func (m *ApplicationInstanceHeartbeat) Reset()         { *m = ApplicationInstanceHeartbeat{} }
func (m *ApplicationInstanceHeartbeat) String() string { return proto.CompactTextString(m) }
func (*ApplicationInstanceHeartbeat) ProtoMessage()    {}
func (*ApplicationInstanceHeartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_59659d1d84d9b7a2, []int{3}
}

func (m *ApplicationInstanceHeartbeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationInstanceHeartbeat.Unmarshal(m, b)
}
func (m *ApplicationInstanceHeartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationInstanceHeartbeat.Marshal(b, m, deterministic)
}
func (m *ApplicationInstanceHeartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationInstanceHeartbeat.Merge(m, src)
}
func (m *ApplicationInstanceHeartbeat) XXX_Size() int {
	return xxx_messageInfo_ApplicationInstanceHeartbeat.Size(m)
}
func (m *ApplicationInstanceHeartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationInstanceHeartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationInstanceHeartbeat proto.InternalMessageInfo

func (m *ApplicationInstanceHeartbeat) GetApplicationInstanceId() int32 {
	if m != nil {
		return m.ApplicationInstanceId
	}
	return 0
}

func (m *ApplicationInstanceHeartbeat) GetHeartbeatTime() int64 {
	if m != nil {
		return m.HeartbeatTime
	}
	return 0
}

type OSInfo struct {
	OsName               string   `protobuf:"bytes,1,opt,name=osName,proto3" json:"osName,omitempty"`
	Hostname             string   `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	ProcessNo            int32    `protobuf:"varint,3,opt,name=processNo,proto3" json:"processNo,omitempty"`
	Ipv4S                []string `protobuf:"bytes,4,rep,name=ipv4s,proto3" json:"ipv4s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OSInfo) Reset()         { *m = OSInfo{} }
func (m *OSInfo) String() string { return proto.CompactTextString(m) }
func (*OSInfo) ProtoMessage()    {}
func (*OSInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_59659d1d84d9b7a2, []int{4}
}

func (m *OSInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OSInfo.Unmarshal(m, b)
}
func (m *OSInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OSInfo.Marshal(b, m, deterministic)
}
func (m *OSInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OSInfo.Merge(m, src)
}
func (m *OSInfo) XXX_Size() int {
	return xxx_messageInfo_OSInfo.Size(m)
}
func (m *OSInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OSInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OSInfo proto.InternalMessageInfo

func (m *OSInfo) GetOsName() string {
	if m != nil {
		return m.OsName
	}
	return ""
}

func (m *OSInfo) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *OSInfo) GetProcessNo() int32 {
	if m != nil {
		return m.ProcessNo
	}
	return 0
}

func (m *OSInfo) GetIpv4S() []string {
	if m != nil {
		return m.Ipv4S
	}
	return nil
}

type ServiceNameCollection struct {
	Elements             []*ServiceNameElement `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ServiceNameCollection) Reset()         { *m = ServiceNameCollection{} }
func (m *ServiceNameCollection) String() string { return proto.CompactTextString(m) }
func (*ServiceNameCollection) ProtoMessage()    {}
func (*ServiceNameCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_59659d1d84d9b7a2, []int{5}
}

func (m *ServiceNameCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceNameCollection.Unmarshal(m, b)
}
func (m *ServiceNameCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceNameCollection.Marshal(b, m, deterministic)
}
func (m *ServiceNameCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceNameCollection.Merge(m, src)
}
func (m *ServiceNameCollection) XXX_Size() int {
	return xxx_messageInfo_ServiceNameCollection.Size(m)
}
func (m *ServiceNameCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceNameCollection.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceNameCollection proto.InternalMessageInfo

func (m *ServiceNameCollection) GetElements() []*ServiceNameElement {
	if m != nil {
		return m.Elements
	}
	return nil
}

type ServiceNameMappingCollection struct {
	Elements             []*ServiceNameMappingElement `protobuf:"bytes,1,rep,name=elements,proto3" json:"elements,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ServiceNameMappingCollection) Reset()         { *m = ServiceNameMappingCollection{} }
func (m *ServiceNameMappingCollection) String() string { return proto.CompactTextString(m) }
func (*ServiceNameMappingCollection) ProtoMessage()    {}
func (*ServiceNameMappingCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_59659d1d84d9b7a2, []int{6}
}

func (m *ServiceNameMappingCollection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceNameMappingCollection.Unmarshal(m, b)
}
func (m *ServiceNameMappingCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceNameMappingCollection.Marshal(b, m, deterministic)
}
func (m *ServiceNameMappingCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceNameMappingCollection.Merge(m, src)
}
func (m *ServiceNameMappingCollection) XXX_Size() int {
	return xxx_messageInfo_ServiceNameMappingCollection.Size(m)
}
func (m *ServiceNameMappingCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceNameMappingCollection.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceNameMappingCollection proto.InternalMessageInfo

func (m *ServiceNameMappingCollection) GetElements() []*ServiceNameMappingElement {
	if m != nil {
		return m.Elements
	}
	return nil
}

type ServiceNameMappingElement struct {
	ServiceId            int32               `protobuf:"varint,1,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	Element              *ServiceNameElement `protobuf:"bytes,2,opt,name=element,proto3" json:"element,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ServiceNameMappingElement) Reset()         { *m = ServiceNameMappingElement{} }
func (m *ServiceNameMappingElement) String() string { return proto.CompactTextString(m) }
func (*ServiceNameMappingElement) ProtoMessage()    {}
func (*ServiceNameMappingElement) Descriptor() ([]byte, []int) {
	return fileDescriptor_59659d1d84d9b7a2, []int{7}
}

func (m *ServiceNameMappingElement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceNameMappingElement.Unmarshal(m, b)
}
func (m *ServiceNameMappingElement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceNameMappingElement.Marshal(b, m, deterministic)
}
func (m *ServiceNameMappingElement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceNameMappingElement.Merge(m, src)
}
func (m *ServiceNameMappingElement) XXX_Size() int {
	return xxx_messageInfo_ServiceNameMappingElement.Size(m)
}
func (m *ServiceNameMappingElement) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceNameMappingElement.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceNameMappingElement proto.InternalMessageInfo

func (m *ServiceNameMappingElement) GetServiceId() int32 {
	if m != nil {
		return m.ServiceId
	}
	return 0
}

func (m *ServiceNameMappingElement) GetElement() *ServiceNameElement {
	if m != nil {
		return m.Element
	}
	return nil
}

type ServiceNameElement struct {
	ServiceName          string     `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	ApplicationId        int32      `protobuf:"varint,2,opt,name=applicationId,proto3" json:"applicationId,omitempty"`
	SrcSpanType          SpanTypeV1 `protobuf:"varint,3,opt,name=srcSpanType,proto3,enum=SpanTypeV1" json:"srcSpanType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ServiceNameElement) Reset()         { *m = ServiceNameElement{} }
func (m *ServiceNameElement) String() string { return proto.CompactTextString(m) }
func (*ServiceNameElement) ProtoMessage()    {}
func (*ServiceNameElement) Descriptor() ([]byte, []int) {
	return fileDescriptor_59659d1d84d9b7a2, []int{8}
}

func (m *ServiceNameElement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceNameElement.Unmarshal(m, b)
}
func (m *ServiceNameElement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceNameElement.Marshal(b, m, deterministic)
}
func (m *ServiceNameElement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceNameElement.Merge(m, src)
}
func (m *ServiceNameElement) XXX_Size() int {
	return xxx_messageInfo_ServiceNameElement.Size(m)
}
func (m *ServiceNameElement) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceNameElement.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceNameElement proto.InternalMessageInfo

func (m *ServiceNameElement) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *ServiceNameElement) GetApplicationId() int32 {
	if m != nil {
		return m.ApplicationId
	}
	return 0
}

func (m *ServiceNameElement) GetSrcSpanType() SpanTypeV1 {
	if m != nil {
		return m.SrcSpanType
	}
	return SpanTypeV1_Entry
}

func init() {
	proto.RegisterType((*ApplicationInstance)(nil), "ApplicationInstance")
	proto.RegisterType((*ApplicationInstanceMapping)(nil), "ApplicationInstanceMapping")
	proto.RegisterType((*ApplicationInstanceRecover)(nil), "ApplicationInstanceRecover")
	proto.RegisterType((*ApplicationInstanceHeartbeat)(nil), "ApplicationInstanceHeartbeat")
	proto.RegisterType((*OSInfo)(nil), "OSInfo")
	proto.RegisterType((*ServiceNameCollection)(nil), "ServiceNameCollection")
	proto.RegisterType((*ServiceNameMappingCollection)(nil), "ServiceNameMappingCollection")
	proto.RegisterType((*ServiceNameMappingElement)(nil), "ServiceNameMappingElement")
	proto.RegisterType((*ServiceNameElement)(nil), "ServiceNameElement")
}

func init() {
	proto.RegisterFile("language-agent/DiscoveryService.proto", fileDescriptor_59659d1d84d9b7a2)
}

var fileDescriptor_59659d1d84d9b7a2 = []byte{
	// 617 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdd, 0x4e, 0xdb, 0x3e,
	0x1c, 0x25, 0x2d, 0x05, 0xf2, 0xeb, 0x9f, 0xbf, 0x26, 0xf3, 0xa1, 0x10, 0x40, 0x44, 0x11, 0x93,
	0x7a, 0x53, 0x77, 0x74, 0x68, 0xd2, 0x2e, 0xb7, 0x31, 0x8d, 0x5e, 0x8c, 0xa1, 0x14, 0x98, 0xb4,
	0x3b, 0x63, 0x4c, 0x1b, 0x35, 0xb1, 0x23, 0xdb, 0x6b, 0xd7, 0xbd, 0xc2, 0xde, 0x81, 0x07, 0xd8,
	0x3b, 0xec, 0xdd, 0xa6, 0x7c, 0xb5, 0x81, 0xa4, 0xd5, 0x2e, 0xb6, 0x4b, 0x9f, 0xdf, 0xc9, 0xf1,
	0xb1, 0x7d, 0x72, 0xe0, 0x79, 0x40, 0xf8, 0xe0, 0x2b, 0x19, 0xb0, 0x36, 0x19, 0x30, 0xae, 0x3b,
	0x67, 0xbe, 0xa2, 0x62, 0xcc, 0xe4, 0xb4, 0xcf, 0xe4, 0xd8, 0xa7, 0x0c, 0x47, 0x52, 0x68, 0x61,
	0x1f, 0x3d, 0xa5, 0x89, 0x09, 0x57, 0x5a, 0x32, 0x12, 0x66, 0x84, 0x3d, 0x2a, 0xc2, 0x50, 0xf0,
	0x8e, 0x96, 0x84, 0xb2, 0x76, 0xba, 0x48, 0x47, 0xee, 0x83, 0x01, 0x5b, 0x6f, 0xa2, 0x28, 0xf0,
	0x29, 0xd1, 0xbe, 0xe0, 0x3d, 0xae, 0x34, 0xe1, 0x94, 0xa1, 0x63, 0xd8, 0x24, 0x05, 0xf8, 0xce,
	0x32, 0x1c, 0xa3, 0xd5, 0xf0, 0x1e, 0x83, 0xe8, 0x00, 0xcc, 0x64, 0xcb, 0xeb, 0xeb, 0xde, 0x99,
	0x55, 0x73, 0x8c, 0x96, 0xe9, 0xcd, 0x01, 0xe4, 0xc2, 0x7f, 0x92, 0x0d, 0x7c, 0xa5, 0x99, 0xbc,
	0xf2, 0x43, 0x66, 0xd5, 0x1d, 0xa3, 0x55, 0xf7, 0x1e, 0x61, 0xe8, 0x08, 0xd6, 0x84, 0xf2, 0xf9,
	0xbd, 0xb0, 0x56, 0x1d, 0xa3, 0xd5, 0xec, 0xae, 0xe3, 0x4f, 0xfd, 0x1e, 0xbf, 0x17, 0x5e, 0x06,
	0xbb, 0xdf, 0xc0, 0xae, 0xf0, 0xf7, 0x91, 0x44, 0x91, 0xcf, 0x07, 0x7f, 0x68, 0xf3, 0x14, 0x76,
	0x48, 0x59, 0xa3, 0x77, 0x97, 0x58, 0x6e, 0x78, 0xd5, 0x43, 0xf7, 0x97, 0x51, 0xb9, 0xb5, 0xc7,
	0x92, 0x37, 0xf8, 0x97, 0x5b, 0xff, 0x9d, 0x9b, 0xfb, 0x0e, 0x07, 0x15, 0xf6, 0xcf, 0x19, 0x91,
	0xfa, 0x96, 0x11, 0xbd, 0xd8, 0x9a, 0xb1, 0xcc, 0xda, 0x31, 0x6c, 0x0e, 0x73, 0x89, 0xc4, 0x5b,
	0x2d, 0xf1, 0xf6, 0x18, 0x74, 0x23, 0x58, 0x4b, 0xdd, 0xa0, 0xdd, 0xd8, 0xe6, 0x05, 0x09, 0x59,
	0x22, 0x6b, 0x7a, 0xd9, 0x0a, 0xd9, 0xb0, 0x31, 0x14, 0x4a, 0x73, 0x92, 0x49, 0x98, 0xde, 0x6c,
	0x1d, 0xc7, 0x2a, 0x92, 0x82, 0x32, 0xa5, 0x2e, 0x44, 0x72, 0xf6, 0x86, 0x37, 0x07, 0xd0, 0x36,
	0x34, 0xfc, 0x68, 0x7c, 0xaa, 0xac, 0x55, 0xa7, 0xde, 0x32, 0xbd, 0x74, 0xe1, 0x9e, 0xc3, 0x4e,
	0xf6, 0x57, 0xc4, 0xf2, 0xef, 0x44, 0x10, 0x30, 0x1a, 0x5b, 0x47, 0x1d, 0xd8, 0x60, 0x01, 0x0b,
	0x19, 0xd7, 0xca, 0x32, 0x9c, 0x7a, 0xab, 0xd9, 0xdd, 0xc2, 0x05, 0xe6, 0xfb, 0x74, 0xe6, 0xcd,
	0x48, 0xee, 0x0d, 0x1c, 0x14, 0xe6, 0x59, 0xd2, 0x0a, 0x82, 0xaf, 0x4a, 0x82, 0x36, 0x2e, 0x7f,
	0x50, 0xd6, 0x1d, 0xc2, 0xde, 0x42, 0x5a, 0x7c, 0x64, 0x95, 0x0e, 0x67, 0x0f, 0x30, 0x07, 0x50,
	0x1b, 0xd6, 0x33, 0x99, 0xe4, 0xae, 0x16, 0x1c, 0x21, 0xe7, 0xb8, 0x3f, 0x0c, 0x40, 0xe5, 0x39,
	0x72, 0xa0, 0xa9, 0xe6, 0x68, 0xf6, 0x1e, 0x45, 0xa8, 0x9c, 0xe9, 0x5a, 0x55, 0xa6, 0xdb, 0xd0,
	0x54, 0x92, 0xf6, 0x23, 0xc2, 0xaf, 0xa6, 0x51, 0x1a, 0xce, 0xff, 0xbb, 0x4d, 0x9c, 0x03, 0x37,
	0x27, 0x5e, 0x71, 0xde, 0x7d, 0x30, 0xc0, 0xca, 0x03, 0xf4, 0xb4, 0xc1, 0xd0, 0x07, 0x78, 0x96,
	0xa7, 0x7a, 0xd6, 0x3d, 0xdb, 0xb8, 0x22, 0xb7, 0xf6, 0x3e, 0x5e, 0xdc, 0x03, 0xee, 0x0a, 0x7a,
	0x0d, 0xe6, 0x2c, 0x82, 0xe8, 0x10, 0x2f, 0x4b, 0xbe, 0xdd, 0xc4, 0xf3, 0x8a, 0x74, 0x57, 0xba,
	0x14, 0xf6, 0x0b, 0xb7, 0x55, 0xb2, 0x78, 0x06, 0xe6, 0x5d, 0x8e, 0xa1, 0x5d, 0x5c, 0x99, 0x32,
	0xfb, 0x10, 0x2f, 0xcb, 0x8c, 0xbb, 0xf2, 0x76, 0x02, 0x2f, 0x84, 0x1c, 0x60, 0x12, 0x11, 0x3a,
	0x64, 0x58, 0x8d, 0xa6, 0x13, 0x12, 0x8c, 0x7c, 0x1e, 0x23, 0x21, 0xe6, 0x4c, 0x4f, 0x84, 0x1c,
	0xe1, 0xbc, 0xc8, 0x71, 0x52, 0xa2, 0x97, 0xc6, 0x97, 0xe3, 0x39, 0xb1, 0x93, 0x91, 0x3a, 0x39,
	0xa9, 0x93, 0xb6, 0xfd, 0xf8, 0xe4, 0x67, 0xcd, 0xee, 0x8f, 0xa6, 0x9f, 0x33, 0xbd, 0x8b, 0x94,
	0x76, 0x19, 0xd7, 0x3b, 0x15, 0xc1, 0xed, 0x5a, 0x52, 0xf4, 0x2f, 0x7f, 0x07, 0x00, 0x00, 0xff,
	0xff, 0xc1, 0x68, 0x7e, 0x01, 0x4d, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InstanceDiscoveryServiceClient is the client API for InstanceDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InstanceDiscoveryServiceClient interface {
	RegisterInstance(ctx context.Context, in *ApplicationInstance, opts ...grpc.CallOption) (*ApplicationInstanceMapping, error)
	Heartbeat(ctx context.Context, in *ApplicationInstanceHeartbeat, opts ...grpc.CallOption) (*Downstream, error)
}

type instanceDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewInstanceDiscoveryServiceClient(cc *grpc.ClientConn) InstanceDiscoveryServiceClient {
	return &instanceDiscoveryServiceClient{cc}
}

func (c *instanceDiscoveryServiceClient) RegisterInstance(ctx context.Context, in *ApplicationInstance, opts ...grpc.CallOption) (*ApplicationInstanceMapping, error) {
	out := new(ApplicationInstanceMapping)
	err := c.cc.Invoke(ctx, "/InstanceDiscoveryService/registerInstance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceDiscoveryServiceClient) Heartbeat(ctx context.Context, in *ApplicationInstanceHeartbeat, opts ...grpc.CallOption) (*Downstream, error) {
	out := new(Downstream)
	err := c.cc.Invoke(ctx, "/InstanceDiscoveryService/heartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstanceDiscoveryServiceServer is the server API for InstanceDiscoveryService service.
type InstanceDiscoveryServiceServer interface {
	RegisterInstance(context.Context, *ApplicationInstance) (*ApplicationInstanceMapping, error)
	Heartbeat(context.Context, *ApplicationInstanceHeartbeat) (*Downstream, error)
}

// UnimplementedInstanceDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedInstanceDiscoveryServiceServer struct {
}

func (*UnimplementedInstanceDiscoveryServiceServer) RegisterInstance(ctx context.Context, req *ApplicationInstance) (*ApplicationInstanceMapping, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterInstance not implemented")
}
func (*UnimplementedInstanceDiscoveryServiceServer) Heartbeat(ctx context.Context, req *ApplicationInstanceHeartbeat) (*Downstream, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Heartbeat not implemented")
}

func RegisterInstanceDiscoveryServiceServer(s *grpc.Server, srv InstanceDiscoveryServiceServer) {
	s.RegisterService(&_InstanceDiscoveryService_serviceDesc, srv)
}

func _InstanceDiscoveryService_RegisterInstance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationInstance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceDiscoveryServiceServer).RegisterInstance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InstanceDiscoveryService/RegisterInstance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceDiscoveryServiceServer).RegisterInstance(ctx, req.(*ApplicationInstance))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceDiscoveryService_Heartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplicationInstanceHeartbeat)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceDiscoveryServiceServer).Heartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/InstanceDiscoveryService/Heartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceDiscoveryServiceServer).Heartbeat(ctx, req.(*ApplicationInstanceHeartbeat))
	}
	return interceptor(ctx, in, info, handler)
}

var _InstanceDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "InstanceDiscoveryService",
	HandlerType: (*InstanceDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "registerInstance",
			Handler:    _InstanceDiscoveryService_RegisterInstance_Handler,
		},
		{
			MethodName: "heartbeat",
			Handler:    _InstanceDiscoveryService_Heartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language-agent/DiscoveryService.proto",
}

// ServiceNameDiscoveryServiceClient is the client API for ServiceNameDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceNameDiscoveryServiceClient interface {
	Discovery(ctx context.Context, in *ServiceNameCollection, opts ...grpc.CallOption) (*ServiceNameMappingCollection, error)
}

type serviceNameDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewServiceNameDiscoveryServiceClient(cc *grpc.ClientConn) ServiceNameDiscoveryServiceClient {
	return &serviceNameDiscoveryServiceClient{cc}
}

func (c *serviceNameDiscoveryServiceClient) Discovery(ctx context.Context, in *ServiceNameCollection, opts ...grpc.CallOption) (*ServiceNameMappingCollection, error) {
	out := new(ServiceNameMappingCollection)
	err := c.cc.Invoke(ctx, "/ServiceNameDiscoveryService/discovery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceNameDiscoveryServiceServer is the server API for ServiceNameDiscoveryService service.
type ServiceNameDiscoveryServiceServer interface {
	Discovery(context.Context, *ServiceNameCollection) (*ServiceNameMappingCollection, error)
}

// UnimplementedServiceNameDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceNameDiscoveryServiceServer struct {
}

func (*UnimplementedServiceNameDiscoveryServiceServer) Discovery(ctx context.Context, req *ServiceNameCollection) (*ServiceNameMappingCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Discovery not implemented")
}

func RegisterServiceNameDiscoveryServiceServer(s *grpc.Server, srv ServiceNameDiscoveryServiceServer) {
	s.RegisterService(&_ServiceNameDiscoveryService_serviceDesc, srv)
}

func _ServiceNameDiscoveryService_Discovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceNameCollection)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceNameDiscoveryServiceServer).Discovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ServiceNameDiscoveryService/Discovery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceNameDiscoveryServiceServer).Discovery(ctx, req.(*ServiceNameCollection))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceNameDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ServiceNameDiscoveryService",
	HandlerType: (*ServiceNameDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "discovery",
			Handler:    _ServiceNameDiscoveryService_Discovery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language-agent/DiscoveryService.proto",
}
