// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: set.proto

package pb

import (
	config "github.com/NordSecurity/nordvpn-linux/config"
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

type SetAutoconnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerTag            string          `protobuf:"bytes,4,opt,name=server_tag,json=serverTag,proto3" json:"server_tag,omitempty"`
	Protocol             config.Protocol `protobuf:"varint,5,opt,name=protocol,proto3,enum=config.Protocol" json:"protocol,omitempty"`
	ThreatProtectionLite bool            `protobuf:"varint,7,opt,name=threat_protection_lite,json=threatProtectionLite,proto3" json:"threat_protection_lite,omitempty"`
	Obfuscate            bool            `protobuf:"varint,8,opt,name=obfuscate,proto3" json:"obfuscate,omitempty"`
	AutoConnect          bool            `protobuf:"varint,9,opt,name=auto_connect,json=autoConnect,proto3" json:"auto_connect,omitempty"`
	Dns                  []string        `protobuf:"bytes,10,rep,name=dns,proto3" json:"dns,omitempty"`
	Whitelist            *Whitelist      `protobuf:"bytes,11,opt,name=whitelist,proto3" json:"whitelist,omitempty"`
}

func (x *SetAutoconnectRequest) Reset() {
	*x = SetAutoconnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_set_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAutoconnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAutoconnectRequest) ProtoMessage() {}

func (x *SetAutoconnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_set_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAutoconnectRequest.ProtoReflect.Descriptor instead.
func (*SetAutoconnectRequest) Descriptor() ([]byte, []int) {
	return file_set_proto_rawDescGZIP(), []int{0}
}

func (x *SetAutoconnectRequest) GetServerTag() string {
	if x != nil {
		return x.ServerTag
	}
	return ""
}

func (x *SetAutoconnectRequest) GetProtocol() config.Protocol {
	if x != nil {
		return x.Protocol
	}
	return config.Protocol(0)
}

func (x *SetAutoconnectRequest) GetThreatProtectionLite() bool {
	if x != nil {
		return x.ThreatProtectionLite
	}
	return false
}

func (x *SetAutoconnectRequest) GetObfuscate() bool {
	if x != nil {
		return x.Obfuscate
	}
	return false
}

func (x *SetAutoconnectRequest) GetAutoConnect() bool {
	if x != nil {
		return x.AutoConnect
	}
	return false
}

func (x *SetAutoconnectRequest) GetDns() []string {
	if x != nil {
		return x.Dns
	}
	return nil
}

func (x *SetAutoconnectRequest) GetWhitelist() *Whitelist {
	if x != nil {
		return x.Whitelist
	}
	return nil
}

type SetGenericRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *SetGenericRequest) Reset() {
	*x = SetGenericRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_set_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetGenericRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetGenericRequest) ProtoMessage() {}

