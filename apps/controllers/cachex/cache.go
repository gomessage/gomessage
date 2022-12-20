package cachex

import (
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

var cc *cache.Cache
var once sync.Once

func init() {
	once.Do(func() {
		cc = cache.New(5*time.Minute, 10*time.Minute)
	})
}

type arbitrarilyJsonData struct {
	RequestJson map[string]any `json:"request_json"`
	Time        time.Time      `json:"time"`
	RequestByte []byte
}

var CacheData arbitrarilyJsonData

func SetCacheData(ns string, value any) {
	key := ns + "_key"
	cc.Set(key, value, cache.NoExpiration)
}

func GetCacheData(ns string) (any, bool) {
	key := ns + "_key"
	return cc.Get(key)
}
