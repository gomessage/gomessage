package es

import (
	"context"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"time"
)

type PushToEsHook struct {
	client *elastic.Client
}

func (esHook *PushToEsHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (esHook *PushToEsHook) Fire(entry *logrus.Entry) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	SyncIndex(JoinIndexName("gomessage-push"))
	_, err := esHook.client.Index().Index(JoinIndexName("gomessage-push")).BodyJson(*CreateRecord(entry)).Do(ctx)
	return err
}

func NewPushToEsHook() *PushToEsHook {
	return &PushToEsHook{
		client: GetEsClient(),
	}
}
