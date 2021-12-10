package parallel_loader

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/SlardarDev/facility/logs"
)

type LoaderContainer []ParallelLoader

func (l LoaderContainer) Load(ctx context.Context, ctxParam interface{}) error {
	var (
		wg   sync.WaitGroup
		lock sync.Mutex
		err  error
	)
	for _, loader := range l {
		wg.Add(1)
		go func(loader ParallelLoader) {
			defer func() {
				wg.Done()
				if r := recover(); r != nil {
					errMsg := fmt.Sprintf("LoaderMgr panic: loader: %v, reason: %v", loader.Name(), r)
					logs.CtxFatal(ctx, "%v", errMsg)
					nerr := errors.New(errMsg)
					lock.Lock()
					if err == nil {
						err = nerr
					}
					lock.Unlock()
					printStackTrace(ctx)
				}
			}()
			logs.Debug("run loader: %v", loader.Name())
			nerr := loader.Load(ctx, ctxParam)
			if nerr != nil {
				lock.Lock()
				if err == nil {
					err = nerr
				}
				lock.Unlock()
			}
		}(loader)
	}
	wg.Wait()
	return err
}

func (l LoaderContainer) Name() string {
	return ""
}
