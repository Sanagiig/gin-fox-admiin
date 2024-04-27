package internal

import (
	"container/list"
	"fmt"
	"strings"
)

type OnEvicted = func(string, Value)

type LRUCache struct {
	maxBytes  int64
	nbytes    int64
	ll        *list.List
	cache     map[string]*list.Element
	OnEvicted OnEvicted
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func NewLRU(maxBytes int64, onEvicted OnEvicted) *LRUCache {
	return &LRUCache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *LRUCache) Data() string {
	sb := strings.Builder{}
	sb.WriteString("[")
	for _, ele := range c.cache {
		val := ele.Value.(*entry)
		sb.WriteString(fmt.Sprintf(" {%s : %s} ", val.key, val.value.(ByteView).String()))
	}
	sb.WriteString("]")
	return sb.String()
}

func (c *LRUCache) Len() int64 {
	return c.nbytes
}

func (c *LRUCache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}

	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.Removeoldest()
	}
}

func (c *LRUCache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

func (c *LRUCache) Removeoldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}
