package httpcache

import (
	"sync"
	"net/http"
	"errors"
	"fmt"
)

type Cache interface {
	Add(value Cacheable) error
	Get(hash string) Cacheable
	Hash(r *http.Request) string
}

type Hashable interface {
}

type lruCache struct {
	sync.RWMutex
	limit int
	size int
	newest CacheElement
	oldest CacheElement
	cache map[string]CacheElement
}

func NewLruCache(limit int) Cache{
	return &lruCache{limit:limit, cache:map[string]CacheElement{}}
}
func (c *lruCache) Add(value Cacheable) error{
	if  value.Size()>c.limit {
		return errors.New(fmt.Sprintf("%s is bigger then cache limit", value.Key()))
	}
	ce := NewCacheElement(value)
	c.Lock()
	if c.size + value.Size() > c.limit {
		i := 0
		for i < ce.Size() {
			delete(c.cache, c.oldest.Key())
			c.size-=c.oldest.Size()
			i+=c.oldest.Size()
			if c.oldest.Next() != nil {
				newOld := c.oldest.Next().SetPrevious(nil)
				c.oldest = newOld
				continue
			}
			c.oldest=ce

		}
	}
	if c.newest == nil {
		c.newest = ce
	} else {
		ce.SetPrevious(c.newest)
		c.newest.SetNext(ce)
		c.newest = ce
	}
	if c.oldest == nil {
		c.oldest = ce
	}
	c.cache[ce.Key()]=ce
	c.size+=ce.Size()
	c.Unlock()
	return nil

}

func (c *lruCache) Get(key string) Cacheable {
	c.RLock()
	defer c.RUnlock()
	return c.cache[key]
}

func (c *lruCache) Hash(r *http.Request) string {
	return r.URL.RawPath
}