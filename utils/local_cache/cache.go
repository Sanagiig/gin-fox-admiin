package geeCache

import (
	"gin-one/utils/local_cache/internal"
	"sync"
)

type Cache struct {
	mu         sync.RWMutex
	lru        *internal.LRUCache
	cacheBytes int64
}

func NewCache(maxbytes int64) *Cache {
	return &Cache{
		mu:         sync.RWMutex{},
		cacheBytes: maxbytes,
	}
}

func (c *Cache) Add(key string, value internal.ByteView) {
	c.mu.RLock()
	if c.lru == nil {
		c.lru = internal.NewLRU(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
	c.mu.RUnlock()
}

func (c *Cache) Get(key string) (value internal.ByteView, ok bool) {
	var v internal.Value
	c.mu.RLock()
	if c.lru != nil {
		v, ok = c.lru.Get(key)
		if ok {
			value = v.(internal.ByteView)
		}
	}
	c.mu.RUnlock()
	return value, ok
}
