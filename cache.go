package httpcache

import (
	"sync"
	"net/http"
)

type Cache interface {
	Add(value Cacheable) error
	Get(hash string) Cacheable
	Hash(r *http.Request) string
}

type Hashable interface {
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
func (c *lruCache) Add(value Cacheable) error{
	if c.size + value.Size() > c.limit {
		c.Lock()
		c.oldest.Next().SetPrevious(nil)
		delete(c.cache, c.oldest.Key())
		c.Unlock()
	}
	ce := NewCacheElement(value)
	c.Lock()
	ce.SetPrevious(c.newest)
	c.newest.SetNext(ce)
	c.cache[ce.Key()]=ce
	c.Unlock()
	return nil

}

func (c *lruCache) Get(key string) Cacheable {
	return c.cache[key]
}

func (c *lruCache) Hash(r *http.Request) string {
	return r.URL.RawPath
}