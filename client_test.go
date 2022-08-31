package httpmock_bench

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var OKHandler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprint(writer, "OK")
})

func Benchmark_httptest(b *testing.B) {
	for n := 0; n < b.N; n++ {
		httptestIt(b)
	}
}

func httptestIt(b *testing.B) {
	s := httptest.NewServer(OKHandler)

	defer s.Close()

	res, err := get(s.URL)

	assert.Nil(b, err)
	assert.Equal(b, "OK", res)
}

func Benchmark_httpmock(b *testing.B) {
	for n := 0; n < b.N; n++ {
		httpmockIt(b)
	}
}

func httpmockIt(b *testing.B) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "http://httpmock",
		httpmock.NewStringResponder(200, "OK"))

	res, err := get("")

	assert.Nil(b, err)
	assert.Equal(b, "OK", res)
}
