package dogapm

import (
	"context"
	"fmt"
	"protos"
	"testing"
)

type helloSvc struct {
	protos.UnimplementedHelloServiceServer
}

func (h *helloSvc) Receive(ctx context.Context, msg *protos.HelloMsg) (*protos.HelloMsg, error) {
	return &protos.HelloMsg{Msg: msg.Msg}, nil
}

func TestGrpc(t *testing.T) {
	go func() {
		s := NewGrpcServer(":8080")
		protos.RegisterHelloServiceServer(s, &helloSvc{})
		s.Start()
	}()
	client := NewGrpcClient("127.0.0.1:8080")
	res, err := protos.NewHelloServiceClient(client).
		Receive(context.Background(), &protos.HelloMsg{Msg: "hello wrold"})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res.GetMsg())
}
