package parallel_loader

import (
	"bytes"
	"context"
	"runtime"
	"sync"

	"github.com/SlardarDev/facility/logs"
)

var stackBuffer = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 4096))
	},
}

func printStackTrace(ctx context.Context) {

	buffer := stackBuffer.Get().(*bytes.Buffer)
	buffer.Reset()
	buf := buffer.Bytes()
	buf = buf[:4096]
	for i := range buf {
		buf[i] = 0
	}
	defer func() {
		buffer.Reset()
		stackBuffer.Put(buffer)
	}()
	buf = buf[:runtime.Stack(buf, false)] //获得当前goroutine的stacktrace
	logs.CtxError(ctx, "%s", string(buf))
}
