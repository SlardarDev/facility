package merge

import (
	"runtime"
	"sync"
	"time"

	"github.com/SlardarDev/facility/afunc"
)

type DataMerger interface {
	Send(data interface{})
	Flush()
}

type dataMerge struct {
	sync.Mutex
	name            string
	mergeSize       int
	flushWaitSecond int
	collectInterval time.Duration
	queue           []interface{}
	queueChan       chan []interface{}
	callBack        func(data []interface{})
	stopSignal      chan struct{}
	ticker          *time.Ticker
}

func (d *dataMerge) Send(data interface{}) {
	d.Lock()
	d.queue = append(d.queue, data)
	d.Unlock()
}

func (d *dataMerge) Flush() {
	for i := 0; i < d.flushWaitSecond; i++ {
		d.collect(false)
		time.Sleep(time.Second)
	}
}

func (d *dataMerge) collect(checkSize bool) bool {
	q := d.flushQueue(checkSize)
	if len(q) == 0 {
		return false
	}
	d.queueChan <- q
	return true
}

func (d *dataMerge) flushQueue(checkSize bool) []interface{} {
	d.Lock()
	defer d.Unlock()
	if checkSize && len(d.queue) < d.mergeSize {
		return nil
	}
	q := d.queue
	d.queue = make([]interface{}, 0, len(q))
	return q
}

func (d *dataMerge) worker() {
	for {
		dataQueue, ok := <-d.queueChan
		if !ok {
			return
		}
		d.callBack(dataQueue)
	}
}

func (d *dataMerge) loop() {
	idle := time.Now()
	for {
		select {
		case <-d.ticker.C:
			if d.collect(time.Since(idle) <= d.collectInterval) {
				idle = time.Now()
			}
		case <-d.stopSignal:
			d.ticker.Stop()
			close(d.queueChan)
			return
		}
	}
}

func stopMerger(d *dataMerge) {
	d.stopSignal <- struct{}{}
}

func NewDataMerger(name string, workerNum, mergeSize, flushWaitSecond int,
	collectInterval time.Duration, callBack func(data []interface{})) DataMerger {

	m := &dataMerge{
		name:            name,
		mergeSize:       mergeSize,
		flushWaitSecond: flushWaitSecond,
		collectInterval: collectInterval,
		callBack:        callBack,
		queue:           make([]interface{}, 0, 100),
		queueChan:       make(chan []interface{}, workerNum*10),
		stopSignal:      make(chan struct{}),
		ticker:          time.NewTicker(time.Millisecond),
	}

	for idx := 0; idx < workerNum; idx++ {
		afunc.DoASync(func() error {
			m.worker()
			return nil
		})
	}
	afunc.DoASync(func() error {
		m.loop()
		return nil
	})

	runtime.SetFinalizer(m, stopMerger)
	return m
}
