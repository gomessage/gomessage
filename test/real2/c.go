package real2

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retrievers struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r2 Retrievers) Get(cur string) string {
	re, err := http.Get(cur)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(re, true)
	if err != nil {
		panic(err)
	}
	return string(result)
}
