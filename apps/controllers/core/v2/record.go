package v2

import (
	"fmt"
)

type GeneralRecord struct {
	Record
}

// RecordData 实现Renders接口的方法，对通用内容体进行渲染
func (g *GeneralRecord) RecordData() {
	fmt.Println("记录一条发送行为...")
}