func (x *SetGenericRequest) ProtoReflect() protoreflect.Message {
	mi := &file_set_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetGenericRequest.ProtoReflect.Descriptor instead.
func (*SetGenericRequest) Descriptor() ([]byte, []int) {
	return file_set_proto_rawDescGZIP(), []int{1}
}

func (x *SetGenericRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type SetUint32Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetUint32Request) Reset() {
	*x = SetUint32Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_set_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUint32Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUint32Request) ProtoMessage() {}

func (x *SetUint32Request) ProtoReflect() protoreflect.Message {
	mi := &file_set_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUint32Request.ProtoReflect.Descriptor instead.
func (*SetUint32Request) Descriptor() ([]byte, []int) {
	return file_set_proto_rawDescGZIP(), []int{2}
}

func (x *SetUint32Request) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type SetThreatProtectionLiteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThreatProtectionLite bool     `protobuf:"varint,2,opt,name=threat_protection_lite,json=threatProtectionLite,proto3" json:"threat_protection_lite,omitempty"`
	Dns                  []string `protobuf:"bytes,3,rep,name=dns,proto3" json:"dns,omitempty"`
}

func (x *SetThreatProtectionLiteRequest) Reset() {
	*x = SetThreatProtectionLiteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_set_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetThreatProtectionLiteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetThreatProtectionLiteRequest) ProtoMessage() {}

func (x *SetThreatProtectionLiteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_set_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetThreatProtectionLiteRequest.ProtoReflect.Descriptor instead.
func (*SetThreatProtectionLiteRequest) Descriptor() ([]byte, []int) {
	return file_set_proto_rawDescGZIP(), []int{3}
}

func (x *SetThreatProtectionLiteRequest) GetThreatProtectionLite() bool {
	if x != nil {
		return x.ThreatProtectionLite
	}
	return false
}

func (x *SetThreatProtectionLiteRequest) GetDns() []string {
	if x != nil {
		return x.Dns
	}
	return nil
}

type SetDNSRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dns                  []string `protobuf:"bytes,2,rep,name=dns,proto3" json:"dns,omitempty"`
	ThreatProtectionLite bool     `protobuf:"varint,3,opt,name=threat_protection_lite,json=threatProtectionLite,proto3" json:"threat_protection_lite,omitempty"`
}

func (x *SetDNSRequest) Reset() {
	*x = SetDNSRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_set_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetDNSRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetDNSRequest) ProtoMessage() {}

func (x *SetDNSRequest) ProtoReflect() protoreflect.Message {
	mi := &file_set_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetDNSRequest.ProtoReflect.Descriptor instead.
func (*SetDNSRequest) Descriptor() ([]byte, []int) {
	return file_set_proto_rawDescGZIP(), []int{4}
}

func (x *SetDNSRequest) GetDns() []string {
	if x != nil {
		return x.Dns
	}
	return nil
}

func (x *SetDNSRequest) GetThreatProtectionLite() bool {
	if x != nil {
		return x.ThreatProtectionLite
	}
	return false
}

type SetKillSwitchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KillSwitch bool       `protobuf:"varint,2,opt,name=kill_switch,json=killSwitch,proto3" json:"kill_switch,omitempty"`
	Whitelist  *Whitelist `protobuf:"bytes,3,opt,name=whitelist,proto3" json:"whitelist,omitempty"`
}

func (x *SetKillSwitchRequest) Reset() {
	*x = SetKillSwitchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_set_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetKillSwitchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetKillSwitchRequest) ProtoMessage() {}

func (x *SetKillSwitchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_set_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetKillSwitchRequest.ProtoReflect.Descriptor instead.
func (*SetKillSwitchRequest) Descriptor() ([]byte, []int) {
	return file_set_proto_rawDescGZIP(), []int{5}
}

func (x *SetKillSwitchRequest) GetKillSwitch() bool {
	if x != nil {
		return x.KillSwitch
	}
	return false
}

func (x *SetKillSwitchRequest) GetWhitelist() *Whitelist {
	if x != nil {
		return x.Whitelist
	}
	return nil
}

type SetNotifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid    int64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Notify bool  `protobuf:"varint,3,opt,name=notify,proto3" json:"notify,omitempty"`
}

func (x *SetNotifyRequest) Reset() {
	*x = SetNotifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_set_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetNotifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetNotifyRequest) ProtoMessage() {}

func (x *SetNotifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_set_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetNotifyRequest.ProtoReflect.Descriptor instead.
func (*SetNotifyRequest) Descriptor() ([]byte, []int) {
	return file_set_proto_rawDescGZIP(), []int{6}
}

func (x *SetNotifyRequest) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *SetNotifyRequest) GetNotify() bool {
	if x != nil {
		return x.Notify
	}
	return false
}

type SetProtocolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol config.Protocol `protobuf:"varint,2,opt,name=protocol,proto3,enum=config.Protocol" json:"protocol,omitempty"`
}

func (x *SetProtocolRequest) Reset() {
	*x = SetProtocolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_set_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetProtocolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetProtocolRequest) ProtoMessage() {}

