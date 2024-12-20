// proto的注释使用C风格的//和/*...*/，同Go.

// 协议版本，必须是proto文件中的首行(注释和空行不算)，如不写默认使用proto2版本
// 由于proto3中删减了一些字段修饰词，因此看起来更简洁些，本文不讲proto2

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: sku.proto

package protos

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
	SkuService_DecreaseStock_FullMethodName = "/SkuService/decreaseStock"
)

// SkuServiceClient is the client API for SkuService liveness.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SkuServiceClient interface {
	DecreaseStock(ctx context.Context, in *Sku, opts ...grpc.CallOption) (*Sku, error)
}

type skuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSkuServiceClient(cc grpc.ClientConnInterface) SkuServiceClient {
	return &skuServiceClient{cc}
}

func (c *skuServiceClient) DecreaseStock(ctx context.Context, in *Sku, opts ...grpc.CallOption) (*Sku, error) {
	out := new(Sku)
	err := c.cc.Invoke(ctx, SkuService_DecreaseStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SkuServiceServer is the server API for SkuService liveness.
// All implementations must embed UnimplementedSkuServiceServer
// for forward compatibility
type SkuServiceServer interface {
	DecreaseStock(context.Context, *Sku) (*Sku, error)
	mustEmbedUnimplementedSkuServiceServer()
}

// UnimplementedSkuServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSkuServiceServer struct {
}

func (UnimplementedSkuServiceServer) DecreaseStock(context.Context, *Sku) (*Sku, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecreaseStock not implemented")
}
func (UnimplementedSkuServiceServer) mustEmbedUnimplementedSkuServiceServer() {}

// UnsafeSkuServiceServer may be embedded to opt out of forward compatibility for this liveness.
// Use of this interface is not recommended, as added methods to SkuServiceServer will
// result in compilation errors.
type UnsafeSkuServiceServer interface {
	mustEmbedUnimplementedSkuServiceServer()
}

func RegisterSkuServiceServer(s grpc.ServiceRegistrar, srv SkuServiceServer) {
	s.RegisterService(&SkuService_ServiceDesc, srv)
}

func _SkuService_DecreaseStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Sku)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SkuServiceServer).DecreaseStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SkuService_DecreaseStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SkuServiceServer).DecreaseStock(ctx, req.(*Sku))
	}
	return interceptor(ctx, in, info, handler)
}

// SkuService_ServiceDesc is the grpc.ServiceDesc for SkuService liveness.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SkuService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "SkuService",
	HandlerType: (*SkuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "decreaseStock",
			Handler:    _SkuService_DecreaseStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sku.proto",
}
