package httpcache

type CacheElement interface {
	CachableResponseWriter
	SetNext(CacheElement) CacheElement
	Next() CacheElement

	SetPrevious(CacheElement) CacheElement
	Previous() CacheElement

}

type cacheElement struct {
	CachableResponseWriter
	next CacheElement
	previous CacheElement
}

func NewCacheElement(c CachableResponseWriter) CacheElement{
	return &cacheElement{c, nil, nil}
}

func (ce *cacheElement) SetNext(next CacheElement) CacheElement {
	ce.next = next
	return ce
}

func (ce *cacheElement) Next() CacheElement {
	return ce.next
}

func (ce *cacheElement) SetPrevious(previous CacheElement) CacheElement{
	ce.previous = previous
	return ce
}

func (ce *cacheElement) Previous() CacheElement{
	return ce.previous
}