func (x *SetProtocolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_set_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetProtocolRequest.ProtoReflect.Descriptor instead.
func (*SetProtocolRequest) Descriptor() ([]byte, []int) {
	return file_set_proto_rawDescGZIP(), []int{7}
}

func (x *SetProtocolRequest) GetProtocol() config.Protocol {
	if x != nil {
		return x.Protocol
	}
	return config.Protocol(0)
}

type SetTechnologyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Technology config.Technology `protobuf:"varint,2,opt,name=technology,proto3,enum=config.Technology" json:"technology,omitempty"`
}

func (x *SetTechnologyRequest) Reset() {
	*x = SetTechnologyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_set_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetTechnologyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetTechnologyRequest) ProtoMessage() {}

func (x *SetTechnologyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_set_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetTechnologyRequest.ProtoReflect.Descriptor instead.
func (*SetTechnologyRequest) Descriptor() ([]byte, []int) {
	return file_set_proto_rawDescGZIP(), []int{8}
}

func (x *SetTechnologyRequest) GetTechnology() config.Technology {
	if x != nil {
		return x.Technology
	}
	return config.Technology(0)
}

type SetWhitelistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Whitelist *Whitelist `protobuf:"bytes,2,opt,name=whitelist,proto3" json:"whitelist,omitempty"`
}

func (x *SetWhitelistRequest) Reset() {
	*x = SetWhitelistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_set_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetWhitelistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetWhitelistRequest) ProtoMessage() {}

