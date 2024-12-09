package dogapm

import (
	"context"
	"net/http"
)

// HttpServer 创建结构体
type HttpServer struct {
	mux *http.ServeMux
	*http.Server
}

// NewHttpServer 创建HttpServer
func NewHttpServer(addr string) *HttpServer {
	mux := http.NewServeMux()
	server := &http.Server{Addr: addr, Handler: mux}
	s := &HttpServer{mux: mux, Server: server}
	globalStarters = append(globalStarters, s)
	globalClosers = append(globalClosers, s)
	return s
}

// Handle ListenAndServe 启动服务
func (h *HttpServer) Handle(pattern string, handler http.Handler) {
	h.mux.Handle(pattern, handler)
}

// HandleFunc 注册路由
func (h *HttpServer) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	h.mux.HandleFunc(pattern, handler)
}

// Start 启动服务
func (h *HttpServer) Start() {
	go func() {
		err := h.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
}

// Close 关闭服务
func (h *HttpServer) Close() {
	h.Shutdown(context.TODO())
}
