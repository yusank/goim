// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: api/config/v1/config.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Level int32

const (
	Level_DEBUG   Level = 0
	Level_INFO    Level = 1
	Level_WARNING Level = 2
	Level_ERROR   Level = 3
	Level_FATAL   Level = 4
)

// Enum value maps for Level.
var (
	Level_name = map[int32]string{
		0: "DEBUG",
		1: "INFO",
		2: "WARNING",
		3: "ERROR",
		4: "FATAL",
	}
	Level_value = map[string]int32{
		"DEBUG":   0,
		"INFO":    1,
		"WARNING": 2,
		"ERROR":   3,
		"FATAL":   4,
	}
)

func (x Level) Enum() *Level {
	p := new(Level)
	*p = x
	return p
}

func (x Level) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Level) Descriptor() protoreflect.EnumDescriptor {
	return file_api_config_v1_config_proto_enumTypes[0].Descriptor()
}

func (Level) Type() protoreflect.EnumType {
	return &file_api_config_v1_config_proto_enumTypes[0]
}

func (x Level) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Level.Descriptor instead.
func (Level) EnumDescriptor() ([]byte, []int) {
	return file_api_config_v1_config_proto_rawDescGZIP(), []int{0}
}

type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Scheme string `protobuf:"bytes,1,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Addr   string `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Port   int32  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_config_v1_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_api_config_v1_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_api_config_v1_config_proto_rawDescGZIP(), []int{0}
}

func (x *Server) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *Server) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Server) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version  string            `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Http     *Server           `protobuf:"bytes,3,opt,name=http,proto3,oneof" json:"http,omitempty"`
	Grpc     *Server           `protobuf:"bytes,4,opt,name=grpc,proto3,oneof" json:"grpc,omitempty"`
	Log      *Log              `protobuf:"bytes,5,opt,name=log,proto3" json:"log,omitempty"`
	Metadata map[string]string `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Redis    *Redis            `protobuf:"bytes,7,opt,name=redis,proto3" json:"redis,omitempty"`
	Mq       *MQ               `protobuf:"bytes,8,opt,name=mq,proto3" json:"mq,omitempty"`
	Mysql    *MySQL            `protobuf:"bytes,9,opt,name=mysql,proto3" json:"mysql,omitempty"`
	// services name
	GatewayService string `protobuf:"bytes,10,opt,name=gatewayService,proto3" json:"gatewayService,omitempty"`
	UserService    string `protobuf:"bytes,11,opt,name=userService,proto3" json:"userService,omitempty"`
	PushService    string `protobuf:"bytes,12,opt,name=pushService,proto3" json:"pushService,omitempty"`
	StoreService   string `protobuf:"bytes,13,opt,name=storeService,proto3" json:"storeService,omitempty"`
	MsgService     string `protobuf:"bytes,14,opt,name=msgService,proto3" json:"msgService,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_config_v1_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_api_config_v1_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_api_config_v1_config_proto_rawDescGZIP(), []int{1}
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Service) GetHttp() *Server {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *Service) GetGrpc() *Server {
	if x != nil {
		return x.Grpc
	}
	return nil
}

func (x *Service) GetLog() *Log {
	if x != nil {
		return x.Log
	}
	return nil
}

func (x *Service) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Service) GetRedis() *Redis {
	if x != nil {
		return x.Redis
	}
	return nil
}

func (x *Service) GetMq() *MQ {
	if x != nil {
		return x.Mq
	}
	return nil
}

func (x *Service) GetMysql() *MySQL {
	if x != nil {
		return x.Mysql
	}
	return nil
}

func (x *Service) GetGatewayService() string {
	if x != nil {
		return x.GatewayService
	}
	return ""
}

func (x *Service) GetUserService() string {
	if x != nil {
		return x.UserService
	}
	return ""
}

func (x *Service) GetPushService() string {
	if x != nil {
		return x.PushService
	}
	return ""
}

func (x *Service) GetStoreService() string {
	if x != nil {
		return x.StoreService
	}
	return ""
}

func (x *Service) GetMsgService() string {
	if x != nil {
		return x.MsgService
	}
	return ""
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogPath       *string `protobuf:"bytes,1,opt,name=log_path,json=logPath,proto3,oneof" json:"log_path,omitempty"`
	Level         Level   `protobuf:"varint,2,opt,name=level,proto3,enum=api.config.v1.Level" json:"level,omitempty"`
	EnableConsole bool    `protobuf:"varint,3,opt,name=enable_console,json=enableConsole,proto3" json:"enable_console,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_config_v1_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_api_config_v1_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_api_config_v1_config_proto_rawDescGZIP(), []int{2}
}

