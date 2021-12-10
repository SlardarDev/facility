package cache

import (
	"github.com/SlardarDev/facility/cache/lfu"
	"github.com/SlardarDev/facility/cache/lru"
	"github.com/SlardarDev/facility/cache/simplettl"
)

type MetricsCounterFunc func(name string, value interface{}, prefix string, tagkv map[string]string)

func NewSimpleTTL(name string) Cache {
	return simplettl.NewSimpleTTLCache(name)
}

func NewLRU(size int, name string, openStat bool, emitCounter MetricsCounterFunc) Cache {
	return lru.NewLockedLRU(size, name, openStat, emitCounter)
}

func NewLFU(maxSize int, evictToSize int, name string, openStat bool, emitCounter MetricsCounterFunc) Cache {
	return lfu.New(maxSize, evictToSize, name, openStat, emitCounter)
}
