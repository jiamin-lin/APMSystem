package dogapm

import (
	"os"
	"os/signal"
	"syscall"
)

type starter interface {
	Start()
}

type closer interface {
	Close()
}

var (
	globalStarters = make([]starter, 0)
	globalClosers  = make([]closer, 0)
)

type endPoint struct {
	stop chan int
}

var EndPoint = &endPoint{stop: make(chan int)}

func (e *endPoint) Start() {
	for _, com := range globalStarters {
		com.Start()
	}
	go func() {
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-quit
		e.ShutDown()
	}()
	<-e.stop
}

func (e *endPoint) ShutDown() {
	for _, com := range globalClosers {
		com.Close()
	}
	e.stop <- 1
}
