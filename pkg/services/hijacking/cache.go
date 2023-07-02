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
	RequestByte  []byte         `json:"-"`              //请求数据[]byte格式
	RequestJson  map[string]any `json:"request_json"`   //请求数据json格式
	RequestTime  time.Time      `json:"request_time"`   //请求时间
	KeyValueMap  map[string]any `json:"key_value_map"`  //展开解析后的key:value
	KeyValueList []Mappings     `json:"key_value_list"` //展开解析后的key:value切片
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
