package main

import (
	"dogapm"
	pb "protos"
	"usrsvc/grpc"
)

const (
	appName = "usersvc"
)

func main() {
	// 初始化db和redis连接,grpcclient
	dogapm.Infra.Init(
		dogapm.InfraDbOption("root:1234567@tcp(134.175.127.240:3308)/usersvc"),
		dogapm.InfraRDBOption("134.175.127.240:6380"),
	)
	// 启动grpc 服务
	grpcserver := dogapm.NewGrpcServer(":8082")
	pb.RegisterUserServiceServer(grpcserver, &grpc.UserServer{})
	dogapm.EndPoint.Start()
}
