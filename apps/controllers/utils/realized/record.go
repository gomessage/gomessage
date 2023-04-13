package realized

import (
	"fmt"
	"gomessage/apps/controllers/utils/base"
)

type GeneralRecord struct {
	base.Record
}

// RecordData 实现Renders接口的方法，对通用内容体进行渲染
func (g *GeneralRecord) RecordData() {
	fmt.Println("记录一条发送行为...")
}
