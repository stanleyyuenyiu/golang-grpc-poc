// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: protos/topk/topk.proto

package topk

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EErrorCode int32

const (
	EErrorCode_INVALID    EErrorCode = 0
	EErrorCode_UNEXPECTED EErrorCode = 1
)

// Enum value maps for EErrorCode.
var (
	EErrorCode_name = map[int32]string{
		0: "INVALID",
		1: "UNEXPECTED",
	}
	EErrorCode_value = map[string]int32{
		"INVALID":    0,
		"UNEXPECTED": 1,
	}
)

func (x EErrorCode) Enum() *EErrorCode {
	p := new(EErrorCode)
	*p = x
	return p
}

func (x EErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_topk_topk_proto_enumTypes[0].Descriptor()
}

func (EErrorCode) Type() protoreflect.EnumType {
	return &file_protos_topk_topk_proto_enumTypes[0]
}

func (x EErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EErrorCode.Descriptor instead.
func (EErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_protos_topk_topk_proto_rawDescGZIP(), []int{0}
}

type IncreaseCounterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *IncreaseCounterReq) Reset() {
	*x = IncreaseCounterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_topk_topk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncreaseCounterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncreaseCounterReq) ProtoMessage() {}

func (x *IncreaseCounterReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_topk_topk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncreaseCounterReq.ProtoReflect.Descriptor instead.
func (*IncreaseCounterReq) Descriptor() ([]byte, []int) {
	return file_protos_topk_topk_proto_rawDescGZIP(), []int{0}
}

func (x *IncreaseCounterReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type IncreaseCounterRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool   `protobuf:"varint,1,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
	Message   string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *IncreaseCounterRes) Reset() {
	*x = IncreaseCounterRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_topk_topk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncreaseCounterRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncreaseCounterRes) ProtoMessage() {}

func (x *IncreaseCounterRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_topk_topk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncreaseCounterRes.ProtoReflect.Descriptor instead.
func (*IncreaseCounterRes) Descriptor() ([]byte, []int) {
	return file_protos_topk_topk_proto_rawDescGZIP(), []int{1}
}

func (x *IncreaseCounterRes) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *IncreaseCounterRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListTopKReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	K uint32 `protobuf:"varint,1,opt,name=k,proto3" json:"k,omitempty"`
}

func (x *ListTopKReq) Reset() {
	*x = ListTopKReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_topk_topk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTopKReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopKReq) ProtoMessage() {}

func (x *ListTopKReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_topk_topk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopKReq.ProtoReflect.Descriptor instead.
func (*ListTopKReq) Descriptor() ([]byte, []int) {
	return file_protos_topk_topk_proto_rawDescGZIP(), []int{2}
}

func (x *ListTopKReq) GetK() uint32 {
	if x != nil {
		return x.K
	}
	return 0
}

type ListTopKRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool              `protobuf:"varint,1,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
	Message   map[string]uint32 `protobuf:"bytes,2,rep,name=message,proto3" json:"message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *ListTopKRes) Reset() {
	*x = ListTopKRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_topk_topk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTopKRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTopKRes) ProtoMessage() {}

func (x *ListTopKRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_topk_topk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTopKRes.ProtoReflect.Descriptor instead.
func (*ListTopKRes) Descriptor() ([]byte, []int) {
	return file_protos_topk_topk_proto_rawDescGZIP(), []int{3}
}

func (x *ListTopKRes) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *ListTopKRes) GetMessage() map[string]uint32 {
	if x != nil {
		return x.Message
	}
	return nil
}

type ErrorStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string       `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Details []*anypb.Any `protobuf:"bytes,2,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *ErrorStatus) Reset() {
	*x = ErrorStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_topk_topk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorStatus) ProtoMessage() {}

func (x *ErrorStatus) ProtoReflect() protoreflect.Message {
	mi := &file_protos_topk_topk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorStatus.ProtoReflect.Descriptor instead.
func (*ErrorStatus) Descriptor() ([]byte, []int) {
	return file_protos_topk_topk_proto_rawDescGZIP(), []int{4}
}

func (x *ErrorStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorStatus) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_protos_topk_topk_proto protoreflect.FileDescriptor

var file_protos_topk_topk_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x6f, 0x70, 0x6b, 0x2f, 0x74, 0x6f,
	0x70, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f, 0x70, 0x6b, 0x1a, 0x19,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x12, 0x49, 0x6e, 0x63,
	0x72, 0x65, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x4c, 0x0a, 0x12, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x1b, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x4b, 0x52, 0x65, 0x71, 0x12, 0x0c,
	0x0a, 0x01, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x6b, 0x22, 0xa1, 0x01, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x4b, 0x52, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x38, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x6f,
	0x70, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x4b, 0x52, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x3a, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x57, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x2a, 0x29, 0x0a, 0x0a, 0x45, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x32, 0x8b, 0x02, 0x0a, 0x04, 0x54, 0x6f, 0x70, 0x6b, 0x12, 0x47, 0x0a,
	0x0f, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x12, 0x18, 0x2e, 0x74, 0x6f, 0x70, 0x6b, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x74, 0x6f, 0x70,
	0x6b, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4a, 0x0a, 0x15, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x18, 0x2e, 0x74, 0x6f, 0x70, 0x6b, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x74, 0x6f, 0x70, 0x6b,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x4b, 0x52, 0x65, 0x73, 0x22, 0x00, 0x28, 0x01,
	0x30, 0x01, 0x12, 0x32, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x4b, 0x12, 0x11,
	0x2e, 0x74, 0x6f, 0x70, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x4b, 0x52, 0x65,
	0x71, 0x1a, 0x11, 0x2e, 0x74, 0x6f, 0x70, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70,
	0x4b, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x4b, 0x12, 0x11, 0x2e, 0x74, 0x6f, 0x70, 0x6b, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x4b, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x74, 0x6f,
	0x70, 0x6b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x6f, 0x70, 0x4b, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x30, 0x01, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x6f, 0x70,
	0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_topk_topk_proto_rawDescOnce sync.Once
	file_protos_topk_topk_proto_rawDescData = file_protos_topk_topk_proto_rawDesc
)

func file_protos_topk_topk_proto_rawDescGZIP() []byte {
	file_protos_topk_topk_proto_rawDescOnce.Do(func() {
		file_protos_topk_topk_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_topk_topk_proto_rawDescData)
	})
	return file_protos_topk_topk_proto_rawDescData
}

var file_protos_topk_topk_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_topk_topk_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_topk_topk_proto_goTypes = []interface{}{
	(EErrorCode)(0),            // 0: topk.EErrorCode
	(*IncreaseCounterReq)(nil), // 1: topk.IncreaseCounterReq
	(*IncreaseCounterRes)(nil), // 2: topk.IncreaseCounterRes
	(*ListTopKReq)(nil),        // 3: topk.ListTopKReq
	(*ListTopKRes)(nil),        // 4: topk.ListTopKRes
	(*ErrorStatus)(nil),        // 5: topk.ErrorStatus
	nil,                        // 6: topk.ListTopKRes.MessageEntry
	(*anypb.Any)(nil),          // 7: google.protobuf.Any
}
var file_protos_topk_topk_proto_depIdxs = []int32{
	6, // 0: topk.ListTopKRes.message:type_name -> topk.ListTopKRes.MessageEntry
	7, // 1: topk.ErrorStatus.details:type_name -> google.protobuf.Any
	1, // 2: topk.Topk.IncreaseCounter:input_type -> topk.IncreaseCounterReq
	1, // 3: topk.Topk.StreamIncreaseCounter:input_type -> topk.IncreaseCounterReq
	3, // 4: topk.Topk.ListTopK:input_type -> topk.ListTopKReq
	3, // 5: topk.Topk.StreamListTopK:input_type -> topk.ListTopKReq
	2, // 6: topk.Topk.IncreaseCounter:output_type -> topk.IncreaseCounterRes
	4, // 7: topk.Topk.StreamIncreaseCounter:output_type -> topk.ListTopKRes
	4, // 8: topk.Topk.ListTopK:output_type -> topk.ListTopKRes
	4, // 9: topk.Topk.StreamListTopK:output_type -> topk.ListTopKRes
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_topk_topk_proto_init() }
func file_protos_topk_topk_proto_init() {
	if File_protos_topk_topk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_topk_topk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncreaseCounterReq); i {
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
		file_protos_topk_topk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncreaseCounterRes); i {
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
		file_protos_topk_topk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTopKReq); i {
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
		file_protos_topk_topk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTopKRes); i {
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
		file_protos_topk_topk_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorStatus); i {
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
			RawDescriptor: file_protos_topk_topk_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_topk_topk_proto_goTypes,
		DependencyIndexes: file_protos_topk_topk_proto_depIdxs,
		EnumInfos:         file_protos_topk_topk_proto_enumTypes,
		MessageInfos:      file_protos_topk_topk_proto_msgTypes,
	}.Build()
	File_protos_topk_topk_proto = out.File
	file_protos_topk_topk_proto_rawDesc = nil
	file_protos_topk_topk_proto_goTypes = nil
	file_protos_topk_topk_proto_depIdxs = nil
}