func (x *Log) GetLogPath() string {
	if x != nil && x.LogPath != nil {
		return *x.LogPath
	}
	return ""
}

func (x *Log) GetLevel() Level {
	if x != nil {
		return x.Level
	}
	return Level_DEBUG
}

func (x *Log) GetEnableConsole() bool {
	if x != nil {
		return x.EnableConsole
	}
	return false
}

type Redis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr         string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	Password     string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	MaxConns     int32  `protobuf:"varint,3,opt,name=max_conns,json=maxConns,proto3" json:"max_conns,omitempty"`
	MinIdleConns int32  `protobuf:"varint,4,opt,name=min_idle_conns,json=minIdleConns,proto3" json:"min_idle_conns,omitempty"`
	// range: [10ms,10s]
	DialTimeout *durationpb.Duration `protobuf:"bytes,5,opt,name=dial_timeout,json=dialTimeout,proto3" json:"dial_timeout,omitempty"`
	// range: [10ms,10s]
	IdleTimeout *durationpb.Duration `protobuf:"bytes,6,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
}

func (x *Redis) Reset() {
	*x = Redis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_config_v1_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Redis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Redis) ProtoMessage() {}

func (x *Redis) ProtoReflect() protoreflect.Message {
	mi := &file_api_config_v1_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Redis.ProtoReflect.Descriptor instead.
func (*Redis) Descriptor() ([]byte, []int) {
	return file_api_config_v1_config_proto_rawDescGZIP(), []int{3}
}

func (x *Redis) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Redis) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Redis) GetMaxConns() int32 {
	if x != nil {
		return x.MaxConns
	}
	return 0
}

func (x *Redis) GetMinIdleConns() int32 {
	if x != nil {
		return x.MinIdleConns
	}
	return 0
}

func (x *Redis) GetDialTimeout() *durationpb.Duration {
	if x != nil {
		return x.DialTimeout
	}
	return nil
}

func (x *Redis) GetIdleTimeout() *durationpb.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

type MQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr     []string `protobuf:"bytes,1,rep,name=addr,proto3" json:"addr,omitempty"`
	MaxRetry int32    `protobuf:"varint,2,opt,name=max_retry,json=maxRetry,proto3" json:"max_retry,omitempty"`
}

func (x *MQ) Reset() {
	*x = MQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_config_v1_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MQ) ProtoMessage() {}

func (x *MQ) ProtoReflect() protoreflect.Message {
	mi := &file_api_config_v1_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MQ.ProtoReflect.Descriptor instead.
func (*MQ) Descriptor() ([]byte, []int) {
	return file_api_config_v1_config_proto_rawDescGZIP(), []int{4}
}

func (x *MQ) GetAddr() []string {
	if x != nil {
		return x.Addr
	}
	return nil
}

func (x *MQ) GetMaxRetry() int32 {
	if x != nil {
		return x.MaxRetry
	}
	return 0
}

type MySQL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Addr         string `protobuf:"bytes,1,opt,name=addr,proto3" json:"addr,omitempty"`
	User         string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	Password     string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Db           string `protobuf:"bytes,4,opt,name=db,proto3" json:"db,omitempty"`
	MaxIdleConns int32  `protobuf:"varint,5,opt,name=max_idle_conns,json=maxIdleConns,proto3" json:"max_idle_conns,omitempty"`
	MaxOpenConns int32  `protobuf:"varint,6,opt,name=max_open_conns,json=maxOpenConns,proto3" json:"max_open_conns,omitempty"`
	// range: [10ms,10s]
	IdleTimeout *durationpb.Duration `protobuf:"bytes,7,opt,name=idle_timeout,json=idleTimeout,proto3" json:"idle_timeout,omitempty"`
	// range: [10ms,10s]
	OpenTimeout *durationpb.Duration `protobuf:"bytes,8,opt,name=open_timeout,json=openTimeout,proto3" json:"open_timeout,omitempty"`
}

func (x *MySQL) Reset() {
	*x = MySQL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_config_v1_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MySQL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MySQL) ProtoMessage() {}

func (x *MySQL) ProtoReflect() protoreflect.Message {
	mi := &file_api_config_v1_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MySQL.ProtoReflect.Descriptor instead.
func (*MySQL) Descriptor() ([]byte, []int) {
	return file_api_config_v1_config_proto_rawDescGZIP(), []int{5}
}

func (x *MySQL) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *MySQL) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *MySQL) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *MySQL) GetDb() string {
	if x != nil {
		return x.Db
	}
	return ""
}

func (x *MySQL) GetMaxIdleConns() int32 {
	if x != nil {
		return x.MaxIdleConns
	}
	return 0
}

func (x *MySQL) GetMaxOpenConns() int32 {
	if x != nil {
		return x.MaxOpenConns
	}
	return 0
}

func (x *MySQL) GetIdleTimeout() *durationpb.Duration {
	if x != nil {
		return x.IdleTimeout
	}
	return nil
}

func (x *MySQL) GetOpenTimeout() *durationpb.Duration {
	if x != nil {
		return x.OpenTimeout
	}
	return nil
}

var File_api_config_v1_config_proto protoreflect.FileDescriptor

var file_api_config_v1_config_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2e,
	0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16,
	0xfa, 0x42, 0x13, 0x72, 0x11, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x52, 0x04, 0x67, 0x72, 0x70,
	0x63, 0x52, 0x03, 0x74, 0x63, 0x70, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x1f,
	0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42,
	0x08, 0x72, 0x06, 0xa8, 0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12,
	0x20, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0c, 0xfa,
	0x42, 0x09, 0x1a, 0x07, 0x10, 0xf7, 0xd8, 0x03, 0x20, 0x90, 0x4e, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x22, 0x9c, 0x06, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xfa, 0x42, 0x10,
	0x72, 0x0e, 0x3a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x6f, 0x69, 0x6d,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x68, 0x74, 0x74,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x00,
	0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x04, 0x67, 0x72, 0x70,
	0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x48, 0x01,
	0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x03, 0x6c, 0x6f, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x03, 0x6c, 0x6f, 0x67, 0x12,
	0x40, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x2a, 0x0a, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x52, 0x05, 0x72, 0x65, 0x64, 0x69, 0x73, 0x12, 0x21, 0x0a,
	0x02, 0x6d, 0x71, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x51, 0x52, 0x02, 0x6d, 0x71,
	0x12, 0x2a, 0x0a, 0x05, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x79, 0x53, 0x51, 0x4c, 0x52, 0x05, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x12, 0x43, 0x0a, 0x0e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xfa, 0x42, 0x18, 0x72, 0x16, 0x0a, 0x14, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x6f, 0x69, 0x6d, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x52, 0x0e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0xfa, 0x42, 0x15, 0x72, 0x13, 0x0a, 0x11, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x6f, 0x69, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x52, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a,
	0x0b, 0x70, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x18, 0xfa, 0x42, 0x15, 0x72, 0x13, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x67, 0x6f, 0x69, 0x6d, 0x2e, 0x70, 0x75, 0x73, 0x68, 0x52, 0x0b, 0x70, 0x75,
	0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x74, 0x6f,
	0x72, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x19, 0xfa, 0x42, 0x16, 0x72, 0x14, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x67, 0x6f, 0x69, 0x6d, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x0c, 0x73, 0x74, 0x6f, 0x72,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x6d, 0x73, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xfa, 0x42,
	0x14, 0x72, 0x12, 0x0a, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x67, 0x6f, 0x69,
	0x6d, 0x2e, 0x6d, 0x73, 0x67, 0x52, 0x0a, 0x6d, 0x73, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x22, 0x8f, 0x01, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x1e, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x5f,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x6f,
	0x67, 0x50, 0x61, 0x74, 0x68, 0x88, 0x01, 0x01, 0x12, 0x34, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x42, 0x08, 0xfa,
	0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x25,
	0x0a, 0x0e, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x22, 0x9a, 0x02, 0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x6d, 0x61, 0x78, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x6d, 0x61, 0x78, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x12,
	0x4e, 0x0a, 0x0c, 0x64, 0x69, 0x61, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x10, 0xfa, 0x42, 0x0d, 0xaa, 0x01, 0x0a, 0x22, 0x02, 0x08, 0x0a, 0x32, 0x04, 0x10, 0xc0,
	0x84, 0x3d, 0x52, 0x0b, 0x64, 0x69, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12,
	0x4e, 0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x42, 0x10, 0xfa, 0x42, 0x0d, 0xaa, 0x01, 0x0a, 0x22, 0x02, 0x08, 0x0a, 0x32, 0x04, 0x10, 0xc0,
	0x84, 0x3d, 0x52, 0x0b, 0x69, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22,
	0x35, 0x0a, 0x02, 0x4d, 0x51, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x78,
	0x5f, 0x72, 0x65, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6d, 0x61,
	0x78, 0x52, 0x65, 0x74, 0x72, 0x79, 0x22, 0xc7, 0x02, 0x0a, 0x05, 0x4d, 0x79, 0x53, 0x51, 0x4c,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x64, 0x64, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x64, 0x62, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x78, 0x5f, 0x69, 0x64, 0x6c, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61,
	0x78, 0x49, 0x64, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61,
	0x78, 0x5f, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x4f, 0x70, 0x65, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x73,
	0x12, 0x4e, 0x0a, 0x0c, 0x69, 0x64, 0x6c, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0xaa, 0x01, 0x0a, 0x22, 0x02, 0x08, 0x0a, 0x32, 0x04, 0x10,
	0xc0, 0x84, 0x3d, 0x52, 0x0b, 0x69, 0x64, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x12, 0x4e, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x10, 0xfa, 0x42, 0x0d, 0xaa, 0x01, 0x0a, 0x22, 0x02, 0x08, 0x0a, 0x32, 0x04, 0x10,
	0xc0, 0x84, 0x3d, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x2a, 0x3f, 0x0a, 0x05, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x45, 0x42,
	0x55, 0x47, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x46, 0x4f, 0x10, 0x01, 0x12, 0x0b,
	0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x41, 0x54, 0x41, 0x4c, 0x10,
	0x04, 0x42, 0x12, 0x5a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_config_v1_config_proto_rawDescOnce sync.Once
	file_api_config_v1_config_proto_rawDescData = file_api_config_v1_config_proto_rawDesc
)

func file_api_config_v1_config_proto_rawDescGZIP() []byte {
	file_api_config_v1_config_proto_rawDescOnce.Do(func() {
		file_api_config_v1_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_config_v1_config_proto_rawDescData)
	})
	return file_api_config_v1_config_proto_rawDescData
}

var file_api_config_v1_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_config_v1_config_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_config_v1_config_proto_goTypes = []interface{}{
	(Level)(0),                  // 0: api.config.v1.Level
	(*Server)(nil),              // 1: api.config.v1.Server
	(*Service)(nil),             // 2: api.config.v1.Service
	(*Log)(nil),                 // 3: api.config.v1.Log
	(*Redis)(nil),               // 4: api.config.v1.Redis
	(*MQ)(nil),                  // 5: api.config.v1.MQ
	(*MySQL)(nil),               // 6: api.config.v1.MySQL
	nil,                         // 7: api.config.v1.Service.MetadataEntry
	(*durationpb.Duration)(nil), // 8: google.protobuf.Duration
}
var file_api_config_v1_config_proto_depIdxs = []int32{
	1,  // 0: api.config.v1.Service.http:type_name -> api.config.v1.Server
	1,  // 1: api.config.v1.Service.grpc:type_name -> api.config.v1.Server
	3,  // 2: api.config.v1.Service.log:type_name -> api.config.v1.Log
	7,  // 3: api.config.v1.Service.metadata:type_name -> api.config.v1.Service.MetadataEntry
	4,  // 4: api.config.v1.Service.redis:type_name -> api.config.v1.Redis
	5,  // 5: api.config.v1.Service.mq:type_name -> api.config.v1.MQ
	6,  // 6: api.config.v1.Service.mysql:type_name -> api.config.v1.MySQL
	0,  // 7: api.config.v1.Log.level:type_name -> api.config.v1.Level
	8,  // 8: api.config.v1.Redis.dial_timeout:type_name -> google.protobuf.Duration
	8,  // 9: api.config.v1.Redis.idle_timeout:type_name -> google.protobuf.Duration
	8,  // 10: api.config.v1.MySQL.idle_timeout:type_name -> google.protobuf.Duration
	8,  // 11: api.config.v1.MySQL.open_timeout:type_name -> google.protobuf.Duration
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_api_config_v1_config_proto_init() }
func file_api_config_v1_config_proto_init() {
	if File_api_config_v1_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_config_v1_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Server); i {
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
		file_api_config_v1_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
		file_api_config_v1_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_api_config_v1_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Redis); i {
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
		file_api_config_v1_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MQ); i {
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
		file_api_config_v1_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MySQL); i {
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
	file_api_config_v1_config_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_api_config_v1_config_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_config_v1_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_config_v1_config_proto_goTypes,
		DependencyIndexes: file_api_config_v1_config_proto_depIdxs,
		EnumInfos:         file_api_config_v1_config_proto_enumTypes,
		MessageInfos:      file_api_config_v1_config_proto_msgTypes,
	}.Build()
	File_api_config_v1_config_proto = out.File
	file_api_config_v1_config_proto_rawDesc = nil
	file_api_config_v1_config_proto_goTypes = nil
	file_api_config_v1_config_proto_depIdxs = nil
}
