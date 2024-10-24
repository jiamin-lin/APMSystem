package dogalarm

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"protos"
	"testing"
	"time"
)

func TestInit(t *testing.T) {
	Infra.Init(InfraDbOption("root:1234567@tcp(127.0.0.1:3307)/ordersvc"),
		InfraRdbOption("127.0.0.1:6379"))

}

func TestNewHttpServer(t *testing.T) {
	s := NewHttpServer(":8080")
	s.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})
	s.Start()
	time.Sleep(time.Hour)
}

type helloSvc struct {
	protos.UnimplementedHelloServiceServer
}

// Receive 实现 gRPC 服务接口
func (h *helloSvc) Receive(ctx context.Context, msg *protos.HelloMsg) (*protos.HelloMsg, error) {
	return msg, nil
}

func StartGrpcServer(addr string, ready chan<- struct{}) {
	// 创建监听器
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 创建 gRPC 服务器
	s := grpc.NewServer()
	protos.RegisterHelloServiceServer(s, &helloSvc{})

	// 服务器启动通知
	close(ready)
	log.Println("Starting gRPC server on", addr)

	// 启动服务器
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func TestGrpc(t *testing.T) {
	ready := make(chan struct{}) // 用于同步服务器启动状态

	// 启动 gRPC 服务器
	go StartGrpcServer(":8080", ready)

	// 等待服务器启动完成
	<-ready
	time.Sleep(2 * time.Second) // 确保服务完全准备就绪

	// 创建 gRPC 客户端
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(10*time.Second))
	if err != nil {
		t.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := protos.NewHelloServiceClient(conn)

	// 调用服务的 Receive 方法
	res, err := client.Receive(context.TODO(), &protos.HelloMsg{Msg: "hello world"})
	if err != nil {
		t.Fatalf("Error calling service: %v", err)
	}
	fmt.Println("Received:", res.Msg)
}
