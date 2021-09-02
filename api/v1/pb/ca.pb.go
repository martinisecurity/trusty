// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.6.1
// source: ca.proto

package pb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "github.com/gogo/googleapis/google/api"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type CertProfileInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label   string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Profile string `protobuf:"bytes,2,opt,name=profile,proto3" json:"profile,omitempty"`
}

func (x *CertProfileInfoRequest) Reset() {
	*x = CertProfileInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertProfileInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertProfileInfoRequest) ProtoMessage() {}

func (x *CertProfileInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertProfileInfoRequest.ProtoReflect.Descriptor instead.
func (*CertProfileInfoRequest) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{0}
}

func (x *CertProfileInfoRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *CertProfileInfoRequest) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

// IssuersInfoResponse provides response for Issuers Info request
type IssuersInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issuers []*IssuerInfo `protobuf:"bytes,1,rep,name=issuers,proto3" json:"issuers,omitempty"`
}

func (x *IssuersInfoResponse) Reset() {
	*x = IssuersInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssuersInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssuersInfoResponse) ProtoMessage() {}

func (x *IssuersInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssuersInfoResponse.ProtoReflect.Descriptor instead.
func (*IssuersInfoResponse) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{1}
}

func (x *IssuersInfoResponse) GetIssuers() []*IssuerInfo {
	if x != nil {
		return x.Issuers
	}
	return nil
}

// SignCertificateRequest specifies certificate sign request
type SignCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RequestFormat provides the certificate request format: CSR, CMS
	RequestFormat EncodingFormat `protobuf:"varint,1,opt,name=request_format,json=requestFormat,proto3,enum=pb.EncodingFormat" json:"request_format,omitempty"`
	// Request provides the certificate request
	Request []byte `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	// Profile specifies the certificate profile: client, server, spiffe
	Profile string `protobuf:"bytes,3,opt,name=profile,proto3" json:"profile,omitempty"`
	// IssuerLabel specifies which Issuer to be appointed to sign the request
	IssuerLabel string `protobuf:"bytes,4,opt,name=issuer_label,json=issuerLabel,proto3" json:"issuer_label,omitempty"`
	// San specifies Subject Alternative Names
	San []string `protobuf:"bytes,5,rep,name=san,proto3" json:"san,omitempty"`
	// Subject specifies name
	Subject *X509Subject `protobuf:"bytes,6,opt,name=subject,proto3" json:"subject,omitempty"`
	// Token provides the authorization token for the request
	Token string `protobuf:"bytes,7,opt,name=token,proto3" json:"token,omitempty"`
	// OrgId provides the ID of Organization that certificate belongs to
	OrgId uint64 `protobuf:"varint,8,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// NotBefore is the time when the validity period starts
	NotBefore *timestamp.Timestamp `protobuf:"bytes,9,opt,name=not_before,proto3" json:"not_before,omitempty"`
	// NotAfter is the time when the validity period ends
	NotAfter *timestamp.Timestamp `protobuf:"bytes,10,opt,name=not_after,proto3" json:"not_after,omitempty"`
}

func (x *SignCertificateRequest) Reset() {
	*x = SignCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignCertificateRequest) ProtoMessage() {}

func (x *SignCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignCertificateRequest.ProtoReflect.Descriptor instead.
func (*SignCertificateRequest) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{2}
}

func (x *SignCertificateRequest) GetRequestFormat() EncodingFormat {
	if x != nil {
		return x.RequestFormat
	}
	return EncodingFormat_PEM
}

func (x *SignCertificateRequest) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *SignCertificateRequest) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *SignCertificateRequest) GetIssuerLabel() string {
	if x != nil {
		return x.IssuerLabel
	}
	return ""
}

func (x *SignCertificateRequest) GetSan() []string {
	if x != nil {
		return x.San
	}
	return nil
}

func (x *SignCertificateRequest) GetSubject() *X509Subject {
	if x != nil {
		return x.Subject
	}
	return nil
}

func (x *SignCertificateRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SignCertificateRequest) GetOrgId() uint64 {
	if x != nil {
		return x.OrgId
	}
	return 0
}

func (x *SignCertificateRequest) GetNotBefore() *timestamp.Timestamp {
	if x != nil {
		return x.NotBefore
	}
	return nil
}

