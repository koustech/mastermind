// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProxyService_GetTelemetry_FullMethodName       = "/proxy.v1.ProxyService/GetTelemetry"
	ProxyService_GetEnemyTelemetry_FullMethodName  = "/proxy.v1.ProxyService/GetEnemyTelemetry"
	ProxyService_SendLockPacket_FullMethodName     = "/proxy.v1.ProxyService/SendLockPacket"
	ProxyService_SendKamikazePacket_FullMethodName = "/proxy.v1.ProxyService/SendKamikazePacket"
	ProxyService_GetRedZonePacket_FullMethodName   = "/proxy.v1.ProxyService/GetRedZonePacket"
)

// ProxyServiceClient is the client API for ProxyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProxyServiceClient interface {
	// Returns a normal telemetry stream when activated. Should be used by
	// services that don't interact with the competition server
	GetTelemetry(ctx context.Context, in *GetTelemetryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetTelemetryResponse], error)
	// GetEnemyTelemetry returns the enemy's telemetry stream when
	// activated
	GetEnemyTelemetry(ctx context.Context, in *GetEnemyTelemetryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetEnemyTelemetryResponse], error)
	// SendLockPacket sends a successful lock notification to the competition server
	SendLockPacket(ctx context.Context, in *SendLockPacketRequest, opts ...grpc.CallOption) (*SendLockPacketResponse, error)
	// SendKamikazePacket sends a successful kamikaze QR code reading to the competition server
	SendKamikazePacket(ctx context.Context, in *SendKamikazePacketRequest, opts ...grpc.CallOption) (*SendKamikazePacketResponse, error)
	GetRedZonePacket(ctx context.Context, in *GetRedZonePacketRequest, opts ...grpc.CallOption) (*GetRedZonePacketResponse, error)
}

type proxyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyServiceClient(cc grpc.ClientConnInterface) ProxyServiceClient {
	return &proxyServiceClient{cc}
}

func (c *proxyServiceClient) GetTelemetry(ctx context.Context, in *GetTelemetryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetTelemetryResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ProxyService_ServiceDesc.Streams[0], ProxyService_GetTelemetry_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetTelemetryRequest, GetTelemetryResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProxyService_GetTelemetryClient = grpc.ServerStreamingClient[GetTelemetryResponse]

func (c *proxyServiceClient) GetEnemyTelemetry(ctx context.Context, in *GetEnemyTelemetryRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetEnemyTelemetryResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ProxyService_ServiceDesc.Streams[1], ProxyService_GetEnemyTelemetry_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetEnemyTelemetryRequest, GetEnemyTelemetryResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProxyService_GetEnemyTelemetryClient = grpc.ServerStreamingClient[GetEnemyTelemetryResponse]

func (c *proxyServiceClient) SendLockPacket(ctx context.Context, in *SendLockPacketRequest, opts ...grpc.CallOption) (*SendLockPacketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendLockPacketResponse)
	err := c.cc.Invoke(ctx, ProxyService_SendLockPacket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyServiceClient) SendKamikazePacket(ctx context.Context, in *SendKamikazePacketRequest, opts ...grpc.CallOption) (*SendKamikazePacketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendKamikazePacketResponse)
	err := c.cc.Invoke(ctx, ProxyService_SendKamikazePacket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *proxyServiceClient) GetRedZonePacket(ctx context.Context, in *GetRedZonePacketRequest, opts ...grpc.CallOption) (*GetRedZonePacketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetRedZonePacketResponse)
	err := c.cc.Invoke(ctx, ProxyService_GetRedZonePacket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyServiceServer is the server API for ProxyService service.
// All implementations should embed UnimplementedProxyServiceServer
// for forward compatibility.
type ProxyServiceServer interface {
	// Returns a normal telemetry stream when activated. Should be used by
	// services that don't interact with the competition server
	GetTelemetry(*GetTelemetryRequest, grpc.ServerStreamingServer[GetTelemetryResponse]) error
	// GetEnemyTelemetry returns the enemy's telemetry stream when
	// activated
	GetEnemyTelemetry(*GetEnemyTelemetryRequest, grpc.ServerStreamingServer[GetEnemyTelemetryResponse]) error
	// SendLockPacket sends a successful lock notification to the competition server
	SendLockPacket(context.Context, *SendLockPacketRequest) (*SendLockPacketResponse, error)
	// SendKamikazePacket sends a successful kamikaze QR code reading to the competition server
	SendKamikazePacket(context.Context, *SendKamikazePacketRequest) (*SendKamikazePacketResponse, error)
	GetRedZonePacket(context.Context, *GetRedZonePacketRequest) (*GetRedZonePacketResponse, error)
}

// UnimplementedProxyServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProxyServiceServer struct{}

func (UnimplementedProxyServiceServer) GetTelemetry(*GetTelemetryRequest, grpc.ServerStreamingServer[GetTelemetryResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetTelemetry not implemented")
}
func (UnimplementedProxyServiceServer) GetEnemyTelemetry(*GetEnemyTelemetryRequest, grpc.ServerStreamingServer[GetEnemyTelemetryResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetEnemyTelemetry not implemented")
}
func (UnimplementedProxyServiceServer) SendLockPacket(context.Context, *SendLockPacketRequest) (*SendLockPacketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendLockPacket not implemented")
}
func (UnimplementedProxyServiceServer) SendKamikazePacket(context.Context, *SendKamikazePacketRequest) (*SendKamikazePacketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendKamikazePacket not implemented")
}
func (UnimplementedProxyServiceServer) GetRedZonePacket(context.Context, *GetRedZonePacketRequest) (*GetRedZonePacketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRedZonePacket not implemented")
}
func (UnimplementedProxyServiceServer) testEmbeddedByValue() {}

// UnsafeProxyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProxyServiceServer will
// result in compilation errors.
type UnsafeProxyServiceServer interface {
	mustEmbedUnimplementedProxyServiceServer()
}

func RegisterProxyServiceServer(s grpc.ServiceRegistrar, srv ProxyServiceServer) {
	// If the following call pancis, it indicates UnimplementedProxyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProxyService_ServiceDesc, srv)
}

func _ProxyService_GetTelemetry_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTelemetryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServiceServer).GetTelemetry(m, &grpc.GenericServerStream[GetTelemetryRequest, GetTelemetryResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProxyService_GetTelemetryServer = grpc.ServerStreamingServer[GetTelemetryResponse]

func _ProxyService_GetEnemyTelemetry_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetEnemyTelemetryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProxyServiceServer).GetEnemyTelemetry(m, &grpc.GenericServerStream[GetEnemyTelemetryRequest, GetEnemyTelemetryResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ProxyService_GetEnemyTelemetryServer = grpc.ServerStreamingServer[GetEnemyTelemetryResponse]

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

func _ProxyService_GetRedZonePacket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRedZonePacketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyServiceServer).GetRedZonePacket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProxyService_GetRedZonePacket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyServiceServer).GetRedZonePacket(ctx, req.(*GetRedZonePacketRequest))
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
		{
			MethodName: "GetRedZonePacket",
			Handler:    _ProxyService_GetRedZonePacket_Handler,
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
