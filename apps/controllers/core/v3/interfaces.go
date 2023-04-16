package v3

import "gomessage/apps/models"

type ClientAction interface {
	RendersMessages(client *models.Client, isMerge bool, contentList []string) []any
	PushMessages(messages []any)
}
