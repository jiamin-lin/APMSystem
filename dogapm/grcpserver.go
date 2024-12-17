package dogapm

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	*grpc.Server
	addr string
}

func NewGrpcServer(addr string) *GrpcServer {
	svc := grpc.NewServer(grpc.UnaryInterceptor(unaryServerInterceptor()))
	s := &GrpcServer{Server: svc, addr: addr}
	globalStarters = append(globalStarters, s)
	globalClosers = append(globalClosers, s)
	return s
}

func unaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		return handler(ctx, req)
	}
}

func (g *GrpcServer) Start() {
	l, err := net.Listen("tcp", g.addr)
	if err != nil {
		panic(err)
	}
	go func() {
		err = g.Serve(l)
		if err != nil {
			panic(err)
		}
	}()
}

func (g *GrpcServer) Close() {
	g.Server.GracefulStop()
}
