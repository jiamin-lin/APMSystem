package dogapm

import (
	"net/http"
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
