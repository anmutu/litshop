package runtime

var shutdownHandles []func()

func init() {
	shutdownHandles = make([]func(), 10)
}

func RegisterShutdown(f func()) {
	shutdownHandles = append(shutdownHandles, f)
}

func ShutDown() {
	if len(shutdownHandles) > 0 {
		for _, f := range shutdownHandles {
			go f()
		}
	}
}
