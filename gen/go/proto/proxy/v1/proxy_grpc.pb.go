// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/proxy/v1/proxy.proto

package proxyv1

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

const (
	ProxyService_GetTelemetry_FullMethodName       = "/proxy.v1.ProxyService/GetTelemetry"
	ProxyService_GetEnemyTelemetry_FullMethodName  = "/proxy.v1.ProxyService/GetEnemyTelemetry"
	ProxyService_SendLockPacket_FullMethodName     = "/proxy.v1.ProxyService/SendLockPacket"
	ProxyService_SendKamikazePacket_FullMethodName = "/proxy.v1.ProxyService/SendKamikazePacket"
)

// ProxyServiceClient is the client API for ProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyServiceClient interface {
	// Returns a normal telemetry stream when activated. Should be used by
	// services that don't interact with the competition server
	GetTelemetry(ctx context.Context, in *GetTelemetryRequest, opts ...grpc.CallOption) (ProxyService_GetTelemetryClient, error)
	// GetEnemyTelemetry returns the enemy's telemetry stream when
	// activated
	GetEnemyTelemetry(ctx context.Context, in *GetEnemyTelemetryRequest, opts ...grpc.CallOption) (ProxyService_GetEnemyTelemetryClient, error)
	// SendLockPacket sends a successful lock notification to the competition server
	SendLockPacket(ctx context.Context, in *SendLockPacketRequest, opts ...grpc.CallOption) (*SendLockPacketResponse, error)
	// SendKamikazePacket sends a successful kamikaze QR code reading to the competition server
	SendKamikazePacket(ctx context.Context, in *SendKamikazePacketRequest, opts ...grpc.CallOption) (*SendKamikazePacketResponse, error)
}

type proxyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyServiceClient(cc grpc.ClientConnInterface) ProxyServiceClient {
	return &proxyServiceClient{cc}
}

