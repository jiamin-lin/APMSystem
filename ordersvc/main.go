package main

import (
	"dogapm"
	"net/http"
	"ordersvc/grpcclient"
	"protos"
)

func main() {

	// 初始化db， httpserver， grpcclient
	dogapm.Infra.Init(
		dogapm.InfraDbOption("root:1234567@tcp(134.175.127.240:3308)/ordersvc"),
	)

	//  grpcclient 初始化
	grpcclient.SkuClient = protos.NewSkuServiceClient(dogapm.NewGrpcClient(":8081"))
	grpcclient.UserClient = protos.NewUserServiceClient(dogapm.NewGrpcClient(":8082"))

	httpserver := dogapm.NewHttpServer(":8080")
	httpserver.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("OK"))
		return
	})

	dogapm.EndPoint.Start()

}
