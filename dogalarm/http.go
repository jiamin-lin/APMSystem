package dogalarm

import (
	"context"
	"net/http"
)

type HttpServer struct {
	mux *http.ServeMux
	*http.Server
}

func NewHttpServer(addr string) *HttpServer {
	mux := http.NewServeMux()
	server := &http.Server{Addr: addr, Handler: mux}
	return &HttpServer{mux, server}
}

func (h *HttpServer) Handle(pattern string, handler http.Handler) {
	h.mux.Handle(pattern, handler)
}

func (h *HttpServer) HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	h.mux.HandleFunc(pattern, handler)
}

func (h *HttpServer) Start() {
	go func() {
		err := h.Server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
}

func (h *HttpServer) Close() error {
	return h.Shutdown(context.TODO())
}
