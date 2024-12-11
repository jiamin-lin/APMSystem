package dogapm

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

// GrpcServer grpc服务端
type GrpcServer struct {
	*grpc.Server
	addr string
}

// NewGrpcServer 创建GrpcServer
func NewGrpcServer(addr string) *GrpcServer {
	svc := grpc.NewServer(grpc.UnaryInterceptor(unaryServerInterceptor()))
	server := &GrpcServer{
		Server: svc,
		addr:   addr,
	}
	// 服务注册
	globalClosers = append(globalClosers, server)
	globalStarters = append(globalStarters, server)
	return server
}

func (g *GrpcServer) Close() {
	g.Server.GracefulStop()
}

// Start 启动服务
func (g *GrpcServer) Start() {
	l, err := net.Listen("tcp", g.addr)
	if err != nil {
	}
	go func() {
		err = g.Serve(l)
		if err != nil {
			panic(err)
		}
	}()
}

// 拦截器
func unaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(ctx, req)
	}
}
