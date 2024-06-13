// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: router/Handlers/routes.proto

package handlers

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

// RoutesServiceClient is the client API for RoutesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RoutesServiceClient interface {
	Home(ctx context.Context, in *HomeRequest, opts ...grpc.CallOption) (*HomeResponse, error)
	Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type routesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRoutesServiceClient(cc grpc.ClientConnInterface) RoutesServiceClient {
	return &routesServiceClient{cc}
}

func (c *routesServiceClient) Home(ctx context.Context, in *HomeRequest, opts ...grpc.CallOption) (*HomeResponse, error) {
	out := new(HomeResponse)
	err := c.cc.Invoke(ctx, "/RoutesService/Home", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routesServiceClient) Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, "/RoutesService/Post", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *routesServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/RoutesService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoutesServiceServer is the server API for RoutesService service.
// All implementations must embed UnimplementedRoutesServiceServer
// for forward compatibility
type RoutesServiceServer interface {
	Home(context.Context, *HomeRequest) (*HomeResponse, error)
	Post(context.Context, *PostRequest) (*PostResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedRoutesServiceServer()
}

// UnimplementedRoutesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRoutesServiceServer struct {
}

func (UnimplementedRoutesServiceServer) Home(context.Context, *HomeRequest) (*HomeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Home not implemented")
}
func (UnimplementedRoutesServiceServer) Post(context.Context, *PostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedRoutesServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRoutesServiceServer) mustEmbedUnimplementedRoutesServiceServer() {}

// UnsafeRoutesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RoutesServiceServer will
// result in compilation errors.
type UnsafeRoutesServiceServer interface {
	mustEmbedUnimplementedRoutesServiceServer()
}

func RegisterRoutesServiceServer(s grpc.ServiceRegistrar, srv RoutesServiceServer) {
	s.RegisterService(&RoutesService_ServiceDesc, srv)
}

func _RoutesService_Home_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutesServiceServer).Home(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoutesService/Home",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutesServiceServer).Home(ctx, req.(*HomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutesService_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutesServiceServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoutesService/Post",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutesServiceServer).Post(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RoutesService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoutesServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RoutesService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoutesServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RoutesService_ServiceDesc is the grpc.ServiceDesc for RoutesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RoutesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RoutesService",
	HandlerType: (*RoutesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Home",
			Handler:    _RoutesService_Home_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _RoutesService_Post_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RoutesService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "router/Handlers/routes.proto",
}