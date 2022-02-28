// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PushMessagerClient is the client API for PushMessager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushMessagerClient interface {
	PushMessage(ctx context.Context, in *PushMessageReq, opts ...grpc.CallOption) (*PushMessageResp, error)
}

type pushMessagerClient struct {
	cc grpc.ClientConnInterface
}

func NewPushMessagerClient(cc grpc.ClientConnInterface) PushMessagerClient {
	return &pushMessagerClient{cc}
}

func (c *pushMessagerClient) PushMessage(ctx context.Context, in *PushMessageReq, opts ...grpc.CallOption) (*PushMessageResp, error) {
	out := new(PushMessageResp)
	err := c.cc.Invoke(ctx, "/message.api.v1.PushMessager/PushMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushMessagerServer is the server API for PushMessager service.
// All implementations must embed UnimplementedPushMessagerServer
// for forward compatibility
type PushMessagerServer interface {
	PushMessage(context.Context, *PushMessageReq) (*PushMessageResp, error)
	mustEmbedUnimplementedPushMessagerServer()
}

// UnimplementedPushMessagerServer must be embedded to have forward compatible implementations.
type UnimplementedPushMessagerServer struct {
}

func (UnimplementedPushMessagerServer) PushMessage(context.Context, *PushMessageReq) (*PushMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMessage not implemented")
}
func (UnimplementedPushMessagerServer) mustEmbedUnimplementedPushMessagerServer() {}

// UnsafePushMessagerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushMessagerServer will
// result in compilation errors.
type UnsafePushMessagerServer interface {
	mustEmbedUnimplementedPushMessagerServer()
}

func RegisterPushMessagerServer(s grpc.ServiceRegistrar, srv PushMessagerServer) {
	s.RegisterService(&PushMessager_ServiceDesc, srv)
}

func _PushMessager_PushMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushMessagerServer).PushMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.api.v1.PushMessager/PushMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushMessagerServer).PushMessage(ctx, req.(*PushMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PushMessager_ServiceDesc is the grpc.ServiceDesc for PushMessager service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PushMessager_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "message.api.v1.PushMessager",
	HandlerType: (*PushMessagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushMessage",
			Handler:    _PushMessager_PushMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/message/v1/message.proto",
}
