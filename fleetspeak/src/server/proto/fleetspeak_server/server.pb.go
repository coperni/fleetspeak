// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fleetspeak/src/server/proto/fleetspeak_server/server.proto

package fleetspeak_server

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Describes a server's configuration. If unset, all values default to values
// reasonable for a unit test or small installation. Larger installations may
// need to tune these.
type ServerConfig struct {
	// The collection of services that this server should include.
	Services []*ServiceConfig `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"`
	// The approximate time to wait between checking for new broadcasts. If unset,
	// a default of 1 minute is used.
	BroadcastPollTime    *duration.Duration `protobuf:"bytes,2,opt,name=broadcast_poll_time,json=broadcastPollTime,proto3" json:"broadcast_poll_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ServerConfig) Reset()         { *m = ServerConfig{} }
func (m *ServerConfig) String() string { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()    {}
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_server_2271f3b4a5c265fb, []int{0}
}
func (m *ServerConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServerConfig.Unmarshal(m, b)
}
func (m *ServerConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServerConfig.Marshal(b, m, deterministic)
}
func (dst *ServerConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerConfig.Merge(dst, src)
}
func (m *ServerConfig) XXX_Size() int {
	return xxx_messageInfo_ServerConfig.Size(m)
}
func (m *ServerConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ServerConfig proto.InternalMessageInfo

func (m *ServerConfig) GetServices() []*ServiceConfig {
	if m != nil {
		return m.Services
	}
	return nil
}

func (m *ServerConfig) GetBroadcastPollTime() *duration.Duration {
	if m != nil {
		return m.BroadcastPollTime
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerConfig)(nil), "fleetspeak.server.ServerConfig")
}

func init() {
	proto.RegisterFile("fleetspeak/src/server/proto/fleetspeak_server/server.proto", fileDescriptor_server_2271f3b4a5c265fb)
}

var fileDescriptor_server_2271f3b4a5c265fb = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4a, 0xcb, 0x49, 0x4d,
	0x2d, 0x29, 0x2e, 0x48, 0x4d, 0xcc, 0xd6, 0x2f, 0x2e, 0x4a, 0xd6, 0x2f, 0x4e, 0x2d, 0x2a, 0x4b,
	0x2d, 0xd2, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x47, 0xc8, 0xc5, 0x43, 0xc5, 0x21, 0x94, 0x1e,
	0x58, 0x5a, 0x48, 0x10, 0x21, 0xaf, 0x07, 0x91, 0x90, 0x92, 0x4b, 0xcf, 0xcf, 0x4f, 0xcf, 0x49,
	0x85, 0xe8, 0x4f, 0x2a, 0x4d, 0xd3, 0x4f, 0x29, 0x2d, 0x4a, 0x2c, 0xc9, 0xcc, 0xcf, 0x83, 0x68,
	0x91, 0xb2, 0x21, 0xdd, 0xba, 0xcc, 0xe4, 0xd4, 0x62, 0x88, 0x6e, 0xa5, 0xe9, 0x8c, 0x5c, 0x3c,
	0xc1, 0x60, 0x19, 0xe7, 0xfc, 0xbc, 0xb4, 0xcc, 0x74, 0x21, 0x1b, 0x2e, 0x0e, 0x98, 0x12, 0x09,
	0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x05, 0x3d, 0x0c, 0x47, 0xe9, 0x05, 0x43, 0x94, 0x40, 0xf4,
	0x04, 0xc1, 0x75, 0x08, 0x79, 0x72, 0x09, 0x27, 0x15, 0xe5, 0x27, 0xa6, 0x24, 0x27, 0x16, 0x97,
	0xc4, 0x17, 0xe4, 0xe7, 0xe4, 0xc4, 0x97, 0x64, 0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70,
	0x1b, 0x49, 0xea, 0x41, 0xbc, 0xa2, 0x07, 0xf3, 0x8a, 0x9e, 0x0b, 0xd4, 0x2b, 0x41, 0x82, 0x70,
	0x5d, 0x01, 0xf9, 0x39, 0x39, 0x21, 0x99, 0xb9, 0xa9, 0x49, 0x6c, 0x60, 0x55, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x0d, 0x23, 0xfe, 0xe7, 0x4f, 0x01, 0x00, 0x00,
}
