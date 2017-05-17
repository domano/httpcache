package httpcache

import (
	"bytes"
	"strconv"
	"net/http"
)

type CachableResponseWriter interface {
	http.ResponseWriter
	Key() string
	Size() int
}

type cacheableResponseWriter struct {
	http.ResponseWriter
	key string
	header int
	buffer *bytes.Buffer
}

func NewCachableResponseWriter(key string, rw http.ResponseWriter) CachableResponseWriter{
	return &cacheableResponseWriter{rw, key, 200, &bytes.Buffer{}}
}

func (hc *cacheableResponseWriter) Key() string {
	return strconv.Itoa(hc.buffer.Len())
}

func (hc *cacheableResponseWriter) Size() int {
	return hc.buffer.Len()
}

func (hc *cacheableResponseWriter) WriteHeader(header int) {
	hc.header = header
}

func (hc *cacheableResponseWriter) Write(p []byte) (int, error) {
	hc.buffer.Write(p)
}
