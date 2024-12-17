package main

import (
	"dogapm"
	"protos"
	"skusvc/grpc"
)

func main() {
	// 初始化db，grpc server
	dogapm.Infra.Init(
		dogapm.InfraDbOption("root:1234567@tcp(134.175.127.240:3308)/skusvc"),
	)
	grpcserver := dogapm.NewGrpcServer(":8081")
	protos.RegisterSkuServiceServer(grpcserver, &grpc.SkuServer{})
	dogapm.EndPoint.Start()
}