func (x *SetWhitelistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_set_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetWhitelistRequest.ProtoReflect.Descriptor instead.
func (*SetWhitelistRequest) Descriptor() ([]byte, []int) {
	return file_set_proto_rawDescGZIP(), []int{9}
}

func (x *SetWhitelistRequest) GetWhitelist() *Whitelist {
	if x != nil {
		return x.Whitelist
	}
	return nil
}

var File_set_proto protoreflect.FileDescriptor

var file_set_proto_rawDesc = []byte{
	0x0a, 0x09, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a,
	0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x74, 0x65, 0x63,
	0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x02,
	0x0a, 0x15, 0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x54, 0x61, 0x67, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x34, 0x0a, 0x16, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x74, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x62,
	0x66, 0x75, 0x73, 0x63, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x6f,
	0x62, 0x66, 0x75, 0x73, 0x63, 0x61, 0x74, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x6f,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x61, 0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x6e, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6e, 0x73, 0x12, 0x2b, 0x0a,
	0x09, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x52,
	0x09, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x11, 0x53, 0x65,
	0x74, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x28, 0x0a, 0x10, 0x53, 0x65, 0x74,
	0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x22, 0x68, 0x0a, 0x1e, 0x53, 0x65, 0x74, 0x54, 0x68, 0x72, 0x65, 0x61, 0x74,
	0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x50, 0x72, 0x6f,
	0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6e, 0x73, 0x22, 0x57, 0x0a,
	0x0d, 0x53, 0x65, 0x74, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x64, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6e, 0x73,
	0x12, 0x34, 0x0a, 0x16, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x14, 0x74, 0x68, 0x72, 0x65, 0x61, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x74, 0x65, 0x22, 0x64, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x4b, 0x69, 0x6c,
	0x6c, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f,
	0x0a, 0x0b, 0x6b, 0x69, 0x6c, 0x6c, 0x5f, 0x73, 0x77, 0x69, 0x74, 0x63, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0a, 0x6b, 0x69, 0x6c, 0x6c, 0x53, 0x77, 0x69, 0x74, 0x63, 0x68, 0x12,
	0x2b, 0x0a, 0x09, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73,
	0x74, 0x52, 0x09, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x3c, 0x0a, 0x10,
	0x53, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x22, 0x42, 0x0a, 0x12, 0x53, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2c, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x4a,
	0x0a, 0x14, 0x53, 0x65, 0x74, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0a, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f,
	0x6c, 0x6f, 0x67, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x52, 0x0a,
	0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x22, 0x42, 0x0a, 0x13, 0x53, 0x65,
	0x74, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2b, 0x0a, 0x09, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x57, 0x68, 0x69, 0x74, 0x65, 0x6c,
	0x69, 0x73, 0x74, 0x52, 0x09, 0x77, 0x68, 0x69, 0x74, 0x65, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x31,
	0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x6f, 0x72,
	0x64, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x6e, 0x6f, 0x72, 0x64, 0x76, 0x70,
	0x6e, 0x2d, 0x6c, 0x69, 0x6e, 0x75, 0x78, 0x2f, 0x64, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_set_proto_rawDescOnce sync.Once
	file_set_proto_rawDescData = file_set_proto_rawDesc
)

func file_set_proto_rawDescGZIP() []byte {
	file_set_proto_rawDescOnce.Do(func() {
		file_set_proto_rawDescData = protoimpl.X.CompressGZIP(file_set_proto_rawDescData)
	})
	return file_set_proto_rawDescData
}

var file_set_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_set_proto_goTypes = []interface{}{
	(*SetAutoconnectRequest)(nil),          // 0: pb.SetAutoconnectRequest
	(*SetGenericRequest)(nil),              // 1: pb.SetGenericRequest
	(*SetUint32Request)(nil),               // 2: pb.SetUint32Request
	(*SetThreatProtectionLiteRequest)(nil), // 3: pb.SetThreatProtectionLiteRequest
	(*SetDNSRequest)(nil),                  // 4: pb.SetDNSRequest
	(*SetKillSwitchRequest)(nil),           // 5: pb.SetKillSwitchRequest
	(*SetNotifyRequest)(nil),               // 6: pb.SetNotifyRequest
	(*SetProtocolRequest)(nil),             // 7: pb.SetProtocolRequest
	(*SetTechnologyRequest)(nil),           // 8: pb.SetTechnologyRequest
	(*SetWhitelistRequest)(nil),            // 9: pb.SetWhitelistRequest
	(config.Protocol)(0),                   // 10: config.Protocol
	(*Whitelist)(nil),                      // 11: pb.Whitelist
	(config.Technology)(0),                 // 12: config.Technology
}
var file_set_proto_depIdxs = []int32{
	10, // 0: pb.SetAutoconnectRequest.protocol:type_name -> config.Protocol
	11, // 1: pb.SetAutoconnectRequest.whitelist:type_name -> pb.Whitelist
	11, // 2: pb.SetKillSwitchRequest.whitelist:type_name -> pb.Whitelist
	10, // 3: pb.SetProtocolRequest.protocol:type_name -> config.Protocol
	12, // 4: pb.SetTechnologyRequest.technology:type_name -> config.Technology
	11, // 5: pb.SetWhitelistRequest.whitelist:type_name -> pb.Whitelist
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_set_proto_init() }
func file_set_proto_init() {
	if File_set_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_set_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAutoconnectRequest); i {
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
		file_set_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetGenericRequest); i {
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
		file_set_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUint32Request); i {
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
		file_set_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetThreatProtectionLiteRequest); i {
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
		file_set_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetDNSRequest); i {
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
		file_set_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetKillSwitchRequest); i {
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
		file_set_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetNotifyRequest); i {
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
		file_set_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetProtocolRequest); i {
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
		file_set_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetTechnologyRequest); i {
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
		file_set_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetWhitelistRequest); i {
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
			RawDescriptor: file_set_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_set_proto_goTypes,
		DependencyIndexes: file_set_proto_depIdxs,
		MessageInfos:      file_set_proto_msgTypes,
	}.Build()
	File_set_proto = out.File
	file_set_proto_rawDesc = nil
	file_set_proto_goTypes = nil
	file_set_proto_depIdxs = nil
}
