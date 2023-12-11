// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.20.3
// source: wifi_service.proto

package pb

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
	WifiService_SendCredentials_FullMethodName = "/WifiService/SendCredentials"
	WifiService_GetCredentials_FullMethodName  = "/WifiService/GetCredentials"
)

// WifiServiceClient is the client API for WifiService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WifiServiceClient interface {
	// Sends WiFi credentials
	SendCredentials(ctx context.Context, in *WifiCredentials, opts ...grpc.CallOption) (*Confirmation, error)
	// Receives WiFi credentials from the server
	GetCredentials(ctx context.Context, in *Request, opts ...grpc.CallOption) (*WifiCredentials, error)
}

type wifiServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWifiServiceClient(cc grpc.ClientConnInterface) WifiServiceClient {
	return &wifiServiceClient{cc}
}

func (c *wifiServiceClient) SendCredentials(ctx context.Context, in *WifiCredentials, opts ...grpc.CallOption) (*Confirmation, error) {
	out := new(Confirmation)
	err := c.cc.Invoke(ctx, WifiService_SendCredentials_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *wifiServiceClient) GetCredentials(ctx context.Context, in *Request, opts ...grpc.CallOption) (*WifiCredentials, error) {
	out := new(WifiCredentials)
	err := c.cc.Invoke(ctx, WifiService_GetCredentials_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WifiServiceServer is the server API for WifiService service.
// All implementations must embed UnimplementedWifiServiceServer
// for forward compatibility
type WifiServiceServer interface {
	// Sends WiFi credentials
	SendCredentials(context.Context, *WifiCredentials) (*Confirmation, error)
	// Receives WiFi credentials from the server
	GetCredentials(context.Context, *Request) (*WifiCredentials, error)
	mustEmbedUnimplementedWifiServiceServer()
}

// UnimplementedWifiServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWifiServiceServer struct {
}

func (UnimplementedWifiServiceServer) SendCredentials(context.Context, *WifiCredentials) (*Confirmation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCredentials not implemented")
}
func (UnimplementedWifiServiceServer) GetCredentials(context.Context, *Request) (*WifiCredentials, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCredentials not implemented")
}
func (UnimplementedWifiServiceServer) mustEmbedUnimplementedWifiServiceServer() {}

// UnsafeWifiServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WifiServiceServer will
// result in compilation errors.
type UnsafeWifiServiceServer interface {
	mustEmbedUnimplementedWifiServiceServer()
}

func RegisterWifiServiceServer(s grpc.ServiceRegistrar, srv WifiServiceServer) {
	s.RegisterService(&WifiService_ServiceDesc, srv)
}

func _WifiService_SendCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WifiCredentials)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WifiServiceServer).SendCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WifiService_SendCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WifiServiceServer).SendCredentials(ctx, req.(*WifiCredentials))
	}
	return interceptor(ctx, in, info, handler)
}

func _WifiService_GetCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WifiServiceServer).GetCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WifiService_GetCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WifiServiceServer).GetCredentials(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// WifiService_ServiceDesc is the grpc.ServiceDesc for WifiService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WifiService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WifiService",
	HandlerType: (*WifiServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendCredentials",
			Handler:    _WifiService_SendCredentials_Handler,
		},
		{
			MethodName: "GetCredentials",
			Handler:    _WifiService_GetCredentials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wifi_service.proto",
}