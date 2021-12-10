package afunc

import (
	"errors"
	"sync/atomic"
)

type Pool struct {
	currentWorkerNumber int64
	metricsFunc         func(int64)
	ch                  chan struct{}
}

func NewPool(workerNumber int64, metricsFunc ...func(currentWorker int64)) (*Pool, error) {
	if workerNumber <= 0 {
		return nil, errors.New("worker number of pool must great than zero")
	}
	var m func(int64) = nil
	if len(metricsFunc) != 0 {
		m = metricsFunc[0]
	}
	p := &Pool{
		currentWorkerNumber: 0,
		metricsFunc:         m,
		ch:                  make(chan struct{}, workerNumber),
	}
	return p, nil
}

func (p *Pool) Do(f func()) {
	p.ch <- struct{}{}
	go func() {
		if p.metricsFunc != nil {
			atomic.AddInt64(&p.currentWorkerNumber, 1)
			p.metricsFunc(atomic.LoadInt64(&p.currentWorkerNumber))
		}
		defer func() {
			if p.metricsFunc != nil {
				atomic.AddInt64(&p.currentWorkerNumber, -1)
				p.metricsFunc(atomic.LoadInt64(&p.currentWorkerNumber))
			}
			<-p.ch
		}()
		f()
	}()
}
