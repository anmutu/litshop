package runtime

type RunTime struct {
	shutdownHandles []func()
	RunApp          string
}

var (
	R = RunTime{}
)

func SetRunApp(s string) {
	R.RunApp = s
}

func RegisterShutdown(f func()) {
	R.shutdownHandles = append(R.shutdownHandles, f)
}

func ShutDown() {
	if len(R.shutdownHandles) > 0 {
		for _, f := range R.shutdownHandles {
			if f != nil {
				f()
			}
		}
	}
}
