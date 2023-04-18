package hijacking

import (
	"github.com/patrickmn/go-cache"
	"sync"
	"time"
)

var cc *cache.Cache
var once sync.Once

func init() {
	once.Do(func() {
		cc = cache.New(5*time.Minute, 5*time.Minute)
	})
}

type Mappings struct {
	Key   string `json:"key"`
	Value any    `json:"value"`
}

type arbitrarilyJsonData struct {
	RequestJson  map[string]any `json:"request_json"`
	RequestTime  time.Time      `json:"request_time"`
	RequestByte  []byte         `json:"-"`
	KeyValueMap  map[string]any `json:"-"`
	KeyValueList []Mappings     `json:"key_value_list"`
}

var CacheData arbitrarilyJsonData

func SetCacheData(ns string, value any) {
	key := ns + "_key"
	cc.Set(key, value, cache.DefaultExpiration)
}

func GetCacheData(ns string) (any, bool) {
	key := ns + "_key"
	return cc.Get(key)
}
