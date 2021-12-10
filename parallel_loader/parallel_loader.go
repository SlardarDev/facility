package parallel_loader

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/SlardarDev/facility/logs"

	"errors"
)

type ParallelLoader interface {
	Load(ctx context.Context, ctxParam interface{}) error
	Name() string
}

var ErrorBreak = errors.New("break")

type LoadManager []LoaderContainer

func (mgr LoadManager) Load(ctx context.Context, ctxParam interface{}) error {
	var stop int32 = 0
	var err error
	var lock sync.Mutex
	for _, loaders := range mgr {
		if atomic.LoadInt32(&stop) > 0 {
			break
		}
		var wg sync.WaitGroup
		for _, loader := range loaders {
			if atomic.LoadInt32(&stop) > 0 {
				break
			}
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
						atomic.AddInt32(&stop, 1)
					}
				}()
				logs.Debug("run loader: %v", loader.Name())
				nerr := loader.Load(ctx, ctxParam)
				if nerr != nil {
					atomic.AddInt32(&stop, 1)
					lock.Lock()
					if err == nil {
						err = nerr
					}
					lock.Unlock()
				}
			}(loader)
		}
		wg.Wait()
	}
	if err != nil && err.Error() == ErrorBreak.Error() {
		return nil
	}
	return err
}
