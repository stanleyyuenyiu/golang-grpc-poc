// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: protos/simple/simple.proto

package simple

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

// CallClient is the client API for Call service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CallClient interface {
	ReqRes(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error)
}

type callClient struct {
	cc grpc.ClientConnInterface
}

func NewCallClient(cc grpc.ClientConnInterface) CallClient {
	return &callClient{cc}
}

func (c *callClient) ReqRes(ctx context.Context, in *CallRequest, opts ...grpc.CallOption) (*CallResponse, error) {
	out := new(CallResponse)
	err := c.cc.Invoke(ctx, "/simple.Call/ReqRes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CallServer is the server API for Call service.
// All implementations must embed UnimplementedCallServer
// for forward compatibility
type CallServer interface {
	ReqRes(context.Context, *CallRequest) (*CallResponse, error)
	mustEmbedUnimplementedCallServer()
}

// UnimplementedCallServer must be embedded to have forward compatible implementations.
type UnimplementedCallServer struct {
}

func (UnimplementedCallServer) ReqRes(context.Context, *CallRequest) (*CallResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReqRes not implemented")
}
func (UnimplementedCallServer) mustEmbedUnimplementedCallServer() {}

// UnsafeCallServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CallServer will
// result in compilation errors.
type UnsafeCallServer interface {
	mustEmbedUnimplementedCallServer()
}

func RegisterCallServer(s grpc.ServiceRegistrar, srv CallServer) {
	s.RegisterService(&Call_ServiceDesc, srv)
}

func _Call_ReqRes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CallServer).ReqRes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simple.Call/ReqRes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CallServer).ReqRes(ctx, req.(*CallRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Call_ServiceDesc is the grpc.ServiceDesc for Call service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Call_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simple.Call",
	HandlerType: (*CallServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReqRes",
			Handler:    _Call_ReqRes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/simple/simple.proto",
}
