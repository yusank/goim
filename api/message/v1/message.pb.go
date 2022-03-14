// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: api/message/v1/message.proto

package v1

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

type MessageContentType int32

const (
	MessageContentType_UnknownContentType MessageContentType = 0
	MessageContentType_Text               MessageContentType = 1
	MessageContentType_Image              MessageContentType = 2
	MessageContentType_Voice              MessageContentType = 3
)

// Enum value maps for MessageContentType.
var (
	MessageContentType_name = map[int32]string{
		0: "UnknownContentType",
		1: "Text",
		2: "Image",
		3: "Voice",
	}
	MessageContentType_value = map[string]int32{
		"UnknownContentType": 0,
		"Text":               1,
		"Image":              2,
		"Voice":              3,
	}
)

func (x MessageContentType) Enum() *MessageContentType {
	p := new(MessageContentType)
	*p = x
	return p
}

func (x MessageContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_message_v1_message_proto_enumTypes[0].Descriptor()
}

func (MessageContentType) Type() protoreflect.EnumType {
	return &file_api_message_v1_message_proto_enumTypes[0]
}

func (x MessageContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageContentType.Descriptor instead.
func (MessageContentType) EnumDescriptor() ([]byte, []int) {
	return file_api_message_v1_message_proto_rawDescGZIP(), []int{0}
}

type PushMessageType int32

const (
	PushMessageType_UnknownUserType PushMessageType = 0
	// user to user
	PushMessageType_User PushMessageType = 1
	// user to group
	PushMessageType_Group PushMessageType = 2
	// global broadcast
	PushMessageType_Broadcast PushMessageType = 3
)

// Enum value maps for PushMessageType.
var (
	PushMessageType_name = map[int32]string{
		0: "UnknownUserType",
		1: "User",
		2: "Group",
		3: "Broadcast",
	}
	PushMessageType_value = map[string]int32{
		"UnknownUserType": 0,
		"User":            1,
		"Group":           2,
		"Broadcast":       3,
	}
)

func (x PushMessageType) Enum() *PushMessageType {
	p := new(PushMessageType)
	*p = x
	return p
}

func (x PushMessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PushMessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_message_v1_message_proto_enumTypes[1].Descriptor()
}

func (PushMessageType) Type() protoreflect.EnumType {
	return &file_api_message_v1_message_proto_enumTypes[1]
}

func (x PushMessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PushMessageType.Descriptor instead.
func (PushMessageType) EnumDescriptor() ([]byte, []int) {
	return file_api_message_v1_message_proto_rawDescGZIP(), []int{1}
}

type PushMessageRespStatus int32

const (
	PushMessageRespStatus_OK                 PushMessageRespStatus = 0
	PushMessageRespStatus_ConnectionNotFound PushMessageRespStatus = 1
	PushMessageRespStatus_ConnectionLost     PushMessageRespStatus = 2
	PushMessageRespStatus_ConnectionTimeout  PushMessageRespStatus = 3
	PushMessageRespStatus_WriteTimeout       PushMessageRespStatus = 4
	PushMessageRespStatus_Unknown            PushMessageRespStatus = 5
)

// Enum value maps for PushMessageRespStatus.
var (
	PushMessageRespStatus_name = map[int32]string{
		0: "OK",
		1: "ConnectionNotFound",
		2: "ConnectionLost",
		3: "ConnectionTimeout",
		4: "WriteTimeout",
		5: "Unknown",
	}
	PushMessageRespStatus_value = map[string]int32{
		"OK":                 0,
		"ConnectionNotFound": 1,
		"ConnectionLost":     2,
		"ConnectionTimeout":  3,
		"WriteTimeout":       4,
		"Unknown":            5,
	}
)

func (x PushMessageRespStatus) Enum() *PushMessageRespStatus {
	p := new(PushMessageRespStatus)
	*p = x
	return p
}

func (x PushMessageRespStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PushMessageRespStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_api_message_v1_message_proto_enumTypes[2].Descriptor()
}

func (PushMessageRespStatus) Type() protoreflect.EnumType {
	return &file_api_message_v1_message_proto_enumTypes[2]
}

func (x PushMessageRespStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PushMessageRespStatus.Descriptor instead.
func (PushMessageRespStatus) EnumDescriptor() ([]byte, []int) {
	return file_api_message_v1_message_proto_rawDescGZIP(), []int{2}
}

// SendMessageReq receive data from gateway
type SendMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUser    string             `protobuf:"bytes,1,opt,name=from_user,json=fromUser,proto3" json:"from_user,omitempty"`
	ToUser      string             `protobuf:"bytes,2,opt,name=to_user,json=toUser,proto3" json:"to_user,omitempty"`
	ContentType MessageContentType `protobuf:"varint,3,opt,name=content_type,json=contentType,proto3,enum=api.message.v1.MessageContentType" json:"content_type,omitempty"`
	Content     string             `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendMessageReq) Reset() {
	*x = SendMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_message_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageReq) ProtoMessage() {}

func (x *SendMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_message_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageReq.ProtoReflect.Descriptor instead.
func (*SendMessageReq) Descriptor() ([]byte, []int) {
	return file_api_message_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageReq) GetFromUser() string {
	if x != nil {
		return x.FromUser
	}
	return ""
}

func (x *SendMessageReq) GetToUser() string {
	if x != nil {
		return x.ToUser
	}
	return ""
}

func (x *SendMessageReq) GetContentType() MessageContentType {
	if x != nil {
		return x.ContentType
	}
	return MessageContentType_UnknownContentType
}

func (x *SendMessageReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// PushMessage use for push a message to persistence connection server
type PushMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUser        string             `protobuf:"bytes,1,opt,name=from_user,json=fromUser,proto3" json:"from_user,omitempty"`
	ToUser          string             `protobuf:"bytes,2,opt,name=to_user,json=toUser,proto3" json:"to_user,omitempty"`
	AgentId         string             `protobuf:"bytes,3,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	PushMessageType PushMessageType    `protobuf:"varint,4,opt,name=push_message_type,json=pushMessageType,proto3,enum=api.message.v1.PushMessageType" json:"push_message_type,omitempty"`
	ContentType     MessageContentType `protobuf:"varint,5,opt,name=content_type,json=contentType,proto3,enum=api.message.v1.MessageContentType" json:"content_type,omitempty"`
	Content         string             `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *PushMessageReq) Reset() {
	*x = PushMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_message_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMessageReq) ProtoMessage() {}

func (x *PushMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_message_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMessageReq.ProtoReflect.Descriptor instead.
func (*PushMessageReq) Descriptor() ([]byte, []int) {
	return file_api_message_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *PushMessageReq) GetFromUser() string {
	if x != nil {
		return x.FromUser
	}
	return ""
}

func (x *PushMessageReq) GetToUser() string {
	if x != nil {
		return x.ToUser
	}
	return ""
}

func (x *PushMessageReq) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *PushMessageReq) GetPushMessageType() PushMessageType {
	if x != nil {
		return x.PushMessageType
	}
	return PushMessageType_UnknownUserType
}

func (x *PushMessageReq) GetContentType() MessageContentType {
	if x != nil {
		return x.ContentType
	}
	return MessageContentType_UnknownContentType
}

func (x *PushMessageReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type PushMessageResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status PushMessageRespStatus `protobuf:"varint,1,opt,name=status,proto3,enum=api.message.v1.PushMessageRespStatus" json:"status,omitempty"`
	Reason string                `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
}

func (x *PushMessageResp) Reset() {
	*x = PushMessageResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_message_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMessageResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMessageResp) ProtoMessage() {}

func (x *PushMessageResp) ProtoReflect() protoreflect.Message {
	mi := &file_api_message_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMessageResp.ProtoReflect.Descriptor instead.
func (*PushMessageResp) Descriptor() ([]byte, []int) {
	return file_api_message_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *PushMessageResp) GetStatus() PushMessageRespStatus {
	if x != nil {
		return x.Status
	}
	return PushMessageRespStatus_OK
}

func (x *PushMessageResp) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

var File_api_message_v1_message_proto protoreflect.FileDescriptor

var file_api_message_v1_message_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xa7,
	0x01, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x45, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x8f, 0x02, 0x0a, 0x0e, 0x50, 0x75, 0x73,
	0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x11,
	0x70, 0x75, 0x73, 0x68, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f, 0x70, 0x75, 0x73, 0x68, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x68, 0x0a, 0x0f, 0x50, 0x75,
	0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x3d, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x2a, 0x4c, 0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x55, 0x6e,
	0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x6f, 0x69, 0x63, 0x65,
	0x10, 0x03, 0x2a, 0x4a, 0x0a, 0x0f, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x10, 0x02, 0x12,
	0x0d, 0x0a, 0x09, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x10, 0x03, 0x2a, 0x81,
	0x01, 0x0a, 0x15, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00,
	0x12, 0x16, 0x0a, 0x12, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x6f,
	0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x6f, 0x73, 0x74, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x05, 0x32, 0x5e, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x4e, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x42, 0x13, 0x5a, 0x11, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_message_v1_message_proto_rawDescOnce sync.Once
	file_api_message_v1_message_proto_rawDescData = file_api_message_v1_message_proto_rawDesc
)

func file_api_message_v1_message_proto_rawDescGZIP() []byte {
	file_api_message_v1_message_proto_rawDescOnce.Do(func() {
		file_api_message_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_message_v1_message_proto_rawDescData)
	})
	return file_api_message_v1_message_proto_rawDescData
}

var file_api_message_v1_message_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_api_message_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_message_v1_message_proto_goTypes = []interface{}{
	(MessageContentType)(0),    // 0: api.message.v1.MessageContentType
	(PushMessageType)(0),       // 1: api.message.v1.PushMessageType
	(PushMessageRespStatus)(0), // 2: api.message.v1.PushMessageRespStatus
	(*SendMessageReq)(nil),     // 3: api.message.v1.SendMessageReq
	(*PushMessageReq)(nil),     // 4: api.message.v1.PushMessageReq
	(*PushMessageResp)(nil),    // 5: api.message.v1.PushMessageResp
}
var file_api_message_v1_message_proto_depIdxs = []int32{
	0, // 0: api.message.v1.SendMessageReq.content_type:type_name -> api.message.v1.MessageContentType
	1, // 1: api.message.v1.PushMessageReq.push_message_type:type_name -> api.message.v1.PushMessageType
	0, // 2: api.message.v1.PushMessageReq.content_type:type_name -> api.message.v1.MessageContentType
	2, // 3: api.message.v1.PushMessageResp.status:type_name -> api.message.v1.PushMessageRespStatus
	4, // 4: api.message.v1.PushMessager.PushMessage:input_type -> api.message.v1.PushMessageReq
	5, // 5: api.message.v1.PushMessager.PushMessage:output_type -> api.message.v1.PushMessageResp
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_message_v1_message_proto_init() }
func file_api_message_v1_message_proto_init() {
	if File_api_message_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_message_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageReq); i {
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
		file_api_message_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMessageReq); i {
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
		file_api_message_v1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMessageResp); i {
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
			RawDescriptor: file_api_message_v1_message_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_message_v1_message_proto_goTypes,
		DependencyIndexes: file_api_message_v1_message_proto_depIdxs,
		EnumInfos:         file_api_message_v1_message_proto_enumTypes,
		MessageInfos:      file_api_message_v1_message_proto_msgTypes,
	}.Build()
	File_api_message_v1_message_proto = out.File
	file_api_message_v1_message_proto_rawDesc = nil
	file_api_message_v1_message_proto_goTypes = nil
	file_api_message_v1_message_proto_depIdxs = nil
}
