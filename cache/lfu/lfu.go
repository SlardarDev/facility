package lfu

import (
	"container/list"
	"runtime"
	"sync"
	"time"
)

/*
简单的O(1) LFU实现，http://dhruvbird.com/lfu.pdf，改自：https://github.com/dgrijalva/lfu-go
*/

type Eviction struct {
	Key   interface{}
	Value interface{}
}

type simpleLFU struct {
	// If len > upperBound, cache will automatically evict
	// down to lowerBound.  If either value is 0, this behavior
	// is disabled.
	name        string
	openStat    bool
	emitCounter func(name string, value interface{}, prefix string, tagkv map[string]string)
	upperBound  int
	lowerBound  int
	values      map[interface{}]*cacheEntry
	freqs       *list.List
	len         int
	lock        sync.Mutex
	callback    func(k, v interface{})
	stopSignal  chan struct{}
}

func (c *simpleLFU) Name() string {
	return c.name
}

func (c *simpleLFU) Remove(k interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if entry, ok := c.values[k]; ok {
		c.del(entry)
	}
}

func (c *simpleLFU) SetEvictCallBack(cb func(k, v interface{})) {
	c.callback = cb
}

type cacheEntry struct {
	key      interface{}
	value    interface{}
	expireAt time.Time
	freqNode *list.Element
}

type listEntry struct {
	entries map[*cacheEntry]byte
	freq    int
}

func runStat(lfu *simpleLFU) {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-lfu.stopSignal:
			ticker.Stop()
			return
		case <-ticker.C:
			lfu.emitCounter(lfu.name+".size", lfu.Len(), "", nil)
		}
	}
}

func New(upperBounds, lowerBound int, name string, openStat bool,
	emitCounter func(name string, value interface{}, prefix string, tagkv map[string]string)) *simpleLFU {
	c := &simpleLFU{
		name:        name,
		openStat:    openStat,
		emitCounter: emitCounter,
		upperBound:  upperBounds,
		lowerBound:  lowerBound,
		values:      make(map[interface{}]*cacheEntry),
		freqs:       list.New(),
		stopSignal:  make(chan struct{}),
	}

	if openStat && emitCounter != nil {
		go runStat(c)
		runtime.SetFinalizer(c, stopStat)
	}
	return c
}

func stopStat(lru *simpleLFU) {
	lru.stopSignal <- struct{}{}
}

func (c *simpleLFU) Get(key interface{}) (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	v, exist := func() (interface{}, bool) {
		if e, ok := c.values[key]; ok {
			if time.Now().After(e.expireAt) {
				c.del(e)
				return nil, false
			}
			c.increment(e)
			return e.value, true
		}
		return nil, false
	}()
	if c.openStat && c.emitCounter != nil {
		if exist {
			c.emitCounter(c.name+".hit", 1, "", nil)
		} else {
			c.emitCounter(c.name+".miss", 1, "", nil)
		}
	}
	return v, exist
}

func (c *simpleLFU) del(entry *cacheEntry) {
	delete(c.values, entry.key)
	c.remEntry(entry.freqNode, entry)
	c.len--
	if c.callback != nil {
		c.callback(entry.key, entry.value)
	}
}

func (c *simpleLFU) Set(key, value interface{}, ttl time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if e, ok := c.values[key]; ok {
		// value already exists for key.  overwrite
		e.value = value
		e.expireAt = time.Now().Add(ttl)
		c.increment(e)
	} else {
		// value doesn't exist.  insert
		e := new(cacheEntry)
		e.key = key
		e.value = value
		e.expireAt = time.Now().Add(ttl)
		c.values[key] = e
		c.increment(e)
		c.len++
		// bounds mgmt
		if c.upperBound > 0 && c.lowerBound > 0 {
			if c.len > c.upperBound {
				c.evict(c.len - c.lowerBound)
			}
		}
	}
}

func (c *simpleLFU) Len() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.len
}

func (c *simpleLFU) Evict(count int) int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return c.evict(count)
}

func (c *simpleLFU) evict(count int) int {
	// No lock here so it can be called
	// from within the lock (during Set)
	var evicted int
	for i := 0; i < count; {
		if place := c.freqs.Front(); place != nil {
			for entry := range place.Value.(*listEntry).entries {
				if i < count {
					c.del(entry)
					evicted++
					i++
				}
			}
		}
	}
	return evicted
}

func (c *simpleLFU) increment(e *cacheEntry) {
	currentPlace := e.freqNode
	var nextFreq int
	var nextPlace *list.Element
	if currentPlace == nil {
		// new entry
		nextFreq = 1
		nextPlace = c.freqs.Front()
	} else {
		// move up
		nextFreq = currentPlace.Value.(*listEntry).freq + 1
		nextPlace = currentPlace.Next()
	}

	if nextPlace == nil || nextPlace.Value.(*listEntry).freq != nextFreq {
		// create a new list entry
		li := new(listEntry)
		li.freq = nextFreq
		li.entries = make(map[*cacheEntry]byte)
		if currentPlace != nil {
			nextPlace = c.freqs.InsertAfter(li, currentPlace)
		} else {
			nextPlace = c.freqs.PushFront(li)
		}
	}
	e.freqNode = nextPlace
	nextPlace.Value.(*listEntry).entries[e] = 1
	if currentPlace != nil {
		// remove from current position
		c.remEntry(currentPlace, e)
	}
}

func (c *simpleLFU) remEntry(place *list.Element, entry *cacheEntry) {
	entries := place.Value.(*listEntry).entries
	delete(entries, entry)
	if len(entries) == 0 {
		c.freqs.Remove(place)
	}
}
