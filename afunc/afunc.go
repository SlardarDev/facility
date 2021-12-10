package afunc

import (
	"runtime"

	"github.com/SlardarDev/facility/logs"
)

type runFunc func() error

func DoASync(run runFunc) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				const size = 4096
				buf := make([]byte, size)
				buf = buf[:runtime.Stack(buf, false)]
				logs.Error("ASync routine panic: %v\n %s\n", err, buf)
			}
		}()
		run()
	}()
}

func DoSync(run runFunc) {
	defer func() {
		if err := recover(); err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			logs.Error("Sync routine panic: %v\n %s\n", err, buf)
		}
	}()
	run()
}
