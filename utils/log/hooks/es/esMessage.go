package es

// Mapping 定义索引
const Mapping = `
{
  "mappings": {
    "properties": {
	  "@timestamp": {
        "type": "date"
      },
      "start_time": {
        "type": "keyword"
      },
      "end_time": {
        "type": "keyword"
      },
	  "latency": {
        "type": "keyword"
      },
      "status_code": {
        "type": "integer"
      },
      "client_ip": {
        "type": "keyword"
      },
      "method": {
        "type": "keyword"
      },
      "path": {
        "type": "keyword"
      },
	  "router": {
        "type": "keyword"
      },
	  "request_body": {
        "type": "keyword"
      },
	  "server_host": {
        "type": "keyword"
      },
	  "levels": {
        "type": "keyword"
      }
    }
  }
}`

//type DocsToEs struct {
//	Timestamp   time.Time `json:"@timestamp"`
//	StartTime   string    `json:"start_time"`
//	EndTime     string    `json:"end_time"`
//	Latency     string    `json:"latency"`
//	StatusCode  int       `json:"status_code"`
//	ClientIp    string    `json:"client_ip"`
//	Method      string    `json:"method"`
//	Path        string    `json:"path"`
//	Router      string    `json:"router"`
//	RequestBody string    `json:"request_body"`
//}
