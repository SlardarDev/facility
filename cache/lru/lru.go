package lru

import (
	"container/list"
	"runtime"
	"sync"
	"time"
)

/*
简单的lru，带逐出功能，实现为map加链表，复杂度为o(1)
*/

type lockedLRU struct {
	c           simpleLRU
	stopSignal  chan struct{}
	m           sync.Mutex
	name        string
	openStat    bool
	emitCounter func(name string, value interface{}, prefix string, tagkv map[string]string)
}

func (l *lockedLRU) SetEvictCallBack(cb func(k, v interface{})) {
	l.c.callback = cb
}

// NewLocked constructs a new Cache of the given size that is safe for
// concurrent use. It will panic if size is not a positive integer.
func NewLockedLRU(size int, name string, openStat bool,
	emitCounter func(name string, value interface{}, prefix string, tagkv map[string]string)) *lockedLRU {
	if size <= 0 {
		panic("inmem: must provide a positive size")
	}
	lru := &lockedLRU{
		openStat:   openStat,
		stopSignal: make(chan struct{}),
		name:       name,
		c: simpleLRU{
			size:  size,
			lru:   list.New(),
			items: make(map[interface{}]*list.Element),
		},
		emitCounter: emitCounter,
	}

	if openStat && emitCounter != nil {
		go runStat(lru)
		runtime.SetFinalizer(lru, stopStat)

	}
	return lru
}

func stopStat(lru *lockedLRU) {
	lru.stopSignal <- struct{}{}
}

func runStat(lru *lockedLRU) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-lru.stopSignal:
			ticker.Stop()
			return
		case <-ticker.C:
			lru.emitCounter(lru.name+".size", lru.Len(), "", nil)
		}
	}
}

func (l *lockedLRU) Get(key interface{}) (interface{}, bool) {
	l.m.Lock()
	v, ok := l.c.Get(key)
	l.m.Unlock()
	if l.openStat && l.emitCounter != nil {
		if ok {
			l.emitCounter(l.name+".hit", 1, "", nil)
		} else {
			l.emitCounter(l.name+".miss", 1, "", nil)
		}
	}
	return v, ok
}

func (l *lockedLRU) Remove(key interface{}) {
	l.m.Lock()
	l.c.Remove(key)
	l.m.Unlock()
}

func (l *lockedLRU) Len() int {
	l.m.Lock()
	c := l.c.Len()
	l.m.Unlock()
	return c
}

func (l *lockedLRU) Name() string {
	return l.name
}

func (l *lockedLRU) Set(k, v interface{}, ttl time.Duration) {
	l.m.Lock()
	l.c.Add(k, v, time.Now().Add(ttl))
	l.m.Unlock()
}
