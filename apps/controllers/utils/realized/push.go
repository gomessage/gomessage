package realized

import (
	"fmt"
	"gomessage/apps/controllers/send"
	"gomessage/apps/controllers/utils/base"
)

type GeneralPush struct {
	base.Push
}

func (d *GeneralPush) PushData(url string, data []any) {
	fmt.Println("调用钉钉客户端的推送数据的方法...")
	for _, d := range data {
		send.Push(d, url)
	}
}
