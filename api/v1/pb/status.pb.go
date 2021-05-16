// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.6.1
// source: status.proto

package pb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "github.com/gogo/googleapis/google/api"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ServerVersion provides server build and runtime version
type ServerVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Build is the server build version.
	Build string `protobuf:"bytes,1,opt,name=build,proto3" json:"build,omitempty"`
	// Runtime is the runtime version.
	Runtime string `protobuf:"bytes,2,opt,name=runtime,proto3" json:"runtime,omitempty"`
}

func (x *ServerVersion) Reset() {
	*x = ServerVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerVersion) ProtoMessage() {}

func (x *ServerVersion) ProtoReflect() protoreflect.Message {
	mi := &file_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerVersion.ProtoReflect.Descriptor instead.
func (*ServerVersion) Descriptor() ([]byte, []int) {
	return file_status_proto_rawDescGZIP(), []int{0}
}

func (x *ServerVersion) GetBuild() string {
	if x != nil {
		return x.Build
	}
	return ""
}

func (x *ServerVersion) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

// ServerStatus provides server status information
type ServerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the server or application.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Nodename is the human-readable name of the cluster member,
	// or empty for single host.
	Nodename string `protobuf:"bytes,2,opt,name=nodename,proto3" json:"nodename,omitempty"`
	// Hostname is operating system's host name.
	Hostname string `protobuf:"bytes,3,opt,name=hostname,proto3" json:"hostname,omitempty"`
	// ListenURLs is the list of URLs the service is listening on.
	ListenUrls []string `protobuf:"bytes,4,rep,name=listen_urls,proto3" json:"listen_urls,omitempty"`
	// StartedAt is the time when the server has started.
	StartedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=started_at,proto3" json:"started_at,omitempty"`
}

func (x *ServerStatus) Reset() {
	*x = ServerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_status_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatus) ProtoMessage() {}

func (x *ServerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_status_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStatus.ProtoReflect.Descriptor instead.
func (*ServerStatus) Descriptor() ([]byte, []int) {
	return file_status_proto_rawDescGZIP(), []int{1}
}

func (x *ServerStatus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServerStatus) GetNodename() string {
	if x != nil {
		return x.Nodename
	}
	return ""
}

func (x *ServerStatus) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *ServerStatus) GetListenUrls() []string {
	if x != nil {
		return x.ListenUrls
	}
	return nil
}

func (x *ServerStatus) GetStartedAt() *timestamp.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

// ServerStatusResponse returns status and version
type ServerStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status of the server.
	Status *ServerStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// Version of the server.
	Version *ServerVersion `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ServerStatusResponse) Reset() {
	*x = ServerStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_status_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerStatusResponse) ProtoMessage() {}

func (x *ServerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_status_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerStatusResponse.ProtoReflect.Descriptor instead.
func (*ServerStatusResponse) Descriptor() ([]byte, []int) {
	return file_status_proto_rawDescGZIP(), []int{2}
}

func (x *ServerStatusResponse) GetStatus() *ServerStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *ServerStatusResponse) GetVersion() *ServerVersion {
	if x != nil {
		return x.Version
	}
	return nil
}

// CallerStatusResponse returns the caller information
type CallerStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of the caller.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the caller.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Role of the caller.
	Role string `protobuf:"bytes,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *CallerStatusResponse) Reset() {
	*x = CallerStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_status_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallerStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallerStatusResponse) ProtoMessage() {}

func (x *CallerStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_status_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallerStatusResponse.ProtoReflect.Descriptor instead.
func (*CallerStatusResponse) Descriptor() ([]byte, []int) {
	return file_status_proto_rawDescGZIP(), []int{3}
}

func (x *CallerStatusResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CallerStatusResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CallerStatusResponse) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

var File_status_proto protoreflect.FileDescriptor

var file_status_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f,
	0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x22,
	0xb8, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x3a,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x6d, 0x0a, 0x14, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2b, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x4e, 0x0a, 0x14, 0x43, 0x61, 0x6c,
	0x6c, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x32, 0x8f, 0x02, 0x0a, 0x0d, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x07, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a,
	0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x12, 0x55, 0x0a, 0x06, 0x43, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x61, 0x6c, 0x6c,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2f, 0x63, 0x61, 0x6c, 0x6c, 0x65, 0x72, 0x42, 0xc5, 0x01, 0x5a, 0x23,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6b, 0x73, 0x70, 0x61,
	0x6e, 0x64, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x79, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x62, 0x92, 0x41, 0x9c, 0x01, 0x12, 0x73, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x20, 0x41, 0x50, 0x49, 0x22, 0x3e, 0x0a, 0x06, 0x54, 0x72, 0x75, 0x73, 0x74, 0x79, 0x12,
	0x21, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6b, 0x73, 0x70, 0x61, 0x6e, 0x64, 0x2f, 0x74, 0x72, 0x75, 0x73,
	0x74, 0x79, 0x1a, 0x11, 0x64, 0x65, 0x6e, 0x69, 0x73, 0x40, 0x65, 0x6b, 0x73, 0x70, 0x61, 0x6e,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x3a, 0x20, 0x0a, 0x15, 0x78, 0x2d,
	0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x6f, 0x6d, 0x65, 0x74, 0x68,
	0x69, 0x6e, 0x67, 0x12, 0x07, 0x1a, 0x05, 0x79, 0x61, 0x64, 0x64, 0x61, 0x2a, 0x01, 0x01, 0x32,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_status_proto_rawDescOnce sync.Once
	file_status_proto_rawDescData = file_status_proto_rawDesc
)

