package dogapm

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcClient grpc客户端
type GrpcClient struct {
	*grpc.ClientConn
}

// NewGrpcClient 创建GrpcClient
func NewGrpcClient(addr string) *GrpcClient {
	conn, err := grpc.Dial(addr,
		grpc.WithUnaryInterceptor(unaryInterceptor()),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &GrpcClient{ClientConn: conn}
}

// 拦截器
func unaryInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}
