// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: gRPC/proto/auth.proto

package auth

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

type ResponseType int32

const (
	ResponseType_SUCCESS      ResponseType = 0
	ResponseType_FAILED       ResponseType = 1
	ResponseType_EXPIRED_DATE ResponseType = 2
)

// Enum value maps for ResponseType.
var (
	ResponseType_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILED",
		2: "EXPIRED_DATE",
	}
	ResponseType_value = map[string]int32{
		"SUCCESS":      0,
		"FAILED":       1,
		"EXPIRED_DATE": 2,
	}
)

func (x ResponseType) Enum() *ResponseType {
	p := new(ResponseType)
	*p = x
	return p
}

func (x ResponseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseType) Descriptor() protoreflect.EnumDescriptor {
	return file_gRPC_proto_auth_proto_enumTypes[0].Descriptor()
}

func (ResponseType) Type() protoreflect.EnumType {
	return &file_gRPC_proto_auth_proto_enumTypes[0]
}

func (x ResponseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseType.Descriptor instead.
func (ResponseType) EnumDescriptor() ([]byte, []int) {
	return file_gRPC_proto_auth_proto_rawDescGZIP(), []int{0}
}

type AuthData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address     string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PasetoToken string `protobuf:"bytes,2,opt,name=pasetoToken,proto3" json:"pasetoToken,omitempty"`
	CreatedDate int64  `protobuf:"varint,3,opt,name=createdDate,proto3" json:"createdDate,omitempty"`
	ExpireDate  int64  `protobuf:"varint,4,opt,name=expireDate,proto3" json:"expireDate,omitempty"`
}

func (x *AuthData) Reset() {
	*x = AuthData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthData) ProtoMessage() {}

func (x *AuthData) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthData.ProtoReflect.Descriptor instead.
func (*AuthData) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthData) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AuthData) GetPasetoToken() string {
	if x != nil {
		return x.PasetoToken
	}
	return ""
}

func (x *AuthData) GetCreatedDate() int64 {
	if x != nil {
		return x.CreatedDate
	}
	return 0
}

func (x *AuthData) GetExpireDate() int64 {
	if x != nil {
		return x.ExpireDate
	}
	return 0
}

type VerifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status []ResponseType `protobuf:"varint,1,rep,packed,name=status,proto3,enum=ResponseType" json:"status,omitempty"`
	Auth   *AuthData      `protobuf:"bytes,2,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *VerifyResponse) Reset() {
	*x = VerifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyResponse) ProtoMessage() {}

func (x *VerifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyResponse.ProtoReflect.Descriptor instead.
func (*VerifyResponse) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *VerifyResponse) GetStatus() []ResponseType {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *VerifyResponse) GetAuth() *AuthData {
	if x != nil {
		return x.Auth
	}
	return nil
}

type CreateNewPasetoTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth *AuthData `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *CreateNewPasetoTokenRequest) Reset() {
	*x = CreateNewPasetoTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewPasetoTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewPasetoTokenRequest) ProtoMessage() {}

func (x *CreateNewPasetoTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewPasetoTokenRequest.ProtoReflect.Descriptor instead.
func (*CreateNewPasetoTokenRequest) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *CreateNewPasetoTokenRequest) GetAuth() *AuthData {
	if x != nil {
		return x.Auth
	}
	return nil
}

type CreateNewPasetoTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Auth *AuthData `protobuf:"bytes,1,opt,name=auth,proto3" json:"auth,omitempty"`
}

func (x *CreateNewPasetoTokenResponse) Reset() {
	*x = CreateNewPasetoTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNewPasetoTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNewPasetoTokenResponse) ProtoMessage() {}

func (x *CreateNewPasetoTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNewPasetoTokenResponse.ProtoReflect.Descriptor instead.
func (*CreateNewPasetoTokenResponse) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *CreateNewPasetoTokenResponse) GetAuth() *AuthData {
	if x != nil {
		return x.Auth
	}
	return nil
}

type VerifyPasetoTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PasetoToken string `protobuf:"bytes,1,opt,name=pasetoToken,proto3" json:"pasetoToken,omitempty"`
}

func (x *VerifyPasetoTokenRequest) Reset() {
	*x = VerifyPasetoTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyPasetoTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPasetoTokenRequest) ProtoMessage() {}

func (x *VerifyPasetoTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPasetoTokenRequest.ProtoReflect.Descriptor instead.
func (*VerifyPasetoTokenRequest) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_auth_proto_rawDescGZIP(), []int{4}
}

func (x *VerifyPasetoTokenRequest) GetPasetoToken() string {
	if x != nil {
		return x.PasetoToken
	}
	return ""
}

type VerifyPasetoTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *VerifyResponse `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *VerifyPasetoTokenResponse) Reset() {
	*x = VerifyPasetoTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyPasetoTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyPasetoTokenResponse) ProtoMessage() {}

func (x *VerifyPasetoTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyPasetoTokenResponse.ProtoReflect.Descriptor instead.
func (*VerifyPasetoTokenResponse) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_auth_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyPasetoTokenResponse) GetStatus() *VerifyResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_gRPC_proto_auth_proto protoreflect.FileDescriptor

var file_gRPC_proto_auth_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x08, 0x41, 0x75, 0x74, 0x68,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x70, 0x61, 0x73, 0x65, 0x74, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x73, 0x65, 0x74, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x44, 0x61,
	0x74, 0x65, 0x22, 0x56, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x04, 0x61,
	0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x3c, 0x0a, 0x1b, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x65, 0x74, 0x6f, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x04, 0x61, 0x75, 0x74,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x3d, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x73, 0x65, 0x74, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x3c, 0x0a, 0x18, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x50, 0x61, 0x73, 0x65, 0x74, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x73, 0x65, 0x74, 0x6f, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x73, 0x65, 0x74, 0x6f,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x44, 0x0a, 0x19, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50,
	0x61, 0x73, 0x65, 0x74, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x39, 0x0a, 0x0c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x5f,
	0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x32, 0x9d, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x75, 0x74, 0x68, 0x12, 0x1c, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77,
	0x50, 0x61, 0x73, 0x65, 0x74, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x50, 0x61,
	0x73, 0x65, 0x74, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x43, 0x0a, 0x0a, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x19, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x65, 0x74, 0x6f, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x50, 0x61, 0x73, 0x65, 0x74, 0x6f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gRPC_proto_auth_proto_rawDescOnce sync.Once
	file_gRPC_proto_auth_proto_rawDescData = file_gRPC_proto_auth_proto_rawDesc
)

func file_gRPC_proto_auth_proto_rawDescGZIP() []byte {
	file_gRPC_proto_auth_proto_rawDescOnce.Do(func() {
		file_gRPC_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_gRPC_proto_auth_proto_rawDescData)
	})
	return file_gRPC_proto_auth_proto_rawDescData
}

var file_gRPC_proto_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gRPC_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_gRPC_proto_auth_proto_goTypes = []interface{}{
	(ResponseType)(0),                    // 0: ResponseType
	(*AuthData)(nil),                     // 1: AuthData
	(*VerifyResponse)(nil),               // 2: VerifyResponse
	(*CreateNewPasetoTokenRequest)(nil),  // 3: CreateNewPasetoTokenRequest
	(*CreateNewPasetoTokenResponse)(nil), // 4: CreateNewPasetoTokenResponse
	(*VerifyPasetoTokenRequest)(nil),     // 5: VerifyPasetoTokenRequest
	(*VerifyPasetoTokenResponse)(nil),    // 6: VerifyPasetoTokenResponse
}
var file_gRPC_proto_auth_proto_depIdxs = []int32{
	0, // 0: VerifyResponse.status:type_name -> ResponseType
	1, // 1: VerifyResponse.auth:type_name -> AuthData
	1, // 2: CreateNewPasetoTokenRequest.auth:type_name -> AuthData
	1, // 3: CreateNewPasetoTokenResponse.auth:type_name -> AuthData
	2, // 4: VerifyPasetoTokenResponse.status:type_name -> VerifyResponse
	3, // 5: AuthService.CreateAuth:input_type -> CreateNewPasetoTokenRequest
	5, // 6: AuthService.VerifyAuth:input_type -> VerifyPasetoTokenRequest
	4, // 7: AuthService.CreateAuth:output_type -> CreateNewPasetoTokenResponse
	6, // 8: AuthService.VerifyAuth:output_type -> VerifyPasetoTokenResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_gRPC_proto_auth_proto_init() }
func file_gRPC_proto_auth_proto_init() {
	if File_gRPC_proto_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gRPC_proto_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthData); i {
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
		file_gRPC_proto_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyResponse); i {
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
		file_gRPC_proto_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewPasetoTokenRequest); i {
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
		file_gRPC_proto_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNewPasetoTokenResponse); i {
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
		file_gRPC_proto_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyPasetoTokenRequest); i {
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
		file_gRPC_proto_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyPasetoTokenResponse); i {
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
			RawDescriptor: file_gRPC_proto_auth_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gRPC_proto_auth_proto_goTypes,
		DependencyIndexes: file_gRPC_proto_auth_proto_depIdxs,
		EnumInfos:         file_gRPC_proto_auth_proto_enumTypes,
		MessageInfos:      file_gRPC_proto_auth_proto_msgTypes,
	}.Build()
	File_gRPC_proto_auth_proto = out.File
	file_gRPC_proto_auth_proto_rawDesc = nil
	file_gRPC_proto_auth_proto_goTypes = nil
	file_gRPC_proto_auth_proto_depIdxs = nil
}