func (c *proxyServiceClient) GetTelemetry(ctx context.Context, in *GetTelemetryRequest, opts ...grpc.CallOption) (ProxyService_GetTelemetryClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProxyService_ServiceDesc.Streams[0], ProxyService_GetTelemetry_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyServiceGetTelemetryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProxyService_GetTelemetryClient interface {
	Recv() (*GetTelemetryResponse, error)
	grpc.ClientStream
}

type proxyServiceGetTelemetryClient struct {
	grpc.ClientStream
}

func (x *proxyServiceGetTelemetryClient) Recv() (*GetTelemetryResponse, error) {
	m := new(GetTelemetryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *proxyServiceClient) GetEnemyTelemetry(ctx context.Context, in *GetEnemyTelemetryRequest, opts ...grpc.CallOption) (ProxyService_GetEnemyTelemetryClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProxyService_ServiceDesc.Streams[1], ProxyService_GetEnemyTelemetry_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &proxyServiceGetEnemyTelemetryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProxyService_GetEnemyTelemetryClient interface {
	Recv() (*GetEnemyTelemetryResponse, error)
	grpc.ClientStream
}

type proxyServiceGetEnemyTelemetryClient struct {
	grpc.ClientStream
}

func (x *proxyServiceGetEnemyTelemetryClient) Recv() (*GetEnemyTelemetryResponse, error) {
	m := new(GetEnemyTelemetryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *proxyServiceClient) SendLockPacket(ctx context.Context, in *SendLockPacketRequest, opts ...grpc.CallOption) (*SendLockPacketResponse, error) {
	out := new(SendLockPacketResponse)
	err := c.cc.Invoke(ctx, ProxyService_SendLockPacket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyServiceClient) SendKamikazePacket(ctx context.Context, in *SendKamikazePacketRequest, opts ...grpc.CallOption) (*SendKamikazePacketResponse, error) {
	out := new(SendKamikazePacketResponse)
	err := c.cc.Invoke(ctx, ProxyService_SendKamikazePacket_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyServiceServer is the server API for ProxyService service.
// All implementations should embed UnimplementedProxyServiceServer
// for forward compatibility
type ProxyServiceServer interface {
	// Returns a normal telemetry stream when activated. Should be used by
	// services that don't interact with the competition server
	GetTelemetry(*GetTelemetryRequest, ProxyService_GetTelemetryServer) error
	// GetEnemyTelemetry returns the enemy's telemetry stream when
	// activated
	GetEnemyTelemetry(*GetEnemyTelemetryRequest, ProxyService_GetEnemyTelemetryServer) error
	// SendLockPacket sends a successful lock notification to the competition server
	SendLockPacket(context.Context, *SendLockPacketRequest) (*SendLockPacketResponse, error)
	// SendKamikazePacket sends a successful kamikaze QR code reading to the competition server
	SendKamikazePacket(context.Context, *SendKamikazePacketRequest) (*SendKamikazePacketResponse, error)
}

// UnimplementedProxyServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProxyServiceServer struct {
}

func (UnimplementedProxyServiceServer) GetTelemetry(*GetTelemetryRequest, ProxyService_GetTelemetryServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTelemetry not implemented")
}
func (UnimplementedProxyServiceServer) GetEnemyTelemetry(*GetEnemyTelemetryRequest, ProxyService_GetEnemyTelemetryServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEnemyTelemetry not implemented")
}
func (UnimplementedProxyServiceServer) SendLockPacket(context.Context, *SendLockPacketRequest) (*SendLockPacketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLockPacket not implemented")
}
func (UnimplementedProxyServiceServer) SendKamikazePacket(context.Context, *SendKamikazePacketRequest) (*SendKamikazePacketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendKamikazePacket not implemented")
}

// UnsafeProxyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyServiceServer will
// result in compilation errors.
type UnsafeProxyServiceServer interface {
	mustEmbedUnimplementedProxyServiceServer()
}

func RegisterProxyServiceServer(s grpc.ServiceRegistrar, srv ProxyServiceServer) {
	s.RegisterService(&ProxyService_ServiceDesc, srv)
}

func _ProxyService_GetTelemetry_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTelemetryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServiceServer).GetTelemetry(m, &proxyServiceGetTelemetryServer{stream})
}

type ProxyService_GetTelemetryServer interface {
	Send(*GetTelemetryResponse) error
	grpc.ServerStream
}

type proxyServiceGetTelemetryServer struct {
	grpc.ServerStream
}

func (x *proxyServiceGetTelemetryServer) Send(m *GetTelemetryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProxyService_GetEnemyTelemetry_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetEnemyTelemetryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServiceServer).GetEnemyTelemetry(m, &proxyServiceGetEnemyTelemetryServer{stream})
}

type ProxyService_GetEnemyTelemetryServer interface {
	Send(*GetEnemyTelemetryResponse) error
	grpc.ServerStream
}

type proxyServiceGetEnemyTelemetryServer struct {
	grpc.ServerStream
}

func (x *proxyServiceGetEnemyTelemetryServer) Send(m *GetEnemyTelemetryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProxyService_SendLockPacket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendLockPacketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServiceServer).SendLockPacket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProxyService_SendLockPacket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServiceServer).SendLockPacket(ctx, req.(*SendLockPacketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProxyService_SendKamikazePacket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendKamikazePacketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServiceServer).SendKamikazePacket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProxyService_SendKamikazePacket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServiceServer).SendKamikazePacket(ctx, req.(*SendKamikazePacketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProxyService_ServiceDesc is the grpc.ServiceDesc for ProxyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProxyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proxy.v1.ProxyService",
	HandlerType: (*ProxyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendLockPacket",
			Handler:    _ProxyService_SendLockPacket_Handler,
		},
		{
			MethodName: "SendKamikazePacket",
			Handler:    _ProxyService_SendKamikazePacket_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTelemetry",
			Handler:       _ProxyService_GetTelemetry_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetEnemyTelemetry",
			Handler:       _ProxyService_GetEnemyTelemetry_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/proxy/v1/proxy.proto",
}