func (x *SignCertificateRequest) GetNotAfter() *timestamp.Timestamp {
	if x != nil {
		return x.NotAfter
	}
	return nil
}

// PublishCrlsRequest allows to publish CRLs on demand
type PublishCrlsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// IKID specifies Issuer, or empty to publish for all issuers
	Ikid string `protobuf:"bytes,1,opt,name=ikid,proto3" json:"ikid,omitempty"`
}

func (x *PublishCrlsRequest) Reset() {
	*x = PublishCrlsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublishCrlsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublishCrlsRequest) ProtoMessage() {}

func (x *PublishCrlsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublishCrlsRequest.ProtoReflect.Descriptor instead.
func (*PublishCrlsRequest) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{3}
}

func (x *PublishCrlsRequest) GetIkid() string {
	if x != nil {
		return x.Ikid
	}
	return ""
}

// RevokeCertificateRequest specifies revocation request
type RevokeCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id specifies certificate ID.
	// If it's not set, then SKID must be provided
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// SKID specifies Subject Key ID to search
	Skid string `protobuf:"bytes,2,opt,name=skid,proto3" json:"skid,omitempty"`
	// Reason for revocation
	Reason Reason `protobuf:"varint,3,opt,name=reason,proto3,enum=pb.Reason" json:"reason,omitempty"`
}

func (x *RevokeCertificateRequest) Reset() {
	*x = RevokeCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ca_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeCertificateRequest) ProtoMessage() {}

