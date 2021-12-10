package cache

import "time"

type Cache interface {
	Name() string
	Len() int
	Get(k interface{}) (interface{}, bool)
	Set(k, v interface{}, ttl time.Duration)
	Remove(k interface{})
	SetEvictCallBack(func(k, v interface{}))
}
