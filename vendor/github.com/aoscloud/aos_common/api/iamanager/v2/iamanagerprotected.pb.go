// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: iamanager/v2/iamanagerprotected.proto

package iamanager

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type ClearRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *ClearRequest) Reset() {
	*x = ClearRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClearRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClearRequest) ProtoMessage() {}

func (x *ClearRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClearRequest.ProtoReflect.Descriptor instead.
func (*ClearRequest) Descriptor() ([]byte, []int) {
	return file_iamanager_v2_iamanagerprotected_proto_rawDescGZIP(), []int{0}
}

func (x *ClearRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type SetOwnerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SetOwnerRequest) Reset() {
	*x = SetOwnerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetOwnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetOwnerRequest) ProtoMessage() {}

func (x *SetOwnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetOwnerRequest.ProtoReflect.Descriptor instead.
func (*SetOwnerRequest) Descriptor() ([]byte, []int) {
	return file_iamanager_v2_iamanagerprotected_proto_rawDescGZIP(), []int{1}
}

func (x *SetOwnerRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SetOwnerRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateKeyRequest) Reset() {
	*x = CreateKeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKeyRequest) ProtoMessage() {}

func (x *CreateKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKeyRequest.ProtoReflect.Descriptor instead.
func (*CreateKeyRequest) Descriptor() ([]byte, []int) {
	return file_iamanager_v2_iamanagerprotected_proto_rawDescGZIP(), []int{2}
}

func (x *CreateKeyRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateKeyRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Csr  string `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (x *CreateKeyResponse) Reset() {
	*x = CreateKeyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateKeyResponse) ProtoMessage() {}

func (x *CreateKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateKeyResponse.ProtoReflect.Descriptor instead.
func (*CreateKeyResponse) Descriptor() ([]byte, []int) {
	return file_iamanager_v2_iamanagerprotected_proto_rawDescGZIP(), []int{3}
}

func (x *CreateKeyResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateKeyResponse) GetCsr() string {
	if x != nil {
		return x.Csr
	}
	return ""
}

type ApplyCertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Cert string `protobuf:"bytes,2,opt,name=cert,proto3" json:"cert,omitempty"`
}

func (x *ApplyCertRequest) Reset() {
	*x = ApplyCertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyCertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyCertRequest) ProtoMessage() {}

func (x *ApplyCertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyCertRequest.ProtoReflect.Descriptor instead.
func (*ApplyCertRequest) Descriptor() ([]byte, []int) {
	return file_iamanager_v2_iamanagerprotected_proto_rawDescGZIP(), []int{4}
}

func (x *ApplyCertRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ApplyCertRequest) GetCert() string {
	if x != nil {
		return x.Cert
	}
	return ""
}

type ApplyCertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	CertUrl string `protobuf:"bytes,2,opt,name=cert_url,json=certUrl,proto3" json:"cert_url,omitempty"`
}

func (x *ApplyCertResponse) Reset() {
	*x = ApplyCertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyCertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyCertResponse) ProtoMessage() {}

func (x *ApplyCertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyCertResponse.ProtoReflect.Descriptor instead.
func (*ApplyCertResponse) Descriptor() ([]byte, []int) {
	return file_iamanager_v2_iamanagerprotected_proto_rawDescGZIP(), []int{5}
}

func (x *ApplyCertResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ApplyCertResponse) GetCertUrl() string {
	if x != nil {
		return x.CertUrl
	}
	return ""
}

type RegisterInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instance    *InstanceIdent          `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
	Permissions map[string]*Permissions `protobuf:"bytes,2,rep,name=permissions,proto3" json:"permissions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RegisterInstanceRequest) Reset() {
	*x = RegisterInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterInstanceRequest) ProtoMessage() {}

func (x *RegisterInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterInstanceRequest.ProtoReflect.Descriptor instead.
func (*RegisterInstanceRequest) Descriptor() ([]byte, []int) {
	return file_iamanager_v2_iamanagerprotected_proto_rawDescGZIP(), []int{6}
}

func (x *RegisterInstanceRequest) GetInstance() *InstanceIdent {
	if x != nil {
		return x.Instance
	}
	return nil
}

func (x *RegisterInstanceRequest) GetPermissions() map[string]*Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type RegisterInstanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret string `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
}

func (x *RegisterInstanceResponse) Reset() {
	*x = RegisterInstanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterInstanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterInstanceResponse) ProtoMessage() {}

func (x *RegisterInstanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterInstanceResponse.ProtoReflect.Descriptor instead.
func (*RegisterInstanceResponse) Descriptor() ([]byte, []int) {
	return file_iamanager_v2_iamanagerprotected_proto_rawDescGZIP(), []int{7}
}

func (x *RegisterInstanceResponse) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type UnregisterInstanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instance *InstanceIdent `protobuf:"bytes,1,opt,name=instance,proto3" json:"instance,omitempty"`
}

func (x *UnregisterInstanceRequest) Reset() {
	*x = UnregisterInstanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnregisterInstanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnregisterInstanceRequest) ProtoMessage() {}

func (x *UnregisterInstanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnregisterInstanceRequest.ProtoReflect.Descriptor instead.
func (*UnregisterInstanceRequest) Descriptor() ([]byte, []int) {
	return file_iamanager_v2_iamanagerprotected_proto_rawDescGZIP(), []int{8}
}

func (x *UnregisterInstanceRequest) GetInstance() *InstanceIdent {
	if x != nil {
		return x.Instance
	}
	return nil
}

type EncryptDiskRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *EncryptDiskRequest) Reset() {
	*x = EncryptDiskRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptDiskRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptDiskRequest) ProtoMessage() {}

func (x *EncryptDiskRequest) ProtoReflect() protoreflect.Message {
	mi := &file_iamanager_v2_iamanagerprotected_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptDiskRequest.ProtoReflect.Descriptor instead.
func (*EncryptDiskRequest) Descriptor() ([]byte, []int) {
	return file_iamanager_v2_iamanagerprotected_proto_rawDescGZIP(), []int{9}
}

func (x *EncryptDiskRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_iamanager_v2_iamanagerprotected_proto protoreflect.FileDescriptor

var file_iamanager_v2_iamanagerprotected_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x32, 0x2f, 0x69,
	0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x69, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x32, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x22, 0x69, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x76, 0x32,
	0x2f, 0x69, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x0c, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x41, 0x0a, 0x0f, 0x53, 0x65,
	0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x42, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x39, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x73,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x73, 0x72, 0x22, 0x3a, 0x0a, 0x10,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x65, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x65, 0x72, 0x74, 0x22, 0x42, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x65, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x65, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x87, 0x02, 0x0a,
	0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x61, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x12, 0x58, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x69, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x59, 0x0a, 0x10, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x69, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x32, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x54, 0x0a, 0x19, 0x55, 0x6e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x61, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x22, 0x30, 0x0a, 0x12, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x32, 0x8a, 0x05, 0x0a, 0x13, 0x49, 0x41, 0x4d, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x53, 0x65,
	0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x69, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x05, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x12, 0x1a, 0x2e, 0x69, 0x61, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4e,
	0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x1e, 0x2e, 0x69, 0x61,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x61,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e,
	0x0a, 0x09, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x43, 0x65, 0x72, 0x74, 0x12, 0x1e, 0x2e, 0x69, 0x61,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x69, 0x61,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x43, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49,
	0x0a, 0x0b, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x12, 0x20, 0x2e,
	0x69, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x12, 0x46, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x63, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x25, 0x2e, 0x69, 0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x69,
	0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x12, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x27, 0x2e, 0x69,
	0x61, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x2e, 0x55, 0x6e, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_iamanager_v2_iamanagerprotected_proto_rawDescOnce sync.Once
	file_iamanager_v2_iamanagerprotected_proto_rawDescData = file_iamanager_v2_iamanagerprotected_proto_rawDesc
)

func file_iamanager_v2_iamanagerprotected_proto_rawDescGZIP() []byte {
	file_iamanager_v2_iamanagerprotected_proto_rawDescOnce.Do(func() {
		file_iamanager_v2_iamanagerprotected_proto_rawDescData = protoimpl.X.CompressGZIP(file_iamanager_v2_iamanagerprotected_proto_rawDescData)
	})
	return file_iamanager_v2_iamanagerprotected_proto_rawDescData
}

var file_iamanager_v2_iamanagerprotected_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_iamanager_v2_iamanagerprotected_proto_goTypes = []interface{}{
	(*ClearRequest)(nil),              // 0: iamanager.v2.ClearRequest
	(*SetOwnerRequest)(nil),           // 1: iamanager.v2.SetOwnerRequest
	(*CreateKeyRequest)(nil),          // 2: iamanager.v2.CreateKeyRequest
	(*CreateKeyResponse)(nil),         // 3: iamanager.v2.CreateKeyResponse
	(*ApplyCertRequest)(nil),          // 4: iamanager.v2.ApplyCertRequest
	(*ApplyCertResponse)(nil),         // 5: iamanager.v2.ApplyCertResponse
	(*RegisterInstanceRequest)(nil),   // 6: iamanager.v2.RegisterInstanceRequest
	(*RegisterInstanceResponse)(nil),  // 7: iamanager.v2.RegisterInstanceResponse
	(*UnregisterInstanceRequest)(nil), // 8: iamanager.v2.UnregisterInstanceRequest
	(*EncryptDiskRequest)(nil),        // 9: iamanager.v2.EncryptDiskRequest
	nil,                               // 10: iamanager.v2.RegisterInstanceRequest.PermissionsEntry
	(*InstanceIdent)(nil),             // 11: iamanager.v2.InstanceIdent
	(*Permissions)(nil),               // 12: iamanager.v2.Permissions
	(*empty.Empty)(nil),               // 13: google.protobuf.Empty
}
var file_iamanager_v2_iamanagerprotected_proto_depIdxs = []int32{
	11, // 0: iamanager.v2.RegisterInstanceRequest.instance:type_name -> iamanager.v2.InstanceIdent
	10, // 1: iamanager.v2.RegisterInstanceRequest.permissions:type_name -> iamanager.v2.RegisterInstanceRequest.PermissionsEntry
	11, // 2: iamanager.v2.UnregisterInstanceRequest.instance:type_name -> iamanager.v2.InstanceIdent
	12, // 3: iamanager.v2.RegisterInstanceRequest.PermissionsEntry.value:type_name -> iamanager.v2.Permissions
	1,  // 4: iamanager.v2.IAMProtectedService.SetOwner:input_type -> iamanager.v2.SetOwnerRequest
	0,  // 5: iamanager.v2.IAMProtectedService.Clear:input_type -> iamanager.v2.ClearRequest
	2,  // 6: iamanager.v2.IAMProtectedService.CreateKey:input_type -> iamanager.v2.CreateKeyRequest
	4,  // 7: iamanager.v2.IAMProtectedService.ApplyCert:input_type -> iamanager.v2.ApplyCertRequest
	9,  // 8: iamanager.v2.IAMProtectedService.EncryptDisk:input_type -> iamanager.v2.EncryptDiskRequest
	13, // 9: iamanager.v2.IAMProtectedService.FinishProvisioning:input_type -> google.protobuf.Empty
	6,  // 10: iamanager.v2.IAMProtectedService.RegisterInstance:input_type -> iamanager.v2.RegisterInstanceRequest
	8,  // 11: iamanager.v2.IAMProtectedService.UnregisterInstance:input_type -> iamanager.v2.UnregisterInstanceRequest
	13, // 12: iamanager.v2.IAMProtectedService.SetOwner:output_type -> google.protobuf.Empty
	13, // 13: iamanager.v2.IAMProtectedService.Clear:output_type -> google.protobuf.Empty
	3,  // 14: iamanager.v2.IAMProtectedService.CreateKey:output_type -> iamanager.v2.CreateKeyResponse
	5,  // 15: iamanager.v2.IAMProtectedService.ApplyCert:output_type -> iamanager.v2.ApplyCertResponse
	13, // 16: iamanager.v2.IAMProtectedService.EncryptDisk:output_type -> google.protobuf.Empty
	13, // 17: iamanager.v2.IAMProtectedService.FinishProvisioning:output_type -> google.protobuf.Empty
	7,  // 18: iamanager.v2.IAMProtectedService.RegisterInstance:output_type -> iamanager.v2.RegisterInstanceResponse
	13, // 19: iamanager.v2.IAMProtectedService.UnregisterInstance:output_type -> google.protobuf.Empty
	12, // [12:20] is the sub-list for method output_type
	4,  // [4:12] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_iamanager_v2_iamanagerprotected_proto_init() }
func file_iamanager_v2_iamanagerprotected_proto_init() {
	if File_iamanager_v2_iamanagerprotected_proto != nil {
		return
	}
	file_iamanager_v2_iamanagercommon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_iamanager_v2_iamanagerprotected_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClearRequest); i {
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
		file_iamanager_v2_iamanagerprotected_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetOwnerRequest); i {
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
		file_iamanager_v2_iamanagerprotected_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKeyRequest); i {
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
		file_iamanager_v2_iamanagerprotected_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateKeyResponse); i {
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
		file_iamanager_v2_iamanagerprotected_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyCertRequest); i {
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
		file_iamanager_v2_iamanagerprotected_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyCertResponse); i {
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
		file_iamanager_v2_iamanagerprotected_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterInstanceRequest); i {
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
		file_iamanager_v2_iamanagerprotected_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterInstanceResponse); i {
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
		file_iamanager_v2_iamanagerprotected_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnregisterInstanceRequest); i {
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
		file_iamanager_v2_iamanagerprotected_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptDiskRequest); i {
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
			RawDescriptor: file_iamanager_v2_iamanagerprotected_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_iamanager_v2_iamanagerprotected_proto_goTypes,
		DependencyIndexes: file_iamanager_v2_iamanagerprotected_proto_depIdxs,
		MessageInfos:      file_iamanager_v2_iamanagerprotected_proto_msgTypes,
	}.Build()
	File_iamanager_v2_iamanagerprotected_proto = out.File
	file_iamanager_v2_iamanagerprotected_proto_rawDesc = nil
	file_iamanager_v2_iamanagerprotected_proto_goTypes = nil
	file_iamanager_v2_iamanagerprotected_proto_depIdxs = nil
}
