package simplettl

import (
	"sync"
	"time"
)

/*
简单的ttl cache，没有淘汰逐出功能，不设置尺寸，适合类似于配置文件，etcd这种场景，高效，O(1) 复杂度
*/

type entry struct {
	value     interface{}
	expiresAt time.Time
}

type simpleTTL struct {
	sync.RWMutex
	name     string
	content  map[interface{}]*entry
	callBack func(k, v interface{})
}

func (s *simpleTTL) SetEvictCallBack(cb func(k, v interface{})) {
	s.callBack = cb
}

func (s *simpleTTL) Name() string {
	return s.name
}

func (s *simpleTTL) Len() int {
	s.RLock()
	l := len(s.content)
	s.RUnlock()
	return l
}

func (s *simpleTTL) Get(k interface{}) (interface{}, bool) {
	s.RLock()
	defer s.RUnlock()
	if v, ok := s.content[k]; ok {
		if time.Now().After(v.expiresAt) {
			return nil, false
		}
		return v.value, true
	}
	return nil, false
}

func (s *simpleTTL) Set(k, v interface{}, ttl time.Duration) {
	s.Lock()
	s.content[k] = &entry{
		v,
		time.Now().Add(ttl),
	}
	s.Unlock()
}

func (s *simpleTTL) Remove(k interface{}) {
	var v interface{}
	var ok bool
	if s.callBack != nil {
		v, ok = s.Get(k)
	}
	s.Lock()
	delete(s.content, k)
	s.Unlock()
	if s.callBack != nil && ok {
		s.callBack(k, v)
	}
}

func NewSimpleTTLCache(name string) *simpleTTL {

	return &simpleTTL{
		name:    name,
		content: make(map[interface{}]*entry),
	}

}
