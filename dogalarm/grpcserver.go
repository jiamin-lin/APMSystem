package dogalarm

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
	return &GrpcServer{
		Server: svc,
		addr:   addr,
	}
}

func (g *GrpcServer) Start() {
	l, err := net.Listen("tcp", g.addr)
	if err != nil {
		go func() {
			err := g.Serve(l)
			if err != nil {
				panic(err)
			}
		}()
	}

}

func unaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		return handler(ctx, req)
	}
}
