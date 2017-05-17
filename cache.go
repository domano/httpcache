package httpcache

import (
	"sync"
	"github.com/hashicorp/golang-lru/simplelru"
)

type Cache interface {
	Put(value CachableResponseWriter)
	Pull(hash string) CachableResponseWriter
}

type lruCache struct {
	sync.Mutex
	limit int
	size int
	newest CacheElement
	oldest CacheElement
	cache map[string]CacheElement
}

func NewLruCache(limit int) Cache{
	return &lruCache{limit:limit}
}
func (c *lruCache) Put(value CachableResponseWriter) {
	c.Lock()
	defer c.Unlock()
	if c.size + value.Size() > c.limit {
		c.oldest.Next().SetPrevious(nil)
		c.cache[c.oldest.Key()] = nil
	}
	ce := NewCacheElement(value)
	ce.SetPrevious(c.newest)
	c.newest.SetNext(ce)
	c.cache[ce.Key()]=ce

}

func (c *lruCache) Pull(key string) CachableResponseWriter {
	return c.cache[key]
}