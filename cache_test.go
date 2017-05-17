package httpcache

import (
	"testing"
	"httpcache-go/mocks"
	"github.com/golang/mock/gomock"
)

//go:generate go get github.com/golang/mock/gomock
//go:generate go get github.com/golang/mock/mockgen
//go:generate mockgen -destination mocks/mock_http.go net/http ResponseWriter
func Test_Put(t *testing.T) {
	//given
	ctrl := gomock.NewController(t)
	rwMock := mock_http.NewMockResponseWriter(ctrl)

	limit:= 100
	cache := NewLruCache(limit)

	cacheableRW := cacheableResponseWriter{}
	//expect

	//when
	cache.Put()

	//then
}
