package dogapm

import (
	"context"
	"fmt"
	"net/http"
	"protoc"
	"testing"
	"time"
)

func TestInfra_InitInfra(t *testing.T) {
	Infra.InitInfra(InfraDbOption("root:1234567@tcp(134.175.127.240:3308)/ordersvc"),
		InfraRDBOption("134.175.127.240:6380"))
}

func TestNewHttpServer(t *testing.T) {
	s := NewHttpServer(":8080")
	s.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	s.Start()
	time.Sleep(time.Hour)
}

// helloSvc hello
type helloSvc struct {
	protoc.UnimplementedHelloServiceServer
}

// Receive 接收消息
func (h *helloSvc) Receive(ctx context.Context, msg *protoc.HelloMsg) (*protoc.HelloMsg, error) {
	return msg, nil
}

// TestGrpc 测试grpc
func TestGrpc(t *testing.T) {
	go func() {
		s := NewGrpcServer(":8080")
		protoc.RegisterHelloServiceServer(s, &helloSvc{})
		s.Start()
	}()
	client := NewGrpcClient("127.0.0.1:8080")
	res, err := protoc.NewHelloServiceClient(client).
		Receive(context.TODO(), &protoc.HelloMsg{Msg: "hello"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.Msg)
}
