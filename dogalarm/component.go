package dogalarm

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

func (e *endPoint) Start() {
	for _, com := range globalStarters {
		com.Start()
	}
	go func() {
		// todo 监听服务的结束
	}()

	func (e *endPoint)ShutDown() {
		for _, com := range globalClosers {

		}
	}
}
