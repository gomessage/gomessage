package es

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
        "type": "text"
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

const Mapping2 = `
{
  "mappings": {
	"dynamic": true,
    "properties": {}
  }
}`
