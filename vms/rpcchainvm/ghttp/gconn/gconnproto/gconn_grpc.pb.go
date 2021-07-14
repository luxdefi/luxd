// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gconnproto

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

// ConnClient is the client API for Conn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error)
	SetDeadline(ctx context.Context, in *SetDeadlineRequest, opts ...grpc.CallOption) (*SetDeadlineResponse, error)
	SetReadDeadline(ctx context.Context, in *SetReadDeadlineRequest, opts ...grpc.CallOption) (*SetReadDeadlineResponse, error)
	SetWriteDeadline(ctx context.Context, in *SetWriteDeadlineRequest, opts ...grpc.CallOption) (*SetWriteDeadlineResponse, error)
}

type connClient struct {
	cc grpc.ClientConnInterface
}

func NewConnClient(cc grpc.ClientConnInterface) ConnClient {
	return &connClient{cc}
}

func (c *connClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/gconnproto.Conn/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/gconnproto.Conn/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connClient) Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error) {
	out := new(CloseResponse)
	err := c.cc.Invoke(ctx, "/gconnproto.Conn/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connClient) SetDeadline(ctx context.Context, in *SetDeadlineRequest, opts ...grpc.CallOption) (*SetDeadlineResponse, error) {
	out := new(SetDeadlineResponse)
	err := c.cc.Invoke(ctx, "/gconnproto.Conn/SetDeadline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connClient) SetReadDeadline(ctx context.Context, in *SetReadDeadlineRequest, opts ...grpc.CallOption) (*SetReadDeadlineResponse, error) {
	out := new(SetReadDeadlineResponse)
	err := c.cc.Invoke(ctx, "/gconnproto.Conn/SetReadDeadline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connClient) SetWriteDeadline(ctx context.Context, in *SetWriteDeadlineRequest, opts ...grpc.CallOption) (*SetWriteDeadlineResponse, error) {
	out := new(SetWriteDeadlineResponse)
	err := c.cc.Invoke(ctx, "/gconnproto.Conn/SetWriteDeadline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnServer is the server API for Conn service.
// All implementations must embed UnimplementedConnServer
// for forward compatibility
type ConnServer interface {
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	Close(context.Context, *CloseRequest) (*CloseResponse, error)
	SetDeadline(context.Context, *SetDeadlineRequest) (*SetDeadlineResponse, error)
	SetReadDeadline(context.Context, *SetReadDeadlineRequest) (*SetReadDeadlineResponse, error)
	SetWriteDeadline(context.Context, *SetWriteDeadlineRequest) (*SetWriteDeadlineResponse, error)
	mustEmbedUnimplementedConnServer()
}

// UnimplementedConnServer must be embedded to have forward compatible implementations.
type UnimplementedConnServer struct{}

func (UnimplementedConnServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}

func (UnimplementedConnServer) Write(context.Context, *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}

func (UnimplementedConnServer) Close(context.Context, *CloseRequest) (*CloseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}

func (UnimplementedConnServer) SetDeadline(context.Context, *SetDeadlineRequest) (*SetDeadlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeadline not implemented")
}

func (UnimplementedConnServer) SetReadDeadline(context.Context, *SetReadDeadlineRequest) (*SetReadDeadlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetReadDeadline not implemented")
}

func (UnimplementedConnServer) SetWriteDeadline(context.Context, *SetWriteDeadlineRequest) (*SetWriteDeadlineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetWriteDeadline not implemented")
}
func (UnimplementedConnServer) mustEmbedUnimplementedConnServer() {}

// UnsafeConnServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnServer will
// result in compilation errors.
type UnsafeConnServer interface {
	mustEmbedUnimplementedConnServer()
}

func RegisterConnServer(s grpc.ServiceRegistrar, srv ConnServer) {
	s.RegisterService(&Conn_ServiceDesc, srv)
}

func _Conn_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gconnproto.Conn/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conn_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gconnproto.Conn/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conn_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gconnproto.Conn/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).Close(ctx, req.(*CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conn_SetDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).SetDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gconnproto.Conn/SetDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).SetDeadline(ctx, req.(*SetDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conn_SetReadDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetReadDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).SetReadDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gconnproto.Conn/SetReadDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).SetReadDeadline(ctx, req.(*SetReadDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Conn_SetWriteDeadline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetWriteDeadlineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnServer).SetWriteDeadline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gconnproto.Conn/SetWriteDeadline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnServer).SetWriteDeadline(ctx, req.(*SetWriteDeadlineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Conn_ServiceDesc is the grpc.ServiceDesc for Conn service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Conn_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gconnproto.Conn",
	HandlerType: (*ConnServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _Conn_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _Conn_Write_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _Conn_Close_Handler,
		},
		{
			MethodName: "SetDeadline",
			Handler:    _Conn_SetDeadline_Handler,
		},
		{
			MethodName: "SetReadDeadline",
			Handler:    _Conn_SetReadDeadline_Handler,
		},
		{
			MethodName: "SetWriteDeadline",
			Handler:    _Conn_SetWriteDeadline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gconn.proto",
}
