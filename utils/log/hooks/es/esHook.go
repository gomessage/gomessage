package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"time"
)

// excludedKeys is the keys which should not be sent to ElasticSearch
//var excludedKeys = []string{"SendToElasticsearch"}

type Hook struct {
	client *elastic.Client
	host   string
	index  string
	levels []logrus.Level
}

func (esHook *Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (esHook *Hook) Fire(entry *logrus.Entry) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	entry.Data["server_host"] = esHook.host
	entry.Data["levels"] = esHook.levels
	_, err := esHook.client.Index().Index(esHook.index).BodyJson(entry.Data).Do(ctx)
	return err
}

func NewEsHook(levels []logrus.Level) *Hook {
	return &Hook{
		client: ElasticsearchClient,
		host:   MyHost,
		index:  IndexName,
		levels: levels,
	}
}