func (x *RevokeCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ca_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeCertificateRequest.ProtoReflect.Descriptor instead.
func (*RevokeCertificateRequest) Descriptor() ([]byte, []int) {
	return file_ca_proto_rawDescGZIP(), []int{4}
}

func (x *RevokeCertificateRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RevokeCertificateRequest) GetSkid() string {
	if x != nil {
		return x.Skid
	}
	return ""
}

func (x *RevokeCertificateRequest) GetReason() Reason {
	if x != nil {
		return x.Reason
	}
	return Reason_UNSPECIFIED
}

var File_ca_proto protoreflect.FileDescriptor

var file_ca_proto_rawDesc = []byte{
	0x0a, 0x08, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0a,
	0x70, 0x6b, 0x69, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x16, 0x43, 0x65, 0x72, 0x74, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x22, 0x3f, 0x0a, 0x13, 0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72,
	0x73, 0x22, 0x8a, 0x03, 0x0a, 0x16, 0x53, 0x69, 0x67, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x39, 0x0a, 0x0e,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69,
	0x6e, 0x67, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x10,
	0x0a, 0x03, 0x73, 0x61, 0x6e, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x73, 0x61, 0x6e,
	0x12, 0x29, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x58, 0x35, 0x30, 0x39, 0x53, 0x75, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x72, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x6e, 0x6f, 0x74, 0x5f,
	0x62, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6e, 0x6f, 0x74, 0x5f, 0x62, 0x65,
	0x66, 0x6f, 0x72, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65,
	0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x22, 0x28,
	0x0a, 0x12, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x72, 0x6c, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6b, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x69, 0x6b, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x18, 0x52, 0x65, 0x76, 0x6f,
	0x6b, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x6b, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x32, 0x80, 0x03, 0x0a,
	0x09, 0x43, 0x41, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0b, 0x50, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x07, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x72, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0f, 0x53, 0x69, 0x67, 0x6e, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x39, 0x0a, 0x0b, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x72, 0x6c, 0x73, 0x12, 0x16,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x43, 0x72, 0x6c, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x11, 0x52, 0x65,
	0x76, 0x6f, 0x6b, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12,
	0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6b,
	0x73, 0x70, 0x61, 0x6e, 0x64, 0x2f, 0x74, 0x72, 0x75, 0x73, 0x74, 0x79, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ca_proto_rawDescOnce sync.Once
	file_ca_proto_rawDescData = file_ca_proto_rawDesc
)

func file_ca_proto_rawDescGZIP() []byte {
	file_ca_proto_rawDescOnce.Do(func() {
		file_ca_proto_rawDescData = protoimpl.X.CompressGZIP(file_ca_proto_rawDescData)
	})
	return file_ca_proto_rawDescData
}

var file_ca_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_ca_proto_goTypes = []interface{}{
	(*CertProfileInfoRequest)(nil),     // 0: pb.CertProfileInfoRequest
	(*IssuersInfoResponse)(nil),        // 1: pb.IssuersInfoResponse
	(*SignCertificateRequest)(nil),     // 2: pb.SignCertificateRequest
	(*PublishCrlsRequest)(nil),         // 3: pb.PublishCrlsRequest
	(*RevokeCertificateRequest)(nil),   // 4: pb.RevokeCertificateRequest
	(*IssuerInfo)(nil),                 // 5: pb.IssuerInfo
	(EncodingFormat)(0),                // 6: pb.EncodingFormat
	(*X509Subject)(nil),                // 7: pb.X509Subject
	(*timestamp.Timestamp)(nil),        // 8: google.protobuf.Timestamp
	(Reason)(0),                        // 9: pb.Reason
	(*empty.Empty)(nil),                // 10: google.protobuf.Empty
	(*CertProfileInfo)(nil),            // 11: pb.CertProfileInfo
	(*CertificateResponse)(nil),        // 12: pb.CertificateResponse
	(*CrlsResponse)(nil),               // 13: pb.CrlsResponse
	(*RevokedCertificateResponse)(nil), // 14: pb.RevokedCertificateResponse
}
var file_ca_proto_depIdxs = []int32{
	5,  // 0: pb.IssuersInfoResponse.issuers:type_name -> pb.IssuerInfo
	6,  // 1: pb.SignCertificateRequest.request_format:type_name -> pb.EncodingFormat
	7,  // 2: pb.SignCertificateRequest.subject:type_name -> pb.X509Subject
	8,  // 3: pb.SignCertificateRequest.not_before:type_name -> google.protobuf.Timestamp
	8,  // 4: pb.SignCertificateRequest.not_after:type_name -> google.protobuf.Timestamp
	9,  // 5: pb.RevokeCertificateRequest.reason:type_name -> pb.Reason
	0,  // 6: pb.CAService.ProfileInfo:input_type -> pb.CertProfileInfoRequest
	10, // 7: pb.CAService.Issuers:input_type -> google.protobuf.Empty
	2,  // 8: pb.CAService.SignCertificate:input_type -> pb.SignCertificateRequest
	3,  // 9: pb.CAService.PublishCrls:input_type -> pb.PublishCrlsRequest
	4,  // 10: pb.CAService.RevokeCertificate:input_type -> pb.RevokeCertificateRequest
	11, // 11: pb.CAService.ProfileInfo:output_type -> pb.CertProfileInfo
	1,  // 12: pb.CAService.Issuers:output_type -> pb.IssuersInfoResponse
	12, // 13: pb.CAService.SignCertificate:output_type -> pb.CertificateResponse
	13, // 14: pb.CAService.PublishCrls:output_type -> pb.CrlsResponse
	14, // 15: pb.CAService.RevokeCertificate:output_type -> pb.RevokedCertificateResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_ca_proto_init() }
func file_ca_proto_init() {
	if File_ca_proto != nil {
		return
	}
	file_pkix_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ca_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertProfileInfoRequest); i {
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
		file_ca_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssuersInfoResponse); i {
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
		file_ca_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignCertificateRequest); i {
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
		file_ca_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublishCrlsRequest); i {
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
		file_ca_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeCertificateRequest); i {
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
			RawDescriptor: file_ca_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ca_proto_goTypes,
		DependencyIndexes: file_ca_proto_depIdxs,
		MessageInfos:      file_ca_proto_msgTypes,
	}.Build()
	File_ca_proto = out.File
	file_ca_proto_rawDesc = nil
	file_ca_proto_goTypes = nil
	file_ca_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CAServiceClient is the client API for CAService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CAServiceClient interface {
	// ProfileInfo returns the certificate profile info
	ProfileInfo(ctx context.Context, in *CertProfileInfoRequest, opts ...grpc.CallOption) (*CertProfileInfo, error)
	// Issuers returns the issuing CAs
	Issuers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IssuersInfoResponse, error)
	// SignCertificate returns the certificate
	SignCertificate(ctx context.Context, in *SignCertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	// PublishCrls returns published CRLs
	PublishCrls(ctx context.Context, in *PublishCrlsRequest, opts ...grpc.CallOption) (*CrlsResponse, error)
	// RevokeCertificate returns the revoked certificate
	RevokeCertificate(ctx context.Context, in *RevokeCertificateRequest, opts ...grpc.CallOption) (*RevokedCertificateResponse, error)
}

type cAServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCAServiceClient(cc grpc.ClientConnInterface) CAServiceClient {
	return &cAServiceClient{cc}
}

func (c *cAServiceClient) ProfileInfo(ctx context.Context, in *CertProfileInfoRequest, opts ...grpc.CallOption) (*CertProfileInfo, error) {
	out := new(CertProfileInfo)
	err := c.cc.Invoke(ctx, "/pb.CAService/ProfileInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cAServiceClient) Issuers(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IssuersInfoResponse, error) {
	out := new(IssuersInfoResponse)
	err := c.cc.Invoke(ctx, "/pb.CAService/Issuers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cAServiceClient) SignCertificate(ctx context.Context, in *SignCertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/pb.CAService/SignCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cAServiceClient) PublishCrls(ctx context.Context, in *PublishCrlsRequest, opts ...grpc.CallOption) (*CrlsResponse, error) {
	out := new(CrlsResponse)
	err := c.cc.Invoke(ctx, "/pb.CAService/PublishCrls", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cAServiceClient) RevokeCertificate(ctx context.Context, in *RevokeCertificateRequest, opts ...grpc.CallOption) (*RevokedCertificateResponse, error) {
	out := new(RevokedCertificateResponse)
	err := c.cc.Invoke(ctx, "/pb.CAService/RevokeCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CAServiceServer is the server API for CAService service.
type CAServiceServer interface {
	// ProfileInfo returns the certificate profile info
	ProfileInfo(context.Context, *CertProfileInfoRequest) (*CertProfileInfo, error)
	// Issuers returns the issuing CAs
	Issuers(context.Context, *empty.Empty) (*IssuersInfoResponse, error)
	// SignCertificate returns the certificate
	SignCertificate(context.Context, *SignCertificateRequest) (*CertificateResponse, error)
	// PublishCrls returns published CRLs
	PublishCrls(context.Context, *PublishCrlsRequest) (*CrlsResponse, error)
	// RevokeCertificate returns the revoked certificate
	RevokeCertificate(context.Context, *RevokeCertificateRequest) (*RevokedCertificateResponse, error)
}

// UnimplementedCAServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCAServiceServer struct {
}

func (*UnimplementedCAServiceServer) ProfileInfo(context.Context, *CertProfileInfoRequest) (*CertProfileInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProfileInfo not implemented")
}
func (*UnimplementedCAServiceServer) Issuers(context.Context, *empty.Empty) (*IssuersInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Issuers not implemented")
}
func (*UnimplementedCAServiceServer) SignCertificate(context.Context, *SignCertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignCertificate not implemented")
}
func (*UnimplementedCAServiceServer) PublishCrls(context.Context, *PublishCrlsRequest) (*CrlsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishCrls not implemented")
}
func (*UnimplementedCAServiceServer) RevokeCertificate(context.Context, *RevokeCertificateRequest) (*RevokedCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeCertificate not implemented")
}

func RegisterCAServiceServer(s *grpc.Server, srv CAServiceServer) {
	s.RegisterService(&_CAService_serviceDesc, srv)
}

func _CAService_ProfileInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertProfileInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAServiceServer).ProfileInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CAService/ProfileInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAServiceServer).ProfileInfo(ctx, req.(*CertProfileInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CAService_Issuers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAServiceServer).Issuers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CAService/Issuers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAServiceServer).Issuers(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CAService_SignCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAServiceServer).SignCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CAService/SignCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAServiceServer).SignCertificate(ctx, req.(*SignCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CAService_PublishCrls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishCrlsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAServiceServer).PublishCrls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CAService/PublishCrls",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAServiceServer).PublishCrls(ctx, req.(*PublishCrlsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CAService_RevokeCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CAServiceServer).RevokeCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CAService/RevokeCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CAServiceServer).RevokeCertificate(ctx, req.(*RevokeCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CAService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CAService",
	HandlerType: (*CAServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProfileInfo",
			Handler:    _CAService_ProfileInfo_Handler,
		},
		{
			MethodName: "Issuers",
			Handler:    _CAService_Issuers_Handler,
		},
		{
			MethodName: "SignCertificate",
			Handler:    _CAService_SignCertificate_Handler,
		},
		{
			MethodName: "PublishCrls",
			Handler:    _CAService_PublishCrls_Handler,
		},
		{
			MethodName: "RevokeCertificate",
			Handler:    _CAService_RevokeCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ca.proto",
}
