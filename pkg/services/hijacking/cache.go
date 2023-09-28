package hijacking

import (
	"gomessage/pkg/utils/log/loggers"
	"time"
)

//var cc *cache.Cache
//var once sync.Once

var DataList map[string]any

func init() {
	//once.Do(func() {
	//	cc = cache.New(5*time.Minute, 5*time.Minute)
	//})
	DataList = make(map[string]any)
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

// SetCacheData TODO: 由于缓存库出现了BUG，因此这里临时用map顶一下，以后在修复，上下文中与cache相关的内容临时注释一下
func SetCacheData(ns string, value any) {
	if len(ns) != 0 {
		key := ns + "_key"
		//cc.Set(key, value, cache.DefaultExpiration)
		DataList[key] = value
	} else {
		loggers.DefaultLogger.Errorln("写入cache时，Namespace获取失败...")
	}

}

func GetCacheData(ns string) (any, bool) {
	if len(ns) != 0 {
		key := ns + "_key"
		//return cc.Get(key)
		return DataList[key], true
	} else {
		return nil, false
	}

}
