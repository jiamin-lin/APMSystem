package dogalarm

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GrpcClient struct {
	*grpc.ClientConn
}

func NewGroupClient(add string) *GrpcClient {
	conn, err := grpc.Dial(add,
		grpc.WithUnaryInterceptor(unaryInterceptor()),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	return &GrpcClient{conn}
}

func unaryInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}
