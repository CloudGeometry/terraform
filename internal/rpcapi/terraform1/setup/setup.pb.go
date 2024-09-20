// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.15.6
// source: setup.proto

package setup

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Handshake struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Handshake) Reset() {
	*x = Handshake{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handshake) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handshake) ProtoMessage() {}

func (x *Handshake) ProtoReflect() protoreflect.Message {
	mi := &file_setup_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handshake.ProtoReflect.Descriptor instead.
func (*Handshake) Descriptor() ([]byte, []int) {
	return file_setup_proto_rawDescGZIP(), []int{0}
}

type Stop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Stop) Reset() {
	*x = Stop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stop) ProtoMessage() {}

func (x *Stop) ProtoReflect() protoreflect.Message {
	mi := &file_setup_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stop.ProtoReflect.Descriptor instead.
func (*Stop) Descriptor() ([]byte, []int) {
	return file_setup_proto_rawDescGZIP(), []int{1}
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Credentials map[string]*HostCredential `protobuf:"bytes,1,rep,name=credentials,proto3" json:"credentials,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_setup_proto_msgTypes[2]
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
	return file_setup_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetCredentials() map[string]*HostCredential {
	if x != nil {
		return x.Credentials
	}
	return nil
}

type HostCredential struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *HostCredential) Reset() {
	*x = HostCredential{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setup_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostCredential) ProtoMessage() {}

func (x *HostCredential) ProtoReflect() protoreflect.Message {
	mi := &file_setup_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostCredential.ProtoReflect.Descriptor instead.
func (*HostCredential) Descriptor() ([]byte, []int) {
	return file_setup_proto_rawDescGZIP(), []int{3}
}

func (x *HostCredential) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

// The capabilities that the client wishes to advertise to the server during
// handshake.
type ClientCapabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ClientCapabilities) Reset() {
	*x = ClientCapabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setup_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientCapabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientCapabilities) ProtoMessage() {}

func (x *ClientCapabilities) ProtoReflect() protoreflect.Message {
	mi := &file_setup_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientCapabilities.ProtoReflect.Descriptor instead.
func (*ClientCapabilities) Descriptor() ([]byte, []int) {
	return file_setup_proto_rawDescGZIP(), []int{4}
}

// The capabilities that the server wishes to advertise to the client during
// handshake. Fields in this message can also be used to acknowledge and
// confirm support for client capabilities advertised in ClientCapabilities,
// in situations where the client must vary its behavior based on the server's
// level of support.
type ServerCapabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerCapabilities) Reset() {
	*x = ServerCapabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setup_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerCapabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerCapabilities) ProtoMessage() {}

func (x *ServerCapabilities) ProtoReflect() protoreflect.Message {
	mi := &file_setup_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerCapabilities.ProtoReflect.Descriptor instead.
func (*ServerCapabilities) Descriptor() ([]byte, []int) {
	return file_setup_proto_rawDescGZIP(), []int{5}
}

type Handshake_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capabilities *ClientCapabilities `protobuf:"bytes,1,opt,name=capabilities,proto3" json:"capabilities,omitempty"`
	Config       *Config             `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Handshake_Request) Reset() {
	*x = Handshake_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setup_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handshake_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handshake_Request) ProtoMessage() {}

func (x *Handshake_Request) ProtoReflect() protoreflect.Message {
	mi := &file_setup_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handshake_Request.ProtoReflect.Descriptor instead.
func (*Handshake_Request) Descriptor() ([]byte, []int) {
	return file_setup_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Handshake_Request) GetCapabilities() *ClientCapabilities {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

func (x *Handshake_Request) GetConfig() *Config {
	if x != nil {
		return x.Config
	}
	return nil
}

type Handshake_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capabilities *ServerCapabilities `protobuf:"bytes,2,opt,name=capabilities,proto3" json:"capabilities,omitempty"`
}

func (x *Handshake_Response) Reset() {
	*x = Handshake_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setup_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Handshake_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Handshake_Response) ProtoMessage() {}

func (x *Handshake_Response) ProtoReflect() protoreflect.Message {
	mi := &file_setup_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Handshake_Response.ProtoReflect.Descriptor instead.
func (*Handshake_Response) Descriptor() ([]byte, []int) {
	return file_setup_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Handshake_Response) GetCapabilities() *ServerCapabilities {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

type Stop_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Stop_Request) Reset() {
	*x = Stop_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setup_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stop_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stop_Request) ProtoMessage() {}

func (x *Stop_Request) ProtoReflect() protoreflect.Message {
	mi := &file_setup_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stop_Request.ProtoReflect.Descriptor instead.
func (*Stop_Request) Descriptor() ([]byte, []int) {
	return file_setup_proto_rawDescGZIP(), []int{1, 0}
}

type Stop_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Stop_Response) Reset() {
	*x = Stop_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_setup_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stop_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stop_Response) ProtoMessage() {}

func (x *Stop_Response) ProtoReflect() protoreflect.Message {
	mi := &file_setup_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stop_Response.ProtoReflect.Descriptor instead.
func (*Stop_Response) Descriptor() ([]byte, []int) {
	return file_setup_proto_rawDescGZIP(), []int{1, 1}
}

var File_setup_proto protoreflect.FileDescriptor

var file_setup_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x73, 0x65, 0x74, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x74,
	0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x73, 0x65, 0x74, 0x75, 0x70, 0x22,
	0xe9, 0x01, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x1a, 0x85, 0x01,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x48, 0x0a, 0x0c, 0x63, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x73, 0x65, 0x74,
	0x75, 0x70, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31,
	0x2e, 0x73, 0x65, 0x74, 0x75, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x54, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x48, 0x0a, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66,
	0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x73, 0x65, 0x74, 0x75, 0x70, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x0c, 0x63,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x22, 0x1d, 0x0a, 0x04, 0x53,
	0x74, 0x6f, 0x70, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xb7, 0x01, 0x0a, 0x06, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x4b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x74, 0x65, 0x72,
	0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x73, 0x65, 0x74, 0x75, 0x70, 0x2e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x1a, 0x60, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x36, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66,
	0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x73, 0x65, 0x74, 0x75, 0x70, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x26, 0x0a, 0x0e, 0x48, 0x6f, 0x73, 0x74, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x14, 0x0a, 0x12,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x32, 0xa8, 0x01, 0x0a, 0x05, 0x53, 0x65, 0x74,
	0x75, 0x70, 0x12, 0x56, 0x0a, 0x09, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x12,
	0x23, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e, 0x73, 0x65, 0x74,
	0x75, 0x70, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d,
	0x31, 0x2e, 0x73, 0x65, 0x74, 0x75, 0x70, 0x2e, 0x48, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x04, 0x53, 0x74,
	0x6f, 0x70, 0x12, 0x1e, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e,
	0x73, 0x65, 0x74, 0x75, 0x70, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x31, 0x2e,
	0x73, 0x65, 0x74, 0x75, 0x70, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_setup_proto_rawDescOnce sync.Once
	file_setup_proto_rawDescData = file_setup_proto_rawDesc
)

func file_setup_proto_rawDescGZIP() []byte {
	file_setup_proto_rawDescOnce.Do(func() {
		file_setup_proto_rawDescData = protoimpl.X.CompressGZIP(file_setup_proto_rawDescData)
	})
	return file_setup_proto_rawDescData
}

var file_setup_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_setup_proto_goTypes = []any{
	(*Handshake)(nil),          // 0: terraform1.setup.Handshake
	(*Stop)(nil),               // 1: terraform1.setup.Stop
	(*Config)(nil),             // 2: terraform1.setup.Config
	(*HostCredential)(nil),     // 3: terraform1.setup.HostCredential
	(*ClientCapabilities)(nil), // 4: terraform1.setup.ClientCapabilities
	(*ServerCapabilities)(nil), // 5: terraform1.setup.ServerCapabilities
	(*Handshake_Request)(nil),  // 6: terraform1.setup.Handshake.Request
	(*Handshake_Response)(nil), // 7: terraform1.setup.Handshake.Response
	(*Stop_Request)(nil),       // 8: terraform1.setup.Stop.Request
	(*Stop_Response)(nil),      // 9: terraform1.setup.Stop.Response
	nil,                        // 10: terraform1.setup.Config.CredentialsEntry
}
var file_setup_proto_depIdxs = []int32{
	10, // 0: terraform1.setup.Config.credentials:type_name -> terraform1.setup.Config.CredentialsEntry
	4,  // 1: terraform1.setup.Handshake.Request.capabilities:type_name -> terraform1.setup.ClientCapabilities
	2,  // 2: terraform1.setup.Handshake.Request.config:type_name -> terraform1.setup.Config
	5,  // 3: terraform1.setup.Handshake.Response.capabilities:type_name -> terraform1.setup.ServerCapabilities
	3,  // 4: terraform1.setup.Config.CredentialsEntry.value:type_name -> terraform1.setup.HostCredential
	6,  // 5: terraform1.setup.Setup.Handshake:input_type -> terraform1.setup.Handshake.Request
	8,  // 6: terraform1.setup.Setup.Stop:input_type -> terraform1.setup.Stop.Request
	7,  // 7: terraform1.setup.Setup.Handshake:output_type -> terraform1.setup.Handshake.Response
	9,  // 8: terraform1.setup.Setup.Stop:output_type -> terraform1.setup.Stop.Response
	7,  // [7:9] is the sub-list for method output_type
	5,  // [5:7] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_setup_proto_init() }
func file_setup_proto_init() {
	if File_setup_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_setup_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Handshake); i {
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
		file_setup_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Stop); i {
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
		file_setup_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
		file_setup_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*HostCredential); i {
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
		file_setup_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ClientCapabilities); i {
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
		file_setup_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ServerCapabilities); i {
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
		file_setup_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Handshake_Request); i {
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
		file_setup_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Handshake_Response); i {
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
		file_setup_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Stop_Request); i {
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
		file_setup_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*Stop_Response); i {
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
			RawDescriptor: file_setup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_setup_proto_goTypes,
		DependencyIndexes: file_setup_proto_depIdxs,
		MessageInfos:      file_setup_proto_msgTypes,
	}.Build()
	File_setup_proto = out.File
	file_setup_proto_rawDesc = nil
	file_setup_proto_goTypes = nil
	file_setup_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SetupClient is the client API for Setup service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SetupClient interface {
	// Clients must call Handshake before any other function of any other
	// service, to complete the capability negotiation step that may
	// then affect the behaviors of subsequent operations.
	//
	// This function can be called only once per RPC server.
	Handshake(ctx context.Context, in *Handshake_Request, opts ...grpc.CallOption) (*Handshake_Response, error)
	// At any time after handshaking, clients may call Stop to initiate a
	// graceful shutdown of the server.
	Stop(ctx context.Context, in *Stop_Request, opts ...grpc.CallOption) (*Stop_Response, error)
}

type setupClient struct {
	cc grpc.ClientConnInterface
}

func NewSetupClient(cc grpc.ClientConnInterface) SetupClient {
	return &setupClient{cc}
}

func (c *setupClient) Handshake(ctx context.Context, in *Handshake_Request, opts ...grpc.CallOption) (*Handshake_Response, error) {
	out := new(Handshake_Response)
	err := c.cc.Invoke(ctx, "/terraform1.setup.Setup/Handshake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *setupClient) Stop(ctx context.Context, in *Stop_Request, opts ...grpc.CallOption) (*Stop_Response, error) {
	out := new(Stop_Response)
	err := c.cc.Invoke(ctx, "/terraform1.setup.Setup/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SetupServer is the server API for Setup service.
type SetupServer interface {
	// Clients must call Handshake before any other function of any other
	// service, to complete the capability negotiation step that may
	// then affect the behaviors of subsequent operations.
	//
	// This function can be called only once per RPC server.
	Handshake(context.Context, *Handshake_Request) (*Handshake_Response, error)
	// At any time after handshaking, clients may call Stop to initiate a
	// graceful shutdown of the server.
	Stop(context.Context, *Stop_Request) (*Stop_Response, error)
}

// UnimplementedSetupServer can be embedded to have forward compatible implementations.
type UnimplementedSetupServer struct {
}

func (*UnimplementedSetupServer) Handshake(context.Context, *Handshake_Request) (*Handshake_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Handshake not implemented")
}
func (*UnimplementedSetupServer) Stop(context.Context, *Stop_Request) (*Stop_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}

func RegisterSetupServer(s *grpc.Server, srv SetupServer) {
	s.RegisterService(&_Setup_serviceDesc, srv)
}

func _Setup_Handshake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Handshake_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetupServer).Handshake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terraform1.setup.Setup/Handshake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetupServer).Handshake(ctx, req.(*Handshake_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Setup_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stop_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SetupServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/terraform1.setup.Setup/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SetupServer).Stop(ctx, req.(*Stop_Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Setup_serviceDesc = grpc.ServiceDesc{
	ServiceName: "terraform1.setup.Setup",
	HandlerType: (*SetupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Handshake",
			Handler:    _Setup_Handshake_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _Setup_Stop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "setup.proto",
}
