package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"time"
)

type AccessToEsHook struct {
	client *elastic.Client
}

func (esHook *AccessToEsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (esHook *AccessToEsHook) Fire(entry *logrus.Entry) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	SyncIndex(JoinIndexName("gomessage-access"))
	_, err := esHook.client.Index().Index(JoinIndexName("gomessage-access")).BodyJson(*CreateRecord(entry)).Do(ctx)
	return err
}

func NewAccessToEsHook() *AccessToEsHook {
	return &AccessToEsHook{
		client: GetEsClient(),
	}
}
