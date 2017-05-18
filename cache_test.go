package httpcache

import (
	"github.com/golang/mock/gomock"
	"testing"
	"github.com/stretchr/testify/assert"
	"sort"
)

//go:generate go get github.com/golang/mock/gomock
//go:generate go get github.com/golang/mock/mockgen
//go:generate mockgen -destination mocks/http/mock_http.go net/http ResponseWriter
//go:generate mockgen -destination mocks/httpcache/mock_cache.go httpcache CacheableResponseWriter

type testAddData struct {
	testCase string
	hashes   []string
	expected []string
	size     int
	limit    int
}

var addTestCases = []testAddData{
	{"single entry, fitting", []string{"abc"}, []string{"abc"}, 100, 100},
	{"multiple entries, fitting", []string{"a", "b", "c"}, []string{"a", "b", "c"}, 33, 99},
	{"single entry, too large", []string{"a"}, []string{}, 101, 100},
	{"multiple entries, sum too large", []string{"a", "b"}, []string{"b"}, 50, 50},
}

// Data driven test for updates
func Test_Add(t *testing.T) {
	for _, test := range addTestCases {
		t.Run(test.testCase, func(t *testing.T) {
			addTestFunction(t, test)
		})
	}
}

func addTestFunction(t *testing.T, data testAddData) {
	//given
	ctrl := gomock.NewController(t)
	cache := NewLruCache(data.limit)
	crwMocks := []Cacheable{}


	//expect

	//when
	for
	cache.Add()

	//then
	for _, expected := range data.expected {
		assert.NotNil(t,cache.(lruCache).cache[expected])
	}

	for _, hash := range data.hashes {

		if data.expected[hash] == "" {

		}
	}
}
