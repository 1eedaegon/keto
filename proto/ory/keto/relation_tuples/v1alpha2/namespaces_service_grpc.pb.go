// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: ory/keto/relation_tuples/v1alpha2/namespaces_service.proto

package rts

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
	NamespacesService_ListNamespaces_FullMethodName = "/ory.keto.relation_tuples.v1alpha2.NamespacesService/ListNamespaces"
)

// NamespacesServiceClient is the client API for NamespacesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NamespacesServiceClient interface {
	// Lists Namespaces
	ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error)
}

type namespacesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNamespacesServiceClient(cc grpc.ClientConnInterface) NamespacesServiceClient {
	return &namespacesServiceClient{cc}
}

func (c *namespacesServiceClient) ListNamespaces(ctx context.Context, in *ListNamespacesRequest, opts ...grpc.CallOption) (*ListNamespacesResponse, error) {
	out := new(ListNamespacesResponse)
	err := c.cc.Invoke(ctx, NamespacesService_ListNamespaces_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NamespacesServiceServer is the server API for NamespacesService service.
// All implementations should embed UnimplementedNamespacesServiceServer
// for forward compatibility
type NamespacesServiceServer interface {
	// Lists Namespaces
	ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error)
}

// UnimplementedNamespacesServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNamespacesServiceServer struct {
}

func (UnimplementedNamespacesServiceServer) ListNamespaces(context.Context, *ListNamespacesRequest) (*ListNamespacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNamespaces not implemented")
}

// UnsafeNamespacesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NamespacesServiceServer will
// result in compilation errors.
type UnsafeNamespacesServiceServer interface {
	mustEmbedUnimplementedNamespacesServiceServer()
}

func RegisterNamespacesServiceServer(s grpc.ServiceRegistrar, srv NamespacesServiceServer) {
	s.RegisterService(&NamespacesService_ServiceDesc, srv)
}

func _NamespacesService_ListNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNamespacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespacesServiceServer).ListNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NamespacesService_ListNamespaces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespacesServiceServer).ListNamespaces(ctx, req.(*ListNamespacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NamespacesService_ServiceDesc is the grpc.ServiceDesc for NamespacesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NamespacesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ory.keto.relation_tuples.v1alpha2.NamespacesService",
	HandlerType: (*NamespacesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNamespaces",
			Handler:    _NamespacesService_ListNamespaces_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ory/keto/relation_tuples/v1alpha2/namespaces_service.proto",
}
