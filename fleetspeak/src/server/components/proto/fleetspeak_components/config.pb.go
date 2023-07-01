// Copyright 2019 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: fleetspeak/src/server/components/proto/fleetspeak_components/config.proto

package fleetspeak_components

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Mysql connection string. Required.
	//
	// https://github.com/go-sql-driver/mysql#dsn-data-source-name
	MysqlDataSourceName string `protobuf:"bytes,1,opt,name=mysql_data_source_name,json=mysqlDataSourceName,proto3" json:"mysql_data_source_name,omitempty"`
	// The parameters required to stand up an https server.
	HttpsConfig *HttpsConfig `protobuf:"bytes,2,opt,name=https_config,json=httpsConfig,proto3" json:"https_config,omitempty"`
	// Parameters required to stand up an admin server. Either this or
	// "https_config", or both, have to be specified.
	AdminConfig *AdminConfig `protobuf:"bytes,7,opt,name=admin_config,json=adminConfig,proto3" json:"admin_config,omitempty"`
	// Parameters required to set up a stats collector.
	StatsConfig *StatsConfig `protobuf:"bytes,8,opt,name=stats_config,json=statsConfig,proto3" json:"stats_config,omitempty"`
	// Parameters required to stand up a http health check service. Optional.
	HealthCheckConfig *HealthCheckConfig `protobuf:"bytes,9,opt,name=health_check_config,json=healthCheckConfig,proto3" json:"health_check_config,omitempty"`
	// If set, expects connections to arrive through a load balance implementing
	// the PROXY protocol.
	//
	// https://www.haproxy.org/download/1.8/doc/proxy-protocol.txt
	ProxyProtocol bool `protobuf:"varint,3,opt,name=proxy_protocol,json=proxyProtocol,proto3" json:"proxy_protocol,omitempty"`
	// If set, only clients reporting this label will be allowed to
	// connect. Meant as a sanity check that the client and server are for the
	// same Fleetspeak installation.
	RequiredLabel string `protobuf:"bytes,4,opt,name=required_label,json=requiredLabel,proto3" json:"required_label,omitempty"`
	// If set, the bind address to listen on to receive notifications from other
	// fleetspeak servers. Optional, but strongly recommended for installations
	// involving multiple servers. e.g. ":8080", "localhost:1234".
	NotificationListenAddress string `protobuf:"bytes,5,opt,name=notification_listen_address,json=notificationListenAddress,proto3" json:"notification_listen_address,omitempty"`
	// If set, other servers will be told to use this address in order to connect
	// with this server's notification port. Has no effect when
	// notification_listen_address is unset.
	NotificationPublicAddress string `protobuf:"bytes,6,opt,name=notification_public_address,json=notificationPublicAddress,proto3" json:"notification_public_address,omitempty"`
	// If set, a HTTP notifier implementation is used for sending notifications.
	// Set this if running a pure admin server (without a notification listener)
	// in a distributed setup.
	NotificationUseHttpNotifier bool `protobuf:"varint,10,opt,name=notification_use_http_notifier,json=notificationUseHttpNotifier,proto3" json:"notification_use_http_notifier,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDescGZIP(), []int{0}
}

func (x *Config) GetMysqlDataSourceName() string {
	if x != nil {
		return x.MysqlDataSourceName
	}
	return ""
}

func (x *Config) GetHttpsConfig() *HttpsConfig {
	if x != nil {
		return x.HttpsConfig
	}
	return nil
}

func (x *Config) GetAdminConfig() *AdminConfig {
	if x != nil {
		return x.AdminConfig
	}
	return nil
}

func (x *Config) GetStatsConfig() *StatsConfig {
	if x != nil {
		return x.StatsConfig
	}
	return nil
}

func (x *Config) GetHealthCheckConfig() *HealthCheckConfig {
	if x != nil {
		return x.HealthCheckConfig
	}
	return nil
}

func (x *Config) GetProxyProtocol() bool {
	if x != nil {
		return x.ProxyProtocol
	}
	return false
}

func (x *Config) GetRequiredLabel() string {
	if x != nil {
		return x.RequiredLabel
	}
	return ""
}

func (x *Config) GetNotificationListenAddress() string {
	if x != nil {
		return x.NotificationListenAddress
	}
	return ""
}

func (x *Config) GetNotificationPublicAddress() string {
	if x != nil {
		return x.NotificationPublicAddress
	}
	return ""
}

func (x *Config) GetNotificationUseHttpNotifier() bool {
	if x != nil {
		return x.NotificationUseHttpNotifier
	}
	return false
}

type HttpsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bind address to listen on for client connections, e.g. ":443" or
	// "localhost:1234". Required.
	ListenAddress string `protobuf:"bytes,1,opt,name=listen_address,json=listenAddress,proto3" json:"listen_address,omitempty"`
	// A certificate chain which identifies the server to clients. Must lead to a
	// certificate known to the clients. x509 format. Required.
	Certificates string `protobuf:"bytes,2,opt,name=certificates,proto3" json:"certificates,omitempty"`
	// The private key used to identify the server. Must match the first entry in
	// certificates. x509 format. Required.
	Key string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	// If set, disables long running (streaming) connections. This type of
	// connection causes more active connections but can reduce database load and
	// server->client communications latency.
	DisableStreaming bool `protobuf:"varint,4,opt,name=disable_streaming,json=disableStreaming,proto3" json:"disable_streaming,omitempty"`
	// If set, the server will validate the client certificate from the request header.
	// This should be used if TLS is terminated at the load balancer and client certificates
	// can be passed upstream to the fleetspeak server as an http header.
	ClientCertificateHeader string `protobuf:"bytes,5,opt,name=client_certificate_header,json=clientCertificateHeader,proto3" json:"client_certificate_header,omitempty"`
}

func (x *HttpsConfig) Reset() {
	*x = HttpsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpsConfig) ProtoMessage() {}

func (x *HttpsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpsConfig.ProtoReflect.Descriptor instead.
func (*HttpsConfig) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDescGZIP(), []int{1}
}

func (x *HttpsConfig) GetListenAddress() string {
	if x != nil {
		return x.ListenAddress
	}
	return ""
}

func (x *HttpsConfig) GetCertificates() string {
	if x != nil {
		return x.Certificates
	}
	return ""
}

func (x *HttpsConfig) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *HttpsConfig) GetDisableStreaming() bool {
	if x != nil {
		return x.DisableStreaming
	}
	return false
}

func (x *HttpsConfig) GetClientCertificateHeader() string {
	if x != nil {
		return x.ClientCertificateHeader
	}
	return ""
}

type AdminConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bind address to listen on for connections, e.g. ":443" or
	// "localhost:1234". Required.
	ListenAddress string `protobuf:"bytes,1,opt,name=listen_address,json=listenAddress,proto3" json:"listen_address,omitempty"`
}

func (x *AdminConfig) Reset() {
	*x = AdminConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminConfig) ProtoMessage() {}

func (x *AdminConfig) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminConfig.ProtoReflect.Descriptor instead.
func (*AdminConfig) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDescGZIP(), []int{2}
}

func (x *AdminConfig) GetListenAddress() string {
	if x != nil {
		return x.ListenAddress
	}
	return ""
}

type StatsConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bind address to listen on for Prometheus http metric collection in the
	// form "<host>:<port>", e.g. "localhost:2112".
	// Optional; if no address is configured, then no stats collector
	// will be used (i.e. noopStatsCollector).
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *StatsConfig) Reset() {
	*x = StatsConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatsConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatsConfig) ProtoMessage() {}

func (x *StatsConfig) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatsConfig.ProtoReflect.Descriptor instead.
func (*StatsConfig) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDescGZIP(), []int{3}
}

func (x *StatsConfig) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type HealthCheckConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The bind address to listen on for http health check probes in the
	// form "<host>:<port>", e.g. "localhost:8080".
	ListenAddress string `protobuf:"bytes,1,opt,name=listen_address,json=listenAddress,proto3" json:"listen_address,omitempty"`
}

func (x *HealthCheckConfig) Reset() {
	*x = HealthCheckConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckConfig) ProtoMessage() {}

func (x *HealthCheckConfig) ProtoReflect() protoreflect.Message {
	mi := &file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckConfig.ProtoReflect.Descriptor instead.
func (*HealthCheckConfig) Descriptor() ([]byte, []int) {
	return file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDescGZIP(), []int{4}
}

func (x *HealthCheckConfig) GetListenAddress() string {
	if x != nil {
		return x.ListenAddress
	}
	return ""
}

var File_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto protoreflect.FileDescriptor

var file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDesc = []byte{
	0x0a, 0x49, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x73, 0x72, 0x63,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70,
	0x65, 0x61, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x66, 0x6c, 0x65,
	0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x22, 0xff, 0x04, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x33, 0x0a,
	0x16, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6d,
	0x79, 0x73, 0x71, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x48, 0x74, 0x74, 0x70, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x45, 0x0a, 0x0c, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x22, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x45, 0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70,
	0x65, 0x61, 0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x58, 0x0a, 0x13, 0x68, 0x65, 0x61, 0x6c, 0x74,
	0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61,
	0x6b, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x11,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x3e, 0x0a, 0x1b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x3e, 0x0a, 0x1b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x43, 0x0a, 0x1e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x75, 0x73, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1b, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x73, 0x65, 0x48, 0x74, 0x74, 0x70, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x22, 0xd3, 0x01, 0x0a, 0x0b, 0x48, 0x74, 0x74, 0x70, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69,
	0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x64, 0x69,
	0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x12, 0x3a,
	0x0a, 0x19, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x17, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x34, 0x0a, 0x0b, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x22, 0x27, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x3a, 0x0a, 0x11, 0x48, 0x65, 0x61,
	0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25,
	0x0a, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x5b, 0x5a, 0x59, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74,
	0x73, 0x70, 0x65, 0x61, 0x6b, 0x2f, 0x66, 0x6c, 0x65, 0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b,
	0x2f, 0x73, 0x72, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x65,
	0x65, 0x74, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDescOnce sync.Once
	file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDescData = file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDesc
)

func file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDescGZIP() []byte {
	file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDescOnce.Do(func() {
		file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDescData)
	})
	return file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDescData
}

var file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_goTypes = []interface{}{
	(*Config)(nil),            // 0: fleetspeak.components.Config
	(*HttpsConfig)(nil),       // 1: fleetspeak.components.HttpsConfig
	(*AdminConfig)(nil),       // 2: fleetspeak.components.AdminConfig
	(*StatsConfig)(nil),       // 3: fleetspeak.components.StatsConfig
	(*HealthCheckConfig)(nil), // 4: fleetspeak.components.HealthCheckConfig
}
var file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_depIdxs = []int32{
	1, // 0: fleetspeak.components.Config.https_config:type_name -> fleetspeak.components.HttpsConfig
	2, // 1: fleetspeak.components.Config.admin_config:type_name -> fleetspeak.components.AdminConfig
	3, // 2: fleetspeak.components.Config.stats_config:type_name -> fleetspeak.components.StatsConfig
	4, // 3: fleetspeak.components.Config.health_check_config:type_name -> fleetspeak.components.HealthCheckConfig
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_init() }
func file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_init() {
	if File_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpsConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatsConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_goTypes,
		DependencyIndexes: file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_depIdxs,
		MessageInfos:      file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_msgTypes,
	}.Build()
	File_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto = out.File
	file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_rawDesc = nil
	file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_goTypes = nil
	file_fleetspeak_src_server_components_proto_fleetspeak_components_config_proto_depIdxs = nil
}
