package v3

import (
	"fmt"
	"gomessage/apps/models"
)

type ClientActionEmail struct{}

func (c *ClientActionEmail) RendersMessages(client *models.Client, isMerge bool, contentList []string) []any {
	fmt.Println("Email消息体成功...")
	return nil
}

func (c *ClientActionEmail) PushMessages(messages []any) {
	fmt.Println("Email消息体成功...")
}
