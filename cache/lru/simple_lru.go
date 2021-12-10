package lru

import (
	"container/list"
	"time"
)

type simpleLRU struct {
	size     int
	lru      *list.List
	items    map[interface{}]*list.Element
	callback func(k, v interface{})
}

// entry in the simpleLRU.
type entry struct {
	key       interface{}
	value     interface{}
	expiresAt time.Time
}

func (c *simpleLRU) Add(key, value interface{}, expiresAt time.Time) {
	if ent, ok := c.items[key]; ok {
		// update existing entry
		c.lru.MoveToFront(ent)
		v := ent.Value.(*entry)
		v.value = value
		v.expiresAt = expiresAt
		return
	}

	// add new entry
	c.items[key] = c.lru.PushFront(&entry{
		key:       key,
		value:     value,
		expiresAt: expiresAt,
	})

	// remove oldest
	if c.lru.Len() > c.size {
		ent := c.lru.Back()
		if ent != nil {
			c.removeElement(ent)
		}
	}
}

func (c *simpleLRU) Get(key interface{}) (interface{}, bool) {
	if ent, ok := c.items[key]; ok {
		v := ent.Value.(*entry)

		if v.expiresAt.After(time.Now()) {
			// found good entry
			c.lru.MoveToFront(ent)
			return v.value, true
		}

		// ttl expired
		c.removeElement(ent)
	}
	return nil, false
}

func (c *simpleLRU) Remove(key interface{}) {
	if ent, ok := c.items[key]; ok {
		c.removeElement(ent)
	}
}

func (c *simpleLRU) Len() int {
	return c.lru.Len()
}

// removeElement is used to remove a given list element from the simpleLRU
func (c *simpleLRU) removeElement(e *list.Element) {
	c.lru.Remove(e)
	kv := e.Value.(*entry)
	delete(c.items, kv.key)
	if c.callback != nil {
		c.callback(kv.key, kv.value)
	}
}
