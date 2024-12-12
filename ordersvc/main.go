package main

import (
	"dogapm"
	"net/http"
	"ordersvc/api"
)

func main() {
	// 初始化 db，httpserver，grpcserver

	dogapm.Infra.InitInfra(
		dogapm.InfraDbOption("root:1234567@tcp(134.175.127.240:3308)/ordersvc"),
	)

	//Todo: 初始化grpc server

	httpserver := dogapm.NewHttpServer(":8080")
	httpserver.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	httpserver.HandleFunc("/order/add", api.Order.Add)

	dogapm.EndPoint.Start()

}
