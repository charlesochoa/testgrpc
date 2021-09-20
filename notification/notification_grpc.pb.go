// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package testgrpc

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

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationServiceClient interface {
	Send(ctx context.Context, in *SendItemReq, opts ...grpc.CallOption) (*SendItemResp, error)
	SendAgain(ctx context.Context, in *SendItemReq, opts ...grpc.CallOption) (*SendItemResp, error)
	Click(ctx context.Context, in *ClickReq, opts ...grpc.CallOption) (*ClickResp, error)
}

type notificationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationServiceClient(cc grpc.ClientConnInterface) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Send(ctx context.Context, in *SendItemReq, opts ...grpc.CallOption) (*SendItemResp, error) {
	out := new(SendItemResp)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) SendAgain(ctx context.Context, in *SendItemReq, opts ...grpc.CallOption) (*SendItemResp, error) {
	out := new(SendItemResp)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/SendAgain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) Click(ctx context.Context, in *ClickReq, opts ...grpc.CallOption) (*ClickResp, error) {
	out := new(ClickResp)
	err := c.cc.Invoke(ctx, "/notification.NotificationService/Click", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
// All implementations must embed UnimplementedNotificationServiceServer
// for forward compatibility
type NotificationServiceServer interface {
	Send(context.Context, *SendItemReq) (*SendItemResp, error)
	SendAgain(context.Context, *SendItemReq) (*SendItemResp, error)
	Click(context.Context, *ClickReq) (*ClickResp, error)
	mustEmbedUnimplementedNotificationServiceServer()
}

// UnimplementedNotificationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (UnimplementedNotificationServiceServer) Send(context.Context, *SendItemReq) (*SendItemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedNotificationServiceServer) SendAgain(context.Context, *SendItemReq) (*SendItemResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendAgain not implemented")
}
func (UnimplementedNotificationServiceServer) Click(context.Context, *ClickReq) (*ClickResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Click not implemented")
}
func (UnimplementedNotificationServiceServer) mustEmbedUnimplementedNotificationServiceServer() {}

// UnsafeNotificationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationServiceServer will
// result in compilation errors.
type UnsafeNotificationServiceServer interface {
	mustEmbedUnimplementedNotificationServiceServer()
}

func RegisterNotificationServiceServer(s grpc.ServiceRegistrar, srv NotificationServiceServer) {
	s.RegisterService(&NotificationService_ServiceDesc, srv)
}

func _NotificationService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Send(ctx, req.(*SendItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_SendAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).SendAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/SendAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).SendAgain(ctx, req.(*SendItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_Click_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClickReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Click(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/notification.NotificationService/Click",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Click(ctx, req.(*ClickReq))
	}
	return interceptor(ctx, in, info, handler)
}

// NotificationService_ServiceDesc is the grpc.ServiceDesc for NotificationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NotificationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notification.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _NotificationService_Send_Handler,
		},
		{
			MethodName: "SendAgain",
			Handler:    _NotificationService_SendAgain_Handler,
		},
		{
			MethodName: "Click",
			Handler:    _NotificationService_Click_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notification/notification.proto",
}