func file_status_proto_rawDescGZIP() []byte {
	file_status_proto_rawDescOnce.Do(func() {
		file_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_status_proto_rawDescData)
	})
	return file_status_proto_rawDescData
}

var file_status_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_status_proto_goTypes = []interface{}{
	(*ServerVersion)(nil),        // 0: pb.ServerVersion
	(*ServerStatus)(nil),         // 1: pb.ServerStatus
	(*ServerStatusResponse)(nil), // 2: pb.ServerStatusResponse
	(*CallerStatusResponse)(nil), // 3: pb.CallerStatusResponse
	(*timestamp.Timestamp)(nil),  // 4: google.protobuf.Timestamp
	(*empty.Empty)(nil),          // 5: google.protobuf.Empty
}
var file_status_proto_depIdxs = []int32{
	4, // 0: pb.ServerStatus.started_at:type_name -> google.protobuf.Timestamp
	1, // 1: pb.ServerStatusResponse.status:type_name -> pb.ServerStatus
	0, // 2: pb.ServerStatusResponse.version:type_name -> pb.ServerVersion
	5, // 3: pb.StatusService.Version:input_type -> google.protobuf.Empty
	5, // 4: pb.StatusService.Server:input_type -> google.protobuf.Empty
	5, // 5: pb.StatusService.Caller:input_type -> google.protobuf.Empty
	0, // 6: pb.StatusService.Version:output_type -> pb.ServerVersion
	2, // 7: pb.StatusService.Server:output_type -> pb.ServerStatusResponse
	3, // 8: pb.StatusService.Caller:output_type -> pb.CallerStatusResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_status_proto_init() }
func file_status_proto_init() {
	if File_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerVersion); i {
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
		file_status_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStatus); i {
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
		file_status_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerStatusResponse); i {
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
		file_status_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallerStatusResponse); i {
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
			RawDescriptor: file_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_status_proto_goTypes,
		DependencyIndexes: file_status_proto_depIdxs,
		MessageInfos:      file_status_proto_msgTypes,
	}.Build()
	File_status_proto = out.File
	file_status_proto_rawDesc = nil
	file_status_proto_goTypes = nil
	file_status_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// StatusServiceClient is the client API for StatusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatusServiceClient interface {
	// Version returns the server version.
	Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ServerVersion, error)
	// Server returns the server status.
	Server(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ServerStatusResponse, error)
	// Caller returns the caller status.
	Caller(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CallerStatusResponse, error)
}

type statusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStatusServiceClient(cc grpc.ClientConnInterface) StatusServiceClient {
	return &statusServiceClient{cc}
}

func (c *statusServiceClient) Version(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ServerVersion, error) {
	out := new(ServerVersion)
	err := c.cc.Invoke(ctx, "/pb.StatusService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statusServiceClient) Server(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ServerStatusResponse, error) {
	out := new(ServerStatusResponse)
	err := c.cc.Invoke(ctx, "/pb.StatusService/Server", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statusServiceClient) Caller(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*CallerStatusResponse, error) {
	out := new(CallerStatusResponse)
	err := c.cc.Invoke(ctx, "/pb.StatusService/Caller", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatusServiceServer is the server API for StatusService service.
type StatusServiceServer interface {
	// Version returns the server version.
	Version(context.Context, *empty.Empty) (*ServerVersion, error)
	// Server returns the server status.
	Server(context.Context, *empty.Empty) (*ServerStatusResponse, error)
	// Caller returns the caller status.
	Caller(context.Context, *empty.Empty) (*CallerStatusResponse, error)
}

// UnimplementedStatusServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStatusServiceServer struct {
}

func (*UnimplementedStatusServiceServer) Version(context.Context, *empty.Empty) (*ServerVersion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (*UnimplementedStatusServiceServer) Server(context.Context, *empty.Empty) (*ServerStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Server not implemented")
}
func (*UnimplementedStatusServiceServer) Caller(context.Context, *empty.Empty) (*CallerStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Caller not implemented")
}

func RegisterStatusServiceServer(s *grpc.Server, srv StatusServiceServer) {
	s.RegisterService(&_StatusService_serviceDesc, srv)
}

func _StatusService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StatusService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServiceServer).Version(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatusService_Server_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServiceServer).Server(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StatusService/Server",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServiceServer).Server(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _StatusService_Caller_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatusServiceServer).Caller(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.StatusService/Caller",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatusServiceServer).Caller(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _StatusService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.StatusService",
	HandlerType: (*StatusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _StatusService_Version_Handler,
		},
		{
			MethodName: "Server",
			Handler:    _StatusService_Server_Handler,
		},
		{
			MethodName: "Caller",
			Handler:    _StatusService_Caller_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "status.proto",
}