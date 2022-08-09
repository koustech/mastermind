// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
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
}

type proxyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyServiceClient(cc grpc.ClientConnInterface) ProxyServiceClient {
	return &proxyServiceClient{cc}
}

func (c *proxyServiceClient) GetTelemetry(ctx context.Context, in *GetTelemetryRequest, opts ...grpc.CallOption) (ProxyService_GetTelemetryClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProxyService_ServiceDesc.Streams[0], "/proxy.v1.ProxyService/GetTelemetry", opts...)
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
	stream, err := c.cc.NewStream(ctx, &ProxyService_ServiceDesc.Streams[1], "/proxy.v1.ProxyService/GetEnemyTelemetry", opts...)
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

// ProxyService_ServiceDesc is the grpc.ServiceDesc for ProxyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProxyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proxy.v1.ProxyService",
	HandlerType: (*ProxyServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
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