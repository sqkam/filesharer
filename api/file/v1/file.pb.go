// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.19.6
// source: api/file/v1/file.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListByAddrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *ListByAddrRequest) Reset() {
	*x = ListByAddrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_v1_file_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListByAddrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListByAddrRequest) ProtoMessage() {}

func (x *ListByAddrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_v1_file_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListByAddrRequest.ProtoReflect.Descriptor instead.
func (*ListByAddrRequest) Descriptor() ([]byte, []int) {
	return file_api_file_v1_file_proto_rawDescGZIP(), []int{0}
}

func (x *ListByAddrRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *ListByAddrRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type ListByAddrReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ListByAddrReplyItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListByAddrReply) Reset() {
	*x = ListByAddrReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_v1_file_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListByAddrReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListByAddrReply) ProtoMessage() {}

func (x *ListByAddrReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_v1_file_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListByAddrReply.ProtoReflect.Descriptor instead.
func (*ListByAddrReply) Descriptor() ([]byte, []int) {
	return file_api_file_v1_file_proto_rawDescGZIP(), []int{1}
}

func (x *ListByAddrReply) GetData() []*ListByAddrReplyItem {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetDetailByAddrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *GetDetailByAddrRequest) Reset() {
	*x = GetDetailByAddrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_v1_file_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDetailByAddrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetailByAddrRequest) ProtoMessage() {}

func (x *GetDetailByAddrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_v1_file_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetailByAddrRequest.ProtoReflect.Descriptor instead.
func (*GetDetailByAddrRequest) Descriptor() ([]byte, []int) {
	return file_api_file_v1_file_proto_rawDescGZIP(), []int{2}
}

func (x *GetDetailByAddrRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *GetDetailByAddrRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type GetDetailByAddrReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path      string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Size      int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	UpdatedAt int64  `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GetDetailByAddrReply) Reset() {
	*x = GetDetailByAddrReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_v1_file_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDetailByAddrReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDetailByAddrReply) ProtoMessage() {}

func (x *GetDetailByAddrReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_v1_file_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDetailByAddrReply.ProtoReflect.Descriptor instead.
func (*GetDetailByAddrReply) Descriptor() ([]byte, []int) {
	return file_api_file_v1_file_proto_rawDescGZIP(), []int{3}
}

func (x *GetDetailByAddrReply) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *GetDetailByAddrReply) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetDetailByAddrReply) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type DownloadByAddrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DownloadByAddrRequest) Reset() {
	*x = DownloadByAddrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_v1_file_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadByAddrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadByAddrRequest) ProtoMessage() {}

func (x *DownloadByAddrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_v1_file_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadByAddrRequest.ProtoReflect.Descriptor instead.
func (*DownloadByAddrRequest) Descriptor() ([]byte, []int) {
	return file_api_file_v1_file_proto_rawDescGZIP(), []int{4}
}

func (x *DownloadByAddrRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *DownloadByAddrRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DownloadByAddrReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DownloadByAddrReply) Reset() {
	*x = DownloadByAddrReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_v1_file_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadByAddrReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadByAddrReply) ProtoMessage() {}

func (x *DownloadByAddrReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_v1_file_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadByAddrReply.ProtoReflect.Descriptor instead.
func (*DownloadByAddrReply) Descriptor() ([]byte, []int) {
	return file_api_file_v1_file_proto_rawDescGZIP(), []int{5}
}

func (x *DownloadByAddrReply) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type DownloadDirByAddrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *DownloadDirByAddrRequest) Reset() {
	*x = DownloadDirByAddrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_v1_file_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadDirByAddrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadDirByAddrRequest) ProtoMessage() {}

func (x *DownloadDirByAddrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_v1_file_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadDirByAddrRequest.ProtoReflect.Descriptor instead.
func (*DownloadDirByAddrRequest) Descriptor() ([]byte, []int) {
	return file_api_file_v1_file_proto_rawDescGZIP(), []int{6}
}

func (x *DownloadDirByAddrRequest) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *DownloadDirByAddrRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type DownloadDirByAddrReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DownloadDirByAddrReply) Reset() {
	*x = DownloadDirByAddrReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_v1_file_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadDirByAddrReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadDirByAddrReply) ProtoMessage() {}

func (x *DownloadDirByAddrReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_v1_file_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadDirByAddrReply.ProtoReflect.Descriptor instead.
func (*DownloadDirByAddrReply) Descriptor() ([]byte, []int) {
	return file_api_file_v1_file_proto_rawDescGZIP(), []int{7}
}

func (x *DownloadDirByAddrReply) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListNodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListNodeRequest) Reset() {
	*x = ListNodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_v1_file_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodeRequest) ProtoMessage() {}

func (x *ListNodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_v1_file_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodeRequest.ProtoReflect.Descriptor instead.
func (*ListNodeRequest) Descriptor() ([]byte, []int) {
	return file_api_file_v1_file_proto_rawDescGZIP(), []int{8}
}

type ListNodeReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*ListNodeReplyItem `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListNodeReply) Reset() {
	*x = ListNodeReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_v1_file_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNodeReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodeReply) ProtoMessage() {}

func (x *ListNodeReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_v1_file_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodeReply.ProtoReflect.Descriptor instead.
func (*ListNodeReply) Descriptor() ([]byte, []int) {
	return file_api_file_v1_file_proto_rawDescGZIP(), []int{9}
}

func (x *ListNodeReply) GetData() []*ListNodeReplyItem {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListByAddrReplyItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path  string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	IsDir bool   `protobuf:"varint,2,opt,name=is_dir,json=isDir,proto3" json:"is_dir,omitempty"`
	Size  int64  `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListByAddrReplyItem) Reset() {
	*x = ListByAddrReplyItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_v1_file_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListByAddrReplyItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListByAddrReplyItem) ProtoMessage() {}

func (x *ListByAddrReplyItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_v1_file_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListByAddrReplyItem.ProtoReflect.Descriptor instead.
func (*ListByAddrReplyItem) Descriptor() ([]byte, []int) {
	return file_api_file_v1_file_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListByAddrReplyItem) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ListByAddrReplyItem) GetIsDir() bool {
	if x != nil {
		return x.IsDir
	}
	return false
}

func (x *ListByAddrReplyItem) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListNodeReplyItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceID      string `protobuf:"bytes,1,opt,name=ServiceID,proto3" json:"ServiceID,omitempty"`
	ServiceAddress string `protobuf:"bytes,2,opt,name=ServiceAddress,proto3" json:"ServiceAddress,omitempty"`
	ServicePort    int64  `protobuf:"varint,3,opt,name=ServicePort,proto3" json:"ServicePort,omitempty"`
}

func (x *ListNodeReplyItem) Reset() {
	*x = ListNodeReplyItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_file_v1_file_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNodeReplyItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNodeReplyItem) ProtoMessage() {}

func (x *ListNodeReplyItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_file_v1_file_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNodeReplyItem.ProtoReflect.Descriptor instead.
func (*ListNodeReplyItem) Descriptor() ([]byte, []int) {
	return file_api_file_v1_file_proto_rawDescGZIP(), []int{9, 0}
}

func (x *ListNodeReplyItem) GetServiceID() string {
	if x != nil {
		return x.ServiceID
	}
	return ""
}

func (x *ListNodeReplyItem) GetServiceAddress() string {
	if x != nil {
		return x.ServiceAddress
	}
	return ""
}

func (x *ListNodeReplyItem) GetServicePort() int64 {
	if x != nil {
		return x.ServicePort
	}
	return 0
}

var File_api_file_v1_file_proto protoreflect.FileDescriptor

var file_api_file_v1_file_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x35, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x69, 0x74, 0x65, 0x6d, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x1a, 0x45, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x69, 0x73, 0x44, 0x69, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x40, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x5d, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x48, 0x0a, 0x15,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x29, 0x0a, 0x13, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x4b, 0x0a, 0x18, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x72,
	0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64,
	0x72, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x2c,
	0x0a, 0x16, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x72, 0x42, 0x79, 0x41,
	0x64, 0x64, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x11, 0x0a, 0x0f,
	0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xb4, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x33, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x69, 0x74, 0x65, 0x6d,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x6e, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x26, 0x0a, 0x0e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x32, 0xb5, 0x07, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x4a, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x5b, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x12, 0x23,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x41, 0x64, 0x64,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5a, 0x0a, 0x0e, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x30, 0x01, 0x12, 0x63, 0x0a, 0x11, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x44, 0x69, 0x72, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x44, 0x69, 0x72, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x72, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x30, 0x01, 0x12, 0x67, 0x0a, 0x0e, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x48, 0x74, 0x74, 0x70, 0x12, 0x1e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x79,
	0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x7b, 0x0a, 0x12, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x79,
	0x41, 0x64, 0x64, 0x72, 0x48, 0x74, 0x74, 0x70, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x42,
	0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22, 0x14, 0x76, 0x31, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x66, 0x69, 0x6c, 0x65, 0x12,
	0x83, 0x01, 0x0a, 0x15, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x72, 0x42,
	0x79, 0x41, 0x64, 0x64, 0x72, 0x48, 0x74, 0x74, 0x70, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x44, 0x69, 0x72, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x44, 0x69, 0x72, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a,
	0x22, 0x13, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x64, 0x69, 0x72, 0x12, 0x5d, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64,
	0x65, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x78, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x48, 0x74, 0x74, 0x70, 0x12, 0x23, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x79, 0x41, 0x64, 0x64, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e,
	0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x42, 0x2a,
	0x0a, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x19, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x68, 0x61, 0x72, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_file_v1_file_proto_rawDescOnce sync.Once
	file_api_file_v1_file_proto_rawDescData = file_api_file_v1_file_proto_rawDesc
)

func file_api_file_v1_file_proto_rawDescGZIP() []byte {
	file_api_file_v1_file_proto_rawDescOnce.Do(func() {
		file_api_file_v1_file_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_file_v1_file_proto_rawDescData)
	})
	return file_api_file_v1_file_proto_rawDescData
}

var file_api_file_v1_file_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_api_file_v1_file_proto_goTypes = []any{
	(*ListByAddrRequest)(nil),        // 0: api.file.v1.ListByAddrRequest
	(*ListByAddrReply)(nil),          // 1: api.file.v1.ListByAddrReply
	(*GetDetailByAddrRequest)(nil),   // 2: api.file.v1.GetDetailByAddrRequest
	(*GetDetailByAddrReply)(nil),     // 3: api.file.v1.GetDetailByAddrReply
	(*DownloadByAddrRequest)(nil),    // 4: api.file.v1.DownloadByAddrRequest
	(*DownloadByAddrReply)(nil),      // 5: api.file.v1.DownloadByAddrReply
	(*DownloadDirByAddrRequest)(nil), // 6: api.file.v1.DownloadDirByAddrRequest
	(*DownloadDirByAddrReply)(nil),   // 7: api.file.v1.DownloadDirByAddrReply
	(*ListNodeRequest)(nil),          // 8: api.file.v1.ListNodeRequest
	(*ListNodeReply)(nil),            // 9: api.file.v1.ListNodeReply
	(*ListByAddrReplyItem)(nil),      // 10: api.file.v1.ListByAddrReply.item
	(*ListNodeReplyItem)(nil),        // 11: api.file.v1.ListNodeReply.item
}
var file_api_file_v1_file_proto_depIdxs = []int32{
	10, // 0: api.file.v1.ListByAddrReply.data:type_name -> api.file.v1.ListByAddrReply.item
	11, // 1: api.file.v1.ListNodeReply.data:type_name -> api.file.v1.ListNodeReply.item
	0,  // 2: api.file.v1.File.ListByAddr:input_type -> api.file.v1.ListByAddrRequest
	2,  // 3: api.file.v1.File.GetDetailByAddr:input_type -> api.file.v1.GetDetailByAddrRequest
	4,  // 4: api.file.v1.File.DownloadByAddr:input_type -> api.file.v1.DownloadByAddrRequest
	6,  // 5: api.file.v1.File.DownloadDirByAddr:input_type -> api.file.v1.DownloadDirByAddrRequest
	0,  // 6: api.file.v1.File.ListByAddrHttp:input_type -> api.file.v1.ListByAddrRequest
	4,  // 7: api.file.v1.File.DownloadByAddrHttp:input_type -> api.file.v1.DownloadByAddrRequest
	6,  // 8: api.file.v1.File.DownloadDirByAddrHttp:input_type -> api.file.v1.DownloadDirByAddrRequest
	8,  // 9: api.file.v1.File.ListNode:input_type -> api.file.v1.ListNodeRequest
	2,  // 10: api.file.v1.File.GetDetailByAddrHttp:input_type -> api.file.v1.GetDetailByAddrRequest
	1,  // 11: api.file.v1.File.ListByAddr:output_type -> api.file.v1.ListByAddrReply
	3,  // 12: api.file.v1.File.GetDetailByAddr:output_type -> api.file.v1.GetDetailByAddrReply
	5,  // 13: api.file.v1.File.DownloadByAddr:output_type -> api.file.v1.DownloadByAddrReply
	7,  // 14: api.file.v1.File.DownloadDirByAddr:output_type -> api.file.v1.DownloadDirByAddrReply
	1,  // 15: api.file.v1.File.ListByAddrHttp:output_type -> api.file.v1.ListByAddrReply
	5,  // 16: api.file.v1.File.DownloadByAddrHttp:output_type -> api.file.v1.DownloadByAddrReply
	7,  // 17: api.file.v1.File.DownloadDirByAddrHttp:output_type -> api.file.v1.DownloadDirByAddrReply
	9,  // 18: api.file.v1.File.ListNode:output_type -> api.file.v1.ListNodeReply
	3,  // 19: api.file.v1.File.GetDetailByAddrHttp:output_type -> api.file.v1.GetDetailByAddrReply
	11, // [11:20] is the sub-list for method output_type
	2,  // [2:11] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_file_v1_file_proto_init() }
func file_api_file_v1_file_proto_init() {
	if File_api_file_v1_file_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_file_v1_file_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListByAddrRequest); i {
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
		file_api_file_v1_file_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListByAddrReply); i {
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
		file_api_file_v1_file_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetDetailByAddrRequest); i {
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
		file_api_file_v1_file_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetDetailByAddrReply); i {
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
		file_api_file_v1_file_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DownloadByAddrRequest); i {
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
		file_api_file_v1_file_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DownloadByAddrReply); i {
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
		file_api_file_v1_file_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DownloadDirByAddrRequest); i {
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
		file_api_file_v1_file_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DownloadDirByAddrReply); i {
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
		file_api_file_v1_file_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ListNodeRequest); i {
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
		file_api_file_v1_file_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ListNodeReply); i {
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
		file_api_file_v1_file_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ListByAddrReplyItem); i {
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
		file_api_file_v1_file_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*ListNodeReplyItem); i {
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
			RawDescriptor: file_api_file_v1_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_file_v1_file_proto_goTypes,
		DependencyIndexes: file_api_file_v1_file_proto_depIdxs,
		MessageInfos:      file_api_file_v1_file_proto_msgTypes,
	}.Build()
	File_api_file_v1_file_proto = out.File
	file_api_file_v1_file_proto_rawDesc = nil
	file_api_file_v1_file_proto_goTypes = nil
	file_api_file_v1_file_proto_depIdxs = nil
}
