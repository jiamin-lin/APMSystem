package skusvc

import (
	"dogapm"
	"protoc"
	"skusvc/grpc"
)

func main() {
	// 初始化 db，httpserver，grpcserver

	dogapm.Infra.InitInfra(
		dogapm.InfraDbOption("root:1234567@tcp(134.175.127.240:3308)/skusvc"))

	grpcserver := dogapm.NewGrpcServer(":8081")
	protoc.RegisterHelloServiceServer(grpcserver, &grpc.SkuServer{})

	dogapm.EndPoint.Start()

}
