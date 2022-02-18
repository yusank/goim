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
	MessageContentType_Text    MessageContentType = 0
	MessageContentType_Image   MessageContentType = 1
	MessageContentType_Voice   MessageContentType = 2
	MessageContentType_UNKNOWN MessageContentType = -1
)

// Enum value maps for MessageContentType.
var (
	MessageContentType_name = map[int32]string{
		0:  "Text",
		1:  "Image",
		2:  "Voice",
		-1: "UNKNOWN",
	}
	MessageContentType_value = map[string]int32{
		"Text":    0,
		"Image":   1,
		"Voice":   2,
		"UNKNOWN": -1,
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
	// user to user
	PushMessageType_ToUser PushMessageType = 0
	// user to group
	PushMessageType_ToGroup PushMessageType = 1
	// gloabl broadcast
	PushMessageType_ToChannel PushMessageType = 2
)

// Enum value maps for PushMessageType.
var (
	PushMessageType_name = map[int32]string{
		0: "ToUser",
		1: "ToGroup",
		2: "ToChannel",
	}
	PushMessageType_value = map[string]int32{
		"ToUser":    0,
		"ToGroup":   1,
		"ToChannel": 2,
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

// Message receive data from gateway
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUser    string             `protobuf:"bytes,1,opt,name=from_user,json=fromUser,proto3" json:"from_user,omitempty"`
	ToUser      string             `protobuf:"bytes,2,opt,name=to_user,json=toUser,proto3" json:"to_user,omitempty"`
	ContentType MessageContentType `protobuf:"varint,3,opt,name=content_type,json=contentType,proto3,enum=message.api.v1.MessageContentType" json:"content_type,omitempty"`
	Content     string             `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_message_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_api_message_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetFromUser() string {
	if x != nil {
		return x.FromUser
	}
	return ""
}

func (x *Message) GetToUser() string {
	if x != nil {
		return x.ToUser
	}
	return ""
}

func (x *Message) GetContentType() MessageContentType {
	if x != nil {
		return x.ContentType
	}
	return MessageContentType_Text
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// PushMessage use for push a message to persistence connection server
type PushMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromUser        string             `protobuf:"bytes,1,opt,name=from_user,json=fromUser,proto3" json:"from_user,omitempty"`
	ToUser          string             `protobuf:"bytes,2,opt,name=to_user,json=toUser,proto3" json:"to_user,omitempty"`
	PushMessageType PushMessageType    `protobuf:"varint,3,opt,name=push_message_type,json=pushMessageType,proto3,enum=message.api.v1.PushMessageType" json:"push_message_type,omitempty"`
	ContentType     MessageContentType `protobuf:"varint,4,opt,name=content_type,json=contentType,proto3,enum=message.api.v1.MessageContentType" json:"content_type,omitempty"`
	Content         string             `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *PushMessage) Reset() {
	*x = PushMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_message_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMessage) ProtoMessage() {}

func (x *PushMessage) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use PushMessage.ProtoReflect.Descriptor instead.
func (*PushMessage) Descriptor() ([]byte, []int) {
	return file_api_message_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *PushMessage) GetFromUser() string {
	if x != nil {
		return x.FromUser
	}
	return ""
}

func (x *PushMessage) GetToUser() string {
	if x != nil {
		return x.ToUser
	}
	return ""
}

func (x *PushMessage) GetPushMessageType() PushMessageType {
	if x != nil {
		return x.PushMessageType
	}
	return PushMessageType_ToUser
}

func (x *PushMessage) GetContentType() MessageContentType {
	if x != nil {
		return x.ContentType
	}
	return MessageContentType_Text
}

func (x *PushMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_api_message_v1_message_proto protoreflect.FileDescriptor

var file_api_message_v1_message_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0xa0,
	0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x75, 0x73,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x45, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0xf1, 0x01, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17,
	0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x11, 0x70, 0x75, 0x73, 0x68, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0f, 0x70, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2a, 0x4a, 0x0a, 0x12, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x54,
	0x65, 0x78, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x10, 0x01,
	0x12, 0x09, 0x0a, 0x05, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x07, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0x01, 0x2a, 0x39, 0x0a, 0x0f, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x10, 0x00,
	0x12, 0x0b, 0x0a, 0x07, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x10, 0x01, 0x12, 0x0d, 0x0a,
	0x09, 0x54, 0x6f, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x10, 0x02, 0x42, 0x13, 0x5a, 0x11,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_api_message_v1_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_message_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_message_v1_message_proto_goTypes = []interface{}{
	(MessageContentType)(0), // 0: message.api.v1.MessageContentType
	(PushMessageType)(0),    // 1: message.api.v1.PushMessageType
	(*Message)(nil),         // 2: message.api.v1.Message
	(*PushMessage)(nil),     // 3: message.api.v1.PushMessage
}
var file_api_message_v1_message_proto_depIdxs = []int32{
	0, // 0: message.api.v1.Message.content_type:type_name -> message.api.v1.MessageContentType
	1, // 1: message.api.v1.PushMessage.push_message_type:type_name -> message.api.v1.PushMessageType
	0, // 2: message.api.v1.PushMessage.content_type:type_name -> message.api.v1.MessageContentType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_message_v1_message_proto_init() }
func file_api_message_v1_message_proto_init() {
	if File_api_message_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_message_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			switch v := v.(*PushMessage); i {
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
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
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
